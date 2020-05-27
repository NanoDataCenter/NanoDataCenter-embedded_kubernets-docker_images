from op_monitor_lib.common_class_py3 import Common_Class
import json
import time


class Redis_Monitor(Common_Class):
  
   def __init__(self, subsystem_name,common_obj ):
       Common_Class.__init__(self,subsystem_name,common_obj )
       self.construct_data_structures()
       self.handlers = common_obj.handlers
       self.handlers["SYSTEM_STATUS"].hset(self.subsystem_name,True)
       
       
     
   def execute_day(self):
       print("execute day")  
       self.handlers["SYSTEM_STATUS"].hset(self.subsystem_name,True) # new day rollover
       temp = {} 
       temp["CLIENTS"]         = self.common_obj.general_stream_handler(self.analyize_client,self.watch_handlers["REDIS_MONITOR_CLIENT_STREAM"],duration=self.common_obj.one_day, )# all keys
       temp["MEMORY"]  = self.common_obj.general_stream_handler(self.analyize_memory,self.watch_handlers["REDIS_MONITOR_MEMORY_STREAM"],duration=self.common_obj.one_day)# all keys
       error_count = self.common_obj.count_errors(temp)
       new_data = [error_count,temp]
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
           
 
  

   def construct_data_structures(self):
       
       
       search_list = [["PACKAGE","REDIS_MONITORING"]]
       data_structures = ["REDIS_MONITOR_CLIENT_STREAM","REDIS_MONITOR_MEMORY_STREAM"]
       self.watch_handlers = self.common_obj.generate_structures_without_processor(search_list,data_structures,hash_flag = False)
      
 
              


   def analyize_client( self,data):
       field_keys = ['connected_clients','blocked_clients']
       filter_data = self.common_obj.filter_stream_values(field_keys,data)
       stat_data = {}
       status = True
     
       stat_data['connected_clients'] = self.common_obj.determine_statistics(filter_data['connected_clients']) 
       if stat_data['connected_clients']["max"] > 200:
          status == False
       stat_data['blocked_clients'] = self.common_obj.determine_statistics(filter_data['blocked_clients']) 
       if stat_data['blocked_clients']["max"] > 200:
          status == False       
             
       return_value = [status,json.dumps(stat_data)]
 
       return return_value


                

   def analyize_memory( self,data):

       field_keys = ['used_memory','used_memory_rss']
       filter_data = self.common_obj.filter_stream_values(field_keys,data)
       stat_data = {}
       status = True
     
       stat_data['used_memory'] = self.common_obj.determine_statistics(filter_data['used_memory']) 
       if stat_data['used_memory']["max"] > 60000000:
          status == False
       stat_data['used_memory_rss'] = self.common_obj.determine_statistics(filter_data['used_memory_rss']) 
       if stat_data['used_memory_rss']["max"]  > 60000000:
          status == False       
             
       return_value = [status,json.dumps(stat_data)]
 
       return return_value
           
   



           
   




           
   
