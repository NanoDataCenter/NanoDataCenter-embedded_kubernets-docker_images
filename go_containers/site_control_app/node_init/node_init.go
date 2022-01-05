
package node_init

import ( 

"context"
//"fmt"

"lacima.com/site_control_app/docker_control"
"lacima.com/redis_support/graph_query"


)


var site_data *map[string]interface{}
//type Database struct {
//   Client *redis.Client
//}



var ctx    = context.TODO()


var node_container_list = make([]string,0)
var node_container_properties = make([]map[string]string,0)


func rebuild_container(container_value map[string]string){
    
	if docker_control.Image_Exists(container_value["container_image"]) == false{
	   
	   panic("container image doesnot exit "+container_value["container_image"])
	}
	
	docker_control.Container_up(container_value["name"],container_value["startup_command"])

}


func start_node_containers(){
  for _,value := range node_container_properties{
     if docker_control.Container_is_running(value["name"] )== false{
         if docker_control.Exists(value["name"]) == false{
             rebuild_container(value)
         }else{
            
             docker_control.Container_start(value["name"])
         }
     }
  }
}  
	
   

func find_node_container_properties(){
    
    for _,container := range node_container_list{
	     var item = make(map[string]string,0)
	     var search_list = []string{"CONTAINER"+":"+container}
		 var container_nodes = graph_query.Common_qs_search(&search_list)
         var container_node = container_nodes[0]
		
		 item["name"] =  graph_query.Convert_json_string(container_node["name"])
 		 item["container_image"] =graph_query.Convert_json_string(container_node["container_image"])
		 item["startup_command"] = graph_query.Convert_json_string(container_node["startup_command"])

		 node_container_properties = append(node_container_properties,item)
		 
	}
	
	
}



func  find_node_containers(){
    var search_list = []string{ "NODE:"+(*site_data)["local_node"].(string) }
    var site_nodes = graph_query.Common_qs_search(&search_list)
    var site_node = site_nodes[0]
 
    node_container_list = graph_query.Convert_json_string_array(site_node["containers"] ) 
      



}







	   
   
	   
func Node_Init(  site_map *map[string]interface{} ){ 
 
   site_data = site_map                 
   find_node_containers()
   
   find_node_container_properties()
   start_node_containers()

   docker_control.Prune()

}	
						 




    




