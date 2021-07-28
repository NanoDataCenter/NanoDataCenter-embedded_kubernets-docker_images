package site_web_server

import (
   
    "net/http"
    "html/template"
    "lacima.com/Patterns/web_server_support/jquery_react_support"
    
)







func node_status_init(){
     ; // fetch data handlers
    
}    



func node_status(w http.ResponseWriter, r *http.Request) {
   
   node_status_template ,_ := base_templates.Clone()
   
   display_data := generate_node_status_data()
   node_status_html := web_support.Generate_list_link( "container_1","Node_Status", display_data )
   template.Must(node_status_template.New("application").Parse(node_status_html))
 
    
   data := make(map[string]interface{})
   data["Title"] = "Node Status"
   node_status_template.ExecuteTemplate(w,"bootstrap", data)
   
    
    
}


func generate_node_status_data()[]string {
    
   return_value := make([]string,0)



   return return_value   
    
    
    
}
