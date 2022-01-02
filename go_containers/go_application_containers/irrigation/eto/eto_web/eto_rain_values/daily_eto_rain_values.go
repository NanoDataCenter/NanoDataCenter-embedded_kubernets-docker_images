package eto_rain_values


import (
  "fmt"
  //"os"
  //"strings"  
  //"html/template"
  //"net"
  "sort"
  "net/http"
  "lacima.com/Patterns/web_server_support/jquery_react_support"
   "lacima.com/go_application_containers/irrigation/eto/eto_calc/eto_support"
  "html/template"
)



var base_templates *template.Template  


func Init(templates *template.Template ){
    base_templates = templates
    
    
    
}




func Generate_page(w http.ResponseWriter, r *http.Request){
    html := generate_html()
    working_template,_ := base_templates.Clone()
    template.Must(working_template.New("application").Parse(html))
    data := make(map[string]interface{})
    data["Title"] = "ETO Daily Values"
    working_template.ExecuteTemplate(w,"bootstrap", data)    
    
}


func generate_html()string{
    return generate_ETO_html()+generate_Rain_html()
}

func generate_ETO_html()string {
    
    html_title := "<CENTER><H3>DAILY ETO VALUES</H3></CENTER><BR>"
    keys       := eto_support.ETO_HKeys()
    display_list := make([][]string,len(keys))
    sort.Strings(keys)
    for index,key := range keys {
       data := eto_support.ETO_HGet(key)
       priority_string := fmt.Sprintf("%.2f",data.Priority)
       eto_string := fmt.Sprintf("%.2f",data.Value)
       display_list[index] = []string{key,priority_string,eto_string}  
       
    }
    return html_title+web_support.Setup_data_table("ETO Data",[]string{"Station","Priority","ETO"},display_list)
}
    
func generate_Rain_html()string {
    
    html_title := "<CENTER><H3>DAILY Rain VALUES</H3></CENTER><BR>"
    keys       := eto_support.Rain_HKeys()
    display_list := make([][]string,len(keys))
    sort.Strings(keys)
    for index,key := range keys {
       data := eto_support.Rain_HGet(key)
       priority_string := fmt.Sprintf("%.2f",data.Priority)
       eto_string := fmt.Sprintf("%.2f",data.Value)
       display_list[index] = []string{key,priority_string,eto_string}  
       
    }
    return html_title+web_support.Setup_data_table("ETO Data",[]string{"Station","Priority","Rain"},display_list)
}
        
