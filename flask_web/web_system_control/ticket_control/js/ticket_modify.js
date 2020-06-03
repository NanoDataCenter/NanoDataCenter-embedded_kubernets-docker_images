
<script type="text/javascript" >
  function cancel_modify_ticket(event)
  {
      $("#display_entry").hide();
      $("#modify_entry").hide();
      $("#main_panel").show()
      
  }
 
  //"fields":[ "id INTEGER PRIMARY KEY  AUTOINCREMENT","active Int","create_timestamp FLOAT","close_timestamp FLOAT","type Int","subtype Text",
   //                            "title Text","description TEXT","resolution TEXT"   ]} )
  function save_modify_ticket(event)
  {
     
      
      table_item["resolution"] =$("#modify_problem_resolution").val()
      if( $("#modify-active").is(':checked') == true )
      {
          table_item["active"] = 1
      }
      else
      {
          table_item["active"] = 0
      }
      //alert(table_item["active"])
      var result = confirm("do you wish to make changes?");  
      if( result == true )
           ajax_post_get(modify_link,table_item, success_function, "entry not modified") 
      
      
  }  
  
  function success_function()
  {
      
      window.location = full_link
  }
           
  function initialize_modify_entry()
  {
      
      $("#cancel_modify_ticket").bind("click",cancel_modify_ticket)
      
      $("#save_modify_ticket").bind("click",save_modify_ticket)
      
      $("#return_display_ticket").bind("click",cancel_modify_ticket)
      
  }
  
  function load_display_data(entry)
  {

      $("#display_entry").show();
      $("#modify_entry").hide();
      $("#main_panel").hide()
      $("#main_panel").hide()  
      $("#display_title").html("Title: "+entry.title)      
      $("#display_type").html("Type: "+entry.type)
      $("#display_subtype").html("Subtype: "+entry.subtype)
      $("#display_creation_time_stamp").html("Creation Time: "+entry.create_timestamp)
      
      
      $("#display_problem_description").val(entry.description)
      $("#display_status").html("Status: "+entry.active)
      if(entry.active != "active")
      {
      $("#display_resolution_time_stamp").html("Resolution Time: "+entry.close_timestamp)
      }
      $("#display_problem_resolution").val(entry.resolution)
  }
  function load_modify_data(entry)
  {
   
      $("#display_entry").hide();
      $("#modify_entry").show();
      $("#main_panel").hide()  
      $("#modify_title").html("Title: "+entry.title)   
      $("#modify_type").html("Type: "+entry.type)
      $("#modify_subtype").html("Subtype: "+entry.subtype)
      
      
      $("#modify-active").prop("checked", true)
     
      $("#modify_creation_time_stamp").html("Creation Time: "+entry.create_timestamp)
      
      $("#modify_problem_description").val(entry.description)
      $("#modify_problem_resolution").val(entry.resolution)

  }     
  
 
</script>