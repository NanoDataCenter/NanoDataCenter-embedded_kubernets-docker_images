 
 
from templates.common.redis_stream_manager_py3 import Redis_Stream_Manager
from .redis_base_self_py3 import Redis_Stream_Base
class Redis_Memory_Stream(Redis_Stream_Base,Redis_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Stream_Manager.__init__(self,base_self,parameters)
       Redis_Stream_Base.__init__(self,base_self)
       
       
       
   def application_generation(self):
     temp_data = self.handlers["REDIS_MONITOR_MEMORY_STREAM"].revrange("+","-" , count=1000)
     temp_data.reverse()

     chart_title = " Redis Memory Utilization : "
     self.stream_keys,self.stream_range,self.stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="Memory",title_x="Date")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0
    
 
 
   def redis_memory_stream(self):
       temp_data = self.handlers["REDIS_MONITOR_MEMORY_STREAM"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " Redis Memory Utilization : "
       
       stream_keys,stream_range,stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="Memory Stream",title_x="Date")
       
      
       return self.render_template( "streams/base_stream",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     max_value = 10000000,
                                     min_value = 0

                                     
                                     )
