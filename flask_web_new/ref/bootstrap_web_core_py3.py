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
from flask import request, session, url_for

from redis_support_py3.graph_query_support_py3 import  Query_Support
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
from web_core.load_static_pages_py3     import  Load_Static_Files 
from web_core.load_redis_access_py3     import  Load_Redis_Access

from redis_support_py3.construct_data_handlers_py3 import Redis_RPC_Client

from bootstrap_web_system_control_py3 import PI_Web_System_Control
from bootstrap_web_monitoring_py3     import PI_Web_Monitor_Server
from bootstrap_mqtt_client_py3        import PI_MQTT_Client_Monitor
from bootstrap_eto_py3                import ETO_Management
from file_server_library.file_server_lib_py3 import Construct_RPC_Library
from bootstrap_irrigation_scheduling_py3 import Irrigation_Scheduling
from irrigation_control.load_irrigation_control_py3   import Load_Irrigation_Control
class URL_Rule_Class(object):

   def __init__(self,app,auth):
       self.subsystems = {}
       self.subsystem_order = []
       self.app = app
       self.auth = auth


   def add_get_rules(self,subsystem_name,function_list,url_list):
       slash_name = "/"+subsystem_name+"/"
       assert(len(function_list)==len(url_list))
       menu_list = []
       menu_data = {}
       for i in range(0,len(function_list)):
           a1 = self.auth.login_required( function_list[i] )
           self.app.add_url_rule(slash_name+url_list[i][0]+url_list[i][1],slash_name+url_list[i][0],a1)
           menu_data[url_list[i][0]] =[a1,url_list[i][0]+url_list[i][2],url_list[i][3]] 
           menu_list.append(url_list[i][0])
       self.subsystems[subsystem_name] = {"menu_list":menu_list,"menu_data":menu_data}      
       self.subsystem_order.append(subsystem_name)
       
   def move_directories(self,path):
       #print("move directory path",path)
       path_test = path.split("/")
       if len(path_test) != 1:
          path_dest = path_test[1]
       else:
          path_dest = path
       #print(path)
       #print(path_dest)
       
       os.system('mkdir flask_templates/'+path_dest)
       os.system('mkdir flask_templates/js/'+path_dest)
       #os.system("ls flask_templates")
       #print("path",path,path_dest)
       os.system('cp -r ' +path+'/templates/* flask_templates/'+path_dest)
       
       os.system('cp -r ' +path+'/js/* flask_templates/js/'+path_dest)
       return path_dest




class Load_App_Sys_Files(object):

   def __init__( self, app, auth, request, file_server_library ):
       self.app      = app
       self.auth     = auth
       self.request  = request
      
       self.file_server_library = file_server_library

       a1 = auth.login_required( self.get_system_file )
       app.add_url_rule("/ajax/get_system_file/<path:file_name>","get_system_file",a1)
       a1 = auth.login_required( self.get_app_file )
       app.add_url_rule("/ajax/get_app_file/<path:file_name>","get_app_file",a1)
       a1 = auth.login_required( self.save_app_file )
       app.add_url_rule("/ajax/save_app_file/<path:file_name>","save_app_file",a1,methods=["POST"])
       a1 = auth.login_required( self.save_sys_file )
       app.add_url_rule("/ajax/save_sys_file/<path:file_name>","save_sys_file",a1,methods=["POST"])
               


   def get_system_file(self, file_name):   
       data = self.file_server_library.load_file( "application_files",file_name)
      
       return json.dumps(data)

   def get_app_file(self,file_name):
       data = self.file_server_library.load_file( "system_files",file_name)
       return json.dumps(data )
               
   def save_app_file(self,file_name):
       json_object = self.request.json
      
       if type(json_object) != str:
          json_object = json.dumps(json_object)
       self.file_server_library.save_file("application_files",file_name, json_object );
       return json.dumps('SUCCESS')

   def save_sys_file(self,file_name):
       json_object = self.request.json
       if type(json_object) != str:
          json_object = json.dumps(json_object)
       self.file_server_library.save_file( "system_files",file_name, json_object );
       return json.dumps('SUCCESS') 
       




