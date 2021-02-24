



import redis
import time
import base64
from datetime import datetime
import sys
import json
from redis_support_py3.graph_query_support_py3 import  Query_Support
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers


#
#
#  we need this class because we are using a different redis server
#
#
#

import uuid
import msgpack
class Redis_RPC_Client(object):

   def __init__( self,redis_handle ):
       self.redis_handle = redis_handle
      
       
   
   def set_rpc_queue(self,queue):
       self.rpc_queue = queue
       
       
   def send_rpc_message( self, method,parameters,timeout=30 ):
        request = {}
        request["method"] = method
        request["params"] = parameters
        request["id"]   = str(uuid.uuid1())    
        request_msg = msgpack.packb( request )
        self.redis_handle.delete(request["id"] )
        self.redis_handle.lpush(self.rpc_queue, request_msg)
        data =  self.redis_handle.brpop(request["id"],timeout = timeout )
        
        self.redis_handle.delete(request["id"] )
        if data == None:
            raise ValueError("No Communication with RPC SERVER")
        response = msgpack.unpackb(data[1])
        
        return response 
        
        
class Modbus_Server( object ):
    
   def __init__( self, generate_handlers,data_structures,remote_ip,remote_queue):  # fill in proceedures
       self.redis_handle = redis.StrictRedis(remote_ip, 
                                           6379, 
                                           db=4)
                                           

       self.redis_rpc_client = Redis_RPC_Client(self.redis_handle)
       self.redis_rpc_client.set_rpc_queue(remote_queue)
       self.rpc_server_handle = generate_handlers.construct_rpc_sever(data_structures["PLC_RPC_SERVER"] )

       
      
       self.rpc_server_handle.register_call_back( "modbus_relay", self.process_modbus_message)
       self.rpc_server_handle.register_call_back( "ping_message", self.process_ping_message)
       self.rpc_server_handle.add_time_out_function(self.process_null_msg)
       self.rpc_server_handle.start()
 
 
   def process_ping_message(self, address):    
        try:
            output_message = self.redis_rpc_client.send_rpc_message("ping_message",address,timeout=30 )
            print("ping*******************",address,output_message)
            return output_message      
        except:
            print("ping failure +++++++++++++++++",address)        
        
   def process_modbus_message( self,input_message ):
       try:
           output_message = self.redis_rpc_client.send_rpc_message("modbus_relay",input_message,timeout=30 )
           print("process message***************",input_message,output_message)
           return output_message
       except:
           raise
           print("process failure ++++++++++++++++",input_message)
       
        

   def process_null_msg( self ):
       print("time out message")






if __name__ == "__main__":
   plc_server_name =  sys.argv[1]
   remote_ip       = sys.argv[2]
   remote_queue    = sys.argv[3]
   file_handle = open("/data/redis_server.json",'r')
   data = file_handle.read()
   file_handle.close()
   redis_site = json.loads(data)
   qs = Query_Support( redis_site )
   
   #  find data structures
   query_list = []   
   query_list = qs.add_match_relationship( query_list,relationship="SITE",label=redis_site["site"] )
   query_list = qs.add_match_relationship( query_list,relationship="PLC_SERVER",label=plc_server_name )
   query_list = qs.add_match_terminal( query_list, 
                                           relationship = "PACKAGE", 
                                           property_mask={"name":"PLC_SERVER_DATA"} )
                                           
   package_sets, package_sources = qs.match_list(query_list)
       
   package = package_sources[0] 
   generate_handlers = Generate_Handlers(package,qs)   
   data_structures = package["data_structures"]
   
   
   
   
   Modbus_Server( generate_handlers,data_structures,remote_ip,remote_queue )    
 
   