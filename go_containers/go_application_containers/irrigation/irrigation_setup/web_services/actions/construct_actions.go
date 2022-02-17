package construct_actions


import(
    //"fmt"
    
    //"strings"
    "net/http"
    "html/template"
    
    "lacima.com/go_application_containers/irrigation/irrigation_libraries/postgres_access/schedule_access"
    "lacima.com/Patterns/web_server_support/jquery_react_support"
)

var base_templates                    *template.Template

var control_block irr_sched_access.Irr_sched_access_type


 
 

func Page_init(input *template.Template){
    
    
    base_templates = input
    
    control_block = irr_sched_access.Construct_irr_schedule_access()
   
}      


    



func Generate_page(w http.ResponseWriter, r *http.Request){
    
     
    page_template ,_ := base_templates.Clone()
    page_html := generate_html_js()
    template.Must(page_template.New("application").Parse(page_html))        
    data := make(map[string]interface{})
    data["Title"] = "Edit Action Schedules"
    page_template.ExecuteTemplate(w,"bootstrap", data)
    
    
}

 
func generate_html_js()string{
    web_variables := make(map[string]string)
    web_variables["master_sub_server"]  = control_block.Master_table_list_json
    web_variables["valve_list"]         = control_block.Valve_list_json
    
    ajax_variables := make(map[string]string)
    ajax_variables["add_action"]    =   "ajax/irrigation/irrigation_schedules/add_action"
    ajax_variables["delete_action"] =   "ajax/irrigation/irrigation_schedules/delete_action" 
    ajax_variables["get_actions"]   =   "ajax/irrigation/irrigation_schedules/get_actions" 
    global_js                         :=  web_support.Load_jquery_ajax_components()
    global_js                         +=  web_support.Jquery_components_js()
    global_js                         +=  js_generate_global_js()
   
    
    top_list := web_support.Construct_web_components("main_section","main_form",web_variables,ajax_variables,global_js)    
    
    main_component := generate_main_component()
    top_list.Add_section(main_component)
    
    //get_schedule_name := generate_get_schedule_name()
    //top_list.Add_section(get_schedule_name)
    
    //edit_schedule_name := generate_edit_table_html()
    //top_list.Add_section(edit_schedule_name)
    
    
    //edit_step_name := generate_edit_a_step()
    //top_list.Add_section(edit_step_name)
    
    //edit_valve_name := generate_edit_a_valve()
    //top_list.Add_section(edit_valve_name)
    
    //step_time_change := generate_step_time_change()
    //top_list.Add_section(step_time_change)
    
    //copy_sched_name := generate_copy_schedule_name()
    //top_list.Add_section(copy_sched_name)
    
    return top_list.Generate_ending()
    
}   





func js_generate_global_js()string{

 
  return_value := 
    `
    // global data
    var schedule_data      =  []
    var schedule_data_map  =  {}
    
    master_sub_server_json ='`+ control_block.Master_table_list_json+`'
    master_sub_server = JSON.parse(master_sub_server_json)

    valve_list_json ='`+control_block.Valve_list_json+`'
    valve_list = JSON.parse(valve_list_json)
  
    ajax_add_schedule    = "ajax/irrigation/irrigation_schedules/add_schedule" 
    ajax_delete_schedule = "ajax/irrigation/irrigation_schedules/delete_schedule" 
    ajax_get_schedule    = "ajax/irrigation/irrigation_schedules/get_schedules" 



    
    
    
    
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
  
    


