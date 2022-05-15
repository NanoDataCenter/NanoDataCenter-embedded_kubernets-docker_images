package irrigation_operations

import (
   //"fmt"
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
var control_block irr_sched_access.Irr_sched_access_type

var io_map                                  string
var valve_io                             string
var valve_group_names         string

//go:embed js/irrigation_operations.js
var irrigation_operations_js string

func Page_init(input *template.Template){
  
    base_templates = input

   
    control_block =  irr_sched_access.Construct_irr_schedule_access()
 

    temp               := control_block.Valve_group_data     
   
    io_map                                = temp["io_map"]
    valve_io                               = temp["valve_io"]
   valve_group_names          = temp["valve_group_names"]
  
   
   
  
}




func Generate_page_adjust(w http.ResponseWriter, r *http.Request){
    
     
    page_template ,_ := base_templates.Clone()
    page_html := generate_html_js()
    template.Must(page_template.New("application").Parse(page_html))        
    data := make(map[string]interface{})
    data["Title"] = "Irrigation Operations"
    page_template.ExecuteTemplate(w,"bootstrap", data)
    
    
}
 

func generate_html_js()string{
    web_variables := make(map[string]string)
    //web_variables["master_sub_server"]  = control_block.Master_table_list_json
    //web_variables["valve_list"]         = control_block.Valve_list_json
    
    ajax_variables := make(map[string]string)
    
    global_js                         :=  web_support.Load_jquery_ajax_components()
    global_js                         +=  web_support.Jquery_components_js()
    global_js                         +=  js_generate_global_js()
    global_js                         +=  irrigation_web_support.Queue_irrigation_jobs()
    
    top_list := web_support.Construct_web_components("main_section","main_form",web_variables,ajax_variables,global_js)    
    
    main_component := generate_main_component()
    top_list.Add_section(main_component)
    schedule_component :=  generate_schedule_component()
    top_list.Add_section(schedule_component)
    time_component :=  irrigation_web_support.Generate_step_time_change()
    top_list.Add_section(time_component)
    valve_group := generate_valve_group_component_component()
    top_list.Add_section(valve_group)
    station_channel := generate_station_channel()
    top_list.Add_section(station_channel)
    return irrigation_web_support.Attach_status_panel()+top_list.Generate_ending()  
    
}   





func js_generate_global_js()string{
    
 

  return_value := 
    ` valve_group_names_json =' `+valve_group_names +`'
     valve_group_names = JSON.parse(valve_group_names_json)
    //console.log(valve_group_names)
  
     valve_io_json = '`+ valve_io  +`'
    valve_io         = JSON.parse( valve_io_json )
    //console.log("valve io",valve_io)
    io_map_json =' `+ io_map     +`'
    io_map = JSON.parse(io_map_json)
    master_sub_server_json ='`+ control_block.Master_table_list_json+`'
    master_sub_server = JSON.parse(master_sub_server_json)
    ajax_get_actions                         =  "ajax/irrigation/irrigation_manage/get_actions" 
    ajax_get_schedule                       = "ajax/irrigation/irrigation_manage/get_schedules" 
   
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
     
    return_value := irrigation_web_support.Generate_controller_select("main_form","Manual Operations")
     return_value.Append_line(web_support.Generate_space("50"))
    
    return_value.Append_line(web_support.Generate_sub_title("sub_title","Select Appropriate Operations"))
    return_value.Append_line(web_support.Generate_space("10"))

    return_value.Append_line(web_support.Generate_space("10"))
    return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Manage Irrigation Queue","manage_select"))
    return_value.Append_line(web_support.Generate_button("Valve Io Management","manage_valve_group_io"))
    return_value.Append_line(web_support.Generate_button("Direct Io Management","manage_direct_io"))
    return_value.Append_line("</div>")
     return_value.Append_line(web_support.Generate_space("10"))
     null_list := make([]string,0)
    return_value.Append_line(web_support.Generate_select("Select Action","action_select",null_list,null_list))
     return_value.Append_line(web_support.Generate_space("10"))
     
    return_value.Append_line(web_support.Generate_select("Select Irrigation Schedule","irrigation_schedule_select",null_list,null_list))
     return_value.Append_line(web_support.Generate_space("10"))
    return_value.Append_line("</div>")
   
   
    return_value.Append_line(irrigation_web_support.Generate_controller_select_js())
    return_value.Append_line(js_generate_top_js())
    
    return return_value

}

func js_generate_top_js()string{

  return_value := `<script type="text/javascript"> 
  ` +  irrigation_operations_js +` 
  </script>`
  
  return return_value
    
    
}




