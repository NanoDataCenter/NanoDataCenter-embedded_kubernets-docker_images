package su


func Construct_processor(name string, containers []string){

      properties := make(map[string]interface{})
      properties["containers"] = containers
      Bc_Rec.Add_header_node("PROCESSOR",name, properties )
   

      Construct_incident_logging("NODE_REBOOT")

      Bc_Rec.Add_header_node("PROCESSOR_MONITORING","PROCESSOR_MONITORING", make(map[string]interface{}))
	  
	  Construct_streaming_logs("processor_monitor")
	  Bc_Rec.End_header_node("PROCESSOR_MONITORING","PROCESSOR_MONITORING")

      Bc_Rec.Add_header_node("PROCESSOR_WATCHDOG","PROCESSOR_WATCHDOG", make(map[string]interface{}))
      Construct_watchdog_logging("PROCESSOR_WATCHDOG",60)   
	  Bc_Rec.End_header_node("PROCESSOR_WATCHDOG","PROCESSOR_WATCHDOG")

      register_containers(containers)
      Bc_Rec.End_header_node("PROCESSOR",name)

 
}    
