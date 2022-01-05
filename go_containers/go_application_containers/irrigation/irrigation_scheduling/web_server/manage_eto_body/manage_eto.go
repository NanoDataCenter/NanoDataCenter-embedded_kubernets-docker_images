package manage_eto


import (
  "fmt"
  //"os"
  //"strings"  
  //"html/template"
  //"net"
  "lacima.com/go_application_containers/irrigation/eto/eto_calc/eto_support"
  "net/http"
  "lacima.com/redis_support/redis_handlers" 
  "html/template"
)


var  base_templates     *template.Template
var  eto_accumulation   redis_handlers.Redis_Hash_Struct

func Init(templates *template.Template, accumulation redis_handlers.Redis_Hash_Struct ){
    base_templates    = templates
    eto_accumulation  = accumulation
    
}


func Generate_page(w http.ResponseWriter, r *http.Request){
    values := eto_support.GetAll_Accumulation_Tables(eto_accumulation)
    fmt.Println("values",values)
    
    
    
}
