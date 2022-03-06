package construct_actions


import(
    //"fmt"
    "lacima.com/Patterns/web_server_support/jquery_react_support"
)



func generate_main_component()web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("main_form")

    null_list := make([]string,0)
    return_value.Append_line(web_support.Generate_title("Manage Irrigation Actions"))
    return_value.Append_line(web_support.Generate_space("25"))
     return_value.Append_line(web_support.Generate_sub_title("master_state","Master Server State"))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_check_box("Select For Master  Controller", "master_controller_select"))
     return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_select("Select Master Server","master_server",null_list,null_list))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_div_start("sub_controller_select"))
    return_value.Append_line(web_support.Generate_select("Select Sub Server","sub_server",null_list,null_list))
    return_value.Append_line(web_support.Generate_div_end())
   return_value.Append_line(web_support.Generate_space("25"))
    
    values := []string{"null","create","copy","delete","edit_actions","edit_start_time"}
    text   := []string{"Null Action","Create Action","Copy Action","Delete Action","Edit Action Steps","Edit Start Time"}
    
    return_value.Append_line(web_support.Generate_select("Select Action","action_map",values,text))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_table("List of Actions","action_list"))
    return_value.Append_line("</div>")
    return_value.Append_line(js_generate_top_js())
    
    return return_value

}








