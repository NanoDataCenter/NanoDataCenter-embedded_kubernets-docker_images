package eto_daily_variation_page

import (
  //"fmt"
  //"os"
  //"strings"  
  //"html/template"
  //"net"
  "net/http"
  "lacima.com/redis_support/redis_handlers" 
  "html/template"
)

var base_templates       *template.Template
var  eto_stream_data     redis_handlers.Redis_Hash_Struct

func Init(templates *template.Template, stream_data redis_handlers.Redis_Hash_Struct){
    base_templates = templates
    eto_stream_data    = stream_data
    
}


func Generate_page(w http.ResponseWriter, r *http.Request){
    
    
    
    
}

