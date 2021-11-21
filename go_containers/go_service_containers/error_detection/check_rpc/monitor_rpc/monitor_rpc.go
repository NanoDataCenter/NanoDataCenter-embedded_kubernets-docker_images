package monitor_rpc


import "fmt"
import "time"
import "strings"
import "encoding/json"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/Patterns/logging_support"
import "lacima.com/Patterns/msgpack_2"
import "lacima.com/server_libraries/postgres"

 
type monitor_rpc_type struct {
    stream_keys                      []string
    sample_time                      int64
    trim_time                        int64
    keys                             map[string]rpc_records_type 
    description                      redis_handlers.Redis_Hash_Struct
    contact_time                     redis_handlers.Redis_Hash_Struct
    status                           redis_handlers.Redis_Hash_Struct
    loading                          redis_handlers.Redis_Hash_Struct
    length                           redis_handlers.Redis_Hash_Struct

    incident_log                     pg_drv.Postgres_Stream_Driver
    trim_handle                      pg_drv.Postgres_Stream_Driver
    
    
    rpc_incident_log                 *logging_support.Incident_Log_Type
    load_incident_log                *logging_support.Incident_Log_Type
    performance_log                   pg_drv.Postgres_Stream_Driver
}

type rpc_records_type struct {
    namespace          string
    key_array          []string
    key                string
    description        string
    rpc_server         redis_handlers.Redis_RPC_Struct
    
   
}





var monitor_control                 monitor_rpc_type
var rpc_records                     []rpc_records_type
var rpc_state                       map[string]bool
var rpc_bad_number                  int64


var keys []string



func Init_data_structures(){
   
    rpc_state                          = make(map[string]bool)
    
    construct_monitor_control()
    construct_construct_rpc_servers()
    construct_keys()
   
    
    
    
}



func construct_monitor_control() {
   
    rpc_nodes  := []string{"ERROR_DETECTION:ERROR_DETECTION", "RPC_ANALYSIS:RPC_ANALYSIS"   }
    nodes := graph_query.Common_qs_search(&rpc_nodes)
    node  := nodes[0]
    monitor_control.stream_keys         = []string{"queue_depth","utilization"}    
    monitor_control.sample_time         = graph_query.Convert_json_int(node["sample_time"])
    monitor_control.trim_time            = graph_query.Convert_json_int(node["trim_time"])
    monitor_control.rpc_incident_log    = logging_support.Construct_incident_log([]string{"ERROR_DETECTION:ERROR_DETECTION", "INCIDENT_LOG:RPC_FAILURE" ,"INCIDENT_LOG"} )
    monitor_control.load_incident_log   = logging_support.Construct_incident_log([]string{"ERROR_DETECTION:ERROR_DETECTION", "INCIDENT_LOG:RPC_LOADING" ,"INCIDENT_LOG"} )
    
    data_node_search  := []string{"ERROR_DETECTION:ERROR_DETECTION", "RPC_ANALYSIS:RPC_ANALYSIS" ,"RPC_ANALYSIS_DATA" }
    handlers := data_handler.Construct_Data_Structures(&data_node_search)
    monitor_control.description              = (*handlers)["DESCRIPTION"].(redis_handlers.Redis_Hash_Struct)
    monitor_control.contact_time             = (*handlers)["TIME"].(redis_handlers.Redis_Hash_Struct)
    monitor_control.status                   = (*handlers)["STATUS"].(redis_handlers.Redis_Hash_Struct)
    monitor_control.loading                  = (*handlers)["LOADING"].(redis_handlers.Redis_Hash_Struct)
    monitor_control.length                   = (*handlers)["LENGTH"].(redis_handlers.Redis_Hash_Struct)
    monitor_control.incident_log             = (*handlers)["INCIDENT_LOG"].(pg_drv.Postgres_Stream_Driver)
    monitor_control.trim_handle              = (*handlers)["INCIDENT_LOG"].(pg_drv.Postgres_Stream_Driver)  
    monitor_control.performance_log          = logging_support.Find_stream_logging_driver()
}



func construct_construct_rpc_servers(){
    rpc_records = make([]rpc_records_type,0)
    incident_nodes  := []string{"RPC_SERVER"}
    nodes := graph_query.Common_qs_search(&incident_nodes)
    //fmt.Println("nodes",len(nodes),nodes)
    
    for _,node := range nodes{
        var item  rpc_records_type
        
        item.description        = graph_query.Convert_json_string(node["description"])
        item.namespace          = graph_query.Convert_json_string(node["namespace"])
        item.key_array          = graph_query.Generate_key(item.namespace)
        key_array               := append(item.key_array,"RPC_SERVER")
        item.key                = strings.Join(item.key_array,"/")
        
        handlers                := data_handler.Construct_Data_Structures(&key_array)
        item.rpc_server         = (*handlers)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
       


        
        
        rpc_records             = append(rpc_records,item)
    }
    
}

