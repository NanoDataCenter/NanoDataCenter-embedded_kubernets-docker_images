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

from sqlite_library.sqlite_sql_support_py3 import SQLITE_Client_Support
from bootstrap_web_monitoring_py3 import PI_Web_Monitor_Server
from web_system_control.load_ticket_control_py3 import Load_Ticket_Control

class PI_Web_System_Control(PI_Web_Monitor_Server):

   def __init__(self , name, site_data ):
       PI_Web_Monitor_Server.__init__(self,name,site_data)
       self.sqlite_library = SQLITE_Client_Support(self.qs, site_data)
       Load_Ticket_Control( self.app, self.auth, request, render_template,self.qs,site_data,self.url_rule_class,'web_site_control/ticket_control',self.sqlite_library)

 
   
   

if __name__ == "__main__":

   file_handle = open("/data/redis_server.json",'r')
   data = file_handle.read()
   file_handle.close()
  
   redis_site_data = json.loads(data)


   pi_web_server = PI_Web_System_Control(__name__, redis_site_data  )
   pi_web_server.generate_menu_page()
   #pi_web_server.generate_index_page("Pod_Control_Processes","start_and_stop_processes")
   
   pi_web_server.generate_site_map()
   
   pi_web_server.generate_default_index_page()
   pi_web_server.run_http()
   
   
