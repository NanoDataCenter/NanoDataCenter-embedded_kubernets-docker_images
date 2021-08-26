package site_web_server


import (
    //"fmt"
    
   
    //"strconv"
    "net/http"
    "html/template"
    "lacima.com/Patterns/web_server_support/jquery_react_support"
    
    //"github.com/msgpack/msgpack-go"

)

var class_page_template            *template.Template
var device_page_template           *template.Template


var mqtt_incident_page_template    *template.Template

    
func class_page_init(){
    class_page_template ,_   = base_templates.Clone()
    class_page_html          := web_support.Generate_accordian("intro_1","Description of Web Pages",  generate_intro_data())
    template.Must(class_page_template.New("application").Parse(class_page_html))
}   
    
func class_page_generate(w http.ResponseWriter, r *http.Request) {
   
   data := make(map[string]interface{})
   data["Title"] = "Introduction"
   class_page_template.ExecuteTemplate(w,"bootstrap", data)
}    
    
    
func device_page_init(){
    introduction_page_template ,_ = base_templates.Clone()
    introduction_page_html := web_support.Generate_accordian("intro_1","Description of Web Pages",  generate_intro_data())
    template.Must(class_page_template.New("application").Parse(introduction_page_html))    
    
}    
    
    
func device_page_generate(w http.ResponseWriter, r *http.Request) {
   
   data := make(map[string]interface{})
   data["Title"] = "Introduction"
   introduction_page_template.ExecuteTemplate(w,"bootstrap", data)
}   
    


func topic_page_init(){
    introduction_page_template ,_ = base_templates.Clone()
    introduction_page_html := web_support.Generate_accordian("intro_1","Description of Web Pages",  generate_intro_data())
    template.Must(introduction_page_template.New("application").Parse(introduction_page_html))
    
    
}    
    
 func topic_page_generate(w http.ResponseWriter, r *http.Request) {
   
   data := make(map[string]interface{})
   data["Title"] = "Introduction"
   introduction_page_template.ExecuteTemplate(w,"bootstrap", data)
}   
   


func topic_map_init(){
     introduction_page_template ,_ = base_templates.Clone()
    introduction_page_html := web_support.Generate_accordian("intro_1","Description of Web Pages",  generate_intro_data())
    template.Must(introduction_page_template.New("application").Parse(introduction_page_html))
    
    
}    

func topic_map_generate(w http.ResponseWriter, r *http.Request) {
   
   data := make(map[string]interface{})
   data["Title"] = "Introduction"
   introduction_page_template.ExecuteTemplate(w,"bootstrap", data)
}   

func device_incident_init(){
    introduction_page_template ,_ = base_templates.Clone()
    introduction_page_html := web_support.Generate_accordian("intro_1","Description of Web Pages",  generate_intro_data())
    template.Must(introduction_page_template.New("application").Parse(introduction_page_html))
    
    
}    

func device_incident_generate(w http.ResponseWriter, r *http.Request) {
   
   data := make(map[string]interface{})
   data["Title"] = "Introduction"
   introduction_page_template.ExecuteTemplate(w,"bootstrap", data)
}   




func mqtt_incident_init(){
     introduction_page_template ,_ = base_templates.Clone()
    introduction_page_html := web_support.Generate_accordian("intro_1","Description of Web Pages",  generate_intro_data())
    template.Must(introduction_page_template.New("application").Parse(introduction_page_html))
    
    
}    

func mqtt_incident_generate(w http.ResponseWriter, r *http.Request) {
   
   data := make(map[string]interface{})
   data["Title"] = "Introduction"
   mqtt_incident_page_template.ExecuteTemplate(w,"bootstrap", data)
}   









  





     
        
 
    





