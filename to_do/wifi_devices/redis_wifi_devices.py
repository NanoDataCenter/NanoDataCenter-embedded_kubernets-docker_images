from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
from   redis_support_py3.raw_data_structures_py3 import Raw_Key_Data_Handlers
import time
import statistics

SCAN_INTERVAL = .25 #hours
  

class WIFI_Devices(object):

   def __init__(self,wifi_source,handlers,wifi_devices ):
       
       self.device_status = {} #dummy initialization
       self.internal_queue = {}
       self.wifi_source = wifi_source
       self.handlers = handlers
       self.wifi_devices = wifi_devices
       self.redis_device_handle  = redis.StrictRedis( host = self.wifi_source["ip"], port = self.wifi_source["port"],db = self.wifi_source["db"]) 
       self.wifi_data_handlers =  Raw_Key_Data_Handlers(self.redis_device_handle)
    
      
   def trim_device_data_structures(self,*params):
       trim_length_heart_beat = len(self.wifi_devices)*4*60*SCAN_INTERVAL

       self.wifi_data_handlers.stream_trim(self.wifi_source["queues"]["heart_beat"],trim_length_heart_beat)
       self.wifi_data_handlers.stream_trim(self.wifi_source["queues"]["reboot"],trim_length_heart_beat)
       self.wifi_data_handlers.stream_trim(self.wifi_source["queues"]["internal_streams"],trim_length_heart_beat*4)
       for key,item in self.wifi_devices.items():
            print("i",item)
            key = item["input_queue"]
            length = item["input_queue_length"]*2
            self.wifi_data_handlers.job_queue_trim(key,data,length)

       
   def monitor_device_input_queue(self,*params):
      data = self.wifi_data_handlers.job_queue_pop(self.wifi_source["queues"]["client_input"])
      if data != None:
         if 'NODE_ID' in data:
             node_id = data["NODE_ID"]
             if node_id in self.wifi_devices:
                 queue = self.wifi_devices [node_id]["alert_queue"]
                 #print("queue",queue)
                 if queue in self.handlers:
                    self.handlers[queue].push(data)
                 else:
                    print("bad queue")
             else:
                self.handlers["UNKNOWN_DEVICES"].hset(node_id,node_id)
         
         else:
             print("bad data packet")         
       
   
   def monitior_device_status(self,*params):
      
       self.reboot_streams = self.get_reboot_stream()
       
       self.heart_beat_stream = self.get_heart_beat_stream()
       device_status = self.sort_streams(self.reboot_streams, self.heart_beat_stream)
       device_status = self.check_for_device_status(device_status) 
       device_status =self.analyize_reboot_times(device_status)
       device_status = self.analyize_heap_values(device_status)
       
       for i in self.wifi_devices:
           device_status[i]["raw_reboot"] = None
           device_status[i]["raw_heart_beat"] = None
           self.handlers["DEVICE_STATUS"].hset(i,device_status[i])
       print("device_status",device_status) 
       self.device_status = device_status
  
 
   def monitor_external_queue(self,*params):
       data = self.get_external_streams()
       
       external_queue = {}
       for key,data_item in data.items():
           fields = self.find_fields(data_item)
           external_queue[key] = {}
           temp = data[key]
           print("fields",fields)
           for j in fields:
               internal_queue[key][j] = {}
               stream_data = self.analyize_internal_stream(j,temp)
               external_queue[key][j]["max"] = max(stream_data)
               external_queue[key][j]["min"] =min(stream_data)                 
               external_queue[key][j]["mean"] = statistics.mean(stream_data)
               external_queue[key][j]["median"] = statistics.median(stream_data) 
               if len(stream_data) >1:
                  external_queue[key][j]["std"] = statistics.stdev(stream_data) 
               else:
                  external_queue[key][j]["std"] = 0
                      
       for i in self.wifi_devices:
           self.handlers["EXTERNAL_SENSORS"].hset(i,external_queue[i])  
           
       print("external_queue",external_queue)           
       self.external_queue = external_queue
      
   def monitor_internal_queue(self,*params):
      
       data = self.get_internal_streams()
       
       internal_queue = {}
       for key,data_item in data.items():
           fields = self.find_fields(data_item)
           internal_queue[key] = {}
           temp = data[key]
           print("fields",fields)
           for j in fields:
               internal_queue[key][j] = {}
               stream_data = self.analyize_internal_stream(j,temp)
               internal_queue[key][j]["max"] = max(stream_data)
               internal_queue[key][j]["min"] =min(stream_data)                 
               internal_queue[key][j]["mean"] = statistics.mean(stream_data)
               internal_queue[key][j]["median"] = statistics.median(stream_data) 
               if len(stream_data) >1:
                  internal_queue[key][j]["std"] = statistics.stdev(stream_data) 
               else:
                  sinternal_queue[key][j]["std"] = 0
                      
       for i in self.wifi_devices:
           self.handlers["INTERNAL_SENSORS"].hset(i,internal_queue[i])  
       print("internal_queue",internal_queue)           
       self.internal_queue = internal_queue
     
 

   def monitor_alert_queue(self,*params):

       data = self.handlers["ALERT_PUSH"].pop()
      
       if data[0] == False:
          return 
       data = data[1]
      
       if 'NODE_ID' in data:
          node_id = data["NODE_ID"]
          if node_id in self.wifi_devices:
              queue = self.wifi_devices [node_id]["input_queue"]
              #print("queue",queue)
              length = self.wifi_devices[node_id]["input_queue_length"]
              self.wifi_data_handlers.job_queue_push(queue,data,length)
           
          else:
              self.handlers["UNKNOWN_DEVICES"].hset(node_id,node_id)
         
       else:
             print("bad data packet")         
      
                 

   def monitor_well_queue(self,*params):
       pass

   #### Internal functions
   
   def get_reboot_stream(self):
       current_time_stamp = time.time()
       data = self.wifi_data_handlers.stream_range(self.wifi_source["queues"]["reboot"], current_time_stamp-3600*SCAN_INTERVAL, current_time_stamp)
       #print("reboot data",data)
       return data
       
   def get_heart_beat_stream(self):
       current_time_stamp = time.time()
       data = self.wifi_data_handlers.stream_range(self.wifi_source["queues"]["heart_beat"], current_time_stamp-3600*SCAN_INTERVAL, current_time_stamp)
       #print("reboot data",data)
       return data

   def get_internal_streams(self):
       
       current_time_stamp = time.time()
       internal_stream = self.wifi_data_handlers.stream_range(self.wifi_source["queues"]["internal_streams"], current_time_stamp-3600*SCAN_INTERVAL, current_time_stamp)
       #print("internal_stream",internal_stream)   
       internal_queue = {}
       for i in self.wifi_devices:
           internal_queue[i] = []
           
           
       for i in internal_stream:
           time_stamp = i['timestamp']
           data = i["data"]
           node_id = data['Node_Id']
           if node_id in self.wifi_devices:
               data["time_stamp"] = time_stamp
               internal_queue[node_id].append(data)
           else:
              self.handlers["UNKNOWN_DEVICES"].hset(node_id,node_id)
      
       return internal_queue   


   def get_external_streams(self):
       
       current_time_stamp = time.time()
       external_stream = self.wifi_data_handlers.stream_range(self.wifi_source["queues"]["external_streams"], current_time_stamp-3600*SCAN_INTERVAL, current_time_stamp)
       #print("internal_stream",internal_stream)   
       external_queue = {}
       for i in self.wifi_devices:
           external_queue[i] = []
           
           
       for i in external_stream:
           time_stamp = i['timestamp']
           data = i["data"]
           node_id = data['Node_Id']
           if node_id in self.wifi_devices:
               data["time_stamp"] = time_stamp
               external_queue[node_id].append(data)
           else:
              self.handlers["UNKNOWN_DEVICES"].hset(node_id,node_id)
      
       return external_queue    

       
   def sort_streams(self, reboot_streams, heart_beat_stream):
       device_summary = {}
       for i in self.wifi_devices:
          
           device_summary[i] = {}
           device_summary[i]["raw_reboot"] =[]
           device_summary[i]["raw_heart_beat"] = []
           
       for i in reboot_streams:
           time_stamp = i['timestamp']
           data = i["data"]
           node_id = data['NODE_ID']
           if node_id in device_summary:
               data["time_stamp"] = time_stamp
               device_summary[node_id]["raw_reboot"].append(data)
           else:
              self.handlers["UNKNOWN_DEVICES"].hset(node_id,node_id)
       for i in heart_beat_stream:
           time_stamp = i['timestamp']
           data = i["data"]
           node_id = data['NODE_ID']
           if node_id in device_summary:
               data["time_stamp"] = time_stamp
               device_summary[node_id]["raw_heart_beat"].append(data)
           else:
              self.handlers["UNKNOWN_DEVICES"].hset(node_id,node_id)

       return device_summary    


   def check_for_device_status(self,device_status):
        time_stamp = time.time()

    
        for i in self.wifi_devices:
            if len(device_status[i]["raw_heart_beat"]) <1:
               device_status[i]["active"] = False
            else:
               last_hb = device_status[i]["raw_heart_beat"][-1]
               last_hb = last_hb["time_stamp"]
            
               if last_hb + 60 *3 > time_stamp:
                    device_status[i]["active"] = True
               else:
                     device_status[i]["active"] = False
        
        return device_status        
        statistics.mean(data1) 
        
   def analyize_reboot_times(self,device_status):
       for i in self.wifi_devices:
           reboot_data = device_status[i]["raw_reboot"]
           reboot_number = len(reboot_data)
           device_status[i]["reboot_number"] = reboot_number
           if reboot_number > 0:
              device_status[i]["latest_reboot"] =    device_status[i]["raw_reboot"][-1]["time_stamp"]
           else:
               device_status[i]["latest_reboot"] = 0           
           reboot_times = []
           if reboot_number > 2:
              for j in range(0,reboot_number-1):
                  delta_t = device_status[i]["raw_reboot"][j+1]["time_stamp"] - device_status[i]["raw_reboot"][j]["time_stamp"]
                  reboot_times.append(delta_t)
              device_status[i]["reboot_mean"] = statistics.mean(reboot_times) 
              device_status[i]["reboot_std"] = statistics.stdev(reboot_times) 
              device_status[i]["reboot_median"] = statistics.median(reboot_times) 
              device_status[i]["reboot_min"]           = min(reboot_times)
              device_status[i]["reboot_max"]           = max(reboot_times)
           else:
              device_status[i]["reboot_mean"] = 0
              device_status[i]["reboot_std"] = 0
              device_status[i]["reboot_median"] = 0
              device_status[i]["reboot_min"]           = 0
              device_status[i]["reboot_max"]           = 0
           
         
       return device_status
       
       
       
   def analyize_heap_values(self,device_status):
       
       for i in self.wifi_devices:
           heart_beat_values = []
           raw_heart_beat = device_status[i]["raw_heart_beat"]
           for j in raw_heart_beat:
               heart_beat_values.append(j['HEAP'])
           if len(heart_beat_values) > 0:
               device_status[i]["heap_mean"] = statistics.mean(heart_beat_values)
               device_status[i]["heap_min"] =  min(heart_beat_values)
               device_status[i]["heap_max"] = max(heart_beat_values)
           else:
               device_status[i]["heap_mean"] = 0
               device_status[i]["heap_min"] =  0
               device_status[i]["heap_max"] = 0
       
       return device_status   


   def analyize_internal_stream(self,field,data):
       return_value = []
       for i in data:
          if field in i:
             return_value.append(float(i[field]))
       return return_value

   def log_device_status(self,*params):
        
         self.handlers["DEVICE_STATUS_LOG"].push(self.device_status) 
         
   def log_internal_queue(self,*params):
       self.handlers["INTERNAL_SENSORS_LOG"].push(self.internal_queue) 
       
   def log_external_queue(self,*params):
       self.handlers["EXTERNAL_SENSORS_LOG"].push(self.external_queue) 
       


   def find_fields(self,data):
       #print("data",data)
       return_value = set()
       for i in data:
          for field,data in i.items():
              if (field != "time_stamp") and (field != "Node_Id") :
                   return_value.add(field)
       
       return list(return_value)
       
 
