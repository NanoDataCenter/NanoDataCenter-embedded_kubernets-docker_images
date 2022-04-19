package irrigation_operations

import (
   // "fmt"
    //"strings"
   // "net/http"
  //  "html/template"
    //"encoding/json"
     _ "embed"
   // "lacima.com/go_application_containers/irrigation/irrigation_web_utilities"
     "lacima.com/Patterns/web_server_support/jquery_react_support"
   // "lacima.com/redis_support/generate_handlers"
	//"lacima.com/redis_support/graph_query"
  //  "lacima.com/redis_support/redis_handlers"
  
  //  "lacima.com/Patterns/web_server_support/jquery_react_support"
    //"lacima.com/Patterns/msgpack_2"
    //"github.com/vmihailenco/msgpack/v5"
     
   // "lacima.com/go_application_containers/irrigation/irrigation_libraries/postgres_access/schedule_access"
)

//go:embed js/irrigation_schedule.js
var  irrigation_schedule_select_js string

func generate_schedule_component()web_support.Sub_component_type{

    return_value := web_support.Construct_subsystem("queue_schedule") 
   return_value.Append_line(web_support.Generate_title("Queue Irrigation Job"))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Back","schedule_cancel_id"))
    return_value.Append_line("</div>")
    return_value.Append_line(web_support.Generate_space("25"))
     null_list := make([]string,0)
    return_value.Append_line(web_support.Generate_select("Select Action","schedule_action_select",null_list,null_list))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_table("List of Schedule Step","schedule_table"))
    return_value.Append_line("</div>")

    return_value.Append_line(js_generate_schedule_queue())
    
    
    return return_value

}

func js_generate_schedule_queue()string{

  return_value := `<script type="text/javascript"> 
  ` +  irrigation_schedule_select_js +` 
  </script>`
  
  return return_value
    
    
}