func js_generate_top_js()string{

  return_value := 
  ` <script type="text/javascript"> 
 
    /***************************** Main Setup Screen ************************************************/
    
    function main_form_start(){
       hide_all_sections()
       show_section("main_form")
     
        populate_action_list()
    }
  
    function main_form_init(){
      
      master_key = Object.keys(master_sub_server)
      master_key.sort()
      jquery_populate_select('#master_server',master_key,master_key,master_server_change)
      let sub_key  = master_key[0]
      let sub_data = master_sub_server[sub_key]
      sub_data.sort()
      jquery_populate_select("#sub_server",sub_data,sub_data,sub_server_change)
      jquery_initalize_select("#action_map",main_menu)
      create_action_list_table()
      $('#master_controller_select').change(master_controller_select_function)
      $("#master_state").html("Sub Server State")
       $("#sub_controller_select").show()
      
    }
    
    /**************************** Control Javascript Function **********************/
    
   function master_controller_select_function(){
     if(this.checked) {
       $("#sub_controller_select").hide()
      $("#master_state").html("Master Server State")
    } else {
        $("#sub_controller_select").show()
        $("#master_state").html("Sub Server State")
   }
 
    populate_action_list()
   }
   

   
   
function master_server_change(event,ui){
      let sub_key  = $("#master_server").val()
      let sub_data = master_sub_server[sub_key]
      sub_data.sort()
      jquery_populate_select("#sub_server",sub_data,sub_data,null)
     
      populate_action_list()   
   }
    
    
   function sub_server_change(event,ui){
    
     populate_action_list()
   }

   /********************************** Main Action Dispacther ************************************/
    function main_menu(event,ui){
       var index
       var choice
       choice = $("#action_map").val()
       
       if( choice == "create"){
           
           
           add_action_start()
       }
       
       
           
       if( choice == "copy"){
         copy_handlers()
       }
       if( choice == "delete"){
           
          delete_handler()
       }
     if(choice == "edit_start_time"){
              edit_start_time_handler()
      }
       if(choice == "edit_actions"){
              edit_start_step_handler()
      }
     $("#action_map")[0].selectedIndex = 0;
              
}      
   
   
  /*******************************************************  action handlers ***********************/

function edit_start_step_handler(){
     let select_index = find_select_index("Action_display_",key_list.length)
     if( select_index  == -1){
           alert("no action selected")
    }else{
         if ($("#master_controller_select").is(':checked') == true) {
           schedule_map = {}
           alert("no schedules for master controller")
         
       }else{
       modify_action_steps(select_index)
       }
    }
}

  
  
  
function copy_handlers(){
      let select_index = find_select_index("Action_display_",key_list.length)
      if( select_index  == -1){
           alert("no action selected")
     }else{
        key = key_list[select_index]
       copy_action_start(key)
    }
}         

function delete_handler(){
     let select_index = find_select_index("Action_display_",key_list.length)
     if( select_index  == -1){
              alert("no action selected")
     }else{
         let key = key_list[select_index]
          let item = action_data[key]
          let name = item["name"]
          
          if( confirm("Delete Action "+name)== true){
              
                   let data = {}
                   let master_flag = $("#master_controller_select").is(':checked')
                   data["master_flag"]   = master_flag
                   data["main_controller"] = $("#master_server").val()
                   data["sub_controller"]    = $("#sub_server").val()
                   data["name"]     =  name
                  
                   ajax_post_get( ajax_delete_action,data, populate_action_list, "action not deleted")
        
         }
   }
}
     
 function edit_start_time_handler(){
     let select_index = find_select_index("Action_display_",key_list.length)
     if( select_index  == -1){
              alert("no action selected")
     }else{
         modify_start_time(select_index)
    }
    
}
   
   
/********************************************** table handlers *******************************************************/   
   
   
   function create_action_list_table(){
   
      create_table( "#action_list",["Select","Name","Description" ,"Start Time","End Time","# of Steps"])
   
   
   }
   
  
   
   function populate_action_list(){
       let data = {}
       data["master_flag"]           = $("#master_controller_select").is(':checked')
       data["master_controller"] = $("#master_server").val()
       data["sub_controller"]    = $("#sub_server").val()
      
       ajax_post_get(ajax_get_actions  , data, ajax_get_function,  "Action Data Not Loaded")
       
    }
   function ajax_get_function(data){
      action_data  = {}
      
      //console.log(data)
      
      action_data_map = {}
      set_status_bar("Action Data Downloaded")
      let row_data = []
      let i = 0
      for (i = 0;i<data.length;i++){
        let temp = JSON.parse(data[i])
        let key = temp["name"]
        action_data[key]  = temp
     }
     //console.log(action_data)
     keys = Object.keys(action_data)
     keys.sort()
    // console.log("keys",keys)
     key_list = keys
     for( let i= 0;i<keys.length;i++){
         key = keys[i]
         let temp = action_data[key]
        // console.log(key,temp)
         let entry               =   []
         let name            =   temp["name"]
         let description  = temp["description"]
         let start_time        = temp["start_time_hr"]+":"+temp["start_time_min"]
         let end_time          = temp["end_time_hr"]+":"+temp["end_time_min"]
         let number_of_steps   = temp["steps"].length
         action_data_map[name] = true 
         entry.push(radio_button_element("Action_display_"+i))
         
         entry.push(name)
         entry.push(description)
         entry.push(start_time)
         entry.push(end_time)
         entry.push(number_of_steps)
         row_data.push(entry)
        
      }
     //console.log(row_data)
     load_table('#action_list', row_data)
      get_schedules()
   }
   
  
function get_schedules(){
       if ($("#master_controller_select").is(':checked') == true) {
           schedule_map = {}
           return
       }
        let data = {}
       data["master_server"] = $("#master_server").val()
       data["sub_server"]    = $("#sub_server").val()
      
       ajax_post_get(ajax_get_schedule , data, ajax_get_schedule_function,  "Schedule Data Not Loaded")
       
    }





 
   function ajax_get_schedule_function(data){
   
     
      schedule_data  = data
      
      console.log(schedule_data)
      
      schedule_data_map = {}
      set_status_bar("Schedule Data Downloaded")
      let row_data = []
      let i = 0
      for (i = 0;i< schedule_data.length;i++){
         let entry =[]
         let name = schedule_data[i]["name"]
         schedule_data_map[name] = schedule_data[i]["description"]
        
   
      }
     console.log("schedule_data_map",schedule_data_map)
     
      
   }
   
   
    </script>`
    
  return return_value
    
    
}
