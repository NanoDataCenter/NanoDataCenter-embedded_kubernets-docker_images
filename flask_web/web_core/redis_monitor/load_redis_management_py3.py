
import os
import json
from datetime import datetime
import time
from base_stream_processing.base_stream_processing_py3 import Base_Stream_Processing
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
class Load_Redis_Monitoring(Base_Stream_Processing):

   def __init__(self, app, auth, request, render_template,qs,site_data,url_rule_class,subsystem_name,path):
       self.app      = app
       self.auth     = auth
       self.request  = request
       self.render_template = render_template
       self.path = path
       self.qs = qs
       self.site_data = site_data
       self.url_rule_class = url_rule_class
       self.subsystem_name = subsystem_name
       
       self.assemble_handlers()
       self.assemble_url_rules()
       self.path_dest = self.url_rule_class.move_directories(self.path)     

       
   def assemble_url_rules(self):
    
       self.slash_name = "/"+self.subsystem_name
       self.menu_data = {}
       self.menu_list = []
       
       

       
       function_list =   [ self.redis_key_stream,
                           self.redis_client_stream,
                           self.redis_memory_stream,
                           self.redis_call_stream,                      
                           self.redis_cmd_time_stream,
                           self.redis_server_time_stream ]
                           
 
                               
   
      
       url_list = [
                      [ 'redis_key_stream' ,'','',"Redis Key Stream"  ],
                      [ 'redis_client_stream' ,'','',"Redis Client Stream"  ],
                      [ 'redis_memory_stream' ,'','',"Redis Memory Stream"   ],
                      [ 'redis_call_stream' ,'','',"Redis Call Stream"  ],    
                      [ 'redis_cmd_time_stream','','',"Redis Command Time Stream"  ],
                      [ 'redis_server_time_stream','','',"Redis Server Time Stream"  ] 
        ]

       self.url_rule_class.add_get_rules(self.subsystem_name,function_list,url_list)
            

   def assemble_handlers(self):
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
       query_list = self.qs.add_match_relationship( query_list,relationship="CONTAINER",label="monitor_redis" )
       
       
       query_list = self.qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", label = "REDIS_MONITORING" )
                                           
       package_sets, package_sources = self.qs.match_list(query_list)  
      
       package = package_sources[0]
       generate_handlers = Generate_Handlers(package,self.qs)
       data_structures = package["data_structures"]     
 
       self.handlers = {}
       self.handlers["REDIS_MONITOR_KEY_STREAM"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_KEY_STREAM"])
       self.handlers["REDIS_MONITOR_CLIENT_STREAM"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_CLIENT_STREAM"])
       self.handlers["REDIS_MONITOR_MEMORY_STREAM"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_MEMORY_STREAM"])
       self.handlers["REDIS_MONITOR_CALL_STREAM"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_CALL_STREAM"])
       self.handlers["REDIS_MONITOR_CMD_TIME_STREAM"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_CMD_TIME_STREAM"])
       self.handlers["REDIS_MONITOR_SERVER_TIME"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_SERVER_TIME"])
       


   def redis_key_stream(self):
       
       temp_data = self.handlers["REDIS_MONITOR_KEY_STREAM"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " Number of Redis Key in : "
      
       stream_keys,stream_range,stream_data = self.format_data_specific_key(temp_data,title=chart_title,title_y="Deg F",title_x="Date",specific_key = "keys")
       
       
       return self.render_template( "streams/base_stream",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     max_value = 10000000,
                                     min_value = 0
                                     
                                     )

      



  

   def redis_client_stream(self):
       temp_data = self.handlers["REDIS_MONITOR_CLIENT_STREAM"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " Redis Client : "
       
       stream_keys,stream_range,stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       
      
       return self.render_template( "streams/base_stream",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     max_value = 10000000,
                                     min_value = 0
                                     
                                     )



   def redis_memory_stream(self):
       temp_data = self.handlers["REDIS_MONITOR_MEMORY_STREAM"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " Redis Memory Utilization : "
       
       stream_keys,stream_range,stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       
      
       return self.render_template( "streams/base_stream",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     max_value = 10000000,
                                     min_value = 0

                                     
                                     )




   def redis_call_stream(self):
       temp_data = self.handlers["REDIS_MONITOR_CALL_STREAM"].revrange("+","-" , count=1000)
       #print(temp_data)
       temp_data.reverse()
       chart_title = " Number of Redis Command Calls/hour : "
       
       stream_keys,stream_range,stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       
       #print("made it here",stream_keys,stream_data)
       return self.render_template( "streams/base_stream",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     max_value = 10000000,
                                     min_value = 0
                                     
                                     )

        
 
   def redis_cmd_time_stream (self):
       temp_data = self.handlers["REDIS_MONITOR_CMD_TIME_STREAM"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " Redis Command Time in us : "
       
       stream_keys,stream_range,stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       
      
       return self.render_template( "streams/base_stream",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     max_value = 10000000,
                                     min_value = 0
                                     
                                     )

    
        
   def redis_server_time_stream(self):
       temp_data = self.handlers["REDIS_MONITOR_SERVER_TIME"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " Redis Execution time/hour : "
       
       stream_keys,stream_range,stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       
      
       return self.render_template( "streams/base_stream",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     max_value = 10000000,
                                     min_value = 0
                                     
                                     )
       
 
      
 
 
