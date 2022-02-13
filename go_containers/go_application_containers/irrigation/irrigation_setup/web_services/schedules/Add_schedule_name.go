
package construct_schedule


import(
    //"fmt"
   "lacima.com/Patterns/web_server_support/jquery_react_support"
)


func generate_get_schedule_name()web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("add_schedule")

    
    
   
    return_value.Append_line(web_support.Generate_title("Enter New Schedule's Name and Description"))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Continue","add_schedule_save_id"))
    return_value.Append_line(web_support.Generate_button("Back","add_schedule_cancel_id"))
    return_value.Append_line("</div>")
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_input("Enter Name", "add_schedule_input_id"))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_input("Enter Description","add_schedule_description_id"))    
    return_value.Append_line("</div>")
    return_value.Append_line(js_generate_create_schedule_name())
    
    return return_value

}



func js_generate_create_schedule_name()string{
  return_value := 
    ` <script type="text/javascript"> 
    function add_schedule_start(){
       hide_all_sections()
       show_section("add_schedule")
       $("#add_schedule_input_id").val("")
       $("#add_schedule_description_id").val("")
    }
  
    function add_schedule_init(){
      
      attach_button_handler("#add_schedule_save_id" ,add_schedule_continue)
      attach_button_handler("#add_schedule_cancel_id" ,add_schedule_cancel)
      
    }
    function add_schedule_continue(){
       let schedule_name = $("#add_schedule_input_id").val()
       let description   = $("#add_schedule_description_id").val()
       schedule_name.trim()
       description.trim()
       if (schedule_name.length == 0){
           alert("invalid schedule")
           return
       }
       console.log("schedule_name",schedule_name)
       console.log(schedule_data_map)
       if (schedule_name in schedule_data_map){
           alert("duplicate schedule")
           return
       }
       
       
         
           add_schedule(schedule_name,description)
           
     
    }
    function add_schedule_cancel(){
      start_section("main_form")
    }
    </script>`
    
  return return_value
 
    
}
/*
 var str = $("#myInput").val();
        alert(str);
*/
