package mqtt_publish_server_lib





import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"
import  b64 "encoding/base64"

type MQTT_Publish_Client_Type struct{

   driver redis_handlers.Redis_RPC_Struct
}




func MQTT_Publish_Init(search_list *[]string)MQTT_Publish_Client_Type{

  var return_value MQTT_Publish_Client_Type
  handlers := data_handler.Construct_Data_Structures(search_list)  
  return_value.driver = (*handlers)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
  return return_value
}  

func (v MQTT_Publish_Client_Type)Ping()bool{
  

       parameters := make(map[string]interface{})
       
       result := v.driver.Send_json_rpc_message( "ping", parameters ) 
       return result["status"].(bool)

}

func (v MQTT_Publish_Client_Type)Publish(topic,payload string)bool {
       
      
       payload_b64           := b64.StdEncoding.EncodeToString([]byte(payload))
       
       parameters            := make(map[string]interface{})
       parameters["topic"]    = topic
       parameters["payload"]  = payload_b64
       result := v.driver.Send_json_rpc_message( "publish", parameters ) 
	   status := result["status"].(bool)
       return status

}







