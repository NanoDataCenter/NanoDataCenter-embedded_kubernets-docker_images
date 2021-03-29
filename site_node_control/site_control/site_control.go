package site_control

/*

import "fmt"
import "time"
import "site_control.com/site_data"
import "site_control.com/docker_management"
import "site_control.com/redis_support/graph_query"
import  "site_control.com/redis_support/generate_handlers"
import  "site_control.com/redis_support/redis_handlers"




var site_data map[string]interface{}








*/

import "time"
import "site_control.com/docker_management"
import "site_control.com/cf_control"

var cf_control  cf.CF_SYSTEM

func Site_control_startup(site_data *map[string]interface{}){
   init_docker_management(site_data)


}


var docker_handle docker_management.Docker_Handle_Type
  
 
func init_docker_management(site_data *map[string]interface{}){

    
	var container_search_list = []string{"SITE_CONTROL:SITE_CONTROL"}
    var display_struct_search_list = []string{"SITE_CONTROL:SITE_CONTROL","DOCKER_CONTROL"}
    docker_management.Initialize_Docker_Monitor( &docker_handle, &container_search_list, &display_struct_search_list,site_data)
}

func  Initialize_CF(){


   cf.Init(&cf_control)
   
   cf.Add_Chain(&cf_control,"container_monitoring",true)
   cf.Cf_add_log_link(&cf_control,"container_monitor_loop")
	  // add one step
   cf.Cf_add_wait_interval(&cf_control,int64(time.Minute*5)  )
   cf.Cf_add_reset(&cf_control)
   
   cf.Add_Chain(&cf_control,"container_performance_logs",true)
   cf.Cf_add_log_link(&cf_control,"container_monitor_loop")
	 // add one step
   cf.Cf_add_wait_interval(&cf_control,int64(time.Minute*15)  )
   cf.Cf_add_reset(&cf_control)
   panic("done")
   
}	