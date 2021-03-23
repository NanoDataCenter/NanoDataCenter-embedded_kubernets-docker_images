

import redis
import time
import base64
import msgpack
import uuid

from .redis_stream_utilities_py3 import Redis_Stream_Utilities


'''

class SQL_LOG_TABLES(object):
   def __init__(self,properties    ):
       self.rpc_handler = None
       self.properties = properties
       self.required_properties = ["database_name","table_name","field_names"]
       self.verify_properties()
    
   
   def set_rpc_handler(self,handler):
       print("made it here")
       self.rpc_handler = handler  
       self.check_db()
       self.check_table()
       
       

    
       
   def push(self,data):
       field_values = self.orders_fields(data)
       self.rpc_handler.insert_composite(self.database_name,self.table_name,self.field_names,[field_values])
       
       
   def trim_stream(self,time):
       self.rpc_handler.delete(self.database_name,self.table_name,"time < "+str(time));
   
   def select(self,where_clause=None,distinct_flag = False):
      return self.rpc_handler.select_composite(self.database_name,self.table_name,self.field_names,where_clause=where_clause,distinct_flag=distinct_flag)
 
   def verify_properties(self):
       for i in self.required_properties:
          if i not in self.properties:
              raise ValueError("property "+str(i)+" is not defined ")
       self.database_name = self.properties["database_name"]
       self.table_name    = self.properties["table_name"]
       self.field_names   = self.properties["field_names"]       
       
       
   def check_db(self):
       db_tables = self.rpc_handler.list_data_bases().keys()
       if self.database_name not in db_tables:
          self.rpc_handler.create_database(self.database_name)
   
   def check_table(self):
       current_tables = self.rpc_handler.list_tables(self.database_name)
       if self.table_name not in current_tables:
            self.rpc_handler.create_table(self.database_name,self.table_name,self.field_names)
       
   def orders_fields(self, data):
       return_value = []
       for i in self.field_names:
           return_value.append(data[i])
       return return_value
           

   
class SQL_TEXT_SEARCH_LOG_TABLES(SQL_LOG_TABLES):

    def __init__(self,properties  ):
        SQL_LOG_TABLES.__init__(self,properties )
        
    def check_table(self):
       current_tables = self.rpc_handler.list_tables(self.database_name)
       if self.table_name not in current_tables:
            self.rpc_handler.create_text_search_table(self.database_name,self.table_name,self.field_names)
   
    def select_text_search_general(self,search_term):
        where_clause = self.table_name+ "  MATCH  '"+search_term+"'"
        return self.rpc_handler.select_composite(self.database_name,self.table_name,self.field_names,where_clause=where_clause)    
     
     
''' 

class Field_Not_Defined(Exception):
    pass       
      
class FIELD_TYPE_ERROR(Exception):
   pass



class Redis_RPC_Client(object):

   def __init__( self,redis_handle,properties ):
       self.properties = properties
       key = properties["queue"]
       self.redis_handle = redis_handle
       self.set_rpc_queue(key)
       
   
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
                
class RPC_Server(object):
    def __init__( self, redis_handle ,properties, redis_rpc_queue ):
       self.properties = properties
       self.redis_handle = redis_handle
       self.redis_rpc_queue = redis_rpc_queue
       self.handler = {}
       self.redis_handle.delete(redis_rpc_queue)
       self.timeout_function = None
       self.timeout_value  = properties["timeout"]
       

    def add_time_out_function(self,time_out_function):
       self.timeout_function = time_out_function
       

    def register_call_back( self, method_name, handler):
        self.handler[method_name] = handler
    
    def start( self ):
        while True:
            try:
               input = self.redis_handle.brpop(self.redis_rpc_queue,self.timeout_value)
              
               if input == None:
                    if self.timeout_function != None:
                        self.timeout_function()
               else:
                   input = msgpack.unpackb(input[1])  # 0 parameter is the queue
                   #print("input_message",input)
                   self.process_message(  input )
                       
            except:
                raise
 
    def process_message( self, input):

        id      = input["id"]
        method  =  input["method"]
        params  = input["params"]
        response = self.handler[method](params)
       
        self.redis_handle.lpush( id, msgpack.packb(response))        
        self.redis_handle.expire(id, 30)

