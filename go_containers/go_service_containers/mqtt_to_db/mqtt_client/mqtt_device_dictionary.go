package mqtt_client

import "lacima.com/redis_support/graph_query"


type topic_type struct {
    name          string
    description   string
    handler_type  string
}

type class_type struct {
   
    name          string
    description   string
    device_list   []string
    topic_list    []string
    contact_time  int64
    
}

type device_type struct {
    name        string
    description string
    class       string
}


var class_map    map[string]class_type
var topic_map    map[string]topic_type
var device_map   map[string]device_type
var handler_map  map[string]string



func construct_mqtt_device_enviroment(){
    
    
   data_nodes := graph_query.Common_qs_search(&[]string{"MQTT_DEVICES:MQTT_DEVICES"})
   data_node  := data_nodes[0]

   class_dict :=  graph_query.Convert_json_nested_dictionary_interface(data_node["classes"])
   device_dict := graph_query.Convert_json_nested_dictionary_interface(data_node["devices"])
   topic_dict   := graph_query.Convert_json_nested_dictionary_interface(data_node["topics"])

   register_topics(topic_dict)
   register_classes(class_dict)
   register_devices(device_dict)
   
}




func register_topics( topic_interface map[string]map[string]interface{}){
   topic_map    = make(map[string]topic_type)
   handler_map  = make(map[string]string)
   for key ,element := range topic_interface{
      var item topic_type
      item.name         = element["name"].(string)
      item.description  = element["description"].(string)
      item.handler_type = element["handler_type"].(string)
      handler_map[key]  = item.handler_type
      topic_map[key]    = item
   }
}

func generate_list_array( input []interface{})[]string{
    return_value := make([]string,len(input))
    for index,value := range input {
        return_value[index] = value.(string)
    }
    return return_value
}

func register_classes( class_interface map[string]map[string]interface{}){
   class_map = make(map[string]class_type)
   
   for key ,element := range class_interface{
      
      var item class_type
      item.name         = element["name"].(string)
      item.description  = element["description"].(string)
      item.topic_list   = generate_list_array(element["topic_list"].([]interface{}))
      item.device_list  = generate_list_array(element["device_list"].([]interface{}))
      item.contact_time = int64(element["contact_time"].(float64))
      class_map[key] = item
   }
}



func register_devices( topic_interface map[string]map[string]interface{}){
   device_map = make(map[string]device_type)
   for key ,element := range topic_interface{
      var item device_type
      item.name         = element["name"].(string)
      item.description  = element["description"].(string)
      item.class        = element["class"].(string)
      device_map[key] = item
   }
}


