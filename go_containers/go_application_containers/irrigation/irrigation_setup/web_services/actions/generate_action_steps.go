package construct_actions


import(
    //"fmt"
   "lacima.com/Patterns/web_server_support/jquery_react_support"
)


func generate_action_steps_setup()web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("generate_action_steps")

    return_value.Append_line(web_support.Generate_title("Manage Action Steps"))
    return_value.Append_line(web_support.Generate_sub_title("action_list_display_id","For Action "))
    return_value.Append_line(web_support.Generate_space("25"))
    
     return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Save","action_step_save_id"))
    return_value.Append_line(web_support.Generate_button("Back","action_step_cancel_id"))
    return_value.Append_line("</div>")
    values := []string{"null","add_schedule","add_action","delete_entry"}
    text   := []string{"Null Action","Add Schedule","Add Action","Delete Entry"}
    
    return_value.Append_line(web_support.Generate_select("Select Function","action_step_select",values,text))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_table("List of Actions","action_step_list"))
    return_value.Append_line("</div>")
    
    return_value.Append_line(web_support.Generate_div_end())
 
    return_value = generate_schedule_select(return_value)
    
    return_value = generate_action_select(return_value)
     
     return_value.Append_line(web_support.Generate_div_end())
        
     return_value.Append_line(js_action_step_top_level())     
     
 
     
    return return_value

}



func  js_action_step_top_level()string{
  return_value := 
    ` <script type="text/javascript"> 
    
    var time_action_step_copy
    
     function generate_action_steps_start(){
       hide_all_sections()
    
      show_section("generate_action_steps")
      $("#add_schedule_select_select").hide()
      $("#add_action_select_select").hide()
      
    }
  
   function modify_action_steps(select_index){
      generate_action_steps_start()
      key = keys[select_index]
      $("#action_list_display_id").html("For Action   "+key)
       time_action_step_copy = deepCopyObject( action_data[key])
   
   }
  
    function generate_action_steps_init(main_controller,sub_controller,master_flag){
       attach_button_handler("#action_step_save_id" ,generate_action_steps_handler)
       attach_button_handler("#action_step_cancel_id" ,generate_action_steps_cancel_handler)
       jquery_initalize_select("#action_step_select",action_steps_menu)
       create_action_step_list_table()
       add_action_window_init()
       add_schedule_window_init()
        $("#add_schedule_select_select").hide()
      $("#add_action_select_select").hide()
    }
    function generate_action_steps_cancel_handler(){
         if(confirm("do you wish to leave") == true){
           start_section("main_form")
        }
   }
    
    
    function  generate_action_steps_handler(){
    
       time_action_step_copy["steps"]  = []
     
    
    ajax_post_get(ajax_add_action ,time_action_step_copy, generate_action_steps_complete, "error action not saved") 
     
     }
     
    function generate_action_steps_complete(){
       start_section("main_form")
    }
    
     function action_steps_menu(event,ui){
       var index
       var choice
       choice = $("#action_step_select").val()
       $("#action_map")[0].selectedIndex = 0;
       if( choice ==  "add_schedule"){
            $("#generate_action_steps").hide()
        $("#add_schedule_select_select").show()
         $("#add_action_select_select").hide()
           
       }
       
       
           
       if( choice == "add_action"){
          $("#generate_action_steps").hide()
        $("#add_schedule_select_select").hide()
         $("#add_action_select_select").show()
       }
      
     if(choice == "delete_entry"){
            alert("delete_entry")
      }
    
     
              
}      
  function create_action_step_list_table(){
   
      create_table( "#action_step_list",["Select","Type","Name" ,"Description"])
   
   
   }
   
    </script>
     `
     return return_value
}

 
 
func generate_schedule_select(return_value web_support.Sub_component_type)web_support.Sub_component_type{
   null_list := make([]string,0)
    return_value.Append_line(web_support.Generate_div_start("add_schedule_select_select"))
     return_value.Append_line(web_support.Generate_title("Add Schedule"))
     return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Add Schedule","action_add_schedule_save_id"))
    return_value.Append_line(web_support.Generate_button("Cancel","action_add_schedule_cancel_id"))
    return_value.Append_line("</div>")
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_select("Select Schedule","add_schedule_select",null_list,null_list))
    return_value.Append_line(generate_schedule_select_js( ))
    return_value.Append_line(web_support.Generate_div_end())
    return return_value
}
func generate_schedule_select_js( )string{
  return_value :=  ` <script type="text/javascript">  
      function add_schedule_window_init(){
      
       attach_button_handler("#action_add_schedule_save_id" ,save_added_schedule)
       attach_button_handler("#action_add_schedule_cancel_id" ,common_add_step_sub_window_return)
       
     }
     function save_added_schedule(){
         common_add_step_sub_window_return()
     
     }
     function common_add_step_sub_window_return(){
       $("#generate_action_steps").show()
        $("#add_schedule_select_select").hide()
         $("#add_action_select_select").hide()
     }
     </script>
     `
    return return_value
}

