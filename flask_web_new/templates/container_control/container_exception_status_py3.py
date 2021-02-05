from .docker_container_base_py3 import Docker_Base_Class
from templates.Base_Multi_Template_Class_py3  import Base_Multi_Template_Class
from flask import request
import json
import datetime

class Container_Exception_Status(Base_Multi_Template_Class,Docker_Base_Class):
   def __init__(self,base_self,parameters = None):
       Docker_Base_Class.__init__(self,base_self)
       Base_Multi_Template_Class.__init__(self,base_self,parameters)
  
   def application_page_generation(self,container_id,data):
       if container_id >= len(self.managed_container_names):
            container_id = len(self.managed_container_names)-1
       self.container_id = container_id
       self.container_names = self.managed_container_names
       self.container_name = self.managed_container_names[container_id]
      
       self.container_exceptions = self.handlers[self.container_name]["ERROR_HASH"].hgetall()

       for i in self.container_exceptions.keys():
           if "time" in self.container_exceptions[i]:
               temp =self.container_exceptions[i]["time"]
               self.container_exceptions[i]["time"] = datetime.datetime.utcfromtimestamp(temp).strftime('%Y-%m-%d %H:%M:%S')
           
           temp = self.container_exceptions[i]["error_output"]
           self.container_exceptions[i]["error_output"] = [temp]

       return self.generate_template()

     
   def generate_template(self):
       return_value = []
       return_value.append(self.process_html())
       return_value.append(self.process_load_javascript())
       return "\n".join(return_value)
 
 
 
   def generate_log_data(self):
       return_value = []
       for i in self.container_exceptions.keys():
           return_value.append('<div style="margin-top:10px"></div>')
           return_value.append('<h5>'+self.container_exceptions[i]["script"]+'</h5>')
           return_value.append('<ul>')  
           return_value.append('<li>name: '+self.container_exceptions[i]["script"]+ '</li>')
           if "time" in self.container_exceptions[i]:
                  return_value.append('<li>time: '+self.container_exceptions[i]["time"]+ '</li>')
           return_value.append('<li>exception stack:</li>')
           return_value.append("<ul>")
           for j in self.container_exceptions[i]["error_output"]:
               return_value.append('<li>'+j+'</li>')
           return_value.append("</ul></ul>")
           
       return "\n".join(return_value)
 
   def process_html(self):
       return_value = []
       return_value.append(self.load_container_selection_html())
       self.mp.generate_log_data = self.generate_log_data
       self.mp.container_name = self.container_name
       return_value.append(self.mp.macro_expand_start("{{","}}",self.load_html()))
       return "\n".join(return_value)

   def load_html(self):
       return '''
<div style="margin-top:20px"></div>
<h4>Exception Status for Container {{self.container_name}} </h4>
{{(self.generate_log_data) }}
</div>
     '''


   def process_load_javascript(self):
       
       return_value = []
       self.mp.container_id = self.container_id
       return_value.append(self.load_container_control_javascript())
       return_value.append( self.mp.macro_expand_start("{{","}}",self.load_javascript()))
       return "\n".join(return_value)


   def load_javascript(self):
       return '''
   
<script type="text/javascript" >
False = false
True = true
None = null
container_id = {{self.container_id}}                         
                               
      


 $(document).ready(
 function()
 {
   
   
   $("#container_select").val( {{ self.container_id }});
   $("#container_select").bind('change',change_container)  
  


 }
)
 

</script>
       '''
   
   
   
 