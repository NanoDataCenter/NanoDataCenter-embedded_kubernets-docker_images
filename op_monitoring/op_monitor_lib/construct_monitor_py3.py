
from .core_systems.core_systems_py3  import Generate_Core_Monitoring
#from .monitor_redis_py3              import Generate_Monitor_Redis
from .common_functions_py3   import Common_Functions
class Construct_Monitors(object):

   def __init__(self,site_data,qs,monitoring_list,handlers):
      self.common_functions = Common_Functions(site_data,qs,handlers)
      
      self.monitoring_list = monitoring_list
      self.monitors = {}
      for i in self.monitoring_list:
 
           if i == "CORE_OPS":
               print(i)
               self.monitors[i] = Generate_Core_Monitoring(self.common_functions)
           #if i == "monitor_redis":
           #    print(i)
           #    self.monitors[i] = Generate_Monitor_Redis(self.common_functions)
              
          
                  
           else:
               raise         
          

   def execute_minute(self):
       for i in self.monitoring_list:
           self.monitors[i].execute_minute()  

   def execute_15_minutes(self):
       for i in self.monitoring_list:
           self.monitors[i].execute_15_minutes()  
   
   def execute_hour(self):
       for i in self.monitoring_list:
           self.monitors[i].execute_hour()  
           
   def execute_day(self):
       for i in self.monitoring_list:
           self.monitors[i].execute_day()  


   def execute_alarm_daily_count(self):
       for i in self.monitoring_list:
           self.monitors[i].execute_alarm_daily_count()  
   