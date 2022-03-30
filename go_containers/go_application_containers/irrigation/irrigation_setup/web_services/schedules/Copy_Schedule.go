package construct_schedule


import(
    //"fmt"
   "lacima.com/Patterns/web_server_support/jquery_react_support"
)


func generate_copy_schedule_name()web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("copy_schedule")

    
    
   
    return_value.Append_line(web_support.Generate_title("Enter New Schedule's Name and Description"))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Continue","copy_schedule_save_id"))
    return_value.Append_line(web_support.Generate_button("Back","copy_schedule_cancel_id"))
    return_value.Append_line("</div>")
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_input("Enter Name", "copy_schedule_input_id"))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_input("Enter Description","copy_schedule_description_id"))    
    return_value.Append_line("</div>")
    return_value.Append_line(js_generate_copy_schedule_name())
    
    return return_value

}



func js_generate_copy_schedule_name()string{
  return_value := 
    ` <script type="text/javascript"> 
    var  copy_select_index
    function copy_schedule_start(){
       hide_all_sections()
       show_section("copy_schedule")
       $("#copy_schedule_input_id").val("")
       $("#copy_schedule_description_id").val("")
    }
  
    function copy_schedule_init(){
      
      attach_button_handler("#copy_schedule_save_id" ,copy_schedule_save)
      attach_button_handler("#copy_schedule_cancel_id" ,add_schedule_cancel)
      
    }
    
    function copy_schedule_go(select_index){
        copy_select_index = select_index
        copy_schedule_start()
    
    }
    
    function copy_schedule_save(){
       let schedule_name = $("#copy_schedule_input_id").val()
       let description   = $("#copy_schedule_description_id").val()
       schedule_name.trim()
       description.trim()
       if (schedule_name.length == 0){
           alert("invalid schedule")
           return
       }
       
       if (schedule_name in schedule_data_map){
           alert("duplicate schedule")
           return
       }
       let ed_sch_working_schedule = deepcopy(schedule_data[copy_select_index])
       ed_sch_working_schedule["name"]        = schedule_name
       ed_sch_working_schedule["description"] = description
       ed_sch_working_schedule["server_key"] =  g_server_key
       ed_sch_working_schedule["json_steps"] = JSON.stringify([])
       ajax_post_get(ajax_add_schedule, ed_sch_working_schedule, copy_schedule_complete, "error copy schedule not saved") 
     }
     
    function copy_schedule_complete(){
       start_section("main_form")
    }
       

    function add_schedule_cancel(){
      start_section("main_form")
    }
    </script>`
    
  return return_value
 
    
}
    
