

package docker_management


import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/graph_query"
import "lacima.com/Patterns/logging_support"
import "lacima.com/server_libraries/postgres"





type Docker_Handle_Type struct {
    
  containers []string
  container_set map[string]bool
  hash_status redis_handlers.Redis_Hash_Struct  
  incident_stream *logging_support.Incident_Log_Type
  logging_stream  pg_drv.Postgres_Stream_Driver
  
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
   
   
 
}




func (docker_handle *Docker_Handle_Type) initialize_docker_container_monitoring(display_struct_search_list, incident_search_list  *[]string ){
  
  
  
  (docker_handle).container_set = make(map[string]bool)
  var data_structures =  data_handler.Construct_Data_Structures(display_struct_search_list)

   
  (docker_handle).hash_status = (*data_structures)["DOCKER_DISPLAY_DICTIONARY"].(redis_handlers.Redis_Hash_Struct)
  (docker_handle).incident_stream = logging_support.Construct_incident_log((*incident_search_list))
  (docker_handle).logging_stream  = logging_support.Find_stream_logging_driver()
  (docker_handle).hash_status.Delete_All()  // table will be repopulated later on
}















