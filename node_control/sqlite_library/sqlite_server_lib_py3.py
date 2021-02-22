from redis_support_py3.graph_query_support_py3 import  Query_Support
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
import datetime
import msgpack
from Pattern_tools_py3.builders.common_directors_py3 import construct_all_handlers
  

class Construct_RPC_Library(object):

   def __init__( self, qs, site_data ):
          
       self.site_data = site_data
       self.qs        = qs
        
       search_list = ["SQL_CLIENT","SQL_CLIENT"]
    
       self.ds_handlers = construct_all_handlers(site_data,qs,search_list)
       
       self.rpc_client = self.ds_handlers['SQL_SERVER_RPC_CLIENT']
       self.time_out = 15
    
   def filter_result(self, return_value):
       if return_value[0] == True:
           return return_value[1]
       else:
           raise ValueError(return_value[1])
  

  
   def list_data_bases(self):
       
       parameters = {}
       return_value = self.rpc_client.send_rpc_message( method="list_list_data_bases",parameters=parameters,timeout=self.time_out )
       
       return self.filter_result(return_value)       
       
   def create_database(self,database):
       parameters = {}
       parameters["database"] = database
       return_value = self.rpc_client.send_rpc_message( method="create_database",parameters=parameters ,timeout=self.time_out )
       return self.filter_result(return_value)       
           
   def delete_database(self,database):
       parameters = {}
       parameters["database"] = database      
       return_value = self.rpc_client.send_rpc_message( method="delete_database", parameters=parameters ,timeout=self.time_out )
       return self.filter_result(return_value)   
 
 
   def close_database(self,database):
       parameters = {}
       parameters["database"] = database
       return_value = self.rpc_client.send_rpc_message( method="close_database",parameters=parameters ,timeout=self.time_out )
       return self.filter_result(return_value)              


            

   def vacuum(self,database):
       parameters = {}
       parameters["database"] = database
       return_value = self.rpc_client.send_rpc_message( method="vacuum",parameters=parameters ,timeout=self.time_out )
       return self.filter_result(return_value)       

       
   def version(self):
       parameters = {}
       return_value = self.rpc_client.send_rpc_message( method="version",parameters=parameters ,timeout=self.time_out )
       return self.filter_result(return_value)  
       
   def set_text(self,database,text_state):
       parameters = {}
       parameters["database"] = database  
       parameters["text_state"] = text_state       
       return_value = self.rpc_client.send_rpc_message( method="set_txt",parameters=parameters ,timeout=self.time_out)
       return self.filter_result(return_value)

       
   def get_text(self,database):
       parameters = {}
       parameters["database"] = database
       return_value = self.rpc_client.send_rpc_message( method="get_txt",parameters=parameters ,timeout=self.time_out )
       return self.filter_result(return_value)    
       
   def backup(self,database,backup_db,pages= 0):
       parameters = {}
       parameters["database"] = database
       parameters["backup_db"] = backup_db
       parameters["pages"] = pages
       return_value = self.rpc_client.send_rpc_message( method="backup",parameters=parameters ,timeout=self.time_out )
       return self.filter_result(return_value) 
       
   def ex_exec(self,database,script):
       parameters = {}
       parameters["database"] = database
       parameters["script"] = script
       return_value = self.rpc_client.send_rpc_message( method="execute",parameters=parameters ,timeout=self.time_out )
       return self.filter_result(return_value)    
       
   def ex_script(self,database,script):
       parameters = {}
       parameters["database"] = database
       parameters["script"] = script
       return_value = self.rpc_client.send_rpc_message( method="execute_script",parameters=parameters ,timeout=self.time_out)
       return self.filter_result(return_value)              
       
   def commit(self,database):
       parameters = {}
       parameters["database"] = database
       return_value = self.rpc_client.send_rpc_message( method="commit",parameters=parameters ,timeout=self.time_out )
       return self.filter_result(return_value)    
       
   def select(self,database,script):
       parameters = {}
       parameters["database"] = database
       parameters["script"] = script
       return_value = self.rpc_client.send_rpc_message( method="select",parameters=parameters ,timeout=self.time_out )
       return self.filter_result(return_value)    
   
   
 
 
       

 

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
    
    print(sqlite_client.list_data_bases())
    #os.system("ls /sqlite")
    try:
        print(sqlite_client.create_database("test"))
    except:
        print("duplicate db")
    try:
       print(sqlite_client.create_database("backup"))
    except:
        print("duplicate db")
    print(sqlite_client.version())
    print(sqlite_client.get_text("test"))
    print(sqlite_client.set_text("test","bytes"))
    print(sqlite_client.get_text("test"))
    print(sqlite_client.set_text("test","string"))
    print(sqlite_client.get_text("test"))
    print(sqlite_client.backup("test","backup"))
    temp = '''create table recipe( name text, ingredients text);'''
    print(sqlite_client.ex_exec("test",temp))
    temp = """
    insert into recipe (name, ingredients) values ('broccoli stew', 'broccoli peppers cheese tomatoes');
    insert into recipe (name, ingredients) values ('pumpkin stew', 'pumpkin onions garlic celery');
    insert into recipe (name, ingredients) values ('broccoli pie', 'broccoli cheese onions flour');
    insert into recipe (name, ingredients) values ('pumpkin pie', 'pumpkin sugar flour butter');
    """
    print(sqlite_client.ex_script("test",temp))
    print(sqlite_client.commit("test"))
    print(sqlite_client.select("test","select * from recipe"))
    print(sqlite_client.select("test","select rowid,name,ingredients from recipe"))
    # intentional bad sql
    
    #print(sqlite_client.select("test","select "))
    print(sqlite_client.close_database("test"))
    print(sqlite_client.close_database("backup"))
    print(sqlite_client.delete_database("test"))
    print(sqlite_client.delete_database("backup"))
    