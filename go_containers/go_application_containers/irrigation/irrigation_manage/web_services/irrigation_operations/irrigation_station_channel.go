package irrigation_operations


import (
   // "fmt"
    //"strings"
    //"net/http"
   // "html/template"
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


//go:embed js/irrigation_station_channel.js
var js_station_channel string



func generate_station_channel()web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("station_channel")

    null_list := make([]string,0)
    return_value.Append_line(web_support.Generate_title("Station Channel Irrigation Diagnostic"))

    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Back","station_channel_cancel_id"))
    return_value.Append_line("</div>")
    return_value.Append_line(web_support.Generate_space("25"))
  
    return_value.Append_line(web_support.Generate_select("Select Station","stations",null_list,null_list))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_select("Select Channel","channels",null_list,null_list))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_select("Select Irrigation Time","station_step_time_time_select",null_list,null_list))
    return_value.Append_line("</div>")
    return_value.Append_line(js_generate_station_channel_js())
    
    return return_value

}

func js_generate_station_channel_js()string{

  return_value := `<script type="text/javascript">
  `+ js_station_channel+ `
  </script>`
  
  return return_value
    
    
}




