

package docker_management

//import "fmt"
import  "site_control.com/redis_support/generate_handlers"
import  "site_control.com/redis_support/redis_handlers"
import "site_control.com/redis_support/graph_query"






type Docker_Handle_Type struct {
    
  containers []string
  container_set map[string]bool
  hash_status redis_handlers.Redis_Hash_Struct  
  error_stream redis_handlers.Redis_Stream_Struct
  docker_performance_drivers map[string]map[string]redis_handlers.Redis_Stream_Struct
}



   
var site_ptr *map[string]interface{}


func (docker_handle *Docker_Handle_Type)Initialize_Docker_Monitor( container_search_list *[]string ,display_struct_search_list  *[]string ,site_data *map[string]interface{}){

  
  
   site_ptr = site_data
   (docker_handle).initialize_docker_container_monitoring(container_search_list,display_struct_search_list)
   (docker_handle).initialize_docker_performance_monitor()
   
 
}


func (docker_handle *Docker_Handle_Type) initialize_docker_container_monitoring(container_search_list *[]string ,display_struct_search_list  *[]string, ){
  
  
  (docker_handle).find_containers( container_search_list,&(docker_handle).containers )
  (docker_handle).container_set = make(map[string]bool)
  var data_structures =  data_handler.Construct_Data_Structures(display_struct_search_list)

   
  (docker_handle).hash_status = (*data_structures)["DOCKER_DISPLAY_DICTIONARY"].(redis_handlers.Redis_Hash_Struct)
  (docker_handle).error_stream = (*data_structures)["ERROR_STREAM"].(redis_handlers.Redis_Stream_Struct)
 
}


func  (docker_handle *Docker_Handle_Type) find_containers(search_list *[]string, containers *[]string){
    
    var site_nodes = graph_query.Common_qs_search(search_list)
    var site_node = site_nodes[0]
    *containers = graph_query.Convert_json_string_array(	site_node["containers"] ) 
}

func (docker_handle *Docker_Handle_Type) initialize_docker_performance_monitor(){


   
  // docker status handlers are needed for performance handling
  (*docker_handle).docker_performance_drivers = make(map[string]map[string]redis_handlers.Redis_Stream_Struct)
   (*docker_handle).find_container_data_structures( )
 

}



func (docker_handle *Docker_Handle_Type) find_container_data_structures( ){
   
	
    for _,container := range (*docker_handle).containers{
	     (docker_handle).container_set[container] = true // true is dummy value for the set
	     var search_list = []string{"CONTAINER"+":"+container,"DATA_STRUCTURES"}
		 var data_element = data_handler.Construct_Data_Structures(&search_list)
		 
		 (docker_handle).docker_performance_add_driver( container, "PROCESS_VSZ",  (*data_element)["PROCESS_VSZ"].(redis_handlers.Redis_Stream_Struct) )

		 (docker_handle).docker_performance_add_driver(container, "PROCESS_RSS",  (*data_element)["PROCESS_RSS"].(redis_handlers.Redis_Stream_Struct) )
		 (docker_handle).docker_performance_add_driver(container, "PROCESS_CPU",  (*data_element)["PROCESS_CPU"].(redis_handlers.Redis_Stream_Struct) )
	}
	
}


func (docker_handle *Docker_Handle_Type) docker_performance_add_driver(container,key string, value  redis_handlers.Redis_Stream_Struct){

   if _,ok := (*docker_handle).docker_performance_drivers[container];ok==false{
        (*docker_handle).docker_performance_drivers[container] = make(map[string]redis_handlers.Redis_Stream_Struct)
   }
   (*docker_handle).docker_performance_drivers[container][key] = value
   
}





