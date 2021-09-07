package mqtt_monitor_devices

import "fmt"

import "time"
import "strconv"
import "lacima.com/redis_support/graph_query"
import "lacima.com/server_libraries/postgres"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"

type class_type struct {
   
    name          string
    device_list   []string
    contact_time  int64
    
}

type device_type struct {
    name          string
    class         string
    contact_time  int64
}


var class_map    map[string]class_type
var device_map   map[string]device_type



var postgres_incident_stream    pg_drv.Postgres_Stream_Driver      // time stream for all device changes
                                                       // tag1 class
                                                       // tag2 device
                                                       // tag3 status
                                                       // tag4 handler 
                                                       // date time string
                                                       // data msgpack data
                                                       
var redis_device_status          redis_handlers.Redis_Hash_Struct
var redis_contact_time           redis_handlers.Redis_Hash_Struct


var device_list []string

func Monitor_int(){
    
   data_search_list              := []string{ "MQTT_IN_SETUP:mqtt_in_setup","TOPIC_STATUS"}
   data_element                  := data_handler.Construct_Data_Structures(&data_search_list)
   redis_device_status           = (*data_element)["DEVICE_STATUS"].(redis_handlers.Redis_Hash_Struct)
   redis_contact_time            = (*data_element)["DEVICE_TIME_STAMP"].(redis_handlers.Redis_Hash_Struct)
   postgres_incident_stream      = (*data_element)["POSTGRES_INCIDENT_STREAM"].(pg_drv.Postgres_Stream_Driver)
   device_list                   = redis_device_status.HKeys()
  
   construct_mqtt_device_enviroment()
   
   mark_all_devices_true() 
   
    
}


func construct_mqtt_device_enviroment(){
    
    
   data_nodes := graph_query.Common_qs_search(&[]string{"MQTT_IN_SETUP:mqtt_in_setup","MQTT_DEVICES:MQTT_DEVICES"})
   data_node  := data_nodes[0]

   class_dict :=  graph_query.Convert_json_nested_dictionary_interface(data_node["classes"])
   device_dict := graph_query.Convert_json_nested_dictionary_interface(data_node["devices"])
   

   register_classes(class_dict)
   register_devices(device_dict)
   
}


func register_classes( class_interface map[string]map[string]interface{}){
   class_map = make(map[string]class_type)
   
   for key ,element := range class_interface{
      
      var item class_type
      item.name           = element["name"].(string)
      item.device_list    = generate_list_array(element["device_list"].([]interface{}))
      item.contact_time   = int64(element["contact_time"].(float64))
      class_map[key]      = item
   }
}



func register_devices( topic_interface map[string]map[string]interface{}){
   device_map = make(map[string]device_type)
   for key ,element := range topic_interface{
       
       
      var item device_type
      item.name              = element["name"].(string)
      item.class             = element["class"].(string)
      item.contact_time      = class_map[item.class].contact_time
      
      device_map[key] = item
   }
}

func generate_list_array( input []interface{})[]string{
    return_value := make([]string,len(input))
    for index,value := range input {
        return_value[index] = value.(string)
    }
    return return_value
}  
    


func mark_all_devices_true(){
   for _,device := range device_list {
     
      postgres_incident_stream.Insert(device_map[device].class,device,"true","","","")
     
   }
    
    
}
func Monitor_devices(){
    for true {
        check_all_devices()
        time.Sleep(time.Second *15)
    }
}



func check_all_devices(){
    time_stamp := time.Now().Unix()
    for device, item := range device_map {
        
         last_contact_string := redis_contact_time.HGet(device)
         last_contact , _ := strconv.Atoi(last_contact_string)
         time_out_value := int64(last_contact) + item.contact_time
         //fmt.Println("device",device,last_contact_string,last_contact,time_out_value,time_stamp)
         if time_stamp > time_out_value {
            if  redis_device_status.HGet(device) == "true" {
              
              fmt.Println("device",device,"false")
              postgres_incident_stream.Insert(device_map[device].class,device,"false","","","") 
              redis_device_status.HSet(device,"false")

            }
            

         }else {
            if redis_device_status.HGet(device) == "false" {
               
                fmt.Println("device",device,"true")
                postgres_incident_stream.Insert(device_map[device].class,device,"true","","","")
                redis_device_status.HSet(device,"true")
            }
            
          
         }
    }
}
        

