package support

import "fmt"
import "lacima.com/redis_support/redis_handlers"
import mqtt "github.com/eclipse/paho.mqtt.golang"
import "lacima.com/Patterns/logging_support"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/server_libraries/postgres"

var top_topic_string string

var topic_hash      redis_handlers.Redis_Hash_Struct
var error_hash      redis_handlers.Redis_Hash_Struct
var time_hash       redis_handlers.Redis_Hash_Struct
var incident_log    *logging_support.Incident_Log_Type
var topic_stream    pg_drv.Postgres_Stream_Driver


func Construct_event_registry_tasks(topic string){
   
   top_topic_string = "/"+topic+"/#"
   construct_incident_log()
   construct_topic_hash()
   construct_postgres_driver()
   construct_device_classes()
   setup_topic_hash()
   setup_other_hashes()
 
}

func construct_incident_log(){
    
    incident_log  = logging_support.Construct_incident_log([]string{"INCIDENT_LOG:MQTT_LOG","INCIDENT_LOG"} )
    
}

func construct_topic_hash(){
    
   data_search_list := []string{ "TOPIC_STATUS"}
   data_element := data_handler.Construct_Data_Structures(&data_search_list)
   topic_hash   = (*data_element)["TOPIC_STATUS"].(redis_handlers.Redis_Hash_Struct)
   error_hash   = (*data_element)["BAD_TOPIC_STATUS"].(redis_handlers.Redis_Hash_Struct)
   time_hash    = (*data_element)["TIME_STATUS"].(redis_handlers.Redis_Hash_Struct)  
   topic_stream = (*data_element)["POSTGRES_STREAM"].(pg_drv.Postgres_Stream_Driver)
   //fmt.Println("incident_log",incident_log) 
   //fmt.Println("topic_hash",topic_hash)
   //fmt.Println("error_hash",error_hash)
   //fmt.Println("time_hash",time_hash)
   //fmt.Println("topic_stream",topic_stream)
}


func construct_postgres_driver(){
    
}


func construct_device_classes(){
    
    
}


func setup_topic_hash(){
    

    
}

func setup_other_hashes(){

    time_hash.Delete_All()
    error_hash.Delete_All()  
    
}



func get_monitoring_topic()string{
    fmt.Println("top_topic_string",top_topic_string)
    return top_topic_string
}



func receive_mqtt_packet(msg mqtt.Message){
    
     // id class from topic 
    // id device from topic
     fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
    
}
