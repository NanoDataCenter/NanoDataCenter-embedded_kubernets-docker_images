
import datetime
import msgpack
from op_monitor_lib.common_class_py3 import Common_Class
import json
import time


class Rpi_Mosquitto_Client_Monitor(Common_Class):
  
   def __init__(self, subsystem_name,common_obj ):
       Common_Class.__init__(self,subsystem_name,common_obj )
       
       self.handlers = common_obj.handlers
       self.setup_test_client()
       
       
       
       
       
     
   def execute_15_minutes(self):
       print("execute_15_minutes")  
       self.handlers["SYSTEM_STATUS"].hset(self.subsystem_name,True)
       new_data = self.do_mosquitto_test()
       
  
       
       
       self.compare_and_log_data(new_data)
       
  
   def compare_and_log_data(self,new_data):   
       
       #print("new_data",new_data)
       
       
       ref_total_data = self.handlers["MONITORING_DATA"].hget(self.subsystem_name)
       #print("ref_total_data",ref_total_data)
       if ref_total_data == None:
          ref_total_data = new_data
       status = True   
 
       status = status and self.common_obj.detect_new_alert(self.subsystem_name,new_data,ref_total_data)
              
       self.handlers["MONITORING_DATA"].hset(self.subsystem_name,new_data)   
 
       if status == False: # change is monitoring status
           print("log alert")
           self.common_obj.log_alert(self.subsystem_name,new_data)
  
  

   def setup_test_client(self):
       search_list = [["MQTT_DEVICES","MQTT_DEVICES"],["PACKAGE","MQTT_DEVICES_DATA"]]
       data_structures = ["MQTT_CONTACT_LOG"]
       self.contact_log = self.common_obj.generate_structures_without_processor(search_list,data_structures,hash_flag = True)       
       print("contact_log",self.contact_log)
      
   
   def do_mosquitto_test(self):
       error_count = 0
       return_data = {}
       for i in self.contact_log["MQTT_CONTACT_LOG"].hkeys():
           data = self.contact_log["MQTT_CONTACT_LOG"].hget(i)
           if data["status"]  == True:
              return_data[i] = [True,json.dumps("success")]
           else:
              return_data[i] = [False,json.dumps("failure")]
       return [error_count,return_data]
         
       
       
       

