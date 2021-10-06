package telegram_rpc_interface



//import "fmt"

import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"


type Telegram_Client_Type struct{

   driver redis_handlers.Redis_RPC_Struct
   
}




 

    
func Site_Server_Init()Telegram_Client_Type{

  var return_value Telegram_Client_Type
  temp := data_handler.Construct_Data_Structures(&[]string{"RPC_SERVER:TELEGRAPH_RPC","RPC_SERVER"} )
  return_value.driver = (*temp)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
  
  return return_value
}  

func (v Telegram_Client_Type)Ping()bool{
      parameters := make(map[string]interface{})
       
       result := v.driver.Send_json_rpc_message( "ping", parameters ) 
       //fmt.Println(result)
       if result != nil {
          return result["status"].(bool)
       }
       return false
}


func (v Telegram_Client_Type)Send_message(message string)bool{
  

       parameters := make(map[string]interface{})
       parameters["message"] = message
       result := v.driver.Send_json_rpc_message( "send_message", parameters ) 
       //fmt.Println("result",result)
       if result != nil {
          return result["status"].(bool)
       }
       return false
      

}
