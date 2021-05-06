package scheduling_utilities

import      "fmt"
import      "time"
//import      "reflect"
import     "bytes"
import      "encoding/json"

import "github.com/msgpack/msgpack-go"


import	    "lacima.com/go_application_containers/irrigation/irrigation_libraries/irrigation_queue_library"
import		"lacima.com/go_application_containers/irrigation/irrigation_libraries/irrigation_files_library"
import		"lacima.com/cf_control"
import      "lacima.com/redis_support/redis_handlers"
import      "lacima.com/Patterns/logging_support"
import      "lacima.com/redis_support/generate_handlers"

var  CF_site_node_control_cluster *cf.CF_CLUSTER_TYPE

type system_scheduling_type struct {
    scheduling_file       string
    completion_hash       redis_handlers.Redis_Hash_Struct
    scheduling_array     []map[string]interface{}
	incident_log          *logging_support.Incident_Log_Type
    iq                   irrigation_rpc.Irrigation_Client_Type
}


type Scheduling_Type struct {
  ok_flag              bool
  fs                  irr_files.Irrigation_File_Manager_Type 
  iq                  irrigation_rpc.Irrigation_Client_Type
  system_control      system_scheduling_type
  irrigation_control  system_scheduling_type
  scheduling_control  redis_handlers.Redis_Hash_Struct  

}

 
 
 
var data_str  Scheduling_Type 
 
 
 
func Setup_Scheduling( ip string, port,file_db int, cf *cf.CF_CLUSTER_TYPE ){
    CF_site_node_control_cluster = cf
 
    (&data_str).irrigation_initialize_setup(ip , port,file_db )
}
 
 
func execute(){
  
  (CF_site_node_control_cluster).CF_Fork()
} 

 
 
func ( v* Scheduling_Type ) irrigation_initialize_setup(ip string, port,file_db int){

    v.fs = irr_files.Initialization( ip , port,file_db )
    v.iq = irrigation_rpc.Irrigation_RPC_Client_Init(&[]string{"IRRIGIGATION_CONTROL"})
	v.system_control.iq = v.iq
	v.irrigation_control.iq = v.iq
	v.iq.Ping()
	
    search_path := []string{"IRRIGIGATION_SCHEDULING:IRRIGIGATION_SCHEDULING","IRRIGIGATION_SCHEDULING"}
    handlers := data_handler.Construct_Data_Structures(&search_path)
    
	v.irrigation_control.completion_hash = (*handlers)["IRRIGATION_COMPLETION_DICTIONARY"].(redis_handlers.Redis_Hash_Struct)
	v.system_control.completion_hash     = (*handlers)["SYSTEM_COMPLETION_DICTIONARY"].(redis_handlers.Redis_Hash_Struct)
    v.scheduling_control                 = (*handlers)["SCHEDULING_CONTROL"].(redis_handlers.Redis_Hash_Struct)
    v.irrigation_control.scheduling_file = "sprinkler_ctrl.json"
    v.system_control.scheduling_file     = "system_actions.json"
	
	
	v.irrigation_control.incident_log  = logging_support.Construct_incident_log([]string{"IRRIGIGATION_SCHEDULING","IRRIGATION_ERROR_LOG","INCIDENT_LOG"} )
	v.system_control.incident_log  = logging_support.Construct_incident_log([]string{"IRRIGIGATION_SCHEDULING","SYSTEM_ERROR_LOG","INCIDENT_LOG"} )
    v.construct_chain()
	
}

func (v *system_scheduling_type) error_logger(){

    if r := recover(); r != nil {
	     fmt.Println("Done flag in f", r)
         value := fmt.Sprint(r)  
		 var b bytes.Buffer	
         msgpack.Pack(&b,value)
         current_value := b.String()
 
		 v.incident_log.Log_data( false,  current_value, current_value )
    }



} 
	
func (v *Scheduling_Type)system_error_logger(){

    if r := recover(); r != nil {
	     fmt.Println("System Recovered in f", r)
         value := fmt.Sprint(r)  
		 var b bytes.Buffer	
         msgpack.Pack(&b,value)
         current_value := b.String()
 
		 v.system_control.incident_log.Log_data( false,  current_value, current_value )
    }



}


