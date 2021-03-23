package docker_management

import "fmt"
import "time"
import "site_control.com/redis_support/graph_query"
import  "site_control.com/redis_support/generate_handlers"
//import "site_control.com/docker_control"

const monitor_delay = time.Second*15
const performance_delay = time.Minute*15
var monitored_containers  = make([]string,0)
var monitored_container_properties = make([]map[string]string,0)

func Initialize_Docker_Monitor( container_search_list *[]string ){

   Find_Containers( container_search_list, &monitored_containers )
   
   Find_Container_Properties( &monitored_containers,&monitored_container_properties)
   
   fmt.Println("containers",monitored_containers)
   fmt.Println("properties",monitored_container_properties)
   Find_Container_Data_Structures( &monitored_containers)
}

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

func  Find_Containers(search_list *[]string, containers *[]string){
    fmt.Println(*search_list)
    var site_nodes = graph_query.Common_qs_search(search_list)
    var site_node = site_nodes[0]
	
    *containers = graph_query.Convert_json_string_array(	site_node["containers"] ) 
}

func Find_Container_Properties(container_list *[]string, container_properties *[]map[string]string){
    
    for _,container := range *container_list{
	     var item = make(map[string]string,0)
	     var search_list = []string{"CONTAINER"+":"+container}
		 var container_nodes = graph_query.Common_qs_search(&search_list)
         var container_node = container_nodes[0]
		 //fmt.Println(container_node)
		 item["name"] =  graph_query.Convert_json_string(container_node["name"])
 		 item["container_image"] = graph_query.Convert_json_string(container_node["container_image"])
		 item["startup_command"] = graph_query.Convert_json_string(container_node["startup_command"])

		 *container_properties = append(*container_properties,item)
		 
	}
	
}


func Find_Container_Data_Structures(container_list *[]string){
    for _,container := range *container_list{
	     
	     var search_list = []string{"CONTAINER"+":"+container,"DATA_STRUCTURES"}
		 data_handler.Construct_handler_definitions(&search_list)
		 
	}
	
}



