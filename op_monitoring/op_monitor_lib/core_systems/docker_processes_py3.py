from op_monitor_lib.common_class_py3 import Common_Class
import json
import time

class Docker_Processes(Common_Class):
  
   def __init__(self, subsystem_name,common_obj ):
       Common_Class.__init__(self,subsystem_name,common_obj )
       self.construct_data_structures()
       self.handlers = common_obj.handlers
       
      
       
       

             
          
   
   def execute_15_minutes(self):
       print("execute_15_minutes")
       web_display = {}
       error_hash  = {}
       self.handlers["SYSTEM_STATUS"].hset(self.subsystem_name,True)
       new_data = {}
       for i in self.processors:
           container_data = {}
           for j in self.containers[i]: 
               web_display = self.common_obj.general_hash_iterator(self.subsystem_name,self.analyize_web_display,self.watch_handlers[i][j]["WEB_DISPLAY_DICTIONARY"])
               error_display  = self.common_obj.general_hash_iterator(self.subsystem_name,self.analyize_error_display,self.watch_handlers[i][j]["ERROR_HASH"])
               
               container_data[j] = {"web_display":web_display,"error_hash":error_display}
              

           new_data[i] = container_data

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
               for j in self.containers[i]:
                   if j == "op_monitor":
                      continue
                   ref_container = ref_data[j]
                   #print("ref_container",j,ref_container)
                   if ref_container == None:
                      print("continue",i)
                   else:
                      status = status and self.common_obj.detect_new_alert(self.subsystem_name,new_data[i][j]["web_display"],ref_container["web_display"])
                      status = status and self.common_obj.check_for_error_flag(self.subsystem_name,new_data[i][j]["error_hash"])

                
       print("status",status)
       if status == False: # change is monitoring status
           print("log alert")
           self.common_obj.log_alert(self.subsystem_name,new_data)
       self.handlers["MONITORING_DATA"].hset(self.subsystem_name,new_data)


   def construct_data_structures(self):
       
       self.processors = self.common_obj.find_processors()
       self.containers = self.common_obj.find_all_containters()
       data_structures = ["WEB_DISPLAY_DICTIONARY","ERROR_HASH"]
       self.watch_handlers = self.common_obj.generate_structures_with_processor_container(self.processors,data_structures)
       


   def analyize_web_display( self,data):
 
       #print("WEB_DISPLAY_DICTIONARY",data)
       if data["error"] == False:
          flag  = True
       else:
          flag = False
       return [True,flag, json.dumps(data)]
       


   def analyize_error_display( self,data):
       #print("error_display",data)

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


           
   
