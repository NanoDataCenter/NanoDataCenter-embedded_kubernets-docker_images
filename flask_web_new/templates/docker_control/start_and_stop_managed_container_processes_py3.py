   def process_control(self,container_id):
      container_name = self.managed_container_names[container_id]
      display_list = self.docker_performance_data_structures[container_name]["WEB_DISPLAY_DICTIONARY"].hkeys()
      
      return self.render_template(self.path_dest+"/docker_process_control",
                                  display_list = display_list, 
                                  command_queue_key = "WEB_COMMAND_QUEUE",
                                  process_data_key = "WEB_DISPLAY_DICTIONARY",
                                  container_id = container_id,
                                  containers = self.managed_container_names,
                                  load_process = '"'+self.slash_name+'/load_processes/load_process'+'"',
                                  manage_process =  '"'+self.slash_name+'/manage_processes/change_process'+'"' )
                                  
                                  
{% extends "base_template" %}

{% block application_javascript %}
  <script type="text/javascript" >
       False = false
       True = true
       None = null
       display_list = {{display_list}} 
       command_queue_key = "{{ command_queue_key}}"
       process_data_key = "{{ process_data_key}}"
       container_id = {{container_id}}                         
       load_process = {{load_process}}
       manage_process = {{manage_process}}                          
      
      {% include "js/docker_process_control/docker_process_control.js" %}
      </script>

  
{% endblock %}

{% block application %}
<div class="container">
<center>
<h4>Select Container</h4>
</center>

<div id="select_tag">
<center>
<select id="container_select">
  {% for item in containers %}
  
  <option value="{{loop.index0}}">{{item}}</option>
  {% endfor %}
  
</select>
</center>
</div>
<div style="margin-top:20px"></div>

  <h4>Refresh State</h4>       
   <button type="button" id="refresh_b">Refresh</button>
   <div style="margin-top:20px"></div>
   <h4>Edit Managed Processes</h4>
   
   <h4>Toggle Check Box to Change State  -- Check to Enable  Uncheck to Disable</h4>
   
   <button type="button" id="change_state">Click to Change State</button> 
   <div style="margin-top:20px"></div>
   <div id="queue_elements">
   </div>
</div>

{% endblock %}function refresh_data(event,ui)
{
       
       load_data();
}  

function load_data()
{
   json_object = {}
   json_object["container"]  = container_id

  
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
          data1 = 'Process: '+temp.name+" -- Enabled: "+temp.enabled+"  -- Active: "+
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
        



function  change_process_status(event,ui)
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
  json_object["container"]  = container_id
  json_object["process_data"] = temp_json
  ajax_post_confirmation(manage_process, json_object,"Do you want to start/kill selected processes ?",
                            "Changes Made", "Changes Not Made") 
  

}
      


 
function change_container(event,ui)
{
  current_page = window.location.href
  
 
  current_page = current_page.slice(0,-2)
  
  current_page = current_page+"/"+$("#container_select")[0].selectedIndex
  window.location.href = current_page
}
 

$(document).ready(
 function()
 {
   load_data() 
   
   $("#container_select").val( {{ container_id|int  }});
   $("#container_select").bind('change',change_container)   
   $("#refresh_b").bind("click",refresh_data)
   $("#change_state").bind("click",change_process_status)

 }
)
     
     
     

      


