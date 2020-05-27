from op_monitor_lib.common_class_py3 import Common_Class
import json
import time


class Processor_Utilization(Common_Class):
  
   def __init__(self, subsystem_name,common_obj ):
       Common_Class.__init__(self,subsystem_name,common_obj )
       self.construct_data_structures()
       self.handlers = common_obj.handlers
       self.handlers["SYSTEM_STATUS"].hset(self.subsystem_name,True)
       
       
     
   def execute_day(self):
       print("execute day")  
       self.handlers["SYSTEM_STATUS"].hset(self.subsystem_name,True) # new day rollover
       new_data = {} 
       for i in self.processors:
          temp = {}
          temp["ram"]         = self.common_obj.general_stream_handler(self.analyize_ram,self.watch_handlers[i]["RAM"],duration=self.common_obj.one_day, )# all keys
          temp["disk_space"]  = self.common_obj.general_stream_handler(self.analyize_disk_space,self.watch_handlers[i]["DISK_SPACE"],duration=self.common_obj.one_day)# all keys
          temp["free_cpu"]    = self.common_obj.general_stream_handler( self.analyize_free_cpu,self.watch_handlers[i]["FREE_CPU"],duration=self.common_obj.one_day)# all keys
          temp["temperature"] = self.common_obj.general_stream_handler(self.analyize_temperature,self.watch_handlers[i]["TEMPERATURE"],duration=self.common_obj.one_day)# all keys
          temp["edev"] = self.common_obj.general_stream_handler(self.analyize_edev,self.watch_handlers[i]["EDEV"],duration=self.common_obj.one_day)
          error_count = self.common_obj.count_errors(temp)
          new_data[i] = [error_count,temp]
       
       self.compare_and_log_data(new_data)
       
   
 
   
   def compare_and_log_data(self,new_data):   
       
       #print("new_data",new_data)
       
       
       ref_total_data = self.handlers["MONITORING_DATA"].hget(self.subsystem_name)
       #print("ref_total_data",ref_total_data)
       if ref_total_data == None:
          ref_total_data = new_data
       status = True   
       for i in self.processors:
           
           ref_data = ref_total_data[i]
           new_ref  = new_data[i]
           
           if ref_data == None:
              print("continue",i)
           else:   
               status = status and self.common_obj.detect_new_alert(self.subsystem_name,new_ref,ref_data)
              
                   
       self.handlers["MONITORING_DATA"].hset(self.subsystem_name,new_data)   
 
       if status == False: # change is monitoring status
           print("log alert")
           self.common_obj.log_alert(self.subsystem_name,new_data)
           
     
  

   def construct_data_structures(self):
       
       self.processors = self.common_obj.find_processors()
       search_list = [["PACKAGE","SYSTEM_MONITORING"]]
       data_structures = ["FREE_CPU","RAM","DISK_SPACE","TEMPERATURE","EDEV"]
       self.watch_handlers = self.common_obj.generate_structures_with_processor(self.processors,search_list,data_structures,hash_flag = False)
      
              


   def analyize_free_cpu( self,data):
       
       filter_data = self.common_obj.filter_stream_values(["%idle"],data)
   
       stat_data = self.common_obj.determine_statistics(filter_data["%idle"])
       if stat_data["mean"] > .7:
          status = True
       else:
          status = False
       return_value = [status,json.dumps(stat_data)]
       
       return return_value
      
       
       
   def analyize_ram( self,data):
    
       filter_data = self.common_obj.filter_stream_values(['MemTotal','MemAvailable'],data)
       stat_data_total = self.common_obj.determine_statistics(filter_data['MemTotal'])
       stat_data_available = self.common_obj.determine_statistics(filter_data['MemAvailable'])
       #print(stat_data_total)
       #print(stat_data_available)
       ratio = stat_data_available["min"]/stat_data_total["mean"]
       #print(ratio)
       if ratio > .5 :
          status = True
       else:
          status = False
       stat_data_available["total"] = stat_data_total["mean"]   
       return_value = [status,json.dumps(stat_data_available)]
       #print(return_value)
       return return_value
         

   def analyize_disk_space( self,data):
       #print("disk_space",data[0].keys())
       field_keys = []
       keys = data[0].keys()
       if '/dev/root' in keys:
           field_keys.append('/dev/root')
       if "/dev/sda" in keys:
           field_keys.append('/dev/sda')
       if "/dev/sda1" in keys:
           field_keys.append('sda1')  

       #print('field_keys',field_keys)
       filter_data = self.common_obj.filter_stream_values(field_keys,data)
       stat_data = {}
       status = True
       for i in field_keys:
           stat_data[i] = self.common_obj.determine_statistics(filter_data[i]) 
           if stat_data[i]["max"] > .5:
              status = False
              
       return_value = [status,json.dumps(stat_data)]
       return return_value

      
   def analyize_temperature( self,data):
       filter_data = self.common_obj.filter_stream_values(['TEMP_F'],data)
       stat_data = self.common_obj.determine_statistics(filter_data['TEMP_F'])
       status = True
       if stat_data["mean"] > 130.:
          status = False
       if stat_data["max"] > 140.:
          status = False      
       return_value = [status,json.dumps(stat_data)]
    
       return return_value

   def analyize_edev( self,data):
     
       interfaces = data[0].keys()
       stat_data = {}
       status = True
       for i in interfaces:
           filter_data = self.common_obj.filter_stream_values([i],data)
           stat_data[i] = self.common_obj.determine_statistics(filter_data[i])
           if stat_data[i]["max"] > 0:
               status = False
               
       return_value = [status, json.dumps(stat_data)]
       return return_value       

           
   




           
   



           
   




           
   
