from common_tools.redis_support_py3.graph_query_support_py3 import  Query_Support
from common_tools.redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
from docker_control.docker_interface_py3 import Docker_Interface
from common_tools.py_cf_new_py3.chain_flow_py3 import CF_Base_Interpreter
import json 
import time
import os
from common_tools.system_error_log_py3 import  System_Error_Logging

from common_tools.Pattern_tools_py3.builders.common_directors_py3 import construct_all_handlers
from common_tools.Pattern_tools_py3.factories.graph_search_py3 import common_qs_search
from common_tools.Pattern_tools_py3.factories.get_site_data_py3 import get_site_data


class Control_Containers(object):
   def __init__(self):
   
       self.site_data = get_site_data("/mnt/ssd/site_config/redis_server.json")
       qs =  Query_Support( self.site_data )
      
      
       self.system_error_logging = System_Error_Logging(qs,"Site_Control",self.site_data)
  
       search_list = [ ["SITE_CONTROL","SITE_CONTROL"] , "DOCKER_CONTROL" ]
       self.ds_handlers = construct_all_handlers(self.site_data,qs,search_list,rpc_client=None)
       self.ds_handlers["DOCKER_COMMAND_QUEUE"].delete_all()
       
       
 
       search_list = [ ["SITE","LACIMA_SITE"]]
       processor_nodes = common_qs_search(self.site_data,qs,search_list)
       processor_node = processor_nodes[0]
      
                                        
 
       
       
       self.starting_scripts = {}
       self.container_images = {}
       self.find_starting_scripts(qs,"CONTAINER",processor_node["services"],self.starting_scripts,self.container_images)
       self.find_starting_scripts(qs,"CONTAINER",processor_node["containers"],self.starting_scripts,self.container_images)
       
       

       services = set(processor_node["services"])
       containers = set(processor_node["containers"])
       containers_set = containers.union(services)
       self.container_list = list(self.container_images.keys())  
       
       
       
       self.docker_interface =  Docker_Interface()  
       self.verify_container_images()       
    
       
       self.setup_environment()
       self.check_for_allocated_containers()
       self.docker_performance_data_structures= {}
       for i in services:
           self.docker_performance_data_structures[i] = self.assemble_container_data_structures(qs,i)
           
           
   def assemble_container_data_structures(self,qs,container_name):
       print("container name",container_name)
       search_list = [  ["CONTAINER",container_name], "DATA_STRUCTURES" ]
       return_value = construct_all_handlers(self.site_data,qs,search_list,rpc_client=None)
       
       return return_value
        
   def find_starting_scripts(self,qs,relationship_label, entry_list,startup_scripts,container_images):
       
       for entry in entry_list:
           
           search_list = [ [relationship_label,entry] ]
           service_sets = common_qs_search(self.site_data,qs,search_list)
          
           service_node = service_sets[0]
           startup_scripts[entry] = service_node["startup_command"]
           container_images[entry] = service_node["container_image"]
          
   


          
 
       
       
   def verify_container_images(self):
       self.system_images = self.docker_interface.images()
       required_images = self.container_images.values()

       for i in required_images:
           if i not in self.system_images:
                raise # images should be present?
                self.docker_interface.pull(i)
       
      
       
       


   



      
   def setup_environment(self):
        for i in self.container_list:
           self.ds_handlers["DOCKER_DISPLAY_DICTIONARY"].hset(i,{"name":i,"enabled":True,"active":True,"error":False,"defined":True})


   
  
   def check_for_allocated_containers(self):
       running_containers = self.docker_interface.containers_ls_runing()

       for i in self.container_list:
           
           if i not in running_containers:
               self.docker_interface.container_up(i,self.starting_scripts[i])
               
               
               
           
   
   def monitor(self,*unused):  # try to start stop container
       running_containers = self.docker_interface.containers_ls_runing()
       for i in self.container_list:
          
           temp = self.ds_handlers["DOCKER_DISPLAY_DICTIONARY"].hget(i)
           if i not in running_containers:

               if (temp["defined"] == True) and (temp["enabled"] == True) :
                 
                  if temp["active"] == True:
                     
                     temp["active"] = False
                     self.ds_handlers["DOCKER_DISPLAY_DICTIONARY"].hset(i,{"name":i,"enabled":True,"active":False,"error":True,"defined":True})
                 
                  self.docker_interface.container_up(i,self.starting_scripts[i])
           else:
               self.ds_handlers["DOCKER_DISPLAY_DICTIONARY"].hset(i,{"name":i,"enabled":True,"active":True,"error":False,"defined":True})  
           

 
   def change_container_state(self,data):
       
       for container,new_state in data.items():
           
           old_state = self.ds_handlers["DOCKER_DISPLAY_DICTIONARY"].hget(container)
        
           try:
               print(container,new_state,old_state)
               if new_state["enabled"] == False:
                  print("stop")
                  if (old_state["enabled"] == True)or(old_state["active"] == True):
                      old_state["enabled"] = False 
                      old_state["active"] = False
                      self.docker_interface.container_stop(container) # try to start container
                      self.ds_handlers["DOCKER_DISPLAY_DICTIONARY"].hset(container,old_state)
               else:
                 print("start")
                 if (old_state["enabled"] == False)or(old_state["active"] == False):
                     old_state["enabled"] = True 
                     old_state["active"] = True
                     print(self.starting_scripts[container])
                     self.docker_interface.container_up(container,self.starting_scripts[container]) # start container                         
                     self.ds_handlers["DOCKER_DISPLAY_DICTIONARY"].hset(container,old_state)
                 
                  
           except:
                   raise    


   def upgrade_a_container(self,container):
          
               try:
                  
                  self.docker_interface.container_stop(container) # stop container
                  self.docker_interface.container_rm(container) # rm container
                  container_image = self.container_images[container]
                  print("container_image",container_image)
                  print("pull",self.docker_interface.pull(container_image)) 
                  self.docker_interface.container_up(container,self.starting_scripts[container]) # start container\
                  
               except:
                   raise

                   
   def upgrade_containers(self,data):
       print("upgrade container")
       for container,upgrade in data.items():
          if upgrade["enabled"] == True:
              self.upgrade_a_container(container)

   def upgrade_all_containers(self,data):
       print("upgrade all")
       for container,upgrade in data.items():
           self.upgrade_a_container(container)    

   def  prune_container(self):
        self.docker_interface.prune()   
           
   def process_web_queue( self, *unused ):
       data = self.ds_handlers["DOCKER_COMMAND_QUEUE"].pop()
      
       if data[0] == True :
           print("receivd message",data[1])
           
           if data[1]["command"] == 1:  #"CONTAINER_START/STOP":
              self.change_container_state(data[1]["items"])           
           if data[1]["command"] == 2 : #"UPGRADE":
              self.upgrade_containers(data[1]["items"])    
           if data[1]["command"] == 3 : #"UPGRADE_ALL":
              self.upgrade_all_containers(data[1]["items"])
           if data[1]["command"] == 4 : #"Prune Images":
              self.prune_container()          
           if data[1]["command"] == 5: #"REBOOT":
              print("starting to reboot")
              time.sleep(15)         
              os.system("reboot")
          

   def measure_container_processes(self,*args):
       
       for i in self.container_list:
           
           self.measure_ps_parameter(i,"%CPU","PROCESS_CPU")
           self.measure_ps_parameter(i,"VSZ","PROCESS_VSZ"),
           self.measure_ps_parameter(i,"RSS","PROCESS_RSS")

   def measure_ps_parameter( self , container_name,field_name,stream_name ):
       handlers = self.docker_performance_data_structures[container_name]
       headers = [ "USER","PID","%CPU","%MEM","VSZ","RSS","TTY","STAT","START","TIME","COMMAND", "PARAMETER1", "PARAMETER2" ]
       f = os.popen("docker top "+container_name+ "  -aux | grep python")
       data = f.read()
       f.close()
       
       lines = data.split("\n")
       return_value = {}
       for i in range(0,len(lines)):
           fields = lines[i].split()
           temp_value = {}
           if len(fields) <= len(headers):
               for i in range(0,len(fields)):
                   temp_value[headers[i]] = fields[i]
               
               if "PARAMETER1" in temp_value:
                   if temp_value["COMMAND"] == "python":
                       key = temp_value["PARAMETER1"]
                       return_value[key] = temp_value[field_name]
                       
                       
       #print(return_value,container_name)   
       print("data",return_value)
       handlers[stream_name].push( data = return_value,local_node = container_name )
             




      
   def add_chains(self,cf):

       
       cf.define_chain("monitor_web_command_queue", True)
       cf.insert.wait_event_count( event = "TIME_TICK", count = 1)
       cf.insert.one_step(self.process_web_queue)
       cf.insert.reset()
       
       cf.define_chain("monitor_active_containers",True)
       #cf.insert.log("starting docker monitoring")     
       cf.insert.one_step(self.monitor)
       cf.insert.wait_event_count( event = "TIME_TICK",count = 10)
       cf.insert.reset()
       

       cf.define_chain("process_monitor", True)
       cf.insert.log("starting docker process measurements")     
       cf.insert.one_step(self.measure_container_processes)
       cf.insert.log("ending docker measurements")
       cf.insert.wait_event_count( event = "MINUTES",count = 5)
       cf.insert.reset()

 

if __name__ == "__main__":
   
   cf = CF_Base_Interpreter()
    #
    #
    # Read Boot File
    # expand json file
    # 
 
 
 
   system_control =  Control_Containers()
   cf = CF_Base_Interpreter()
   system_control.add_chains(cf)
   #
   # Executing chains
   #
   try: 
       cf.execute()
   except:
       
       raise
else:
   pass

    