class PI_Web_Server_Core(object):

   def __init__(self , name, site_data ):
       redis_handle_pw = redis.StrictRedis(site_data["host"], 
                                           site_data["port"], 
                                           db=site_data["redis_password_db"], 
                                           decode_responses=True)
                               


       
       
       self.site_data = site_data                                       
       startup_dict = redis_handle_pw.hgetall("web")
       
       
       
       self.qs = Query_Support( site_data)
       self.file_server_library = Construct_RPC_Library(self.qs,self.site_data)
       
       self.app         = Flask(name) 
       self.auth = HTTPDigestAuth()
       self.url_rule_class = URL_Rule_Class(self.app,self.auth)
       
       self.auth.get_password( self.get_pw )
       
       self.startup_dict = startup_dict
      
      
       self.app.template_folder       =   'flask_templates'
       self.app.static_folder         =   'static'  
       self.app.config['SECRET_KEY']      = startup_dict["SECRET_KEY"]

 
       self.users                    = json.loads(startup_dict["users"])
       
       
       
       Load_Static_Files(self.app,self.auth) #enable static files to be fetched
       self.redis_access = Load_Redis_Access(self.app, self.auth, request ) #enable web access for redis operations
       Load_App_Sys_Files( self.app, self.auth, request, self.file_server_library    )
       self.subsystems = []
       self.modules = {}
       self.load_specified_modules()
       
      
   def load_specified_modules(self):
       results=self.common_qs_search(["WEB_SERVER","WEB_SERVER"])
       result = results[0]
       
       modules = result["modules"]
       for i in modules:
           if i == "monitoring":
               print(i)
               PI_Web_Monitor_Server(self)
           elif i == "system_control":
                print(i)
                PI_Web_System_Control(self)
                
           elif i == "mqtt_client":
               print(i)
               PI_MQTT_Client_Monitor(self )

           elif i == "eto":
               print(i)
               ETO_Management(self)
               
           elif i == "irrigation_scheduling":
               print(i)
               Irrigation_Scheduling(self)
               
           elif i == "irrigation_control":
               print(i)
               Load_Irrigation_Control(self)
               
           elif i == "modbus_control":
               print("do nothing right now")
              
 

               
           else:
             raise ValueError("bad web module")         
       self.result = result 
       if "status_function" in self.result:
          print(self.result["status_function"])
       else:
          self.result["status_function"] = ""
          print("status function not defined")
       file_handle = open("flask_templates/js/status_definition.js","w")   
       
       file_handle.write('__status_option__ = "'+self.result["status_function"]+'"; \n') 
       
       file_handle.close()
       

   def common_qs_search(self,search_list): # generalized graph search
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
       for i in range(0,len(search_list)-1):
           if type(search_list[i]) == list:
               query_list = self.qs.add_match_relationship( query_list,relationship = search_list[i][0],label = search_list[i][1] )
           else:
               query_list = self.qs.add_match_relationship( query_list,relationship = search_list[i] )
           
       if type(search_list[-1]) == list:
          query_list = self.qs.add_match_terminal( query_list,relationship = search_list[-1][0],label = search_list[-1][1] )
       else:
           query_list = self.qs.add_match_terminal( query_list,relationship = search_list[-1] )
       
       node_sets, node_sources = self.qs.match_list(query_list)        
       return node_sources
                                         
 
   def get_pw( self,username):
       
      
       if username in self.users:
          
           return self.users[username]
       return None
 
   def generate_menu_page(self):
      
       self.subsystems.sort()
       self.generate_menu_template()
       self.generate_modal_template()
       
       
   def generate_default_index_page(self):
       self.app.add_url_rule("/","home_page",self.links_a1)
       
   def generate_index_page(self,module,element):
       menu_data = self.url_rule_class.subsystems[module]["menu_data"]
   
       menu_element = menu_data[element]
       self.app.add_url_rule("/","home page",menu_element[0])
       
   def generate_site_map(self):
       self.links_a1 = self.auth.login_required( self.site_map_function )
       self.app.add_url_rule("/link_page","/links_page",self.links_a1)
   
   
   def site_map_function(self):   
   
       links = []
       for rule in self.app.url_map.iter_rules():
           # Filter out rules we can't navigate to in a browser
           # and rules that require parameters
           
           #url = url_for(rule.endpoint, **(rule.defaults or {}))
           links.append((rule.endpoint))
           links.sort()
       return render_template("list_of_endpoints",endpoints = links)
         

   def run_http( self):
       self.app.run(threaded=True , use_reloader=True, host='0.0.0.0',port=self.port,debug =self.debug )

   def run_https( self ):
       startup_dict          = self.startup_dict
      
       self.app.run(threaded=True , use_reloader=True, host='0.0.0.0',debug =self.debug,
           port=self.port ,ssl_context=("/data/cert.pem", "/data/key.pem"))
       
 

   
   def generate_menu_template(self):
       f = open( self.app.template_folder+'/menu', 'w')

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
       self.url_rule_class.subsystems 
       for i in self.url_rule_class.subsystems:
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
            </nav>
       '''
       f.write(output_string)
       f.close()
       
       

  
 
   def generate_modal_template(self):
      
       f = open(self.app.template_folder+'/modals', 'w')
       for i in self.url_rule_class.subsystem_order:
           #print("generate_modal_template - i",i)
           output_string = '<!–'+i+' –>\n'
           f.write(output_string)
           
 
           output_string ='<div class="modal fade" id='+i+' tabindex="-1" role="dialog" aria-labelledby="accountModalLabel" aria-hidden="true">\n'
           f.write(output_string)
           
           output_string = '''
        <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
            '''
           f.write(output_string)
            
           f.write('    <h5 class="modal-title" id="accountModalLabel">'+i+'</h5>\n')
           output_string = '''
                <button type="button" class="close" data-dismiss="modal" aria-label="close">
                    <span aria-hidden="true">&times;</span>
                </button>
                 
            </div>
            <div class="modal-body">
                <ul >
                
           '''           
           
           f.write(output_string)
           # <li><a href ='/control/display_past_system_alerts'  target="_self">Current System State</a></li>     
           sub_system_data = self.url_rule_class.subsystems[i]
           temp = sub_system_data["menu_data"]
           
           #
        
           for j in sub_system_data['menu_list']:
                 data = temp[j]      
                 #print("data",data)  
                 format_output = '<li><a href='+'"/'+i+'/'+data[1]+'" target="_self">'+data[2]+'</a></li>\n'                 
                 f.write(format_output)
                  
       

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
  
   
   pi_web_server.generate_site_map()
   
   pi_web_server.generate_default_index_page()
   port = pi_web_server.result["port"]
   pi_web_server.port = port
   debug = pi_web_server.result["debug"]
   pi_web_server.debug = debug
   https_flag = pi_web_server.result["https"]
   if https_flag == False:
       pi_web_server.run_https()
   else:
       pi_web_server.run_https()
   
