package construct_actions


import(
    //"fmt"
   "lacima.com/Patterns/web_server_support/jquery_react_support"
)


func generate_get_action_name()web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("add_action")

    
    
   
    return_value.Append_line(web_support.Generate_title("Enter New Action's Name and Description"))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Continue","add_action_save_id"))
    return_value.Append_line(web_support.Generate_button("Back","add_action_cancel_id"))
    return_value.Append_line("</div>")
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_input("Enter Name", "add_action_input_id"))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_input("Enter Description","add_action_description_id"))    
    return_value.Append_line("</div>")
    return_value.Append_line(js_generate_create_action_name())
    
    return return_value

}



func js_generate_create_action_name()string{
  return_value := 
    ` <script type="text/javascript"> 
    function add_action_start(){
       hide_all_sections()
       show_section("add_action")
       $("#add_action_input_id").val("")
       $("#add_action_description_id").val("")
    }
  
    function add_action_init(main_controller,sub_controller,master_flag){
      
      attach_button_handler("#add_action_save_id" ,add_action_continue)
      attach_button_handler("#add_action_cancel_id" ,add_action_cancel)
      
    }
    function add_action_continue(){
       let action_name = $("#add_action_input_id").val()
       let description   = $("#add_action_description_id").val()
       action_name.trim()
       description.trim()
       if (action_name.length == 0){
           alert("invalid action")
           return
       }
      
       if (action_name in action_data_map){
           alert("duplicate action")
           return
       }
          
           let new_action = blank_new_action()
           
           new_action["name"]                       = action_name
            new_action["description"]             = description  
            save_action_save(new_action)  
        
    }
    function add_action_cancel(){
      start_section("main_form")
    }
    
    
    function blank_new_action(){
      let return_value = {}
      return_value["server_key"]             =  g_server_key
      return_value["name"] = ""
      return_value["description"] = ""
      return_value["steps"] = []
      return_value["start_time_hr"] = 0
      return_value["start_time_min"] = 0
      return_value["end_time_hr"] = 0
      return_value["end_time_min"] = 0
      return_value["day_mask"] = [false,false,false,false,false,false,false]
      return_value["dow_week_flag"] = false
      return_value["doy_divisor"] = 2
      return_value["doy_modulus"] =0
      return return_value
   }
    
    
   function save_action_save(working_action){
      // console.log("working action",working_action)
       ajax_post_get(ajax_add_action , working_action, add_action_complete, "error action not saved") 
       
     }
     
    function add_action_complete(){
       start_section("main_form")
    }
     
    
    
    </script>`
    
  return return_value
 
    
}
