package construct_schedule


import(
    //"fmt"
    
)


func generate_top_html()string{
    
  return_value :=
  `
  <div class="container" id="main_section">
 
     
    <h3>Mange Irrigation Schedules</h3>
    <h4>Select Master Server</h4> 
    <select id="master_server">
     </select>
    <h4>Select Sub Server</h4> 
    <select id="sub_server">
    </select>
    <div style="margin-top:20px"></div>
    <h4>Select Select Action</h4> 
    <select id="schedule_action">
    <option value="null">Null Action</option>
    <option value="create">Create Schedule</option>
    <option value="edit">Edit Schedule</option>
    <option value="copy">Copy Schedule</option>
    <option value="delete">Delete Schedule</option>
    </select>
    <div style="margin-top:20px"></div>
   
     <div style="margin-top:20px"></div>
     <h4>List of Schedules</h4>
     <div style="margin-top:20px"></div>
    
     <table id="schedule_list" class="display" width="100%"></table>
    
    
    </div>
    
    
`
 return return_value
}


func js_generate_top_js()string{

  return_value := 
    ` <script type="text/javascript"> 
    function initialize_main_panel(){

      populate_master_select()
      
      attach_action_handler()
      populate_table()

    }
       
    function start_main_panel(){
       hide_all_sections()
       $("#main_section").show()
    }
    
    // supporting function
    
    function populate_master_select(){
      
      master_key = Object.keys(master_sub_server)
      master_key.sort()
      
      for(let i=0; i<master_key.length; i++){
        $('#master_server').append($('<option>').val(master_key[i]).text(master_key[i]));
      }
      let sub_key  = master_key[0]
      let sub_data = master_sub_server[sub_key]
      populate_sub_server_select(sub_data)
      $("#master_server").bind('change', master_server_change)
      $("#sub_server")[0].selectedIndex = 0;
     
    }
    
    
    function populate_sub_server_select(sub_select_list){
    
        
        for(let i=0; i<sub_select_list.length; i++){
           $('#sub_server').append($('<option>').val(sub_select_list[i]).text(sub_select_list[i]));
        }
        $("#master_server").bind('change', sub_server_change)
        $("#sub_server")[0].selectedIndex = 0;
    }
    function attach_action_handler(){
    
      $("#schedule_action").bind('change',main_menu)
      $("#schedule_action")[0].selectedIndex = 0;
    }
    
   
    
    function main_menu(event,ui){
       var index
       var choice
       choice = $("#schedule_action").val()
       if( choice == "create"){
   
           start_table_entry_panel()
           
       }
       
       if( choice == "edit"){
           
           alert("edit")
           
       }
       if( choice == "copy"){
           
           // check if selected
           copy_table_generate_panel()
           
       }
       if( choice == "delete"){
           
           alert("delete")
           
       }
       $("#schedule_action")[0].selectedIndex = 0;
              
   }      
   
   function master_server_change(event,ui){
      let sub_key  = $("#master_server").val()
      let sub_data = master_sub_server[sub_key]
      populate_sub_server_select(sub_data)
      
   
   }
    
   function sub_server_change(event,ui){
      populate_table()
   
   }
    </script>`
    
  return return_value
    
    
}
