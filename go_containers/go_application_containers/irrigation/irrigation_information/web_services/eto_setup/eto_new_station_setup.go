package eto_setup

import (
   
    "strings"
)

func create_new_station_html(input string)string{
    html_array := make([]string,8)
    html_array[0] = create_new_station_main_html(input)
    html_array[1] = create_change_recharge_level_add("recharge_div_add")
    html_array[2] = create_crop_utilization_add("crop_utilization_div_add")
    html_array[3] = create_change_salt_flush_add("salt_flush_div_add")
    html_array[4] = create_change_sprayer_efficiency_add("sprayer_efficiency_div_add")
    html_array[5] = create_change_sprayer_rate_add("sprayer_rate_div_add")
    html_array[6] = create_change_tree_radius_add("tree_radius_div_add")
    html_array[7] = `</div>  `
    return strings.Join(html_array,"\n")
}

func create_new_station_main_html(top_div string)string{

html := `
<div id="`+top_div+`" class="container">
    <div id="station_save_div">
      <h3 id=edit_panel_header> Add New ETO Resource </h3>
       <div>
        <input type="button" id = "station_save" value="Save"  data-inline="true"  /> 
        <input type="button" id = "station_cancel" value="Cancel" data-inline="true"  /> 
        <input type="button" id = "station_reset" value="Reset" data-inline="true"  />
       
       </div>
      </div>

      <div id="station_add_controls">
            <h3>Select Controller</h3>
    	    <select  name="station_controller" id="station_controller" data-inline="true" >
            </select>
            <h3>Select Valve </h3>
    	    <select  name="station_valve" id="station_valve"  data-inline="true"  >
            </select>

      </div>

`
  return html + create_new_station_java_script()
}

func create_new_station_java_script()string{
    
 js := `<script type="text/javascript"> 
 function setup_station_controls(){
   
   
   $("#station_save").bind('click',station_save_control)
   $("#station_cancel").bind('click',station_cancel_control)
   $("#station_reset").bind('click',station_reset_control)
   $("#station_controller").bind('change',station_controller_change)
   recharge_div_add_load_controls()
   crop_utilization_div_add_load_controls()
   salt_flush_div_add_load_controls()
   sprayer_efficiency_div_add_load_controls()
   sprayer_rate_div_add_load_controls()
   tree_radius_div_add_load_controls()
   populate_controllers()
   populate_valves()
  
  }
  
  function station_controller_change()
   {
      
      populate_valves()
      
   }

   function station_reset_control()
   {
       var result = confirm("Do reset operation?");  
      
       if( result == true )
       {
          recharge_div_add_reset()
          crop_utilization_div_add_reset()
          salt_flush_div_add_reset()
          sprayer_efficiency_div_add_reset()
          sprayer_rate_div_add_reset()
          tree_radius_div_add_reset()

       }
    }
   
   function station_cancel_control()
   {
       var result = confirm("Do abandon operation?");  
      
       if( result == true )
       {
          setup_main_screen()

       }
    }

   function station_save_control()
   {
      
      
      
          let return_value = []
          let temp = []
          
          temp = recharge_div_add_get_value()
          return_value[temp[0]] = temp[1]
          
          temp = crop_utilization_div_add_get_value()
          return_value[temp[0]] = temp[1]
          
          temp = salt_flush_div_add_get_value()
          return_value[temp[0]] = temp[1]
          
          temp = sprayer_efficiency_div_add_get_value()
          return_value[temp[0]] = temp[1]
          
          temp = sprayer_rate_div_add_get_value()
          return_value[temp[0]] = temp[1]
          
          temp = tree_radius_div_add_get_value()
          return_value[temp[0]] = temp[1]
       
          return_value["station"] = $("#station_controller").val()
          return_value["valve"]   = parseFloat($("#station_valve").val())
          
          if (add_new_station(return_value) == true)  {
             setup_main_screen()
          }

       
    }
    
   function populate_controllers(  ){ 
       let key_list =  Object.keys(station_config_data) 
       key_list.sort()
       $("#station_controller").empty() 
       for(let  i= 0; i< key_list.length; i++)
       {
    
          $("#station_controller").append($("<option></option>").val(key_list[i]).html(key_list[i]));
       }
       $("#station_controller")[0].selectedIndex = 0;
       
   }
    
    function populate_valves() {
    
    
        var controller_index;

        let controller_value = $("#station_controller").val()
        
        let select_pins = station_config_data[controller_value]
    
        $("#station_valve").empty()
        for( let i = 0; i < select_pins.length; i++)
        {      
          $("#station_valve").append($("<option></option>").val(i+1).html("valve: "+(i+1)))
        } 
        $("#station_valve")[0].selectedIndex = 0; 
      
       
   }
   
  
  
  </script> `
    

return js
}

func  create_change_recharge_level_add(div_name string)string{
    size := 3
    return_array := make([]string,size)
    return_array[0] =  `<div id="`+div_name+`">`
    return_array[1] = create_recharge_html_add(div_name)
    return_array[size-1] = `</div>`    
    return strings.Join(return_array,"\n")
}

