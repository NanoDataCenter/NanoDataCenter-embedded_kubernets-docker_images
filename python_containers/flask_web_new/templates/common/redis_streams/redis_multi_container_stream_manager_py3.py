
import json
from templates.Base_Multi_Template_Class_py3  import Base_Multi_Template_Class
from .redis_stream_base_py3 import Redis_Stream_Base
class Redis_Multi_Container_Stream_Manager( Redis_Stream_Base,Base_Multi_Template_Class):
   def __init__(self,base_self,parameters):
      Base_Multi_Template_Class.__init__(self,base_self,parameters)
      Redis_Stream_Base.__init__(self)
      
   



       
   def generate_data(self):
        self.stream_data = []
        self.stream_keys = []
        self.titles  = []
        self.stream_range  = 0
        self.max_value    = []
        self.min_value    = []        

   def application_page_generation(self,container, not_used_data):
       if container >= len(self.containers):
           container = len(self.containers)-1
       self.application_generation(container,not_used_data)
       return_value = []
       return_value.append('<script src="/static/js/plotly-latest.min.js"></script>') 
       return_value.append('<script type="text/javascript">')
       return_value.append('container_id = '+str(container))
       return_value.append('function change_container(event,ui)')
       return_value.append('{')
       return_value.append('    current_page = window.location.pathname')
 
       return_value.append('    current_page = current_page+"?"+$("#container_select")[0].selectedIndex')
       return_value.append('   window.location.href = current_page')
       return_value.append('}')
       return_value.append('   $(function () ')
       return_value.append('{')
       return_value.append('$("#container_select").val( container_id );')
       return_value.append('$("#container_select").bind("change",change_container)')   
       return_value.append('}')
       return_value.append(')')
       return_value.append("</script>")
       return_value.append('<div style="margin-top:20px"></div>')
       return_value.append('<div id="container_tag">')
       return_value.append('<center>')
       return_value.append('<h4>Select Container</h4>')
       return_value.append('</center>')
       return_value.append('</div>')
       return_value.append('<div id="select_tag">')
       return_value.append('<center>')
       return_value.append('<select id="container_select">')
       
       for i in range(0,len(self.containers)):
          return_value.append('<option value="'+str(i)+'">'+self.containers[i]+'</option>')
       return_value.append('</select>')
       return_value.append('</center>')
       return_value.append('</div>')
       for i in self.stream_range:
           title_name = self.title[i].replace(".","_").replace("-","_")
          
           data = self.stream_data[self.stream_keys[i]]
        
           layout = data["layout"]
           trace= data["data"]
           return_value.append('<div style="margin-top:50px"></div>')
           return_value.append('<div class="container">')
           return_value.append('<div id="'+title_name+'" style="width:100%;height:600pts;"></div>')
           return_value.append('<script>')
           
           
           
           return_value.append('layout_json = '+"'"+json.dumps(layout)+"'")
           return_value.append('layout = JSON.parse(layout_json)')
           return_value.append('title_name= "'+title_name+'"')
           return_value.append('trace_json = '+"'"+json.dumps(trace)+"'")
           return_value.append('trace = JSON.parse(trace_json)')
           return_value.append('for( i = 0;i< trace.x.length; i++)')
           return_value.append('{')
           return_value.append('   trace.x[i] = new Date(trace.x[i]*1000) ')
           return_value.append('}')           
           return_value.append('Plotly.newPlot(title_name, [trace],layout);')
           return_value.append('</script>')
           return_value.append('</div>')
       return "\n".join(return_value)
   



    
