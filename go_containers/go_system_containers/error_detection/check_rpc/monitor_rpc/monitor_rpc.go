package monitor_rpc


import "fmt"
import "time"
import "strings"

import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/Patterns/logging_support"



type monitor_rpc_type struct {
    
    sample_time                      int64
    incident_log                     *logging_support.Incident_Log_Type
    
}

type rpc_records_type struct {
    namespace    string
    key_array    []string
    key          string
    rpc_server   redis_handlers.Redis_RPC_Struct
}






var monitor_control  monitor_rpc_type
var rpc_records      []rpc_records_type




func Init_data_structures(){
 
    construct_monitor_control()
    construct_construct_rpc_servers()
    
    
    
}



func construct_monitor_control() {
   
    wd_nodes  := []string{"ERROR_DETECTION:ERROR_DETECTION", "RPC_ANALYSIS:RPC_ANALYSIS"   }
    nodes := graph_query.Common_qs_search(&wd_nodes)
    node  := nodes[0]
    
    monitor_control.sample_time     = graph_query.Convert_json_int(node["sample_time"])
    monitor_control.incident_log    = logging_support.Construct_incident_log([]string{"ERROR_DETECTION:ERROR_DETECTION", "RPC_ANALYSIS:RPC_ANALYSIS" ,"INCIDENT_LOG"} )
   
}
    




func construct_construct_rpc_servers(){
    rpc_records = make([]rpc_records_type,0)
    incident_nodes  := []string{"RPC_SERVER"}
    nodes := graph_query.Common_qs_search(&incident_nodes)
    fmt.Println("nodes",len(nodes),nodes)
   
    for _,node := range nodes{
        var item  rpc_records_type
               
        item.namespace          = graph_query.Convert_json_string(node["namespace"])
        item.key_array          = graph_query.Generate_key(item.namespace)
        item.key_array          = append(item.key_array,"RPC_SERVER")
        item.key                = strings.Join(item.key_array,"/")
        
        handlers                := data_handler.Construct_Data_Structures(&item.key_array)
        item.rpc_server         = (*handlers)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
        rpc_records             = append(rpc_records,item)
    }
    
}


func Process_functions(){
        
   go ping_rpc_server_loop()
  
}






func ping_rpc_server_loop(){
    timeout   := time.Duration(monitor_control.sample_time)*time.Second
    //time.Sleep(time.Minute)
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
}       
    
   
func ping_rpc_server(item  rpc_records_type ){
    
   parameters := make(map[string]interface{})
   result := item.rpc_server.Send_json_rpc_message( "info", parameters ) 
   if result == nil {
       fmt.Println("rpc not active")
   }else{
      fmt.Println("result", result )
   }
    
}
