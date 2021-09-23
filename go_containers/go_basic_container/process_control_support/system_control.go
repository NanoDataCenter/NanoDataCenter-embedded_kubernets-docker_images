package system_control

import "fmt"
import "time"
import  "encoding/json"
import "lacima.com/redis_support/graph_query"
import "lacima.com/cf_control"
import "lacima.com/Patterns/logging_support"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/Patterns/msgpack_2"


type incident_log_type struct{
    
    process_status           map[string]bool
    error_status             map[string]string  
    
}

type System_Control_Type struct {
  
  container_name           string
  incident_log             *logging_support.Incident_Log_Type
  watch_dog_log            *logging_support.Watch_Dog_Log_Type
  status_dict              redis_handlers.Redis_Hash_Struct
  process_map              map[string]string
  process_ctrl             []*Process_Manager_Type
  incident_data            incident_log_type
  
}



func Construct_System_Control(  container_name string ) *System_Control_Type {

   var return_value  System_Control_Type
   return_value.container_name   = container_name
   return_value.process_map      = make(map[string]string)
   return_value.incident_data.process_status   = make(map[string]bool)
   return_value.incident_data.error_status     = make(map[string]string)
   
   return &return_value
}   




func ( v *System_Control_Type) Init(cf_cluster *cf.CF_CLUSTER_TYPE){

   var search_path = []string{"CONTAINER:"+v.container_name,"CONTAINER_STRUCTURES"}
   handlers          := data_handler.Construct_Data_Structures(&search_path)
   v.status_dict     = (*handlers)["PROCESS_STATUS"].(redis_handlers.Redis_Hash_Struct)
   
   v.incident_log    = logging_support.Construct_incident_log([]string{"CONTAINER:"+v.container_name,"INCIDENT_LOG:managed_process_failure","INCIDENT_LOG"} )
   v.watch_dog_log   = logging_support.Construct_watch_data_log([]string{"CONTAINER:"+v.container_name,"WATCH_DOG:process_control","WATCH_DOG"})
    
   search_list := []string{ "CONTAINER:"+v.container_name  }
   
   nodes := graph_query.Common_qs_search( &search_list)
   node:= nodes[0]
   //fmt.Println("node",node) 
   process_map_json := node["command_map"]
   process_map := graph_query.Convert_json_dict(process_map_json)
   //fmt.Println("process_map",process_map) 
   
   v.initialize_process_map(process_map)
   v.construct_chains(cf_cluster)
   
  
}



func ( v *System_Control_Type ) initialize_process_map( process_map map[string]string ) {

   
   
   for key,command := range process_map { 
    
	v.status_dict.HSet(key,msg_pack_utils.Pack_bool(true))
    v.incident_data.process_status[key] = true
	v.incident_data.error_status[key]   = ""
    v.process_map[key] = command
	v.process_ctrl = append(v.process_ctrl,construct_process_manager( key,command) )
   
   }


}




func ( v *System_Control_Type ) construct_chains(cf_cluster *cf.CF_CLUSTER_TYPE) {

   var cf_control  cf.CF_SYSTEM_TYPE

   (cf_control).Init(cf_cluster ,"container_process_monitor",true, time.Second)
   
   (cf_control).Add_Chain("initialization",true)
   //(cf_control).Cf_add_log_link("cf initialization chain")
   var parameters = make(map[string]interface{})
   (cf_control).Cf_add_one_step(v.launch_processes,parameters)
   (cf_control).Cf_add_enable_chains_links( []string{"monitor_active_processes"}  )
   (cf_control).Cf_add_terminate()  


   (cf_control).Add_Chain("strobe_watchdog",true)
   //(cf_control).Cf_add_log_link("strobe_watch_dog")
   var parameters1 = make(map[string]interface{})
   (cf_control).Cf_add_one_step(v.strobe_watch_dog,parameters1)
   (cf_control).Cf_add_wait_interval(time.Second*5  )
   (cf_control).Cf_add_reset()

  
    (cf_control).Add_Chain("monitor_active_processes",false)
   //(cf_control).Cf_add_log_link("strobe_watch_dog")
    var parameters2 = make(map[string]interface{})
   (cf_control).Cf_add_one_step(v.monitor_process,parameters2)
   (cf_control).Cf_add_wait_interval(time.Second*10  )
   (cf_control).Cf_add_reset()

}
   
func ( v *System_Control_Type )launch_processes( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{

  
   for _,element := range v.process_ctrl {
   
     go (*element).run()
   
   }
   
   
   return cf.CF_DISABLE
}

func ( v *System_Control_Type )strobe_watch_dog( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{
   v.watch_dog_log.Strobe_Watch_Dog(  )
   return cf.CF_DISABLE
}

func ( v *System_Control_Type )monitor_process( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{
   update := false
   for _,element := range v.process_ctrl {
       if v.monitor_element(element) == true {
           update = true
       }
      
   }
   if update == true {
       v.log_incident_data()
   }
   return cf.CF_DISABLE
}
 
func ( v *System_Control_Type )monitor_element( element  *Process_Manager_Type) bool{

   update := false
  
   ref := v.incident_data.process_status[element.key]
   
   select {
      case  msg :=  <- element.output:
          v.incident_data.process_status[element.key] = false
          update = true
          if msg != v.incident_data.error_status[element.key]  {
              v.incident_data.error_status[element.key] = msg
             
              v.incident_data.process_status[element.key] = false
          }
        default:
             v.incident_data.error_status[element.key] = ""
             v.incident_data.process_status[element.key] = true
     
   }
   if ref != v.incident_data.process_status[element.key] {
        v.status_dict.HSet(element.key,msg_pack_utils.Pack_bool(v.incident_data.process_status[element.key]))
   }
   return update
   
}



func ( v *System_Control_Type )log_incident_data(){
     
     bytes_1, _ := json.Marshal(v.incident_data.process_status )
     bytes_2, _ := json.Marshal(v.incident_data.error_status )
     incident_data_string := string(bytes_1)+" \n" + string(bytes_2)
     fmt.Println("incident_data",incident_data_string)
	 v.incident_log.Log_data( incident_data_string)
  
}

