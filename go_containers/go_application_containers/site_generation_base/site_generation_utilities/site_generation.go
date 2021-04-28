package su

var working_site string


func End_site_definitions(){

   Bc_Rec.End_header_node("SITE",working_site)

}




func Start_site_definitions(site_name string, system_containers, startup_containers   []string){
    
    working_site = site_name
    properties := make(map[string]interface{})
	properties["startup_containers"] = startup_containers
	properties["containers"] = system_containers
    Bc_Rec.Add_header_node( "SITE",site_name,  properties  )
    
	register_containers(system_containers)
	
	// build site containers
	
    	
    /*
       ---- figure this out later	
    
	   these control structures allow external
	   programs such as web browsers to control 
	   the site
	
    cd.Construct_package("SITE_CONTROL")

    cd.add_job_queue("WEB_COMMAND_QUEUE",10) // commands such as reboot pull container
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
	
	
	    bc.add_header_node("SYSTEM_MONITOR")
    cd.construct_package("SYSTEM_MONITOR")      
    cd.add_hash("SYSTEM_VERBS")
    cd.add_redis_stream("SYSTEM_ALERTS")
    cd.close_package_contruction()
    bc.end_header_node("SYSTEM_MONITOR")
  
    */
   
    /*
       these are application related
	   data structures for site monitoring
	*/
   
    
     properties = make(map[string]interface{})
    Bc_Rec.Add_header_node("REDIS_MONITORING","REDIS_MONITORING",properties)
	Construct_streaming_logs("redis_monitor",[]string{"STREAMING_LOG","KEYS","CLIENTS","MEMORY","REDIS_MONITOR_CMD_TIME_STREAM"})

    Bc_Rec.End_header_node("REDIS_MONITORING","REDIS_MONITORING")
   
   
    Bc_Rec.Add_header_node("FILE_SERVER","FILE_SERVER",properties)
    Cd_Rec.Construct_package("FILE_SERVER")
    Cd_Rec.Add_rpc_server("FILE_SERVER_RPC_SERVER",30,10)
    Cd_Rec.Close_package_contruction()
    Bc_Rec.End_header_node("FILE_SERVER","FILE_SERVER")
    

    
}

