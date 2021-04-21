package main

import "time"
import "context"
import "strings"
import "strconv"
import "bytes"
import "fmt"
import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"
import "github.com/go-redis/redis/v8"

import "github.com/msgpack/msgpack-go"


var site_data_store map[string]interface{}
const config_file = "/data/redis_server.json"


type Redis_Monitor_Type struct {
   
   ctx context.Context;
   client *redis.Client;
   streams map[string]redis_handlers.Redis_Stream_Struct
}

var redis_monitor_structure Redis_Monitor_Type

func main(){

   
 
    site_data_store = get_site_data.Get_site_data(config_file)
    graph_query.Graph_support_init(&site_data_store)
	redis_handlers.Init_Redis_Mutex()
	data_handler.Data_handler_init(&site_data_store)

	(&redis_monitor_structure).Init()
	(&redis_monitor_structure).Exec()
	
   

}





func ( v *Redis_Monitor_Type)Init(){


	v.streams = make(map[string]redis_handlers.Redis_Stream_Struct)
  	data_search_list := []string{ "REDIS_MONITORING"}
	data_element := data_handler.Construct_Data_Structures(&data_search_list)
	
	v.streams["KEYS"]                           = (*data_element)["KEYS"].(redis_handlers.Redis_Stream_Struct)
	v.streams["CLIENTS"]                        = (*data_element)["CLIENTS"].(redis_handlers.Redis_Stream_Struct)
	v.streams["MEMORY"]                         = (*data_element)["MEMORY"].(redis_handlers.Redis_Stream_Struct)
	v.streams["REDIS_MONITOR_CMD_TIME_STREAM"]  = (*data_element)["REDIS_MONITOR_CMD_TIME_STREAM"].(redis_handlers.Redis_Stream_Struct)
    v.client = v.streams["KEYS"].Get_client()
	v.ctx = v.streams["KEYS"].Get_context()

}
  

func (v *Redis_Monitor_Type)Exec(){

    for true {
	   v.log_data()
	   time.Sleep(time.Minute)
	 }
}

func ( v *Redis_Monitor_Type)log_data(){
     
      v.log_keyspace()
	  v.log_clients()
	    
      v.log_memory()
	 
      v.commandstats()
	  fmt.Println("log data ")   
}








func ( v *Redis_Monitor_Type)log_keyspace(){
   value,err := v.client.Info(v.ctx,"Keyspace").Result()
   if err == nil {
       log_value_data := generate_map(value)
	   
      
	   process_key_data(&log_value_data)
	   log_data := generate_floats(&log_value_data)
	   v.push_data( "KEYS" , &log_data)	  
	}else{
	   ;
	}
    
}



func ( v *Redis_Monitor_Type)log_clients(){
   value,err := v.client.Info(v.ctx,"Clients").Result()
   if err == nil {
       log_value_data := generate_map(value)
	   all_data := generate_floats(&log_value_data)
	   log_data := make(map[string]float64)
       log_data["connected_clients"] =all_data["connected_clients"]
       log_data["blocked_clients"] = all_data["blocked_clients"]	   
	   v.push_data( "CLIENTS" , &log_data)	  
	}else{
	   ;
	}
    
}

func ( v *Redis_Monitor_Type)log_memory(){
   value,err := v.client.Info(v.ctx,"Memory").Result()
   if err == nil {
       log_value_data := generate_map(value)
	   all_data := generate_floats(&log_value_data)
	   log_data := make(map[string]float64)
       log_data["maxmemory"] =all_data["maxmemory"]
       log_data["used_memory"] = all_data["used_memory"]	   
	   v.push_data( "MEMORY" , &log_data)	  
	}else{
	   ;
	}
    
}

 


func ( v *Redis_Monitor_Type)commandstats(){

    
      value,err := v.client.Info(v.ctx,"commandstats").Result()
	  
      if err == nil {
	   log_data := make(map[string]float64)
       log_value_data := generate_map(value)
	 
       for key, value := range log_value_data {
	      log_data[key] = extract_commandstats_data(value)
	     
	   }
       v.push_data( "REDIS_MONITOR_CMD_TIME_STREAM" , &log_data)	    
	}else{
	   ;
	}
    


}


func ( v *Redis_Monitor_Type)push_data(key string, log_value *map[string]float64){

  var b bytes.Buffer	
  msgpack.Pack(&b,(*log_value))
  current_value := b.String()
  v.streams[key].Xadd(current_value)


}	


func extract_commandstats_data( input string) float64 {
  var return_value float64 = 0
  temp1 := strings.Split(input,",")
  if len(temp1) >= 3 {
     temp2 := strings.Split(temp1[2],"=")
	 if len(temp2) == 2 {
	    value  ,err := strconv.ParseFloat(temp2[1], 64)
	    if err == nil{
		   return_value = value
		}
   	}
  }
  return return_value


}



func generate_map (input string )map[string]string{
   return_value := make(map[string]string)
   lines := strings.Split(input,"\r")
   for _,i := range lines{
     temp := strings.Split(i,":")
     if len(temp) == 2{
	   return_value[temp[0]] = temp[1]
	 }
  }
  return return_value


}


func process_key_data( input *map[string]string ) {

    for key,value := range (*input) {
	  temp := strings.Split(value,"keys=")
	  fields := strings.Split(temp[1],",")
	  (*input)[key] = fields[0]
		
	}


}


func generate_floats( input *map[string]string ) map[string]float64 {

   return_value := make(map[string]float64)
   for key ,value := range (*input) {
   
      temp  ,err := strconv.ParseFloat(value, 64)
	  
	  if err == nil {
	      return_value[key] = temp
	  }
	 
   
   }
   return return_value

}