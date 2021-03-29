

package docker_management

import "fmt"
import "bytes"
import "time"
import "github.com/msgpack/msgpack-go"
import "site_control.com/redis_support/graph_query"
import  "site_control.com/redis_support/generate_handlers"
import  "site_control.com/redis_support/redis_handlers"


//type set_hash_dictionary_type func(driver *redis_handlers.Redis_Hash_Struct,field *string, value *map[string]bool)
//type log_handler_type func(driver *redis_handlers.Redis_Stream_Struct, value *map[string]interface{})



type Docker_Monitoring_Handle struct {

   
   hash_status redis_handlers.Redis_Hash_Struct  
   error_stream redis_handlers.Redis_Stream_Struct
   set_handler func(driver *redis_handlers.Redis_Hash_Struct,field string,value *map[string]bool)
   get_handler func(driver *redis_handlers.Redis_Hash_Struct,field string) *map[string]bool
   log_handler func(driver *redis_handlers.Redis_Stream_Struct, container string, action,status bool)
}

type docker_performance_log_driver_type  func(driver *redis_handlers.Redis_Stream_Struct,data map[string]interface{} )



type Docker_Performance_Handle struct {
 
  docker_performance_drivers map[string]map[string]redis_handlers.Redis_Stream_Struct
  log_driver docker_performance_log_driver_type
  
}
  

type Docker_Handle_Type struct {
    
  Containers []string
  Docker_monitoring  Docker_Monitoring_Handle
  docker_performance Docker_Performance_Handle
}



   



func Initialize_Docker_Monitor( docker_handle *Docker_Handle_Type, container_search_list *[]string ,display_struct_search_list  *[]string, site_data *map[string]interface{} ){

   site_ptr = site_data
  
   
   initialize_docker_container_monitoring(docker_handle,container_search_list,display_struct_search_list)
   initialize_docker_performance_monitor(docker_handle)
   fmt.Println("monitored containers",(*docker_handle).Containers)
   var driver = (*docker_handle).Docker_monitoring.hash_status
   var test = make(map[string]bool)
   test["active"] = false
   test["status"] = true
   (*docker_handle).Docker_monitoring.set_handler(&driver,"test",&test)
   var stream_driver = (*docker_handle).Docker_monitoring.error_stream
   (*docker_handle).Docker_monitoring.log_handler(&stream_driver,"test_container",false,true)
}


func initialize_docker_container_monitoring(docker_handle *Docker_Handle_Type, container_search_list *[]string ,display_struct_search_list  *[]string){
  

  find_containers( container_search_list,&(*docker_handle).Containers )
  var data_structures =  data_handler.Construct_Data_Structures(display_struct_search_list)

  (*docker_handle).Docker_monitoring.hash_status = (*data_structures)["DOCKER_DISPLAY_DICTIONARY"].(redis_handlers.Redis_Hash_Struct)
  (*docker_handle).Docker_monitoring.error_stream = (*data_structures)["ERROR_STREAM"].(redis_handlers.Redis_Stream_Struct)
  (*docker_handle).Docker_monitoring.set_handler = set_docker_display_dictionary
  (*docker_handle).Docker_monitoring.log_handler = log_docker_error_stream
  (*docker_handle).Docker_monitoring.get_handler = get_docker_display_dictionary
}

func  find_containers(search_list *[]string, containers *[]string){
    //fmt.Println(*search_list)
    var site_nodes = graph_query.Common_qs_search(search_list)
    var site_node = site_nodes[0]
	//fmt.Print("site_node",site_node)
    *containers = graph_query.Convert_json_string_array(	site_node["containers"] ) 
}

//                              func(driver *redis_handlers.Redis_Hash_Struct,field *string, value *map[string]bool)

func set_docker_display_dictionary(driver *redis_handlers.Redis_Hash_Struct,field string,value *map[string]bool){

    // convert bool to msgpack
    var b bytes.Buffer
	
	msgpack.Pack(&b,*value)
	var output = b.String()
	
    (*driver).HSet(field,output)

}