func (v *Scheduling_Type)irrigation_error_logger(){

    if r := recover(); r != nil {
	     fmt.Println("Irrigation Recovered in f", r)
         value := fmt.Sprint(r)  
		 var b bytes.Buffer	
         msgpack.Pack(&b,value)
         current_value := b.String()
 
		 v.irrigation_control.incident_log.Log_data( false,  current_value, current_value )
    }



}
	
func (v* Scheduling_Type)construct_chain(){
  
   var cf_control  cf.CF_SYSTEM_TYPE
  (cf_control).Init(CF_site_node_control_cluster , "irrigation_scheduling" ,true, time.Minute )




  (cf_control).Add_Chain("irrigation_scheduling",true)
  //(cf_control).Cf_add_log_link("scheduling")
   
   
    (cf_control).Cf_add_one_step(v.action_check_for_system_activity,  make(map[string]interface{}))
   (cf_control).Cf_add_one_step(v.sched_check_for_schedule_activity, make(map[string]interface{})) 
  
  (cf_control).Cf_add_wait_interval(time.Minute )
  (cf_control).Cf_add_reset()
  

	
	

}



func (v *Scheduling_Type)sched_check_for_schedule_activity( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{

   v.ok_flag = true
   v.check_for_rain_flag()
   v.retrieve_irrigation_data()
   v.irrigation_schedule()
   v.irrigation_check_for_done_flag()

   return cf.CF_DISABLE
	  
}

func (v *Scheduling_Type)action_check_for_system_activity( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{
   fmt.Println("made it here #1")
   v.ok_flag = true
   v.retrieve_system_data()
   v.system_schedule()
   v.system_check_for_done_flag()

 
 
  return cf.CF_DISABLE  
}

func (v *Scheduling_Type)check_for_rain_flag(){
   if v.ok_flag == false{
      return
   }

   data := v.scheduling_control.HGet("RAIN_FLAG")
   if data != "true" {
      v.ok_flag = true
	  v.scheduling_control.HSet("RAIN_FLAG","false")
   }else{
       v.ok_flag = false
   }
}


func (v *Scheduling_Type)retrieve_irrigation_data(){
   if v.ok_flag == false{
      return
   }
  defer v.irrigation_error_logger()
  v.irrigation_control.scheduling_array, v.ok_flag =  v.decode_json_scheduling_files(v.irrigation_control.scheduling_file)
   
}

func (v *Scheduling_Type)irrigation_schedule(){
   if v.ok_flag == false{
      return
   }
   v.check_for_irrigation_activity()
 
   
}

func (v *Scheduling_Type)irrigation_check_for_done_flag(){
   if v.ok_flag == false{
      return
   }
 

   (v.irrigation_control).irrigation_clear_done_flag()
   
}

func (v *Scheduling_Type)retrieve_system_data(){
   fmt.Println("made it here #1",v.ok_flag)
   if v.ok_flag == false{
      return
   }
   defer v.system_error_logger()
   v.system_control.scheduling_array, v.ok_flag =  v.decode_json_scheduling_files(v.system_control.scheduling_file)
   
   
   
}

func (v *Scheduling_Type)system_schedule(){
   if v.ok_flag == false{
      return
   }

   v.check_for_system_activity()
   
}

func (v *Scheduling_Type)system_check_for_done_flag(){
   if v.ok_flag == false{
      return
   }

   (v.system_control).system_clear_done_flag()
   
}


func (v *Scheduling_Type)decode_json_scheduling_files( file_name string) ([]map[string]interface{},bool) {

   data,err := v.fs.Read_App_File( file_name)
   if err == false{
     // add error handler
     return nil,false
    }
   
   fmt.Println("data",err,data)
   var tmp []interface{}
   
   json.Unmarshal([]byte(data), &tmp)

   
   var result []map[string]interface{}
   for _,element  := range tmp   {
      result = append(result, element.(map[string]interface{}) )
   }
   return result,err



}
