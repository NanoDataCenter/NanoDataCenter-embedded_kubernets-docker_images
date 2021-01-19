from .docker_container_base_py3 import Docker_Base_Class
from templates.Base_Multi_Template_Class_py3  import Base_Multi_Template_Class
from flask import request
import json

class Docker_Container_Control(Base_Multi_Template_Class,Docker_Base_Class):
   def __init__(self,base_self,parameters = None):
       Docker_Base_Class.__init__(self,base_self)
       Base_Multi_Template_Class.__init__(self,base_self,parameters)

   def change_processor_state(self):
       param = request.get_json()
      
       processor_id = int(param["processor_id"])
       process_state_json = param["process_data"]
       process_state = json.loads(process_state_json)
      
       if processor_id >= len(self.processor_names):
          
          return "BAD"
       else:
          processor_name = self.processor_names[processor_id]
         
          self.container_control_structure[processor_name]["WEB_COMMAND_QUEUE"].push(process_state)
          return json.dumps("SUCCESS")

   def load_containers(self):
       param = request.get_json()
      
       processor_id = int(param["processor_id"])
       
       if processor_id  >= len(self.processor_names):
          return "BAD"
       else:
          processor_name = self.processor_names[processor_id]
          result = self.container_control_structure[processor_name]["WEB_DISPLAY_DICTIONARY"].hgetall()

          result_json = json.dumps(result)
          
          return result_json.encode()
   

   def application_page_contruction(self):
       add_ajax_handler = self.base_self.add_ajax_handler
       self.ajax_names={}

       self.ajax_names["change_processor_state"] = "/ajax/manage_containers/change_processor_state"
       add_ajax_handler(self.ajax_names["change_processor_state"],self.change_processor_state,methods=["POST"])

       self.ajax_names["load_containers"] = "/ajax/manage_containers/load_containers, change_processor_state"
       add_ajax_handler(self.ajax_names["load_containers"],self.load_containers,methods=["POST"])


   def application_page_generation(self,processor_id,data):
       self.processor_id = processor_id
       self.processor_name = self.processor_names[processor_id]
       self.display_list = self.container_control_structure[self.processor_name]["WEB_DISPLAY_DICTIONARY"].hkeys()
       return self.generate_template()
     
   def generate_template(self):
       return_value = []
       return_value.append(self.process_html())
       return_value.append(self.process_load_javascript())
       return "\n".join(return_value)
 
 
 
   def process_html(self):
       return_value = []
       return_value.append(self.load_processor_selection_html())
       return_value.append(self.mp.macro_expand_start("{{","}}",self.load_html()))
       return "\n".join(return_value)
 
   def load_html(self):
       return '''
       

<div style="margin-top:20px"></div>

  <h4>Refresh State</h4>       
   <button type="button" id="refresh_b">Refresh</button>
   <div style="margin-top:20px"></div>
   <h4>Edit Managed Containers</h4>
   
   <h4>Toggle Check Box to Change State  -- Check to Enable  Uncheck to Disable</h4>
   
   <button type="button" id="change_state">Click to Change State</button> 
   <div style="margin-top:20px"></div>
   <div id="queue_elements">
   </div>
</div>
       '''

   def process_load_javascript(self):
       self.mp.change_processor_state = self.ajax_names["change_processor_state"]
       self.mp.load_containers = self.ajax_names["load_containers"]
       self.mp.processor_id = self.processor_id
       self.mp.display_list_json = json.dumps(self.display_list)
       return_value = []
       return_value.append(self.load_processor_control_javascript())
       return_value.append( self.mp.macro_expand_start("{{","}}",self.load_javascript()))
       return "\n".join(return_value)
       
       
   def load_javascript(self):
      
       return '''
<script>
change_processor_state = "{{self.change_processor_state}}"
load_containers = "{{self.load_containers}}"
processor_id = "{{self.processor_id}}"
display_list_json = '{{self.display_list_json}}'
display_list =  JSON.parse(display_list_json)
True = true
False = false
function refresh_data(event,ui)
{
       
       load_data();
}  

function load_data()
{
   json_object = {}
   json_object["processor_id"]  = processor_id

  
   ajax_post_get(load_containers,json_object, getQueueEntries, "Initialization Error!!!!") 
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
  ajax_post_confirmation(change_processor_state, json_object,"Do you want to start/kill selected containers ?",
                            "Changes Made", "Changes Not Made") 
  

}
      


 

 

$(document).ready(
 function()
 {
   load_data() 
   
   $("#processor_select").val( {{ self.processor_id }});
   $("#processor_select").bind('change',change_processor)   
   $("#refresh_b").bind("click",refresh_data)
   $("#change_state").bind("click",change_container_status)

 }
)
</script>

   '''