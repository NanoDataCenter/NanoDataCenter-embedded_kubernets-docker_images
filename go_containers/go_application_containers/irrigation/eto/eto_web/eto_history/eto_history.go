package eto_history_page


import (
  //"fmt"
  //"os"
  //"strings"  
  //"html/template"
  //"net"
  "net/http"
  "lacima.com/server_libraries/postgres" 
  "html/template"
)

var base_templates  *template.Template
var  eto_history     pg_drv.Postgres_Stream_Driver

func Init(templates *template.Template, history pg_drv.Postgres_Stream_Driver){
    base_templates = templates
    eto_history    = history
    
}


func Generate_page(w http.ResponseWriter, r *http.Request){
    
    
    
    
}
