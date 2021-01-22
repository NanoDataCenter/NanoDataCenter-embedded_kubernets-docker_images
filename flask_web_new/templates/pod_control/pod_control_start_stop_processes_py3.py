from .pod_base_class_py3 import Pod_Base_Class
from templates.Base_Multi_Template_Class_py3  import Base_Multi_Template_Class
from flask import request
import json
import datetime

class Pod_Processor_Control(Base_Multi_Template_Class,Pod_Base_Class):
   def __init__(self,base_self,parameters = None):
       Pod_Base_Class.__init__(self,base_self)
       Base_Multi_Template_Class.__init__(self,base_self,parameters)
       
       
       
   def load_processes(self):
       param = request.get_json()
      
       processor_id = int(param["processor"])
       
       if processor_id >= len(self.processor_names):
          return "BAD"
       else:
          result = self.handlers[processor_id]["WEB_DISPLAY_DICTIONARY"].hgetall()
          result_json = json.dumps(result)
          
          return result_json.encode()
          

   def manage_processes(self):
       param = request.get_json()
      
       processor_id = int(param["processor"])
       process_state_json = param["process_data"]
       process_state = json.loads(process_state_json)
       if processor_id >= len(self.processor_names):
          return "BAD"
       else:
          
          self.handlers[processor_id]["WEB_COMMAND_QUEUE"].push(process_state)
          return json.dumps("SUCCESS")


   def application_page_contruction(self):
       add_ajax_handler = self.base_self.add_ajax_handler
       self.ajax_names={}

       
       self.ajax_names["load_processor"] = "/ajax/pod_control/load_processor_process"
       add_ajax_handler(self.ajax_names["load_processor"],self.load_processes,methods=["POST"])

       self.ajax_names["manage_processors"] = "/ajax/pod_control/manage_processes"
       add_ajax_handler(self.ajax_names["manage_processors"],self.manage_processes,methods=["POST"])



   def load_javascript_preamble(self):
       return_value = []
       return_value.append('<script type="text/javascript" >')
       return_value.append('False = false')
       return_value.append('True = true')
       return_value.append('None = null')
       return_value.append('display_list =' + json.dumps(self.display_list))
       
       return_value.append('command_queue_key ="WEB_COMMAND_QUEUE"')
       return_value.append('process_data_key = "WEB_DISPLAY_DICTIONARY"') 
       return_value.append('processor_id =' + str(self.processor_id))                     
       return_value.append('load_process ="' + self.ajax_names["load_processor"]+'"') 
       return_value.append('manage_process ="' + self.ajax_names["manage_processors"]+'"')                         
       return_value.append('</script>')
       return "\n".join(return_value)
 
   def application_page_generation(self,processor_id,data):
       self.processor_id = processor_id
       self.processor_name = self.processor_names[processor_id]
       self.processor_names = self.processor_names
       self.display_list = self.handlers[ self.processor_id]["WEB_DISPLAY_DICTIONARY"].hkeys()
 
       return_value = []
       return_value.append(self.load_html())
       return_value.append(self.load_javascript())
       return "\n".join(return_value)


       
       

   def load_html(self):
       return_value = []
       return_value.append(self.load_processor_selection_html())
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
       self.mp.processor_id = self.processor_id
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
   json_object["processor"]  = processor_id

  
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
      html +=  "<h3>No processors managed </h3>";
	  
	 
	 
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
  json_object["processor"]  = processor_id
  json_object["process_data"] = temp_json
  ajax_post_confirmation(manage_process, json_object,"Do you want to start/kill selected processes ?",
                            "Changes Made", "Changes Not Made") 
  

}
      


 
function change_processor(event,ui)
{
  current_page = window.location.pathname
  

  
  current_page = current_page+"?"+$("#processor_select")[0].selectedIndex
  window.location.href = current_page
}
 

$(document).ready(
 function()
 {
   load_data() 
   
   $("#processor_select").val( {{ self.processor_id  }});
   $("#processor_select").bind('change',change_processor)   
   $("#refresh_b").bind("click",refresh_data)
   $("#change_state").bind("click",change_process_status)

 }
)
</script>
       '''
       
   