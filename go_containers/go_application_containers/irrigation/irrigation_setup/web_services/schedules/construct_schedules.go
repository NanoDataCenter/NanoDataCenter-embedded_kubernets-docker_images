package construct_schedule


import(
    //"fmt"
    
    "strings"
    "net/http"
    "html/template"
    
    "lacima.com/go_application_containers/irrigation/irrigation_libraries/postgres_access/schedule_access"
)

var base_templates                    *template.Template



var sched_access irr_sched_access.Irr_sched_access_type

 
 

func Page_init(input *template.Template){
    
    
    base_templates = input
    sched_access = irr_sched_access.Construct_irr_schedule_access()
    
    
}      


    



func Generate_page(w http.ResponseWriter, r *http.Request){
    
     
    page_template ,_ := base_templates.Clone()
    page_html := generate_html_js()
    template.Must(page_template.New("application").Parse(page_html))        
    data := make(map[string]interface{})
    data["Title"] = "Edit Irrigation Schedules"
    page_template.ExecuteTemplate(w,"bootstrap", data)
    
}


func generate_html_js()string{
    
    
    return generate_html()+generate_js()
}


func generate_html()string{
   html_array := make([]string,5)
   html_array[0] = generate_top_html()
   html_array[1] = generate_create_schedule_name_html()
   html_array[2] = generate_edit_table_html()
   html_array[3] = generate_valve_table_html()
   html_array[4] = generate_copy_table_html()
   return strings.Join(html_array,"\n")
    
    
}

func generate_js()string{
   js_array := make([]string,7)
   js_array[0] = js_generate_global_js()
   js_array[1] = js_generate_top_js()
   js_array[2] = js_generate_populate_table()
   js_array[3] = js_generate_create_schedule_name()
   js_array[4] = js_generate_edit_table()
   js_array[5] = js_generate_valve_table()
   js_array[6] = js_generate_copy_table()
   return strings.Join(js_array,"\n")    
}
   
func js_generate_global_js()string{

 
  return_value := 
    ` <script type="text/javascript"> 
    master_sub_server_json ='`+ sched_access.Master_table_list_json+`'
    master_sub_server = JSON.parse(master_sub_server_json)

    valve_list_json ='`+sched_access.Valve_list_json+`'
    valve_list = JSON.parse(valve_list_json)
    ajax_add_schedule    = "ajax/irrigation/irrigation_schedules/add_schedule" 
    ajax_delete_scheudle = "ajax/irrigation/irrigation_schedules/delete_schedule" 
    ajax_get_schedule    = "ajax/irrigation/irrigation_schedules/get_schedules" 


    
    
    
    function hide_all_sections(){
    
        $("#main_section").hide()
        $("#table_construction").hide()
        $("#valve_section").hide()
        $("#table_name_section").hide()
        $("#copy_section").hide()
    
    }
    
    $(document).ready(
    function()
    {  
       hide_all_sections()
       initialize_main_panel()
       initialize_schedule_construction_panel()
       initialize_valve_construction_panel()
       initialize_table_name_panel()
       initialize_copy_panel()
       setup_table()
       start_main_panel()
    
    })
    
    
    </script>`
    
  return return_value
    
    
}
  
    


