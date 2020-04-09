
import msgpack


class Local_Queue_Server( object ):
 
   def __init__(self,redis_handle, key):
      
      self.redis_handle = redis_handle
      self.key = key 
 
 
 
 
   def length(self):
       return self.redis_handle.llen(self.key)
       
   def delete(self, index ):
       if index < self.redis_handle.llen(self.key):
           self.redis_handle.lset(self.key, index,"__#####__")
           self.redis_handle.lrem(self.key, 1,"__#####__") 
           
                

 
   def pop(self):
       pack_data = self.redis_handle.rpop(self.key)
        
  

       if pack_data == None:
          return False, None
       else:
         
          return True,msgpack.unpackb(pack_data,encoding='utf-8')
          
   def show_next_job(self):
       pack_data = self.redis_handle.lindex(self.key, -1)
       if pack_data == None:
          return False, None
       else:
          
          return True, msgpack.unpackb(pack_data,encoding='utf-8')

   def push_front(self,data):
       pack_data =  msgpack.packb(data,use_bin_type = True )
       self.redis_handle.rpush(self.key,pack_data)
       self.redis_handle.ltrim(self.key,0,self.depth)
  









class Stream_Event_Monitor(object):

   def __init__(self,redis_handle,cloud_log_stream ):
       self.stream_server = Local_Queue_Server(redis_handle,cloud_log_stream["key"])
       print(self.stream_server.length())
       
   def log_data(self,*parameters):
       pass      
 
       
 
       
def construct_stream_server_instance( qs, site_data ):
          
    
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
    query_list = qs.add_match_relationship( query_list,relationship= "CLOUD_SERVICE_HOST_INTERFACE" )
    query_list = qs.add_match_terminal( query_list,relationship = "HOST_INFORMATION" )
    node_sets,node_list = qs.match_list(query_list)                                       
    node_data = node_list[0]
    
    stream_event_monitor = Stream_Event_Monitor(qs.get_redis_data_handle() , node_data )
    exit()
    
    
  

    return stream_event_monitor




def add_chains(stream_event_monitor, cf):
 
    cf.define_chain("make_measurements", True)
    cf.insert.log("log_cloud_events")
    cf.insert.one_step(stream_event_monitor.log_data)
    cf.insert.wait_event_count( event = "TIME_TICK",count = 1)
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
    from redis_support_py3.graph_query_support_py3 import  Query_Support
    import datetime
    

    from py_cf_new_py3.chain_flow_py3 import CF_Base_Interpreter

    #
    #
    # Read Boot File
    # expand json file
    # 
    file_handle = open("/data/redis_server.json",'r')
    data = file_handle.read()
    file_handle.close()
    redis_site = json.loads(data)
    print("made it here 1")
    #
    # Setup handle
    # open data stores instance
   
    qs = Query_Support( redis_site )
    
    stream_event_monitor = construct_stream_server_instance(qs, redis_site )
    print("made it here 2")
    #
    # Adding chains
    #
    cf = CF_Base_Interpreter()
    add_chains(stream_event_monitor, cf)
    #
    # Executing chains
    #
    print("made it here 3")
    cf.execute()