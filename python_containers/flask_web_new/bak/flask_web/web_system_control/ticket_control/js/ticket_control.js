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
           ajax_post_get(delete_link, item, success_function, "entry not deleted") 
     
      
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

   
   if( choice == "add_entry")
   {   
       
       $("#show_new_entry").show();
       $("#main_panel").hide()
   }


 
   
   if( choice =="modify_entry")
   {   
       index = find_index()
       
       if( index >= 0 )
       {
          table_item = table_data[index]
          item = table_data[index]
          //alert(item.active)
          if( item.active == "active")
          {
             //alert("modify data")
             load_modify_data(item)
          }
          else
          {
             //alert("display data")
             load_display_data(item) 
          }              
         
       }       
       else
       {
           set_status_bar("No Resources Selected !!!!")
       }  
   }
   if( choice =="display_entry")
   {   
       index = find_index()
       //alert(index)
       if( index >= 0 )
       {
         item = table_data[index]
         //alert("display data")
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
   if( choice == "search_entry")
   {   
       ;//
   }



   $("#action-choice")[0].selectedIndex = 0;
   
  }
  
  
  
   $(document).ready(
      function()
      {
        initialize_new_entry()
        initialize_modify_entry()
        $("#show_new_entry").hide();
        $("#search_entry").hide();
        $("#modify_entry").hide();
        $("#display_entry").hide();
        $("#action-choice").bind('change',main_menu)
        $("#action-choice")[0].selectedIndex = 0;
      }
      ) 

 
      </script>
