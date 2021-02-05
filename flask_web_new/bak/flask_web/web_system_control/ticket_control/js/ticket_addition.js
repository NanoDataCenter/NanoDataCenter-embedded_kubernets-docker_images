
<script type="text/javascript" >
  function new_entry_cancel(event)
  {
      $("#show_new_entry").hide();
      $("#main_panel").show()
      
  }
 
  //"fields":[ "id INTEGER PRIMARY KEY  AUTOINCREMENT","active Int","create_timestamp FLOAT","close_timestamp FLOAT","type Int","subtype Text",
   //                            "title Text","description TEXT","resolution TEXT"   ]} )
  function save_new_entry(event)
  {
      // gather data
      data = ["a","b"] // test data
      data = {}
      data["active"] = 1
      data["title"] = $("#title").val();
      data["type"]  = $("#new-choice").val()
      data["subtype"] =    $("#new_subtype").val()
      data["description"] =$("#new_problem_description").val()
      data["resolution"] = ""
      var result = confirm("do you wish to make changes?");  
      if( result == true )
           ajax_post_get(add_link, data, success_function, "entry not added") 
      
      
  }  
  
  function success_function()
  {
      
      window.location = full_link
  }
           
  function initialize_new_entry()
  {
      $("#cancel_new_ticket").bind("click",new_entry_cancel)
      $("#save_new_ticket").bind("click",save_new_entry)
      
  }
</script>