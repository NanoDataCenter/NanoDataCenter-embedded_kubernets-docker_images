
function  change_system_status(event,ui)
{
   var count = 0;
   reset_data = {}
   
   if( $("#pod_reset").is(":checked") == true )
   {  
      count = count+1;
	  reset_data["pod"] = true
   }
   else
   {
	        
	  reset_data["pod"] = false
   }  
   
  
   service_list = []
   for( i=0;i< services.length;i++)
   {
	                  
        name = services[i]
	    
        id = "#service"+ name
         
	     if( $(id).is(":checked") == true )
	     {
             
             count = count+1
	         service_list.push(name)
	     }
	     else
	     {
	         ; // do nothing for now
	     }
     
  }
  reset_data["services"] = service_list
  container_list = []
  for( i=0;i< containers.length;i++)
  {
	                  
        name = containers[i]
	    
        id = "#container"+ name
 
	     if( $(id).is(":checked") == true )
	     {
             
             count = count+1
	         container_list.push(name)
	     }
	     else
	     {
	         ; // do nothing for now
	     }
     
  }
  reset_data["containers"] = container_list
  processor_id
  reset_data["processor_id"] = processor_id
  alert(JSON.stringify(reset_data))
  alert(count)
 
  ajax_post_confirmation(ajax_handler, reset_data,"Do you want to upgrade selected components ?",
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
   
   
   $("#processor_select").val( {{ processor_id|int  }});
   $("#processor_select").bind('change',change_processor)   
   $("#refresh_b").bind("click",change_system_status)
   

 }
)
     
     
     

      


 