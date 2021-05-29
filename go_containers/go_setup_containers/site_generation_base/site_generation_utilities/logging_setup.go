package su
//import "fmt"

func Construct_incident_logging(command_code string){
    //fmt.Println("command_code",command_code)
    properties := make(map[string]interface{})
    Bc_Rec.Add_header_node("INCIDENT_LOG",command_code,properties)
	
    Cd_Rec.Construct_package("INCIDENT_LOG")
	Cd_Rec.Add_single_element("TIME_STAMP") // unix time stamp
    Cd_Rec.Add_single_element("STATUS") // true or false
    Cd_Rec.Add_single_element("CURRENT_STATE") // current value of total data
    Cd_Rec.Add_single_element("LAST_ERROR") // current value of error data
    Cd_Rec.Add_redis_stream("ERROR_LOG",1024)   // log of error data
    Cd_Rec.Close_package_contruction()
	
	Bc_Rec.End_header_node("INCIDENT_LOG",command_code)

}


func Construct_streaming_logs(command_code string,keys []string){

  properties := make(map[string]interface{})
  Bc_Rec.Add_header_node("STREAMING_LOG",command_code,properties)
 
  Cd_Rec.Construct_package("STREAMING_LOG")
  for _,i:= range keys {  
    Cd_Rec.Add_redis_stream(i,1024)
  }
  Cd_Rec.Close_package_contruction()
  Cd_Rec.Construct_package("STREAMING_VALUES")
   for _,i:= range keys {  
    Cd_Rec.Add_single_element(i)  // stores the running monitor values
  } 
  Cd_Rec.Close_package_contruction()
  
  Bc_Rec.End_header_node("STREAMING_LOG",command_code)
}






func Construct_watchdog_logging(command_code string){
 properties := make(map[string]interface{})
 
  Bc_Rec.Add_header_node("WATCH_DOG",command_code,properties)
 
  Cd_Rec.Construct_package("NODE_WATCH_DOG")
  Cd_Rec.Add_single_element("NODE_WATCH_DOG") 
  Cd_Rec.Close_package_contruction()
  Bc_Rec.End_header_node("WATCH_DOG",command_code)

}
