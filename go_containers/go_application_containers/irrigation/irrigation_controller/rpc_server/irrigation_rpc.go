package irrigation_rpc

import (
         // "lacima.com/redis_support/graph_query"
          "lacima.com/redis_support/redis_handlers"
          "lacima.com/redis_support/generate_handlers"
          "fmt"
)


func Start(){
    
//[SYSTEM:farm_system][SITE:LACIMA_SITE][IRRIGATION_DATA_STRUCTURES:IRRIGATION_DATA_STRUCTURES][SCHEDULE_DATA:SCHEDULE_DATA][RPC_SERVER:IRRIGATION_JOB_QUEUE][PACKAGE:RPC_SERVER]
 	  
    fmt.Println("irrigaiton initialization rpc server")
     search_list := []string{ "IRRIGATION_DATA_STRUCTURES:IRRIGATION_DATA_STRUCTURES",  "RPC_SERVER:IRRIGATION_JOB_QUEUE","RPC_SERVER"}

     handlers := data_handler.Construct_Data_Structures(&search_list)
    
     driver := (*handlers)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
	

	
	driver.Add_handler( "QUEUE_ACTION", queue_actions)
    driver.Add_handler(  "QUEUE_MANAGED_IRRIGATION", queue_managed_irrigation)
    driver.Add_handler( "QUEUE_IRRIGATION", queue_irrigation)
    driver.Add_handler("QUEUE_IRRIGATION_DIRECT",queue_irrigation_direct)
	driver.Json_Rpc_start() 
    
    
}    
    
func queue_actions( parameters map[string]interface{} ) map[string]interface{}{
  
  fmt.Println("queue_actions",parameters)
  parameters["status"] = true
  return parameters
}

 func queue_managed_irrigation( parameters map[string]interface{} ) map[string]interface{}{
  
  fmt.Println("queue_managed_irrigation",parameters)
  parameters["status"] = true
  return parameters
}

func queue_irrigation( parameters map[string]interface{} ) map[string]interface{}{
  fmt.Println("queue_irrigation",parameters)
  parameters["status"] = true
  return parameters
}

func queue_irrigation_direct( parameters map[string]interface{} ) map[string]interface{}{
  fmt.Println("queue_irrigation_direct",parameters)
  parameters["status"] = true
  return parameters
}
