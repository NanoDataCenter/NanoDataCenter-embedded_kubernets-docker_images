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
    values := []string{"null","add_schedule","add_action","delete_entry","move_elements"}
    text   := []string{"Null Action","Add Schedule","Add Action","Delete Entry","Move Elements"}
    
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
    var action_table_list = [] 
    var action_table_rows = []
     function generate_action_steps_start(){
       hide_all_sections()
    
      show_section("generate_action_steps")
      $("#add_schedule_select_select").hide()
      $("#add_action_select_select").hide()
       if($("#master_controller_select").is(':checked')==false){
            populate_schedule_table()
             $('#action_step_select option[value="add_schedule"]' ).attr("disabled", false);
       }else{
             $('#action_step_select option[value="add_schedule"]' ).attr("disabled", true);
     }
      populate_action_table()
      
    }
  
  
  
   function modify_action_steps(select_index){
      action_table_list = [] 
      generate_action_steps_start()
      key = keys[select_index]
      $("#action_list_display_id").html("For Action   "+key)
       time_action_step_copy = deepCopyObject( action_data[key])
      load_initial_action_table(time_action_step_copy["steps"])
   }
   
   function load_initial_action_table(steps){
      
      action_table_rows = []
      let temp =[]
      for( i= 0;i<steps.length;i++){
          let temp = []
          temp.push( radio_button_element("Action_display_list_select"+i))
          temp.push(check_box_element("Action_display_list_checkbox"+i))
          let step_data = steps[i]
       
          type = step_data["type"]
          name = step_data["name"]
          description = step_data["description"]
          temp.push(type)
          temp.push(name)
          temp.push(description)
         action_table_rows.push(temp)
      }
      
      load_table("#action_step_list", action_table_rows)
   
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
       for( let i = 0; i< action_table_rows.length;i++){
           let temp = {}
           temp["type"]    = action_table_rows[i][2]
           temp["name"] = action_table_rows[i][3]
           temp["description"] = action_table_rows[i][4]
          time_action_step_copy["steps"].push(temp)
    }
    
    ajax_post_get(ajax_add_action ,time_action_step_copy, generate_action_steps_complete, "error action not saved") 
     
     }
     
    function generate_action_steps_complete(){
       start_section("main_form")
    }
    
     function action_steps_menu(event,ui){
       var index
       var choice
       choice = $("#action_step_select").val()
       $("#action_step_select")[0].selectedIndex = 0;
      
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
           delete_step_entry()
      }
    
     if(choice == "move_elements"){
            move()
      }
              
}      

function move(){
     let step_data    = deepcopy(time_action_step_copy["steps"])
     let select_index = find_select_index("Action_display_list_select",step_data.length)
     if( select_index == -1){
        alert("no move point")
        return
     }
     let move_array = find_check_box_elements("Action_display_list_checkbox",step_data.length)     
     if( move_array.length == 0){
         alert("no points to move ")
         return
     }
     let input = calculate_move(step_data.length,select_index,move_array)
    
     
     for( let i=0;i<input.length;i++){
         time_action_step_copy["steps"][i] = step_data[input[i]]
     }
     
     load_initial_action_table(time_action_step_copy["steps"])
   
   }




function delete_step_entry(){

     let select_index = find_select_index("Action_display_list_select",action_table_rows.length)
     if( select_index  == -1){
              alert("no action selected")
     }else{
          name = action_table_rows[select_index][3] 
          if( confirm("Delete Action "+name)== true){
             action_table_rows.splice(select_index,1)
             load_table("#action_step_list", action_table_rows)
         }
   }
}







  function create_action_step_list_table(){
   
      create_table( "#action_step_list",["Select","Move Select","Type","Name" ,"Description"])
   
   
   }
   
   function add_action_table(type,name,description){
      
      let temp =[]
      let index = action_table_rows.length
      temp.push( radio_button_element("Action_display_list_select"+index))
      temp.push(check_box_element("Action_display_list_checkbox"+index))

      temp.push(type)
      temp.push(name)
      temp.push(description)
  
     
      action_table_rows.push(temp)
      
      load_table("#action_step_list", action_table_rows)
   
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
       populate_schedule_table()
        
     }
     function save_added_schedule(){
        let keys = Object.keys(schedule_data_map)
        if( keys.length >0){
            let schedule =  $("#add_schedule_select").val()
            let description = schedule_data_map[schedule]
            add_action_table("schedule",schedule,description)
         }
         common_add_step_sub_window_return()
     
     }
     function common_add_step_sub_window_return(){
       $("#generate_action_steps").show()
        $("#add_schedule_select_select").hide()
         $("#add_action_select_select").hide()
     }
     function populate_schedule_table(){
        let keys = Object.keys(schedule_data_map)
        keys.sort()
        let display_list = []
        let value_list   = []
        for( let i= 0;i<keys.length;i++){
            let key  = keys[i]
            let description = schedule_data_map[key]
            let temp = key+":"+description
            display_list.push(temp)
            value_list.push(key)
      }
    jquery_populate_select($("#add_schedule_select"),value_list,display_list,null)
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

      populate_action_table()
       
     }
     function save_added_action(){
        let action =  $("#add_action_select").val()
            
        add_action_table("action",action,"")
         common_add_step_sub_window_return()
     
     }
     function populate_action_table(){
     
      jquery_populate_select($("#add_action_select"),action_data_list,action_data_list,null)
    }
     </script>
     `
     
     
    return return_value
}
