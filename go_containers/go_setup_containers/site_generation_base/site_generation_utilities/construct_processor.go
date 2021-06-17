package su


func Construct_processor(name string, containers []string){

      properties := make(map[string]interface{})
      properties["containers"] = containers
      Bc_Rec.Add_header_node("PROCESSOR",name, properties )
   
      var description string
      
      description = name + "node reboot"
      Construct_incident_logging("NODE_REBOOT",description)
      keys := []string{"FREE_CPU","RAM","TEMPERATURE","DISK_SPACE","SWAP_SPACE","CONTEXT_SWITCHES","BLOCK_DEV","IO_SPACE","RUN_QUEUE","EDEV"}
      Bc_Rec.Add_header_node("PROCESSOR_MONITORING","PROCESSOR_MONITORING", make(map[string]interface{}))
	  description = name+" processor_monitor"
	  Construct_streaming_logs("processor_monitor",description,keys) //wait until flush out
	  Bc_Rec.End_header_node("PROCESSOR_MONITORING","PROCESSOR_MONITORING")

      description = name + " processor watchdog"
      Bc_Rec.Add_header_node("PROCESSOR_WATCHDOG","PROCESSOR_WATCHDOG", make(map[string]interface{}))
      Construct_watchdog_logging("PROCESSOR_WATCHDOG",description,60)   
	  Bc_Rec.End_header_node("PROCESSOR_WATCHDOG","PROCESSOR_WATCHDOG")

      
      
      Construct_RPC_Server("NODE_CONTROL","rpc for controlling node",10,15,  make(map[string]interface{}) )
      
      Cd_Rec.Construct_package("DOCKER_CONTROL")
      Cd_Rec.Add_job_queue("DOCKER_COMMAND_QUEUE",10) //temp disable turning of containers
      Cd_Rec.Add_hash("DOCKER_DISPLAY_DICTIONARY")
      Cd_Rec.Add_redis_stream("ERROR_STREAM",1024)
      Cd_Rec.Close_package_contruction()
      register_containers(containers)
      Bc_Rec.End_header_node("PROCESSOR",name)

 
}    
