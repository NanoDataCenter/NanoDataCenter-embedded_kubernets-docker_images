#
#
#  File: flask_web_py3.py
#
#
#
import os
import json
import redis
import urllib
from flask import render_template,jsonify
from flask import request, session, url_for


from  mqtt_client_monitor.load_mqtt_client_monitor_py3 import Load_MQTT_Client_Monitoring



class PI_MQTT_Client_Monitor(object):

   def __init__(base_self ,self  ):
       
 
       Load_MQTT_Client_Monitoring(self.app, self.auth, request, render_template, self.qs,
                                         self.site_data,self.url_rule_class,"MQTT_CLIENT_MONITOR",'mqtt_client_monitor')

 
   
   

if __name__ == "__main__":

   file_handle = open("/data/redis_server.json",'r')
   data = file_handle.read()
   file_handle.close()
  
   redis_site_data = json.loads(data)


   pi_web_server = PI_MQTT_Client_Monitor(__name__, redis_site_data  )
   pi_web_server.generate_menu_page()
   #pi_web_server.generate_index_page("Pod_Control_Processes","start_and_stop_processes")
   
   pi_web_server.generate_site_map()
   
   pi_web_server.generate_default_index_page()
   pi_web_server.run_http()
   
   
