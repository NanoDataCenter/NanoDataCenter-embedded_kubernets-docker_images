

package construct_schedule


import(
    //"fmt"
   
)




func generate_valve_table_html()string{
    
  return_value :=
  `
  <div class="container" id="valve_section">
 
     
    <h3>Mange Valves</h3>
    <h4>Master</h4>
    <h4>Slave</h4>
    <h4>Schedule</h4>
    <h4>Step</h4>
    
    <h3 > Make Valve Adjustments </h3>
       <div>
        <input type="button" id = "valve_save" value="Save"  data-inline="true"  /> 
        <input type="button" id = "valve_reset" value="Reset" data-inline="true"  /> 
       </div>
     
    <div class="container">
     
     <h5>List of Valves</h5>
     <div style="margin-top:10px"></div>
     <table id="valve_list" class="display" width="100%"></table>
    </div>
    
    </div>
    
    </div>
    
`
 return return_value
}

func js_generate_valve_table()string{

  return_value := 
    ` <script type="text/javascript"> 
      
       function initialize_valve_construction_panel(){
       
       
       }
       
       function start_valve_contruction_panel(){
       
       
       }
       
    </script>`
    
  return return_value
    
    
}
   