def construct_wifi_instance( qs, site_data ):

                   
    
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
    query_list = qs.add_match_terminal( query_list, relationship="WIFI_DEVICES" )
    wifi_sets, wifi_sources = qs.match_list(query_list)
    wifi_source = wifi_sources[0]
    
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
    query_list = qs.add_match_relationship( query_list, relationship="WIFI_DEVICES" )
    query_list = qs.add_match_terminal( query_list, relationship="PACKAGE" )
    package_sets, package_sources = qs.match_list(query_list)  
    package = package_sources[0]
    data_structures = package["data_structures"]
    #print("data_structurres",data_structures);
   
    generate_handlers = Generate_Handlers( package, qs )
    '''
    handlers = {}
    handlers["DEVICE_STATUS"] = generate_handlers.construct_hash(data_structures["DEVICE_STATUS"])
    handlers["UNKNOWN_DEVICES"] = generate_handlers.construct_hash(data_structures["UNKNOWN_DEVICES"])
    handlers["ALERT_PUSH"] = generate_handlers.construct_job_queue_server(data_structures["ALERT_PUSH"])
    handlers["INTERNAL_SENSORS"] = generate_handlers.construct_hash(data_structures["INTERNAL_SENSORS"])
    handlers["WELL_SENSORS"] = generate_handlers.construct_hash(data_structures["WELL_SENSORS"])
    handlers["TURN_ON_WATER"] = generate_handlers.construct_job_queue_client(data_structures["TURN_ON_WATER"])
    handlers["DEVICE_STATUS_LOG"] = generate_handlers.construct_redis_stream_writer(data_structures["DEVICE_STATUS_LOG"])
    handlers["INTERNAL_SENSORS_LOG"] = generate_handlers.construct_redis_stream_writer(data_structures["INTERNAL_SENSORS_LOG"])
    '''
    handlers = generate_handlers.generate_all()
   

    #handlers["DEVICE_STATUS_LOG"].delete_all()
    #handlers["INTERNAL_SENSORS_LOG"].delete_all() 
    
   

    
    wifi_devices = {}
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
    query_list = qs.add_match_relationship( query_list, relationship="WIFI_DEVICES" )
    query_list = qs.add_match_terminal( query_list, relationship="WIFI_DEVICE" )
    
    devices_sets, device_sources = qs.match_list(query_list)  
    devices = device_sources
    for i in devices:
       wifi_devices[i["node_id"]] = i
    #print(wifi_devices)
    #print(len(wifi_devices))
    
    return  WIFI_Devices(wifi_source,handlers,wifi_devices)

 
    
    
    
    
    
    '''

    return redis_monitor
    '''



