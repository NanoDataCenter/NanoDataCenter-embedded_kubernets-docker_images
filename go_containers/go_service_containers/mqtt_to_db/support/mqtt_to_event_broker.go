package support

import "fmt"
import mqtt "github.com/eclipse/paho.mqtt.golang"



func Construct_event_registry_tasks(){
   
   // get inclident log
   // construct classes
   // construct drivers for each classes
    
    
    
    
}


func get_monitoring_topic()string{
    return "/topic/#"
}



func receive_mqtt_packet(msg mqtt.Message){
    
     // id class from topic 
    // id device from topic
     fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
    
}