class String_Field(object):
   def __init__(self,handler):
       self.handler = handler
       
   def hget(self,setup,field):
      result = self.handler.hget( field)
      if result == None:
        result = setup["init_value"]
      return str(result)
      
   def hset(self,setup,field,data):
       self.handler.hset(field,str(data))
       
class Float_Field(object):
   def __init__(self,handler):
       self.handler = handler
       
   def hget(self,setup,field):
      result = self.handler.hget(field)
      if result == None:
        result = setup["init_value"]
      return float(result)
      
   def hset(self,setup,field,data):

       self.handler.hset(field,float(data))
 

 
class Binary_Field(object):
   def __init__(self,handler):
       self.handler = handler
       
   def hget(self,setup,field):
      result = self.handler.hget(field)
      if result == None:
         result = setup["init_value"]
      return bool(result) 
      
   def hset(self,setup,field,data):
       if data == 0:
          data = False
       if data == 1:
          data = True
       if isinstance(data,bool):
           self.handler.hset(field,data)
       else:
         raise ValueError("not a boolean type "+str(data))
 

class List_Field(object):
   def __init__(self,handler):
       self.handler = handler
       
   def hget(self,setup,field):
      result = self.handler.hget(field)
      if result == None:
         result = setup["init_value"]
      return bool(result) 
      
   def hset(self,setup,field,data):
       if isinstance(data, list):
           self.handler.hset(field,data)
       else:
         raise ValueError("not a list type "+str(data))
 

 
class Dictionary_Field(object):
   def __init__(self,handler):
       self.handler = handler
       
   def hget(self,setup,field):
      
      temp = self.handler.hget(field)
      if temp == None:
        result = setup["fields"]
      else:
         result = {}
         for i in setup["fields"].keys():
            if i in temp:
               
               result[i] = temp[i]
            else:
               result[i] = setup["fields"][i]
               
      return result
      
   def hset(self,setup,field,data):
       temp = {}
       for i in setup["fields"].keys():
         if i in data:
            temp[i] = data[i]
         else:
            raise ValueError("key: "+str(i)+" not in "+str(data))
       self.handler.hset(field,temp)

                 
class Managed_Redis_Hash(object):
   def __init__(self,redis_handle,properties, key  ):
        self.handler = Redis_Hash_Dictionary(redis_handle,properties, key  )
        self.fields = properties["fields"]
        self.field_handlers = {}
        self.field_handlers["string"] = String_Field(self.handler)
        self.field_handlers["float"] = Float_Field(self.handler)
        self.field_handlers["binary"] = Binary_Field(self.handler)
        self.field_handlers["list"]   = List_Field(self.handler)
        self.field_handlers["dictionary"] = Dictionary_Field(self.handler)      
        self.validate_graph_data()
        self.sanitize_keys()
 

   def get_rid_of_bad_keys(self):
       keys = self.handler.hkeys()
       for i in keys:
         if i not in self.fields:
            print("key "+str(i)+" doesnot belong")
            self.handler.hdelete(i)
            
   def sanitize_keys(self):
       
       for key,item in self.fields.items():
         temp = self.hget(key)
         self.hset(key,temp)
                
       
   def validate_graph_data(self):
       for i,item in self.fields.items():

          instance_type = item["type"]
          if instance_type not in self.field_handlers:
             raise ValueError("improper type:  "+str(key)+"  "+str(item))          
       
   def hget(self,field):
       if field in self.fields:
          item = self.fields[field]
          return self.field_handlers[item["type"]].hget(item,field)
       else:
          raise ValueError("field is not registered:  "+field)
          
   def hget_all(self):
      result = {}
      for i,item in self.fields.items():
         result[i] = self.field_handlers[item["type"]].hget(item,i)
      return result
      
   def hset(self,field,data):
      if field in self.fields:
          item = self.fields[field]
          self.field_handlers[item["type"]].hset(item,field,data)
      else:
          raise ValueError("field is not registered:  "+field)
           
 
