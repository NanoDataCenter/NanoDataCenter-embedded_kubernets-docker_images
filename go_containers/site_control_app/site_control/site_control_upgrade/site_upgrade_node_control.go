package site_control_upgrade


import "fmt"
import "time"
import "bytes"

import "lacima.com/cf_control"
import  "lacima.com/redis_support/generate_handlers"
import  "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/graph_query"
import  "lacima.com/Patterns/msgpack"
import "github.com/msgpack/msgpack-go"





var image_map map[string][]string
var processors map[string]bool

var processor_data_handlers map[string]*map[string]interface{}
var site_data_handlers *map[string]interface{}
var node_data_handlers *map[string]interface{}


var site_container_control_structs map[string]map[string]interface{}


func store_node_initial_status(processor string, value bool ){

   var b bytes.Buffer	
   msgpack.Pack(&b,value)
   driver := (*node_data_handlers)["NODE_STATUS"].(redis_handlers.Redis_Hash_Struct)
   driver.HSet(processor,b.String())
}



func Initialize_site_monitoring_data_structures(site_data *map[string]interface{}){

   find_container_images(site_data)
   find_data_structures(site_data)
   for processor,_ := range processors{
      store_node_initial_status(processor,true)
   }
    driver := (*node_data_handlers)["WEB_COMMAND_QUEUE"].(redis_handlers.Redis_Job_Queue)
	driver.Delete_all()
	driver1 := (*node_data_handlers)["SYSTEM_CONTAINER_IMAGES"].(redis_handlers.Redis_Hash_Struct)
	driver1.Delete_All()
	for image,_ := range image_map{
	   driver1.HSet(image,"")
	}
 
}


func find_data_structures(site_data *map[string]interface{}){

   processor_data_handlers = make( map[string]*map[string]interface{})
   var site_search_list = []string{"SITE_CONTROL"}
   site_data_handlers = data_handler.Construct_Data_Structures( &site_search_list )
   var node_search_list = []string{"NODE_MONITORING"}
   node_data_handlers = data_handler.Construct_Data_Structures( &node_search_list )
   
   for processor,_ := range processors {
     node_search_list := []string{"PROCESSOR:"+processor,"NODE_WATCH_DOG"}
	 processor_data_handlers[processor] = data_handler.Construct_Data_Structures( &node_search_list ) 
     //fmt.Println("processor",processor_data_handlers[processor])
   }
   //fmt.Println("site_data",site_data_handlers)
   //fmt.Println("processor",processor_data_handlers)
   
}



func add_image_map( node string, container_image string){

  if _,ok := image_map[container_image]; ok==false {
	image_map[container_image]= []string{node}
  }else{
    image_map[container_image] = append(image_map[container_image],node)
  }

}


func find_container_images(site_data *map[string]interface{}) {


   image_map          = make(map[string][]string)

   processors         = make(map[string]bool)

   
   
   site_search_list := []string{"SITE:"+(*site_data)["site"].(string)}
   processor_search_list := []string{"PROCESSOR"}
   
   site_nodes := graph_query.Common_qs_search(&site_search_list)
   site_node := site_nodes[0]  
   site_containers := graph_query.Convert_json_string_array(	site_node["containers"] )
   startup_containers := graph_query.Convert_json_string_array(	site_node["startup_containers"] )
   
   for _,container := range site_containers {
      container_image := find_container_image(container)
      add_image_map((*site_data)["local_node"].(string),container_image)
   }	  
   for _,container := range startup_containers {
      container_image := find_container_image(container)
      add_image_map((*site_data)["local_node"].(string),container_image)
   }		
 
   processor_nodes := graph_query.Common_qs_search(&processor_search_list)
   for _,processor_node := range processor_nodes {
      name := graph_query.Convert_json_string(processor_node["name"])
	  
      containers := graph_query.Convert_json_string_array(processor_node["containers"])
      processors[name] = true
	  for _,container := range containers {
	     container_image := find_container_image(container)
		 add_image_map(name,container_image)
	  
	  }

   }   
   //fmt.Println("image_map",image_map)
   //fmt.Println("processors",processors)
   


}




func find_container_image(container string) string{
    
  
   var search_list = []string{"CONTAINER"+":"+container}
   var container_nodes = graph_query.Common_qs_search(&search_list)
   var container_node = container_nodes[0]
   return graph_query.Convert_json_string(container_node["container_image"])

}





