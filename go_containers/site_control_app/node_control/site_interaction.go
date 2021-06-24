package node_control


import "fmt"
import "time"




//import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"
import  "lacima.com/cf_control"
import "lacima.com/Patterns/shell_utils"
import "lacima.com/Patterns/logging_support"


var site_data *map[string]interface{}
var wd_struct *logging_support.Watch_Dog_Log_Type


// local copy of site data

func setup_site_control(cf_cluster *cf.CF_CLUSTER_TYPE , site_data_input *map[string]interface{}){

 site_data = site_data_input
 
 go start_rpc_server()   
    
    
    
}



func strobe_watch_dog( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {

     wd_struct.Strobe_Watch_Dog()
     return cf.CF_DISABLE
}




func start_rpc_server(){
     fmt.Println("made it here")
    
     search_list := []string{"PROCESSOR:"+(*site_data)["local_node"].(string),"RPC_SERVER:NODE_CONTROL","RPC_SERVER"}
     handlers := data_handler.Construct_Data_Structures(&search_list)
     driver := (*handlers)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)    
     driver.Add_handler( "reboot",reboot_system)
     driver.Json_Rpc_start()
}


func reboot_system( parameters map[string]interface{} ) map[string]interface{}{
    
    if (*site_data)["master"] == true {
       time.Sleep(time.Second*15)    
    
    }
    shell_utils.System_shell("reboot")
    return parameters
}