func construct_keys(){
    monitor_control.keys = make(map[string] rpc_records_type)
    monitor_control.description.Delete_All()
    monitor_control.contact_time.Delete_All()
    monitor_control.status.Delete_All()
    monitor_control.loading.Delete_All()
    
    zero_pack   := msg_pack_utils.Pack_int64(0)
    status_pack := msg_pack_utils.Pack_bool(true)
    for _,item := range rpc_records{
        
        monitor_control.keys[item.namespace] = item
        monitor_control.description.HSet(item.namespace,msg_pack_utils.Pack_string(item.description))
        
        monitor_control.contact_time.HSet(item.namespace,zero_pack)
        monitor_control.status.HSet(item.namespace,status_pack)           
        monitor_control.loading.HSet(item.namespace,zero_pack)
        
    }
}





func Process_functions(){
        
   go ping_rpc_server_loop()
   go trim_db()
  
}


func trim_db(){
  
    for true {
      
      monitor_control.trim_handle.Trim(monitor_control.trim_time)
      time.Sleep(time.Hour)
      
    }
}




func ping_rpc_server_loop(){
    timeout   := time.Duration(monitor_control.sample_time)*time.Minute
    
    rpc_bad_number = 0
    for true {
      //fmt.Println("ping rpc server")
      ping_rpc_servers()
      time.Sleep(timeout)
      
    }
}


func ping_rpc_servers(){
   for _,rpc_record := range rpc_records {
       //fmt.Println("key",rpc_record.key)
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
    //fmt.Println("request_json",string(request_json))
    monitor_control.rpc_incident_log.Log_data(string(request_json))
}

func post_performance_incident(key string,time_utilitization float64){
    rpc_state := make(map[string]interface{})
    rpc_state["server"] = key
    rpc_state["utilization"] = time_utilitization
    request_json,err := json.Marshal(&rpc_state)
    if err != nil{
          panic("json marshall error")
    }  
    //fmt.Println("request_json",string(request_json))
    monitor_control.load_incident_log.Log_data(string(request_json))
}


func ping_rpc_server(item  rpc_records_type ){
   key       := item.namespace
   rpc_state[key] = false        
   store_contact_time(item)
   //fmt.Println("rpc server ",key)
   parameters := make(map[string]interface{})
   result := item.rpc_server.Send_json_rpc_message( "info", parameters ) 
   if result == nil {
       rpc_bad_number += 1
       store_bad_status(item)
       fmt.Println("rpc not active")
   }else{
      //fmt.Println("result", result )
      if result["status"].(bool) == false {
          rpc_bad_number += 1
          store_bad_status(item)
          fmt.Println("bad",key,rpc_bad_number)
                    
      }else{
        rpc_state[key]         = true 
        
        length                 :=int64( result["length"].(float64))
        delta_time             :=  result["end_time"].(float64) -  result["start_time"].(float64)
        time_utilitization      := result["total_time"].(float64)/delta_time
        if time_utilitization > .5 {
            post_performance_incident(key,time_utilitization)
        }
        post_data_to_stream(item, length,time_utilitization)
        store_good_status(item,time_utilitization,length)
      }
   }
    
}

 

func post_data_to_stream(item  rpc_records_type,length int64, time_utilitization float64){
    
    // store status
    // store loadin    // store contact_tim
    length_packed  := msg_pack_utils.Pack_int64(length)
    monitor_control.performance_log.Insert( "RPC_ANALYSIS",item.key,"queue_depth","","",length_packed )

    time_utilitization_packed := msg_pack_utils.Pack_float64(time_utilitization)
    monitor_control.performance_log.Insert( "RPC_ANALYSIS",item.key,"utilization","","",time_utilitization_packed )   
  
}

func store_contact_time(item rpc_records_type){
    key        := item.namespace
    time_stamp := time.Now().UnixNano()
    time_pack  := msg_pack_utils.Pack_int64(time_stamp)
    monitor_control.contact_time.HSet(key,time_pack)
}
    
    
func store_bad_status(item rpc_records_type){
    //fmt.Println("bad  key ",item.namespace)
    key := item.namespace
    check_for_status_change(item, false)
    monitor_control.status.HSet(key,msg_pack_utils.Pack_bool(false))
    monitor_control.length.HSet(key,msg_pack_utils.Pack_int64(0))
    monitor_control.loading.HSet(key,msg_pack_utils.Pack_float64(0))    
    
    
}

func store_good_status(item rpc_records_type,time_utilitization float64,length int64){
    
    //fmt.Println("good key ",item.namespace,time_utilitization,length)
    key := item.namespace
    check_for_status_change(item, true)
    monitor_control.status.HSet(key,msg_pack_utils.Pack_bool(true))
    monitor_control.length.HSet(key,msg_pack_utils.Pack_int64(length))
    monitor_control.loading.HSet(key,msg_pack_utils.Pack_float64(time_utilitization)) 
}

func check_for_status_change( item rpc_records_type, current_state bool ){
    key := item.namespace
    past_state, err := msg_pack_utils.Unpack_bool(monitor_control.status.HGet(key))
    if err == false {
        past_state = false
    }
    if past_state != current_state {
        //fmt.Println("ps log key ",key,current_state)
        current_state_packed   :=  msg_pack_utils.Pack_bool(current_state)
        monitor_control.incident_log.Insert( key,"","","","",current_state_packed)
    }
    
}
   
