package construct_actions


import(
    //"fmt"
   "lacima.com/Patterns/web_server_support/jquery_react_support"
)


func generate_time_setup()web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("time_action")

    return_value.Append_line(web_support.Generate_title("Manage Start Time of Action"))
    return_value.Append_line(web_support.Generate_sub_title("time_title_display_id","For Action "))
    return_value.Append_line(web_support.Generate_space("25"))
    
     return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Save","time_save_id"))
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
    
    var time_data_copy
    
     function time_action_start(){
       hide_all_sections()
    
      show_section("time_action")
      $("#hr_min_tag").show()
    }
  
    function time_action_init(main_controller,sub_controller,master_flag){
       attach_button_handler("#time_save_id" ,time_save_handler)
      attach_button_handler("#time_cancel_id" ,time_cancel_handler)
       attach_button_handler("#hour_select_id" ,hour_select_handler)
      attach_button_handler("#day_select_id" ,day_select_handler)
       hour_select_handler()
       initialize_hr_min_controls()
       time_doy_js_init()
    }
    
    function time_save_handler(){
    
     time_data_copy["start_time_hr"]    =   parseFloat($("#start_time_hr").val())
      time_data_copy["start_time_min"] = parseFloat($("#start_time_min").val())
      time_data_copy["end_time_hr"]      = parseFloat($("#end_time_hr").val())
      time_data_copy["end_time_min"]   = parseFloat($("#end_time_min").val())
      
     time_data_copy["day_mask"] = []
     time_data_copy["day_mask"].push($("#dow_0").is(':checked'))                                  
    time_data_copy["day_mask"].push($("#dow_1").is(':checked'))                                        
    time_data_copy["day_mask"].push($("#dow_2").is(':checked'))                          
    time_data_copy["day_mask"].push($("#dow_3").is(':checked'))                      
    time_data_copy["day_mask"].push( $("#dow_4").is(':checked'))                                
    time_data_copy["day_mask"].push( $("#dow_5").is(':checked'))                              
    time_data_copy["day_mask"].push($("#dow_6").is(':checked'))
   
     
     
     
     time_data_copy["dow_week_flag"] =   $("#time_type_select").is(':checked') 
     time_data_copy["doy_divisor"]        =   parseFloat(   $("#doy_divisor").val())
     time_data_copy["doy_modulus"]    =  parseFloat($("#doy_modulus").val())
    
    ajax_post_get(ajax_add_action , time_data_copy, time_data_action_complete, "error action not saved") 
     
     }
     
    function time_data_action_complete(){
       start_section("main_form")
    }
    
    
    function time_cancel_handler(){
       if(confirm("do you wish to leave") == true){
           start_section("main_form")
        }
   }
    function hour_select_handler(){
       $("#dow_tag").hide()
      $("#doy_tag").hide()
      $("#day_select_tag").hide()
      $("#hr_min_tag").show()
     
  }
    
  function  day_select_handler(){
      $("#day_select_tag").show()
      $("#hr_min_tag").hide()
       time_type_change_function()
      
  }
    
   
    function modify_start_time(select_index){
      time_action_start()
      
      key = keys[select_index]
     $("#time_title_display_id").html("For Action   "+key)
      time_data_copy = deepCopyObject( action_data[key])
   
      // set title
      // set time time hour min
      $("#start_time_hr").val(time_data_copy["start_time_hr"])
      $("#start_time_min").val(time_data_copy["start_time_min"] )
      $("#end_time_hr").val(time_data_copy["end_time_hr"])
     $("#end_time_min").val(time_data_copy["end_time_min"])
      
      // set time dow
     $("#dow_0").prop('checked', time_data_copy["day_mask"][0])
    $("#dow_1").prop('checked', time_data_copy["day_mask"][1])
    $("#dow_2").prop('checked', time_data_copy["day_mask"][2])
    $("#dow_3").prop('checked',  time_data_copy["day_mask"][3])
    $("#dow_4").prop('checked', time_data_copy["day_mask"][4])
   $("#dow_5").prop('checked',  time_data_copy["day_mask"][5])
   $("#dow_6").prop('checked',  time_data_copy["day_mask"][6])
   
    
      // set time doy
    $("#doy_divisor").val(time_data_copy["doy_divisor"])
    $("#doy_modulus").val(time_data_copy["doy_modulus"])
    // set time checkbox
    $("#time_type_select").prop('checked', time_data_copy["dow_week_flag"])
   
     
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
return_value.Append_line(`  <input type="checkbox" id="dow_4" name="dow_4" >`)
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
