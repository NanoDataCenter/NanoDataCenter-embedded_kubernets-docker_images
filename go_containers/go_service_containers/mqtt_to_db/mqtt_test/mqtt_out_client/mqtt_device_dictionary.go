package mqtt_out_client

import "lacima.com/redis_support/graph_query"


type topic_type struct {
    name          string
    description   string
    handler_type  string
}

type class_type struct {
   
    name            string
    description     string
    instance_list   []string
    topic_list      []string
   
    
}

type instance_type struct {
    name        string
    description string
    class       string
}


var class_map      map[string]class_type
var topic_map      map[string]topic_type
var instance_map   map[string]instance_type
var handler_map    map[string]string



func construct_mqtt_device_enviroment(){
    
    
   data_nodes := graph_query.Common_qs_search(&[]string{"MQTT_OUTPUT_SETUP:site_out_server","MQTT_INSTANCES:MQTT_INSTANCES"})
   data_node  := data_nodes[0]

   class_dict       :=  graph_query.Convert_json_nested_dictionary_interface(data_node["classes"])
   instance_dict    := graph_query.Convert_json_nested_dictionary_interface(data_node["instances"])
   topic_dict       := graph_query.Convert_json_nested_dictionary_interface(data_node["topics"])

   register_topics(topic_dict)
   register_classes(class_dict)
   register_instances(instance_dict)
   
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
      item.name           = element["name"].(string)
      item.description    = element["description"].(string)
      item.instance_list  = generate_list_array(element["instance_list"].([]interface{} ))
      item.topic_list     = generate_list_array(element["topic_list"].([]interface{}))
      class_map[key]      = item
   }
}



func register_instances( topic_interface map[string]map[string]interface{}){
   instance_map = make(map[string]instance_type)
   for key ,element := range topic_interface{
      var item instance_type
      item.name         = element["name"].(string)
      item.description  = element["description"].(string)
      item.class        = element["class"].(string)
      instance_map[key] = item
   }
}