func Initialize_site_monitoring_chains(cf_cluster *cf.CF_CLUSTER_TYPE){

    var cf_control  cf.CF_SYSTEM_TYPE
   (cf_control).Init(cf_cluster ,"lacima_monitor_nodes",true, time.Second)
   
   (cf_control).Add_Chain("start_up_wait",true)   // watch dog strobe
   (cf_control).Cf_add_wait_interval(time.Second*14  ) // every 15 seconds
   ( cf_control).Cf_add_enable_chains_links( []string{"lacima_monitor_watch_dogs","monitor_site_command_queue"}  )
   (cf_control).Cf_add_terminate()  
   
   
   
   (cf_control).Add_Chain("lacima_monitor_watch_dogs",false)   // watch dog strobe
    var parameters = make(map[string]interface{})
   ( cf_control).Cf_add_one_step(lacima_monitor_watch_dog,parameters)
   (cf_control).Cf_add_wait_interval(time.Second*14  ) // every 15 seconds
   (cf_control).Cf_add_reset()
  
   (cf_control).Add_Chain("monitor_site_command_queue",false) // monitor command from site_contol
   
   var parameters1 = make(map[string]interface{}) 
   (cf_control).Cf_add_unfiltered_element(lacima_input_queue,parameters1)
   (cf_control).Cf_add_reset()
   
   
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func log_processor_status(processor string,status bool){

  
  var driver = (*node_data_handlers)["NODE_STATUS"].(redis_handlers.Redis_Hash_Struct)
  var value =   msgpack_utils.Unpack(driver.HGet(processor)).(bool)
  if value != status {
     var b bytes.Buffer	
     msgpack.Pack(&b,value)    
     driver.HSet(processor,b.String())
	 var log_map = make(map[string]interface{})
	 log_map["value"] = status
	 log_map["time"] =  time.Now().UnixNano()
	 var b1 bytes.Buffer	
     msgpack.Pack(&b,log_map)
     var driver1 = (*node_data_handlers)["ERROR_STREAM"].(redis_handlers.Redis_Stream_Struct)	
     driver1.Xadd(b1.String())	 
	 
  }

}


var watch_dog_value_xxx int64

func unpack_time_recovery() {
    if a := recover(); a != nil {
        fmt.Println("RECOVER", a)
    }
    watch_dog_value_xxx = 0
}

func unpack_time( driver redis_handlers.Redis_Single_Structure ) {
 
    defer unpack_time_recovery()
    watch_dog_value_xxx  =  msgpack_utils.Unpack(driver.Get()).(int64)
   
    
    
}    

func lacima_monitor_watch_dog( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{

  var processor_state bool
  time_stamp := time.Now().UnixNano()
  for processor,_ := range processors{
    driver := (*processor_data_handlers[processor])["NODE_WATCH_DOG"].(redis_handlers.Redis_Single_Structure)
    unpack_time(driver)
	if abs(time_stamp-watch_dog_value_xxx) > int64(time.Minute) {
	  processor_state = false
	}else  {
	   processor_state = true
	}
	
	log_processor_status(processor,processor_state)
    
  }

  
  return cf.CF_DISABLE
  
  
}
  
  
  
func lacima_input_queue( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{

  var driver = (*site_data_handlers)["WEB_COMMAND_QUEUE"].(redis_handlers.Redis_Job_Queue)
  if driver.Length() != 0 {
     var data = msgpack_utils.Unpack( driver.Show_next_job( ) ).(cf.CF_EVENT_TYPE)
	 if data.Name == "reboot" {
	    var b bytes.Buffer	
        msgpack.Pack(&b,data) 
	    for processor,_:= range processors {
		    var node_driver = (*processor_data_handlers[processor])["NODE_COMMAND_QUEUE"].(redis_handlers.Redis_Job_Queue)
   
		    node_driver.Push(b.String())
		}
	 
	 }
	 /*  handled by an ansilbe like facility
	 if  data.Name == "upgrade" {
	    var image_list = data.Value.([]string)
	    for _,image := range image_list {
             var impacted_processors = image_map[image]
			 for _, processor := range impacted_processors {
			    var driver = (*processor_data_handlers[processor])["NODE_UPGRADE_QUEUE"].(redis_handlers.Redis_Job_Queue)
				var b bytes.Buffer	
                msgpack.Pack(&b,image)
			    driver.Push(b.String())
			 }
		
		
		}
		
		
	}
	*/
  
  
  }
  
  
  return cf.CF_HALT
  
}
