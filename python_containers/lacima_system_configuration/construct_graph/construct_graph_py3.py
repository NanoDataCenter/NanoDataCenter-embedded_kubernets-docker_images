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



 
   


def construct_site_definitions(bc,cd,graph_container_image,graph_container_script,containers ):
    
    Construct_Containers(bc,cd,containers)
    properties = {}
   
    properties["containers"] = containers
    properties["graph_container_image"] = graph_container_image
    properties["graph_container_script"] = graph_container_script
    
    properties["command_list"] = [{"file":"docker_control_py3.py","restart":True},{"file":"redis_monitoring_py3.py","restart":True},{"file":"node_monitoring_py3.py","restart":True}]
    bc.add_header_node("SITE_CONTROL","SITE_CONTROL",properties= properties) 
   
    cd.construct_package("SITE_CONTROL")

    cd.add_job_queue("WEB_COMMAND_QUEUE",10)
    cd.add_redis_stream("ERROR_STREAM")
    cd.add_hash("ERROR_HASH")
    cd.add_hash("WEB_DISPLAY_DICTIONARY")
    cd.close_package_contruction()
    
    cd.construct_package("DOCKER_CONTROL")
    cd.add_job_queue("DOCKER_COMMAND_QUEUE",10)
    cd.add_hash("DOCKER_DISPLAY_DICTIONARY")
    cd.close_package_contruction()

    cd.construct_package("NODE_MONITORING")
    cd.add_job_queue("WEB_COMMAND_QUEUE",1)
    cd.add_hash("WEB_DISPLAY_DICTIONARY")
    cd.add_hash("NODE_STATUS")
    cd.close_package_contruction()
   
    '''
    cd.construct_package("DOCKER_MONITORING")
    cd.add_redis_stream("ERROR_STREAM")
    cd.add_hash("ERROR_HASH")
    cd.add_job_queue("WEB_COMMAND_QUEUE",1)
    cd.add_hash("WEB_DISPLAY_DICTIONARY")
    cd.add_hash("PROCESS_CONTROL")
    cd.close_package_contruction()
    '''
   
    cd.construct_package("REDIS_MONITORING")  # redis monitoring
    cd.add_redis_stream("KEYS")
    cd.add_redis_stream("CLIENTS")
    cd.add_redis_stream("MEMORY")
    cd.add_redis_stream("REDIS_MONITOR_CALL_STREAM")
    cd.add_redis_stream("REDIS_MONITOR_CMD_TIME_STREAM")  
    cd.add_redis_stream("REDIS_MONITOR_SERVER_TIME")  
    cd.close_package_contruction()
    bc.end_header_node("SITE_CONTROL")
   
    
    bc.add_header_node("FILE_SERVER")
    cd.construct_package("FILE_SERVER")
    cd.add_rpc_server("FILE_SERVER_RPC_SERVER",{"timeout":30,"queue":"FILE_RPC_SERVER"})
    cd.close_package_contruction()
    bc.end_header_node("FILE_SERVER")

    bc.add_header_node("FILE_SERVER_CLIENT")
    cd.construct_package("FILE_SERVER_CLIENT")
    cd.add_rpc_client("FILE_SERVER_RPC_SERVER",{"timeout":30,"queue":"FILE_RPC_SERVER"})
    cd.close_package_contruction()
    bc.end_header_node("FILE_SERVER_CLIENT")
    
    
    
    bc.add_header_node("SYSTEM_MONITOR")
    cd.construct_package("SYSTEM_MONITOR")      
    cd.add_hash("SYSTEM_VERBS")
    cd.add_redis_stream("SYSTEM_ALERTS")
    cd.close_package_contruction()
    bc.end_header_node("SYSTEM_MONITOR")

 
 

def construct_processor(name,containers):
    properties = {}
    properties["containers"] = containers

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
   '''
   bc.add_header_node( "SITE","CLOUD_SITE",  properties = {"address":"21005 Paseo Montana Murrieta, Ca 92562" } )

   graph_container_image = "nanodatacenter/lacima_system_configuration"
   graph_container_script ="docker run   -it --network host --rm  --name lacima_system_configuration  --mount type=bind,source=/mnt/ssd/site_config,target=/data/ nanodatacenter/lacima_system_configuration /bin/bash construct_graph.bsh"
   construct_site_definitions(bc,cd,graph_container_image,graph_container_script,services=[],containers = [])

   Cloud_Site_Definitons(bc,cd)


   

   
   construct_processor(name="block_chain_server",containers = ["monitor_redis","stream_events_to_log","stream_events_to_cloud","op_monitor"],services=["redis","ethereum_go","sqlite_server","file_server"])
   #
   
   construct_processor(name="gateway_server",containers = ["mqtt_interface"],services=["rpi_mosquitto"])
   #
   #
   #  Add other processes if desired
   #
   
   
   bc.end_header_node("SITE")                                                  

   '''
   properties = {}
   properties["services"] = [ "redis", "file_server" ]
   properties["address"] = "21005 Paseo Montana Murrieta, Ca 92562" 
  
   bc.add_header_node( "SITE","LACIMA_SITE",  properties = properties )

   lacima_containers = [ "redis", "file_server"]
   graph_container_image = "nanodatacenter/lacima_system_configuration"
   graph_container_script ="docker run   -it --network host --rm  --name lacima_system_configuration  --mount type=bind,source=/mnt/ssd/site_config,target=/data/ nanodatacenter/lacima_system_configuration /bin/bash construct_graph.bsh"
   construct_site_definitions(bc,cd,graph_container_image,graph_container_script,containers=lacima_containers)

   
  
   LACIMA_Site_Definitons(bc,cd)


   
   containers = ["eto","irrigation_scheduling"  ]
   
   construct_processor(name="irrigation_controller",containers = containers)
                     
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


 
