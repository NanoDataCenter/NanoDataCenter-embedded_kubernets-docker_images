package construct_schedule




import(
    //"fmt"
   "lacima.com/Patterns/web_server_support/jquery_react_support"
)


func generate_edit_table_html()web_support.Sub_component_type{
    return_value := web_support.Construct_subsystem("edit_schedule")

    
    
    
   
    return_value.Append_line(web_support.Generate_title("Edit Schedule"))
    return_value.Append_line(web_support.Generate_sub_title("edit_schedule_name","Schedule Name"))
    return_value.Append_line(web_support.Generate_sub_title("edit_schedule_description","Description"))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line("<div>")
    return_value.Append_line(web_support.Generate_button("Save","edit_schedule_save_id"))
    return_value.Append_line(web_support.Generate_button("Back","edit_schedule_cancel_id"))
    return_value.Append_line("</div>")
    return_value.Append_line(web_support.Generate_space("25"))
    
    
    values := []string{"null","create","pause","edit","copy","delete","time","move"}
    text   := []string{"Null Action","Create Step","Create Pause","Edit Step","Copy Step","Delete Step","Change Step Time","Move Steps"}
    
    return_value.Append_line(web_support.Generate_select("Select Action","edit_schedule_action",values,text))
    return_value.Append_line(web_support.Generate_space("25"))
    return_value.Append_line(web_support.Generate_table("List of Steps","edit_schedule_step_list"))
    return_value.Append_line("</div>")
    return_value.Append_line(js_generate_edit_schedule())
    
    return return_value

}   

func js_generate_edit_schedule()string{
  return_value := 
    ` <script type="text/javascript"> 
    var ed_sch_working_schedule
    var schedule_table_list
   
    
    function edit_schedule_start(){
       hide_all_sections()
       $("#edit_schedule").show()
    }
    
    function edit_schedule_init(){
      attach_button_handler("#edit_schedule_save_id",edit_schedule_save)
      attach_button_handler("#edit_schedule_cancel_id",edit_schedule_cancel)
      jquery_initalize_select("#edit_schedule_action",edit_schedule_menu)
      let columns = ["Select","Parm Select","Step Number", "Time","Valves"]
      create_table( "#edit_schedule_step_list",columns)
    }
   
    function add_schedule(name,description){
 
      ed_sch_working_schedule = {}
      ed_sch_working_schedule["master_server"] = $("#master_server").val()
      ed_sch_working_schedule["sub_server"]    = $("#sub_server").val()
      ed_sch_working_schedule["name"]        = name
      ed_sch_working_schedule["description"] = description
      ed_sch_working_schedule["steps"]       = []
      $("#edit_schedule_name").html("Schedule Name:  "+name)
      $("#edit_schedule_description").html("Description:  "+description)
      load_schedule_table()
      start_section("edit_schedule")
    }
    function edit_schedule( working_step){
       ed_sch_working_schedule = deepcopy(working_step)
       $("#edit_schedule_name").html("Schedule Name:  "+ed_sch_working_schedule["name"])
       $("#edit_schedule_description").html("Description:  "+ed_sch_working_schedule["description"])
       load_schedule_table()
       start_section("edit_schedule")
   }
    
    function edit_schedule_save(){
    
       ajax_post_get(ajax_add_schedule, ed_sch_working_schedule, edit_schedule_complete, "error message not saved") 
     }
     
    function edit_schedule_complete(){
       start_section("main_form")
    }
    
    function edit_schedule_cancel(){
    
      start_section("main_form")
    
    
    }
    
     function edit_schedule_menu(event,ui){
       var index
       var choice
       
       choice = $("#edit_schedule_action").val()
       if( choice == "create"){
           
           create_step_new()
           
       }
       if( choice == "pause"){
          step_time_activate_function(edit_schedule_start,add_pause_enity) 
        } 
       
       if( choice == "edit"){
          let step_data    = ed_sch_working_schedule["steps"]
           let select_index = find_select_index("edit_table_select_",step_data.length)
           if( select_index == -1){
               alert("no item is select")
               
           }else{
              edit_step_setup(select_index, step_data[select_index])

           }
           
       }
       if( choice == "copy"){
           let step_data    = ed_sch_working_schedule["steps"]
           let select_array = find_check_box_elements("edit_table_check_box_",step_data.length)
           if( select_array.length == 0){
               alert("no item is select")
               
           }else{
               
            for( i=0;i<select_array.length;i++){  
              let copy_data  = JSON.parse(JSON.stringify(step_data[select_array[i]]))
              
              ed_sch_working_schedule["steps"].push(copy_data)
              
              load_schedule_table()
              }
           }
       }
       if( choice == "delete"){
           
           let step_data    = ed_sch_working_schedule["steps"]
           let select_index = find_select_index("edit_table_select_",step_data.length)
           if( select_index == -1){
               alert("no item is select")
               
           }else{
              let temp = select_index +1
              if( confirm("Delete step "+temp+" ?") == true){
                  ed_sch_working_schedule["steps"]= deepslice(step_data,select_index,1) 
                  load_schedule_table()
                }
     
            }
           
       }
        if( choice == "time"){
           let step_data    = ed_sch_working_schedule["steps"]
           let select_array = find_check_box_elements("edit_table_check_box_",step_data.length)
           
           if( select_array.length == 0){
               alert("no item is select")
               
           }else{ step_time_activate_function(edit_schedule_start,step_time_change_bulk) }
           
       }
        if( choice == "move"){
           
           move()
           
       }
        
       $("#edit_schedule_action")[0].selectedIndex = 0;
              
   }      
   
   function load_schedule_table(){
      
     
      step_data = ed_sch_working_schedule["steps"]
      schedule_table_list = []
      schedule_table_count = 0
      
      let i = 0
      for( i= 0;i<step_data.length;i++){
          let valve_data = step_data[i]
          time = valve_data["time"]
          valve_json = JSON.stringify(valve_data["station"])
          select = radio_button_element("edit_table_select_"+i)
          check_box = check_box_element("edit_table_check_box_"+i)
          schedule_table_list.push([select,check_box,i+1,time,valve_json])


       }    
       load_table("#edit_schedule_step_list",  schedule_table_list)
          
  
   }
   
   function move(){
     let step_data    = deepcopy(ed_sch_working_schedule["steps"])
     let select_index = find_select_index("edit_table_select_",step_data.length)
     if( select_index == -1){
        alert("no move point")
        return
     }
     let move_array = find_check_box_elements("edit_table_check_box_",step_data.length)     
     if( move_array.length == 0){
         alert("no points to move")
         return
     }
     let input = calculate_move(step_data.length,select_index,move_array)
    
     
     for( let i=0;i<input.length;i++){
         ed_sch_working_schedule["steps"][i] = step_data[input[i]]
     }
     
     load_schedule_table()
   
   }
   
   
   function step_time_change_bulk(value){
      let select_array = find_check_box_elements("edit_table_check_box_",ed_sch_working_schedule["steps"].length)
      if( select_array.length == 0){
           alert("no item is select")
           return         
      }
      for( i = 0;i<select_array.length;i++){
         let index = select_array[i]
         ed_sch_working_schedule["steps"][index]["time"] = value
      }
      load_schedule_table()  
        
   }
   
   function add_pause_enity(value){
     
     var edit_step_working            = {}
     edit_step_working["time"]    = value
     edit_step_working["station"] = {} 
     edit_step_working["station"]["pause"] = {}
     edit_step_working["station"]["pause"]["1"] = 1
     ed_sch_working_schedule["steps"].push(edit_step_working)
     load_schedule_table() 
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
