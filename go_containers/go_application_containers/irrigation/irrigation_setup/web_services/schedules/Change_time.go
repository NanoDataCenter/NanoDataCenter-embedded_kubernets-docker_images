package construct_schedule




import(
    //"fmt"
   "lacima.com/Patterns/web_server_support/jquery_react_support"
)



func generate_step_time_change()web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("change_step_time")

    
    
    
    return_value.Append_line(web_support.Generate_title("Change Step Time"))
    return_value.Append_line(web_support.Generate_space("25"))
    
    return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Make Change","make_step_time_id"))
    return_value.Append_line(web_support.Generate_button("Back","quit_step_time_id"))
    return_value.Append_line("</div>")
    null_list := make([]string,0)
    
    
    return_value.Append_line(web_support.Generate_select("Select Irrigation Time","step_time_time_select",null_list,null_list))
    return_value.Append_line("</div>")

    return_value.Append_line(js_add_step_time_change())
    
    return return_value

}  

func js_add_step_time_change()string{
  return_value := 
    `<script type="text/javascript"> 
    var step_time_return_function
    var step_time_change_function
    
    function step_time_activate_function(return_function,change_function){
       step_time_return_function = return_function
       step_time_change_function = change_function
       change_step_time_start()
    }
    
    
    function change_step_time_start(){
       hide_all_sections()
       $("#change_step_time").show()
       
    }
   
    function change_step_time_init(){
      attach_button_handler("#make_step_time_id",make_step_time_id)
      attach_button_handler("#quit_step_time_id",quit_step_time_id)
      load_schedule_time("#step_time_time_select",70)
      
     

    } 
   
    function make_step_time_id(){
        new_value_string = $("#step_time_time_select").val()
        new_value        = parseFloat(new_value_string)
        step_time_change_function(new_value)
        step_time_return_function()
        
    
     }
   
   
     function quit_step_time_id(){
     
        step_time_return_function()
     
      }
    
    
    </script>`
    
  return return_value
}
