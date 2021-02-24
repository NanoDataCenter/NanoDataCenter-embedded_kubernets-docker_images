
import sys
import os
from os import listdir
from os.path import isfile, join
import redis
import json
import msgpack
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
from system_error_log_py3 import  System_Error_Logging
from Pattern_tools_py3.builders.common_directors_py3 import construct_all_handlers
from sqlite_library.sqlite_sql_support_py3 import SQLITE_Client_Support
from file_server_library.file_server_lib_py3 import Construct_RPC_File_Library

#
#  DB_CONNECTIONS hash key for store data base file locations
#
#
#


class Construct_RPC_Server(object):

   def __init__(self,rpc_queue_handle ):
       self.rpc_queue_handle = rpc_queue_handle
       self.rpc_queue_handle.register_call_back( "load",self.load_file)
       self.rpc_queue_handle.register_call_back( "save", self.save_file)
       self.rpc_queue_handle.register_call_back( "file_exists",self.file_exists)
       self.rpc_queue_handle.register_call_back( "delete_file", self.delete_file)
       self.rpc_queue_handle.register_call_back( "file_directory", self.file_directory)
       self.rpc_queue_handle.register_call_back("make_dir",self.mkdir)
       self.rpc_queue_handle.add_time_out_function(self.process_null_msg) 
       self.rpc_queue_handle.start()


   def process_null_msg( self ):  
       print("null message")         
 
   
   def load_file(self,input_message):
       try:
           print("load_file")
           path = "/files/"+input_message["path"]+"/"+input_message["file_name"]
           f = open(path, 'r')
           data = f.read()
           f.close()
           return [True,data]
       except:
           return [False,data]
       
   def save_file(self,input_message):
       try:
           print("save_file")
           path = "/files/"+input_message["path"]+"/"+input_message["file_name"]
           f = open(path, 'w')
           data = input_message["data"]
           f.write(data)
           f.close()
           return [True,None]
       except:
           return [False,None]
   
   def file_exists(self,input_message):
       try:
           print("file_exits")
           path = "/files/"+input_message["path"]+"/"+input_message["file_name"]
           return [True,isfile(path)]
       except:
           return [False,None]
        
   def file_directory(self,input_message):
       try:
           print("file_directory")
           path = "/files/"+input_message["path"]
           return [True,listdir(path)]
       except:
           return [False,None]       

   def delete_file(self, input_message):
       try:
           print("delete_file")
           path = "/files/"+input_message["path"]+"/"+input_message["file_name"]
           os.remove(path)
           return [True,None]
       except:
           return [False,None]
           
           
   def mkdir(self,input_message):
       try:
           print("mkdir")
           path = "/files/"+input_message["path"]
           os.makedirs(path)
           return [True,None]
       except:
           return [False,None]
   
 
 
       
def construct_fileserver_instance( qs, site_data ):
          
    '''
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
    query_list = qs.add_match_relationship( query_list,relationship= "FILE_SERVER")
    query_list = qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", property_mask={"name":"FILE_SERVER"} )
    package_sets, package_sources = qs.match_list(query_list) 
    package = package_sources[0]    
    data_structures = package["data_structures"]
    generate_handlers = Generate_Handlers( package, qs )
    rpc_queue = generate_handlers.construct_rpc_sever(data_structures["FILE_SERVER_RPC_SERVER"] )
    Construct_RPC_Server(rpc_queue )

    '''
    search_list = ["FILE_SERVER","FILE_SERVER"]
    rpc_queue = construct_all_handlers(site_data,qs,search_list,field_list=["FILE_SERVER_RPC_SERVER"])
    field_list=None
 



 

if __name__ == "__main__":
 
    import datetime
    import time
    import string
    import urllib.request
    import math
    import redis
    
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
    container_name = os.getenv("CONTAINER_NAME")
    self.sqlite_client = SQLITE_Client_Support(qs,site_data)
    self.error_logging = System_Error_Logging(qs,container_name,site_data,self.sqlite_client)

    
    
    construct_fileserver_instance(qs, redis_site )
    
