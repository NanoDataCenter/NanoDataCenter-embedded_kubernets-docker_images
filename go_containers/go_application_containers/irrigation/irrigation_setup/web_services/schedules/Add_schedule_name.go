
package construct_schedule


import(
    //"fmt"
   
)




func generate_create_schedule_name_html()string{
    
  return_value :=
  `
  <div class="container" id="table_name_section">
 
     
    <h3>Enter New Table Name</h3>
    
    
       <div>
        <input type="button" id = "table_name_continue" value="Contine"  data-inline="true"  /> 
        <input type="button" id = "table_name_abort" value="Abort" data-inline="true"  /> 
       </div>
     <input type="text" id="new_schedule_input">
    
    
    </div>
    
`
 return return_value
}

func js_generate_create_schedule_name()string{
  return_value := 
    ` <script type="text/javascript"> 
      
       function initialize_table_name_panel(){
       
         $("#table_name_continue").bind('click',table_name_continue)
         $("#table_name_abort").bind('click',table_name_abort)
       
    
       
       
       }
       function start_table_entry_panel(){
          hide_all_sections()
          $("#table_name_section").show()
       }
       function table_name_continue(){
         
         hide_all_sections()
         $("#table_construction").show()
       }
       function table_name_abort(){
         hide_all_sections()
         $("#main_section").show()
       }
    </script>`
    
  return return_value
 
    
}
/*
 var str = $("#myInput").val();
        alert(str);
*/
