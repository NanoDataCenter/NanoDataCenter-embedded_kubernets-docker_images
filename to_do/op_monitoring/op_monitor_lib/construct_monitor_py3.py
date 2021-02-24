
from .core_systems.core_systems_py3  import Generate_Core_Monitoring
from .servers.monitor_redis_py3   import Redis_Monitor
from .servers.monitor_sql_server_py3   import SQLITE_Monitor
from .servers.monitor_block_chain_py3  import Block_Chain_Monitor
from .servers.monitor_rpi_mosquitto_py3 import Rpi_Mosquitto_Monitor
from .servers.monitor_rpi_mosquitto_clients_py3  import Rpi_Mosquitto_Client_Monitor
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
           elif i == "MONITOR_REDIS":
               print(i)
               self.monitors[i] = Redis_Monitor("MONITOR_REDIS",self.common_functions)
              
           elif i == "MONITOR_SQLITE":
                print(i)
                self.monitors[i] = SQLITE_Monitor("MONITOR_SQLITE",self.common_functions)

           elif i == "MONITOR_BLOCK_CHAIN":
                print(i)
                self.monitors[i] = Block_Chain_Monitor("MONITOR_BLOCK_CHAIN",self.common_functions) 

           elif i == "MONITOR_RPI_MOSQUITTO":
                print(i)
                self.monitors[i] = Rpi_Mosquitto_Monitor("MONITOR_RPI_MOSQUITTO",self.common_functions)  
                
           elif i == "MONITOR_RPI_MOSQUITTO_CLIENTS":
                print(i)
                self.monitors[i] = Rpi_Mosquitto_Client_Monitor("MONITOR_RPI_MOSQUITTO_CLIENTS",self.common_functions)                  
           else:
               raise ValueError(i)        
          

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
   