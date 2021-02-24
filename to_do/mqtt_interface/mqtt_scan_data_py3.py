#
#  This file will check for heart beat messages and reboot messages
#
#
#
#
#
#




import msgpack
import base64
import redis
import time
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
from mqtt_library.mqtt_to_redis_py3 import MQTT_TO_REDIS_BRIDGE_RETRIEVE


class MQTT_Log(object):

   def __init__(self,
                mqtt_server_data, 
                mqtt_devices, 
                package,
                site_data,
                qs):
               
               
        
        
        self.mqtt_server_data  = mqtt_server_data
        self.mqtt_devices = mqtt_devices
        self.site_data = site_data
        
        

        self.mqtt_bridge = MQTT_TO_REDIS_BRIDGE_RETRIEVE(site_data)
        self.construct_presence_key_name()
        self.construct_reboot_key_name()
        self.generate_data_handlers(package,qs)
        

        self.log_data()
        
        
   def construct_presence_key_name(self):
       for i,item in self.mqtt_devices.items():
       
          topic = item["name"]+"/"+item["HEART_BEAT"]
          self.mqtt_devices[i]["presence_name_space"] = self.mqtt_bridge.construct_name_space(topic)[0]       
       
   def construct_reboot_key_name(self):
       for i,item in self.mqtt_devices.items():
         if "REBOOT_FLAG" in item:
            topic = item["name"]+"/"+item["REBOOT_KEY"]
            self.mqtt_devices[i]["reboot_name_space"] = self.mqtt_bridge.construct_name_space(topic)[0]   
      
        
   def generate_data_handlers(self,package,qs):
        self.handlers = {}
        data_structures = package["data_structures"]
        generate_handlers = Generate_Handlers(package,qs)
        self.ds_handlers = {}
       
        self.ds_handlers["MQTT_PAST_ACTION_QUEUE"] = generate_handlers.construct_redis_stream_writer(data_structures["MQTT_PAST_ACTION_QUEUE"])
        self.ds_handlers["MQTT_CONTACT_LOG"] = generate_handlers.construct_hash(data_structures["MQTT_CONTACT_LOG"])
        self.ds_handlers["MQTT_REBOOT_LOG"] = generate_handlers.construct_hash(data_structures["MQTT_REBOOT_LOG"])
        
        contact_set = set(self.ds_handlers["MQTT_CONTACT_LOG"].hkeys())
        device_set = set(self.mqtt_devices.keys())
        difference_set = contact_set-device_set
        for i in list(difference_set):
           self.ds_handlers["MQTT_CONTACT_LOG"].hdelete(i)
           
        




   def log_data(self):
       
       self.check_heartbeat()
       self.check_reboot()   
      
       
       
       while 1:
           time.sleep(60)
               
              
           self.check_heartbeat()
           self.check_reboot()
           


 
   

   def update_contact(self,name,key,status):
       print("*******",name,key,status)
       old_data = self.ds_handlers["MQTT_CONTACT_LOG"].hget(name)
       print("old_data",old_data)
       
       data = {}
       data["time"] = time.time()
       data["status"] = status
       data["name"] = name
       data["device_id"] = name # redundant with name
       self.ds_handlers["MQTT_CONTACT_LOG"].hset(name,data)  
       update_flag = False
       if old_data == None:
          update_flag = False # already made update
          self.ds_handlers["MQTT_PAST_ACTION_QUEUE"].push({"action":"Device_Change","device_id":name,"status":status})
         
          return
       if old_data["status"] != status:
           update_flag = True
           print("device change")
          
           self.ds_handlers["MQTT_PAST_ACTION_QUEUE"].push({"action":"Device_Change","device_id":name,"status":status})   
       
       
          

   def check_heartbeat(self):
      
      for i,items in self.mqtt_devices.items():
          name = items["name"]
          key = items["presence_name_space"]
          
          if self.mqtt_bridge.stream_exists(key) == False:
            print("stream does not exist")
            self.update_contact(name,key,False)
          else:
            data = self.mqtt_bridge.xrevrange_namespace(key, "+", "-" , count=1)
            print('data',data[0])
            timestamp = data[0]["timestamp"]
            print(name,time.time()-timestamp)
            if items["HEART_BEAT_TIME_OUT"] < time.time()-timestamp:
              
              self.update_contact(name,key,False)
            else:
              
              self.update_contact(name,key,True)

   def update_reboot(self,name,key,timestamp):
       old_data = self.ds_handlers["MQTT_REBOOT_LOG"].hget(name)
       data = {}
       data["timestamp"] = timestamp
       data["name"] = name
       data["device_id"] = name # redundant with name
       update_flag = False
       if (old_data == None) or (old_data["timestamp"] < timestamp):
           print("update reboot log")
           self.ds_handlers["MQTT_REBOOT_LOG"].hset(name,data)  
           self.ds_handlers["MQTT_PAST_ACTION_QUEUE"].push({"action":"Device_Reboot","device_id":name,"status":True})   




      
   def check_reboot(self):
     
      for i,items in self.mqtt_devices.items():
          name = items["name"]
          if "REBOOT_FLAG" in items:
             key = items["reboot_name_space"] 
            
             if self.mqtt_bridge.stream_exists(key) == False:
                pass
                
             else:
                 data = self.mqtt_bridge.xrevrange_namespace(key, "+", "-" , count=1)
                 data = data[0]
                 self.update_reboot(name,key,data["timestamp"])
      


                   
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
     
    qs = Query_Support( redis_site )
 

    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=redis_site["site"] )
    query_list = qs.add_match_terminal( query_list, 
                             relationship = "MQTT_SERVER" )
                                           
    host_sets, host_sources = qs.match_list(query_list)
    host_data = host_sources[0]
    
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=redis_site["site"] )
    query_list = qs.add_match_terminal( query_list, 
                                        relationship =  "MQTT_DEVICE" )
                                        
    mqtt_sets, mqtt_sources = qs.match_list(query_list) 
    mqtt_devices = {}
 
    for i in mqtt_sources:
       mqtt_devices[i["name"]] = i

 
    
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=redis_site["site"] )

    query_list = qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", property_mask={"name":"MQTT_DEVICES_DATA"} )
                                           
    package_sets, package_sources = qs.match_list(query_list)
    package = package_sources[0]
 
    MQTT_Log(mqtt_server_data = host_data,
                 mqtt_devices = mqtt_devices,
                 package = package,
                 site_data = redis_site,
                 qs = qs)
                 
               
                 
                

   
