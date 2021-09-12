package mqtt_web

import "fmt"
import "strings"
import "sort"
import "time"
import "strconv"

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
    accordion_elements[0] = v.assemble_topic_elements(element.name+"topic",element.topic_list)
    accordion_elements[1] = v.assemble_device_name(element.name+"device",element.device_list)
    
    
    title := fmt.Sprintf("Class: %s Description: %s  Contact Time: %d  ",element.name, element.description, element.contact_time)
    return web_support.Generate_accordian(key+"_class",title,  accordion_elements ) 
            
}

func (v *class_page_type)assemble_topic_elements( tag string,topic_list []string)web_support.Accordion_Elements{
    
    var  return_value web_support.Accordion_Elements
    
    return_value.Title = "List of Topics"
    text_array  := make([][]string,len(topic_list))
    for index,value := range topic_list {
        topic_element := topic_map[value]
        text_array[index] = []string{topic_element.name,topic_element.description, topic_element.handler_type}
    }
    return_value.Body = web_support.Setup_data_table(tag , []string{"Name","Description","Handler"},text_array )
    
    return return_value
        


}

func (v *class_page_type)assemble_device_name(tag string, device_list []string)web_support.Accordion_Elements{
    
   var  return_value web_support.Accordion_Elements
    
    return_value.Title = "List of Devices"
    text_array  := make([][]string,len(device_list))
    for index,value := range device_list {
        device_element := device_map[value]
        
        text_array[index] = []string{device_element.name,device_element.description, device_element.class}
    }
    return_value.Body = web_support.Setup_data_table(tag  , []string{"Name","Description","Class"},text_array )
    return return_value
        


}


func (v *topic_map_page_type)generate_html()string {
    topic_map := redis_topic_time_stamp.HGetAll()
    topic_keys := redis_topic_time_stamp.HKeys()
    sort.Strings(topic_keys)
    display_list := make([][]string,len(topic_keys))
    for index,key := range topic_keys {
       last_contact , _ := strconv.Atoi(topic_map[key])
       time_stamp := time.Unix(int64(last_contact),0)
       display_list[index] = []string{key,time_stamp.Format(time.UnixDate)}  
       
    }
    
    return web_support.Setup_data_table("topic_list",[]string{"Topic","Contact Time"},display_list)
}


func (v *device_status_page_type)generate_html()string {
    topic_map := redis_device_status.HGetAll()
    topic_keys := redis_device_status.HKeys()
    sort.Strings(topic_keys)
    display_list := make([][]string,len(topic_keys))
    for index,key := range topic_keys {
        
       display_list[index] = []string{key,topic_map[key]} 
    }
    
    return web_support.Setup_data_table("topic_list",[]string{"Device","Status"},display_list)
}

func (v *bad_topic_page_type)generate_html()string {
    topic_map := redis_topic_error_ts.HGetAll()
    topic_keys := redis_topic_error_ts.HKeys()
    sort.Strings(topic_keys)
    display_list := make([][]string,len(topic_keys))
    for index,key := range topic_keys {
       time_value,_ := strconv.Atoi(topic_map[key])
       time_stamp := time.Unix(int64(time_value),0)
       
       display_list[index] = []string{key,time_stamp.Format(time.UnixDate)} 
    }
    
    return web_support.Setup_data_table("topic_list",[]string{"Topic","Contact Time"},display_list)
}


func (v *recent_mqtt_activitiy_page_type)generate_html()string {
    postgres_data,_ := postges_topic_stream.Select_after_time_stamp_desc(3600) // one hour
    display_list := make([][]string,len(postgres_data))
    for index,data := range postgres_data {
     
      time_sec  := data.Time_stamp / 1e9
      time_nsec := data.Time_stamp % 1e9
    
       time_stamp := time.Unix(time_sec ,time_nsec)
       
       stream_id_string := strconv.FormatInt(data.Stream_id,10) 
       display_list[index] = []string{stream_id_string,data.Tag1,data.Tag2,data.Tag3,data.Tag4,time_stamp.Format(time.UnixDate)} 
    }

    return web_support.Setup_data_table("topic_list",[]string{"Stream ID","Class","Device","Topic","Handler","Time"},display_list)
}
 
  
func (v *device_off_line_incidents_page_type)generate_html()string {
   
    
    postgres_data,_ := postgres_incident_stream.Select_after_time_stamp_desc(3600)

    display_list := make([][]string,len(postgres_data))
    for index,data := range postgres_data {
       t :=  time.Unix(data.Time_stamp/1000000000,0)
       string_date := t.Format(time.UnixDate) 
       display_list[index] = []string{data.Tag1,data.Tag2,data.Tag3,string_date} 
    }
    
    return web_support.Setup_data_table("topic_list",[]string{"Class","Device","Status","Time"},display_list)
}






func (v *sys_history_page_type )generate_html()string {
        
    postgres_data,_ := postgres_sys_stream.Select_after_time_stamp_desc(3600)

    display_list := make([][]string,len(postgres_data))
    for index,data := range postgres_data {
       t :=  time.Unix(data.Time_stamp/1000000000,0)
       string_date := t.Format(time.UnixDate) 
       display_list[index] = []string{data.Tag1,data.Tag2,data.Tag3,data.Tag4,string_date,data.Data} 
    }
    
    return web_support.Setup_data_table("topic_list",[]string{"Topic","Tag1","Tag2","Tag3","Time","DATA"},display_list)
}


