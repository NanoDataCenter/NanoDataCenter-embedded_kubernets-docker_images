

package docker_management


import  "lacima.com/redis_support/generate_handlers"
import  "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/graph_query"
import "lacima.com/Patterns/logging_support"





type Docker_Handle_Type struct {
    
  containers []string
  container_set map[string]bool
  hash_status redis_handlers.Redis_Hash_Struct  
  incident_stream *logging_support.Incident_Log_Type
  docker_performance_drivers map[string]map[string]redis_handlers.Redis_Stream_Struct
}



   
var site_ptr *map[string]interface{}


func Find_containers(search_list *[]string )[]string{
   
   
    site_nodes := graph_query.Common_qs_search(search_list)
    site_node  := site_nodes[0]
    return graph_query.Convert_json_string_array( site_node["containers"] ) 
}

func (docker_handle *Docker_Handle_Type)Initialize_Docker_Monitor( container_list []string ,display_struct_search_list,incident_search_list *[]string, site_data *map[string]interface{}){

  
   docker_handle.containers = container_list
   site_ptr = site_data
   (docker_handle).initialize_docker_container_monitoring(display_struct_search_list,incident_search_list)
   (docker_handle).initialize_docker_performance_monitor()
   
 
}




func (docker_handle *Docker_Handle_Type) initialize_docker_container_monitoring(display_struct_search_list, incident_search_list  *[]string ){
  
  
  
  (docker_handle).container_set = make(map[string]bool)
  var data_structures =  data_handler.Construct_Data_Structures(display_struct_search_list)

   
  (docker_handle).hash_status = (*data_structures)["DOCKER_DISPLAY_DICTIONARY"].(redis_handlers.Redis_Hash_Struct)
  (docker_handle).incident_stream = logging_support.Construct_incident_log((*incident_search_list))
  (docker_handle).hash_status.Delete_All()  // table will be repopulated later on
}




func (docker_handle *Docker_Handle_Type) initialize_docker_performance_monitor(){


   
  // docker status handlers are needed for performance handling
  (*docker_handle).docker_performance_drivers = make(map[string]map[string]redis_handlers.Redis_Stream_Struct)
   (*docker_handle).find_container_data_structures( )
 

}



func (docker_handle *Docker_Handle_Type) find_container_data_structures( ){
   
	
    for _,container := range (*docker_handle).containers{
	     (docker_handle).container_set[container] = true // true is dummy value for the set
	     var search_list = []string{"CONTAINER"+":"+container,"STREAMING_LOG"}
		 var data_element = data_handler.Construct_Data_Structures(&search_list)
		 
		 (docker_handle).docker_performance_add_driver( container, "VSZ",  (*data_element)["vsz"].(redis_handlers.Redis_Stream_Struct) )

		 (docker_handle).docker_performance_add_driver(container, "RSS",  (*data_element)["rss"].(redis_handlers.Redis_Stream_Struct) )
		 (docker_handle).docker_performance_add_driver(container, "CPU",  (*data_element)["cpu"].(redis_handlers.Redis_Stream_Struct) )
	}
	
}


func (docker_handle *Docker_Handle_Type) docker_performance_add_driver(container,key string, value  redis_handlers.Redis_Stream_Struct){

   if _,ok := (*docker_handle).docker_performance_drivers[container];ok==false{
        (*docker_handle).docker_performance_drivers[container] = make(map[string]redis_handlers.Redis_Stream_Struct)
   }
   docker_handle.docker_performance_drivers[container][key] = value
   
}





