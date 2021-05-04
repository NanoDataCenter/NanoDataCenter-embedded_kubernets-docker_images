package irrigation_modules

import "bytes"
import "sync"

import	"github.com/msgpack/msgpack-go"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/Patterns/msgpack"


var irr_rpc_queue sync.Mutex



func lock_irrigation_rpc_queue(){

irr_rpc_queue.Lock()

}


func unlock_irrigation_rpc_queue(){

irr_rpc_queue.Unlock()

}
   


type handler_type func ( parameters map[string]interface{} ) map[string]interface{}

var driver        redis_handlers.Redis_RPC_Struct
var job_queue      redis_handlers.Redis_Job_Queue 
var rpc_command_map   map[string]handler_type  






func Get_next_job()( *map[string]interface{}, bool){
  lock_irrigation_rpc_queue()
  defer unlock_irrigation_rpc_queue()
  if job_queue.Length() == 0{
    return nil,false
  }
  data := job_queue.Pop()
  return_value  := msgpack_utils.Convert_rpc_return(data)
  return &return_value, true
  
}


func Irrigation_rpc_setup( ){
   
   search_list := []string{"IRRIGIGATION_CONTROL:IRRIGIGATION_CONTROL"} // fix this
   handlers := data_handler.Construct_Data_Structures(&search_list)
   driver = (*handlers)["IRRIGATION_JOB_SERVER"].(redis_handlers.Redis_RPC_Struct)
   job_queue = (*handlers)["INTERNAL_IRRIGATION_JOB_QUEUE"].(redis_handlers.Redis_Job_Queue)
   
   driver.Add_handler("PING",ping)
   driver.Add_handler("QUEUE_COMMAND",rpc_command_handler)                                
   driver.Add_handler("QUEUE_IRRIGATION_CONTROLLER_PIN",queue_controller_pin)
   driver.Add_handler("QUEUE_OFFLINE_CONTROLLER_PIN",queue_offline_controller_pin)
   driver.Add_handler("DELETE_IRRIGATION_JOBS",delete_irrigation_jobs) 
  
   rpc_command_map  = make(map[string]handler_type)
   rpc_command_map["CLEAN_FILTER"]              = clean_filter                
   rpc_command_map["OPEN_MASTER_VALVE"]         = open_master_valve          
   rpc_command_map["CLOSE_MASTER_VALVE"]        = close_master_valve                         
   rpc_command_map["CHECK_OFF"]                = check_off       
   rpc_command_map["RESISTANCE_CHECK"]          = check_resistance  
   rpc_command_map["CLEAR_IRRIGATION_QUEUE"]    = clear_queue
   rpc_command_map["CLEAR_OFFLINE_IRRIGATION"]  = clear_offline
   rpc_command_map["SUSPEND"]                   = suspend                                 
   rpc_command_map["RESUME"]                    = resume    
   rpc_command_map["SKIP_STATION"]              = skip_job     
}
 
 

   
   
func Irrigation_rpc_start(){
 
  go driver.Json_Rpc_start()
} 


func rpc_command_handler(parameters map[string]interface{} ) map[string]interface{}{

  command := parameters["COMMAND"].(string)
  if handler, ok :=  rpc_command_map[command]; ok == true{
      return handler(parameters)
  }
  parameters["status"] = false // unregistered command
  return parameters  
}

func ping( parameters map[string]interface{} ) map[string]interface{}{

   parameters["status"] = true
   return parameters

}

