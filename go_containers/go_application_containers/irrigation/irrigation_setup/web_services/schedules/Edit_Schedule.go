package construct_schedule


import(
    //"fmt"
   
)

import(
    //"fmt"
   "lacima.com/Patterns/web_server_support/jquery_react_support"
)


func generate_edit_table_html()web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("edit_schedule")

    
    
    
   
    return_value.Append_line(web_support.Generate_title("Edit Schedule"))
    return_value.Append_line(web_support.Generate_sub_title("edit_schedule_name","Schedule Name"))
    return_value.Append_line(web_support.Generate_sub_title("edit_schedule_description","Description Name"))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Save","edit_schedule_save_id"))
    return_value.Append_line(web_support.Generate_button("Back","edit_schedule_cancel_id"))
    return_value.Append_line("</div>")
    return_value.Append_line(web_support.Generate_space("25"))
    values := []string{"null","create","edit","copy","delete","time","move"}
    text   := []string{"Null Action","Create Step","Edit Step","Copy Step","Delete Step","Change Step Time","Move Steps"}
    
    return_value.Append_line(web_support.Generate_select("Select Action","schedule_action",values,text))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_table("List of Steps","edit_schedule_step_list"))
    return_value.Append_line("</div>")
    return_value.Append_line(js_generate_edit_schedule())
    
    return return_value

}   

func js_generate_edit_schedule()string{
  return_value := 
    ` <script type="text/javascript"> 
    function edit_schedule_start(){
       alert("edit start")
       hide_all_sections()
       //show_section("edit_schedule")
       $("#edit_schedule").show()
    }
  
    function edit_schedule_init(){
       alert("edit init")
    
    
    }
    
    function edit_schedule_cancel(){
    
      start_section("main_form")
    
    
    }
    </script>`
    
  return return_value
 
    
}









/*

func generate_edit_table_html()string{
    
  return_value :=
  `
  <div class="container" id="table_construction">
 
      <h3>Mange Schedules</h3>
      <h4 id= "table_construction_master">Master</h4>
      <h4 id= "table_construction_slave">Slave</h4>
      <h4>Schedule id= "table_construction_schedule</h4>
      <h3 > Make Schedule Adjustments </h3>
      <div>
           <input type="button" id = "schedule_save" value="Save"  data-inline="true"  /> 
           <input type="button" id = "schedule_abort" value="Abort" data-inline="true"  /> 
      </div>
      <select id="schedule_action">
      </select>
      <div class="container">
         <h5>List of Steps</h5>
         <div style="margin-top:10px"></div>
         <table id="step_list" class="display" width="100%"></table>
      </div>
  </div>

    
`
 return return_value
}

func js_generate_edit_table()string{

  return_value := 
    ` <script type="text/javascript"> 
     
       function initialize_schedule_construction_panel(){
          $("#schedule_save").bind('click',schedule_construction_save)
          $("#schedule_abort").bind('click',schedule_construct_abort)
          setup_single_schedule_table()
          
       }
       function setup_single_schedule_table(){
           let columns = [  { title:"Select" },{title:"Step"},{title:"TIME"},{title:"Change Valve"} ,{title:"Valve List"}   ]

         
           $('#schedule_list').DataTable( {
                   pageLength: 50,
                    columns: columns
           } );
       }
       
       function start_construct_table(){
         hide_all_sections()
         $("#main_section").show()
         //setup headers    .html("your new header");
         // setup schedule
       }
       function schedule_construction_save(){
         // make data base save of data save
         // if successful then
         hide_all_sections()
         $("#main_section").show()
       }
       function schedule_construct_abort(){
         hide_all_sections()
         $("#main_section").show()
       }
    </script>`
    
  return return_value
    
    
}   

    
*/
