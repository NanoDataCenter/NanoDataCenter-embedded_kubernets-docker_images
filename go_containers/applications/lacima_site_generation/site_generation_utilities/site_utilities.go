package site_processor




func Construct_site_definitions( bc  , cd , start_up_containers,  container ){

   container.Construct_Containers(bc,cd,containers )
   properties := make(map[string]
   








}
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
    cd.add_hash("WEB_DISPLAY_DICTIONARY")  # for displaying node status
    cd.close_package_contruction()
    
    cd.construct_package("DOCKER_CONTROL")
    cd.add_job_queue("DOCKER_COMMAND_QUEUE",10) #temp disable turning of containers
    cd.add_hash("DOCKER_DISPLAY_DICTIONARY")
    cd.add_redis_stream("ERROR_STREAM")
    cd.close_package_contruction()

    cd.construct_package("NODE_MONITORING")
    cd.add_job_queue("WEB_COMMAND_QUEUE",1)
    cd.add_hash("NODE_STATUS")
    cd.add_redis_stream("ERROR_STREAM")
    cd.add_hash("SYSTEM_CONTAINER_IMAGES") # value list of nodes container is in
    cd.close_package_contruction()
   
    cd.construct_package("SITE_REBOOT_LOG")
    cd.add_redis_stream("SITE_REBOOT_LOG")
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
    
    cd.add_redis_stream("REDIS_MONITOR_CMD_TIME_STREAM")  
    #cd.add_redis_stream("REDIS_MONITOR_SERVER_TIME") 
    #cd.add_redis_stream("REDIS_MONITOR_CALL_STREAM")    
    cd.close_package_contruction()
    bc.end_header_node("SITE_CONTROL")
   
    
    bc.add_header_node("FILE_SERVER")
    cd.construct_package("FILE_SERVER")
    cd.add_rpc_server("FILE_SERVER_RPC_SERVER",{"depth":10,"timeout":30})
    cd.close_package_contruction()
    bc.end_header_node("FILE_SERVER")

    bc.add_header_node("TP_MONITOR_SWITCHES")
    
    properties = {}
    properties["ip"] = "192.168.1.45"
    properties["id"] = "Gr1234gfd"
    bc.add_header_node("TP_SWITCH","switch_office",properties)
    cd.construct_package("LOG_DATA")
    cd.add_single_element("STATUS")
    cd.add_single_element("CURRENT_STATE")
    cd.add_single_element("LAST_ERROR")
    cd.add_redis_stream("ERROR_LOG")
    cd.close_package_contruction()
    bc.end_header_node("TP_SWITCH")

    properties = {}
    properties["ip"] = "192.168.1.56"
    properties["id"] = "Gr1234gfd"
    bc.add_header_node("TP_SWITCH","switch_garage",properties)
    cd.construct_package("LOG_DATA")
    cd.add_single_element("STATUS")
    cd.add_single_element("CURRENT_STATE")
    cd.add_single_element("LAST_ERROR")
    cd.add_redis_stream("ERROR_LOG")
    cd.close_package_contruction()
    bc.end_header_node("TP_SWITCH")
    
    bc.end_header_node("TP_MONITOR_SWITCHES")
    
    
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


    cd.construct_package("SITE_NODE_CONTROL_LOG")
    cd.add_redis_stream("ERROR_STREAM") # for entire processor
    cd.close_package_contruction()
    
    cd.construct_package("PROCESSOR_MONITORING")

    cd.add_redis_stream("FREE_CPU",forward = True) # for entire processor
    cd.add_redis_stream("RAM",forward = True)
    cd.add_redis_stream("DISK_SPACE",forward = True) 
    cd.add_redis_stream("TEMPERATURE",forward = True)
    
    cd.add_redis_stream("SWAP_SPACE")
    cd.add_redis_stream("IO_SPACE")
    cd.add_redis_stream("BLOCK_DEV")
    cd.add_redis_stream("CONTEXT_SWITCHES")
    cd.add_redis_stream("RUN_QUEUE")       
    cd.add_redis_stream("EDEV") 
    cd.close_package_contruction()
    
    cd.construct_package("DOCKER_CONTROL")
    cd.add_job_queue("DOCKER_COMMAND_QUEUE",10) #temp disable turning of containers
    cd.add_hash("DOCKER_DISPLAY_DICTIONARY")
    cd.add_redis_stream("ERROR_STREAM")
    cd.close_package_contruction()
    
    cd.construct_package("NODE_CONTROL")
    cd.add_job_queue("NODE_COMMAND_QUEUE",10) #temp disable turning of containers
    cd.add_single_element("NODE_WATCH_DOG")
    cd.add_job_queue("NODE_UPGRADE_QUEUE",500)
    cd.close_package_contruction()
    
    
    cd.construct_package("NODE_REBOOT_LOG")
    cd.add_redis_stream("NODE_REBOOT_LOG")
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
 
    
   
    bc.end_header_node("NODE_SYSTEM")
    
    '''
     
    bc.add_header_node("DOCKER_MONITOR",name,properties)
    cd.construct_package("DATA_STRUCTURES")
    cd.add_redis_stream("ERROR_STREAM",forward=True)
    cd.add_job_queue("WEB_COMMAND_QUEUE",1)
    cd.add_rpc_server("DOCKER_UPDATE_QUEUE",{"timeout":5,"queue":name+"_DOCKER_RPC_SERVER"})
    cd.add_hash("REBOOT_DATA")
    cd.add_hash("WEB_DISPLAY_DICTIONARY")
    cd.close_package_contruction()
    bc.end_header_node("DOCKER_MONITOR")
    '''
    
    Construct_Containers(bc,cd,containers)
    bc.end_header_node("PROCESSOR")    






