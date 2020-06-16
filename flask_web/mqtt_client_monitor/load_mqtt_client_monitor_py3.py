import os
import json
from datetime import datetime
import time
from base_stream_processing.base_stream_processing_py3 import Base_Stream_Processing
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
from  collections import OrderedDict
class Load_MQTT_Client_Monitoring(Base_Stream_Processing):
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
       
       

       
       function_list =   [ self.mqtt_device_status,
                           self.mqtt_reboot_status,
                           self.mqtt_unknown_devices,
                           self.mqtt_unknown_commands,                      
                           self.mqtt_past_actions ]
                           
 
                               
   
      
       url_list = [
                      [ 'mqtt_device_status' ,'','',"MQTT Device Status"  ],
                      [ 'mqtt_reboot_status' ,'','',"MQTT Device Reboots"  ],
                      [ 'mqtt_unknown_devices' ,'','',"MQTT Unknown Devices"   ],
                      [ 'mqtt_unknown_commands' ,'','',"MQTT Unknown Commands"  ],    
                      [ 'mqtt_past_actions','','',"MQTT_Past_Actions"  ]
        ]

       self.url_rule_class.add_get_rules(self.subsystem_name,function_list,url_list)
            

   def assemble_handlers(self):
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
       query_list = self.qs.add_match_relationship( query_list,relationship="MQTT_DEVICES" )
       
       
       query_list = self.qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", label = "MQTT_DEVICES_DATA" )
                                           
       package_sets, package_sources = self.qs.match_list(query_list)  
      
       package = package_sources[0]
       generate_handlers = Generate_Handlers(package,self.qs)
       data_structures = package["data_structures"]     
 
       self.handlers = {}
       self.handlers["MQTT_PAST_ACTION_QUEUE"] = generate_handlers.construct_redis_stream_reader(data_structures["MQTT_PAST_ACTION_QUEUE"])
       self.handlers["MQTT_CONTACT_LOG"] = generate_handlers.construct_hash(data_structures["MQTT_CONTACT_LOG"])
       self.handlers["MQTT_REBOOT_LOG"] = generate_handlers.construct_hash(data_structures["MQTT_REBOOT_LOG"])
       self.handlers["MQTT_UNKNOWN_DEVICES"] = generate_handlers.construct_hash(data_structures["MQTT_UNKNOWN_DEVICES"])
       self.handlers["MQTT_UNKNOWN_SUBSCRIPTIONS"]  = generate_handlers.construct_hash(data_structures["MQTT_UNKNOWN_SUBSCRIPTIONS"])
                                             
       
 

   def mqtt_device_status( self):
       temp_data = self.handlers["MQTT_CONTACT_LOG"].hgetall()
       
       for key,item in temp_data.items():
 
         item["time"] = str(datetime.fromtimestamp(item["time"]))
         item["detail"] = "--- Device Id: "+ item["device_id"]+"  Date:  "+item['time'] + " Status "+str(item["status"])
       
       return self.render_template(self.path_dest+"/mqtt_status_template" ,time_history = temp_data ,title = "Device Status"  )



   def mqtt_reboot_status( self):
       temp_data = self.handlers["MQTT_REBOOT_LOG"].hgetall()
      
       for key,i in temp_data.items():
           i["time"] = str(datetime.fromtimestamp(int(i["timestamp"]))) 
           i["detail"] = "--- Device Id: "+ i["device_id"]+" Reboot Date:  "+i['time'] 
       return self.render_template(self.path_dest+"/mqtt_status_template" ,time_history = temp_data, title ="Reboot Status" )


   def mqtt_unknown_devices( self):
       temp_data = self.handlers["MQTT_UNKNOWN_DEVICES"].hgetall()
       for i in temp_data:
           temp_data[i]["detail"] = "Device:  "+i+"  Topic: "+temp_data[i]["topic"]
       return self.render_template(self.path_dest+"/mqtt_status_template" ,time_history = temp_data,title="Unknown Devices" )


   def mqtt_unknown_commands( self):
       temp_data = self.handlers["MQTT_UNKNOWN_SUBSCRIPTIONS"].hgetall()
      
       for key,i in temp_data.items():
         i["time"] = str(datetime.fromtimestamp(i["timestamp"]))
         i["detail"] = "--- Date:  "+i['time'] + " Topic "+str(i['topic'])
       return self.render_template(self.path_dest+"/mqtt_status_template" ,time_history = temp_data,title ="Unknown Commands" )



   def mqtt_past_actions( self):
       temp_data = self.handlers["MQTT_PAST_ACTION_QUEUE"].revrange("+","-" , count=1000)
      
       collated_data = []
       for  j in temp_data:
         i= j["data"]
         i["time"] = str(datetime.fromtimestamp(j["timestamp"]))
         i["detail"] = "---Date: "+i["time"]+"  Action: "+i["action"]+"  Device Id: "+i["device_id"] +" Status: "+str(i["status"] )
         collated_data.append(i)
       return self.render_template(self.path_dest+"/mqtt_past_actions_template" ,time_history = collated_data,title="PAST ACTIONS" )