class Redis_Hash_Dictionary( object ):
 
   def __init__(self,redis_handle,properties, key  ):
      self.properties = properties
      self.redis_handle = redis_handle
      self.key = key


      
   def delete_all( self ):
       self.redis_handle.delete(self.key)

      
      
   def hset( self, field, data ):
   
      pack_data = msgpack.packb(data )
      
      if self.redis_handle.hget(self.key,field)== pack_data: # donot propagte identical values
         return
      self.redis_handle.hset(self.key,field,pack_data)


   def hmset(self,dictionary_table):
       for i,items in dictionary_table.items():
          self.hset(i,items)
       

      
     
         
   def hget( self, field):
      
      pack_data = self.redis_handle.hget(self.key,field)
      
      if pack_data == None:
         return None
      
      return  msgpack.unpackb(pack_data)
      

   def hgetall( self ):
      return_value = {}
      keys = self.redis_handle.hkeys(self.key)
      
      for field in keys:
         try:
            new_field = field.decode('utf-8')
         except:
            new_field = field
         return_value[new_field] = self.hget(field)
     
      return return_value
      
   def hkeys(self):
       binary_list = self.redis_handle.hkeys(self.key)
       return_list = []
       for i in binary_list:
          return_list.append(i.decode())
       return return_list
   
   def hexists(self,field):
     return self.redis_handle.hexists(self.key,field)
   

   def hdelete(self,field):
       self.redis_handle.hdel(self.key,field)
          
   def delete_all( self ):
       self.redis_handle.delete(self.key)

class Single_Element(object):

   def __init__(self,redis_handle,properties,key  ):
       self.redis_handle = redis_handle
       self.key          = key
       self.properties    = properties
 

   def get( self):
      
      pack_data = self.redis_handle.get(self.key)
      
      if pack_data == None:
         return None
      
      return  msgpack.unpackb(pack_data)
  
   def set( self, data ):
   
      pack_data = msgpack.packb(data )
      
      if self.redis_handle.get(self.key)== pack_data: # donot propagte identical values
         return
      self.redis_handle.set(self.key,pack_data)

         
class Job_Queue_Client( object ):
 
   def __init__(self,redis_handle,properties, key):
      self.properties = properties
      self.redis_handle = redis_handle
      self.key = key
      self.depth =  properties["depth"]
      
   def delete_all( self ):
       self.redis_handle.delete(self.key)
      
      
   def delete(self, index ):
       if index < self.redis_handle.llen(self.key):
           self.redis_handle.lset(self.key, index,"__#####__")
           self.redis_handle.lrem(self.key, 1,"__#####__") 
      
   def length(self):
       return self.redis_handle.llen(self.key)   


   def list_range(self,start,stop):
      
      list_data =  self.redis_handle.lrange(self.key,0,-1)
     
      if list_data == None:
         return None
      return_value = []
      for pack_data in list_data:
        return_value.append(msgpack.unpackb(pack_data))
      return return_value
      
   def pop(self):
       pack_data = self.redis_handle.rpop(self.key)
        
 
       if pack_data == None:
          return False, None
       else:
         
          return True,msgpack.unpackb(pack_data)     
   def push(self,data):
       pack_data =  msgpack.packb(data )
       self.redis_handle.lpush(self.key,pack_data)
       self.redis_handle.ltrim(self.key,0,self.depth)

           
   def delete_jobs(self,data):
       for i in data:
         self.redis_handle.lset(self.key,i,"__DELETE_ME__")
       self.redis_handle.lrem(self.key,0,"__DELETE_ME__")
         

 
