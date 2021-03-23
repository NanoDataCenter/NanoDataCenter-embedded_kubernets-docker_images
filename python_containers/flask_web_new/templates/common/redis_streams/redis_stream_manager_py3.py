
import json
from templates.Base_Template_Class_py3  import Base_Template_Class
from .redis_stream_base_py3 import Redis_Stream_Base
class Redis_Stream_Manager( Redis_Stream_Base,Base_Template_Class):
   def __init__(self,base_self,parameters):
      Base_Template_Class.__init__(self,base_self,parameters)
      Redis_Stream_Base.__init__(self)
      
      
   


 
   def generate_data(self):
        self.stream_data = []
        self.stream_keys = []
        self.titles  = []
        self.stream_range  = 0
        self.max_value    = []
        self.min_value    = []        

      
   def application_page_generation(self, application_construction):
       self.application_generation()
       return_value = []
       return_value.append('<script src="/static/js/plotly-latest.min.js"></script>') 
       for i in self.stream_range:
         
           return_value.append('<div style="margin-top:50px"></div>')
           return_value.append('<div class="container">')
           title_name = self.title[i]
           
           title_name = title_name.replace(".","_").replace("-","_")
           return_value.append('<div id="'+title_name+'" style="width:100%;height:600pt;"></div>')
           data = self.stream_data[self.title[i]]
           trace= data["data"]
           layout = data["layout"]
           for i  in range(0, len(trace["x"])):
               try:
                   if trace["y"][i] > self.max_value:
                       trace["y"][i] = self.max_value
                   if trace["y"][i] < self.min_value:
                      trace["y"][i] = self.min_value
               except:
                  pass
           
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
    
