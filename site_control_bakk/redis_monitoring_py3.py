
from common_tools.redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
from common_tools.system_error_log_py3 import  System_Error_Logging


from common_tools.Pattern_tools_py3.builders.common_directors_py3 import construct_all_handlers
from common_tools.Pattern_tools_py3.factories.graph_search_py3 import common_qs_search
from common_tools.Pattern_tools_py3.factories.get_site_data_py3 import get_site_data



class Redis_Monitor(object):

   def __init__(self,site_data,qs ):
       self.redis_handle = qs.get_redis_data_handle()
       search_list = [ "REDIS_MONITORING"]
       self.redis_monitoring_streams = construct_all_handlers(site_data,qs,search_list,rpc_client=None)

   
       
       self.system_error_logging = System_Error_Logging(qs,"Node_Control",site_data)
  
       
       
       
       self.cpu_previous = None
      
       self.call_stat_previous = None
       
       
   def log_data(self,*parameters):
       
       redis_data = self.redis_handle.info("Keyspace")
       print("redis_data",redis_data)       
       self.redis_monitoring_streams["KEYS"].push(data=redis_data)      
       
       redis_data = self.redis_handle.info("Clients") 
       self.redis_monitoring_streams["CLIENTS"].push(data=redis_data)  
       
       redis_data = self.redis_handle.info("Memory")            
       self.redis_monitoring_streams["MEMORY"].push(data=redis_data)      

       redis_data = self.redis_handle.info("commandstats")
       temp_time = {}
       temp_calls = {}
       for k, item in redis_data.items():
          temp_time[ k] = item[ 'usec_per_call']
          temp_calls[k]  = item['calls']
          
       self.redis_monitoring_streams["REDIS_MONITOR_CMD_TIME_STREAM"].push(data=temp_time)     
       if self.call_stat_previous != None:
           delta_call = {}
           for k,item in temp_calls.items():
               if k in self.call_stat_previous:
                   delta_call[k] = item - self.call_stat_previous[k]
               else:
                 pass #  raise

           self.redis_monitoring_streams["REDIS_MONITOR_CALL_STREAM"].push(data=delta_call)  
       
       redis_data = self.redis_handle.info("cpu")
       
       if self.cpu_previous != None:
          temp = {}
          for key, item in redis_data.items():
              if key in self.cpu_previous:
                   temp[key] = float(item) - self.cpu_previous[key]
              else:
                pass  #raise

          self.redis_monitoring_streams["REDIS_MONITOR_SERVER_TIME"].push(data=temp)  
       else:
          temp = redis_data
          
       
       self.call_stat_previous = temp_calls
       self.cpu_previous = redis_data
       
 
       
def construct_redis_instance( qs, site_data ):

                   
    
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )

    query_list = qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", property_mask={"name":"REDIS_MONITORING"} )
                                           
    package_sets, package_sources = qs.match_list(query_list)  
 
    package = package_sources[0]
    
    #
    #  do verifications of data package
    #
    #
    #
    data_structures = package["data_structures"]
    print("data_structures",data_structures.keys())
    generate_handlers = Generate_Handlers( package, qs )
    
    redis_monitoring_streams = {}
    redis_monitoring_streams["KEYS"] = generate_handlers.construct_stream_writer(data_structures["KEYS"] )
    redis_monitoring_streams["CLIENTS"] = generate_handlers.construct_stream_writer(data_structures["CLIENTS"] )
    redis_monitoring_streams["MEMORY"] = generate_handlers.construct_stream_writer(data_structures["MEMORY"] )
    
    redis_monitoring_streams["REDIS_MONITOR_CALL_STREAM"] = generate_handlers.construct_stream_writer(data_structures["REDIS_MONITOR_CALL_STREAM"])
    redis_monitoring_streams["REDIS_MONITOR_CMD_TIME_STREAM"] = generate_handlers.construct_stream_writer(data_structures["REDIS_MONITOR_CMD_TIME_STREAM"])
    redis_monitoring_streams["REDIS_MONITOR_SERVER_TIME"] = generate_handlers.construct_stream_writer(data_structures["REDIS_MONITOR_SERVER_TIME"])
 
   
    
    redis_monitor = Redis_Monitor(qs.get_redis_data_handle() , redis_monitoring_streams )
    
    
    
  

    return redis_monitor




def add_chains(redis_monitor, cf):
 
    cf.define_chain("make_measurements", True)
    cf.insert.log("logging_redis_data")
    cf.insert.one_step(redis_monitor.log_data)
    cf.insert.wait_event_count( event = "MINUTE_TICK",count = 15)
    cf.insert.reset()

 

if __name__ == "__main__":

    import datetime
    import time
    import string
    import urllib.request
    import math
    import redis
    import base64
    import json

    import os
    import copy
    #import load_files_py3
    from common_tools.redis_support_py3.graph_query_support_py3 import  Query_Support
    import datetime
    #from redis_support_py3.user_data_tables_py3 import User_Data_Tables

    from common_tools.py_cf_new_py3.chain_flow_py3 import CF_Base_Interpreter

    site_data = get_site_data("/mnt/ssd/site_config/redis_server.json")
    qs = Query_Support( site_data )  
  
    
    redis_monitor = Redis_Monitor(site_data,qs )
    print("made it here 2")
    #
    # Adding chains
    #
    cf = CF_Base_Interpreter()
    add_chains(redis_monitor, cf)
    #
    # Executing chains
    #
    print("made it here 3")
    cf.execute()