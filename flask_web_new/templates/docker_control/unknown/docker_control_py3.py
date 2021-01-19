
from .docker_container_base_py3 import Docker_Base_Class
rom templates.Base_Multi_Template_Class_py3  import Base_Multi_Template_Class

class Docker_Control(Docker_Base_Class):
   def __init__(self,base_self,parameters = None):
       Docker_Base_Class.__init__(self,base_self)
       Base_Multi_Template_Class.__init__(self,base_self,parameters)



   def application_generation(self,processor_id,data):
       self.processor_id = processor_id
       self.processor_name = self.processor_names[processor_id]
       self.display_list = self.container_control_structure[processor_name]["WEB_DISPLAY_DICTIONARY"].hkeys()
       return self.generate_template(processor_id)
     

 
   



   def generate_template(self)
       return_value = []
       return_value.append(self.process_html)
       return_value.append(self.process_load_javascript)
       return "\n".join(return_value)
 
   def generate_processor_names(self):
      return_value = {}
      for i in range(0,len(self.processor_names)):
          return.append('<option value="'+i+'">'+self.processor_names[i]+'</option>')
      return "\n".join(return_value)
 
   def process_html(self):
      self.mp.generate_processor_names = self.generate_processor_names
      return self.mp.macro_expand_start("<<",">>",self.load_html())
   
 
   def html_raw(self)
       return = '''
       
<div class="container">
<center>
<h4>Select Processor</h4>
</center>

<div id="select_tag">
<center>
<select id="processor_select">
  {{(self.generate_processor_names )}}

  
</select>
</center>
</div>
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

   def procecess_load_javascript(self)
       self.mp.manage_process = self.ajax_name["manage_processes"]
       self.mp.load_process = self.ajax_name["load_processes"]
       self.mp.processor_id = self.processor_id
       return self.mp.macro_expand_start("<<",">>",self.load_javascript())
       
       
   def load_javascript(self):
      
       return '''
<script>
manage_process = "{{self.mp.manage_process}}"
load_process = "{{self.load_process}}"

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
  current_page = window.location.path_name
 
  
  
  current_page = current_page+"?"+$("#processor_select")[0].selectedIndex
  window.location.href = current_page
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