class Job_Queue_Server( object ):
 
   def __init__(self,redis_handle,properties, key  ):
      self.properties = properties
      self.redis_handle = redis_handle
      self.key = key 
      

 
   def delete_all( self ):
       self.redis_handle.delete(self.key)
 
 
   def length(self):
       return self.redis_handle.llen(self.key)
       
   def delete(self, index ):
       if index < self.redis_handle.llen(self.key):
           self.redis_handle.lset(self.key, index,"__#####__")
           self.redis_handle.lrem(self.key, 1,"__#####__") 
                

 
   def pop(self):
       pack_data = self.redis_handle.rpop(self.key)
        

       if pack_data == None:
          return False, None
       else:
         
          return True,msgpack.unpackb(pack_data)
          
   def show_next_job(self):
       pack_data = self.redis_handle.lindex(self.key, -1)
       if pack_data == None:
          return False, None
       else:
          
          return True, msgpack.unpackb(pack_data)
   def push_front(self,data):
       pack_data =  msgpack.packb(data )
       self.redis_handle.rpush(self.key,pack_data)
       self.redis_handle.ltrim(self.key,0,self.depth)
 


class Job_Queue(Job_Queue_Server, Job_Queue_Client):
      def __init__(self,redis_handle,properties,key ):
          Job_Queue_Client.__init__(self,redis_handle,properties,key  )
          Job_Queue_Server.__init__(self,redis_handle,properties,key  )     




class Stream_Redis_Writer(Redis_Stream_Utilities):
       
   def __init__(self,redis_handle,   properties,key  ):
      Redis_Stream_Utilities.__init__(self,redis_handle)
      
      self.properties = properties
     
      self.redis_handle = redis_handle

      self.key = key
      self.depth = properties["depth"]
      self.add_pad = "~"
      
      
   def save(self):
       self.redis_handle.save()

   def delete_all( self ):
       self.redis_handle.delete(self.key)
      
   def change_add_flag(self, state):
      if state == True:
          self.add_pad = ""
      else:
          self.add_pad = "~"
   
   def trim(self,key,length):
      self.trim(key,length)


   def push(self,data={} ,id="*",local_node = None ):
       
       
       if len(list(data.keys())) == 0:
           return
       packed_data  =msgpack.packb(data )
       out_data = {}
       out_data["data"] = packed_data
       
       self.xadd(key = self.key, max_len=self.depth,id=id,data_dict=out_data )
       
       
       
class Stream_Redis_Reader(Redis_Stream_Utilities):
       
   def __init__(self,redis_handle,properties, key):
      Redis_Stream_Utilities.__init__(self,redis_handle)
      self.properties = properties
      self.redis_handle = redis_handle
      self.key = key

   def delete_all( self ):
       self.redis_handle.delete(self.key)
      
      
   def range(self,start_timestamp, end_timestamp , count=100):
       if isinstance(start_timestamp,str) == False:
           start_timestamp = int(start_timestamp*1000)
       if isinstance(end_timestamp,str) == False:
           end_timestamp = int(end_timestamp*1000)

       data_list = self.xrange(self.key,start_timestamp,end_timestamp, count)

       return data_list

   def revrange(self,start_timestamp, end_timestamp , count=100):
       
       if isinstance(start_timestamp,str) == False:
           start_timestamp = int(start_timestamp*1000)
       if isinstance(end_timestamp,str) == False:
           end_timestamp = int(end_timestamp*1000)


       data_list = self.xrevrange(self.key,start_timestamp,end_timestamp, count)
       

       return data_list 

class Redis_Stream(Stream_Redis_Writer, Stream_Redis_Reader):
      def __init__(self,redis_handle,properties,key):
          Stream_Redis_Writer.__init__(self,redis_handle,properties,key  )
          Stream_Redis_Reader.__init__(self,redis_handle,properties,key  )
             
