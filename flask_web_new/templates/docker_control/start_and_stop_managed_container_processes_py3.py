from .docker_container_base_py3 import Docker_Base_Class
from templates.Base_Multi_Template_Class_py3  import Base_Multi_Template_Class
from flask import request
import json

class Start_and_Stop_Managed_Container_Processes(Base_Multi_Template_Class,Docker_Base_Class):
   def __init__(self,base_self,parameters = None):
       Docker_Base_Class.__init__(self,base_self)
       Base_Multi_Template_Class.__init__(self,base_self,parameters)

   def load_container_processes(self):
       param = request.get_json()
      
       container_id = int(param["container"])
       
       if container_id >= len(self.managed_container_names):
          return "BAD"
       else:
          container_name = self.managed_container_names[container_id]
          result = self.handlers[container_name]["WEB_DISPLAY_DICTIONARY"].hgetall()
          result_json = json.dumps(result)
          
          return result_json.encode()
          

   def manage_container_processes(self):
       param = request.get_json()
      
       container_id = int(param["container"])
       process_state_json = param["process_data"]
       process_state = json.loads(process_state_json)
       if container_id >= len(self.managed_container_names):
          return "BAD"
       else:
          
          container_name = self.managed_container_names[container_id]
          self.handlers[container_name]["WEB_COMMAND_QUEUE"].push(process_state)
          return json.dumps("SUCCESS")




   def application_page_contruction(self):
       add_ajax_handler = self.base_self.add_ajax_handler
       self.ajax_names={}

       
       self.ajax_names["load_container_process"] = "/ajax/manage_containers/load_container_process"
       add_ajax_handler(self.ajax_names["load_container_process"],self.load_container_processes,methods=["POST"])

       self.ajax_names["manage_container_containers"] = "/ajax/manage_containers/manage_processes"
       add_ajax_handler(self.ajax_names["manage_container_containers"],self.manage_container_processes,methods=["POST"])



   def load_javascript_preamble(self):
       return_value = []
       return_value.append('<script type="text/javascript" >')
       return_value.append('False = false')
       return_value.append('True = true')
       return_value.append('None = null')
       return_value.append('display_list =' + json.dumps(self.display_list))
       
       return_value.append('command_queue_key ="WEB_COMMAND_QUEUE"')
       return_value.append('process_data_key = "WEB_DISPLAY_DICTIONARY"') 
       return_value.append('container_id =' + str(self.container_id))                     
       return_value.append('load_process ="' + self.ajax_names["load_container_process"]+'"') 
       return_value.append('manage_process ="' + self.ajax_names["manage_container_containers"]+'"')                         
       return_value.append('</script>')
       return "\n".join(return_value)
 
   def application_page_generation(self,container_id,data):
       self.container_id = container_id
       self.container_name = self.managed_container_names[container_id]
       self.container_names = self.managed_container_names
       self.display_list = self.handlers[self.container_name]["WEB_DISPLAY_DICTIONARY"].hkeys()
 
       return_value = []
       return_value.append(self.load_html())
       return_value.append(self.load_javascript())
       return "\n".join(return_value)


       
       

   def load_html(self):
       return_value = []
       return_value.append(self.load_container_selection_html())
       return_value.append(self.load_raw_html())
       return "\n".join(return_value)
   
   
   def load_raw_html(self):
       return '''

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

       '''

   def load_javascript(self):
       return_value = []
       return_value.append(self.load_javascript_preamble())
       self.mp.container_id = self.container_id
       return_value.append(self.mp.macro_expand_start("{{","}}",self.load_raw_javascript()))
       
       return "\n".join(return_value)


          
      
       
   def load_raw_javascript(self):
       return '''
       <script>
       function refresh_data(event,ui)
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
  current_page = window.location.pathname
  

  
  current_page = current_page+"?"+$("#container_select")[0].selectedIndex
  window.location.href = current_page
}
 

$(document).ready(
 function()
 {
   load_data() 
   
   $("#container_select").val( {{ self.container_id  }});
   $("#container_select").bind('change',change_container)   
   $("#refresh_b").bind("click",refresh_data)
   $("#change_state").bind("click",change_process_status)

 }
)
</script>
       '''
       
   