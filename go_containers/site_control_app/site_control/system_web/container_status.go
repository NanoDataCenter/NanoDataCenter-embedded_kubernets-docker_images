package site_web_server

import (
    
    "net/http"
    "html/template"
    "lacima.com/Patterns/web_server_support/jquery_react_support"
   
)









func container_status_init(){
   
   ;  // get data handlers 
    
}    



func container_status(w http.ResponseWriter, r *http.Request) {
   
   container_status_template ,_ := base_templates.Clone()
   
   display_data := generate_container_status_data()
   container_status_html := web_support.Generate_list_link( "container_1","Container_Status", display_data )
   
   template.Must(container_status_template.New("application").Parse(container_status_html))
   data := make(map[string]interface{})
   data["Title"] = "Container Status"
   container_status_template.ExecuteTemplate(w,"bootstrap", data)
   
    
    
}


func generate_container_status_data()[]string {
    
   return_value := make([]string,0)



   return return_value   
    
    
    
}






