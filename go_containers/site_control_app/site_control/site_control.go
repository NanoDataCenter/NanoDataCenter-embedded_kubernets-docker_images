package site_control


import "fmt"


import "lacima.com/cf_control"
import "lacima.com/Patterns/logging_support"


var wd_struct *logging_support.Watch_Dog_Log_Type
var wd_array  []*wd_struct


func Site_Startup(cf_cluster *cf.CF_CLUSTER_TYPE , site_data *map[string]interface{}){

   	
	go start_rpc_server() 
}



+
func find_watch_dog_structures(){
    
    
    
}

func find_node_rpc_servers(){
    
    
    
    
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


