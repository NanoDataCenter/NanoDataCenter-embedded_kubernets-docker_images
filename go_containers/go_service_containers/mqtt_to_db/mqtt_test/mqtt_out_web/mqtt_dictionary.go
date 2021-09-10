package mqtt_test_web





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


var topic_map map[string]topic_type
var class_map map[string]class_type
var device_map map[string]device_type

func register_topics( topic_interface map[string]map[string]interface{}){
   topic_map = make(map[string]topic_type)
   for key ,element := range topic_interface{
      var item topic_type
      item.name         = element["name"].(string)
      item.description  = element["description"].(string)
      item.handler_type = element["handler_type"].(string)
      
      topic_map[key] = item
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
