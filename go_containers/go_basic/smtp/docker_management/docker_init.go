

package docker_management

import "fmt"

import "site_control.com/redis_support/graph_query"
import  "site_control.com/redis_support/generate_handlers"





func Initialize_Docker_Monitor( container_search_list *[]string , site_data *map[string]interface{} ){

   site_ptr = site_data
   find_containers( container_search_list, &Monitored_containers )
   
   find_container_properties( &Monitored_containers,&Monitored_container_properties)
   
   //fmt.Println("containers",Monitored_containers)
   //fmt.Println("properties",Monitored_container_properties)
   find_container_data_structures( &Monitored_containers,&Docker_status_handlers)
   fmt.Println("docker status handlers",Docker_status_handlers)
}



func  find_containers(search_list *[]string, containers *[]string){
    fmt.Println(*search_list)
    var site_nodes = graph_query.Common_qs_search(search_list)
    var site_node = site_nodes[0]
	//fmt.Print("site_node",site_node)
    *containers = graph_query.Convert_json_string_array(	site_node["containers"] ) 
}

func find_container_properties(container_list *[]string, Container_properties *map[string]map[string]string){
    var temp = make(map[string]map[string]string)
    for _,container := range *container_list{
	     var item = make(map[string]string)
	     var search_list = []string{"CONTAINER"+":"+container}
		 var container_nodes = graph_query.Common_qs_search(&search_list)
         var container_node = container_nodes[0]
		 //fmt.Println(container_node)
		 item["name"] =  graph_query.Convert_json_string(container_node["name"])
 		 item["container_image"] = graph_query.Convert_json_string(container_node["container_image"])
		 item["startup_command"] = graph_query.Convert_json_string(container_node["startup_command"])

		 temp[container] = item
		 
	}
	*Container_properties = temp
}


func find_container_data_structures(container_list *[]string, docker_handlers **map[string]map[string]interface{}){
    var temp = make(map[string]map[string]interface{})
    for _,container := range *container_list{
	     
	     var search_list = []string{"CONTAINER"+":"+container,"DATA_STRUCTURES"}
		 var data_element = data_handler.Construct_Data_Structures(&search_list)
		 temp[container] = *data_element
	}
	*docker_handlers = &temp
}