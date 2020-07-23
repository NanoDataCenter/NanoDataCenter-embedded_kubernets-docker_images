from .sqlite_server_lib_py3 import Construct_RPC_Library
from redis_support_py3.graph_query_support_py3 import  Query_Support
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
import datetime
import msgpack

class SQLITE_Client_Support(Construct_RPC_Library):

   def __init__( self, qs, site_data ):
       Construct_RPC_Library.__init__(self, qs,site_data)
    
   def list_tables(self,database_name): #tested
       script = 'SELECT name from sqlite_main where type= "table";'
      
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
       
   
 
   
   def create_text_search_table(self,database_name,table_name, fields ): #test
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
          
   
       
   def select_composite(self,database_name,table_name,return_fields,where_clause=None,distinct_flag=False): # tested
       #"select rowid, name, ingredients from recipe where name match 'pie'"
       if distinct_flag != False:
          script = "select distinct "
       else:
          script = "select "
       script = script+"  "+",".join(return_fields) + "  FROM "+table_name+" "
       if where_clause != None:
           script = script+ "  where "+where_clause+" ; "
       else:
           script = script+" ; "
       
      
       return self.select(database_name,script)
       

   def insert_composite(self,database_name,table_name,field_names,field_dictionary): #tested
       field_values = []
       for i in field_names:
           field_values.append(field_dictionary[i])

       script =  'INSERT INTO '+table_name+' ('+",".join(field_names)+" ) "+ " VALUES("+self.join_items(field_values)+");"
       print("insert script  ",script)
       self.ex_script(database_name,script)
   
   def join_items(self,items):
       return_values = []
       for i in items:
           if type(i)== str:
              return_values.append("'"+str(i)+"'")
           else:
              return_values.append(str(i))
       return ",".join(return_values)
               
   def delete(self,database_name,table_name,where_clause):  #tested  
       script = 'DELETE FROM '+table_name+' WHERE '+where_clause+";"
       return self.ex_exec(database_name,script)
  
 
   def update(self,database,table_name,row_id,row_values,where_clause=None):
       if len(row_id) != len(row_values):
          raise ValueError("row id and row values are not same length")          
   
       script = "UPDATE "+table_name+" SET  "
       filtered_values=[]
       for i in row_values:
          if type(i) == str:
             filtered_values.append('"'+i+'"')
          else:
              filtered_values.append(str(i))
       for i in range(0,len(filtered_values)):
           if i != len(filtered_values)-1:
               script = script +" "+row_id[i] +" = "+filtered_values[i]+", "
           else:
               script = script +" "+row_id[i] +" = "+filtered_values[i]+" "
       if where_clause != None:
           script = script + "WHERE  "+where_clause+ " ; "
       else:
           script = script+ " ;"
       print(script)    
       self.ex_exec(database,script)
   
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
    field_values = [
    [.1, 'broccoli peppers cheese tomatoes'],
    [.15, "test data"],
    [.2, 'pumpkin onions garlic celery'],
    [.3, 'broccoli cheese onions flour'],
    [.3, 'duplicate value'],
    [.4, 'pumpkin sugar flour butter']]
    
    print("insert_composite",sqlite_client.insert_composite("test","text",["a","b"],field_values))
    print("select_composite",sqlite_client.select_composite("test","text",["a","b"]))
    print("select_composite",sqlite_client.select_composite("test","text",["a","b"],"a < .2"))
    print("select_composite",sqlite_client.select_composite("test","text",["a"],"a >= .2",True))

    print(sqlite_client.delete("test","text","a <= .2"))
    print("select_composite",sqlite_client.select_composite("test","text",["a","b"]))
    
    #
    #
    #  Now testing alter command
    #
    #
    print("tables",sqlite_client.list_tables("test"))
    print("alter_table_rename",sqlite_client.alter_table_rename("test","text","new_text" ))
    print("tables",sqlite_client.list_tables("test"))
    print(sqlite_client.create_table("test","test",["a text","b text"]))
    print("insert_composite",sqlite_client.insert_composite("test","test",["a","b"],field_values))
    print("select_composite",sqlite_client.select_composite("test","test",["a","b"]))
    print("alter_table_add_column",sqlite_client.alter_table_add_column("test","test","c text"    ))
    print("schema",sqlite_client.get_table_schema("test","test"))
    print("select_composite",sqlite_client.select_composite("test","test",["a","b","c"]))
    
    #
    #
    # Testing Update Command
    #
    #
    print("update ",sqlite_client.update("test","test",["c"],["default_value"],where_clause="a>.2"))
    print("select_composite",sqlite_client.select_composite("test","test",["a","b","c"]))
    print("update ",sqlite_client.update("test","test",["c"],["new default"]))
    print("select_composite",sqlite_client.select_composite("test","test",["a","b","c"]))

    
    print(sqlite_client.close_database("test")) 
    print(sqlite_client.delete_database("test"))
   