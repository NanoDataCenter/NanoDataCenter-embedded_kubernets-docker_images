
from .node_base_class_py3 import Node_Base_Class
from templates.Base_Multi_Template_Class_py3  import Base_Multi_Template_Class
from flask import request
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers 
import json

class Manage_Containers(Base_Multi_Template_Class,Node_Base_Class):
   def __init__(self,base_self,parameters = None):
       Base_Multi_Template_Class.__init__(self,base_self,parameters)
       Node_Base_Class.__init__(self,base_self)
       


   def change_container_state(self):
       param = request.get_json()   
       print("param",param)
      
       
       processor_id = int(param["processor_id"])
      
       
      
       if processor_id >= len(self.processor_names):
          
          return "BAD"
       else:
          processor_name = self.processor_names[processor_id]
         
          self.handlers[processor_id]["DOCKER_COMMAND_QUEUE"].push(param)
          return json.dumps("SUCCESS")

   def load_containers(self):
       param = request.get_json()
      
       processor_id = int(param["processor_id"])
       
       if processor_id  >= len(self.processor_names):
          return "BAD"
       else:
          processor_name = self.processor_names[processor_id]
          result = self.handlers[processor_id]["DOCKER_DISPLAY_DICTIONARY"].hgetall()

          result_json = json.dumps(result)
          
          return result_json.encode()

   def application_page_contruction(self):
       add_ajax_handler = self.base_self.add_ajax_handler
       self.ajax_names={}

       
       self.ajax_names["load_containers"] = "/ajax/node_control/load_containers"
       add_ajax_handler(self.ajax_names["load_containers"],self.load_containers,methods=["POST"])

       self.ajax_names["change_container_state"] = "/ajax/node_control/change_container_state"
       add_ajax_handler(self.ajax_names["change_container_state"],self.change_container_state,methods=["POST"])

   

 


   def application_page_generation(self,processor_id,data):
       self.processor_id = processor_id
       self.processor_name = self.processor_names[processor_id]
       handlers  = self.handlers
       self.display_list = self.handlers[ self.processor_id]["DOCKER_DISPLAY_DICTIONARY"].hkeys()
       return self.generate_template()
     
   def generate_template(self):
       return_value = []
       return_value.append(self.process_html())
       return_value.append(self.process_javascript())
       return "\n".join(return_value)
 
 
 
   def process_html(self):
       return_value = []
       return_value.append(self.load_processor_selection_html())
       return_value.append(self.mp.macro_expand_start("{{","}}",self.load_html()))
       return "\n".join(return_value)
 


   def load_html(self):
       return_value = []
       
       return_value.append(self.load_raw_html())
       return "\n".join(return_value)
   
   
   def load_raw_html(self):
       return '''

<div style="margin-top:20px"></div>

  <h4>Refresh State</h4>       
   <button type="button" id="refresh_b">Refresh</button>
   <div style="margin-top:20px"></div>
   <h4>Select Action </h4>
   <select name="Action" id="action_select">
   <option value="0">No Action</option>
   <option value="1">Start/Stop Containers</option>
   <option value="2">Update Container Image</option>
   <option value="3">Update All Container Image</option>
   <option value="4">Reboot Processor</option>
  </select>
   
   <h5 id="toggle_help">Toggle Check Box to Change State  -- Check to Enable  Uncheck to Disable</h5>
   <div style="margin-top:20px"></div>
   <div id="queue_elements">
   </div>
</div>

       '''
   def load_javascript_preamble(self):
       return_value = []
       return_value.append('<script type="text/javascript" >')
       return_value.append('False = false')
       return_value.append('True = true')
       return_value.append('None = null')
       return_value.append('display_list =' + json.dumps(self.display_list))
       
       return_value.append('command_queue_key ="DOCKER_COMMAND_QUEUE"')
       return_value.append('process_data_key = "DOCKER_DISPLAY_DICTIONARY"') 
       return_value.append('processor_id =' + str(self.processor_id))                     
       return_value.append('load_containers ="' + self.ajax_names["load_containers"]+'"') 
       return_value.append('manage_containers ="' + self.ajax_names["change_container_state"]+'"')                         
       return_value.append('</script>')
       return "\n".join(return_value)
       
   def process_javascript(self):
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
        




      

 
  

function change_processor(event,ui)
{
  current_page = window.location.pathname
  

  
  current_page = current_page+"?"+$("#processor_select")[0].selectedIndex
  window.location.href = current_page
} 
function change_container(event,ui)
{
  current_page = window.location.pathname
  

  
  current_page = current_page+"?"+$("#container_select")[0].selectedIndex
  window.location.href = current_page
}



function  action_handler(event,ui)
{
   json_object = []
   
   data_ref = {}
   for( i=0;i<display_list.length;i++)
   {
	                  
        name = display_list[i]
        data_ref[name] = {}
	  
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

  
  json_object = {}
  json_object["processor_id"]  = parseInt(processor_id);
  json_object["command"] = parseInt($("#action_select").val()) 
  json_object["items"] = data_ref
  
  //let temp_json = JSON.stringify(data_ref)
  
  if( parseInt($("#action_select").val()) == 1)
  {
     $("#action_select").val(0)
     ajax_post_confirmation(manage_containers, json_object,"Do you want to start/stop selected containers ?",
                            "Changes Made", "Changes Not Made") 
  }
  if( parseInt($("#action_select").val()) == 2)
  {
     $("#action_select").val(0)
     ajax_post_confirmation(manage_containers, json_object,"Do you want to upgrade selected containers ?", "Changes Made", "Changes Not Made")
  }
  if( parseInt($("#action_select").val()) == 3)
  {
     $("#action_select").val(0)
     ajax_post_confirmation(manage_containers, json_object,"Do you want to upgrade all containers ?",
                            "Changes Made", "Changes Not Made") 
  }
  if( parseInt($("#action_select").val()) == 4)
  {
     $("#action_select").val(0)
     ajax_post_confirmation(manage_containers, json_object,"Do you want to reboot selected processor ?",
                            "Changes Made", "Changes Not Made") 
  }
}
      
    



$(document).ready(
 function()
 {
   load_data() 

   $("#processor_select").val( {{ self.processor_id  }});
   $("#processor_select").bind('change',change_processor)  
   $("#action_select").val( "0");
   $("#action_select").bind('change',action_handler)   
   $("#refresh_b").bind("click",refresh_data)
 

 }
)
</script>
       '''
       
   