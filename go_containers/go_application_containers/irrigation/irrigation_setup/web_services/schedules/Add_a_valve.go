package construct_schedule




import(
    //"fmt"
   "lacima.com/Patterns/web_server_support/jquery_react_support"
)



func generate_edit_a_valve()web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("edit_a_valve")


    return_value.Append_line(web_support.Generate_title("Add Valves"))
    return_value.Append_line(web_support.Generate_space("25"))
    
    return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Add Valve","add_valve_id"))
    return_value.Append_line(web_support.Generate_button("Back","quit_valve_id"))
    return_value.Append_line("</div>")
    null_list := make([]string,0)
    return_value.Append_line(web_support.Generate_select("Select Station","station_node",null_list,null_list))
    return_value.Append_line(web_support.Generate_space("25"))
    
    return_value.Append_line(web_support.Generate_select("Select Valve","station_valve_list",null_list,null_list))
    return_value.Append_line("</div>")


    return_value.Append_line(js_add_a_valve())
    
    return return_value

}  

func js_add_a_valve()string{
  return_value := 
    `<script type="text/javascript"> 
    var local_valve_list
    
    function edit_a_valve_start(){
      
       hide_all_sections()
       $("#edit_a_valve").show()
       load_station_select()
       load_valve_select()
       
    }
   
    function edit_a_valve_init(){
      attach_button_handler("#add_valve_id",add_valve_click)
      attach_button_handler("#quit_valve_id",cancel_valve_click)
      jquery_initalize_select("#station_node",change_station)
      
     

    } 
   
    
   
    function add_valve_click(){
    
      temp = {}
      station = $("#station_node").val()
      valve   = $("#station_valve_list").val()
      
       if ( ( station in edit_step_working["station"]) == false){
           edit_step_working["station"][station] = {}
       }
       
       edit_step_working["station"][station][valve] = parseFloat(valve)
       
       
      edit_a_step_start()
    }
    
    function cancel_valve_click(){
        edit_a_step_start()
    }
    
    
    
    
    function change_station(event,ui){
      
       load_valve_select()
              
   }  
   
   function load_station_select(){
     master_controller  = $('#master_server').val()
     sub_controller     = $("#sub_server").val()
 
     local_valve_list   = valve_list[master_controller][sub_controller]
     station_list       =  Object.keys(local_valve_list)
     jquery_populate_select("#station_node",station_list,station_list,null)
      
   
   }
   function load_valve_select(){
       var station
       var valve_data
       var valve_array
       
       valve_array = []
       
       station = $("#station_node").val()
       valve_data = local_valve_list[station]
      
       for( let i = 1; i<= valve_data.length;i++){
          valve_array.push(i)
          
       }
       jquery_populate_select("#station_valve_list",valve_array,valve_array,null)
   }
    
   
    
    
    </script>`
    
  return return_value
}
/*
{"main_server":{"sub_server_1":{"station_1":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44],"station_2":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22],"station_3":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22],"station_4":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20]}}}
valve_list_json
*/
