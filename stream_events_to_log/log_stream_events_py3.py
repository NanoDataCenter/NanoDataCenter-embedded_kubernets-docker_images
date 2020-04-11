
import msgpack
import redis
import base64
from  utilities.web3_top_class import Web_Class_IPC
from  utilities.event_listener_top_class import Event_Listner_Class_IPC

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
           
   def list_range(self,start=0,stop=-1):
      
      list_data =  self.redis_handle.lrange(self.key,start,stop)
     
      if list_data == None:
         return None
      return_value = []
      for pack_data in list_data:
        return_value.append(msgpack.unpackb(pack_data))
      return return_value                

 
   def pop(self):
       pack_data = self.redis_handle.rpop(self.key)
        
  

       if pack_data == None:
          return False, None
       else:
         
          return True,msgpack.unpackb(pack_data)
          
   def show_next_job(self):
       pack_data = self.redis_handle.lindex(self.key, -1)
       if pack_data == None:
          return False, None
       else:
          
          return True, msgpack.unpackb(pack_data)

   def push_front(self,data):
       pack_data =  msgpack.packb(data )
       self.redis_handle.rpush(self.key,pack_data)
       self.redis_handle.ltrim(self.key,0,self.depth)
  








class Stream_Event_Monitor(object):

   def __init__(self,redis_handle,key,w3,el ):
       self.stream_server = Local_Queue_Server(redis_handle,key)
       self.w3  = w3
       self.el = el
       self.el_filter = self.el.construct_loop_filter("EventHandler" )
       self.event_object = self.w3.get_contract("EventHandler")
       #print(self.el.get_all_entries("EventHandler"))
       #self.decode_data(self.el.get_all_entries("EventHandler"))
       #exit()
       
   def log_data(self,*parameters):
       print(self.stream_server.length())
       try:       
          if self.stream_server.length() != 0:
              data = self.stream_server.show_next_job()
              self.log_to_block_chain(data)
       except:
          raise
          print("exception")           
 
   def log_to_block_chain(self,data):
       print("data",data)
       data = data[1]
       print("data",data)
       data =data[1]
       print("data",data)
       site = data["site"]
       name = data['name']
       pack_data = msgpack.packb(data)  #json.dumps(data)
       pack_data = base64.b64encode(pack_data).decode()
       parameters = [name,site,pack_data]
       #transmit_event( string memory event_id, string memory sub_event, string memory data)   
       tx_reciept = self.w3.transact_contract_data(self.event_object, "transmit_event" ,parameters)
       print(tx_reciept)
       raise
       
   def decode_data(self,data):  # make sure data can be decode from block chain
       for  i in data:
           #print("i",type(i),i)
           #print(i['args'])
           compress = i['args']["data"]
           print("compress",compress)
           try:
               decode = base64.b64decode(compress.encode())
               print("decode",type(decode),decode)
               print("msg pack ", msgpack.unpackb(decode))
           except:
              print("bad data",i['args']["data"])

       
def construct_stream_server_instance( qs, site_data,w3,el ):
          
    
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
    query_list = qs.add_match_relationship( query_list,relationship= "CLOUD_SERVICE_HOST_INTERFACE" )
    query_list = qs.add_match_terminal( query_list,relationship = "HOST_INFORMATION" )
    node_sets,node_list = qs.match_list(query_list)                                       
    node_data = node_list[0]
    

    redis_handle = redis.StrictRedis( host = node_data["host"] , port=node_data["port"], db=node_data['key_data_base'] )
    stream_event_monitor = Stream_Event_Monitor(redis_handle, node_data["key"],w3,el )
    
    
    
  

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
   
    #
    # Setup handle
    # open data stores instance
   
    qs = Query_Support( redis_site )
    
    
    
    redis_contract_handle = redis.StrictRedis( host = redis_site["host"] , port=redis_site["port"], db=redis_site["redis_contract_db"] )
    ipc_socket = "/ipc/geth.ipc"
    w3 = Web_Class_IPC(ipc_socket,redis_contract_handle)
    el = Event_Listner_Class_IPC(ipc_socket,redis_contract_handle)
    stream_event_monitor = construct_stream_server_instance(qs, redis_site,w3,el )
    
    
    #
    # Adding chains
    #
    cf = CF_Base_Interpreter()
    add_chains(stream_event_monitor, cf)
    #
    # Executing chains
    #
    
    cf.execute()