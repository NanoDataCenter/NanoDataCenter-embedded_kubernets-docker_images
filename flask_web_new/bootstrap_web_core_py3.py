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
#from flask_httpauth import HTTPDigestAuth
from flask import request, session, url_for

from redis_support_py3.graph_query_support_py3 import  Query_Support
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers

from file_server_lib_py3 import Construct_RPC_Library
from load_redis_access_py3  import  Load_Redis_Access
from register_redis_data_structures_py3 import Register_Redis_Data_Structures
from extract_web_map_from_graph_py3 import Extract_and_Build_Web_Structure
from register_template_classes_py3 import Register_Template_Classes
import os

from macro_processor_py3 import Macro_Processor
from construct_and_register_web_page_py3 import Construct_And_Register_Pages


class Construct_Base_Page(object):

    def __init__(self,base_self):
        self.base_self = base_self
        
        
        self.base_page_top = base_self.mp.macro_expand_file("~~","--","templates/base_page_blocks/base_top")
      
        self.base_page_bottom = base_self.mp.macro_expand_file("~~","--","templates/base_page_blocks/base_bottom")
       
        
    


class Load_Static_Files(object):
 
   def __init__( self,app):
       self.app   = app

       app.add_url_rule('/favicon.ico',"get_fav",self.get_fav )
       app.add_url_rule('/js/<path:filename>','get_js',self.get_js)
       app.add_url_rule('/css/<path:filename>','get_css',self.get_css )
       app.add_url_rule('/html/<path:filename>',"get_html",self.get_html )



  
   def get_fav(self):
       return  self.app.send_static_file("favicon.ico")
       
       
   def get_js(self, filename):
       return self.app.send_static_file(os.path.join('js', filename))

 
   def get_css(self, filename):
       return self.app.send_static_file(os.path.join('css', filename))


   def get_html(self, filename):
       return self.app.send_static_file(os.path.join('html', filename))








class Load_App_Sys_Files(object):

   def __init__( self, app, request,qs,site_data ):
       self.app      = app
       self.request  = request
      
       self.file_server_library = Construct_RPC_Library(qs,site_data)
       app.add_url_rule("/ajax/get_system_file/<path:file_name>","get_system_file",self.get_system_file)
       app.add_url_rule("/ajax/get_app_file/<path:file_name>","get_app_file",self.get_app_file)
       app.add_url_rule("/ajax/save_app_file/<path:file_name>","save_app_file",self.save_app_file,methods=["POST"])
       app.add_url_rule("/ajax/save_sys_file/<path:file_name>","save_sys_file",self.save_sys_file,methods=["POST"])
               


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
   
   
       self.mp = Macro_Processor()
       redis_handle_pw = redis.StrictRedis(site_data["host"], 
                                           site_data["port"], 
                                           db=site_data["redis_password_db"], 
                                           decode_responses=True)
                               


       
       #  
       #
       # Setup Web Server Based on Configuration Data
       #
       #
       self.site_data = site_data                                       
       startup_dict = redis_handle_pw.hgetall("web")
       self.app         = Flask(name) 
       self.startup_dict = startup_dict
       self.app.template_folder       =   'flask_templates'
       self.app.static_folder         =   'static'  
       self.app.config['SECRET_KEY']      = startup_dict["SECRET_KEY"]
       self.qs = Query_Support( site_data)
       #
       #
       # Get Server Properties
       #
       #
       results=self.common_qs_search(["WEB_SERVER","WEB_SERVER"])
       self.server_properties = results[0]
       #
       #  Set Up Basic Web URL Services
       #
       
       
    
       # Allows Web server to read/write System and Application Filte
       Load_App_Sys_Files(self.app, request,self.qs,self.site_data )
       

       Load_Static_Files(self.app) #enable static files to be fetched
       self.redis_access = Load_Redis_Access(self.app, request ) #enable web access for redis operations
       Register_Redis_Data_Structures(self)# pass properties of PI_Web_Server
      
      
      
      
       self.class_map = {}
       
       self.class_table = {}
       self.menu_data =  []
       self.url_instanciated_classes = {}
       self.request = request
       
       Register_Template_Classes(self)
       Extract_and_Build_Web_Structure(self) # pass properties of PI_Web_Server
       self.bp = Construct_Base_Page(self)  ## construct base page
       Construct_And_Register_Pages(self)
       
       
       
       
      
 
  

       

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
                                         
 
  
   def run(self):
      if self.server_properties["https"]:
         self.run_https()
      else:
         self.run_http()
        

   def run_http( self):
       debug = self.server_properties["debug"]
       port = self.server_properties["port"]
       self.app.run(threaded=True , use_reloader=True, host='0.0.0.0',port=port,debug=debug)

   def run_https( self ):
       debug = self.server_properties["debug"]
       port = self.server_properties["port"]
     
       self.app.run(threaded=True , use_reloader=True, host='0.0.0.0',debug =debug,
           port= port ,ssl_context=("/data/cert.pem", "/data/key.pem"))
       
 


         



if __name__ == "__main__":

   file_handle = open("/data/redis_server.json",'r')
   data = file_handle.read()
   file_handle.close()
  
   redis_site_data = json.loads(data)


   pi_web_server = PI_Web_Server_Core(__name__, redis_site_data  )
   pi_web_server.run()
   
