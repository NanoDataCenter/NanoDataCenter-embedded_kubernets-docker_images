package su


func Construct_processor(name string, containers []string){
    properties := make(map[string]interface{})
    properties["containers"] = containers
    Bc_Rec.Add_header_node("PROCESSOR",name,properties) 
    register_containers(containers)


    properties = make(map[string]interface{})
    Bc_Rec.Add_header_node("PROCESSOR_MONITORING","PROCESSOR_MONITORING",properties)
	stream_list := []string{"FREE_CPU","RAM","DISK_SPACE","TEMPERATURE","SWAP_SPACE","IO_SPACE","BLOCK_DEV","CONTEXT_SWITCHES","RUN_QUEUE","EDEV"}
	Construct_streaming_logs("processor_monitor",stream_list)
	Bc_Rec.End_header_node("PROCESSOR_MONITORING","PROCESSOR_MONITORING")
    

    
	/*
	
	
	cd.construct_package("SITE_NODE_CONTROL_LOG")
    cd.add_redis_stream("ERROR_STREAM") # for entire processor
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
  
 

    '
    cd.construct_package("DOCKER_MONITORING")
    cd.add_redis_stream("ERROR_STREAM")
    cd.add_hash("ERROR_HASH")
    cd.add_job_queue("WEB_COMMAND_QUEUE",1)
    cd.add_hash("WEB_DISPLAY_DICTIONARY")
    cd.add_hash("PROCESS_CONTROL")
    cd.close_package_contruction()
    
 
    
   

    
    
     
    bc.add_header_node("DOCKER_MONITOR",name,properties)
    cd.construct_package("DATA_STRUCTURES")
    cd.add_redis_stream("ERROR_STREAM",forward=True)
    cd.add_job_queue("WEB_COMMAND_QUEUE",1)
    cd.add_rpc_server("DOCKER_UPDATE_QUEUE",{"timeout":5,"queue":name+"_DOCKER_RPC_SERVER"})
    cd.add_hash("REBOOT_DATA")
    cd.add_hash("WEB_DISPLAY_DICTIONARY")
    cd.close_package_contruction()
    bc.end_header_node("DOCKER_MONITOR")
   */
    
    
    Bc_Rec.End_header_node("PROCESSOR",name)
}    
