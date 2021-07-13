
package node_init

import ( 

"context"
"time"




"lacima.com/Patterns/logging_support"
"lacima.com/site_control_app/docker_control"
"lacima.com/redis_support/graph_query"


)


var site_data *map[string]interface{}
//type Database struct {
//   Client *redis.Client
//}
var ctx    = context.TODO()
var graph_container_script string
var graph_container_image string
var services_json string
//var container []string
var node_init_containers = make([]string,0)
var node_containers = make([]map[string]string,0)





func start_node_containers(){
  for _,value := range node_containers{
    if value["name"] == "redis" {
	  
	  continue
	}
	if docker_control.Image_Exists(value["container_image"]) == false{
	   
	   docker_control.Pull(value["container_image"])
	}
	docker_control.Container_rm(value["name"])
	docker_control.Container_up(value["name"],value["startup_command"])
  }
}  
	
   

func find_node_container_properties(){
    
    for _,container := range node_init_containers{
	     var item = make(map[string]string,0)
	     var search_list = []string{"CONTAINER"+":"+container}
		 var container_nodes = graph_query.Common_qs_search(&search_list)
         var container_node = container_nodes[0]
		
		 item["name"] =  graph_query.Convert_json_string(container_node["name"])
 		 item["container_image"] =graph_query.Convert_json_string(container_node["container_image"])
		 item["startup_command"] = graph_query.Convert_json_string(container_node["startup_command"])

		 node_containers = append(node_containers,item)
		 
	}
	
}



func  find_node_containers(){
    var search_list = []string{ "NODE:"+(*site_data)["local_node"].(string) }
    var site_nodes = graph_query.Common_qs_search(&search_list)
    var site_node = site_nodes[0]
 
    node_init_containers = graph_query.Convert_json_string_array(	site_node["containers"] ) 
    



}

func match_containers(running_containers []string, match_element string )bool{
  for _,value := range running_containers {
     if value == match_element {
	   return true
	 }
  }
  return false
}

func determine_hot_start() bool {

  var running_containers = docker_control.Containers_ls_runing()
  
  for _,name := range node_init_containers{
     if match_containers(running_containers,name) == false {

	   return false
	 }
  }
  
  
  return true
  

}

func verify_node_containers(){

 for _,value := range node_init_containers{
   if docker_control.Container_is_running(value )== false{
     panic("container "+value+"is not running")
   }
 }
}

	   
func Node_Init(  site_map *map[string]interface{} ){ 
   site_data = site_map                 
   find_node_containers()
   
   incident_log := logging_support.Construct_incident_log( []string{"NODE:"+(*site_data)["local_node"].(string),"INCIDENT_LOG:NODE_REBOOT","INCIDENT_LOG"} ) 
   incident_log.Post_event(false,"reboot","reboot")
  
   if determine_hot_start() == false {
      find_node_container_properties()
      start_node_containers()
	  time.Sleep(time.Second*5)
	  verify_node_containers()
      docker_control.Prune()
      
   }
  

}	
						 




    




