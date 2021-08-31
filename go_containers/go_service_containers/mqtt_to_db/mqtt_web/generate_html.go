package mqtt_web

import "strings"
import "sort"
import "fmt"
import  "lacima.com/Patterns/web_server_support/jquery_react_support"


func (v *class_page_type)generate_html()string{
    return_array := make([]string,len(class_map)+1)
    index := 0
    //return_array[index] = v.generate_introduction()
    index = index +1
    for key, element := range class_map {
       return_array[index] = v.generate_class_element(key,element)
       index = index +1
    }
    return strings.Join(return_array,"<br>")
} 


func (v *class_page_type)generate_class_element(key string, element class_type)string{
 
    accordion_elements := make([]web_support.Accordion_Elements,2)    
    accordion_elements[0] = v.assemble_topic_elements(element.topic_list)
    accordion_elements[1] = v.assemble_device_name(element.device_list)
    
    
    title := fmt.Sprintf("Class: %s Description: %s  Contact Time: %d  ",element.name, element.description, element.contact_time)
    return web_support.Generate_accordian(key+"_class",title,  accordion_elements ) 
            
}

func (v *class_page_type)assemble_topic_elements( topic_list []string)web_support.Accordion_Elements{
    
    var  return_value web_support.Accordion_Elements
    
    return_value.Title = "List of Topics"
    text_array  := make([][]string,len(topic_list))
    for index,value := range topic_list {
        topic_element := topic_map[value]
        text_array[index] = []string{topic_element.name,topic_element.description, topic_element.handler_type}
    }
    return_value.Body = web_support.Setup_data_table("topic_tag" , []string{"Name","Description","Handler"},text_array )
    
    return return_value
        


}

func (v *class_page_type)assemble_device_name( device_list []string)web_support.Accordion_Elements{
    
   var  return_value web_support.Accordion_Elements
    
    return_value.Title = "List of Devices"
    text_array  := make([][]string,len(device_list))
    for index,value := range device_list {
        device_element := device_map[value]
        
        text_array[index] = []string{device_element.name,device_element.description, device_element.class}
    }
    return_value.Body = web_support.Setup_data_table("topic_tag" , []string{"Name","Description","Class"},text_array )
    return return_value
        


}


func (v *topic_map_page_type)generate_html()string {
    topic_map := redis_topic_time_stamp.HGetAll()
    topic_keys := redis_topic_time_stamp.HKeys()
    sort.Strings(topic_keys)
    display_list := make([]string,len(topic_keys))
    for index,key := range topic_keys {
        
       display_list[index] = fmt.Sprintf("Topic: %s Contact Time %s ",key,topic_map[key]) 
    }
    return web_support.Generate_list_link("topic_list","<center>Topic Map Display </centr>",display_list)
}


func (v *device_status_page_type)generate_html()string {
    return "Device Status Page"
}

func (v *bad_topic_page_type)generate_html()string {
    return "BAD Topic Page"
}

func (v *recent_mqtt_activitiy_page_type)generate_html()string {
    return "MQTT ACTIVITY"
}

func (v *mqtt_inicident_page_type)generate_html()string {
    return "MQTT INCIDENT PAGE"
}




