package construct_actions


import(
    //"fmt"
   "lacima.com/Patterns/web_server_support/jquery_react_support"
)


func generate_time_setup()web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("time_action")

    return_value.Append_line(web_support.Generate_title("Manage Time Scheduline"))
    return_value.Append_line(web_support.Generate_sub_title("title_display_id","For Action "))
    return_value.Append_line(web_support.Generate_space("25"))
    
     return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Continue","time_save_id"))
    return_value.Append_line(web_support.Generate_button("Back","time_cancel_id"))
    return_value.Append_line("</div>")
    
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Edit Hr/Min","hour_select_id"))
    return_value.Append_line(web_support.Generate_button("Edit Day","day_select_id"))
    return_value.Append_line("</div>")
    return_value.Append_line(web_support.Generate_space("25"))
    return_value =   generate_time_hr_min(return_value)
    
    
    return_value.Append_line(web_support.Generate_div_start("day_select_tag"))
    return_value.Append_line(web_support.Generate_sub_title("Time_state","Schedule By Day of The Year"))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_check_box("Select for Day of The Week Scheduling", "time_type_select"))    
    
    return_value = generate_time_doy(return_value)
    
    return_value = generate_time_dow(return_value)
    return_value.Append_line(web_support.Generate_div_end())
 
     
     return_value.Append_line(web_support.Generate_div_end())
        
     return_value.Append_line(js_time_top_level())     
     
 
     
    return return_value

}



func  js_time_top_level()string{
  return_value := 
    ` <script type="text/javascript"> 
     function time_action_start(){
       hide_all_sections()
      show_section("time_action")
    }
  
    function time_action_init(main_controller,sub_controller,master_flag){
       attach_button_handler("#time_save_id" ,time_save_handler)
      attach_button_handler("#time_cancel_id" ,time_cancel_handler)
       attach_button_handler("#hour_select_id" ,hour_select_handler)
      attach_button_handler("#day_select_id" ,day_select_handler)
       initialize_hr_min_controls()
    }
    
    function time_save_handler(){
       start_section("main_form")
    
    }
    
    function time_cancel_handler(){
       if(confirm("do you wish to leave") == true){
           start_section("main_form")
        }
   }
    function hour_select_handler(){
      $("#day_select_tag").hide()
      $("#hr_min_tag").show()
  }
    
  function  day_select_handler(){
      $("#day_select_tag").show()
      $("#hr_min_tag").hide()
  }
    
    
    function modify_start_time(select_index){
      time_action_start()
      hour_select_handler()
      
      //parseInt()
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
 */
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
    function initialize_hr_min_controls(){
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

/*
 *   dow
 */
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
/*
 *  doy
 * 
 */
func generate_time_doy(return_value web_support.Sub_component_type)web_support.Sub_component_type{
   null_list := make([]string,0)
  return_value.Append_line(web_support.Generate_div_start("doy_tag"))
  return_value.Append_line(web_support.Generate_select("Select DOY Divisor","sub_server",null_list,null_list))
  return_value.Append_line(web_support.Generate_select("Select DOY Modulus","sub_server",null_list,null_list))
    return_value.Append_line(web_support.Generate_div_end())
  return return_value
}
   