func generate_action_select(return_value web_support.Sub_component_type)web_support.Sub_component_type{
    null_list := make([]string,0)
    return_value.Append_line(web_support.Generate_div_start("add_action_select_select"))
     return_value.Append_line(web_support.Generate_title("Add Action"))
     return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Add Action","action_add_action_save_id"))
    return_value.Append_line(web_support.Generate_button("Cancel","action_add_action_cancel_id"))
    return_value.Append_line("</div>")
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_select("Select Action","add_action_select",null_list,null_list))
    return_value.Append_line(generate_action_select_js( ))
    return_value.Append_line(web_support.Generate_div_end())
    return return_value
}
func generate_action_select_js( )string{
  return_value :=  ` <script type="text/javascript">  
      function add_action_window_init(){
      
       attach_button_handler("#action_add_action_save_id" ,save_added_action)
       attach_button_handler("#action_add_action_cancel_id" ,common_add_step_sub_window_return)
       
     }
     function save_added_action(){
         common_add_step_sub_window_return()
     
     }
     </script>
     `
     
     
    return return_value
}
/*
 * 
 * 
 * 
 * 

func generate_time_hr_min(return_value web_support.Sub_component_type)web_support.Sub_component_type{
 null_list := make([]string,0)
 return_value.Append_line(web_support.Generate_div_start("hr_min_tag"))
 return_value.Append_line(web_support.Generate_sub_title("hr_min_display","Enter Earliest and Latest Start Time"))
return_value.Append_line(web_support.Generate_space("25"))
  return_value.Append_line(web_support.Generate_select("Select Earliest Start Time Hr","start_time_hr",null_list,null_list))
  return_value.Append_line(web_support.Generate_select("Select Earliest Start Time Min","start_time_min",null_list,null_list))  
  return_value.Append_line(web_support.Generate_select("Select Latest Start Time Hr","end_time_hr",null_list,null_list))
  return_value.Append_line(web_support.Generate_select("Select Latest Start Time Min","end_time_min",null_list,null_list))
   return_value.Append_line(web_support.Generate_div_end())
  return_value.Append_line(generate_time_js( ))
  return return_value
}

func generate_time_js( )string{
 return_value := 
    ` <script type="text/javascript">    
    
     function time_type_change_function(){
        let checked = $("#time_type_select").is(':checked')
       if( checked == true ){
           $("#dow_tag").show()
           $("#doy_tag").hide()
       }else{
             $("#dow_tag").hide()
           $("#doy_tag").show()
           
    }
    }
    
    function initialize_hr_min_controls(){
    
     $('#time_type_select').change(time_type_change_function)
       $("#start_time_hr").empty()
        for(let i=0; i<24; i++){
           $("#start_time_hr").append($('<option>').val(i).text(i));
        }
        $("#start_time_min").empty()
        for(let i=0; i<60; i++){
           $("#start_time_min").append($('<option>').val(i).text(i));
        }
       $("#end_time_hr").empty()
        for(let i=0; i<24; i++){
           $("#end_time_hr").append($('<option>').val(i).text(i));
        }

       $("#end_time_min").empty()
        for(let i=0; i<60; i++){
           $("#end_time_min").append($('<option>').val(i).text(i));
        }
     }
        </script>
     `
     return return_value
}

/
func generate_time_dow(return_value web_support.Sub_component_type)web_support.Sub_component_type{
  return_value.Append_line(web_support.Generate_div_start("dow_tag"))
  
return_value.Append_line(`  <input type="checkbox" id="dow_0" name="dow_0" >`)
return_value.Append_line(`   <label for="dow_0"> Sunday</label>`)
return_value.Append_line(`   <input type="checkbox" id="dow_1" name="dow_1" >`)
return_value.Append_line(`   <label for="dow_1"> Monday</label>`)
return_value.Append_line(`  <input type="checkbox" id="dow_2" name="dow_2" >`)
return_value.Append_line(`   <label for="dow_2">Tuesday</label>`)
return_value.Append_line(`  <input type="checkbox" id="dow_3" name="dow_3" >`)
return_value.Append_line(`   <label for="dow_3"> Wednesday</label>`)
return_value.Append_line(`  <input type="checkbox" id="dow_0" name="dow_4" >`)
return_value.Append_line(`   <label for="dow_4">Thursday</label>`)
return_value.Append_line(`  <input type="checkbox" id="dow_5" name="dow_5" >`)
return_value.Append_line(`   <label for="dow_5"> Friday</label>`)
return_value.Append_line(`  <input type="checkbox" id="dow_6" name="dow_6" >`)
return_value.Append_line(`   <label for="dow_6"> Saturday</label>`)

 
 
  return_value.Append_line(web_support.Generate_div_end())
  return return_value
}

func generate_time_doy(return_value web_support.Sub_component_type)web_support.Sub_component_type{
   null_list := make([]string,0)
  return_value.Append_line(web_support.Generate_div_start("doy_tag"))
  return_value.Append_line(web_support.Generate_select("Select DOY Divisor","doy_divisor",null_list,null_list))
  return_value.Append_line(web_support.Generate_select("Select DOY Modulus","doy_modulus",null_list,null_list))
  return_value.Append_line(generate_time_doy_js( ))
    return_value.Append_line(web_support.Generate_div_end())
  return return_value
}
   
func generate_time_doy_js( )string{
 return_value := 
    ` <script type="text/javascript">    
    function time_doy_js_init(){
        $("#doy_divisor").empty()
        for(let i=2; i<11; i++){
           $("#doy_divisor").append($('<option>').val(i).text(i));
        }
        $("#doy_modulus").empty()
        for(let i=0; i<10; i++){
           $("#doy_modulus").append($('<option>').val(i).text(i));
        }    
        jquery_initalize_select("#doy_divisor",doy_divisor_change)
      
   }
   
   function doy_divisor_change(){
     let index = $("#doy_divisor").val()
     $("#doy_modulus").empty()
        for(let i=0; i<index; i++){
           $("#doy_modulus").append($('<option>').val(i).text(i));
        }  
      $("#doy_moduls").val(0)
   }
    
    </script>`
    
    return return_value
}
*/
