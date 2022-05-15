package   irr_sched_access

import (
       //"encoding/json"
     "fmt"
      "lacima.com/redis_support/redis_handlers"
     "lacima.com/redis_support/generate_handlers"
)

var rpc_irrigation_driver  redis_handlers.Redis_RPC_Struct

func irrigation_RPC_Client_Init(){
    
    search_list                     := []string{ "IRRIGATION_DATA_STRUCTURES:IRRIGATION_DATA_STRUCTURES",  "RPC_SERVER:IRRIGATION_JOB_QUEUE","RPC_SERVER"}
    handlers                        := data_handler.Construct_Data_Structures(&search_list)
    rpc_irrigation_driver    = (*handlers)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
 
}    
  



 

func rpc_queue_command( rpc_command string, key string, name string,  data map[string]interface{})bool {

   parameters := make(map[string]interface{})
   parameters["COMMAND"]                = rpc_command
   parameters["KEY"]                             = key 
   parameters["NAME"]                         = name 
   parameters["DATA"]                            = data
  
   result := rpc_irrigation_driver.Send_json_rpc_message( "QUEUE_COMMAND",parameters )  	   
   fmt.Printf("%#v \n",result)
   if result == nil {
       return false
   }
   return result["STATUS"].(bool)   

}

func Queue_Action( key,  name string )bool{
    
       parameters := make(map[string]interface{})
       parameters["COMMAND"]                = "QUEUE_ACTION"
       parameters["KEY"]                             = key 
       parameters["NAME"]                         = name 
      result :=   rpc_irrigation_driver.Send_json_rpc_message( "QUEUE_ACTION",parameters )  	   
      
     fmt.Printf("%#v \n",result)
  if result == nil {
       return false
   }
   return result["STATUS"].(bool)   

       
    
}

func Queue_Managed_Irrigation( key string , time float64 , station_io []string  )bool{
       parameters := make(map[string]interface{})
       parameters["COMMAND"]                           = "QUEUE_MANAGED_IRRIGATION"
       parameters["KEY"]                                       = key 
       parameters["TIME"]                                     = time
       parameters["STATION_IO"]                         = station_io
     result :=   rpc_irrigation_driver.Send_json_rpc_message( "QUEUE_MANAGED_IRRIGATION",parameters )  
     fmt.Printf("%#v \n",result)
  if result == nil {
       return false
   }
   return result["STATUS"].(bool)   

}
/*
func Queue_Irrigation( key string , time float64 , station_io []string  )bool{
       panic("should not happen")
       parameters := make(map[string]interface{})
       parameters["COMMAND"]                          = "QUEUE_IRRIGATION"
        parameters["KEY"]                                       = key 
       parameters["TIME"]                                     = time
       parameters["STATION_IO"]                         = station_io
      result :=   rpc_irrigation_driver.Send_json_rpc_message( "QUEUE_IRRIGATION",parameters )
        fmt.Printf("%#v \n",result)
  if result == nil {
       return false
   }
   return result["status"].(bool)   

}
*/
func Queue_Irrigation_Direct( station string , io, time  int64,action bool  )bool{
       parameters := make(map[string]interface{})
       parameters["COMMAND"]                          = "QUEUE_IRRIGATION_DIRECT"
        parameters["STATION"]                             = station 
        parameters["TIME"]                                   = time
       parameters["IO"]                                           = io
      parameters["ACTION"]                                  = action
      result :=   rpc_irrigation_driver.Send_json_rpc_message( "QUEUE_IRRIGATION_DIRECT",parameters )
        fmt.Printf("%#v \n",result)
  if result == nil {
       return false
   }
   return result["STATUS"].(bool)   

}
