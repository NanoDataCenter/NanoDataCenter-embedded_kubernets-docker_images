
import os
import json
from datetime import datetime
import time
import datetime
from redis_support_py3.graph_query_support_py3 import  Query_Support 

from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers 
from base_stream_processing.base_stream_processing_py3 import Base_Stream_Processing
from file_server_library.file_server_lib_py3  import Construct_RPC_Library

class Load_Docker_Processes(Base_Stream_Processing):

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
       
       #self.assemble_handlers()
       self.assemble_url_rules()
       self.path_dest = self.url_rule_class.move_directories(self.path)
       
       self.file_server = Construct_RPC_Library( qs, site_data )           
       self.assemble_containers()
       
       self.docker_performance_data_structures= {}
       for i in self.managed_container_names:
           self.docker_performance_data_structures[i] = self.assemble_container_data_structures(i)


   def assemble_containers(self):  
   
       #
       #
       # First step is to find controllers
       #
       #
   
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
       query_list = self.qs.add_match_terminal( query_list,relationship="PROCESSOR" )
       processor_sets,processor_nodes = self.qs.match_list(query_list) 
       self.processor_names = []
       self.all_container_names = []
       self.managed_container_names = []
       self.container_control_structure = {}
       for i in processor_nodes:
           self.processor_names.append(i["name"])
           self.container_control_structure[i["name"]] = self.determine_container_structure(i["name"])
           
           services = set(i["services"])
           self.services = list(services)
          
           containers = set(i["containers"])
           self.containers = list(containers)
           containers_list = containers.union(services)
           self.all_container_names.extend(list(containers_list))   
           
          
       self.processor_names.sort()  
       self.all_container_names.sort()
       self.containers.sort()
       self.services.sort()
       #print(self.container_control_structure)
       #print(self.all_container_names)
       self.managed_container_names = []
       for i in processor_nodes:
           self.managed_container_names.extend(self.determine_managed_containers(i["name"]))
       self.managed_container_names.sort()
       #print(self.managed_container_names)
       
       
   def determine_managed_containers(self,processor_name):
       query_list = []
       return_value = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
       query_list = self.qs.add_match_relationship( query_list,relationship="PROCESSOR",label=processor_name )
       query_list = self.qs.add_match_terminal( query_list,relationship="CONTAINER")
       container_sets,container_nodes = self.qs.match_list(query_list)
       for i in container_nodes:
          return_value.append(i["name"])
       return return_value          

   
   def determine_container_structure(self,processor_name):
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
       query_list = self.qs.add_match_relationship( query_list,relationship="PROCESSOR",label=processor_name )
       query_list = self.qs.add_match_relationship( query_list,relationship="DOCKER_MONITOR")
       query_list = self.qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", label = "DATA_STRUCTURES" )
                                        
       package_sets, package_nodes = self.qs.match_list(query_list)  
   
       #print("package_nodes",package_nodes)
   
       generate_handlers = Generate_Handlers(package_nodes[0],self.qs)      
       
          
         
       package_node = package_nodes[0]
       data_structures = package_node["data_structures"]
       
       #print(data_structures.keys())
       handlers = {}
       handlers["ERROR_STREAM"]        = generate_handlers.construct_redis_stream_reader(data_structures["ERROR_STREAM"])
      
       handlers["WEB_COMMAND_QUEUE"]   = generate_handlers.construct_job_queue_client(data_structures["WEB_COMMAND_QUEUE"])
       handlers["WEB_DISPLAY_DICTIONARY"] = generate_handlers.construct_hash(data_structures["WEB_DISPLAY_DICTIONARY"])
       queue_name = data_structures["DOCKER_UPDATE_QUEUE"]['queue']
       handlers["DOCKER_UPDATE_QUEUE"] = generate_handlers.construct_rpc_client( )
       handlers["DOCKER_UPDATE_QUEUE"].set_rpc_queue(queue_name)
       return handlers

   def assemble_container_data_structures(self,container_name):
       
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
       query_list = self.qs.add_match_relationship( query_list,relationship="CONTAINER",label=container_name)
       query_list = self.qs.add_match_terminal( query_list, 
                                           relationship = "PACKAGE", label = "DATA_STRUCTURES" )

       package_sets, package_nodes = self.qs.match_list(query_list)  
      
 
           
       #print("package_nodes",package_nodes)
   
       generate_handlers = Generate_Handlers(package_nodes[0],self.qs)
       data_structures = package_nodes[0]["data_structures"]
      
       handlers = {}
       handlers["ERROR_STREAM"]        = generate_handlers.construct_redis_stream_reader(data_structures["ERROR_STREAM"])
       handlers["ERROR_HASH"]        = generate_handlers.construct_hash(data_structures["ERROR_HASH"])
       handlers["WEB_COMMAND_QUEUE"]   = generate_handlers.construct_job_queue_client(data_structures["WEB_COMMAND_QUEUE"])
       handlers["WEB_DISPLAY_DICTIONARY"]   =  generate_handlers.construct_hash(data_structures["WEB_DISPLAY_DICTIONARY"])
       handlers["PROCESS_VSZ"]  = generate_handlers.construct_redis_stream_reader(data_structures["PROCESS_VSZ"])
       handlers["PROCESS_RSS"] = generate_handlers.construct_redis_stream_reader(data_structures["PROCESS_RSS"])
       handlers["PROCESS_CPU"]  = generate_handlers.construct_redis_stream_reader(data_structures["PROCESS_CPU"])
   
       return handlers

       
  

   def assemble_url_rules(self):
       
       
       
       self.slash_name = "/"+self.subsystem_name
       self.menu_data = {}
       self.menu_list = []
       
       function_list = [ self.container_control,
                        
                         self.container_exception_log,
                         self.process_control ,
                         self.display_exception_status,
                         self.display_exception_log,
                         self.display_cpu_percent,
                         self.display_vsz,
                         self.display_rss,
                         self.system_reset_and_docker_upgrade  ]
                         
       url_list = [ [ 'start_and_stop_container','/<int:processor_id>','/0',"Stop/Start Docker Container"  ],
                   
                     [ 'container_exception_log','/<int:processor_id>','/0',"Managed Container Process Exception Log"  ],
                    [ 'start_and_stop_managed_container_processes','/<int:container_id>','/0',"Stop/Start Managed Container Processes"  ],
                    [ 'display_exception_status','/<int:container_id>','/0',"Managed Container Processes Status"  ],
                     [ 'display_exception_log','/<int:container_id>','/0',"Managed Container Process Exception Log"  ],
                     [  'display_cpu','/<int:container_id>','/0',"Managed Container CPU Utilization"    ],
                     [  'display_vsz','/<int:container_id>','/0',"Managed Container VSZ Utilization"    ],
                     [ 'display_rss','/<int:container_id>','/0',"Managed Container RSS Utilization"  ], 
                     [ 'reset_upgrade','/<int:processor_id>','/0',"Reset System and Upgrade Containers" ]
           
                     ]                            

      
       self.url_rule_class.add_get_rules(self.subsystem_name,function_list,url_list)
      

       
       
       # internal callable
       a1 = self.auth.login_required( self.load_containers )
       self.app.add_url_rule(self.slash_name+'/manage_containers/load_containers',self.slash_name+'/manage_containers/load_containers',a1,methods=["POST"])
       
       # internal call
       a1 = self.auth.login_required( self.manage_containers )
       self.app.add_url_rule(self.slash_name+'/manage_containers/change_containers',self.slash_name+'/manage_containers/change_containers',a1,methods=["POST"])

       # internal call
       a1 = self.auth.login_required( self.load_processes )
       self.app.add_url_rule(self.slash_name+'/load_processes/load_process',self.slash_name+'/load_processes/load_process',a1,methods=["POST"])
       
       # internal call
       a1 = self.auth.login_required( self.manage_processes )
       self.app.add_url_rule(self.slash_name+'/manage_processes/change_process',self.slash_name+"/manage_processes/change_process",a1,methods=["POST"])
    
       # internal call
       a1 = self.auth.login_required( self.manage_reset_upgrade )
       self.app.add_url_rule(self.slash_name+'/manage_processes/manage_reset_upgrade',self.slash_name+"/manage_processes/manage_reset_upgrade",a1,methods=["POST"])

 
 
       


   #
   #
   #
   #  Web page handlers
   #
   #
   #

   def container_control(self,processor_id):
   
   
      processor_name = self.processor_names[processor_id]
      display_list = self.container_control_structure[processor_name]["WEB_DISPLAY_DICTIONARY"].hkeys()
      
      return self.render_template(self.path_dest+"/docker_control",
                                  display_list = display_list, 
                                  command_queue_key = "WEB_COMMAND_QUEUE",
                                  process_data_key = "WEB_DISPLAY_DICTIONARY",
                                  processor_id = processor_id,
                                  processor_names = self.processor_names,
                                  load_process = '"'+self.slash_name+'/manage_containers/load_containers'+'"',
                                  manage_process =  '"'+self.slash_name+'/manage_containers/change_containers'+'"' )
   
       

       
   def container_exception_log(self,processor_id):
       processor_name = self.processor_names[processor_id]
       temp_list = self.container_control_structure[processor_name]["ERROR_STREAM"].revrange("+","-" , count=20)

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
       
       return self.render_template(self.path_dest+"/docker_exception_log",                                 
                                  log_data = container_exceptions,
                                  processor_id = processor_id,
                                  processors = self.processor_names )           
 
   def process_control(self,container_id):
      container_name = self.managed_container_names[container_id]
      display_list = self.docker_performance_data_structures[container_name]["WEB_DISPLAY_DICTIONARY"].hkeys()
      
      return self.render_template(self.path_dest+"/docker_process_control",
                                  display_list = display_list, 
                                  command_queue_key = "WEB_COMMAND_QUEUE",
                                  process_data_key = "WEB_DISPLAY_DICTIONARY",
                                  container_id = container_id,
                                  containers = self.managed_container_names,
                                  load_process = '"'+self.slash_name+'/load_processes/load_process'+'"',
                                  manage_process =  '"'+self.slash_name+'/manage_processes/change_process'+'"' )
      
   def load_containers(self):
       param = self.request.get_json()
      
       processor_id = int(param["processor_id"])
       
       if processor_id  >= len(self.processor_names):
          return "BAD"
       else:
          processor_name = self.processor_names[processor_id]
          result = self.container_control_structure[processor_name]["WEB_DISPLAY_DICTIONARY"].hgetall()

          result_json = json.dumps(result)
          
          return result_json.encode()
          

   def manage_containers(self):
       param = self.request.get_json()
      
       processor_id = int(param["processor_id"])
       process_state_json = param["process_data"]
       process_state = json.loads(process_state_json)
      
       if processor_id >= len(self.processor_names):
          
          return "BAD"
       else:
          processor_name = self.processor_names[processor_id]
         
          self.container_control_structure[processor_name]["WEB_COMMAND_QUEUE"].push(process_state)
          return json.dumps("SUCCESS")


   def load_processes(self):
       param = self.request.get_json()
      
       container_id = int(param["container"])
       
       if container_id >= len(self.managed_container_names):
          return "BAD"
       else:
          container_name = self.managed_container_names[container_id]
          result = self.docker_performance_data_structures[container_name]["WEB_DISPLAY_DICTIONARY"].hgetall()
          result_json = json.dumps(result)
          
          return result_json.encode()
          

   def manage_processes(self):
       param = self.request.get_json()
      
       container_id = int(param["container"])
       process_state_json = param["process_data"]
       process_state = json.loads(process_state_json)
       if container_id >= len(self.managed_container_names):
          return "BAD"
       else:
          
          container_name = self.managed_container_names[container_id]
          self.docker_performance_data_structures[container_name]["WEB_COMMAND_QUEUE"].push(process_state)
          return json.dumps("SUCCESS")
          
   def display_exception_status(self,container_id):
       container_name = self.managed_container_names[container_id]
       container_exceptions = self.docker_performance_data_structures[container_name]["ERROR_HASH"].hgetall()

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
                                  containers = self.managed_container_names )
                                  
                                  
   def display_exception_log(self,container_id):
       container_name = self.managed_container_names[container_id]
       temp_list = self.docker_performance_data_structures[container_name]["ERROR_STREAM"].revrange("+","-" , count=20)
      
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
                                  containers = self.managed_container_names )
                                  
  
   def display_cpu_percent(self,container_id):
       container_name = self.managed_container_names[container_id]
       temp_data =self.docker_performance_data_structures[container_name]["PROCESS_CPU"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " %CPU Utilization "
       stream_keys,stream_range,stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="%CPU Utilization",title_x="Date")
       

       return self.render_template( "streams/stream_multi_container",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.managed_container_names,
                                     controller_id = container_id
                                     
                                     )
       
   def display_vsz(self,container_id):
       container_name = self.managed_container_names[container_id]
       temp_data =self.docker_performance_data_structures[container_name]["PROCESS_VSZ"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " VSR Utilization "
       stream_keys,stream_range,stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="VSR",title_x="Date")
       

       return self.render_template( "streams/stream_multi_container",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.managed_container_names,
                                     controller_id = container_id
                                     
                                     )
       
   def display_rss(self,container_id):
       container_name = self.managed_container_names[container_id]
       temp_data =self.docker_performance_data_structures[container_name]["PROCESS_RSS"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " RSS Utilization "
       stream_keys,stream_range,stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="RSS ",title_x="Date")
       

       return self.render_template( "streams/stream_multi_container",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.managed_container_names,
                                     controller_id = container_id
                                     
                                     )
                                     
   def manage_reset_upgrade(self):
       reset_data = self.request.get_json()
       print("reset_data",reset_data)
       processor_name = self.processor_names[reset_data["processor_id"]]
       handlers = self.container_control_structure[processor_name]
       rpc_client = handlers["DOCKER_UPDATE_QUEUE"]
       try:
          response = rpc_client.send_rpc_message("Upgrade",reset_data,timeout=5 )
          
          return json.dumps("SUCCESS")
       except:
         
          raise
    
   def system_reset_and_docker_upgrade(self,processor_id):
       processor_name = self.processor_names[processor_id]
      
       return self.render_template(self.path_dest+"/reset_upgrade",
                                  processor_id = processor_id,
                                  processor_names = self.processor_names,
                                  services = self.services,
                                  containers = self.containers,
                                  services_json = json.dumps(self.services),
                                  container_json = json.dumps(self.containers),
                                  ajax_handler = self.slash_name+'/manage_processes/manage_reset_upgrade' )
                                  
                                  
   

       