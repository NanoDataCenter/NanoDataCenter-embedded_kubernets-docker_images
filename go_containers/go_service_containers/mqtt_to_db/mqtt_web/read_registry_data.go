package mqtt_web



import "fmt"
import "time"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"

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




var redis_topic_value          redis_handlers.Redis_Hash_Struct
var redis_topic_time_stamp     redis_handlers.Redis_Hash_Struct
var redis_device_status        redis_handlers.Redis_Hash_Struct
var redis_topic_handler        redis_handlers.Redis_Hash_Struct   
var redis_topic_error_ts       redis_handlers.Redis_Hash_Struct 
 
 
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
                                                       // tag4 handler 
                                                       // date time string
                                                       // data msgpack data
                                                       



func Construct_event_registry_actions( topic string){
   
   top_topic_string = "/"+topic+"/#"
   base_topic_string = "/"+topic+"/"
   construct_drivers()
   construct_mqtt_enviroment()
   construct_contact_map()
   construct_topic_handlers()

}

func get_monitoring_topic()string{
    fmt.Println("top_topic_string",top_topic_string)
    return top_topic_string
}




func construct_drivers(){
    
   data_search_list         := []string{ "TOPIC_STATUS"}
   data_element             := data_handler.Construct_Data_Structures(&data_search_list)
   redis_topic_value        = (*data_element)["TOPIC_VALUE"].(redis_handlers.Redis_Hash_Struct)
   redis_topic_time_stamp   = (*data_element)["TOPIC_TIME_STAMP"].(redis_handlers.Redis_Hash_Struct)
   redis_device_status      = (*data_element)["DEVICE_STATUS"].(redis_handlers.Redis_Hash_Struct)
   redis_topic_handler      = (*data_element)["TOPIC_HANDLER"].(redis_handlers.Redis_Hash_Struct)
   redis_topic_error_ts     = (*data_element)["TOPIC_ERROR_TIME_STAMP"].(redis_handlers.Redis_Hash_Struct)
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












func construct_contact_map(){
    
    contact_map     = make(map[string]contact_type)
    for key,item := range device_map{
      
      redis_device_status.HSet(key,"true")
      var contact_entry contact_type
      contact_entry.contact_time   = time.Now().Unix()
      contact_entry.delta_time     = class_map[item.class].contact_time
      contact_map[key]             = contact_entry
  }    
    
    
}



func construct_topic_handlers(){
   for class_name,element := range class_map{
       for _, device_name := range element.device_list {
           for _, topic_name := range element.topic_list {
               topic_handler := topic_map[topic_name].handler_type
               full_topic := base_topic_string+class_name+"/"+device_name+"/"+topic_name
               redis_topic_handler.HSet(full_topic,topic_handler)
           }
       }
   }
    
}