func clear_queue( parameters map[string]interface{} ) map[string]interface{}{
   lock_irrigation_rpc_queue()
   defer unlock_irrigation_rpc_queue()
   job_queue.Delete_all()
   send_Command_Channel("clear_queue")
   parameters["status"] = true
   return parameters

}
func clear_offline( parameters map[string]interface{} ) map[string]interface{}{
   lock_irrigation_rpc_queue()
   defer unlock_irrigation_rpc_queue()
   send_Command_Channel("clear_offline")
   parameters["status"] = true
   return parameters

}
func suspend( parameters map[string]interface{} ) map[string]interface{}{
   lock_irrigation_rpc_queue()
   defer unlock_irrigation_rpc_queue()
   send_Command_Channel("suspend")
   parameters["status"] = true
   return parameters

}                                      
func resume( parameters map[string]interface{} ) map[string]interface{}{
   lock_irrigation_rpc_queue()
   defer unlock_irrigation_rpc_queue()
   send_Command_Channel("resume")
   parameters["status"] = true
   return parameters

}                                     
func skip_job( parameters map[string]interface{} ) map[string]interface{}{
   lock_irrigation_rpc_queue()
   defer unlock_irrigation_rpc_queue()
   send_Command_Channel("skip_job")
   parameters["status"] = true
   return parameters

}    
func queue_controller_pin( parameters map[string]interface{} ) map[string]interface{}{
   lock_irrigation_rpc_queue()
   defer unlock_irrigation_rpc_queue()
   command_map := make(map[string]interface{})
   command_map["command"] = "QUEUE_IRRIGATION_CONTROLLER_PIN"
   command_map["io"]     = parameters["io"].(map[string]int64)
   command_map["time"]   = parameters["time"].(int64)
   irrigation_queue_job(&command_map)
   parameters["status"] = true
   return parameters

}
func queue_offline_controller_pin( parameters map[string]interface{} ) map[string]interface{}{
   lock_irrigation_rpc_queue()
   defer unlock_irrigation_rpc_queue()
   if translate_offline_job(parameters){
       translate_offline_job(parameters)
       send_Command_Channel("offline_job")
	   parameters["status"] = true
   }else{
      parameters["status"] = false
   }

   return parameters

}

func clean_filter( parameters map[string]interface{} ) map[string]interface{}{
   lock_irrigation_rpc_queue()
   defer unlock_irrigation_rpc_queue()
   command_map := make(map[string]interface{})
   command_map["command"] = "CLEAN_FILTER"
   irrigation_queue_job(&command_map)
   parameters["status"] = true
   return parameters

}               
func open_master_valve( parameters map[string]interface{} ) map[string]interface{}{
   lock_irrigation_rpc_queue()
   defer unlock_irrigation_rpc_queue()
   send_Command_Channel("open_master_valve")
   parameters["status"] = true
   return parameters

}           
func close_master_valve( parameters map[string]interface{} ) map[string]interface{}{
   lock_irrigation_rpc_queue()
   defer unlock_irrigation_rpc_queue()
   send_Command_Channel("close_master_valve")
   parameters["status"] = true
   return parameters

}                         
func check_off( parameters map[string]interface{} ) map[string]interface{}{
  lock_irrigation_rpc_queue()
  defer unlock_irrigation_rpc_queue()
  command_map := make(map[string]interface{})
  command_map["command"] = "CHECK_OFF"
  irrigation_queue_job(&command_map)
  parameters["status"] = true
  return parameters

}     
func check_resistance( parameters map[string]interface{} ) map[string]interface{}{
   lock_irrigation_rpc_queue()
   defer unlock_irrigation_rpc_queue()
   command_map := make(map[string]interface{})
   command_map["command"] = "RESISTANCE_CHECK"
   irrigation_queue_job(&command_map)
   parameters["status"] = true
   return parameters

}    

func delete_irrigation_jobs( parameters map[string]interface{} ) map[string]interface{}{
   lock_irrigation_rpc_queue()
   defer unlock_irrigation_rpc_queue()
   parameters["status"] = true
   jobs := parameters["index_list"].([]int64)
   
   job_queue.Delete_jobs(jobs)
   parameters["status"] = true
   return parameters

}    





/*
  support routines
  
*/
func irrigation_queue_job(input *map[string]interface{}){

  var b bytes.Buffer	
  msgpack.Pack(&b,input)
  current_value := b.String()
  job_queue.Push(current_value)
}




