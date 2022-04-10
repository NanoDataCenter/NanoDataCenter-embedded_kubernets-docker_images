package irrigation_station_channel


import (
   // "fmt"
    //"strings"
    "net/http"
    "html/template"
    //"encoding/json"
     _ "embed"
    "lacima.com/go_application_containers/irrigation/irrigation_web_utilities"
     "lacima.com/Patterns/web_server_support/jquery_react_support"
   // "lacima.com/redis_support/generate_handlers"
	//"lacima.com/redis_support/graph_query"
  //  "lacima.com/redis_support/redis_handlers"
  
  //  "lacima.com/Patterns/web_server_support/jquery_react_support"
    //"lacima.com/Patterns/msgpack_2"
    //"github.com/vmihailenco/msgpack/v5"
     
        "lacima.com/go_application_containers/irrigation/irrigation_libraries/postgres_access/schedule_access"
)

var base_templates                    *template.Template

var io_map                                  string
var valve_io                             string
var valve_group_names         string


//go:embed js/irrigation_station_channel.js
var js_station_channel string

func Page_init(input *template.Template){
    //initialize_eto_data_structures()
    base_templates = input

    control_block :=  irr_sched_access.Construct_irr_schedule_access()
    temp               := control_block.Valve_group_data     
   
    io_map                                = temp["io_map"]
    valve_io                         = temp["valve_io"]
   valve_group_names         = temp["valve_group_names"]
  
}




func Generate_page_adjust(w http.ResponseWriter, r *http.Request){
    
     
    page_template ,_ := base_templates.Clone()
    page_html := generate_html_js()
    template.Must(page_template.New("application").Parse(page_html))        
    data := make(map[string]interface{})
    data["Title"] = "Irrigation Diagnostics Station Channel"
    page_template.ExecuteTemplate(w,"bootstrap", data)
    
    
}


func generate_html_js()string{
    web_variables := make(map[string]string)
    //web_variables["master_sub_server"]  = control_block.Master_table_list_json
   // web_variables["valve_list"]         = control_block.Valve_list_json
    
    ajax_variables := make(map[string]string)
    ajax_variables["add_action"]    =   "ajax/irrigation/irrigation_schedules/add_action"
    ajax_variables["delete_action"] =   "ajax/irrigation/irrigation_schedules/delete_action" 
    ajax_variables["get_actions"]   =   "ajax/irrigation/irrigation_schedules/get_actions" 
    global_js                         :=  web_support.Load_jquery_ajax_components()
    global_js                         +=  web_support.Jquery_components_js()
    global_js                         +=  js_generate_global_js()
    global_js                         +=  irrigation_web_support.Queue_irrigation_jobs()
    
    top_list := web_support.Construct_web_components("main_section","main_form",web_variables,ajax_variables,global_js)    
    
    main_component := generate_main_component()
    top_list.Add_section(main_component)
   
    return irrigation_web_support.Attach_status_panel()+top_list.Generate_ending()  
    
}   





func js_generate_global_js()string{
  

  return_value := 
    `
 
    io_map_json =' `+ io_map     +`'
    io_map = JSON.parse(io_map_json)
    console.log(io_map_json)
  
   
  
    $(document).ready(
    function()
    {  
       init_sections()
       hide_all_sections()
       start_section("main_form")
    })

    
   `
    
  return return_value
    
  
}


func generate_main_component()web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("main_form")

    null_list := make([]string,0)
    return_value.Append_line(web_support.Generate_title("Station Channel Irrigation Diiagnostic"))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_select("Select Station","stations",null_list,null_list))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_select("Select Channel","channels",null_list,null_list))
    return_value.Append_line("</div>")
    return_value.Append_line(js_generate_top_js())
    
    return return_value

}

func js_generate_top_js()string{

  return_value := `<script type="text/javascript">
  `+ js_station_channel+ `
  </script>`
  
  return return_value
    
    
}




