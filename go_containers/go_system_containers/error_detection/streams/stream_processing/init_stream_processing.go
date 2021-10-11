package monitor_streams


import "fmt"
//import "time"
import "strings"
//import "encoding/json"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/server_libraries/postgres"
//import "lacima.com/Patterns/logging_support"
import "lacima.com/Patterns/msgpack_2"
import "github.com/vmihailenco/msgpack/v5"


type Working_Structure_Type struct {
    count int64
    status bool
}


type stream_records_type struct {
    name                    string
    namespace               string
    description             string
    key_array               []string
    key                     string
    stream_keys             []string
    
}

type monitor_stream_type struct {
    
    sample_time                      int64
    current_time                     int64
    description                      map[string]string
    name                             map[string]string
    key                              map[string]string
    working_table                    redis_handlers.Redis_Hash_Struct
    time_table                       redis_handlers.Redis_Hash_Struct
    status_table                     redis_handlers.Redis_Hash_Struct
    error_value                      redis_handlers.Redis_Hash_Struct
    error_time                       redis_handlers.Redis_Hash_Struct
    process_data_stream              pg_drv.Postgres_Stream_Driver
    process_incident_stream          pg_drv.Postgres_Stream_Driver
    input_stream_map                 map[string]redis_handlers.Redis_Stream_Struct

}



var default_working_structure  Working_Structure_Type 
var monitor_control            monitor_stream_type 







func Init_data_structures(){
   
    
    initialize_default_structure()
    construct_monitor_control()
    construct_construct_stream_records()
    populate_data_structures()
    
    
}


func initialize_default_structure(){
    default_working_structure.count = 0
    default_working_structure.status = true
    
}

func construct_monitor_control() {
   
    wd_nodes  := []string{"ERROR_DETECTION:ERROR_DETECTION", "STREAMING_LOGS:STREAMING_LOGS"   }
    nodes := graph_query.Common_qs_search(&wd_nodes)
    node  := nodes[0]
    
    monitor_control.sample_time     = graph_query.Convert_json_int(node["sample_time"])
    fmt.Println("sample time",monitor_control.sample_time)
    
    search_list := []string{"ERROR_DETECTION:ERROR_DETECTION", "STREAMING_LOGS:STREAMING_LOGS" ,"STREAM_SUMMARY_DATA"}
    data_element := data_handler.Construct_Data_Structures(&search_list)
	 	 
	monitor_control.working_table               =   (*data_element)["WORKING_TABLE"].(redis_handlers.Redis_Hash_Struct) 
    monitor_control.time_table                  =   (*data_element)["TIME_TABLE"].(redis_handlers.Redis_Hash_Struct) 
	monitor_control.status_table                =   (*data_element)["STATUS_TABLE"].(redis_handlers.Redis_Hash_Struct) 
	monitor_control.error_value                 =   (*data_element)["ERROR_TABLE"].(redis_handlers.Redis_Hash_Struct) 
	monitor_control.error_time                  =   (*data_element)["ERROR_TIME"].(redis_handlers.Redis_Hash_Struct) 
	monitor_control.process_data_stream         =   (*data_element)["DATA_STREAM"].(pg_drv.Postgres_Stream_Driver) 
	monitor_control.process_incident_stream     =   (*data_element)["INCIDENT_STREAM"].(pg_drv.Postgres_Stream_Driver) 
	
  
    
}
    
func construct_construct_stream_records(){
    
    
            
    monitor_control.description         = make(map[string]string)
    monitor_control.name                = make(map[string]string)
    monitor_control.key                 = make(map[string]string)
    monitor_control.input_stream_map    = make(map[string]redis_handlers.Redis_Stream_Struct)
    
    stream_nodes  := []string{"STREAMING_LOG"}
    nodes := graph_query.Common_qs_search(&stream_nodes)
    //fmt.Println("nodes",len(nodes),"\n",nodes)
    
    for _,node := range nodes {
        var item  stream_records_type
        item.namespace    = graph_query.Convert_json_string(node["namespace"])
        item.name         = graph_query.Convert_json_string(node["name"])
        item.description  = graph_query.Convert_json_string(node["descrption"])
        item.stream_keys  = graph_query.Convert_json_string_array(node["keys"])
        item.key_array    = graph_query.Generate_key(item.namespace)
        construct_stream_data_structures(item)
        
    }
    
    
    
    
}

func construct_stream_data_structures( item  stream_records_type){
    
 
    
        base_key_array := graph_query.Generate_key(item.namespace)
        base_key       := strings.Join(base_key_array,"/")
        key_temp               := append(item.key_array,"STREAMING_LOG")
        handlers                := data_handler.Construct_Data_Structures(&key_temp)
        for _,key := range item.stream_keys {
            stream_key                         := base_key+"/"+key
            
            monitor_control.key[stream_key]              = key
            monitor_control.description[stream_key]      = item.description
            monitor_control.name[stream_key]             = item.name
            monitor_control.input_stream_map[stream_key] = (*handlers)[key].(redis_handlers.Redis_Stream_Struct)
            
        }
        
    


    
}




