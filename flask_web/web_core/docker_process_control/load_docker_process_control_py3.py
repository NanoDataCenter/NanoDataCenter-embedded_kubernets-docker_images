
import os
import json
from datetime import datetime
import time
import datetime
 
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers 
 
class Load_Docker_Processes(object):

   def __init__( self, app, auth, request, render_template,qs,site_data,url_rule_class,subsystem_name,path):
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
       
       function_list = [ self.process_control ,
                         self.display_exception_status,
                         self.display_exception_log ]
                         
       url_list = [ [ 'start_and_stop_processes','/<int:container_id>','/0',"Stop/Start Docker Processes"  ],
                    [ 'display_exception_status','/<int:container_id>','/0',"Docker Processes Status"  ],
                     [ 'display_exception_log','/<int:container_id>','/0',"Docker Process Exception Log"  ] ]                            

      
       self.url_rule_class.add_get_rules(self.subsystem_name,function_list,url_list)
      

       
       
       # internal callable
       a1 = self.auth.login_required( self.load_processes )
       self.app.add_url_rule(self.slash_name+'/manage_processes/load_process',self.subsystem_name+"docker_process_load_process",a1,methods=["POST"])
       
       # internal call
       a1 = self.auth.login_required( self.manage_processes )
       self.app.add_url_rule(self.slash_name+'/manage_processes/change_process',self.subsystem_name+"docker_process_change_process",a1,methods=["POST"])
    
    
   def assemble_handlers(self):  
   
       #
       #
       # First step is to find controllers
       #
       #
   
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )

       query_list = self.qs.add_match_terminal( query_list, 
                                        relationship = "CONTAINER" )
                                           
       docker_sets, docker_nodes = self.qs.match_list(query_list)  
       self.docker_names = []
       for i in docker_nodes:
           self.docker_names.append(i["name"])
       self.docker_names.sort()
       
       #
       #
       # Assemble data structures for each controller
       #
       #
       self.handlers = []
       for i in self.docker_names:
         
          self.handlers.append(self.assemble_data_structures(i))

   
       

 
   def assemble_data_structures(self,container_name ):
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )

       query_list = self.qs.add_match_relationship( query_list, relationship = "CONTAINER",label = container_name )
      
       query_list = self.qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE",label="DATA_STRUCTURES" )
                                           
       package_sets, package_sources = self.qs.match_list(query_list)  
     
       package = package_sources[0] 
       data_structures = package["data_structures"]
       generate_handlers = Generate_Handlers(package,self.qs)
       
       handlers = {}
       handlers["ERROR_STREAM"]        = generate_handlers.construct_redis_stream_reader(data_structures["ERROR_STREAM"])
       handlers["ERROR_HASH"]        = generate_handlers.construct_hash(data_structures["ERROR_HASH"])
       handlers["WEB_COMMAND_QUEUE"]   = generate_handlers.construct_job_queue_client(data_structures["WEB_COMMAND_QUEUE"])
       
       handlers["WEB_DISPLAY_DICTIONARY"]   =  generate_handlers.construct_hash(data_structures["WEB_DISPLAY_DICTIONARY"])
       return handlers


   #
   #
   #
   #  Web page handlers
   #
   #
   #

   def process_control(self,container_id):
      
      display_list = self.handlers[container_id]["WEB_DISPLAY_DICTIONARY"].hkeys()
      
      return self.render_template(self.path_dest+"/docker_process_control",
                                  display_list = display_list, 
                                  command_queue_key = "WEB_COMMAND_QUEUE",
                                  process_data_key = "WEB_DISPLAY_DICTIONARY",
                                  container_id = container_id,
                                  containers = self.docker_names,
                                  load_process = '"'+self.slash_name+'/manage_processes/load_process'+'"',
                                  manage_process =  '"'+self.slash_name+'/manage_processes/change_process'+'"' )
      


   def load_processes(self):
       param = self.request.get_json()
      
       container = int(param["container"])
       
       if container >= len(self.docker_names):
          return "BAD"
       else:
          result = self.handlers[container]["WEB_DISPLAY_DICTIONARY"].hgetall()
          result_json = json.dumps(result)
          
          return result_json.encode()
          

   def manage_processes(self):
       param = self.request.get_json()
      
       container = int(param["container"])
       process_state_json = param["process_data"]
       process_state = json.loads(process_state_json)
       if container >= len(self.docker_names):
          return "BAD"
       else:
          
          self.handlers[container]["WEB_COMMAND_QUEUE"].push(process_state)
          return json.dumps("SUCCESS")
          
   def display_exception_status(self,container_id):

       container_exceptions = self.handlers[container_id]["ERROR_HASH"].hgetall()

       for i in container_exceptions.keys():
           if "time" in container_exceptions[i]:
               temp =container_exceptions[i]["time"]
               container_exceptions[i]["time"] = datetime.datetime.utcfromtimestamp(temp).strftime('%Y-%m-%d %H:%M:%S')
           
           temp = container_exceptions[i]["error_output"]
           container_exceptions[i]["error_output"] = [temp]
      
       return self.render_template(self.path_dest+"/docker_process_exception_status",
                                  container_keys = container_exceptions.keys(),
                                  container_exceptions = container_exceptions,
                                  container_id = container_id,
                                  containers = self.docker_names )
                                  
                                  
   def display_exception_log(self,container_id):
       temp_list = self.handlers[container_id]["ERROR_STREAM"].revrange("+","-" , count=20)
      
       container_exceptions = []
      
       for j in temp_list:
           i = j["data"]
           i["timestamp"] = j["timestamp"]
           i["datetime"] =  datetime.datetime.fromtimestamp( i["timestamp"]).strftime('%Y-%m-%d %H:%M:%S')

           temp = i["error_output"]
           if len(temp) > 0:
               temp = i["error_output"]
               if len(temp) > 0:
                   temp = [temp]
                   #temp = temp.split("\n")
                   i["error_output"] = temp
                   container_exceptions.append(i)
       
       return self.render_template(self.path_dest+"/docker_process_exception_log",                                 
                                  log_data = container_exceptions,
                                  container_id = container_id,
                                  containers = self.docker_names )
                                  
  
 