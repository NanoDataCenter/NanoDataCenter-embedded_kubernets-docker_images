from .node_base_class_py3 import Node_Base_Class
from templates.Base_Multi_Template_Class_py3  import Base_Multi_Template_Class
from flask import request
import json
import datetime

class Node_Exception_Log(Base_Multi_Template_Class,Node_Base_Class):
   def __init__(self,base_self,parameters = None):
       Node_Base_Class.__init__(self,base_self)
       Base_Multi_Template_Class.__init__(self,base_self,parameters)
  
  


   def application_page_generation(self,processor_id,data):
       if processor_id >= len(self.processor_names):
            processor_id = len(self.processor_names)-1
       
       self.processor_id = processor_id
       self.processor_name = self.processor_names[processor_id]
       print(self.handlers[processor_id])
       processor_exceptions = self.handlers[processor_id]["ERROR_HASH"].hgetall()

       for i in processor_exceptions.keys():
           if "time" in processor_exceptions[i]:
               temp = processor_exceptions[i]["time"]
               processor_exceptions[i]["time"] = datetime.datetime.utcfromtimestamp(temp).strftime('%Y-%m-%d %H:%M:%S')
           
           temp = processor_exceptions[i]["error_output"]
           processor_exceptions[i]["error_output"] = [temp]
       self.processor_exceptions = processor_exceptions
       return self.generate_template()
       
       
       
   def generate_template(self):
       return_value = []
       return_value.append(self.process_html())
       return_value.append(self.process_load_javascript())
       return "\n".join(return_value)
 
 
 
   def generate_log_data(self):
       return_value = []
       for i in self.processor_exceptions.keys():
           return_value.append('<div style="margin-top:10px"></div>')
           return_value.append('<h5>'+self.processor_exceptions[i]["script"]+'</h5>')
           return_value.append('<ul>')  
           return_value.append('<li>name: '+self.processor_exceptions[i]["script"]+ '</li>')
           if "time" in self.processor_exceptions[i]:
                  return_value.append('<li>time: '+self.processor_exceptions[i]["time"]+ '</li>')
           return_value.append('<li>exception stack:</li>')
           return_value.append("<ul>")
           if len(self.processor_exceptions[i]["error_output"]) > 25:
              length = 25
           else:
               length = len(self.processor_exceptions[i]["error_output"])
           for k in range(0,length):
               j = self.processor_exceptions[i]["error_output"][k]
               if len(j) > 2000:
                  j = j[-2000:]
               return_value.append('<li>'+j+'</li>')
           return_value.append("</ul></ul>")
           
       return "\n".join(return_value)
 
   def process_html(self):
       return_value = []
       return_value.append(self.load_processor_selection_html())
       self.mp.generate_log_data = self.generate_log_data
       self.mp.processor_name = self.processor_name
       return_value.append(self.mp.macro_expand_start("{{","}}",self.load_html()))
       return "\n".join(return_value)

   def load_html(self):
       return '''
<div style="margin-top:20px"></div>
<h4>Exception Long for Node Control processes processor {{self.processor_name}} </h4>
{{(self.generate_log_data) }}
</div>
     '''


   def process_load_javascript(self):
       
       return_value = []
       self.mp.processor_id = self.processor_id
       return_value.append(self.load_processor_control_javascript())
       return_value.append( self.mp.macro_expand_start("{{","}}",self.load_javascript()))
       return "\n".join(return_value)


   def load_javascript(self):
       return '''
   
<script type="text/javascript" >
False = false
True = true
None = null
processor_id = {{self.processor_id}}                         
                               
      


 $(document).ready(
 function()
 {
   
   
   $("#processor_select").val( {{ self.processor_id }});
   $("#processor_select").bind('change',change_processor)  
  


 }
)
 

</script>
       '''
   
   
   
 