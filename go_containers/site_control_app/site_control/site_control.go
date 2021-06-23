package site_control


import "fmt"

import "strconv"
import "lacima.com/cf_control"
import "lacima.com/server_libraries/node_control_rpc"

import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"
import "time"

var  node_status_hash   redis_handlers.Redis_Hash_Struct

var  node_rpc_servers   node_control_server_lib.Node_Server_Client_Type


func Site_Startup(cf_cluster *cf.CF_CLUSTER_TYPE , site_data *map[string]interface{}){

   	node_rpc_servers = node_control_server_lib.Node_Server_Init()
    monitor_node_rpc_servers(cf_cluster)
	go start_rpc_server() 
}



func monitor_node_rpc_servers(cf_cluster *cf.CF_CLUSTER_TYPE){
  var cf_control  cf.CF_SYSTEM_TYPE
  
  search_list := []string{"NODE_MAP"}
  handlers := data_handler.Construct_Data_Structures(&search_list)    
  node_status_hash = (*handlers)["NODE_MAP"].(redis_handlers.Redis_Hash_Struct)    
  node_status_hash.Delete_All()

  (cf_control).Init(cf_cluster ,"site_control_node_monitoring",true, time.Second)
  (cf_control).Add_Chain("node_monitoring",true)
  (cf_control).Cf_add_log_link("node_monitoring")
   
  var parameters = make(map[string]interface{})
  (cf_control).Cf_add_one_step(node_monitor,parameters)
  
  (cf_control).Cf_add_wait_interval(time.Second*60 )
  (cf_control).Cf_add_reset()
  
    
}


func node_monitor( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {

	 for node,_ := range node_rpc_servers.Driver_array{
         result := node_rpc_servers.Ping(node)
         node_status_hash.HSet(node,strconv.FormatBool(result))
     }
	 panic("done")
     return cf.CF_DISABLE
}



 
func start_rpc_server(){
     fmt.Println("made it here")
    
     search_list := []string{"RPC_SERVER:SYSTEM_CONTROL","RPC_SERVER"}
     handlers := data_handler.Construct_Data_Structures(&search_list)
     driver := (*handlers)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)    
     driver.Add_handler( "reboot",reboot_system)
     driver.Json_Rpc_start()
}


func reboot_system( parameters map[string]interface{} ) map[string]interface{}{
    
    
    return parameters
}


