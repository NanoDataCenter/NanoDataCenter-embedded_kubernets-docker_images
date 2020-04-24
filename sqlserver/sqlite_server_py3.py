import msgpack
import redis
import base64
import os
import sqlite3
import sys

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
       self.rpc_queue_handle.register_call_back( "delete_database",self.delete_database)
       self.rpc_queue_handle.register_call_back("close_database",self.close_database)
       self.rpc_queue_handle.register_call_back("vacuum",self.vacuum)
       self.rpc_queue_handle.register_call_back("version",self.version)
       self.rpc_queue_handle.register_call_back("set_txt",self.set_text)
       self.rpc_queue_handle.register_call_back("get_txt",self.get_text)
       self.rpc_queue_handle.register_call_back("backup",self.backup)
       self.rpc_queue_handle.register_call_back("execute",self.ex_exec)
       self.rpc_queue_handle.register_call_back("execute_script",self.ex_script)
       self.rpc_queue_handle.register_call_back("commit",self.commit)
       self.rpc_queue_handle.register_call_back("select",self.select)
       self.rpc_queue_handle.add_time_out_function(self.process_null_msg)
       self.rpc_queue_handle.start()
   
   def process_null_msg( self ):  
       print("null message")         
 
   def initialize_data_base_handles(self):
       self.db_handlers = {}
       if self.sql_databases.hexists("default") == False:
           self.sql_databases.hset("default","/sqlite/default.db")
       if self.sql_databases.hexists("memory") == False:
             self.sql_databases.hset("memory",":memory:")
       
       for i in self.sql_databases.hkeys():
           try:
               if i == "default":
                  self.db_handlers[i] = sqlite3.connect("/sqlite/default.db") 
                  self.db_handlers[i].row_factory = sqlite3.Row
               elif i == "memory":
                  self.db_handlers[i] = sqlite3.connect(":memory:") 
                  self.db_handlers[i].row_factory = sqlite3.Row
               else:
                  self.db_handlers[i] = sqlite3.connect("/sqlite/"+i)
                  self.db_handlers[i].row_factory = sqlite3.Row
               
           except:
               print("db connection problem ",i,self.sql_databases.hget(i))           
               self.sql_databases.hdelete(i)
               
                      
 
 
   def process_null_msg( self ):  
       print("null message")   
  
   def list_data_bases(self,input_message):
       print("list data base")
       return_value = self.sql_databases.hgetall()
       return [True,return_value]       
       
   def create_database(self,input_msg):
       print("create database")
       try:
           name = input_msg["database"]
           if self.sql_databases.hexists(name) == False:
              file_path = "/sqlite/"+name+".db"
              connection = sqlite3.connect(file_path)
              self.sql_databases.hset(name,file_path)
              self.db_handlers[name] = connection
              self.db_handlers[name].row_factory = sqlite3.Row
              return [True,""]
           else:
               return [False, "duplicate database" ]
       except :
           return [False,str(sys.exc_info()[:2])]
           
           
   def close_database(self,input_msg):
       try:
           name = input_msg["database"]
           if name  in self.db_handlers:
              connection = self.db_handlers[name]
              connection.close()
              del self.db_handlers[name]
              self.sql_databases.hdelete(name)
              return [True,""]
           else:
              return [False,"non existant database"]
       except :
           return [False,str(sys.exc_info()[:2])]
 
   def delete_database(self,input_msg):
       try:
           name = input_msg["database"]
           if name not in self.db_handlers :  # database must be closed to delete file
              file_path = "/sqlite/"+name+".db"
              try:
                  os.remove(file_path)
              except:
                pass
              return [True,""]
           else:
              return [False,"data base is still active" ]
       except :
           return [False,str(sys.exc_info()[:2])]
           


            

   def vacuum(self,input_msg):
       try:
           db_name = input_msg["database"]
           connection = self.db_handlers[db_name]
           connection.execute("VACUUM")
           return [True,'']
       except :
           return [False,str(sys.exc_info()[:2])]
       
   def version(self,input_msg):
       try:
 
           return [True,sqlite3.version]
       except :
          return [False,str(sys.exc_info()[:2])]
   
   def set_text(self,input_msg):
       try:
           db_name = input_msg["database"]
           state = input_msg["text_state"]
           if state == "bytes":
              state = bytes
           else:
              state = str
           connection = self.db_handlers[db_name]
           connection.text_factory = state
           return [True,'']
       except :
           return [False,str(sys.exc_info()[:2])]
          
   def get_text(self,input_msg):
       try:
           db_name = input_msg["database"]
           connection = self.db_handlers[db_name]
           state = connection.text_factory
           if state == bytes:
              state = "bytes"
           else:
              connection.text_factory = str
              state = "string"
              
           return [True,state]
       except :
           return [False,str(sys.exc_info()[:2])]
  
   def backup(self,input_msg):
       try:
           db_name = input_msg["database"]
           backup_db = input_msg["backup_db"]
           pages     = input_msg["pages"]
           connection = self.db_handlers[db_name]
           backup_connection = self.db_handlers[backup_db]
           connection.backup(target=backup_connection,pages = pages)
           return [True,'']
       except :
          
           return [False,str(sys.exc_info()[:2])]
  
   def ex_exec(self,input_msg):
       try:
           db_name = input_msg["database"]
           script  = input_msg["script"]
           connection = self.db_handlers[db_name]
           connection.execute(script)
           connection.commit()
           return [True,'']
       except :
           return [False,str(sys.exc_info()[:2])]
  
   def ex_script(self,input_msg):
       try:
           db_name = input_msg["database"]
           script  = input_msg["script"]
           connection = self.db_handlers[db_name]
           connection.executescript(script)
           connection.commit()
           return [True,'']
       except :
           return [False,str(sys.exc_info()[:2])]
          
       
   def commit(self,input_msg):
       try:
           db_name = input_msg["database"]       
           connection = self.db_handlers[db_name]        
           connection.commit()
           return [True,'']
       except :
           return [False,str(sys.exc_info()[:2])]

   def select(self,input_msg):
       print("select")
       try:
           db_name = input_msg["database"]   
           script  = input_msg["script"]           
           connection = self.db_handlers[db_name]        
           cursor = connection.execute(script)
          
           
           return_value = []
           while True:
              r = cursor.fetchone()
              print("r",r)
              
              
              if r==None:
                 break
              temp = {}
              for j in r.keys():
                temp[j] = r[j]
              
              return_value.append(temp)
              
           print(return_value)
           return [True,return_value]
       except :
         
           return [False,str(sys.exc_info()[:2])]
             

   
   
 
 
       
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
    
