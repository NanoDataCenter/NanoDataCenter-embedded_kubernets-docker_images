package su
//import "fmt"

/*
 *  used to  
 *   redundant data structures were put in so that logging element doesnot have to dig into the stream
 *   This data structures could be used on arduino type processors
 */

func Construct_incident_logging(command_code ,description string){
    //fmt.Println("command_code",command_code)
    properties := make(map[string]interface{})
    properties["description"] = description
    Bc_Rec.Add_header_node("INCIDENT_LOG",command_code,properties)
	Cd_Rec.Construct_package("INCIDENT_LOG")
	Cd_Rec.Add_redis_stream("ERROR_LOG",1024)
    Cd_Rec.Add_single_element("TIME_STAMP")
    Cd_Rec.Add_single_element("STATUS")
    Cd_Rec.Add_single_element("CURRENT_STATE")
    Cd_Rec.Add_single_element("LAST_ERROR")
    Cd_Rec.Close_package_construction()
	Bc_Rec.End_header_node("INCIDENT_LOG",command_code)

}


/*
 *   
 * 
 * 
 * 
 * 
 */

func Construct_streaming_logs(stream_name , description string, keys []string ){

  properties := make(map[string]interface{})
  properties["keys"] = keys
  properties["descrption"] = description
  Bc_Rec.Add_header_node("STREAMING_LOG",stream_name,properties)
  Cd_Rec.Construct_package("STREAMING_LOG")
  for _,key := range keys{
       Cd_Rec.Add_redis_stream(key,1024)  // stream
       Cd_Rec.Add_hash(key+":ANALYSIS")  // used for aggregates of logs
  }
  Cd_Rec.Close_package_construction()
  Bc_Rec.End_header_node("STREAMING_LOG",stream_name)
}

/*
 * This data structure allows rpc servers to be scanned by diagnostic programs
 * 
 */
func  Construct_RPC_Server( command_code, description string,depth,timeout int64, properties map[string]interface{} ){
    
    
    properties["description"] = description
    Bc_Rec.Add_header_node("RPC_SERVER",command_code,properties)
    Cd_Rec.Construct_package("RPC_SERVER")
    Cd_Rec.Add_rpc_server("RPC_SERVER",depth,timeout)
    Cd_Rec.Add_hash("RPC_STATISTICS")  // used for aggregates of logs
    Cd_Rec.Close_package_construction()
    Bc_Rec.End_header_node("RPC_SERVER",command_code)    
}






func Construct_watchdog_logging(command_code , description string, max_time_interval int){
 properties := make(map[string]interface{})
 properties["description"] = description
 properties["max_time_interval"] = max_time_interval
  Bc_Rec.Add_header_node("WATCH_DOG",command_code,properties)
 
  Cd_Rec.Construct_package("WATCH_DOG")
  Cd_Rec.Add_single_element("WATCH_DOG_TS")   // used to stored timestamp
  Cd_Rec.Add_single_element("WATCH_DOG_STATE")   // used to state
  Cd_Rec.Close_package_construction()
  Bc_Rec.End_header_node("WATCH_DOG",command_code)

}
