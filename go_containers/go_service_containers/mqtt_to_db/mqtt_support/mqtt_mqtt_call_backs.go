package mqtt_support

import "fmt"
import "time"
import "strings"
import "strconv"

import mqtt "github.com/eclipse/paho.mqtt.golang"




func receive_mqtt_packet(msg mqtt.Message){
    
     topic :=  string(msg.Topic())
     data  :=  string(msg.Payload())
    
     time_string := strconv.Itoa(int(time.Now().Unix()) )
     topic_array := strings.Split(topic,"/") // structure  /base_string:1/class:2/device:3/topic:4
     
     if len(topic_array) < 5 {
          redis_topic_error_ts.HSet(topic,time_string) 
         return
     }
     if err := redis_topic_handler.HExists(topic); err == false {
         redis_topic_error_ts.HSet(topic,time_string) 
         return
     }
     

     
    
     
     
     class  := topic_array[2]
     device := topic_array[3]
     
     
     extracted_topic_array    := topic_array[3:]
     extracted_topic          := strings.Join(extracted_topic_array,"/")
     
     contact_value               := contact_map[device]
     contact_value.contact_time  = time.Now().Unix()
     contact_map[device]         = contact_value
     
     
     redis_topic_time_stamp.HSet(topic, time_string)
     redis_topic_value.HSet(topic,string(msg.Payload()))
     handler := handler_map[extracted_topic]
     
     timeT := time.Unix(contact_value.contact_time, 0)
     time_string = fmt.Sprintf("time.Time: %s\n", timeT)
     
     redis_topic_time_stamp.HSet(topic,time_string)
     postges_topic_stream.Insert( class,device,topic,handler,time_string,data  )
     
     fmt.Printf("Received topic: %s message: %s from \n" , msg.Topic(), msg.Payload())
    
}

func log_on_connection(){
    fmt.Println("mqtt on")
    mqtt_incident_log.Log_data( true,"receive_connection","receive_connection")
    
}


func log_off_connection(){
    fmt.Println("mqtt off")
    mqtt_incident_log.Log_data( false, "lost_connection", "lost_connection" )
    
}
