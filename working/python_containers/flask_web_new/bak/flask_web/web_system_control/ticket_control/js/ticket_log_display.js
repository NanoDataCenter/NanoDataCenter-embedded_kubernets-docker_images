
<script type="text/javascript" >
  function cancel_display_ticket(event)
  {
      $("#display_entry").hide();
     
      $("#main_panel").show()
      
  }
 
 
 
           
  function initialize_display_entry()
  {

      $("#display_entry").hide();
      $("#return_display_ticket").bind("click",cancel_display_ticket)
      
  }
  
  function load_display_data(entry)
  {
      
      $("#display_entry").show();
      $("#main_panel").hide()  
      $("#display_title").html("Title: "+entry.title)      
      $("#display_type").html("Type: "+entry.type)
      $("#display_subtype").html("Subtype: "+entry.subtype)
      $("#display_creation_time_stamp").html("Creation Time: "+entry.create_timestamp)
      
      $("#display_problem_description").val(entry.description)
      
   
      $("#display_resolution_time_stamp").html("Resolution Time: "+entry.close_timestamp)
     
      $("#display_problem_resolution").val(entry.resolution)
  }

  
 
</script>