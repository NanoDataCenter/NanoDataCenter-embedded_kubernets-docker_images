# file build system
#
#  The purpose of this file is to load a system configuration
#  in the graphic data base
#

import json

import redis
from  build_configuration_py3 import Build_Configuration
from  construct_data_structures_py3 import Construct_Data_Structures
#from  graph_modules_py3.lacima.construct_applications_py3 import Construct_Lacima_Applications
#from  graph_modules_py3.lacima.construct_controller_py3 import Construct_Lacima_Controllers
#from  graph_modules_py3.lacima.construct_redis_monitor_py3 import Construct_Lacima_Redis_Monitoring
#from  graph_modules_py3.lacima.construct_mqtt_devices_py3  import  Construct_Lacima_MQTT_Devices
#from  graph_modules_py3.lacima.construct_plc_devices_py3   import  Construct_Lacima_PLC_Devices
#from  graph_modules_py3.lacima.construct_cloud_interface_py3 import Construct_Lacima_Cloud_Service
#from  graph_modules_py3.lacima.plc_measurements_py3 import Construct_Lacima_PLC_Measurements

from  graph_modules_py3.containers_py3.construct_container_py3 import Construct_Containers

def construct_processor(name,containers,services):
    properties = {}
    properties["containers"] = containers
    properties["services"] = services
    bc.add_header_node("PROCESSOR",name,properties= properties) 
    

    cd.construct_package("SYSTEM_MONITORING")
    cd.add_redis_stream("FREE_CPU",forward = True) # one month of data
    cd.add_redis_stream("RAM",forward = True)
    cd.add_redis_stream("DISK_SPACE",forward = True) # one month of data
    cd.add_redis_stream("TEMPERATURE",forward = True)
    cd.add_redis_stream("PROCESS_VSZ")
    cd.add_redis_stream("PROCESS_RSS")
    cd.add_redis_stream("PROCESS_CPU")   
    cd.add_redis_stream("CPU_CORE")
    cd.add_redis_stream("SWAP_SPACE")
    cd.add_redis_stream("IO_SPACE")
    cd.add_redis_stream("BLOCK_DEV")
    cd.add_redis_stream("CONTEXT_SWITCHES")
    cd.add_redis_stream("RUN_QUEUE")       
    cd.add_redis_stream("EDEV") 
     

    cd.close_package_contruction()
    properties = {}
    properties["command_list"] = [{"file":"pi_monitoring_py3.py","restart":True},{"file":"docker_monitoring_py3.py","restart":True}]
    bc.add_header_node("NODE_PROCESSES",name,properties=properties)
    cd.construct_package("DATA_STRUCTURES")
    cd.add_redis_stream("ERROR_STREAM",forward=True)
    cd.add_hash("ERROR_HASH")
    cd.add_job_queue("WEB_COMMAND_QUEUE",1)
    cd.add_hash("WEB_DISPLAY_DICTIONARY")
    cd.close_package_contruction()
    bc.end_header_node("NODE_PROCESSES")
    
    bc.add_header_node("DOCKER_MONITOR",name,properties)
    cd.construct_package("DATA_STRUCTURES")
    cd.add_redis_stream("ERROR_STREAM",forward=True)
    cd.add_hash("ERROR_HASH")
    cd.add_job_queue("WEB_COMMAND_QUEUE",1)
    cd.add_hash("WEB_DISPLAY_DICTIONARY")
    cd.close_package_contruction()
    bc.end_header_node("DOCKER_MONITOR")
    print("containers",containers)
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

   


   bc.add_header_node("CLOUD_SERVICE_QUEUE")
   cd.construct_package("CLOUD_SERVICE_QUEUE_DATA")
   cd.add_job_queue("CLOUD_JOB_SERVER",2048,forward=False)
   cd.add_hash("CLOUD_SUB_EVENTS")
   cd.close_package_contruction()
   
   bc.add_header_node("CLOUD_SERVICE_HOST_INTERFACE")
   bc.add_info_node( "HOST_INFORMATION","HOST_INFORMATION",properties={"host":"192.168.1.41" ,"port": 6379, "key_data_base": 6, "key":"_UPLOAD_QUEUE_" ,"depth":1024} )
   bc.end_header_node("CLOUD_SERVICE_HOST_INTERFACE")
   bc.end_header_node("CLOUD_SERVICE_QUEUE")
   
   
   bc.add_header_node("CLOUD_BLOCK_CHAIN_SERVER")
   cd.construct_package("CLOUD_BLOCK_CHAIN_SERVER")
   cd.add_rpc_server("BLOCK_CHAIN_RPC_SERVER",{"timeout":5,"queue":"BLOCK_CHAIN_RPC_SERVER"})
   cd.close_package_contruction()
   bc.end_header_node("CLOUD_BLOCK_CHAIN_SERVER")
   
   bc.add_header_node("SQL_SERVER")
   cd.construct_package("SQL_SERVER")
   cd.add_rpc_server("SQL_SERVER_RPC_SERVER",{"timeout":5,"queue":"SQL_RPC_SERVER"})
   cd.add_hash("SQL_DB_MAPPING")
   cd.close_package_contruction()
   bc.end_header_node("SQL_SERVER")
 
   bc.add_header_node("SYSTEM_MONITOR")
   cd.construct_package("SYSTEM_MONITOR")      
   #cd.add_managed_hash(self,name,fields,forward=False) perfored way to store field how to get field in system
   cd.add_hash("SYSTEM_STATUS")
   
   cd.add_hash("MONITORING_DATA")
   cd.add_redis_stream("SYSTEM_ALERTS")
   cd.add_redis_stream("SYSTEM_PUSHED_ALERTS")
   cd.close_package_contruction()
   bc.end_header_node("SYSTEM_MONITOR")

   bc.add_header_node("TICKET_CONTROL")
   bc.add_info_node( "DATA_BASE","TICKET_CONTROL", properties = {"db":"SYSTEM_CONTROL.db"} )
   bc.add_info_node("TABLE","TICKET_CONTROL",properties = {"name":"TICKET_CONTROL.db",
                    "fields":[ "id INTEGER AUTOINCREMENT","active Int","create_timestamp FLOAT","close_timestamp FLOAT","type Int","subtype Text","description TEXT","resolution TEXT"   ]} )
   bc.add_info_node("VALID_TYPES","TICKET_CONTROL",properties = {"types":["OTHERS","IRRIGATION_ISSUES","IRRIGATION_EQUIPMENT","TRIMMING"]})                   
   bc.end_header_node("TICKET_CONTROL")
    
   
   bc.add_info_node( "OP_MONITOR","OP_MONITOR", properties = {"OP_MONITOR_LIST":["CORE_OPS","MONITOR_REDIS","MONITOR_SQLITE","MONITOR_BLOCK_CHAIN"]} ) 
   
   construct_processor(name="block_chain_server",containers = ["monitor_redis","stream_events_to_log","stream_events_to_cloud","op_monitor"],services=["redis","ethereum_go","sqlite_server"])
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


 
