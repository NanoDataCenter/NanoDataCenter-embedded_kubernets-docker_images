package mqtt_out_client

//import "fmt"


import "lacima.com/redis_support/redis_handlers"

import "lacima.com/Patterns/logging_support"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/server_libraries/postgres"

var top_topic_string string
var base_topic_string string

var input_topic_map         map[string]string
var input_error_map         map[string]int64








var mqtt_incident_log            *logging_support.Incident_Log_Type // incident_log for mqtt server
var redis_topic_value            redis_handlers.Redis_Hash_Struct
var redis_topic_time_stamp       redis_handlers.Redis_Hash_Struct
var redis_device_status          redis_handlers.Redis_Hash_Struct
var redis_topic_handler          redis_handlers.Redis_Hash_Struct   
var redis_topic_error_ts         redis_handlers.Redis_Hash_Struct
var redis_contact_time           redis_handlers.Redis_Hash_Struct

 
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
                                                       



func Construct_event_registry_actions(base_topic string){
   
  base_topic_string = "/"+base_topic+"/"
   
   construct_drivers()
   construct_mqtt_device_enviroment()
   
   construct_topic_handlers()

}




func construct_drivers(){
    
   data_search_list              := []string{"MQTT_OUTPUT_SETUP:site_out_server","TOPIC_STATUS"}
   data_element                  := data_handler.Construct_Data_Structures(&data_search_list)
   redis_topic_value             = (*data_element)["TOPIC_VALUE"].(redis_handlers.Redis_Hash_Struct)
   redis_topic_time_stamp        = (*data_element)["TOPIC_TIME_STAMP"].(redis_handlers.Redis_Hash_Struct)
   redis_topic_handler           = (*data_element)["TOPIC_HANDLER"].(redis_handlers.Redis_Hash_Struct)
   redis_topic_error_ts          = (*data_element)["TOPIC_ERROR_TIME_STAMP"].(redis_handlers.Redis_Hash_Struct)
   postges_topic_stream          = (*data_element)["POSTGRES_DATA_STREAM"].(pg_drv.Postgres_Stream_Driver)
 
   
   
   redis_topic_handler.Delete_All()
   redis_topic_time_stamp.Delete_All()
   redis_topic_value.Delete_All()
   
   redis_topic_error_ts.Delete_All()
  
   
}

   


func construct_topic_handlers(){
   
   for class_name,element := range class_map{
       for _, instance_name := range element.instance_list {
           for _, topic_name := range element.topic_list {
               topic_handler := topic_map[topic_name].handler_type
               full_topic := base_topic_string+class_name+"/"+instance_name+"/"+topic_name
               redis_topic_handler.HSet(full_topic,topic_handler)
               redis_topic_time_stamp.HSet(full_topic,"0")
           }
       }
   }
    
}




