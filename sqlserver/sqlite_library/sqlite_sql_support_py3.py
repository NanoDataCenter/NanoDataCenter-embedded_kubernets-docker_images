from .sqlite_server_lib_py3 import Construct_RPC_Library
from redis_support_py3.graph_query_support_py3 import  Query_Support
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
import datetime
import msgpack

class SQLITE_Client_Support(Construct_RPC_Library):

   def __init__( self, qs, site_data ):
       Construct_RPC_Library.__init__(self, qs,site_data)
    
   def list_tables(self,database_name): #tested
       script = 'SELECT name from sqlite_master where type= "table";'
      
       temp = self.select(database_name,script)
      
       return temp
    
   def create_table(self,database_name,table_name,fields,temp_table=False,not_exists=True):#tested
       script = "create  "
       if temp_table != False:
          script=script+" TEMP "
       script = script+" TABLE "
       if not_exists == True:
          script = script+" IF NOT EXISTS "
       script = script+table_name+" "
       script = script+"( "+",".join(fields)+" ); "
       
       return self.ex_exec(database_name,script)
       
   
 
   
   def create_text_search_table(self,database_name,table_name, fields ):#tested
       script = "create virtual table  "
      
       script = script+table_name+" "
       script = script+" using fts5("+",".join(fields)+",tokenize = 'porter unicode61 remove_diacritics 1'  );"
       print("script",script)
       return self.ex_exec(database_name,script)
       
   def get_table_schema(self,database_name,table_name): #tested
       script = "PRAGMA table_info('"+table_name+"')"
       return self.select(database_name,script)

   def drop_table(self,database_name,table_name): #tested
       script = 'DROP TABLE if exists '+table_name+" ;" 
       return self.ex_exec(database_name,script)

   def alter_table_add_column(self,database_name,table_name,new_column ):
       script = 'ALTER TABLE '+table_name+' ADD COLUMN '+new_column;  
       return self.ex_exec(database_name,script)
 

   def alter_table_rename(self,database_name,old_table,new_table ):
       script = 'ALTER TABLE '+old_table+' RENAME TO '+new_table;  
       return self.ex_exec(database_name,script)
          
   
       
   def select_composite(self,database_name,table_name,return_fields,fields,where_clause,distinct_flag=False):
       #"select rowid, name, ingredients from recipe where name match 'pie'"
       script = "select "+",".join(return_fields)+ "  where "+where_clause+" );"
       return self.select(database_name,script)
       

   def insert_composite(self,database_name,table_name,field_names,field_values):
       for i in field_values:
           if len(field_names) != len(i):
              raise ValueError("field names and field values are not same length")       
           script = 'INSERT INTO '+table_name+'('+",".join(field_names)+ " VALUES("+",".join(i)+");"
           self.ex_exec(database_name,script)
   

   def delete(self,database_name,table_name,where_clause):   
       script = 'DELETE FROM '+table_name+' WHERE '+where_clause+";"
       return self.ex_exec(database_name,script)

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
    
    
    
    sqlite_client = SQLITE_Client_Support(qs,redis_site)
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
 
    print("tables",sqlite_client.list_tables("test"))
    print(sqlite_client.create_table("test","test",["a text","b text"],temp_table=False,not_exists=True))
    print("tables",sqlite_client.list_tables("test"))
    print("drop_table",sqlite_client.drop_table("test","test"))
    print("tables",sqlite_client.list_tables("test"))
    print(sqlite_client.create_table("test","test",["a text","b text"],temp_table=False,not_exists=True))
    print("tables",sqlite_client.list_tables("test"))
    print("schema",sqlite_client.get_table_schema("test","test"))
    print(sqlite_client.create_text_search_table("test","text",["a","b"]))
    print("tables",sqlite_client.list_tables("test"))
    print("schema",sqlite_client.get_table_schema("test","text"))

    temp = '''create table recipe( name text, ingredients text);'''
    
    #print(sqlite_client.ex_exec("test",temp))
    temp = """
    insert into recipe (name, ingredients) values ('broccoli stew', 'broccoli peppers cheese tomatoes');
    insert into recipe (name, ingredients) values ('pumpkin stew', 'pumpkin onions garlic celery');
    insert into recipe (name, ingredients) values ('broccoli pie', 'broccoli cheese onions flour');
    insert into recipe (name, ingredients) values ('pumpkin pie', 'pumpkin sugar flour butter');
    """
    #print(sqlite_client.ex_script("test",temp))

    #print(sqlite_client.select("test","select * from recipe"))
    # print(sqlite_client.select("test","select rowid,name,ingredients from recipe"))
    
    print(sqlite_client.close_database("test")) 
    print(sqlite_client.delete_database("test"))
   