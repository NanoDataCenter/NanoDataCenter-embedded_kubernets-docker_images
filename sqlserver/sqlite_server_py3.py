import msgpack
import redis
import base64

from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers

#
#  DB_CONNECTIONS hash key for store data base file locations
#
#
#


class Construct_RPC_Server(object):

   def __init__(self,rpc_queue_handle,sql_databases  ):
       # create handles for existing databases      
       self.sql_databases = sql_databases
       self.rpc_queue_handle = rpc_queue_handle
       self.initialize_data_base_handles()       
       self.rpc_queue_handle.register_call_back( "list_list_data_bases",self.list_data_bases)
       self.rpc_queue_handle.register_call_back( "create_database", self.create_database)
       self.rpc_queue_handle.register_call_back("close_database",self.close_database)
       self.rpc_queue_handle.register_call_back("vacuum",self.vacuum)
       self.rpc_queue_handle.register_call_back("create_table",self.vacuum)
       self.rpc_queue_handle.register_call_back("create_text_search_table",self.vacuum)
       self.rpc_queue_handle.register_call_back("exec",self.vacuum)
       self.rpc_queue_handle.register_call_back("select",self.vacuum)
       self.rpc_queue_handle.register_call_back("bluk_upload",self.bulk_upload)
       self.rpc_queue_handle.register_call_back("copy",self.copy)
       self.rpc_queue_handle.register_call_back("statistics",self.statistics)
       self.rpc_queue_handle.add_time_out_function(self.process_null_msg)
       self.rpc_queue_handle.start()
      
 
   def initialize_data_base_handles(self):
       self.db_handlers = {}
       if "default" not in self.sql_databases:
           self.sql_databases["default"] = "/sqlite/default.db"
       if "memory" not in self.sql_databases:
             self.sql_database["memory"] = ":memory:"
       
       for i in self.sql_databases.hkeys():
           try:
               self.db_handlers[j] = sqlite.connect(self.sql_databases.hget(i))      
           except:
               print("db connection problem ",i,self.sql_databases.hget(i))           
               self.sql_databases.hdel(i)
               
                      
 
 
   def process_null_msg( self ):  
       print("null message")   
  
   def list_data_bases(self,input_message):
       return_value = self.sql_databases.hgetall()
       return return_value       
       
   def create_database(self,input_msg):
       try:
           name = input_msg["database_name"]
           file_name = input_msg["file_name"]
           file_path = "/sqlite/"+file_name
           connection = sqlite.connect(file_path)
           self.sql_databases.hset(name,file_path)
           self.db_handlers[name] = connection
           return True
       except:
           return False
           
 
   def close_database(self,input_msg):
       try:
           name = input_msg["database_name"]
           connection = self.db_handlers[name]
           connection.close(connection)
           del self.db_handlers[name]
           self.sql_databases.hdel(name)
           return True
       except:
           return True

   def vacuum(self,input_msg):
       try:
           db_name = input_msg["database_name"]
           connection = self.db_handlers[db_name]
           connection.exec("VACUUM")
           return True
       except:
           return False
       
   def create_table(self,input_msg):
      pass
   
   
   def create_text_search_table(self,input_msg):
      pass
   
   
   
   def ex_exec(self,input_msg):
       pass
   
   def select(self,input_msg):
      pass
   

   def bulk_upload(self,input_msg):
       pass
   
   
   
   def copy(self,input_msg):
      pass
   
   
   def statistics(self,input_message):
      pass
 
       
def construct_sql_server_instance( qs, site_data ):
          
    
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
    query_list = qs.add_match_relationship( query_list,relationship= "SQL_SERVER" )
    query_list = qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", property_mask={"name":"SQL_SERVER"} )
    package_sets, package_sources = qs.match_list(query_list) 
    package = package_sources[0]    
    data_structures = package["data_structures"]
    generate_handlers = Generate_Handlers( package, qs )
    rpc_queue = generate_handlers.construct_rpc_sever(data_structures["SQL_SERVER_RPC_SERVER"] )
    sql_databases = generate_handlers.construct_hash(data_structures["SQL_DB_MAPPING"] )
    Construct_RPC_Server(rpc_queue,sql_databases )






 

if __name__ == "__main__":
    import sqlite3
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
    import msgpack

  

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
    
    
    
    construct_sql_server_instance(qs, redis_site )
    
