package site_control_server_lib



//import "fmt"

import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/graph_query"
import "lacima.com/Patterns/logging_support"

type Site_Server_Client_Type struct{

   driver redis_handlers.Redis_RPC_Struct
   
}




 

    
func Site_Server_Init()Site_Server_Client_Type{

  var return_value Site_Server_Client_Type
  temp := data_handler.Construct_Data_Structures(&[]string{"RPC_SERVER:TELEGRAPH_RPC","RPC_SERVER"} )
  return_value.driver = (*temp)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
  
  return return_value
}  



func (v Site_Server_Client_Type)Send_message()bool{
  

       parameters := make(map[string]interface{})
       parameters["message"] = message
       result := v.driver.Send_json_rpc_message( "send_message", parameters ) 
       //fmt.Println("result",result)
       if result != nil {
          return result["status"].(bool)
       }
       return false
      

}
