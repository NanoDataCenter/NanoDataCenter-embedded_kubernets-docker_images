package su
//import "fmt"

/*
 *  used to  
 *    Mark error value
 *    Store in log
 *    Mark time of error
 *    Mark if error is still present
 *    
 */

func Construct_incident_logging(command_code string){
    //fmt.Println("command_code",command_code)
    properties := make(map[string]interface{})
    Bc_Rec.Add_header_node("INCIDENT_LOG",command_code,properties)
	Cd_Rec.Construct_package("INCIDENT_LOG")
	Cd_Rec.Add_Influxdb_Shared_Stream("INCIDENT_LOG",[]string{"command_code"},[]string{"status","msgpack_data"})
    Cd_Rec.Close_package_contruction()
	Bc_Rec.End_header_node("INCIDENT_LOG",command_code)

}


func Construct_streaming_logs(stream_name string, tags, keys []string){

  properties := make(map[string]interface{})
  Bc_Rec.Add_header_node("STREAMING_LOG",command_code,properties)
  Cd_Rec.Construct_package("STREAMING_LOG")
  Cd_Rec.Add_Influxdb_Monitored_Stream("STREAMING_LOG",tags,keys)  // stream
  Cd_Rec.Add_hash("STREAMING_LOG")  // used for aggregates of logs
  Cd_Rec.Close_package_contruction()
  Bc_Rec.End_header_node("STREAMING_LOG",command_code)
}






func Construct_watchdog_logging(command_code string){
 properties := make(map[string]interface{})
 
  Bc_Rec.Add_header_node("WATCH_DOG",command_code,properties)
 
  Cd_Rec.Construct_package("NODE_WATCH_DOG")
  Cd_Rec.Add_single_element("NODE_WATCH_DOG")   // used to stored timestamp
  Cd_Rec.Close_package_contruction()
  Bc_Rec.End_header_node("WATCH_DOG",command_code)

}
