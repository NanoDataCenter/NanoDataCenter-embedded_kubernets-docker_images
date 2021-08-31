package mqtt_web



import (
    
    
    "os"
  
    
    "net/http"
    "html/template"
    "lacima.com/Patterns/web_server_support/jquery_react_support"

    //"github.com/go-redis/redis/v8"
)

var server_id                   string
var base_templates                *template.Template
var introduction_page_template    *template.Template
var class_page                     class_page_type
var topic_map_page                 topic_map_page_type
var device_status_page             device_status_page_type
var bad_topic_page                 bad_topic_page_type
var recent_mqtt_activitiy_page     recent_mqtt_activitiy_page_type
var mqtt_inicident_page            mqtt_inicident_page_type


 
func Init_site_web_server(){
   

   server_id  =  os.Getenv("SERVER_ID")
   if server_id == "" {
       panic("bad server id ")
   }
   web_support.Register_web_page_start(server_id)
   init_web_server_pages()
   web_support.Launch_web_server()
}
  

func init_web_server_pages() {

    web_support.Init_web_support(introduction_page_generate)  // register page
    base_templates = define_web_pages()
    initialize_handlers()
   
}

func initialize_handlers(){
 
    introduction_page_init()
    class_page.init_page()  
    topic_map_page.init_page()
    device_status_page.init_page()
    bad_topic_page.init_page()
    recent_mqtt_activitiy_page.init_page()
    mqtt_inicident_page.init_page()
    web_support.Micro_web_page_init(base_templates)
}


    



func define_web_pages()*template.Template  {
 
    return_value := make(web_support.Menu_array,8)
    return_value[0] = web_support.Construct_Menu_Element( "introduction page","introduction_page",introduction_page_generate)
    return_value[1] = web_support.Construct_Menu_Element( "Class page","class_page", class_page.generate_page)
    return_value[2] = web_support.Construct_Menu_Element( "Topic Map","topic_map", topic_map_page.generate_page)
    return_value[3] = web_support.Construct_Menu_Element( "Device Status Page","device_status_page", device_status_page.generate_page)
    return_value[4] = web_support.Construct_Menu_Element( "Bad Topic Page ","bad_topic_page", bad_topic_page.generate_page)
    return_value[5] = web_support.Construct_Menu_Element( "Recent MQTT History","recent_mqtt_activity", recent_mqtt_activitiy_page.generate_page)
    return_value[6] = web_support.Construct_Menu_Element( "MQTT Incidents","mqtt_incidents", mqtt_inicident_page.generate_page)
    return_value[7] = web_support.Construct_Menu_Element( "application_servers","application_servers", web_support.Micro_web_page)

    web_support.Register_web_pages(return_value)
    return web_support.Generate_single_row_menu(return_value)
}
    




func introduction_page_init( ){
    introduction_page_template ,_ = base_templates.Clone()
    introduction_page_html := web_support.Generate_Introduction()+ web_support.Generate_accordian("intro_1","Description of Web Pages",  generate_intro_data())
    template.Must(introduction_page_template.New("application").Parse(introduction_page_html))
    
    
}    



func introduction_page_generate(w http.ResponseWriter, r *http.Request) {
   
   data := make(map[string]interface{})
   data["Title"] = "Introduction"
   introduction_page_template.ExecuteTemplate(w,"bootstrap", data)
}   



 

const application_server_body string = `
This web page lists all web servers Relating to Site Micro Services<br><br>

Clink the the link opens Web Page for the Micro Service in a separate table.`



const class_page_body  string = `
This web page lists all the Classes of MQTT devices and their respective Properteis.`


const device_page_body  string = `
This web page lists all the devices in the MQTT system and 
their respective properties as well as the status of the device
`


const topic_page_body  string = `
This web page lists all the topics which the MQTT devices use as well
as the properties of the topics `


const master_topic_page_body  string = `
This web page lists the expanded topic map of the MQTT system.  
The design of the topic space is /<site_name>/class_name/device_name/+
the topic which the MQTT device sends.  

The topic fields are sorted alphabetically and the date and time of the latest time 
stamp is shown.
`



const device_incident_log_page_body string = `
This page shows the incident log for the various MQTT devices.
The incident log show when a device dropped off and when the device came back on line`


const mqtt_client_incident_page_body string = `
This page tracks the connections and disconnections of the MQTT server`



    

    
    
    
func generate_intro_data()[]web_support.Accordion_Elements{

  title_array := []string{"Class Page", "Device Page", "Topic Page","Master Topic Page","Device Incident Log","MQTT CLient Incident Page","Application Servers"}
  body_array  := []string{application_server_body,
                          class_page_body,
                          device_page_body,
                          topic_page_body,
                          master_topic_page_body,
                          device_incident_log_page_body,
                          mqtt_client_incident_page_body}
                          
  return web_support.Populate_accordian_elements(title_array,body_array)
    
    
}






