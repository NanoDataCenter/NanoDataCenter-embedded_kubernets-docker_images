
function refresh_data(event,ui)
{
       
       load_data();
}  

function load_data()
{
   json_object = {}
   json_object["processor_id"]  = processor_id

  
   ajax_post_get(load_process,json_object, getQueueEntries, "Initialization Error!!!!") 
}

     

function getQueueEntries( data )
{
  
   var temp_index;
   var temp
   var html;

   data_ref = data
   $("#queue_elements").empty();
   
      
   if( display_list.length == 0 )
   {
      var html = "";
      html +=  "<h3>No Containers managed </h3>";
	  
	 
	 
   }
   else
   {
       var html = "";
       
	      html += '';

       for( i = 0; i < display_list.length; i++ )
       {
          temp_index = i +1;  
          id = "check"+i
          html += "<div>"
          
          name = display_list[i]
	      temp  = data_ref[name]
          data1 = 'Container: '+temp.name+" -- Enabled: "+temp.enabled+"  -- Active: "+
                    temp.active+" --  Error State: "+temp.error 
          data = '<label for='+id+">"+data1+" </label>"
          html += '<div class="btn-group" >'
          html += '<label class=class="btn  btn-toggle" for="'+id+'">'
          html +=  '<input type="checkbox" class="btn  btn-toggle"  id="'+id+'"    name="option"   >'+data
               +'</label>'
          html += '</div>'
          html += '</div>'
           
             
           
        }
        html += "</div>";
        
   } // if
      
     
   $("#queue_elements").append (html)



   for( i = 0; i < display_list.length; i++ )
   {
       name = display_list[i]
	   temp  = data_ref[name]
       id = "#check"+i
       if(temp.enabled == True)
       {
           $(id).prop('checked', true)
       }
       else
       {
           $(id).prop('checked', false)
        }
	 
   }   
 

}
        



function  change_container_status(event,ui)
{
	 
	  
   for( i=0;i<display_list.length;i++)
   {
	                  
        name = display_list[i]
	    temp  = data_ref[name]
        id = "#check"+i
   
	     if( $(id).is(":checked") == true )
	     {
	         data_ref[name].enabled = true
	     }
	     else
	     {
	        
	        data_ref[name].enabled = false
	     }
     
  }

  let temp_json = JSON.stringify(data_ref)
  json_object = {}
  json_object["processor_id"]  = processor_id
  json_object["process_data"] = temp_json
  ajax_post_confirmation(manage_process, json_object,"Do you want to start/kill selected containers ?",
                            "Changes Made", "Changes Not Made") 
  

}
      


 
function change_processor(event,ui)
{
  current_page = window.location.href
  
 
  current_page = current_page.slice(0,-2)
  
  current_page = current_page+"/"+$("#processor_select")[0].selectedIndex
  window.location.href = current_page
}
 

$(document).ready(
 function()
 {
   load_data() 
   
   $("#processor_select").val( {{ processor_id|int  }});
   $("#processor_select").bind('change',change_processor)   
   $("#refresh_b").bind("click",refresh_data)
   $("#change_state").bind("click",change_container_status)

 }
)
     
     
     

      