func populate_data_structures(){
    
    default_time    := msg_pack_utils.Pack_int64(0)
    default_working := Pack_working_structure(default_working_structure)
    default_string  := msg_pack_utils.Pack_string("")
    default_bool    := msg_pack_utils.Pack_bool(true)
    
    clean_redis_hash_table( monitor_control.key, monitor_control.working_table, default_working)
    clean_redis_hash_table( monitor_control.key, monitor_control.time_table, default_time )
    clean_redis_hash_table( monitor_control.key, monitor_control.status_table, default_bool )
    clean_redis_hash_table( monitor_control.key, monitor_control.error_value, default_string )
    clean_redis_hash_table( monitor_control.key, monitor_control.error_time, default_time)
    
    
}
 
func clean_redis_hash_table( key_list map[string]string, table redis_handlers.Redis_Hash_Struct, default_value string){
    
   values := table.HGetAll()
   table.Delete_All()

   for key,value := range values {
       if _,ok := key_list[value]; ok== true {
           table.HSet(key,value)
       }
   }
   for key, _ := range key_list {
      if ok := table.HExists(key); ok == false {
          table.HSet(key,default_value)
      }
       
   }
    
}
    
    
    
func Pack_working_structure( working_structure Working_Structure_Type )string{
    
    b, err := msgpack.Marshal(&working_structure)
    if err != nil {
        panic(err)
    }
    return string(b)
}

    
    
func Unpack_working_structure( input string) (Working_Structure_Type,bool){
    
    item := default_working_structure
    state := true
    err := msgpack.Unmarshal([]byte(input), &item)
    if err != nil {
        state = false
    } 
    
    return item,state
    
}
    
    
    
/*


func construct_construct_rpc_servers(){
    rpc_records = make([]rpc_records_type,0)
    incident_nodes  := []string{"RPC_SERVER"}
    nodes := graph_query.Common_qs_search(&incident_nodes)
    fmt.Println("nodes",len(nodes),nodes)
    
    for _,node := range nodes{
        var item  rpc_records_type
        item.rpc_stream_array   = make( map[string]redis_handlers.Redis_Stream_Struct)
        
        item.namespace          = graph_query.Convert_json_string(node["namespace"])
        item.key_array          = graph_query.Generate_key(item.namespace)
        key_array               := append(item.key_array,"RPC_SERVER")
        item.key                = strings.Join(item.key_array,"/")
        
        handlers                := data_handler.Construct_Data_Structures(&key_array)
        item.rpc_server         = (*handlers)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
       
        key_array               = append(item.key_array,"STREAMING_LOG")
        handlers                = data_handler.Construct_Data_Structures(&key_array)
       

        item.rpc_stream_array["number"]     = (*handlers)["number"].(redis_handlers.Redis_Stream_Struct)
        item.rpc_stream_array["queue_depth"] = (*handlers)["queue_depth"].(redis_handlers.Redis_Stream_Struct)
        item.rpc_stream_array["utilization"] = (*handlers)["utilization"].(redis_handlers.Redis_Stream_Struct)

        
        
        rpc_records             = append(rpc_records,item)
    }
    
}


func Process_functions(){
        
   go ping_rpc_server_loop()
  
}






func ping_rpc_server_loop(){
    timeout   := time.Duration(monitor_control.sample_time)*time.Minute
    
    rpc_bad_number = 0
    for true {
      fmt.Println("ping rpc server")
      ping_rpc_servers()
      time.Sleep(timeout)
      
    }
}


func ping_rpc_servers(){
   for _,rpc_record := range rpc_records {
       fmt.Println("key",rpc_record.key)
       ping_rpc_server( rpc_record )
   }
   if rpc_bad_number > 0 {
       post_incident_report()
   }
   
}       
    
 func post_incident_report(){
    
    request_json,err := json.Marshal(&rpc_state)
    if err != nil{
          panic("json marshall error")
    }  
    fmt.Println("request_json",string(request_json))
    monitor_control.incident_log.Log_data(string(request_json))
}


  
func ping_rpc_server(item  rpc_records_type ){
   key       := item.namespace
   rpc_state[key] = false        
   
   
   parameters := make(map[string]interface{})
   result := item.rpc_server.Send_json_rpc_message( "info", parameters ) 
   if result == nil {
       rpc_bad_number += 1
       
       fmt.Println("rpc not active")
   }else{
      fmt.Println("result", result )
      if result["status"].(bool) == false {
          rpc_bad_number += 1
                    
      }else{
        rpc_state[key]         = true 
        number                 :=int64( result["number"].(float64))
        length                 :=int64( result["length"].(float64))
        delta_time             :=(  result["end_time"].(float64) -  result["start_time"].(float64))
        time_utilitization      := result["total_time"].(float64)/delta_time
        post_data_to_stream(item, number,length,time_utilitization)
      }
   }
    
}

 

func post_data_to_stream(item  rpc_records_type,   number,length int64, time_utilitization float64){
    
    number_packed  := msg_pack_utils.Pack_int64(number)
    item.rpc_stream_array["number"].Xadd( number_packed )
    
    length_packed  := msg_pack_utils.Pack_int64(length)
    item.rpc_stream_array["queue_depth"].Xadd(length_packed)

    time_utilitization_packed := msg_pack_utils.Pack_float64(time_utilitization)
    item.rpc_stream_array["utilization"].Xadd(time_utilitization_packed)   
  
}




*/
