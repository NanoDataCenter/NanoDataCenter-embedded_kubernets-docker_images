package weather_station_problems


import (
  //"fmt"
  //"os"
  //"strings"  
  //"html/template"
  //"net"
  "sort"
  "net/http"
  "lacima.com/redis_support/redis_handlers" 
  "lacima.com/Patterns/web_server_support/jquery_react_support"
  "html/template"
)

var  base_templates     *template.Template
var  eto_exception      redis_handlers.Redis_Hash_Struct

func Init(templates *template.Template, exception redis_handlers.Redis_Hash_Struct ){
    base_templates    = templates
    eto_exception     = exception
    
}

func Generate_page(w http.ResponseWriter, r *http.Request){
    html := generate_html()
    working_template,_ := base_templates.Clone()
    template.Must(working_template.New("application").Parse(html))
    data := make(map[string]interface{})
    data["Title"] = "ETO Daily Values"
    working_template.ExecuteTemplate(w,"bootstrap", data)    
    
}

func generate_html()string {
    data := eto_exception.HGetAll()
    keys := eto_exception.HKeys()
    sort.Strings(keys)
    display_list := make([][]string,len(keys))
    for index,key := range keys {
       problem    := data[key]
       
       display_list[index] = []string{key,problem}  
       
    }
    title := "<CENTER><H3>ETO STATION PROBLEMS</CENTER><BR>"
    return title + web_support.Setup_data_table("ETO STATION Problems",[]string{"Station","Problem"},display_list)
}
