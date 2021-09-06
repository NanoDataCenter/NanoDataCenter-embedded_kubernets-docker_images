package mqtt_client

import "fmt"
import "os"
import "time"
import "strings"
import "strconv"
//import b64 "encoding/base64"
import mqtt "github.com/eclipse/paho.mqtt.golang"




func receive_mqtt_packet(msg mqtt.Message){
    
     topic :=  string(msg.Topic())
     data  :=  string(msg.Payload())
    
     /*
      * Verifying Topic
      * 
      */
      time_string := strconv.Itoa(int(time.Now().Unix()) )
      if err := redis_topic_handler.HExists(topic); err == false {
         redis_topic_error_ts.HSet(topic,time_string) 
         return
     }
     

     
     
     
    
     topic_array := strings.Split(topic,"/") // structure  /base_string:1/class:2/device:3/topic:4
     
     if len(topic_array) < 5 {
          redis_topic_error_ts.HSet(topic,time_string) 
         return
     }

     
    
     
     // identify class device and topic
     class  := topic_array[2]
     device := topic_array[3]
     
     
     extracted_topic_array    := topic_array[4:]
     extracted_topic          := strings.Join(extracted_topic_array,"/")
     
     
     // store in redis hash tables
     redis_contact_time.HSet(device,time_string) 
     redis_topic_time_stamp.HSet(topic, time_string)
     redis_topic_value.HSet(topic,data)
     
     
     
     
     
     // store in postgres db
     handler := handler_map[extracted_topic]
     postges_topic_stream.Insert( class,device,topic,handler,time_string,data  )
     
    
}


func log_on_connection(){
    fmt.Println("mqtt on")
    connection_status = true
  
    
}


func log_off_connection(){
    fmt.Println("mqtt off")
    connection_status = false
   
    fmt.Println("os exit")
    os.Exit(1)
    
}
