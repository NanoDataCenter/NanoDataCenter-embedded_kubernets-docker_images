import msgpack
import redis
import base64
from  utilities.web3_top_class import Web_Class_IPC
from  utilities.event_listener_top_class import Event_Listner_Class_IPC
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers




class Construct_RPC_Server(object):

   def __init__(self,rpc_queue_handle,el  ):
       self.rpc_queue_handle = rpc_queue_handle 
       self.el = el
       self.rpc_queue_handle.register_call_back( "fetch_block_chain_data", self.fetch_block_chain_data)
       self.rpc_queue_handle.register_call_back( "fetch_current_block_number", self.fetch_current_block_number)
       self.rpc_queue_handle.add_time_out_function(self.process_null_msg)
       self.rpc_queue_handle.start()
      
 
   def process_null_msg( self ):  
       print("null message")   
       
       
   def fetch_current_block_number(self,input_msg):
       return self.el.get_block_number()

   
   def fetch_block_chain_data(self,input_msg ):
       start_block = input_msg["start_block"]
       end_block = input_msg["end_block"]
       contract_name = input_msg["contract_name"]
       encoded_data = self.el.get_contract_logs(contract_name,start_block,end_block)
       decoded_data = self.decode_data(encoded_data)
       return decoded_data
       
   def decode_data(self,data):  # make sure data can be decode from block chain
       return_value = []
       for  i in data:
           #print("i",type(i),i)
           #print(i['args'])
           compress = i['args']["data"]
           #print("compress",compress)
           try:
               decode = base64.b64decode(compress.encode())
               temp = msgpack.unpackb(decode)
               temp["blockNumber"] = i["blockNumber"]
               return_value.append(temp)
      
           except:
              print("bad data",i['args']["data"])
              
       return return_value
       
def construct_block_chain_server_instance( qs, site_data,el ):
          
    
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
    query_list = qs.add_match_relationship( query_list,relationship= "CLOUD_BLOCK_CHAIN_SERVER" )
    query_list = qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", property_mask={"name":"CLOUD_BLOCK_CHAIN_SERVER"} )
    package_sets, package_sources = qs.match_list(query_list) 
    package = package_sources[0]    
    data_structures = package["data_structures"]
    generate_handlers = Generate_Handlers( package, qs )
    rpc_queue = generate_handlers.construct_rpc_sever(data_structures["BLOCK_CHAIN_RPC_SERVER"] )
    Construct_RPC_Server(rpc_queue,el )






 

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
 
    from redis_support_py3.graph_query_support_py3 import  Query_Support
    import datetime
    

  

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
    
    el = Event_Listner_Class_IPC(ipc_socket,redis_contract_handle)
    construct_block_chain_server_instance(qs, redis_site,el )
    
    
