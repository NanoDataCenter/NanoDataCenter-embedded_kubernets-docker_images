package construct_schedule


import(
    //"fmt"
   
)






func generate_copy_table_html()string{
    
  return_value :=
  `
 <div class="container" id="copy_section">
 
     
    <h3>Enter New Table Name</h3>
    
    
       <div>
        <input type="button" id = "copy_name_continue" value="Copy"  data-inline="true"  /> 
        <input type="button" id = "copy_name_abort" value="Abort" data-inline="true"  /> 
       </div>
     <input type="text" id="copy_name_input">
    
    
    </div>
    
    
`
 return return_value
}

func js_generate_copy_table()string{

  return_value := 
    ` <script type="text/javascript"> 
      
       function initialize_copy_panel(){
       
         $("#copy_name_continue").bind('click',table_copy_continue)
         $("#copy_name_abort").bind('click',table_copy_abort)
       
    
       
       
       }
       function copy_table_generate_panel(){
          hide_all_sections()
          $("#copy_section").show()
       }
       function table_copy_continue(){
         
         hide_all_sections()
         $("#main_section").show()
       }
       function table_copy_abort(){
         hide_all_sections()
         $("#main_section").show()
       }
    </script>`
    
    
  return return_value
    
    
}   

    
