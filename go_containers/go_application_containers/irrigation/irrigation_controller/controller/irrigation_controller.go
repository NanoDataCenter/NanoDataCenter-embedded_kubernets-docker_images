package main

import (
    
    //"fmt"
    //"reflect"
	"time"
    //"encoding/json"
    "lacima.com/site_data"
    "lacima.com/redis_support/graph_query"
    "lacima.com/redis_support/redis_handlers"
    "lacima.com/redis_support/generate_handlers"
    "lacima.com/cf_control"
	"lacima.com/go_application_containers/irrigation/irrigation_controller/irrigation_modules"


)



var  CF_site_node_control_cluster cf.CF_CLUSTER_TYPE




func main() {
  
 
  var config_file = "/data/redis_server.json"
  var site_data_store map[string]interface{}

  site_data_store = get_site_data.Get_site_data(config_file)
  graph_query.Graph_support_init(&site_data_store)
  redis_handlers.Init_Redis_Mutex()
  data_handler.Data_handler_init(&site_data_store)
  
  
  (CF_site_node_control_cluster).Cf_cluster_init()
  (CF_site_node_control_cluster).Cf_set_current_row("irrigation_scheduling")
  
  /*
   *
   *  Initialize Modules
   *
  */
  irrigation_modules.Initialize_channels()
  irrigation_modules.Irrigation_rpc_setup()
  
  /* 
  ** start modules
  **
  */
  irrigation_modules.Irrigation_rpc_start()
  
  
   for true {
     time.Sleep(time.Minute) //main loop spin
   }

}


func execute(){
  
  (CF_site_node_control_cluster).CF_Fork()
} 
 