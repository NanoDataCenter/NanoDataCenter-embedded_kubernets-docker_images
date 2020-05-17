from op_monitor_lib.common_class_py3 import Common_Class
import json
import time

class Pod_Processes(Common_Class):
  
   def __init__(self, subsystem_name,common_obj ):
       Common_Class.__init__(self,subsystem_name,common_obj )
       self.construct_data_structures()
       self.handlers = common_obj.handlers
       
       
       
     
          
   
   def execute_15_minutes(self):
       print("execute_15_minutes")
       web_display = {}
       error_hash  = {}
       
      
       
       new_data = {} 
       for i in self.processors:
           #print(i)
           web_display = self.common_obj.general_hash_iterator(i,self.analyize_web_display,self.watch_handlers[i]["WEB_DISPLAY_DICTIONARY"])
           error_hash  = self.common_obj.general_hash_iterator(i,self.analyize_error_display,self.watch_handlers[i]["ERROR_HASH"])
           new_data[i] = {"web_display":web_display,"error_hash":error_hash}
      
           
       
       
       
       #print("new_data",new_data)
       status = True
       self.handlers["SYSTEM_STATUS"].hset(self.subsystem_name,True)
       ref_total_data = self.handlers["MONITORING_DATA"].hget(self.subsystem_name)
       if ref_total_data == None:
          ref_total_data = new_data
       for i in self.processors:
           
           ref_data = ref_total_data[i]
           #print("ref_data",i,ref_data)
           if ref_data == None:
              print("continue",i)
           else:   
               status = status and self.common_obj.detect_new_alert(self.subsystem_name,new_data[i]["web_display"],ref_data["web_display"])
               status = status and self.common_obj.check_for_error_flag(self.subsystem_name,new_data[i]["error_hash"])
       self.handlers["MONITORING_DATA"].hset(self.subsystem_name,new_data)
       
       print("status",status)
       if status == False: # change is monitoring status
           print("log alert")
           self.common_obj.log_alert(self.subsystem_name,new_data)
           
         


   def construct_data_structures(self):
       
       self.processors = self.common_obj.find_processors()
       search_list = ["NODE_PROCESSES","PACKAGE"]
       data_structures = ["WEB_DISPLAY_DICTIONARY","ERROR_HASH"]
       self.watch_handlers = self.common_obj.generate_structures_with_processor(self.processors,search_list,data_structures)
      
              


   def analyize_web_display( self,data):
       #print("WEB_DISPLAY_DICTIONARY",data)
       if data["error"] == False:
          flag  = True
       else:
          flag = False
       return [True,flag, json.dumps(data)]
       


   def analyize_error_display( self,data):
       #print("ERROR_DISPLAY",data)
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


           
   
