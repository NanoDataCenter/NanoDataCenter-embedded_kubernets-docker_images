from common_tools.Pattern_tools_py3.builders.common_directors_py3 import construct_all_handlers
import datetime
import msgpack

  

class Construct_RPC_File_Library(object):

   def __init__( self, qs, site_data ):
       search_list = [ "FILE_SERVER_CLIENT", "FILE_SERVER_CLIENT"]
       self.ds_handlers = construct_all_handlers(site_data,qs,search_list)
       self.rpc_client = self.ds_handlers["FILE_SERVER_RPC_SERVER"]
    
    
   def check_file(self,file_name,return_value):
       if return_value[0] == False:
          raise ValueError("file does not exist  "+file_name)
       return return_value[1]          
    
   def load_file(self,path,file_name):

       parameters = {}
       parameters["path"] = path
       parameters["file_name"] = file_name
       return_value = self.rpc_client.send_rpc_message( method="load",parameters=parameters,timeout=3 )
       return self.check_file(file_name,return_value)
       
   def save_file(self,path,file_name,data):
       
       parameters = {}
       parameters["path"] = path
       parameters["file_name"] = file_name
       parameters["data"] = data
       return_value = self.rpc_client.send_rpc_message( method="save",parameters=parameters,timeout=3 )
       return self.check_file(file_name,return_value)
   
   def file_exists(self,path,file_name):
      
       parameters = {}
       parameters["path"] = path
       parameters["file_name"] = file_name
       return_value = self.rpc_client.send_rpc_message( method="file_exists",parameters=parameters,timeout=3 )
       return return_value[0]
        
   def file_directory(self,path):
       
       parameters = {}
       parameters["path"] = path
       return_value = self.rpc_client.send_rpc_message( method="file_directory",parameters=parameters,timeout=3 )
       return self.check_file(path,return_value)
           

   def delete_file(self, path,file_name):
       
       parameters = {}
       parameters["path"] = path
       parameters["file_name"] = file_name
       return_value = self.rpc_client.send_rpc_message( method="delete_file",parameters=parameters,timeout=3 )
       return return_value[0]
   
   def mkdir(self,path):
      
       parameters = {}
       parameters["path"] = path
       return_value = self.rpc_client.send_rpc_message( method="make_dir",parameters=parameters,timeout=3 )
       return return_value[0]
   
 
  

   
 
 
       

 

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