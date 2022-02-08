package construct_schedule




import(
    //"fmt"
   "lacima.com/Patterns/web_server_support/jquery_react_support"
)


func generate_edit_a_step()web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("edit_a_step")

    
    
    
    return_value.Append_line(`<div id="edit_valve_top_section">`)
    return_value.Append_line(web_support.Generate_title("Edit Valves"))
    return_value.Append_line(web_support.Generate_sub_title("edit_step_number",""))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Save","edit_step_save_step_id"))
    return_value.Append_line(web_support.Generate_button("Back","edit_step_cancel_step_id"))
    return_value.Append_line("</div>")
    
    return_value.Append_line(web_support.Generate_space("25"))
    
    null_list := make([]string,0)
    
    return_value.Append_line(web_support.Generate_select("Select Irrigation Time","edit_step_time",null_list,null_list))
    
    return_value.Append_line(web_support.Generate_space("25"))
    values := []string{"null","add_valve","delete_valve"}
    text   := []string{"Null Action","Add Valve","Delete Valve"}
    
    return_value.Append_line(web_support.Generate_select("Select Action","edit_step_action",values,text))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line("</div>")
    
    
  
    
    return_value.Append_line(web_support.Generate_table("Current Selected Valves","edit_valve_list"))
    
    
 
    
    return_value.Append_line("</div>")
    return_value.Append_line(js_edit_a_step())
    
    return return_value

}  
//"key" in obj)
func js_edit_a_step()string{
  return_value := 
    `<script type="text/javascript"> 
    var edit_step_working
    var create_step_flag = false
    var ed_step_number  = 0
    
    
    var list_of_valves
    var table_count
    function edit_a_step_start(){
       hide_all_sections()
       $("#edit_a_step").show()
       $("#edit_step_time").val(edit_step_working["time"])
       load_valve_table("#edit_valve_list")
    }
   
    function edit_a_step_init(){
      attach_button_handler("#edit_step_save_step_id",edit_step_save)
      attach_button_handler("#edit_step_cancel_step_id",edit_step_cancel)
      attach_button_handler("#edit_step_time",edit_time_update)
      jquery_initalize_select("#edit_step_action",edit_step_menu)
      let columns = ["Select","Station","Valve"]
      create_table( "#edit_valve_list",columns)
      load_schedule_time("#edit_step_time",70)
      
      
    }
    
    function edit_step(number,step_data){
      //change step number label
      edit_working_step_data = step_data
      
    }
    
    function create_step_new(){
       create_step_flag = true
       edit_step_working            = {}
       edit_step_working["time"]    = 60
       edit_step_working["station"] = {} 
       list_of_valves               =[]    
       edit_a_step_start()
      
    }
   
     function edit_step_setup(input_step_number, input){
       create_step_flag = false
       ed_step_number  = input_step_number
       edit_step_working            = input
       edit_step_working["time"]    = input["time"]
       edit_step_working["station"] = input["station"]
       list_of_valves               =[]    
       edit_a_step_start()
     
     
     
      }
    
    
    function edit_step_save(){
  
       if( create_step_flag == true){
          ed_sch_working_schedule["steps"].push(edit_step_working)
       }else{
           ed_sch_working_schedule["steps"][ed_step_number] = edit_step_working
        }
       load_schedule_table()  
          
        edit_schedule_start()
     
    
    }
    
    function edit_step_cancel(){
      
        edit_schedule_start()
     
    
    }
    
    function edit_time_update(event,ui){
    
        edit_step_working["time"] = $("#edit_step_time").val()
     }
    
     function edit_step_menu(event,ui){
       var index
       var choice
       
       choice = $("#edit_step_action").val()
       
       if( choice == 'add_valve'){
          
          edit_a_valve_start()
        }
        
       if( choice == 'delete_valve'){
           delete_valve()
        }
                     
       
       $("#edit_step_action")[0].selectedIndex = 0;
              
   }
   function load_schedule_time(id,number){
      load_times = []
      for( let i = 0; i<= number;i++){
          load_times.push(i)
          
      }
      jquery_populate_select(id,load_times,load_times,null)
   }
   
   
   
   function load_valve_table(tag){
       
      list_of_valves  = []    
      table_count     = 0
      let valve_data = edit_step_working["station"]
      let stations = Object.keys(valve_data)
      stations.sort()
     
      for( let i = 0; i < stations.length;i++){
          station = stations[i]
          
          valve_data = edit_step_working["station"][station]
         
          load_valve_table_inner_part(station,valve_data)
      }
       
        load_table(tag,list_of_valves)
   
    }
   
    function load_valve_table_inner_part(station,valve_data){
        let keys = Object.keys(valve_data)
        keys.sort()
        
        for( let j = 0;j<keys.length;j++){
              let entry = []
              entry.push(load_select())
              entry.push(station)
              entry.push(keys[j])
              list_of_valves.push(entry)
        }
        
       }
   
      function load_select(){
        key = "valve_table"+table_count
        let label = " "
        select = radio_button_element(key)
        table_count +=1
        return select
       }
   
    function delete_valve(){
          let result = find_select_index("valve_table",table_count)
          if( result == -1 ){
              alert("no elements selected")
              return
          }
    
          let element = list_of_valves[result]
          let station = element[1]
          let valve   = element[2]
          delete  edit_step_working["station"][station][valve]
          let keys = Keys(edit_step_working["station"][station])
          if( keys.length == 0 ){
              delete  edit_step_working["station"][station]
          }
          load_valve_table("#edit_valve_list")
     }
   
    </script>`
    
  return return_value
 
    
}