def add_chains(wifi_devices, cf):
    cf.define_chain("monitor_input_queue", True)
    #cf.insert.log("monitor_input_queue")
    cf.insert.one_step(wifi_devices.monitor_device_input_queue)
    cf.insert.wait_event_count( event = "TIME_TICK",count = 1)
    cf.insert.reset()

    cf.define_chain("monitor_alert_queue", True)
    #cf.insert.log("monitor_alert_queue")
    cf.insert.one_step(wifi_devices.monitor_alert_queue)
    cf.insert.wait_event_count( event = "TIME_TICK",count = 1)
    cf.insert.reset()


    cf.define_chain("monitor_device_status", True)
    #cf.insert.log("monitor_device_status")
    cf.insert.one_step(wifi_devices.monitior_device_status)
    cf.insert.wait_event_count( event = "MINUTE_TICK",count = 5)
    cf.insert.reset()

    cf.define_chain("monitor_internal_queue", True)
    
    #cf.insert.log("monitor_queue")
    cf.insert.one_step(wifi_devices.monitor_internal_queue)
    cf.insert.one_step(wifi_devices.monitor_external_queue)
    cf.insert.wait_event_count( event = "MINUTE_TICK",count = 5)
    cf.insert.reset()
    
    cf.define_chain("log_data", True)
   
    #cf.insert.log("log_data")
    cf.insert.one_step(wifi_devices.log_device_status)
    cf.insert.one_step(wifi_devices.log_internal_queue)
    cf.insert.one_step(wifi_devices.log_external_queue)
    cf.insert.wait_event_count( event = "MINUTE_TICK",count = 15)
    cf.insert.reset()


    cf.define_chain("monitor_well_monitor", True)
    #cf.insert.log("monitor_well_monitor")
    cf.insert.one_step(wifi_devices.monitor_well_queue)
    cf.insert.wait_event_count( event = "MINUTE_TICK",count = 5)
    cf.insert.reset()


    
    cf.define_chain("trim_streams", True)
    #cf.insert.log("trim_streams")
    cf.insert.one_step(wifi_devices.trim_device_data_structures)
    cf.insert.wait_event_count( event = "MINUTE_TICK",count = 15)
    cf.insert.reset()

 

if __name__ == "__main__":

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
    #import load_files_py3
    from redis_support_py3.graph_query_support_py3 import  Query_Support
    import datetime
    #from redis_support_py3.user_data_tables_py3 import User_Data_Tables

    from py_cf_new_py3.chain_flow_py3 import CF_Base_Interpreter

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
    
    wifi_monitor = construct_wifi_instance(qs, redis_site )
   
    
    
    #
    # Adding chains
    #
    cf = CF_Base_Interpreter()
    add_chains(wifi_monitor, cf)
    #
    # Executing chains
    #
    
    cf.execute()