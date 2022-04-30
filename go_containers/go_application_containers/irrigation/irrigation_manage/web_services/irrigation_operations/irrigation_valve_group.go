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
     
      //  "lacima.com/go_application_containers/irrigation/irrigation_libraries/postgres_access/schedule_access"
)






//go:embed js/irrigation_valve_group.js
var js_valve_group string



func generate_valve_group_component_component()web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("valve_group_components")

    null_list := make([]string,0)
    return_value.Append_line(web_support.Generate_title("Valve Group Irrigation Diagnostic"))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Back","valve_group_cancel_id"))
    return_value.Append_line("</div>")
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_select("Select Valve Group","valve_group",null_list,null_list))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_select("Select Valve","valve_id",null_list,null_list))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_select("Select Irrigation Time","valve_step_time_time_select",null_list,null_list))
  
    return_value.Append_line("</div>")
    return_value.Append_line(js_generate_valve_group_js())
  
    return return_value

}

func js_generate_valve_group_js()string{

  return_value := `<script type="text/javascript">
  `+ js_valve_group+ `
  </script>`
  
  return return_value
    
    
}
