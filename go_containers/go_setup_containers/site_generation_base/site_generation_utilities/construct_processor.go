package su


func Construct_processor(name string, containers []string){

      properties := make(map[string]interface{})
      properties["containers"] = containers
      Bc_Rec.Add_header_node("PROCESSOR",name, properties )
   
      var description string
      
      description = name + "node reboot"
      Construct_incident_logging("NODE_REBOOT",description)
      
      description = name + "processor rpc ping status"
      Construct_incident_logging("NODE_RPC_PING",description)
      
      keys := []string{"FREE_CPU","RAM","TEMPERATURE","DISK_SPACE","SWAP_SPACE","CONTEXT_SWITCHES","BLOCK_DEV","IO_SPACE","RUN_QUEUE","EDEV"}
      Bc_Rec.Add_header_node("PROCESSOR_MONITORING","PROCESSOR_MONITORING", make(map[string]interface{}))
	  description = name+" processor_monitor"
	  Construct_streaming_logs("processor_monitor",description,keys) //wait until flush out
	  Bc_Rec.End_header_node("PROCESSOR_MONITORING","PROCESSOR_MONITORING")

      
      
 

      
      
      Construct_RPC_Server("NODE_CONTROL","rpc for controlling node",10,15,  make(map[string]interface{}) )
      
      Construct_RPC_Server( "NODE_CONTAINER_CONTROL","NODE CONTAINER_CONTROL",5,1, make(map[string]interface{}) )
      Construct_incident_logging("CONTAINER_ERROR_STREAM" ,"container error stream")
      
      Cd_Rec.Construct_package("DOCKER_CONTROL")
      Cd_Rec.Add_hash("DOCKER_DISPLAY_DICTIONARY")
      Cd_Rec.Close_package_contruction()
      
      register_containers(containers)
      Bc_Rec.End_header_node("PROCESSOR",name)

 
}    
