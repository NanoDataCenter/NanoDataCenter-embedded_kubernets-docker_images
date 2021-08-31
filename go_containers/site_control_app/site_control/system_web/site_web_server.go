
package site_web_server

import (
    //"os"
    //"fmt"
    "net/http"
    "html/template"
    "lacima.com/Patterns/web_server_support/jquery_react_support"
)


var base_templates *template.Template
var introduction_page_template *template.Template

const site_id string = "site_controller"

func Init_site_web_server(){
   
   web_support.Register_web_page_start(site_id)
   init_web_server_pages()
   web_support.Launch_web_server()

}



func init_web_server_pages() {

    web_support.Init_web_support(introduction_page)
    base_templates = define_web_pages()
    initialize_handlers()
   
}







func define_web_pages()*template.Template  {
 
    return_value := make(web_support.Menu_array,5)
    

    return_value[0] = web_support.Menu_element{ "introduction page","introduction_page",introduction_page}
    return_value[1] = web_support.Menu_element{ "node status","node_status", node_status}
    return_value[2] = web_support.Menu_element{ "node ip","node_ip", node_ip}
    return_value[3] = web_support.Menu_element{ "container status","container_status",container_status}
    return_value[4] = web_support.Construct_Menu_Element( "application_servers","application_servers", web_support.Micro_web_page)
  
    
    web_support.Register_web_pages(return_value)
    return web_support.Generate_single_row_menu(return_value)
    
}












func initialize_handlers(){
 
    introduction_page_init()
    node_status_init()
    node_ip_init()
    container_status_init()
    web_support.Micro_web_page_init(base_templates)
    
    
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





func introduction_page(w http.ResponseWriter, r *http.Request) {
   
   data := make(map[string]interface{})
   data["Title"] = "Introduction"
   
   introduction_page_template.ExecuteTemplate(w,"bootstrap", data)
}   



 

const application_server_body string = `
This web page lists all web servers Relating to Site Micro Services<br><br>

Clink the the link opens Web Page for the Micro Service in a separate table.`


const node_status_body string = `
This web page lists all nodes in the system and whether the node is active.<br><br>
Active is defined whether the node orchestration node is responding to site controller`

const node_ip_body string = `
This web page lists the ip for all active nodes`

const container_status_body string = `
The web page list the status for all the containers in a system<br><br>

The status contain two parameters<br><br>
The first parameter is whether the container is running<br><br>
The second parameters is whether the container is managed by the node system.<br>  
Containers may be unmanged due to debugging operatons.<br>
Container management is manipulated by Ansible Debugging Scripts`



    

    

    
    
    
func generate_intro_data()[]web_support.Accordion_Elements{

  title_array := []string{"Node Status", "Node Ip", "Container Status","Application Server",}
  body_array  := []string{node_status_body,node_ip_body,container_status_body,application_server_body}
  return web_support.Populate_accordian_elements(title_array,body_array)
    
    
}




 