func get_docker_display_dictionary(driver *redis_handlers.Redis_Hash_Struct,field string)*map[string]bool{

   var msgpack_data string
   
   msgpack_data = (*driver).HGet(field)   

   var buf = bytes.NewBufferString(msgpack_data)
   v, _, err := msgpack.Unpack(buf)
   
	if err != nil {
		panic("bad msgpack data 1")
	}
	
	var return_value = make(map[string]bool)
	for key, value := range v.Interface().(map[interface{}]interface{}) {
	   return_value[key.(string)] = value.(bool)
	}
	return &return_value

}



func log_docker_error_stream(driver *redis_handlers.Redis_Stream_Struct, container string, action,status bool){
    
	var log = make(map[string]interface{})
	log["time"] = time.Now().UnixNano()
    log["container"] = container
	log["action"] = action
	log["status"] = status 
	
   // convert bool to msgpack
    var b bytes.Buffer
	
	msgpack.Pack(&b,log)
	var output = b.String()

	(*driver).Xadd(output)

}




func initialize_docker_performance_monitor(docker_handle *Docker_Handle_Type){


   
  // docker status handlers are needed for performance handling
  (*docker_handle).docker_performance.docker_performance_drivers = make(map[string]map[string]redis_handlers.Redis_Stream_Struct)
   find_container_data_structures( docker_handle)
 
  (*docker_handle).docker_performance.log_driver =   log_performance_data
/*
  var driver = (*docker_handle).docker_performance.docker_performance_drivers["file_server"]["PROCESS_VSZ"]

  var map_data = make(map[string]interface{})
  map_data["test"] = "XXXXX"
  map_data["per"] = 3
  (*docker_handle).docker_performance.log_driver( &driver,map_data)
  driver = (*docker_handle).docker_performance.docker_performance_drivers["file_server"]["PROCESS_RSS"]
  (*docker_handle).docker_performance.log_driver( &driver,map_data)
  driver = (*docker_handle).docker_performance.docker_performance_drivers["file_server"]["PROCESS_CPU"]
  (*docker_handle).docker_performance.log_driver( &driver,map_data)
*/
}



func find_container_data_structures(docker_handle *Docker_Handle_Type ){
   
	
    for _,container := range (*docker_handle).Containers{
	     
	     var search_list = []string{"CONTAINER"+":"+container,"DATA_STRUCTURES"}
		 var data_element = data_handler.Construct_Data_Structures(&search_list)
		 
		 docker_performance_add_driver(docker_handle, container, "PROCESS_VSZ",  (*data_element)["PROCESS_VSZ"].(redis_handlers.Redis_Stream_Struct) )

		 docker_performance_add_driver(docker_handle,container, "PROCESS_RSS",  (*data_element)["PROCESS_RSS"].(redis_handlers.Redis_Stream_Struct) )
		 docker_performance_add_driver(docker_handle,container, "PROCESS_CPU",  (*data_element)["PROCESS_CPU"].(redis_handlers.Redis_Stream_Struct) )
	}
	
}


func docker_performance_add_driver(docker_handle *Docker_Handle_Type, container,key string, value  redis_handlers.Redis_Stream_Struct){

   if _,ok := (*docker_handle).docker_performance.docker_performance_drivers[container];ok==false{
        (*docker_handle).docker_performance.docker_performance_drivers[container] = make(map[string]redis_handlers.Redis_Stream_Struct)
   }
   (*docker_handle).docker_performance.docker_performance_drivers[container][key] = value
   
}


func log_performance_data(driver  *(redis_handlers.Redis_Stream_Struct),data map[string]interface{} ) {
    
	data["time"] =  time.Now().UnixNano()
	   // convert bool to msgpack
    var b bytes.Buffer
	
	msgpack.Pack(&b,data)
	var output = b.String()
	fmt.Println("driver",driver)
	(*driver).Xadd(output)
	
	
}


