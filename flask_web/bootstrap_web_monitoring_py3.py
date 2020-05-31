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

from web_monitor.pod_process_control.load_pod_control_processes_py3 import Load_Pod_Control_Processes
from web_monitor.processor_performance.processor_performance_py3 import Processor_Monitoring
from web_monitor.redis_monitor.load_redis_management_py3 import Load_Redis_Monitoring
from web_monitor.docker_process_control.load_docker_process_control_py3 import Load_Docker_Processes
from web_monitor.subsystem_monitor.load_subsystem_monitor_py3 import Load_Subsystem_Monitor
from bootstrap_web_core_py3 import PI_Web_Server_Core


class PI_Web_Monitor_Server(PI_Web_Server_Core):

   def __init__(self , name, site_data ):
       PI_Web_Server_Core.__init__(self,name,site_data)


       Load_Pod_Control_Processes(self.app, self.auth, request, render_template, self.qs,
                                         self.site_data,self.url_rule_class,"Pod_Control_Processes",'web_monitor/pod_process_control')
                                         
       Load_Docker_Processes(self.app, self.auth, request, render_template, self.qs,
                                         self.site_data,self.url_rule_class,"Docker_Control",'web_monitor/docker_process_control') 
                                         
                                         
       Processor_Monitoring(self.app, self.auth, request, render_template, self.qs,
                                         self.site_data,self.url_rule_class,"Controller_Resoure_Utiliation",'web_monitor/controller_monitor')      
       
       Load_Redis_Monitoring(self.app, self.auth, request, render_template, self.qs,
                                         self.site_data,self.url_rule_class,"Redis_Monitor",'web_monitor/redis_monitor') 

                                         
       Load_Subsystem_Monitor(self.app, self.auth, request, render_template, self.qs,
                                         self.site_data,self.url_rule_class,"Subsystem_Monitor",'web_monitor/subsystem_monitor') 
 

 
   
   

if __name__ == "__main__":

   file_handle = open("/data/redis_server.json",'r')
   data = file_handle.read()
   file_handle.close()
  
   redis_site_data = json.loads(data)


   pi_web_server = PI_Web_Monitor_Server(__name__, redis_site_data  )
   pi_web_server.generate_menu_page()
   #pi_web_server.generate_index_page("Pod_Control_Processes","start_and_stop_processes")
   
   pi_web_server.generate_site_map()
   
   pi_web_server.generate_default_index_page()
   pi_web_server.run_http()
   
   
