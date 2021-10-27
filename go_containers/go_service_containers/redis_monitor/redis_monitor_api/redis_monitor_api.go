package redis_monitor_api


import "fmt"
import "time"
import "strings"
import "context"
import "strconv"
import "encoding/json"
//import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"

import "lacima.com/Patterns/logging_support"
import "lacima.com/Patterns/msgpack_2"
import "lacima.com/server_libraries/postgres"
import "github.com/go-redis/redis/v8"




type Redis_Monitor_Type struct {
   
   ctx                context.Context;
   client             *redis.Client;
   performance_log    pg_drv.Postgres_Stream_Driver
   incident_log       *logging_support.Incident_Log_Type
}

var monitor_structure Redis_Monitor_Type






func Init(){


    monitor_structure.incident_log    = logging_support.Construct_incident_log([]string{"REDIS_MONITORING:REDIS_MONITORING" ,"INCIDENT_LOG"} )    
  	data_search_list := []string{ "REDIS_MONITORING","REDIS_MONITORING","REDIS_MONITORING"}
	data_element := data_handler.Construct_Data_Structures(&data_search_list)
	
	single_element   := (*data_element)["REDIS_MONITORING"].(redis_handlers.Redis_Single_Structure)
	
    monitor_structure.client          = single_element.Get_client()
	monitor_structure.ctx             = single_element.Get_context()
    monitor_structure.performance_log  = logging_support.Find_stream_logging_driver()
}
  

func Exec(){

    for true {
	   log_data()
	   time.Sleep(time.Minute*15)
	 }
}

func log_data(){
     
      log_keyspace()
	  log_clients()
	    
      log_memory()
	 
      commandstats()
	  fmt.Println("log data ")   
}








func log_keyspace(){
   value,err := monitor_structure.client.Info(monitor_structure.ctx,"Keyspace").Result()
   if err == nil {
    
       line_data :=  generate_line_map(value)
       
	   key_data := process_key_data(line_data)
       
	   stream_data := generate_msgpack_floats(key_data)
       post_stream_data("Keyspace",stream_data)
	    
	}else{
	   ;
	}
    
}



func log_clients(){

   value,err := monitor_structure.client.Info(monitor_structure.ctx,"Clients").Result()
   if err == nil {
       line_data := generate_line_map(value)
       
       all_data := generate_msgpack_floats(line_data)
       
       
	   log_data := make(map[string]string)
       log_data["connected_clients"] =all_data["connected_clients"]
       log_data["blocked_clients"] = all_data["blocked_clients"]	  
       log_data["maxclients"]     = all_data["maxclients"]
       max_clients,_ := msg_pack_utils.Unpack_float64(log_data["maxclients"])
       connected_clients,_ := msg_pack_utils.Unpack_float64(log_data["connected_clients"])
       ratio := connected_clients/max_clients
       //fmt.Println("ratio",ratio)
       if ratio > .75 {
           var state map[string]interface{}
           state["system"]              = "redis_monitor"
           state["subsystem"]           = "clients"
           state["connected_clients"]   = connected_clients
           state["max_clients"]         = max_clients
           post_incident_report(state)
       }    
       
       post_stream_data("Keyspace",log_data)
	  
	}else{
	   ;
	}
  
}

func log_memory(){

   value,err := monitor_structure.client.Info(monitor_structure.ctx,"Memory").Result()
   if err == nil {
       line_data := generate_line_map(value)
       
	   all_data := generate_msgpack_floats(line_data)
	   log_data := make(map[string]string)
       log_data["maxmemory"] =all_data["maxmemory"]
       log_data["used_memory"] = all_data["used_memory"]
       max_memory,_ := msg_pack_utils.Unpack_float64(log_data["maxmemory"])
       used_memory,_ := msg_pack_utils.Unpack_float64(log_data["used_memory"])
       ratio := used_memory/max_memory
       //fmt.Println("ratio",ratio)
       if ratio > .75 {
           var state map[string]interface{}
           state["system"]              = "redis_monitor"
           state["subsystem"]           = "memory"
           state["maxmemory"]           = max_memory
           state["used_memory"]         = used_memory
           post_incident_report(state)
       }          
       post_stream_data("Memory",log_data)  
	}else{
	   ;
	}
   
}

 


func commandstats(){


      value,err := monitor_structure.client.Info(monitor_structure.ctx,"commandstats").Result()
	  
      if err == nil {
	   log_data := make(map[string]string)
       line_data := generate_line_map(value)
    
       for key, value := range line_data {
	      log_data[key] = extract_commandstats_data(value)
	     
	   }
	   post_stream_data("REDIS_MONITOR_CMD_TIME_STREAM",log_data)  
      	    
	}else{
	   ;
	}
    
}





func extract_commandstats_data( input string) string {

  var return_value string
  return_value = msg_pack_utils.Pack_float64(0)
  temp1 := strings.Split(input,",")
  if len(temp1) >= 3 {
     temp2 := strings.Split(temp1[2],"=")
	 if len(temp2) == 2 {
	    value  ,err := strconv.ParseFloat(temp2[1], 64)
	    if err == nil{
		   return_value = msg_pack_utils.Pack_float64(value)
		}
   	}
  }
  return return_value


}



func generate_line_map (input string )map[string]string{
   return_value := make(map[string]string)
   lines := strings.Split(input,"\r\n")
   for _,i := range lines{
     temp := strings.Split(i,":")
     if len(temp) == 2{
	   return_value[temp[0]] = temp[1]
	 }
  }
  return return_value


}


func process_key_data( line_map map[string]string ) map[string]string{

    return_value := make(map[string]string)
    for key,value := range line_map {
	  temp := strings.Split(value,"keys=")
	  fields := strings.Split(temp[1],",")
	  return_value[key] = fields[0]
		
	}
    return return_value

}


func generate_msgpack_floats( input map[string]string ) map[string]string {

   return_value := make(map[string]string)
   for key ,value := range input {
   
      temp  ,err := strconv.ParseFloat(value, 64)
	  
	  if err == nil {
	      return_value[key] = msg_pack_utils.Pack_float64(temp)
	  }
	  
	 
   
   }
   return return_value

}

func post_stream_data(subsystem string,log_data map[string]string){
   
    
    for key,value := range log_data{
    
       monitor_structure.performance_log.Insert( "REDIS_ANALYSIS",subsystem,key,"","",value )
    }
    
    
}


func post_incident_report(incident_data map[string]interface{}){
    
    request_json,err := json.Marshal(&incident_data)
    if err != nil{
          panic("json marshall error")
    }  
    fmt.Println("request_json",string(request_json))
    monitor_structure.incident_log.Log_data(string(request_json))
}
