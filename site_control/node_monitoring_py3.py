'''

init startup

    find all nodes in system
    find node which is master
    find all containers in the system
    containers maped to node
    


node command is a web quueue queue server

function update containers images
function reboot node system
'''
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


class Control_Nodes(object):
   def __init__(self):
   
       self.site_data = get_site_data("/mnt/ssd/site_config/redis_server.json")
       qs =  Query_Support( self.site_data )
      
      
       self.system_error_logging = System_Error_Logging(qs,"Site_Control",self.site_data)
  
  
       #
       #  Find all nodes
       #  Identify the master node
       print(self.site_data)   
       local_node = self.site_data['local_node' ]
       print("local_node",local_node)
       #
       # Find all processors
       #
       search_list = [ "PROCESSOR"  ,"PROCESSOR"]
       processor_nodes = common_qs_search(self.site_data,qs,search_list)
       print("processor_nodes",processor_nodes)
       self.nodes = {}
       self.node_containers = {}
       self.site_containers = {}
       for i in processor_nodes:

          name = i{'name']
          containers = i['containers']
          services = i['services']
          self.processor_nodes[name] = {"master_node":False,"containers":containers,"services":services}
          if name == local_node:
              self.processor[name]["master_node"] = True
          self.node_containers[name] = containers
          self.node_containers[name].extend(services)       
    





if __name__ == "__main__":
   
   cf = CF_Base_Interpreter()
    #
    #
    # Read Boot File
    # expand json file
    # 
 
 
 
   system_control =  Control_Nodes()
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