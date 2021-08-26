package mqtt_support


import (
    "fmt"
    
   
    "strings"
    "net/http"
    "html/template"
    "lacima.com/Patterns/web_server_support/jquery_react_support"
    
    //"github.com/msgpack/msgpack-go"

)





type base_type struct{
    working_template *template.Template
    html              string
    title             string
}
    
type class_page_type struct{
    base_type
}
    
    

func (v *base_type)base_init(title string){
    v.title = title
}

func (v  *base_type)basic_generate( html string, w http.ResponseWriter, r *http.Request){
   v.working_template,_ = base_templates.Clone()
   fmt.Println("working template",v.working_template)
   fmt.Println("html",html)
   template.Must(v.working_template.New("application").Parse(html))
   data := make(map[string]interface{})
   data["Title"] = v.title
   v.working_template.ExecuteTemplate(w,"bootstrap", data)    
    
}

    
func ( v *class_page_type)generate_introduction()string{
    
    return  "<center><h3>Display of MQTT Device Class</h3> </center><br>"
}


func (v *class_page_type)generate_class_header(element class_type)string{
 
    return fmt.Sprintln("<center> Class %s %s timeout %d ",element.name,element.description,element.contact_time)
    
}

func (v *class_page_type)assemble_topic_elements( topic_list []string)web_support.Accordion_Elements{
    
    var  return_value web_support.Accordion_Elements
    
    return_value.Title = "List of Topics"
    text_array  := make([]string,len(topic_list))
    for index,value := range topic_list {
        topic_element := topic_map[value]
        text_array[index] = fmt.Sprintf("Name %s Description %s ",topic_element.name,topic_element.description)
    }
    return_value.Body = strings.Join(text_array,"<br>")
    return return_value
        


}

func (v *class_page_type)assemble_device_name( device_list []string)web_support.Accordion_Elements{
    
   var  return_value web_support.Accordion_Elements
    
    return_value.Title = "List of Devices"
    text_array  := make([]string,len(device_list))
    for index,value := range device_list {
        device_element := device_map[value]
        text_array[index] = fmt.Sprintf("Name: %s Description %s ",device_element.name,device_element.description)
    }
    return_value.Body = strings.Join(text_array,"<br>")
    return return_value
        


}


func (v *class_page_type)generate_class_element(key string, element class_type)string{
 
    accordion_elements := make([]web_support.Accordion_Elements,2)    
    accordion_elements[0] = v.assemble_topic_elements(element.topic_list)
    accordion_elements[1] = v.assemble_device_name(element.device_list)
    
    
    title := fmt.Sprintf("Class: %s  ",element.name)
    return web_support.Generate_accordian(key+"_class",title,  accordion_elements ) 
            
}

func (v *class_page_type)generate_html()string{
    return_array := make([]string,len(class_map)+1)
    index := 0
    return_array[index] = v.generate_introduction()
    index = index +1
    for key, element := range class_map {
       return_array[index] = v.generate_class_element(key,element)
       index = index +1
    }
    return strings.Join(return_array,"<br>")
}   



func (v *class_page_type)generate_page(w http.ResponseWriter, r *http.Request){
    html  := v.generate_html()
    v.basic_generate(html,w,r)
}

func (v *class_page_type)init_page(){
    v.base_init("List Classes")
}

