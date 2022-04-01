package irrigation_manage_queue

import (
    //"fmt"
    //"strings"
    "net/http"
    "html/template"
    //"encoding/json"
   // "lacima.com/redis_support/generate_handlers"
	//"lacima.com/redis_support/graph_query"
  //  "lacima.com/redis_support/redis_handlers"
  
  //  "lacima.com/Patterns/web_server_support/jquery_react_support"
    //"lacima.com/Patterns/msgpack_2"
    //"github.com/vmihailenco/msgpack/v5"
)

var base_templates                    *template.Template


func Page_init(input *template.Template){
    //initialize_eto_data_structures()
    base_templates = input

    setup_irrigation_diagnostics()
    
    
}

func setup_irrigation_diagnostics() {

	;
}


func Generate_page_adjust(w http.ResponseWriter, r *http.Request){
    
  
   page_template ,_ := base_templates.Clone()
    page_html := generate_eto_manage()
    template.Must(page_template.New("application").Parse(page_html))        
    data := make(map[string]interface{})
    data["Title"] = "Irrigation Queue Management"
    page_template.ExecuteTemplate(w,"bootstrap", data)
    
}

func generate_eto_manage()string{
    return generate_irrigation_diag_html()+
           generate_irrigation_diag_js()
}

func generate_irrigation_diag_html()string{
  html := `
<div class="container">
 <h3 id=edit_panel_header>Irrigation Queue Management</h3>
 
	        
</div>
`
return html 
}

func  generate_irrigation_diag_js()string{
   js_data := `
        <script type="text/javascript"> 
   </script>`   
        
    
    return js_data
    
    
}
 
