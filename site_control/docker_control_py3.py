from redis_support_py3.graph_query_support_py3 import  Query_Support
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
from docker_control.docker_interface_py3 import Docker_Interface
from py_cf_new_py3.chain_flow_py3 import CF_Base_Interpreter
import json 
import time
import os

class Control_Containers(object):
   def __init__(self,site_data):
       self.site_data = site_data
       qs =  Query_Support( site_data )
      
       query_list = []
       query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
       query_list = qs.add_match_terminal( query_list,relationship="PROCESSOR",label=site_data["local_node"] )
       node_sets, processor_node = qs.match_list(query_list)
  
       services = {}
       
       
       
   
   
   
       query_list = []
       query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
       query_list = qs.add_match_relationship( query_list,relationship="PROCESSOR",label=site_data["local_node"] )
       query_list = qs.add_match_relationship( query_list,relationship="NODE_SYSTEM")
       query_list = qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", label = "DOCKER_CONTROL" )
                                        
                                        
                                           
       package_sets, package_nodes = qs.match_list(query_list)  
   
     
   
       generate_handlers = Generate_Handlers(package_nodes[0],qs)      
       processor_node = processor_node[0]
       self.starting_scripts = {}
       self.container_images = {}
       self.find_starting_scripts(qs,"SERVICE",processor_node["services"],self.starting_scripts,self.container_images)
       self.find_starting_scripts(qs,"CONTAINER",processor_node["containers"],self.starting_scripts,self.container_images)
       
       #print(self.starting_scripts)
       #print(self.container_images)

       services = set(processor_node["services"])
       containers = set(processor_node["containers"])
       containers_set = containers.union(services)
       self.container_list = list(self.container_images.keys())  
       #print("container_list",self.container_list)
       
       self.docker_interface =  Docker_Interface()  
       self.verify_container_images()       
       package_node = package_nodes[0]
       data_structures = package_node["data_structures"]
       
       
       self.ds_handlers = {}
       self.ds_handlers["DOCKER_COMMAND_QUEUE"]   = generate_handlers.construct_job_queue_server(data_structures["DOCKER_COMMAND_QUEUE"])
       self.ds_handlers["DOCKER_DISPLAY_DICTIONARY"]   =  generate_handlers.construct_hash(data_structures["DOCKER_DISPLAY_DICTIONARY"])
       self.ds_handlers["DOCKER_COMMAND_QUEUE"].delete_all()
       
       self.setup_environment()
       self.check_for_allocated_containers()
       self.docker_performance_data_structures= {}
       for i in self.container_list:
           self.docker_performance_data_structures[i] = self.assemble_container_data_structures(qs,i)
       
   

   def find_image_name(self,image_object):
       try:
           temp = image_object.attrs["RepoTags"][0]
           #print(temp)
           return temp
       except:
          return "blank"  # place holder
          
   def strip_name(self,image_name):
       string_list = image_name.split(":")
       return string_list[0]


   def find_system_image_names(self):
          return list(map(self.strip_name,list( map(self.find_image_name,iter(self.docker_interface.images())))))

   def pull_missing_images(self, required_element):
       #print("required element",required_element)
       if required_element not in self.system_images:
           #print("should not happen")
           self.docker_interface.pull(required_element)
       #print(self.docker_interface.pull(required_element)) # for test
       return True
       
       
   def verify_container_images(self):
       self.system_images = self.find_system_image_names()
       required_images = list(self.container_images.values())
       print(list(map(self.pull_missing_images,required_images)))
      
       
       


   
   def assemble_container_data_structures(self,qs,container_name):
       
       query_list = []
       query_list = qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
       query_list = qs.add_match_relationship( query_list,relationship="PROCESSOR",label=self.site_data["local_node"] )
       query_list = qs.add_match_relationship( query_list,relationship="CONTAINER",label=container_name)
       query_list = qs.add_match_terminal( query_list, 
                                           relationship = "PACKAGE", label = "DATA_STRUCTURES" )

       package_sets, package_nodes =qs.match_list(query_list)  
      
 
           
       #print("package_nodes",package_nodes)
   
       generate_handlers = Generate_Handlers(package_nodes[0],qs)
       data_structures = package_nodes[0]["data_structures"]
      
       handlers = {}
       handlers["PROCESS_VSZ"]  = generate_handlers.construct_redis_stream_writer(data_structures["PROCESS_VSZ"])
       handlers["PROCESS_RSS"] = generate_handlers.construct_redis_stream_writer(data_structures["PROCESS_RSS"])
       handlers["PROCESS_CPU"]  = generate_handlers.construct_redis_stream_writer(data_structures["PROCESS_CPU"])
       return handlers  

   def find_starting_scripts(self,qs,relationship_label, entry_list,startup_scripts,container_images):
       
       for entry in entry_list:
           query_list = []
           query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
           query_list = qs.add_match_relationship( query_list,relationship="PROCESSOR",label=site_data["local_node"] )
           query_list = qs.add_match_terminal( query_list,relationship=relationship_label,label=entry )
           service_sets, service_nodes = qs.match_list(query_list)
           startup_scripts[entry] = service_nodes[0]["startup_command"]
           container_images[entry] = service_nodes[0]["container_image"]
          


      
   def setup_environment(self):
        for i in self.container_list:
           self.ds_handlers["DOCKER_DISPLAY_DICTIONARY"].hset(i,{"name":i,"enabled":True,"active":True,"error":False,"defined":True})


   
  
   def check_for_allocated_containers(self):
       running_containers = self.docker_interface.containers_ls_runing()

       for i in self.container_list:
           print(i)
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
   #time.sleep(20) # wait for mqtt server get started
   file_handle = open("/mnt/ssd/site_config/redis_server.json")
   data = file_handle.read()
   file_handle.close()
   site_data = json.loads(data)
  
 
 
   system_control =  Control_Containers(site_data)
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

    