package irrigation_rpc


import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"

type Irrigation_Client_Type struct{

   driver redis_handlers.Redis_RPC_Struct
}



func Irrigation_RPC_Client_Init(search_list *[]string)Irrigation_Client_Type{

    return_value Irrigation_Client_Type
    
    search_list                 := []string{ "IRRIGATION_DATA_STRUCTURES:IRRIGATION_DATA_STRUCTURES",  "RPC_SERVER:IRRIGATION_JOB_QUEUE","RPC_SERVER"}
    handlers                     := data_handler.Construct_Data_Structures(&search_list)
    return_value.driver    := (*handlers)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
   
    return return_value
}    
  





func (v Irrigation_Client_Type)Queue_Command( command string, data interface{})bool {

   if _,ok := rcp_command_map[command];ok==false{
     return false
   }
   parameters := make(map[string]interface{})
   parameters["COMMAND"]        = command
   parameters["PARAMETERS"]  = data
   var result = v.driver.Send_json_rpc_message( "QUEUE_COMMAND",parameters )  	   
   return result["status"].(bool)   


}