class Generate_Handlers(object):
   
   def __init__(self,package,qs ):
      
       self.package = package
       self.redis_handle = qs.get_redis_data_handle()

       self.constructors = {}
       self.constructors["SINGLE_ELEMENT" ] =   self.construct_single_element  
       self.constructors["HASH" ] =       self.construct_hash 
       self.constructors["MANAGED_HASH"] =  self.construct_managed_hash      
       self.constructors["STREAM_REDIS"] = self.construct_redis_stream_writer       
       self.constructors["JOB_QUEUE"] =     self.construct_job_queue_server   
       self.constructors["RPC_SERVER"] =    self.construct_rpc_sever    
       self.constructors["RPC_CLIENT"] =   self.construct_rpc_client  
       self.constructors["SQL_LOG_TABLE"] = self.construct_sql_log
       self.constructors["SEARCH_SQL_LOG_TABLE"] = self.construct_sql_text_search_log
    
   def get_redis_handle(self):
       return self.redis_handle  

   def generate_all(self):
       handlers = {}
       data_structures = self.package["data_structures"]
       for key,data in data_structures.items():
           
           if data["type"] in self.constructors:
               handlers[key] = self.constructors[data["type"]](data)
           else:
               raise ValueError("Bad handler key")
               
       return handlers         

   '''
   def construct_influx_handler(self):
 
       self.influx_handler =  Influx_Handler(self.influx_server,
                                             self.influx_user,
                                             self.influx_password,
                                             self.influx_database,
                                             self.influx_retention)
       self.influx_handler.switch_database(self.influx_database)

   '''
   def construct_single_element(self,data):
       assert(data["type"] == "SINGLE_ELEMENT")
       key = self.package["namespace"]+"["+data["type"]+":"+data["name"] +"]"     
       return   Single_Element( self.redis_handle,data,key )  

       
   def construct_hash(self,data):
         assert(data["type"] == "HASH")
         key = self.package["namespace"]+"["+data["type"]+":"+data["name"] +"]"
         
         return  Redis_Hash_Dictionary( self.redis_handle,data,key )

   def construct_managed_hash(self,data):
         assert(data["type"] == "MANAGED_HASH")
         key = self.package["namespace"]+"["+data["type"]+":"+data["name"] +"]"
         
         return  Managed_Redis_Hash( self.redis_handle,data,key)
   


   def construct_stream_writer(self,data):
       return self.construct_redis_stream_writer(data)
 
   def construct_redis_stream_writer(self,data):
       assert(data["type"] == "STREAM_REDIS")
       key = self.package["namespace"]+"["+data["type"]+":"+data["name"] +"]"
       return Redis_Stream(self.redis_handle,data,key)

   def construct_stream_reader(self,data):
       return self.construct_redis_stream_reader(data)
         
     
   def construct_redis_stream_reader(self,data):
         assert(data["type"] == "STREAM_REDIS")
         
         key = self.package["namespace"]+"["+data["type"]+":"+data["name"] +"]"
         return Stream_Redis_Reader(self.redis_handle,data,key)
  
   def construct_job_queue_client(self,data):
         assert(data["type"] == "JOB_QUEUE")
         key = self.package["namespace"]+"["+data["type"]+":"+data["name"] +"]"
         return Job_Queue(self.redis_handle,data,key )
                                                                                                                                                                                                                                                                                 
   def construct_job_queue_server(self,data):
         assert(data["type"] == "JOB_QUEUE")
         key = self.package["namespace"]+"["+data["type"]+":"+data["name"] +"]"
         return Job_Queue(self.redis_handle,data,key)

   def construct_rpc_client(self,data):
        
         
         
         return Redis_RPC_Client(self.redis_handle,data)

   def construct_rpc_sever(self,data):
         assert(data["type"] ==  "RPC_SERVER")
         key = data["queue"]
         return RPC_Server(self.redis_handle,data,key )
         
   def construct_sql_log(self,data):
       assert(data["type"] ==  "SQL_LOG_TABLE")
       return SQL_LOG_TABLES(data )
       
   
   def construct_sql_text_search_log(self,data):
       assert(data["type"] ==  "SEARCH_SQL_LOG_TABLE")
       return  SQL_TEXT_SEARCH_LOG_TABLES(data )
      
   
if __name__== "__main__":
    pass