func create_recharge_html_add( div string )string{
    var config   ETO_Parameter_Setup
    
    config.Save_flag      = false
    config.Div_name       = div
    config.Title          = "Change Recharge Level"
    config.Field          = "recharge_level"
    config.Parm_name      = "recharge_level"
    
    config.Helper1        = "ETO Recharge Level:"
    config.Helper2        = "1"
    config.Helper3        = " inch"
    config.Max_Range      = "1"
    config.Step           = ".01"
    return_value := generate_parameter_setup(config)
    
    
    return return_value
}
    
    
    
    





func  create_crop_utilization_add(div_name string)string{
    size := 3
    return_array := make([]string,size)
    return_array[0] =  `<div id="`+div_name+`">`
    return_array[1] = create_crop_utilization_html_add(div_name)
    return_array[size-1] = `</div>`    
    return strings.Join(return_array,"\n")

}

func create_crop_utilization_html_add( div string )string{
    var config   ETO_Parameter_Setup
    
    config.Save_flag      = false
    config.Div_name       = div
    config.Title          = "Change Crop Utilization"
    config.Field          = "crop_utilization"
    config.Parm_name      = "crop_utilization"
    
    config.Helper1        = "Crop Utilization: "
    config.Helper2        = "100"
    config.Helper3        = " %"
    config.Max_Range      = "1"
    config.Step           = ".01"
    return_value := generate_parameter_setup(config)
   
    
    return return_value
}
  
func  create_change_salt_flush_add(div_name string)string{
     size := 3
    return_array := make([]string,size)
    return_array[0] =  `<div id="`+div_name+`">`
    return_array[1] = create_salt_flush_html_add(div_name)
    return_array[size-1] = `</div>`    
    return strings.Join(return_array,"\n")
}

func create_salt_flush_html_add( div string )string{
    var config   ETO_Parameter_Setup
   
    config.Save_flag      = false
    config.Div_name       = div
    config.Title          = "Change Salt Flush"
    config.Field          = "salt_flush"
    config.Parm_name      = "salt_flush"
    
    config.Helper1        = "Salt Flush:"
    config.Helper2        = "100"
    config.Helper3        = " %"
    config.Max_Range      = "1"
    config.Step           = ".01"
    return_value := generate_parameter_setup(config)
   
    
    return return_value
}
 


func  create_change_sprayer_efficiency_add(div_name string)string{
    size := 3
    return_array := make([]string,size)
    return_array[0] =  `<div id="`+div_name+`">`
    return_array[1] = create_sprayer_efficiency_html_add(div_name)
    return_array[size-1] = `</div>`    
    return strings.Join(return_array,"\n")

}

func create_sprayer_efficiency_html_add( div string )string{
    var config   ETO_Parameter_Setup
    
    config.Save_flag      = false
    config.Div_name       = div
    config.Title          = "Sprayer Efficiency"
    config.Field          = "sprayer_effiency"
    config.Parm_name      = "sprayer_effiency"
    
    config.Helper1        = "Sprayer Efficiency: "
    config.Helper2        = "100"
    config.Helper3        = " %"
    config.Max_Range      = "1"
    config.Step           = ".01"
    return_value := generate_parameter_setup(config)
    
    
    return return_value
}
 


func  create_change_sprayer_rate_add(div_name string)string{
    size := 3
    return_array := make([]string,size)
    return_array[0] =  `<div id="`+div_name+`">`
    return_array[1] = create_sprayer_rate_html_add(div_name)
    return_array[size-1] = `</div>`    
    return strings.Join(return_array,"\n")

}

func create_sprayer_rate_html_add( div string )string{
    var config   ETO_Parameter_Setup
    
    config.Save_flag      = false
    config.Div_name       = div
    config.Title          = "Sprayer Rate "
    config.Field          = "sprayer_rate"
    config.Parm_name      = "sprayer_rate"
    
    config.Helper1        = "Sprayer Rate: "
    config.Helper2        = "1"
    config.Helper3        = " GPM"
    config.Max_Range      = "50"
    config.Step           = ".1"
    return_value := generate_parameter_setup(config)
    
    
    return return_value
}
 
 
func create_change_tree_radius_add( div_name string )string{
   size := 3
    return_array := make([]string,size)
    return_array[0] =  `<div id="`+div_name+`">`
    return_array[1] = create_tree_radius_html_add(div_name)
    return_array[size-1] = `</div>`    
    return strings.Join(return_array,"\n")

}

    
func create_tree_radius_html_add( div  string )string{
    var config   ETO_Parameter_Setup
    
    config.Save_flag      = false
    config.Div_name       = div
    config.Title          = "Change Tree Radius"
    config.Field          = "tree_radius"
    config.Parm_name      = "tree_radius"
    
    config.Helper1        = "Tree Radius: "
    config.Helper2        = "1"
    config.Helper3        = " feet"
    config.Max_Range      = "15"
    config.Step           = ".1"
    return_value := generate_parameter_setup(config)
    
    
    return return_value
}    
    

    
    
    


