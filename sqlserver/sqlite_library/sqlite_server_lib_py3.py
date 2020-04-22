class Construct_RPC_Library(object):

   def __init__( self, qs, site_data ):
          
       self.site_data = site_data
       self.qs        = qs
       query_list = []
       query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
       query_list = qs.add_match_relationship( query_list,relationship= "SQL_SERVER" )
       query_list = qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", property_mask={"name":"SQL_SERVER"} )
       package_sets, package_sources = qs.match_list(query_list) 
       package = package_sources[0]    
       data_structures = package["data_structures"]
       queue_name = data_structures["SQL_SERVER_RPC_SERVER"]['queue']
       generate_handlers = Generate_Handlers( package, qs )
       self.rpc_client = generate_handlers.construct_rpc_client( )
       self.rpc_client.set_rpc_queue(queue_name)
    
   def filter_result(self, return_value):
       if return_value[0] == True:
           return return_value[1]
       else:
           raise ValueError(return_value[1])
  
   def list_data_bases(self):
       print("list_data_bases")
       parameters = {}
       return_value = self.rpc_client.send_rpc_message( method="list_list_data_bases",parameters=parameters,timeout=3 )
       print("return_value",return_value)
       return self.filter_result(return_value)       
       
   def create_database(self,database):
       parameters = {}
       parameters["database"] = database
       return_value = self.rpc_client.send_rpc_message( method="create_database",parameters=parameters ,timeout=3 )
       return self.filter_result(return_value)       
           
   def close_database(self,database):
       parameters = {}
       parameters["database"] = database      
       return_value = self.rpc_client.send_rpc_message( method="delete_database", parameters=parameters ,timeout=3 )
       return self.filter_result(return_value)   
 
 
   def delete_database(self,database):
       parameters = {}
       parameters["database"] = database
       return_value = self.rpc_client.send_rpc_message( method="close_database",parameters=parameters ,timeout=3 )
       return self.filter_result(return_value)              


            

   def vacuum(self,database):
       parameters = {}
       parameters["database"] = database
       return_value = self.rpc_client.send_rpc_message( method="vacuum",parameters=parameters ,timeout=3 )
       return self.filter_result(return_value)       

       
   def version(self):
       parameters = {}
       return_value = self.rpc_client.send_rpc_message( method="version",parameters=parameters ,timeout=30 )
       return self.filter_result(return_value)  
       
   def set_text(self,database,text_state):
       parameters = {}
       parameters["database"] = database  
       parameters["text_state"] = text_state       
       return_value = self.rpc_client.send_rpc_message( method="set_txt",parameters=parameters ,timeout=30 )
       return self.filter_result(return_value)

       
   def get_text(self,database):
       parameters = {}
       parameters["database"] = database
       return_value = self.rpc_client.send_rpc_message( method="get_txt",parameters=parameters ,timeout=30 )
       return self.filter_result(return_value)    
       
   def backup(self,database,backup_db,pages= 0):
       parameters = {}
       parameters["database"] = database
       parameters["backup_db"] = backup_db
       parameters["pages"] = pages
       return_value = self.rpc_client.send_rpc_message( method="backup",parameters=parameters ,timeout=30 )
  
   def ex_exec(self,database,script):
       parameters = {}
       parameters["database"] = database
       parameters["script"] = script
       return_value = self.rpc_client.send_rpc_message( method="execute",parameters=parameters ,timeout=30 )
  
   def ex_script(self,database,script):
       parameters = {}
       parameters["database"] = database
       parameters["script"] = script
       return_value = self.rpc_client.send_rpc_message( method="execute_script",parameters=parameters ,timeout=30 )
          
       
   def commit(self,database):
       parameters = {}
       parameters["database"] = database
       return_value = self.rpc_client.send_rpc_message( method="commit",parameters=parameters ,timeout=30 )
       
   def select(self,database,script):
       parameters = {}
       parameters["database"] = database
       parameters["script"] = script
       return_value = self.rpc_client.send_rpc_message( method="select",parameters=parameters ,timeout=30 )
   
   
 
 
       

 

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
    from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
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
    
    
    
    sqlite_client = Construct_RPC_Library(qs,redis_site)
    print(sqlite_client.list_data_bases())
    try:
       print(sqlite_client.create_database("test"))
    except:
        print("duplicate db")
    print(sqlite_client.list_data_bases())
    print(sqlite_client.close_database("test"))
    print(sqlite_client.delete_database("test"))
