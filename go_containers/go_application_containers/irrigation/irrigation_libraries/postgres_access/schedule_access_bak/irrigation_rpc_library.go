package   irr_sched_access


import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"


var rpc_irrigation_driver  redis_handlers.Redis_RPC_Struct

func irrigation_RPC_Client_Init(){
    
    search_list                     := []string{ "IRRIGATION_DATA_STRUCTURES:IRRIGATION_DATA_STRUCTURES",  "RPC_SERVER:IRRIGATION_JOB_QUEUE","RPC_SERVER"}
    handlers                        := data_handler.Construct_Data_Structures(&search_list)
    rpc_irrigation_driver    = (*handlers)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
 
}    
  





func  RPC_Queue_Command( command string, data interface{})bool {

   
   parameters := make(map[string]interface{})
   parameters["COMMAND"]        = command
   parameters["PARAMETERS"]  = data
   result := rpc_irrigation_driver.Send_json_rpc_message( "QUEUE_COMMAND",parameters )  	   
   return result["status"].(bool)   


}

