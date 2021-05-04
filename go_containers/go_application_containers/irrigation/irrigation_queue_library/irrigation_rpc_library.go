package irrigation_rpc


import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"

type Irrigation_Client_Type struct{

   driver redis_handlers.Redis_RPC_Struct
}

var rcp_command_map map[string]bool


func Irrigation_RPC_Client_Init(search_list *[]string)Irrigation_Client_Type{

  rcp_command_map = make(map[string]bool)
  rcp_command_map["CLEAN_FILTER"]              = true                
  rcp_command_map["OPEN_MASTER_VALVE"]         = true          
  rcp_command_map["CLOSE_MASTER_VALVE"]        = true                      
  rcp_command_map["CHECK_OFF"]                = true    
  rcp_command_map["RESISTANCE_CHECK"]          = true  
  rcp_command_map["CLEAR_IRRIGATION_QUEUE"]    = true
  rcp_command_map["CLEAR_OFFLINE_IRRIGATION"]  = true
  rcp_command_map["SUSPEND"]                   = true                                
  rcp_command_map["RESUME"]                    = true    
  rcp_command_map["SKIP_STATION"]              = true 
  
  var return_value Irrigation_Client_Type
  var handlers = data_handler.Construct_Data_Structures(search_list)  
  return_value.driver = (*handlers)["IRRIGATION_JOB_SERVER"].(redis_handlers.Redis_RPC_Struct)
 
  return return_value
}  



func (v Irrigation_Client_Type)Ping()bool{
   
    var result = v.driver.Send_json_rpc_message( "PING", make(map[string]interface{}) ) 
                              
   return result["status"].(bool)
}

func (v Irrigation_Client_Type)Queue_Command( command string)bool {

   if _,ok := rcp_command_map[command];ok==false{
     return false
   }
   parameters := make(map[string]interface{})
   parameters["COMMAND"] = command
   var result = v.driver.Send_json_rpc_message( "QUEUE_COMMAND",parameters )  	   
   return result["status"].(bool)   


}

func (v Irrigation_Client_Type)Queue_controller_pin(time int, io map[string]int)bool{
  
   parameters := make(map[string]interface{})
   parameters["time"] = time
   parameters["io"] = io
   var result = v.driver.Send_json_rpc_message( "QUEUE_IRRIGATION_CONTROLLER_PIN", parameters ) 
 	   
   return result["status"].(bool)
}
  
func (v Irrigation_Client_Type)Queue_offline_controller_pin(time int, io map[string]int)bool{
  
   parameters := make(map[string]interface{})
   parameters["time"] = time
   parameters["io"] = io
   var result = v.driver.Send_json_rpc_message( "QUEUE_OFFLINE_CONTROLLER_PIN", parameters ) 
 	   
   return result["status"].(bool)
}



func (v Irrigation_Client_Type)Queue_schedule( schedule string ){


}


func ( v Irrigation_Client_Type)Queue_schedule_step(schedule string, index int64 ){



}

func ( v Irrigation_Client_Type)Queue_schedule_step_time(schedule string, index int, time int64 ){



}






  
