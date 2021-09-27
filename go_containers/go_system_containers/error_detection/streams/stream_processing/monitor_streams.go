package monitor_streams


import "fmt"
//import "time"
import "strings"
//import "encoding/json"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/Patterns/logging_support"
//import "lacima.com/Patterns/msgpack_2"



type monitor_stream_type struct {
    
    sample_time                      int64
    current_time                     int64
}

type stream_records_type struct {
    name                    string
    namespace               string
    description             string
    key_array               []string
    key                     string
    stream_keys             []string
    input_stream_map        map[string]redis_handlers.Redis_Stream_Struct
    peak_detection_map      map[string]redis_handlers.Redis_Hash_Struct
    peak_output_map         map[string]redis_handlers.Redis_Stream_Struct
    incident_log_map        map[string]*logging_support.Incident_Log_Type
}




var monitor_control        monitor_stream_type 
var stream_records         []stream_records_type 






func Init_data_structures(){
   
    
    
    construct_monitor_control()
    construct_construct_stream_records()
    
    
    
}



func construct_monitor_control() {
   
    wd_nodes  := []string{"ERROR_DETECTION:ERROR_DETECTION", "STREAMING_LOGS:STREAMING_LOGS"   }
    nodes := graph_query.Common_qs_search(&wd_nodes)
    node  := nodes[0]
    
    monitor_control.sample_time     = graph_query.Convert_json_int(node["sample_time"])
    fmt.Println("sample time",monitor_control.sample_time)
}
    
func construct_construct_stream_records(){
    
    
    stream_nodes  := []string{"STREAMING_LOG"}
    nodes := graph_query.Common_qs_search(&stream_nodes)
    //fmt.Println("nodes",len(nodes),"\n",nodes)
    stream_records = make([]stream_records_type,len(nodes))
    for index,node := range nodes {
        var item  stream_records_type
        item.namespace    = graph_query.Convert_json_string(node["namespace"])
        item.name         = graph_query.Convert_json_string(node["name"])
        item.description  = graph_query.Convert_json_string(node["descrption"])
        item.stream_keys  = graph_query.Convert_json_string_array(node["keys"])
        item.key_array    = graph_query.Generate_key(item.namespace)
        item  = construct_stream_data_structures(item)
        stream_records[index] = item
    }
    //fmt.Println("stream records ",stream_records)
    panic("done")
    
    
}

func construct_stream_data_structures( item  stream_records_type)  stream_records_type{
    
 
    
        item.input_stream_map     = make(map[string]redis_handlers.Redis_Stream_Struct)
        item.peak_detection_map   = make(map[string]redis_handlers.Redis_Hash_Struct)
        item.peak_output_map      = make(map[string]redis_handlers.Redis_Stream_Struct)
        
        key_temp               := append(item.key_array,"STREAMING_LOG")
        handlers                := data_handler.Construct_Data_Structures(&key_temp)
        for _,key := range item.stream_keys {
            item.input_stream_map[key]         = (*handlers)[key].(redis_handlers.Redis_Stream_Struct)
            item.peak_detection_map[key]       = (*handlers)[key+":ANALYSIS"].(redis_handlers.Redis_Hash_Struct)
            item.peak_output_map[key]          = (*handlers)[key+":PEAK_OUTPUT"].(redis_handlers.Redis_Stream_Struct)
        }
        
        item.incident_log_map = construct_incident_log_map(item)
        
        
        
        return item

    
}


func construct_incident_log_map( item  stream_records_type ) map[string]*logging_support.Incident_Log_Type {
    
    return_value := make(map[string]*logging_support.Incident_Log_Type)
    
    key_array  := item.key_array
    key_array  = append(key_array,"STREAMING_LOG:"+item.name,"INCIDENT_LOG")
   
    nodes      := graph_query.Common_qs_search(&key_array)    
    //fmt.Println(len(nodes))
    for _,node := range nodes {
       name  := graph_query.Convert_json_string(node["name"])
       name_list := strings.Split(name,"^")
       key       := name_list[1]
       //fmt.Println("name",name,key)
       //fmt.Println("item.key_array",item.key_array)
       temp_array   := append(item.key_array,"INCIDENT_LOG:"+item.name+"^"+key,"INCIDENT_LOG")
       fmt.Println("key_array",temp_array)
       return_value[key] = logging_support.Construct_incident_log( temp_array)
    }
    return return_value
    
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
