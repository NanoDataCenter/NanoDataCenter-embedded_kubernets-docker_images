from op_monitor_lib.common_class_py3 import Common_Class
import json
import time

class Docker_Processes(Common_Class):
  
   def __init__(self, subsystem_name,common_obj ):
       Common_Class.__init__(self,subsystem_name,common_obj )
       self.construct_data_structures()
       self.handlers = common_obj.handlers
      
       
       
     
          
   
   def execute_15_minutes(self):
       web_display = {}
       error_hash  = {}
       self.handlers["SYSTEM_STATUS"].hset(self.subsystem_name,True)
      
       for i in self.processors:
           temp_web = {}
           temp_error = {}
           for j in self.containers[i]:           
               temp_web[j] = self.common_obj.general_hash_iterator(self.subsystem_name,self.analyize_web_display,self.watch_handlers[i]["WEB_DISPLAY_DICTIONARY"])
               temp_error[j]  = self.common_obj.general_hash_iterator(self.subsystem_name,self.analyize_error_display,self.watch_handlers[i]["ERROR_HASH"])
           web_display[i] = temp_web
           error_hash[i] = temp_error
       
       status = True
       data = self.handlers["MONITORING_DATA"].hget(self.subsystem_name)       
       if data == None:
          self.handlers["MONITORING_DATA"].hset(self.subsystem_name,[web_display,error_hash]) 
          return          
 
       status = True
       for i in self.processors:
           for j in self.containers[i]
               status = status and self.common_obj.detect_new_alert(self.subsystem_name,data[0][i][j],web_display[i][j])
               status = status and self.common_obj.check_for_error_flag(self.subsystem_name,error_hash[i][j])      
       print("status",status)
       if status == False: # change is monitoring status
           print("log alert")
           #self.common_obj.log_alert([web_display,error_hash])
           
       self.handlers["MONITORING_DATA"].hset(self.subsystem_name,[web_display,error_hash])  
       print("MONITORING_DATA",self.handlers["MONITORING_DATA"].hget(self.subsystem_name))
       print("SYSTEM_STATUS",self.handlers["SYSTEM_STATUS"].hget(self.subsystem_name))


   def construct_data_structures(self):
       
       self.processors = self.common_obj.find_processors()
       self.containers = self.common_obj.find_containers()
       
       search_list = ["PACKAGE"]
       data_structures = ["WEB_DISPLAY_DICTIONARY","ERROR_HASH"]
       self.watch_handlers = self.common_obj.generate_hash_structures_with_processor_container(self.processors,search_list,data_structures)
       


   def analyize_web_display( self,data):
       print("web_display",data)
       exit()
       print("WEB_DISPLAY_DICTIONARY",data)
       if data["error"] == False:
          flag  = True
       else:
          flag = False
       return [True,flag, json.dumps(data)]
       


   def analyize_error_display( self,data):
       print("error_display",data)
       exit()
       print("ERROR_DISPLAY",data)
       if 'time' not in data:
          return [False,True,json.dumps("")]
       ref_time = time.time()-15*60  # 15 minutes in past
       if data["time"] > ref_time:
          store_flag = True
          flag = False
          field = data
          return [True,False,json.dumps(field)]
       else:
          return [False,True,json.dumps("")]


           
   
