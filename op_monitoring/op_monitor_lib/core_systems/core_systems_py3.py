
# core subsystems
#    pod_processes
#    docker container
#    docker process
#    cpu utlization



from .pod_processes_py3 import Pod_Processes
from .docker_containers_py3 import Docker_Containers
class Generate_Core_Monitoring(object):

   def __init__(self,common_functions):
       self.subsystems = {}
       self.subsystem_order = []
       self.common_functions = common_functions
       subsystem_list = [ 
                          ["pod_processes",Pod_Processes],
                          ["docker_containers",Docker_Containers]                          
       
                         ]   
       self.add_subsystems(subsystem_list)
       #
       #
       #
    
       
   def add_subsystems(self,subsystem_list):
       for i in subsystem_list:
            subsystem_name = i[0]
            class_instance = i[1]
            self.subsystem_order.append(subsystem_name)
            self.subsystems[subsystem_name]= class_instance(subsystem_name,self.common_functions)       
       
   def execute_minute(self):
       for i in self.subsystem_order:
           self.subsystems[i].execute_minute()  

   def execute_15_minutes(self):
       for i in self.subsystem_order:
           self.subsystems[i].execute_15_minutes()  
   
   def execute_hour(self):
       for i in self.subsystem_order:
           self.subsystems[i].execute_hour()  

   def execute_alarm_daily_count(self):
       for i in self.subsystem_order:
           self.subsystems[i].execute_alarm_daily_count()  
   