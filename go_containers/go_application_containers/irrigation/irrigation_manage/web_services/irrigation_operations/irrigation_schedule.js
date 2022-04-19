


    var working_scheduling
    
    function start_schedule_select(choice){
         working_scheduling = choice
         queue_schedule_start()
       
    }

    
    function queue_schedule_start(){
       hide_all_sections()
       fill_in_schedule_table()
       show_section("queue_schedule")
       
    }
  
    function queue_schedule_init(main_controller,sub_controller,master_flag){
     // attach select handler 
     create_table( "#schedule_table",["Ref Point","Select Point","Step","Time","Valves"  ])
      attach_button_handler("#schedule_cancel_id" ,schedule_cancel_idl)
      let select_array = ["select entry","Queue Entries","Move Entries"]
       jquery_populate_select("#schedule_action_select",select_array,select_array, schedule_select_handler)
    }
    
    function schedule_cancel_idl(){
      start_section("main_form")
    }
    
    function schedule_select_handler(){
       var index
       var choice
       index    = $("#schedule_action_select")[0].selectedIndex
       choice = $("#schedule_action_select").val()
       $("#schedule_action_select")[0].selectedIndex = 0;
       if (index == 0) {
           return
       }
     if(   $("#schedule_action_select")[0].selectedIndex == 2){
         move()
     }
     
    }
        
    
function fill_in_schedule_table(){
    
     let steps = schedule_step_map[working_scheduling] 

     let table_data = []
      for( i= 0;i<steps.length;i++){
          let temp = []
          temp.push( radio_button_element("Sched_display_list_select"+i))
          temp.push(check_box_element("Sched_display_list_checkbox"+i))   
         temp.push(steps[i]["step"])
          temp.push(steps[i]["time"])
          temp.push(steps[i]["steps"])
          table_data.push(temp)
      }
         load_table("#schedule_table", table_data)
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

    
