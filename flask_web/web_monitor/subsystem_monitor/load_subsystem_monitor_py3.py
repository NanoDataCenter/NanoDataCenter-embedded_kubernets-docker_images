import os
import json
from datetime import datetime
import time
from base_stream_processing.base_stream_processing_py3 import Base_Stream_Processing
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
class Load_Subsystem_Monitor(object):

   def __init__(self, app, auth, request, render_template,qs,site_data,url_rule_class,subsystem_name,path):
       self.app      = app
       self.auth     = auth
       self.request  = request
       self.render_template = render_template
       self.path = path
       self.qs = qs
       self.site_data = site_data
       self.url_rule_class = url_rule_class
       self.subsystem_name = subsystem_name
       
       self.assemble_handlers()
       self.assemble_url_rules()
       self.path_dest = self.url_rule_class.move_directories(self.path)     

       
   def assemble_url_rules(self):
    
       self.slash_name = "/"+self.subsystem_name
       self.menu_data = {}
       self.menu_list = []
       
       

       
       function_list =   [ self.system_status,
                         ]
                           
 
                               
   
      
       url_list = [
                      [ 'system_status' ,'','',"System_Status"  ] 
                  ]

       self.url_rule_class.add_get_rules(self.subsystem_name,function_list,url_list)
            
 
   def assemble_handlers(self):
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
       query_list = self.qs.add_match_relationship( query_list,relationship="SYSTEM_MONITOR" )
       
       
       query_list = self.qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", label = "SYSTEM_MONITOR" )
                                           
       package_sets, package_sources = self.qs.match_list(query_list)  
      
       package = package_sources[0]
       generate_handlers = Generate_Handlers(package,self.qs)
       data_structures = package["data_structures"]     
 
       self.handlers = {}
       self.handlers["SYSTEM_STATUS"] = generate_handlers.construct_hash(data_structures["SYSTEM_STATUS"])
       self.handlers["MONITORING_DATA"] = generate_handlers.construct_hash(data_structures["MONITORING_DATA"])
       self.handlers["SYSTEM_ALERTS"] = generate_handlers.construct_redis_stream_reader(data_structures["SYSTEM_ALERTS"])
       


   def  system_status(self):
       
       subsystems = self.handlers["SYSTEM_STATUS"].hkeys()
       states    = self.handlers["SYSTEM_STATUS"].hgetall()
       values    =  self.handlers["MONITORING_DATA"].hgetall()
       subsystems.sort()
       print("subsystems",subsystems)
       links = self.generate_links(subsystems,values)
       number_of_errors = self.generate_errors(subsystems,values)
       
       
       
       return self.render_template( self.path_dest+"/list_subsystem_status",
                                    subsystems = subsystems,
                                    states     = states, 
                                    links      = links,
                                    number_of_errors = number_of_errors
                                     )

   def generate_links(self,subsystems,values):
       return_value = {}
       for i in subsystems:
           return_value[i] = self.generate_a_link(i,values[i])
       return return_value
       
   def generate_a_link(self,subsystem,value):
       if subsystem in ['MONITOR_BLOCK_CHAIN','MONITOR_REDIS']:
          print(subsystem)
          print(value)
          link = "/"+self.subsystem_name+"/single_level/"+subsystem
          
       elif subsystem in ['MONITOR_SQLITE']:
          print(subsystem)
          print(value)       
          link = "/"+self.subsystem_name+"/single_level/"+subsystem
          
       else:
           print(subsystem)
           print(value)
           
           raise
       print("link",link)
       return link
       
   def generate_errors(self,subsystems,values):
       return_values = []
       for i in subsystems:
          return_value = return_value + values[i][0]
       return return_value