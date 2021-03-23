

























from .docker_container_base_py3 import Docker_Base_Class
from templates.Base_Multi_Template_Class_py3  import Base_Multi_Template_Class
from templates.common.redis_streams.redis_multi_container_stream_manager_py3 import Redis_Multi_Container_Stream_Manager
from flask import request
import json
import datetime









class Container_Cpu_Loading(Redis_Multi_Container_Stream_Manager,Docker_Base_Class):
   def __init__(self,base_self,parameters = None):
       Docker_Base_Class.__init__(self,base_self)
       Redis_Multi_Container_Stream_Manager.__init__(self,base_self,parameters)

   def application_generation(self,container_id,data):
       container_name = self.managed_container_names[container_id]
       temp_data =self.handlers[container_name]["PROCESS_CPU"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " %CPU Utilization "
       self.stream_keys,self.stream_range,self.stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="%CPU Utilization",title_x="Date")
       self.title = self.stream_keys


class Container_Vsz_Loading(Redis_Multi_Container_Stream_Manager,Docker_Base_Class):
   def __init__(self,base_self,parameters = None):
       Docker_Base_Class.__init__(self,base_self)
       Redis_Multi_Container_Stream_Manager.__init__(self,base_self,parameters)

   def application_generation(self,container_id,data):
       container_name = self.managed_container_names[container_id]
       temp_data =self.handlers[container_name]["PROCESS_VSZ"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " %VSZ Useage "
       self.stream_keys,self.stream_range,self.stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="%CPU Utilization",title_x="Date")
       self.title = self.stream_keys
       
       
class Container_Rss_Loading(Redis_Multi_Container_Stream_Manager,Docker_Base_Class):
   def __init__(self,base_self,parameters = None):
       Docker_Base_Class.__init__(self,base_self)
       Redis_Multi_Container_Stream_Manager.__init__(self,base_self,parameters)

   def application_generation(self,container_id,data):
       container_name = self.managed_container_names[container_id]
       temp_data =self.handlers[container_name]["PROCESS_RSS"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " %RSS Useage "
       self.stream_keys,self.stream_range,self.stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="%CPU Utilization",title_x="Date")
       self.title = self.stream_keys
