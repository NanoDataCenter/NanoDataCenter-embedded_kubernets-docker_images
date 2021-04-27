

from templates.common.redis_stream_manager_py3 import Redis_Stream_Manager
from .redis_base_self_py3 import Redis_Stream_Base
class Redis_Server_Time_Stream(Redis_Stream_Base,Redis_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Stream_Manager.__init__(self,base_self,parameters)
       Redis_Stream_Base.__init__(self,base_self)

   def application_generation(self):
     temp_data = self.handlers["REDIS_MONITOR_SERVER_TIME"].revrange("+","-" , count=1000)
     temp_data.reverse()

     chart_title = " Redis Execution time/hour : "
     self.stream_keys,self.stream_range,self.stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="Server Time",title_x="Date")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0
    
 
        
 
      
 
 
