package mqtt_test_web



//import "fmt"

import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"

import "lacima.com/redis_support/generate_handlers"
import "lacima.com/server_libraries/postgres"


var top_topic_string string
var base_topic_string string

var output_topic_map         map[string]string
var output_error_map         map[string]int64





var redis_topic_value          redis_handlers.Redis_Hash_Struct
var redis_topic_time_stamp     redis_handlers.Redis_Hash_Struct
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
                                                       



func construct_event_registry_actions( ){
   

   construct_drivers()
   construct_mqtt_enviroment()
   
}




func construct_drivers(){
    
   data_search_list         := []string{ "MQTT_OUTPUT_SETUP:site_out_server","TOPIC_STATUS"}
   data_element             := data_handler.Construct_Data_Structures(&data_search_list)
   redis_topic_value        = (*data_element)["TOPIC_VALUE"].(redis_handlers.Redis_Hash_Struct)
   redis_topic_time_stamp   = (*data_element)["TOPIC_TIME_STAMP"].(redis_handlers.Redis_Hash_Struct)
   redis_topic_error_ts     = (*data_element)["TOPIC_ERROR_TIME_STAMP"].(redis_handlers.Redis_Hash_Struct)
   postges_topic_stream     = (*data_element)["POSTGRES_DATA_STREAM"].(pg_drv.Postgres_Stream_Driver)
   
   redis_topic_value.Delete_All()
   redis_topic_time_stamp.Delete_All()
   redis_topic_error_ts.Delete_All() 

   
}
 
   
func construct_mqtt_enviroment(){
    
    
   data_nodes := graph_query.Common_qs_search(&[]string{"MQTT_OUTPUT_SETUP:site_out_server","MQTT_INSTANCES:MQTT_INSTANCES"})
   data_node  := data_nodes[0]

   class_dict    :=  graph_query.Convert_json_nested_dictionary_interface(data_node["classes"])
   instance_dict := graph_query.Convert_json_nested_dictionary_interface(data_node["instances"])
   topic_dict    := graph_query.Convert_json_nested_dictionary_interface(data_node["topics"])

   register_topics(topic_dict)
   register_classes(class_dict)
   register_instances(instance_dict)
   
}








