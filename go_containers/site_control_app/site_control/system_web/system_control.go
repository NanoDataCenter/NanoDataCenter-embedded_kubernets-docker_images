package site_web_server

import (

    "net/http"
    "html/template"
   //"lacima.com/Patterns/web_server_support/jquery_react_support"
)




var system_control_template *template.Template

var system_control_html = `
<div class="container">
  <div class="jumbotron">
    <h1>System Control</h1>
    <p>System Control.</p>
  </div>
</div>
`


func system_control_init(){
    system_control_template ,_ = base_templates.Clone()
    
    template.Must(system_control_template.New("application").Parse(system_control_html))
    
    
}    



func system_control(w http.ResponseWriter, r *http.Request) {
   data := make(map[string]interface{})
   data["Title"] = "System Control"
   system_control_template.ExecuteTemplate(w,"bootstrap", data)
   
    
    
}








