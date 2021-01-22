from .pod_base_class_py3 import Pod_Base_Class
from templates.Base_Multi_Template_Class_py3  import Base_Multi_Template_Class
from flask import request
import json
import datetime

class Pod_Exception_Log(Base_Multi_Template_Class,Pod_Base_Class):
   def __init__(self,base_self,parameters = None):
       Pod_Base_Class.__init__(self,base_self)
       Base_Multi_Template_Class.__init__(self,base_self,parameters)
  
   def application_page_generation(self,processor_id,data):
       if processor_id >= len(self.processor_names):
            processor_id = len(self.processor_names)-1
       
       self.processor_id = processor_id
       self.processor_name = self.processor_names[processor_id]
      
       self.processor_exception_log = self.handlers[processor_id]["ERROR_STREAM"].revrange("+","-" , count=20)
       
       processor_exceptions = []
       for j in self.processor_exception_log:
           i = j["data"]
           i["timestamp"] = j["timestamp"]
           i["datetime"] =  datetime.datetime.fromtimestamp( i["timestamp"]).strftime('%Y-%m-%d %H:%M:%S')

           temp = i["error_output"]
           if len(temp) > 0:
               temp = i["error_output"]
               if len(temp) > 0:
                   temp = [temp]
                   #temp = temp.split("\n")
                   i["error_output"] = temp
                   processor_exceptions.append(i)

       self.processor_exceptions = processor_exceptions
       return self.generate_template()

     
   def generate_template(self):
       return_value = []
       return_value.append(self.process_html())
       return_value.append(self.process_javascript())
       return "\n".join(return_value)

   def process_html(self):
       self.row_index = []
       
       self.mp.generate_header_rows = self.generate_header_rows
       return_value = []
       return_value.append(self.load_processor_selection_html())
       return_value.append( self.mp.macro_expand_start("{{","}}",self.process_html_raw()))
       return "\n".join(return_value)

   

   def generate_header_rows(self):
       return_value = []
     
       for i in range(0,len(self.processor_exceptions)):
           
           datetime = self.processor_exceptions[i]["datetime"]
           script   = self.processor_exceptions[i]["script"]
          
           error_lines = self.processor_exception_log[i]["data"]["error_output"]
           return_value.append('<tr data-tt-id="'+str(i+1)+'">')
           return_value.append('<td>'+datetime+'</td>')
           return_value.append('<td>'+script+'</td>')
           return_value.append('</tr>')
           #<tr data-tt-id="1.1" data-tt-parent-id="1">
           for j in range(0,len(error_lines)):
               return_value.append('<tr data-tt-id="'+str(i+1)+"."+str(j+1)+'"  data-tt-parent-id="'+str(i+1)+'" >')
               return_value.append('<td>')
               return_value.append('-->')
               return_value.append('</td>')
               return_value.append('<td>')
               return_value.append(error_lines[j])
               return_value.append('</td>')
               
               return_value.append('</tr>')
       return "\n".join(return_value)

   def process_html_raw(self):
       return '''
<link href="/static/css/jquery.treetable.css" rel="stylesheet" type="text/css" />
<link href="/static/css/jquery.treetable.theme.default.css" rel="stylesheet" type="text/css" />
<link href="/static/css/screen.css" rel="stylesheet" type="text/css" />

<script src="/static/js/jquery.treetable.js"></script>
<h4>processor Exception Status </h4>


<table id="example-basic">
  
  <thead>
    <tr>
      <th>Time Stamp</th>
      <th>Process </th>
    </tr>
  </thead>
 
  <tbody>
     {{ (self.generate_header_rows  ) }}
  </tbody>
</table>
       '''

   def process_javascript(self):
       return_value = []
       self.mp.processor_id = self.processor_id
      
       return_value.append(self.load_processor_control_javascript())
       return_value.append( self.mp.macro_expand_start("{{","}}",self.process_javascript_raw()))
       
       return "\n".join(return_value)

   def process_javascript_raw(self):      
       return '''

<script>
$(document).ready(
 function()
 {
   
   
   $("#processor_select").val( {{ self.processor_id }});
   $("#processor_select").bind('change',change_processor)  
  



 $("#example-basic").treetable({
  expandable:true
});
 })
 

</script>
       '''





''' 
       
<tr data-tt-id="1">
      <td>Node 1: Click on the icon in front of me to expand this branch.</td>
      <td>I live in the second column.</td>
    </tr>
    <tr data-tt-id="1.1" data-tt-parent-id="1">
      <td>Node 1.1: Look, I am a table row <em>and</em> I am part of a tree!</td>
      <td>Interesting.</td>
    </tr>
    <tr data-tt-id="1.1.1" data-tt-parent-id="1.1">
      <td>Node 1.1.1: I am part of the tree too!</td>
      <td>That's it!</td>
    </tr>
    <tr data-tt-id="2">
      <td>Node 2: I am another root node, but without children</td>
      <td>Hurray!</td>
    </tr>
'''    