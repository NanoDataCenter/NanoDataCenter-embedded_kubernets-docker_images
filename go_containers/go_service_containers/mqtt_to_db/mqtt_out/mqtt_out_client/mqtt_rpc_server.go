package mqtt_out_client


import "strconv"
import "strings"
//import "fmt"
//import "os"
import "time"
//import b64 "encoding/base64"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"
import  b64 "encoding/base64"

func set_up_rpc_server(){

    
     search_list := []string{"MQTT_OUTPUT_SETUP:site_out_server","RPC_SERVER:MQTT_OUT_RPC_SERVER","RPC_SERVER"}

     handlers := data_handler.Construct_Data_Structures(&search_list)
    
     driver := (*handlers)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
	
	
	driver.Add_handler( "publish",publish_message)
	driver.Json_Rpc_start()
	
	
   
}


func publish_message( parameters map[string]interface{} ) map[string]interface{}{
   topic                  := parameters["topic"].(string)
   payload_b64            := parameters["payload"].(string)
   parameters["results"]  = ""
   
   
   topic                  = base_topic_string+topic
   
   if verify_topic( topic ) == false {
       parameters["status"] = false
       return parameters
   }
   
   
   
   
    payload,_    := b64.StdEncoding.DecodeString(payload_b64)
    

    //fmt.Println("payload",string(payload))
   
    client.Publish(topic, 2 , false, payload )
   
    parameters["status"] = true
    
    update_transmit_data_structures(topic,string(payload))
  
    return parameters

}
   
 
func verify_topic( topic string )bool {
    
  
  if err := redis_topic_handler.HExists(topic); err == false {
       time_string := strconv.Itoa(int(time.Now().Unix()) )
       redis_topic_error_ts.HSet(topic,time_string) 
       return false
   }  
   return true 
    
}


func  update_transmit_data_structures(topic,payload string){
    
     topic_array := strings.Split(topic,"/") // structure  /base_string:1/class:2/device:3/topic:4
     time_string := strconv.Itoa(int(time.Now().Unix()) )
     if len(topic_array) < 4 {
          redis_topic_error_ts.HSet(topic,time_string) 
         return
     }

     
    
     
     // identify class instance and topic
     class     := topic_array[1]
     instance  := topic_array[2]
    
     extracted_topic_array    := topic_array[3:]
     extracted_topic          := strings.Join(extracted_topic_array,"/")
     
     

     redis_topic_time_stamp.HSet(topic, time_string)
     redis_topic_value.HSet(topic,payload)
     
     postges_topic_stream.Insert( class,instance,extracted_topic,"",time_string,payload  )
     
    
}
