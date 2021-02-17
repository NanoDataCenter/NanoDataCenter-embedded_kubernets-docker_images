# file build system
#
#  The purpose of this file is to load a system configuration
#  in the graphic data base
#

import json

import redis
from  build_configuration_py3 import Build_Configuration
from  construct_data_structures_py3 import Construct_Data_Structures
from   graph_modules_py3.cloud_site.site_definitions_py3 import Cloud_Site_Definitons
from   graph_modules_py3.lacima.site_definitions_py3 import LACIMA_Site_Definitons



from  graph_modules_py3.containers_py3.construct_container_py3 import Construct_Containers
from  graph_modules_py3.services_py3.construct_services_py3    import Construct_Services


 
   


def construct_site_definitions(bc,cd,services):
    properties = {}
    properties["services"] = services
    properties["command_list"] = [{"file":"docker_control_py3.py","restart":True},{"file":"redis_monitoring_py3.py","restart":True}]
    bc.add_header_node("SITE_CONTROL","SITE_CONTROL",properties= properties) 
   
    cd.construct_package("SITE_CONTROL")
    cd.add_job_queue("SYSTEM_COMMAND_QUEUE",1)
    cd.add_single_element("SYSTEM_STATE")
    cd.add_job_queue("WEB_COMMAND_QUEUE",1)
    cd.add_redis_stream("ERROR_STREAM",forward=True)
    cd.add_hash("ERROR_HASH")
    cd.add_hash("WEB_DISPLAY_DICTIONARY")
    cd.close_package_contruction()
 
 
    cd.construct_package("REDIS_MONITORING")  # redis monitoring
    cd.add_redis_stream("KEYS")
    cd.add_redis_stream("CLIENTS")
    cd.add_redis_stream("MEMORY")
    cd.add_redis_stream("REDIS_MONITOR_CALL_STREAM")
    cd.add_redis_stream("REDIS_MONITOR_CMD_TIME_STREAM")  
    cd.add_redis_stream("REDIS_MONITOR_SERVER_TIME")  
    cd.close_package_contruction()
    

       
    
    
    bc.end_header_node("SITE_CONTROL")
    
    bc.add_header_node("SYSTEM_MONITOR")
    cd.construct_package("SYSTEM_MONITOR")      
    #cd.add_managed_hash(self,name,fields,forward=False) perfored way to store field how to get field in system
    cd.add_hash("SYSTEM_STATUS")
    cd.add_hash("MONITORING_DATA")
    cd.add_redis_stream("SYSTEM_ALERTS")
    cd.add_redis_stream("SYSTEM_PUSHED_ALERTS")
    cd.close_package_contruction()
    bc.end_header_node("SYSTEM_MONITOR")

 
 

