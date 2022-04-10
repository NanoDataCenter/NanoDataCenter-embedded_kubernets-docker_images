package irrigation_web_support

/*

func generate_controller_select(title string )web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("main_form")

    null_list := make([]string,0)
   
    return_value.Append_line(web_support.Generate_space("25"))
     return_value.Append_line(web_support.Generate_sub_title("master_state","Master Server State"))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_check_box("Select For Master  Controller", "master_controller_select"))
     return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_select("Select Master Server","master_server",null_list,null_list))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_div_start_plain("sub_controller_select"))
    return_value.Append_line(web_support.Generate_select("Select Sub Server","sub_server",null_list,null_list))
    return_value.Append_line(web_support.Generate_div_end())
   return_value.Append_line(web_support.Generate_space("25"))
    
   
    
    return return_value

}

func generate_controller_select(
` <script type="text/javascript"> 
 
    
    
   
  
    function  controller_init()
      
      master_key = Object.keys(master_sub_server)
      master_key.sort()
      jquery_populate_select('#master_server',master_key,master_key,master_server_change)
      let sub_key  = master_key[0]
      let sub_data = master_sub_server[sub_key]
      sub_data.sort()
      jquery_populate_select("#sub_server",sub_data,sub_data,sub_server_change)
      
    }
    
    
   function master_controller_select_function(){
     if(this.checked) {
       $("#sub_controller_select").hide()
      $("#master_state").html("Master Server State")
    } else {
        $("#sub_controller_select").show()
        $("#master_state").html("Sub Server State")
   }
 
    populate_table()
   }
   

   
   
function master_server_change(event,ui){
      let sub_key  = $("#master_server").val()
      let sub_data = master_sub_server[sub_key]
      sub_data.sort()
      jquery_populate_select("#sub_server",sub_data,sub_data,null)
     
      populate_table()   
   }
    
    
   function sub_server_change(event,ui){
    
     populate_table()
     }`
   
}
*/
