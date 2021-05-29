package su


func Construct_processor(name string, containers []string){

    properties := make(map[string]interface{})
    properties["containers"] = containers
    Bc_Rec.Add_header_node("PROCESSOR",name, properties )
    register_containers(containers)



    Bc_Rec.Add_header_node("PROCESSOR_MONITORING","PROCESSOR_MONITORING", make(map[string]interface{}))
	stream_list := []string{"FREE_CPU","RAM","DISK_SPACE","TEMPERATURE","SWAP_SPACE","IO_SPACE","BLOCK_DEV","CONTEXT_SWITCHES","RUN_QUEUE","EDEV"}
	Construct_streaming_logs("processor_monitor",stream_list)
	Bc_Rec.End_header_node("PROCESSOR_MONITORING","PROCESSOR_MONITORING")

   Bc_Rec.Add_header_node("PROCESSOR_WATCHDOG","PROCESSOR_WATCHDOG", make(map[string]interface{}))
    Construct_watchdog_logging("PROCESSOR_WATCHDOG")   
	Bc_Rec.End_header_node("PROCESSOR_WATCHDOG","PROCESSOR_WATCHDOG")

      Bc_Rec.End_header_node("PROCESSOR",name)

 
}    
