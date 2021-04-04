package site_control


import "fmt"
import "time"

import "site_control.com/cf_control"
import  "site_control.com/redis_support/redis_handlers"
//import "site_control.com/redis_support/graph_query"
//import "github.com/msgpack/msgpack-go"


var site_control_node_list []string
var site_container_map map[string][]string   // each container a map of nodes that it is in

var site_input_queue *redis_handlers.Redis_Single_Structure
var site_node_status *redis_handlers.Redis_Hash_Struct
var site_monitoring_containers  *redis_handlers. Redis_Hash_Struct

var site_container_control_structs map[string]map[string]interface{}


func initialize_site_monitoring_data_structures(site_data *map[string]interface{}){

   site_monitoring_find_site_container_map(site_data)
   site_monitoring_find_local_structures(site_data)
   site_monitoring_find_node_data_structures(site_data)
  
  
}


func  site_monitoring_find_site_container_map(site_data *map[string]interface{}){
  site_container_map = make(map[string][]string)
  // find nodes amd 
  // check node for site["local"]
  // get system containers put in local array
  // for each node  -- also fill up container map
  //   fill up map
}

func     site_monitoring_find_local_structures(site_data *map[string]interface{}){


}

func     site_monitoring_find_node_data_structures(site_data *map[string]interface{}){


   
}




func initialize_site_monitoring_chains(cf_cluster *cf.CF_CLUSTER_TYPE){

  var cf_control  cf.CF_SYSTEM_TYPE

   (cf_control).Init(cf_cluster ,"site_control_monitor_nodes",true, int64(time.Second))
   
   
   
   (cf_control).Add_Chain("site_control_monitor_watch_dogs",true)   // watch dog strobe
   
   var parameters = make(map[string]interface{})
   ( cf_control).Cf_add_one_step(site_control_monitor_watch_dog,parameters)
   (cf_control).Cf_add_wait_interval(int64(time.Second*14)  ) // every 15 seconds
   (cf_control).Cf_add_reset()
  
   (cf_control).Add_Chain("monitor_site_command_queue",true) // monitor command from site_contol
   (cf_control).Cf_add_log_link("monitor_site_command_queue")
   parameters = make(map[string]interface{}) 
   (cf_control).Cf_add_unfiltered_element(site_control_input_queue,parameters)
   (cf_control).Cf_add_reset()
   
   
}

func site_control_monitor_watch_dog( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{

  fmt.Println("site_control check watch dog ")
  return cf.CF_DISABLE
  
}
  
  
  
func site_control_input_queue( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{

  fmt.Println(time.Now())
  return cf.CF_HALT
  
}