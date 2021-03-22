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
from common_tools.Pattern_tools_py3.factories.iterators_py3 import pattern_iter_strip_list_dict

class Control_Nodes(object):
   def __init__(self,config_file):
   
       self.site_data = get_site_data("/mnt/ssd/site_config/redis_server.json")
       self.qs =  Query_Support( self.site_data )
      
      
       self.system_error_logging = System_Error_Logging(self.qs,"Site_Control",self.site_data)
  
  
       
       #
       # Find all nodes
       #
       search_list = [ "PROCESSOR" ]
       nodes = common_qs_search(self.site_data,self.qs,search_list)
       print("processor_nodes",nodes)
       self.monitoring_nodes = pattern_iter_strip_list_dict(nodes,"name")
       print(self.monitoring_nodes)
       self.command_queue = self.find_command_queue()
       print(self.command_queue)
       quit()
     
    
   def find_command_queue(self):
       command_queue = {}
       for i in self.monitoring_nodes:
            search_list = [ ["PROCESSOR" ,i   ] ,"NODE_SYSTEM", "DOCKER_CONTROL" ]
            handlers = construct_all_handlers(self.site_data,self.qs,search_list,rpc_client=None,field_list=["DOCKER_COMMAND_QUEUE"])
            command_queue[i] = handlers["DOCKER_COMMAND_QUEUE"]
       return command_queue


bc.add_header_node("SITE_CONTROL","SITE_CONTROL",properties= properties) 
   
    cd.construct_package("SITE_CONTROL")
    cd.add_job_queue("SYSTEM_COMMAND_QUEUE",1)
    cd.add_single_element("SYSTEM_STATE")
    
    
    cd.add_job_queue("WEB_COMMAND_QUEUE",1)#used
    cd.add_redis_stream("ERROR_STREAM") #
    cd.add_hash("ERROR_HASH") #
    cd.add_hash("WEB_DISPLAY_DICTIONARY")#
    cd.close_package_contruction()
    
    cd.construct_package("DOCKER_CONTROL")

    cd.add_job_queue("DOCKER_COMMAND_QUEUE",1)
    cd.add_hash("DOCKER_DISPLAY_DICTIONARY")

if __name__ == "__main__":
   
   cf = CF_Base_Interpreter()
    #
    #
    # Read Boot File
    # expand json file
    # 
 
 
   config_file = "/mnt/ssd/site_config/redis_server.json"
   system_control =  Control_Nodes(config_file)
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