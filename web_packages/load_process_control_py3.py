
import os
import json
from datetime import datetime
import time
import zlib
import datetime

   def assemble_process_controller_handlers(self,label ):
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.redis_site_data["site"] )

       query_list = self.qs.add_match_relationship( query_list, relationship = "PROCESSOR", label = label )
       query_list = self.qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", property_mask={"name":"DATA_STRUCTURES"} )
                                           
       package_sets, package_sources = self.qs.match_list(query_list)  
     
       package = package_sources[0] 
       data_structures = package["data_structures"]
       generate_handlers = Generate_Handlers(package,self.qs)
       ds_handlers = {}
       ds_handlers["ERROR_STREAM"]        = generate_handlers.construct_redis_stream_reader(data_structures["ERROR_STREAM"])
       ds_handlers["ERROR_HASH"]        = generate_handlers.construct_hash(data_structures["ERROR_HASH"])
       ds_handlers["WEB_COMMAND_QUEUE"]   = generate_handlers.construct_job_queue_client(data_structures["WEB_COMMAND_QUEUE"])
       
       ds_handlers["WEB_DISPLAY_DICTIONARY"]   =  generate_handlers.construct_hash(data_structures["WEB_DISPLAY_DICTIONARY"])
       return ds_handlers
class Load_Process_Management(object):

   def __init__( self, app, auth, request, render_template,controller_names, handlers):
       self.app      = app
       self.auth     = auth
       self.request  = request
       self.render_template = render_template
       self.controller_names = controller_names
       self.handlers = handlers
       
       #
       #  Adding Rules
       #

       a1 = auth.login_required( self.process_control )
       app.add_url_rule('/start_and_stop_processes/<int:controller_id>',"start_and_stop_processes",a1)

       a1 = auth.login_required( self.display_exception_status )
       app.add_url_rule('/display_exception_status/<int:controller_id>',"display_exception_status",a1)
       
       a1 = auth.login_required( self.display_exception_log )
       app.add_url_rule('/display_exception_log/<int:controller_id>',"display_exception_log",a1)


       
       # internal callable
       a1 = auth.login_required( self.load_processes )
       app.add_url_rule('/manage_processes/load_process',"load_process",a1,methods=["POST"])
       
       # internal call
       a1 = auth.login_required( self.manage_processes )
       app.add_url_rule('/manage_processes/change_process',"change_process",a1,methods=["POST"])



   def process_control(self,controller_id):
      
      display_list = self.handlers[controller_id]["WEB_DISPLAY_DICTIONARY"].hkeys()
      
      return self.render_template("process_control/process_control",
                                  display_list = display_list, 
                                  command_queue_key = "WEB_COMMAND_QUEUE",
                                  process_data_key = "WEB_DISPLAY_DICTIONARY",
                                  controller_id = controller_id,
                                  controllers = self.controller_names )
      


   def load_processes(self):
       param = self.request.get_json()
      
       controller = int(param["controller"])
       
       if controller >= len(self.controller_names):
          return "BAD"
       else:
          result = self.handlers[controller]["WEB_DISPLAY_DICTIONARY"].hgetall()
          result_json = json.dumps(result)
          
          return result_json.encode()
          

   def manage_processes(self):
       param = self.request.get_json()
      
       controller = int(param["controller"])
       process_state_json = param["process_data"]
       process_state = json.loads(process_state_json)
       if controller >= len(self.controller_names):
          return "BAD"
       else:
          
          self.handlers[controller]["WEB_COMMAND_QUEUE"].push(process_state)
          return json.dumps("SUCCESS")
          
   def display_exception_status(self,controller_id):

       controller_exceptions = self.handlers[controller_id]["ERROR_HASH"].hgetall()

       for i in controller_exceptions.keys():
           if "time" in controller_exceptions[i]:
               temp = controller_exceptions[i]["time"]
               controller_exceptions[i]["time"] = datetime.datetime.utcfromtimestamp(temp).strftime('%Y-%m-%d %H:%M:%S')
           
           temp = controller_exceptions[i]["error_output"]
           controller_exceptions[i]["error_output"] = [temp]
      
       return self.render_template("process_control/exception_status",
                                  controller_keys = controller_exceptions.keys(),
                                  controller_exceptions = controller_exceptions,
                                  controller_id = controller_id,
                                  controllers = self.controller_names )
                                  
                                  
   def display_exception_log(self,controller_id):
       temp_list = self.handlers[controller_id]["ERROR_STREAM"].revrange("+","-" , count=20)
       
       controller_exceptions = []
      
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
                   controller_exceptions.append(i)
       
       return self.render_template("process_control/exception_log",                                 
                                  log_data = controller_exceptions,
                                  controller_id = controller_id,
                                  controllers = self.controller_names )
                                  
  
 