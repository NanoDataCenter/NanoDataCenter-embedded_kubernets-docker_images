package mqtt_web


import (
    "fmt"
    
   
    //"strings"
    "net/http"
    "html/template"
    //"lacima.com/Patterns/web_server_support/jquery_react_support"
    
    //"github.com/msgpack/msgpack-go"

)

/*
 * 
 * class display web site
 * 
 * 
 * 
 * 
 */



type base_type struct{
    working_template *template.Template
    html              string
    title             string
}
    
type class_page_type struct{
    base_type
}
 
type topic_map_page_type struct{
    base_type
}

type device_status_page_type struct{
    base_type
}

type bad_topic_page_type struct{
    base_type
}

type recent_mqtt_activitiy_page_type struct{
    base_type
}

type mqtt_inicident_page_type struct{
    base_type
}


func (v *base_type)base_init(title string){
    v.title = title
}

func (v  *base_type)basic_generate( html string, w http.ResponseWriter, r *http.Request){
   v.working_template,_ = base_templates.Clone()
   template.Must(v.working_template.New("application").Parse(html))
   data := make(map[string]interface{})
   data["Title"] = v.title
   v.working_template.ExecuteTemplate(w,"bootstrap", data)    
    
}


/*
 * 
 *  class page setup
 * 
 * 
 * 
 */

func (v *class_page_type)init_page(){
    fmt.Println("made it init")
    v.base_init("List Classes")
}

 
func (v *class_page_type)generate_page(w http.ResponseWriter, r *http.Request){
    fmt.Println("geneate page")
    html  := v.generate_html()
    v.basic_generate(html,w,r)
}




func (v *topic_map_page_type)init_page(){
    v.base_init("List of Valid Topics")
}

 
func (v *topic_map_page_type)generate_page(w http.ResponseWriter, r *http.Request){
    html  := v.generate_html()
    v.basic_generate(html,w,r)
}



func (v *device_status_page_type)init_page(){
    v.base_init("Device Status List")
}

 
func (v *device_status_page_type)generate_page(w http.ResponseWriter, r *http.Request){
    html  := v.generate_html()
    v.basic_generate(html,w,r)
}


func (v *bad_topic_page_type)init_page(){
    v.base_init("Bad Topics")
}

 
func (v *bad_topic_page_type)generate_page(w http.ResponseWriter, r *http.Request){
    html  := v.generate_html()
    v.basic_generate(html,w,r)
}

func (v *recent_mqtt_activitiy_page_type)init_page(){
    v.base_init("Recent MQTT Activity")
}

 
func (v *recent_mqtt_activitiy_page_type)generate_page(w http.ResponseWriter, r *http.Request){
    html  := v.generate_html()
    v.basic_generate(html,w,r)
}



func (v *mqtt_inicident_page_type)init_page(){
    v.base_init("List Classes")
}

 
func (v *mqtt_inicident_page_type)generate_page(w http.ResponseWriter, r *http.Request){
    html  := v.generate_html()
    v.basic_generate(html,w,r)
}

