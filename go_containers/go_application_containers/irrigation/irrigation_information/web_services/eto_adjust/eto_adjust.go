package eto_adjust

import (
    //"fmt"
    //"strings"
    "net/http"
    "html/template"
    //"encoding/json"
    //"lacima.com/redis_support/generate_handlers"
	//"lacima.com/redis_support/graph_query"
   // "lacima.com/redis_support/redis_handlers"
  
    //"lacima.com/Patterns/web_server_support/jquery_react_support"
    //"lacima.com/Patterns/msgpack_2"
    //"github.com/vmihailenco/msgpack/v5"
)

var base_templates        *template.Template

func Page_init(input *template.Template){
    //initialize_eto_data_structures()
    base_templates = input
    //file_loader = irr_files.Initialization()
    //setup_eto_handlers()
    //import_eto_setup_file()
    
}



func Generate_page_adjust(w http.ResponseWriter, r *http.Request){
    
   page_template ,_ := base_templates.Clone()
    page_html := generate_eto_manage_html()
    template.Must(page_template.New("application").Parse(page_html))        
    data := make(map[string]interface{})
    data["Title"] = "Modify ETO Resources"
    page_template.ExecuteTemplate(w,"bootstrap", data)
    
}

func generate_eto_manage_html()string{
    return "<H1>ETO adjust</H1>"
}
