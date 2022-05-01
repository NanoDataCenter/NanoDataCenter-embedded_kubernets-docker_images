 
 
 description_array = []
function valve_group_components_start(){
      initialize_direct_io_control()
       hide_all_sections()
       show_section("valve_group_components")     
     
    }
  
    function valve_group_components_init(){
      
        valve_group_description_map = []
        
        for( let i = 0; i< valve_group_names.length;i++){
            let temp = valve_io[ valve_group_names[i] ]
            valve_group_description_map.push(valve_group_names[i]+":"+temp["description"])
        }
   
        attach_button_handler("#valve_group_cancel_id" , station_channel_cancel_id)
     //console.log("valve_group_names",valve_group_names)
     // console.log("valve_group_description_map",valve_group_description_map)
      
      jquery_populate_select('#valve_group',valve_group_names,valve_group_description_map ,valve_group_change)
      let valve_group_id        = valve_group_names[0]

      valves_index           =   make_valves(  valve_io [valve_group_id])
      valve_choice = valve_group_names[0]
       description_array_a  = make_description_array( valve_choice)
    
      jquery_populate_select("#valve_id",  valves_index,  description_array_a, valve_change )
     Time_load_schedule_time("#valve_step_time_time_select",60)
     $("#valve_step_time_time_select").val('15').change()
    }
    function valve_group_cancel_id(){
      start_section("main_form")
    }
    function Time_load_schedule_time(id,number){
      load_times = []
      for( let i = 1; i<= number;i++){
          load_times.push(i)
          
      }
    
      jquery_populate_select(id,load_times,load_times,null)
   }
 
   function make_description_array( valve_choice ){
     return_value = []
     return_value.push("select valve")
     //console.log('valve_choice',valve_choice)
     //console.log("temp", valve_io[valve_choice]["valve_descriptions"])
     temp = valve_io[valve_choice]["valve_descriptions"]
     //console.log("temp",temp)
     for(i = 0; i < temp.length;i++){
         
         return_value.push((i+1)+ ":"+ temp[i])
        
      }
       
      return return_value
}
 
   function make_valves( input ){
      temp = input["valve_descriptions"]
      return_value = []
      return_value.push(0)
      for( i= 1; i < temp.length+1;i++){
        return_value.push(i)     
     }
    return return_value
}
   

function valve_group_change(event,ui){

      let valve_choice = $("#valve_group").val()
     
      let valves_index           =   make_valves(  valve_io [valve_choice])
   
       description_array_a  = make_description_array( valve_choice)
    
      jquery_populate_select("#valve_id",   valves_index,  description_array_a, valve_change )
    
 
}

    function valve_change(event,ui){
       var index
       var choice
       choice = $("#valve_id").val()
      $("#valve_id")[0].selectedIndex = 0;
       if( choice == 0 ){
           return
       }
     let master_controller_id          = $("#valve_group")[0].selectedIndex
     let master_controller_name   = valve_group_names[master_controller_id]
    
     let valve_group_data               = valve_io[master_controller_name]
   
     let stations                                = valve_group_data["stations"]
     let io                                           = valve_group_data["io"]
     let selected_station                = stations[choice]
     let selected_io                          = io[choice]
     let time                                     = parseInt($("#step_time_time_select").val())
     let message = "Queue Valve Group  "+master_controller_name +" Valve Id "+choice
    queue_irrigation_direct(selected_station ,selected_io,time,message)
}      
   
  
