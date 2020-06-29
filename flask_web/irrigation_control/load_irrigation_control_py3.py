
import os
import json
from datetime import datetime
import time
from base_stream_processing.base_stream_processing_py3 import Base_Stream_Processing
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
from  collections import OrderedDict

from flask import render_template,jsonify
from flask import request, session, url_for


def generate_irrigation_control(redis_site_data,qs ):
       

       query_list = []
       query_list = qs.add_match_relationship( query_list,relationship="SITE",label=redis_site_data["site"] )

       query_list = qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", property_mask={"name":"IRRIGATION_CONTROL_MANAGEMENT"} )
                                           
       package_sets, package_sources = qs.match_list(query_list)  
     
       package = package_sources[0] 
       data_structures = package["data_structures"]
       generate_handlers = Generate_Handlers(package,qs)
    
       return generate_handlers.construct_managed_hash(data_structures["IRRIGATION_CONTROL"])   

class Load_Irrigation_Control(object):
      
       def __init__(base_self,self):
           query_list = []
           query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )

           query_list = self.qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", property_mask={"name":"IRRIGIGATION_SCHEDULING_CONTROL_DATA"} )
                                           
           package_sets, package_sources = self.qs.match_list(query_list)  
     
           package = package_sources[0] 
           data_structures = package["data_structures"]
           generate_handlers = Generate_Handlers(package,self.qs)
           ds_handlers = {}
           ds_handlers["IRRIGATION_JOB_SCHEDULING"] = generate_handlers.construct_job_queue_client(data_structures["IRRIGATION_JOB_SCHEDULING"])
           ds_handlers["IRRIGATION_PENDING"] = generate_handlers.construct_job_queue_client(data_structures["IRRIGATION_PENDING"])
           ds_handlers["IRRIGATION_PAST_ACTIONS"] = generate_handlers.construct_redis_stream_reader(data_structures["IRRIGATION_PAST_ACTIONS"])
           query_list = []
           query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )

           query_list = self.qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", property_mask={"name":"MQTT_DEVICES_DATA"} )
                                           
           package_sets, package_sources = self.qs.match_list(query_list)
           package = package_sources[0]
           generate_handlers = Generate_Handlers(package,self.qs)
           data_structures = package["data_structures"]
           ds_handlers["MQTT_SENSOR_QUEUE"] = generate_handlers.construct_redis_stream_reader(data_structures["MQTT_SENSOR_QUEUE"])
    
           irrigation_control = generate_irrigation_control(self.site_data,self.qs)
           query_list = []   
           query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
           query_list = self.qs.add_match_relationship( query_list,relationship="PLC_MEASUREMENTS" )
           query_list = self.qs.add_match_terminal( query_list, 
                                           relationship = "PACKAGE", 
                                           property_mask={"name":"PLC_MEASUREMENTS_PACKAGE"} )
                                           
           package_sets, package_sources = self.qs.match_list(query_list)
       
           package = package_sources[0]       
   
        
           data_structures = package["data_structures"]
           generate_handlers = Generate_Handlers(package,self.qs)
       
           ds_handlers["PLC_MEASUREMENTS_STREAM"] = generate_handlers.construct_redis_stream_reader(data_structures["PLC_MEASUREMENTS_STREAM"])              
       
       
           irrigation_control = generate_irrigation_control(self.site_data,self.qs)
           Load_Irrigation_Pages(self.app, self.auth,render_template,request, file_server_library=self.file_server_library, 
                              handlers= ds_handlers ,irrigation_control=irrigation_control,subsystem_name="Irrigation_Control",path="irrigation_control",url_rule_class = self.url_rule_class)


