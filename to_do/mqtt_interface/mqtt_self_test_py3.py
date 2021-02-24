import paho.mqtt.client as mqtt
import ssl
from redis_support_py3.graph_query_support_py3 import  Query_Support
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
import time
import msgpack
class MQTT_Server_Test(object):

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
       self.job_queue_client = generate_handlers.construct_job_queue_client(data_structures["MQTT_PUBLISH_QUEUE"])
       self.send_request("REBOOT")
       while 1:
           self.send_request("HEART_BEAT")
           self.send_request("SERVER_CHECK")
           time.sleep(15.)
 
   def send_request(self,topic):
       msg_dict = {}
       msg_dict["tx_topic"] = "MQTT_SERVER_CHECK"+"/"+topic
       self.job_queue_client.push(msg_dict)
 
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
    MQTT_Server_Test(redis_site)
 
 






