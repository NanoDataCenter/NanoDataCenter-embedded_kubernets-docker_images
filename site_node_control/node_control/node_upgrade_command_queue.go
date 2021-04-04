package node_control

import "fmt"
import "time"
import "site_control.com/cf_control"


func node_command_queue_structures(site_data *map[string]interface{}){

   // initialize data structures
   
   
}




func  initialize_node_job_server_watch_dog_cf(cf_cluster *cf.CF_CLUSTER_TYPE){

   var cf_control  cf.CF_SYSTEM_TYPE

   (cf_control).Init(cf_cluster ,"node_command_queue_watch_dog",true, int64(time.Second))
   
   
   
   (cf_control).Add_Chain("container_monitoring",true)   // watch dog strobe
   
   var parameters = make(map[string]interface{})
   ( cf_control).Cf_add_one_step(node_strobe_watch_dog,parameters)
   (cf_control).Cf_add_wait_interval(int64(time.Second*14)  ) // every 15 seconds
   (cf_control).Cf_add_reset()
  
   (cf_control).Add_Chain("monitor_node_command_queue",true) // monitor command from site_contol
   (cf_control).Cf_add_log_link("monitor_node_command_queue")
   parameters = make(map[string]interface{}) 
   (cf_control).Cf_add_unfiltered_element(node_process_job_queue,parameters)
   (cf_control).Cf_add_reset()
   
   
}

func node_strobe_watch_dog( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{

  fmt.Println("node stobe watch dog ")
  return cf.CF_DISABLE
  
}
  
  
  
func node_process_job_queue( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{

  fmt.Println("monitor and process job queue ")
  return cf.CF_HALT
  
}