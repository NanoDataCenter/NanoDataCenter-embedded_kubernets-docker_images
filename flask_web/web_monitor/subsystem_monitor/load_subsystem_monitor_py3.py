import os
import json
from datetime import datetime
import time
from base_stream_processing.base_stream_processing_py3 import Base_Stream_Processing
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
from flatten_dict import flatten
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
       self.subsystems = self.handlers["SYSTEM_STATUS"].hkeys()
       self.subsystems.sort()  
     
   def assemble_url_rules(self):
    
       self.slash_name = "/"+self.subsystem_name
       self.menu_data = {}
       self.menu_list = []
       
       

       
       function_list =   [ self.subsystem_status,
                           self.subsystem_error_status,
                           self.subsystem_error_stream_log
                         ]
                           
 
                               
   
      
       url_list = [
                      [ 'subsystem_status' ,'','',"Subsystem_Status"  ],
                      [ 'subsystem_error_status', '/<int:subsystem_id>','/0',"Subsystem_Error_Status"  ],
                      [ 'subsystem_error_log' ,'','',"Subsystem_Error_Log"  ],
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
       
      

   def  subsystem_status(self):
       
       
       error_states    = self.handlers["SYSTEM_STATUS"].hgetall()
       
       
       subsystem_error_display =  self.slash_name+"/subsystem_error_status/"
       error_stream_display = self.slash_name+"/subsystem_error_log"
       
      
       
       
       
       return self.render_template( self.path_dest+"/list_subsystem_status",
                                    subsystems = self.subsystems,
                                    states     = error_states,                                   
                                    sub_error_display = subsystem_error_display,
                                    error_stream      = error_stream_display
                                     )
                                     
   def subsystem_error_status(self,subsystem_id):

       subsystem = self.subsystems[subsystem_id]
       value    =  self.handlers["MONITORING_DATA"].hget(subsystem)
       if type(value) != dict:
          temp = value
          value = {}
          value[subsystem] = temp
       flatten_value = flatten(value,reducer='path')
       print("flatten value",flatten_value)
       json_value = {}
       for key,value in flatten_value.items():
           if type(value) == list:
               key = key+" Total Errors: "+str(value[0])
               json_value[key] = value[1]
           else:
                json_value[key] = value
       return self.render_template(self.path_dest+"/subsystem_error_display",                                 
                                   error_status = json_value,
                                   subsystem_id = subsystem_id,
                                   subsystems = self.subsystems,
                                   subsystem_name = subsystem)
                                         
                 
   def subsystem_error_stream_log(self):
       
       temp_list = self.handlers["SYSTEM_ALERTS"].revrange("+","-" , count=20)
       print("temp_list",temp_list)
       error_log = []
      
       for j in temp_list:
           i = {}
           i["data"] = json.dumps(j["data"])
        
           i["datetime"] =  datetime.datetime.fromtimestamp( j["timestamp"]).strftime('%Y-%m-%d %H:%M:%S')
           error_log.append(i)
    
       
       return self.render_template(self.path_dest+"/subsystem_error_log",                                 
                                   error_log = error_log )
                                  


   ''' 
       
   def flatten(self,input_dict, separator='_', prefix=''):
       output_dict = {}
       for key, value in input_dict.items():
           if isinstance(value, dict) and value:
               deeper = self.flatten(value, separator, prefix+key+separator)
               output_dict.update({key2: val2 for key2, val2 in deeper.items()})
           elif isinstance(value, list) and value:
               for index, sublist in enumerate(value, start=1):
                   if isinstance(sublist, dict) and sublist:
                       deeper = self.flatten(sublist, separator, prefix+key+separator+str(index)+separator)
                       output_dict.update({key2: val2 for key2, val2 in deeper.items()})
                   else:
                       output_dict[prefix+key+separator+str(index)] = value
           else:
               output_dict[prefix+key] = value
       return output_dict
       
       
   def flatten_a(self,d,separator='_'):
       out = {}
       for key, val in d.items():
           if isinstance(val, dict):
               val = [val]
           if isinstance(val, list):
               for subdict in val:
                   deeper = self.flatten_a(subdict,separator).items()
                   out.update({key + separator + key2: val2 for key2, val2 in deeper})
           else:
               out[key] = val
       return out
   '''    