def construct_processor(name,containers,services):
    properties = {}
    properties["containers"] = containers
    properties["services"] = services
    bc.add_header_node("PROCESSOR",name,properties= properties) 
    

   
    properties = {}
    properties["command_list"] = [{"file":"pi_monitoring_py3.py","restart":True},{"file":"docker_control_py3.py","restart":True}]
    bc.add_header_node("NODE_SYSTEM",properties=properties)
    
    cd.construct_package("PROCESSOR_MONITORING")
    cd.add_redis_stream("PROCESS_VSZ")  # for processes of node controller
    cd.add_redis_stream("PROCESS_RSS")  # for processes of node controller
    cd.add_redis_stream("PROCESS_CPU")  # for processes of node controller

    cd.add_redis_stream("FREE_CPU",forward = True) # for entire processor
    cd.add_redis_stream("RAM",forward = True)
    cd.add_redis_stream("DISK_SPACE",forward = True) 
    cd.add_redis_stream("TEMPERATURE",forward = True)
    cd.add_redis_stream("CPU_CORE")
    cd.add_redis_stream("SWAP_SPACE")
    cd.add_redis_stream("IO_SPACE")
    cd.add_redis_stream("BLOCK_DEV")
    cd.add_redis_stream("CONTEXT_SWITCHES")
    cd.add_redis_stream("RUN_QUEUE")       
    cd.add_redis_stream("EDEV") 
    cd.close_package_contruction()
    
    
    cd.construct_package("DOCKER_CONTROL")

    cd.add_job_queue("DOCKER_COMMAND_QUEUE",1)
    cd.add_hash("DOCKER_DISPLAY_DICTIONARY")

    cd.close_package_contruction()


    cd.construct_package("DOCKER_MONITORING")
    cd.add_redis_stream("ERROR_STREAM")
    cd.add_hash("ERROR_HASH")
    cd.add_job_queue("WEB_COMMAND_QUEUE",1)
    cd.add_hash("WEB_DISPLAY_DICTIONARY")
    cd.add_hash("PROCESS_CONTROL")
    cd.close_package_contruction()
    
 
    
   
    bc.end_header_node("NODE_SYSTEM")
    
    
     
    bc.add_header_node("DOCKER_MONITOR",name,properties)
    cd.construct_package("DATA_STRUCTURES")
    cd.add_redis_stream("ERROR_STREAM",forward=True)
    cd.add_job_queue("WEB_COMMAND_QUEUE",1)
    cd.add_rpc_server("DOCKER_UPDATE_QUEUE",{"timeout":5,"queue":name+"_DOCKER_RPC_SERVER"})
    cd.add_hash("REBOOT_DATA")
    cd.add_hash("WEB_DISPLAY_DICTIONARY")
    cd.close_package_contruction()
    bc.end_header_node("DOCKER_MONITOR")
    #print("containers",containers)
    Construct_Services(bc,cd,services)
    Construct_Containers(bc,cd,containers)
    bc.end_header_node("PROCESSOR")    





if __name__ == "__main__" :

   file_handle = open("/data/redis_server.json",'r')
   data = file_handle.read()
   file_handle.close()
   redis_site = json.loads(data)


   bc = Build_Configuration(redis_site)
   cd = Construct_Data_Structures(redis_site["site"],bc)
   
   #
   #
   # Construct Systems
   #
   #
   bc.add_header_node( "SYSTEM","main_operations" )
   #
   #
   #  Construct Master Site
   #
   #
   
   bc.add_header_node( "SITE","CLOUD_SITE",  properties = {"address":"21005 Paseo Montana Murrieta, Ca 92562" } )

   construct_site_definitions(bc,cd,[])

   Cloud_Site_Definitons(bc,cd)


   

   
   construct_processor(name="block_chain_server",containers = ["monitor_redis","stream_events_to_log","stream_events_to_cloud","op_monitor"],services=["redis","ethereum_go","sqlite_server"])
   #
   
   construct_processor(name="gateway_server",containers = ["mqtt_interface"],services=["rpi_mosquitto"])
   #
   #
   #  Add other processes if desired
   #
   
   
   bc.end_header_node("SITE")                                                  

   bc.add_header_node( "SITE","LACIMA_SITE",  properties = {"address":"21005 Paseo Montana Murrieta, Ca 92562" } )

   lacima_services = [ "redis", "file_server" ]
   construct_site_definitions(bc,cd,services = lacima_services)
   LACIMA_Site_Definitons(bc,cd)


   
   containers = ["eto"   ]
   
   construct_processor(name="irrigation_controller",containers = containers,
                      services=[])
   #
   
   
   #
   #
   #  Add other processes if desired
   #
   
   
   bc.end_header_node("SITE")                  

   bc.end_header_node("SYSTEM")
   bc.check_namespace()
   bc.store_keys()
   #bc.extract_db()
   #bc.save_extraction("../code/system_data_files/extraction_file.pickle")
   #bc.delete_all()
   #bc.restore_extraction("extraction_file.pickle")
   #bc.delete_all()


 
