
from templates.common.redis_streams.redis_stream_manager_py3 import Redis_Stream_Manager
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers

class Redis_Stream_Base(object):
    def __init__(self,base_self):
       qs = base_self.qs
       query_list = []
       query_list = qs.add_match_relationship( query_list,relationship="SITE",label=base_self.site_data["site"] )
       query_list = qs.add_match_relationship( query_list,relationship="CONTAINER",label="monitor_redis" )
       
       
       query_list = qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", label = "REDIS_MONITORING" )
                                           
       package_sets, package_sources = qs.match_list(query_list)  
      
       package = package_sources[0]
       generate_handlers = Generate_Handlers(package,qs)
       data_structures = package["data_structures"]     
 
       self.handlers = {}
       self.handlers["REDIS_MONITOR_KEY_STREAM"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_KEY_STREAM"])
       self.handlers["REDIS_MONITOR_CLIENT_STREAM"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_CLIENT_STREAM"])
       self.handlers["REDIS_MONITOR_MEMORY_STREAM"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_MEMORY_STREAM"])
       self.handlers["REDIS_MONITOR_CALL_STREAM"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_CALL_STREAM"])
       self.handlers["REDIS_MONITOR_CMD_TIME_STREAM"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_CMD_TIME_STREAM"])
       self.handlers["REDIS_MONITOR_SERVER_TIME"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_SERVER_TIME"])
       

class Redis_Call_Stream(Redis_Stream_Base,Redis_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Stream_Manager.__init__(self,base_self,parameters)
       Redis_Stream_Base.__init__(self,base_self)


   def application_generation(self):
     temp_data = self.handlers["REDIS_MONITOR_CALL_STREAM"].revrange("+","-" , count=1000)
     temp_data.reverse()

     chart_title = " Number of Redis Command Calls/hour : "
     self.stream_keys,self.stream_range,self.stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="Call Number",title_x="Date")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0
    


 
class Redis_Client_Stream(Redis_Stream_Base,Redis_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Stream_Manager.__init__(self,base_self,parameters)
       Redis_Stream_Base.__init__(self,base_self)
	   

   def application_generation(self):
     temp_data = self.handlers["REDIS_MONITOR_CLIENT_STREAM"].revrange("+","-" , count=1000)
     temp_data.reverse()

     chart_title =  " Redis Client : "
     self.stream_keys,self.stream_range,self.stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="Client Number",title_x="Date")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0
    



class Redis_Command_Time_Stream(Redis_Stream_Base,Redis_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Stream_Manager.__init__(self,base_self,parameters)
       Redis_Stream_Base.__init__(self,base_self)

   def application_generation(self):
     temp_data = self.handlers["REDIS_MONITOR_CMD_TIME_STREAM"].revrange("+","-" , count=1000)
     temp_data.reverse()

     chart_title =" Redis Command Time in us : "
     self.stream_keys,self.stream_range,self.stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="Command Time",title_x="Date")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0
    





class Redis_Key_Stream(Redis_Stream_Base,Redis_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Stream_Manager.__init__(self,base_self,parameters)
       Redis_Stream_Base.__init__(self,base_self)
      
       
       
   def application_generation(self):
     temp_data = self.handlers["REDIS_MONITOR_KEY_STREAM"].revrange("+","-" , count=1000)
     temp_data.reverse()

     chart_title = " Number of Redis Key in : "
     self.stream_keys,self.stream_range,self.stream_data = self.format_data_specific_key(temp_data,title=chart_title,title_y="Key Number",title_x="Date",specific_key = "keys")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0
    





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
    
 
        
 
      
 
 
