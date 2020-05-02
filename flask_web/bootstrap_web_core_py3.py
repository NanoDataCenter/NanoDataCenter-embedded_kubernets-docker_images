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

import flask
from flask import Flask
from flask import render_template,jsonify
from flask_httpauth import HTTPDigestAuth
from flask import request, session

from redis_support_py3.graph_query_support_py3 import  Query_Support
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
from web_packages.load_static_pages_py3     import  Load_Static_Files 
from web_packages.load_redis_access_py3     import  Load_Redis_Access
from web_packages.load_pod_control_processes_py3 import Load_Pod_Control_Processes
from redis_support_py3.construct_data_handlers_py3 import Redis_RPC_Client

class PI_Web_Server_Core(object):

   def __init__(self , name, site_data ):
       redis_handle_pw = redis.StrictRedis(site_data["host"], 
                                           site_data["port"], 
                                           db=site_data["redis_password_db"], 
                                           decode_responses=True)
                               

       self.site_data = site_data                                       
       startup_dict = redis_handle_pw.hgetall("web")
       
       
       
       self.qs = Query_Support( redis_site_data )
       
       self.app         = Flask(name) 
       self.auth = HTTPDigestAuth()
       self.auth.get_password( self.get_pw )
       
       self.startup_dict = startup_dict
      
      
       self.app.template_folder       =   'web_packages/templates'
       self.app.static_folder         =   'web_packages/static'  
       self.app.config['SECRET_KEY']      = startup_dict["SECRET_KEY"]

 
       self.users                    = json.loads(startup_dict["users"])
       
       
       
       Load_Static_Files(self.app,self.auth) #enable static files to be fetched
       self.redis_access = Load_Redis_Access(self.app, self.auth, request ) #enable web access for redis operations
      
       self.subsystems = []
       self.modules = {}
       
       subsystem_name = "Pod_Control_Processes"
     
       self.modules[subsystem_name] = Load_Pod_Control_Processes(self.app, self.auth, request, render_template, self.qs,self.site_data,"/"+subsystem_name)
       self.subsystems.append(subsystem_name)
       
       #self.menu.append(self.load_node_process_interface())
       #self.menu.append(self.load_docker_interface())
       #self.menu.append(self.load_processor_utilization_interface())
       #self.menu_append(self.self.load_redis_monitoring())
 
 
   def get_pw( self,username):
       
       print(username)
       if username in self.users:
           print("sucess")
           return self.users[username]
       return None
 
   def generate_menu_page(self):
      
       self.subsystems.sort()
       self.generate_menu_template(self.subsystems)
       self.generate_modal_template(self.subsystems,self.modules)
       
   def generate_index_page(self):
       pass
       
   def generate_status_function(self):
       pass

   def run_http( self):
       self.app.run(threaded=True , use_reloader=True, host='0.0.0.0',port=80,debug = True)

   def run_https( self ):
       startup_dict          = self.startup_dict
      
       self.app.run(threaded=True , use_reloader=True, host='0.0.0.0',debug = True,
           port=443 ,ssl_context=("/data/cert.pem", "/data/key.pem"))
       
 

   
   def generate_menu_template(self,list_of_subsystems):
       f = open('web_packages/templates/menu', 'w')

       output_string = '''
       <nav class="navbar navbar-expand-sm bg-dark navbar-dark">
       <!-- Links -->
       <ul class="navbar-nav">
       <!-- Dropdown -->
       <li class="nav-item dropdown">
       <a class="nav-link dropdown-toggle" href="#" id="navbardrop" data-toggle="dropdown">Menu</a>
       <div class="dropdown-menu">  
       '''
       f.write(output_string)
       for i in list_of_subsystems:
          temp =  '    <a class="dropdown-item" href="#"  data-toggle="modal" data-target="#'+i+'">'+i+"</a>\n"
          f.write(temp)
       output_string = '''        
        
                 </div>
                 </li>
                </ul>
                <ul class="navbar-nav">

               <button id="status_panel", class="btn " type="submit">Status</button>
                </ul>
               <nav class="navbar navbar-light bg-dark navbar-dark">
               <span class="navbar-text" >
               <h4 id ="status_display"> Status: </h4>
              </span>
            </nav>
       '''
       f.write(output_string)
       f.close()
       
       
       
       
   def generate_modal_template(self,subsystems,modules):
   
       f = open('web_packages/templates/modals', 'w')
       for i in subsystems:
           rule = subsystems[i].rule_data
           
           
           [sub_system_name+'/start_and_stop_processes'] = [a1,sub_system_name+'/start_and_stop_processes/0' ]
           output_string = '<!–'+i+' –>\n'
           f.write(output_string)
           output_string = '''
 
           <div class="modal fade" id="field_operation" tabindex="-1" role="dialog" aria-labelledby="accountModalLabel" aria-hidden="true">
           <div class="modal-dialog" role="document">
           <div class="modal-content">
           <div class="modal-header">
           <h5 class="modal-title" id="accountModalLabel">Field Operations</h5>
           <button type="button" class="close" data-dismiss="modal" aria-label="close">
                    <span aria-hidden="true">&times;</span>
           </button>
                 
           </div>
           <div class="modal-body">
               <ul >
           '''       
           f.write(output_string)
           for i in           
                        
                   <li><a href="/diagnostics/schedule_control" target="_self">Turn On/Off By Schedule Step</a></li>
                   <li><a href="/diagnostics/controller_pin" target="_self">Turn On/Off By Controller Pin</a></li>
                   <li><a href="/diagnostics/valve_group" target="_self">Turn On/Off By Valve Group</a></li>
        
           output_string = '''
                </ul>
            </div>
            <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
                    
            </div>
            </div>
           </div>
           </div>
           '''
           f.write(output_string)
       f.close()           
     
if __name__ == "__main__":

   file_handle = open("/data/redis_server.json",'r')
   data = file_handle.read()
   file_handle.close()
  
   redis_site_data = json.loads(data)


   pi_web_server = PI_Web_Server_Core(__name__, redis_site_data  )
   pi_web_server.generate_menu_page()
   pi_web_server.generate_index_page()
   
   pi_web_server.generate_status_function()
   pi_web_server.run_http()
   
   
