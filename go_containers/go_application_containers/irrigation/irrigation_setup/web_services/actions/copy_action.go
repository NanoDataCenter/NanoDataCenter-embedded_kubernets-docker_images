package construct_actions


import(
    //"fmt"
   "lacima.com/Patterns/web_server_support/jquery_react_support"
)


func generate_copy_action_name()web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("copy_action")

    
    
   
    return_value.Append_line(web_support.Generate_title("Enter New Actions Name and Description"))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Continue","copy_action_save_id"))
    return_value.Append_line(web_support.Generate_button("Back","copy_action_cancel_id"))
    return_value.Append_line("</div>")
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_input("Enter Name", "copy_action_input_id"))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_input("Enter Description","copy_action_description_id"))    
    return_value.Append_line("</div>")
    return_value.Append_line(js_generate_copy_action_name())
    
    return return_value

}



func js_generate_copy_action_name()string{
  return_value := 
    ` <script type="text/javascript"> 
    
    var ref_name
    
    function copy_action_start(name){
      ref_name = name
       hide_all_sections()
       show_section("copy_action")
       $("#copy_action_input_id").val("")
       $("#copy_action_description_id").val("")
    }
  
    function copy_action_init(main_controller,sub_controller,master_flag){
      
      attach_button_handler("#copy_action_save_id" ,copy_action_continue)
      attach_button_handler("#copy_action_cancel_id" ,copy_action_cancel)
      
    }
    function copy_action_continue(){
       let action_name = $("#copy_action_input_id").val()
       let description   = $("#copy_action_description_id").val()
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
          let main_controller  = $("#master_server").val()
           let sub_controller  = $("#sub_server").val()
           
           let master_flag = $("#master_controller_select").is(':checked')
           let new_action = action_data[ref_name]
           new_action["main_controller"]    = main_controller
            new_action["sub_controller"]       = sub_controller
           new_action["master_flag"]          = master_flag
           new_action["name"]                       = action_name
            new_action["description"]             = description  
            save_action_save(new_action)  
        
    }
    function copy_action_cancel(){
      start_section("main_form")
    }
    
    
    
   function save_action_save(working_action){
      // console.log("working action",working_action)
       ajax_post_get(ajax_add_action , working_action, copy_action_complete, "error action not saved") 
       
     }
     
    function copy_action_complete(){
       start_section("main_form")
    }
     
    
    
    </script>`
    
  return return_value
 
    
}
