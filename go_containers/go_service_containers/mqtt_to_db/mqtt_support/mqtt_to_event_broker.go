package mqtt_support

import "fmt"
import "time"
import "strings"
import "strconv"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import mqtt "github.com/eclipse/paho.mqtt.golang"
import "lacima.com/Patterns/logging_support"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/server_libraries/postgres"

var top_topic_string string
var base_topic_string string

var input_topic_map         map[string]string
var input_error_map         map[string]int64
var device_status           map[string]bool
var handler_map             map[string]string


type contact_type struct {
    contact_time int64
    delta_time   int64
}

var contact_map map[string]contact_type



var mqtt_incident_log          *logging_support.Incident_Log_Type // incident_log for mqtt server
var redis_topic_value          redis_handlers.Redis_Hash_Struct
var redis_topic_time_stamp     redis_handlers.Redis_Hash_Struct
var redis_device_status        redis_handlers.Redis_Hash_Struct
var redis_topic_handler        redis_handlers.Redis_Hash_Struct   
 
 
 
var postges_topic_stream    pg_drv.Postgres_Stream_Driver      // time stream for all topics
                                                       // tag1 class
                                                       // tag2 device
                                                       // tag3 topic
                                                       // tag4 msgpack handler 
                                                       // tag5 not used
                                                       // data msgpack data
var postgres_incident_stream    pg_drv.Postgres_Stream_Driver      // time stream for all device changes
                                                       // tag1 class
                                                       // tag2 device
                                                       // tag3 status
                                                       // tag4 not used 
                                                       // tag5 not used
                                                       // data not used
                                                       
// need postgress inventory table for topics
// need postgress inventory table for classes 
// need postgress inventory table for devices



func Construct_event_registry_tasks(topic string){
   
   top_topic_string = "/"+topic+"/#"
   base_topic_string = "/"+topic+"/"
   construct_incident_log()
   construct_drivers()
   construct_mqtt_enviroment()
   construct_mqtt_maps()
   
   
   
 
}

func construct_incident_log(){
    
    mqtt_incident_log   = logging_support.Construct_incident_log([]string{"INCIDENT_LOG:MQTT_LOG","INCIDENT_LOG"} )
    
}

func construct_drivers(){
    
   data_search_list         := []string{ "TOPIC_STATUS"}
   data_element             := data_handler.Construct_Data_Structures(&data_search_list)
   redis_topic_value        = (*data_element)["TOPIC_VALUE"].(redis_handlers.Redis_Hash_Struct)
   redis_topic_time_stamp   = (*data_element)["TOPIC_TIME_STAMP"].(redis_handlers.Redis_Hash_Struct)
   redis_device_status      = (*data_element)["DEVICE_STATUS"].(redis_handlers.Redis_Hash_Struct)
   redis_topic_handler      = (*data_element)["TOPIC_HANDLER"].(redis_handlers.Redis_Hash_Struct)
   postges_topic_stream     = (*data_element)["POSTGRES_DATA_STREAM"].(pg_drv.Postgres_Stream_Driver)
   postgres_incident_stream = (*data_element)["POSTGRES_INCIDENT_STREAM"].(pg_drv.Postgres_Stream_Driver)
}

   
func construct_mqtt_enviroment(){
    
    
   data_nodes := graph_query.Common_qs_search(&[]string{"MQTT_DEVICES:MQTT_DEVICES"})
   data_node  := data_nodes[0]

   class_dict :=  graph_query.Convert_json_nested_dictionary_interface(data_node["classes"])
   device_dict := graph_query.Convert_json_nested_dictionary_interface(data_node["devices"])
   topic_dict   := graph_query.Convert_json_nested_dictionary_interface(data_node["topics"])

   register_topics(topic_dict)
   register_classes(class_dict)
   register_devices(device_dict)
   
}





func construct_mqtt_maps(){
  input_topic_map = make(map[string]string)
  input_error_map = make(map[string]int64)
  contact_map     = make(map[string]contact_type)
  handler_map     = make(map[string]string)
  device_status   = make(map[string]bool)
  
  for key,item := range device_map{
      
      for key,item := range topic_map {
          redis_topic_handler.HSet(key,item.handler_type)
      }
          
      
      
      redis_device_status.HSet(key,"true")
      device_status[key]  = true
      class := item.class
      name  := item.name
      for _,element := range class_map[class].topic_list {
          topic := base_topic_string+class+"/"+name+"/"+element
          input_topic_map[topic] = "true"
          handler_map[topic] = topic_map[element].handler_type
      }
      
      var contact_entry contact_type
      contact_entry.contact_time   = time.Now().Unix()
      contact_entry.delta_time     = class_map[class].contact_time
      contact_map[name] = contact_entry
  }
}
  




func get_monitoring_topic()string{
    fmt.Println("top_topic_string",top_topic_string)
    return top_topic_string
}


func receive_mqtt_packet(msg mqtt.Message){
    
     topic :=  string(msg.Topic())
     data  :=  string(msg.Payload())
     
     topic_array := strings.Split(topic,"/")
     
     if len(topic_array) < 5 {
         input_error_map[topic] = time.Now().Unix()
         return
     }
     if _,err := input_topic_map[topic]; err == false {
         input_error_map[topic] = time.Now().Unix()
         return
     }
     
     
     class  := topic_array[3]
     device := topic_array[4]
     contact_value := contact_map[device]
     contact_value.contact_time  = time.Now().Unix()
     contact_map[device] = contact_value
     
     
     redis_topic_time_stamp.HSet(topic, strconv.Itoa(int(time.Now().Unix()) ))
     redis_topic_handler.HSet(msg.Topic(),string(msg.Payload()))
     handler := handler_map[topic]
     postges_topic_stream.Insert( class,device,topic,handler,"",data  )
     
     fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
    
}

func log_on_connection(){
    fmt.Println("mqtt on")
    mqtt_incident_log.Log_data( true,"receive_connection","receive_connection")
    
}


func log_off_connection(){
    fmt.Println("mqtt off")
    mqtt_incident_log.Log_data( false, "lost_connection", "lost_connection" )
    
}



