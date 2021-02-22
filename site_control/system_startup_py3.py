# this file is 





import time
import json
import redis
import subprocess
from subprocess import Popen, check_output
import shlex
import os
from py_cf_new_py3.chain_flow_py3 import CF_Base_Interpreter
from redis_support_py3.graph_query_support_py3 import  Query_Support
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
from docker_control.docker_interface_py3 import Docker_Interface
import msgpack
import pickle
import zlib



container_run_script = "docker run -d   --name redis -p 6379:6379 --mount type=bind,source=/mnt/ssd/redis,target=/data  " 
container_run_script = container_run_script + " --mount type=bind,source=/mnt/ssd/redis/config/redis.conf,target=/usr/local/etc/redis/redis.conf redis"


    
   

def wait_for_redis_db(site_data):
   
    while True:
        try:
            redis_handle = redis.StrictRedis( host = site_data["host"] , port=site_data["port"], db=site_data["graph_db"])
            temp = redis_handle.ping()
            #print(temp)
            if temp == True:
              
              
               return
            else:
               raise
        except:
           print("exception")
           time.sleep(10)




def find_system_status(site_data,qs):
    
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
    query_list = qs.add_match_relationship( query_list,relationship="SYSTEM_CONTROL",label="SYSTEM_CONTROL" )
    query_list = qs.add_match_terminal( query_list, relationship = "PACKAGE", label = "SYSTEM_CONTROL" )
    package_sets, package_nodes = qs.match_list(query_list)
    package_node = package_nodes[0]
    generate_handlers = Generate_Handlers(package_node,qs)
    data_structures = package_node["data_structures"]
    system_state = generate_handlers.construct_single_element(data_structures["SYSTEM_STATE"])
    return system_state




def find_processors(site_data,qs):
    
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
    query_list = qs.add_match_terminal( query_list,relationship="NODE_PROCESSES" )
    node_sets, nodes = qs.match_list(query_list)
    return_value = []
    for i in nodes:
        return_value.append(i["name"])
    return return_value




def find_node_state(site_data,processor_name,qs):
   
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )
    query_list = qs.add_match_relationship( query_list,relationship="NODE_PROCESSES",label = processor_name )
    query_list = qs.add_match_terminal( query_list, relationship = "PACKAGE", label = "DATA_STRUCTURES" )
    package_sets, package_nodes = qs.match_list(query_list)
    return_value = {}
    package_node = package_nodes[0]
    generate_handlers = Generate_Handlers(package_node,qs)
    data_structures = package_node["data_structures"]
    return_value = generate_handlers.construct_single_element(data_structures["NODE_STATE"])
    
    return return_value

def find_all_node_states(site_data,qs):
    
    processors = find_processors(site_data,qs)
    return_value = {}
    for i in processors:
        return_value[i] = find_node_state(site_data,i,qs)
    
    return return_value

  
    
node_states = []    
system_state = None   
    
def set_status_handler(status_code):
    global system_state
    system_state.set(status_code)



def check_status_handler(status_code_list):
    global node_states
    return_value = True
    for i,handler in node_states.items():
        response = handler.get()
        print("response",response)
        if response in status_code_list:
            pass
        else:
           return_value = False
   
    return return_value
      
def wait_for_response(status_code_list):
  
    loop_flag = True
    while True:
        time.sleep(1)
        print("checking")
        if check_status_handler(status_code_list):
            return

          
          
           
  

if __name__ == "__main__":
  
   docker_control = Docker_Interface()
   
   file_handle = open("/mnt/ssd/site_config/redis_server.json")
   data = file_handle.read()
   file_handle.close()
   site_data = json.loads(data)

   
   if 'master' not in site_data:
       if site_data["master"] != True:  
           while True:               ### the node is not a system control node
              print("not master")
              time.sleep(10)
          
   



   docker_control.container_up("redis",container_run_script) 
      
   wait_for_redis_db(site_data)
   
   qs = Query_Support( site_data )
   system_state = find_system_status(site_data,qs)
   
   processors = find_processors(site_data,qs)
   node_states = find_all_node_states(site_data,qs)
   
   print("system_state",system_state)
   print("node_states",node_states)
   
   

   system_state.set("STARTUP")
   wait_for_response(["STARTUP_CONFIRMED","SERVICES_LOADED","OPERATIONAL"])
 
   system_state.set("LOAD_SERVICE_CONTAINERS")
   wait_for_response(["SERVICES_LOADED","OPERATIONAL"])
 
   system_state.set("LOAD_APP_CONTAINERS")
   wait_for_response(["OPERATIONAL"])  
   
   system_state.set("OPERATIONAL")
   
   print("made it here")
   quit()    
   
else:
   pass









 