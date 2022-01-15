package irrigation_schedules

import (
    "net/http"
    "html/template"
   // "lacima.com/Patterns/web_server_support/jquery_react_support"
)

var base_templates  *template.Template


func Page_init(input *template.Template){
    
    base_templates = input
    
    
    
}

func Generate_page(w http.ResponseWriter, r *http.Request){
    
    page_template ,_ := base_templates.Clone()
    page_html := generate_page_html()
    template.Must(page_template.New("application").Parse(page_html))     
    data := make(map[string]interface{})
    data["Title"] = "Setup Irrigation Schedules"
    page_template.ExecuteTemplate(w,"bootstrap", data)
    
}

func generate_page_html()string{
    return "<H1>Irrigation Setup</H1>"
}
