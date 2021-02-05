 <script type="text/javascript" >
  
  table_item = ""
  function find_index()
  {
   
   var i;
   
   for( i=0; i< table_data.length; i++)
   {  
       if( $("#ticket_log_"+i).is(':checked') == true )
       {
          return i;
       }
    }
    return  -1;  // no item selected
    
   }

  function delete_entry(item)
  {
      var result = confirm("do you wish to delete item?");  
      if( result == true )
      {
           ajax_post_get(delete_link, item, delete_success_function, "entry not deleted") 
      }
      
  } 
  function delete_success_function()
  {
      
      window.location = full_link
  }  
  function reset_entry()
  {
    var result = confirm("do you wish to delete item?");  
    if( result == true )
    {
         window.location = full_link  
    }  
  }
  
  function delete_success_function()
  {
      
      window.location = full_link
  }  
  
  function main_menu(event,ui)
{
   var index
   var choice

   choice = $("#action-choice").val()
   //alert(choice)
   if( choice == "nop")
   {
       ; // do nothing
   }

   

   if( choice =="display_entry")
   {   
       index = find_index()
       //alert(index)
       if( index >= 0 )
       {
         item = table_data[index]
         load_display_data(item) 
          
       }       
       else
       {
           set_status_bar("No Resources Selected !!!!")
       }  
   }
   if( choice =="delete_entry")
   {   
       index = find_index()
       //alert(index)
       if( index >= 0 )
       {
         item = table_data[index]
         //alert("delete_entry")
         delete_entry(item) 
          
       }       
       else
       {
           set_status_bar("No Resources Selected !!!!")
       }  
   }
   
   if( choice =="reset_entry")
   {   
     reset_entry()
   }   
   
   if( choice == "search_entry")
   {   
       ;//
   }



   $("#action-choice")[0].selectedIndex = 0;
   
  }
  
  
  
   $(document).ready(
      function()
      {
       
       
        initialize_display_entry()
        $("#action-choice").bind('change',main_menu)
        $("#action-choice")[0].selectedIndex = 0;
      }
      ) 

 
      </script>
