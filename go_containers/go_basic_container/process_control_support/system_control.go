package system_control

import "fmt"
import "time"
import "bytes"
//import "lacima.com/redis_support/redis_handlers"
//import "lacima.com/redis_support/generate_handlers"
//import "lacima.com/system_error_logging"
import "lacima.com/redis_support/graph_query"
import "lacima.com/cf_control"
import "github.com/msgpack/msgpack-go"
import "lacima.com/Patterns/logging_support"

type System_Control_Type struct {
  status                   bool
  container_name           string
  incident_log             *logging_support.Incident_Log_Type
  process_map              map[string]string
  process_ctrl             []*Process_Manager_Type
  process_status           map[string]bool
  error_status             map[string]string
  good_count               map[string]int
}



func Construct_System_Control(  container_name string ) *System_Control_Type {

   var return_value  System_Control_Type
   return_value.container_name = container_name
   return_value.process_map = make(map[string]string)
   return_value.process_status  = make(map[string]bool)
   return_value.error_status  = make(map[string]string)
   return_value.good_count = make(map[string]int)
   return &return_value
}   

func ( v *System_Control_Type) Init(cf_cluster *cf.CF_CLUSTER_TYPE){

   v.incident_log = logging_support.Construct_incident_log([]string{"CONTAINER:"+v.container_name,"INCIDENT_LOG:managed_process_failure","INCIDENT_LOG"} )
    
   search_list := []string{ "CONTAINER:"+v.container_name  }
   
   nodes := graph_query.Common_qs_search( &search_list)
   node:= nodes[0]
   fmt.Println("node",node) 
   process_map_json := node["command_map"]
   process_map := graph_query.Convert_json_dict(process_map_json)
   fmt.Println("process_map",process_map) 
   
   v.verify_process_map(process_map)
   v.construct_chains(cf_cluster)
   
  
}



func ( v *System_Control_Type ) verify_process_map( process_map map[string]string ) {

   
   
   for key,command := range process_map { 
 
	v.good_count[key] = 10
    v.process_status[key] = true
	v.error_status[key]   = ""
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
/*
   var data = time.Now().UnixNano()
   var b bytes.Buffer	
   msgpack.Pack(&b,data)
   v.watchdog_strobe.Set(b.String())

*/
   return cf.CF_DISABLE
}

func ( v *System_Control_Type )monitor_process( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{
   v.status = true
   for _,element := range v.process_ctrl {
       v.monitor_element(element)
      
   }
   v.summarize_data()
   return cf.CF_DISABLE
}

func ( v *System_Control_Type )monitor_element( element  *Process_Manager_Type){

   key := (*element).key
   if v.good_count[key] == 10 {
        v.process_status[key] = true
   }else{
      v.process_status[key] = false
   }
   v.error_status[key] = ""
   if (*element).failed == true {
      v.status = false
      v.process_status[key] = false
	  v.error_status[key] = (*element).error_log
	  v.good_count[key] = 0
   }else{
     v.good_count[key] +=1
   }
   
}



func ( v *System_Control_Type )summarize_data(){
  
     var b bytes.Buffer	
     msgpack.Pack(&b,v.error_status)
	 current_error := b.String()
	 
     var b1 bytes.Buffer	
     msgpack.Pack(&b1,v.process_status)
	 new_value := b1.String()	 
	 v.incident_log.Log_data( v.status, new_value, current_error)
  
}
  

