import paho.mqtt.client as mqtt
import ssl
from redis_support_py3.graph_query_support_py3 import  Query_Support
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
import time
import msgpack
class MQTT_Publish(object):

   def __init__(self,redis_site) :
       
       
       qs = Query_Support( redis_site )
       query_list = []
       query_list = qs.add_match_relationship( query_list,relationship="SITE",label=redis_site["site"] )

       query_list = qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", property_mask={"name":"MQTT_DEVICES_DATA"} )
                                           
       package_sets, package_sources = qs.match_list(query_list)
       package = package_sources[0] 
       data_structures = package["data_structures"]
       generate_handlers = Generate_Handlers(package,qs)
       self.job_queue_server = generate_handlers.construct_job_queue_server(data_structures["MQTT_PUBLISH_QUEUE"])
       query_list = []
       query_list = qs.add_match_relationship( query_list,relationship="SITE",label=redis_site["site"] )
       query_list = qs.add_match_terminal( query_list, 
                             relationship = "MQTT_SERVER" )
                                           
       host_sets, host_sources = qs.match_list(query_list)
       self.mqtt_server_data = host_sources[0]

       
       self.client = mqtt.Client(client_id="", clean_session=True, userdata=None,  transport="tcp")
 
       self.client.on_connect = self.on_connect
       
       self.client.on_disconnect = self.on_disconnect
       self.client.on_publish = self.on_publish
       self.connection_flag = False
       print("connection attempting")
       while self.connection_flag == False:
           try:
                self.client.connect(self.mqtt_server_data["HOST"],self.mqtt_server_data["PORT"], 60)
           except:
                
                time.sleep(5)
           else:
              self.connection_flag = True
       print("connection achieved")
       self.client.loop_start()
       self.server_job_queue()
       
   def on_connect(self,client, userdata, flags, rc):
       if rc != 0:
          time.sleep()
          raise ValueError("Bad connection")
       print("Connected with result code "+str(rc),)
      

   def on_disconnect(self, client, userdata, rc):
          exit(-1) # self.client.loop_start automatically reconnects
          

   def on_publish(self,client,userdata,result):
      
       self.publish_flag = True

       
   def server_job_queue(self):
       try:
         while True:
           if self.job_queue_server.length() > 0:
               data = self.job_queue_server.show_next_job()
               data = data[1]
               topic = data["tx_topic"]
               del data["tx_topic"] 
               binary_data = msgpack.packb(data, use_bin_type=True)
               print("topic",topic,data)
               self.publish_flag = False
               self.client.publish(topic,binary_data)
               self.job_queue_server.pop()
           else:
               time.sleep(.5)
       except Exception as tst:
          self.client.loop_stop()
          
          raise ValueError(tst)
             

     



 
 
 
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
    MQTT_Publish(redis_site)
 
 