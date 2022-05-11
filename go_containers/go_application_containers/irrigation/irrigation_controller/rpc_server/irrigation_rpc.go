package irrigation_rpc

import (
         // "lacima.com/redis_support/graph_query"
          "lacima.com/redis_support/redis_handlers"
          "lacima.com/redis_support/generate_handlers"
          "fmt"
)


func Start(){
    
   // get  action data structures 
  // irrigation controllers and sub controllers
  // get station and io data
    // map station/io to master and sub controllers
 	  
    fmt.Println("irrigaiton initialization rpc server")
     search_list := []string{ "IRRIGATION_DATA_STRUCTURES:IRRIGATION_DATA_STRUCTURES",  "RPC_SERVER:IRRIGATION_JOB_QUEUE","RPC_SERVER"}

     handlers := data_handler.Construct_Data_Structures(&search_list)
    
     driver := (*handlers)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
	

	
	driver.Add_handler( "QUEUE_ACTION", handler_actions)
    driver.Add_handler(  "QUEUE_MANAGED_IRRIGATION", handle_managed_irrigation)
    driver.Add_handler("QUEUE_IRRIGATION_DIRECT",handler_irrigation_direct)
	driver.Json_Rpc_start() 
    
    
}    
