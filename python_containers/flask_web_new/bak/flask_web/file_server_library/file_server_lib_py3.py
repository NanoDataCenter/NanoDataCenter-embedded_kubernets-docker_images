from redis_support_py3.graph_query_support_py3 import  Query_Support
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
import datetime
import msgpack

  

class Construct_RPC_Library(object):

   def __init__( self, qs, site_data ):
          
       self.site_data = site_data
       self.qs        = qs
       query_list = []
       query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
       query_list = qs.add_match_relationship( query_list,relationship= "FILE_SERVER" )
       query_list = qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", property_mask={"name":"FILE_SERVER"} )
       package_sets, package_sources = qs.match_list(query_list) 
       package = package_sources[0]    
       data_structures = package["data_structures"]
       queue_name = data_structures["FILE_SERVER_RPC_SERVER"]['queue']
       generate_handlers = Generate_Handlers( package, qs )
       self.rpc_client = generate_handlers.construct_rpc_client( )
       self.rpc_client.set_rpc_queue(queue_name)
    
   def load_file(self,path,file_name):
       #print("load_file")
       parameters = {}
       parameters["path"] = path
       parameters["file_name"] = file_name
       return_value = self.rpc_client.send_rpc_message( method="load",parameters=parameters,timeout=3 )
      
       if return_value[0] != True:
          raise ValueError("load file failure")
       return return_value[1]  
       
   def save_file(self,path,file_name,data):
       #print("save_file")
       
       parameters = {}
       parameters["path"] = path
       parameters["file_name"] = file_name
       parameters["data"] = data
       return_value = self.rpc_client.send_rpc_message( method="save",parameters=parameters,timeout=3 )
       if return_value[0] != True:
          raise ValueError("load file failure")
       return return_value[1]  
   
   def file_exists(self,path,file_name):
       #print("file_exits")
       parameters = {}
       parameters["path"] = path
       parameters["file_name"] = file_name
       return_value = self.rpc_client.send_rpc_message( method="file_exists",parameters=parameters,timeout=3 )
       if return_value[0] != True:
          raise ValueError("load file failure")
       return return_value[1]    
        
   def file_directory(self,path):
       #print("file_directory")
       parameters = {}
       parameters["path"] = path
       return_value = self.rpc_client.send_rpc_message( method="file_directory",parameters=parameters,timeout=3 )
       if return_value[0] != True:
          raise ValueError("load file failure")
       return return_value[1]  
           

   def delete_file(self, path,file_name):
       #print("delete_file")
       parameters = {}
       parameters["path"] = path
       parameters["file_name"] = file_name
       return_value = self.rpc_client.send_rpc_message( method="delete_file",parameters=parameters,timeout=3 )
       if return_value[0] != True:
          raise ValueError("load file failure")
       return return_value[1]  
   
   def mkdir(self,path):
       #print("mkdir")
       parameters = {}
       parameters["path"] = path
       return_value = self.rpc_client.send_rpc_message( method="make_dir",parameters=parameters,timeout=3 )
       if return_value[0] != True:
          raise ValueError("load file failure")
       return return_value[1]  
   
 
  

   
 
 
       

 

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
    
    
    
    file_client = Construct_RPC_Library(qs,redis_site)
    print(file_client.file_directory(""))
    print(file_client.mkdir("test_path"))
    print(file_client.file_directory(""))
    print(file_client.save_file("test_path","test_file.test","hi\nthere\nbrown\ncow"))
    print(file_client.load_file("test_path","test_file.test"))
    print(file_client.file_directory("test_path"))
    print(file_client.file_exists("test_path","test_file.test"))
    print(file_client.delete_file("test_path","test_file.test"))
    print(file_client.file_directory("test_path"))