class Load_Irrigation_Pages(Base_Stream_Processing):

   def __init__( self,app, auth,render_template,request,file_server_library,url_rule_class,subsystem_name,path,handlers,irrigation_control):
       self.app      = app
       self.auth     = auth
       self.request  = request
      
       self.render_template = render_template
       
       self.handlers = handlers
       self.irrigation_control = irrigation_control
       self.file_server_library = file_server_library
       self.url_rule_class = url_rule_class
       self.subsystem_name = subsystem_name
       self.path = path
       self.assemble_url_rules()
       self.path_dest = self.url_rule_class.move_directories(self.path)     
 


       



   def assemble_url_rules(self):
    
       self.slash_name = "/"+self.subsystem_name
       self.menu_data = {}
       self.menu_list = []
       
       

       
       function_list =   [ self.queue_irrigation_jobs,
                           self.irrigation_queue,
                           self.display_past_actions,
                           self.plc_stream,
                           self.manage_parameters,
                           self.diagnostic_schedule,
                           self.diagnostic_schedule_valve  ]                      
                        
                           
 
                               
   
      
       url_list = [
                      [ "irrigation_control" ,'','',"Irrigation Control"  ],
                      [ "irrigation_queue" ,'','',"Irrigation Queue"   ],   
                      [ "past_actions"     ,"","","Past Actions"],                      
                      [ "plc_stream",'','',"PLC Data Stream"  ],
                      [ "manage_parameters" ,'','',"Manage Parameters"  ],
                      [ "diagnostic_schedule" ,'','',"Diagnostic Schedule"  ],
                      [ "diagnostic_remote_valve" ,'','',"Diagnostic Remote/Valve"  ]
        ]

       self.url_rule_class.add_get_rules(self.subsystem_name,function_list,url_list)
                  

       '''
             Ajax Handlers **********************************************************************
             
       '''
          

       a1 = self.auth.login_required( self.mode_change )
       self.app.add_url_rule('/ajax/mode_change',"get_mode_change",a1, methods=["POST"]) 


       a1 = self.auth.login_required( self.parameter_update )
       self.app.add_url_rule('/ajax/parameter_update',"get_parameter_update",a1, methods=["POST"]) 
       
       a1 = self.auth.login_required( self.irrigation_job_delete )
       self.app.add_url_rule('/ajax/irrigation_job_delete',"irrigation_job_delete",a1, methods=["POST"]) 
 
       a1 = self.auth.login_required( self.irrigation_status_update )
       self.app.add_url_rule('/ajax/status_update',"status_update",a1, methods=["GET"])
       

   
                          
       
 
   #### need a no data page
   def plc_stream(self):
       
       temp_data = self.handlers["PLC_MEASUREMENTS_STREAM"].revrange("+","-" , count=2160) # 1.5 days
       
       temp_data.reverse()
       '''
       filtered_data = []
       for i in temp_data:
          temp = i["data"]
          temp["timestamp"] = i["timestamp"]
          filtered_data.append(temp)
       '''   
       chart_title = ""
       
       stream_keys,stream_range,stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="",title_x="Date")
       stream_keys.sort()      
       
       return self.render_template( "streams/base_stream",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range ,
                                     max_value = 10000000.,
                                     min_value = -10000000,)
  
   def queue_irrigation_jobs(self):
       schedule_data = self.get_schedule_data()
      
       return self.render_template(self.path+"/irrigation_control",schedule_data = json.dumps(schedule_data))

   def irrigation_queue(self):
       jobs = self.get_queued_irrigation_jobs()
      
       return self.render_template(self.path+"/irrigation_queue",jobs = jobs )


   def manage_parameters(self):
      
       control_data = {}
       control_data["RAIN_FLAG"] = self.irrigation_control.hget("RAIN_FLAG")
       control_data["ETO_MANAGEMENT"] = self.irrigation_control.hget("ETO_MANAGEMENT")
       control_data["FLOW_CUT_OFF"]   =    self.irrigation_control.hget("FLOW_CUT_OFF")
       control_data["CLEANING_INTERVAL"] = self.irrigation_control.hget("CLEANING_INTERVAL")
 
       control_data_json = json.dumps(control_data)
     
       return self.render_template( self.path+"/manage_parameters",
                                    title = "Manage Irrigation Parameters",
                                    control_data_json = control_data_json  )
       

      

   def display_irrigation_queue(self):

       return self.render_template(self.path+"/display_irrigation_queue" )
    
   def display_past_actions( self):
       temp_data = self.handlers["IRRIGATION_PAST_ACTIONS"].revrange("+","-" , count=1000)
      
       for i in temp_data:
         i["time"] = str(datetime.fromtimestamp(i["timestamp"]))
       return self.render_template(self.path+"/irrigation_history_queue" ,time_history = temp_data )
                 
   def diagnostic_schedule(self):
       return self.get_diagnostic_page("schedule_control")
       
   def diagnostic_schedule_valve (self):
       return self.get_diagnostic_page("controller_pin")       
       
       
      
   def get_diagnostic_page(self, filename): 
       if filename == "schedule_control":
          title = 'Irrigation Diagnostics Turn On by Schedule'
          
       if filename == "controller_pin":
          title = 'Irrigation Diagnostics Turn On by Controller/Pin'
          
       if filename == "valve_group":
          title = 'Irrigation Diagnostics Turn On by Valve Group'
       schedule_data = self.get_schedule_data()

       controller_pin = self.file_server_library.load_file( "system_files", "controller_cable_assignment.json" )
       controller_pin_json = json.dumps(controller_pin)

       controller_valve_group = self.file_server_library.load_file( "system_files", "valve_group_assignments.json" )
       return self.render_template(self.path+'/irrigation_diagnostics', 
             filename = filename,
             title = title,
             schedule_data = schedule_data ,
             controller_pin = controller_pin, 
             controller_valve_group = controller_valve_group             )                          

   #  
   #  Function serves a post operation
   #

   def get_schedule_data(self):
     sprinkler_ctrl           = json.loads(self.file_server_library.load_file( "application_files", "sprinkler_ctrl.json" ))   
     
     returnValue = []
     for j in sprinkler_ctrl:
         temp          = json.loads(self.file_server_library.load_file( "application_files", j["link"] )) 

         j["step_number"], j["steps"], j["controller_pins"] = self.generate_steps(temp)
         returnValue.append(j)
     return json.dumps(returnValue)   

   def generate_steps( self, file_data):
  
       returnValue = []
       controller_pins = []
       if file_data["schedule"] != None:
           schedule = file_data["schedule"]
           for i  in schedule:
               returnValue.append(i[0][2])
               temp = []
               for l in  i:
                   temp.append(  [ l[0], l[1][0] ] )
               controller_pins.append(temp)
  
  
       return len(returnValue), returnValue, controller_pins
   def generate_steps( self, file_data):
  
       returnValue = []
       controller_pins = []
       if file_data["schedule"] != None:
           schedule = file_data["schedule"]
           for i  in schedule:
               returnValue.append(i[0][2])
               temp = []
               for l in  i:
                   temp.append(  [ l[0], l[1][0] ] )
               controller_pins.append(temp)
  
  
       return len(returnValue), returnValue, controller_pins

   def mode_change( self):
       json_object = self.request.json
       
       self.handlers["IRRIGATION_JOB_SCHEDULING"].push(json_object)
      
       return json.dumps("SUCCESS")

   def parameter_update(self):
       json_object = self.request.json
      
       field = json_object["field"]
       data = json_object["data"]
       data = float(data)
       self.irrigation_control.hset(field,data)
      
       return json.dumps("SUCCESS")

        
   def get_queued_irrigation_jobs(self):
      
       results = self.handlers["IRRIGATION_PENDING"].list_range(0,-1)
       results.reverse()
       return_value = []
       for i in results:
          temp = {}
          temp["schedule_name"] = i["schedule_name"]
          temp["step"]   = i["step"]
          temp["run_time"] = i["run_time"]
          return_value.append(temp)
       
       return return_value
       
   def irrigation_job_delete(self):
        json_object = self.request.json
        list_index = json_object["list_indexes"]        
       
        self.handlers["IRRIGATION_PENDING"].delete_jobs(list_index)
        return json.dumps("SUCCESS")
  
   def irrigation_status_update(self):
       temp = self.irrigation_control.hget_all()
       return json.dumps(temp)
       
       
       
"""
  def irrigation_sensor_stream(self):

       temp_data = self.handlers["MQTT_SENSOR_QUEUE"].revrange("+","-" , count=2160) # 1.5 days
       temp_data.reverse()
       '''
       filtered_data = []
       for i in temp_data:
          temp = i["data"]
          temp["timestamp"] = i["timestamp"]
          filtered_data.append(temp)
       '''   
       chart_title = ""
       
       stream_keys,stream_range,stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="",title_x="Date")
       stream_keys.sort()      
       
       return self.render_template( "streams/base_stream",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range ,
                                     max_value = 10000000.,
                                     min_value = -10000000,)
                                
   def queue_irrigation_jobs(self):
       schedule_data = self.schedule_data()
      
       return self.render_template("irrigation_templates/irrigation_control",schedule_data = schedule_data)
      
 

                      
"""