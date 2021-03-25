package docker_management

import "fmt"
import "time"
//import "site_control.com/redis_support/graph_query"
//import  "site_control.com/redis_support/generate_handlers"
//import "site_control.com/docker_control"
;;
const monitor_delay = time.Second*15
const performance_delay = time.Minute*15
var Monitored_containers  = make([]string,0)
var Docker_Display_Structures *map[string]interface{}
var Docker_status_handlers *map[string]map[string]interface{}
var site_ptr *map[string]interface{}

Docker_Display_Structures


func Docker_Monitor(){

  var loop_flag = true
  for loop_flag {
   fmt.Println("Docker_Monitor")
   time.Sleep(monitor_delay)
  }

}


func Docker_Performance_Monitor(){

  var loop_flag = true
  for loop_flag {
   fmt.Println("Docker_Performance_Monitor")
   time.Sleep(performance_delay)
  }

}

