
package site_init

import ( 
"fmt"
"context"
"time"
"strconv"

"lacima.com/redis_support/generate_handlers"
"lacima.com/redis_support/redis_handlers"
"lacima.com/site_control_app/docker_control"
"lacima.com/redis_support/graph_query"
"lacima.com/Patterns/logging_support"
"github.com/go-redis/redis/v8"

)

type site_data_type map[string]interface{}
var site_data map[string]interface{}
//type Database struct {
//   Client *redis.Client
//}
var ctx    = context.TODO()
var graph_container_image string
var graph_container_script string
var services_json string
//var container []string
var containers = make([]string,0)
var startup_containers = make([]string,0)
var site_run_once_containers = make([]map[string]string,0)
var site_containers = make([]map[string]string,0)


func verify_system_containers(){

 for _,value := range site_containers{
   if docker_control.Container_is_running(value["name"] )== false{
     panic("container "+value["name"]+"is not running")
   }
 }
}

func start_stopped_system_containers(){
  for _,value := range site_containers{
    if value["name"] == "redis" {
	  fmt.Println("found redis")
	  continue
	}
	
	 if docker_control.Container_is_running(value["name"]) == false{
         
	     if docker_control.Image_Exists(value["container_image"]) == false{
	       
	       docker_control.Pull(value["container_image"])
	      }
	      docker_control.Container_rm(value["name"])
	     docker_control.Container_up(value["name"],value["startup_command"])
	}
  }
}

func start_run_once_containers(){
  //fmt.Println("run once ",site_run_once_containers)
  for _,value := range site_run_once_containers{
    if value["name"] == "redis" {
	  fmt.Println("found redis")
	  continue
	}

	if docker_control.Image_Exists(value["container_image"]) == false{
	   
	   
	   docker_control.Pull(value["container_image"])
	}
	
	docker_control.Container_rm(value["name"])
    
	docker_control.Container_up(value["name"],value["startup_command"])
  }
}

func start_system_containers(){
  for _,value := range site_containers{
    if value["name"] == "redis" {
	  fmt.Println("found redis")
	  continue
	}
	if docker_control.Image_Exists(value["container_image"]) == false{
	   
	   docker_control.Pull(value["container_image"])
	}
	
	docker_control.Container_rm(value["name"])
	docker_control.Container_up(value["name"],value["startup_command"])
  }
}  
	
 
func find_startup_conatiners(){
    
    for _,service := range startup_containers{
         //fmt.Println("service",service)
	     var item = make(map[string]string,0)
	     var search_list = []string{"CONTAINER"+":"+service}
		 var container_nodes = graph_query.Common_qs_search(&search_list)
         var container_node = container_nodes[0]
		 //fmt.Println(container_node)
		 item["name"] =  graph_query.Convert_json_string(container_node["name"])
 		 item["container_image"] =graph_query.Convert_json_string(container_node["container_image"])
		 item["startup_command"] = graph_query.Convert_json_string(container_node["startup_command"])

		 site_run_once_containers = append(site_run_once_containers,item)
		 
	}
	
	//fmt.Println("run_once", site_run_once_containers)
} 
 

func find_site_containers(){
    
    for _,service := range containers{
         //fmt.Println("service",service)
	     var item = make(map[string]string,0)
	     var search_list = []string{"CONTAINER"+":"+service}
		 var container_nodes = graph_query.Common_qs_search(&search_list)
         var container_node = container_nodes[0]
		 //fmt.Println(container_node)
		 item["name"] =  graph_query.Convert_json_string(container_node["name"])
 		 item["container_image"] =graph_query.Convert_json_string(container_node["container_image"])
		 item["startup_command"] = graph_query.Convert_json_string(container_node["startup_command"])

		 site_containers = append(site_containers,item)
		 
	}
	
	//fmt.Println("site", site_containers)
}

func  determine_system_containers(site string){
    search_list := []string{ "SITE:"+site }
    //mt.Println(search_list)
    site_nodes := graph_query.Common_qs_search(&search_list)
    site_node  := site_nodes[0]
    //fmt.Println("site_node",site_node)
    
    containers               = graph_query.Convert_json_string_array(	site_node["containers"] )
    startup_containers       = graph_query.Convert_json_string_array(	site_node["startup_containers"] )
    //fmt.Println("startup_containers",startup_containers)
}


func wait_for_redis_connection(address string, port int ) {
   var address_port = address+":"+strconv.Itoa(port)
   fmt.Println("address",address_port)
   fmt.Println("wait_for_redis_connection",port)
   var loop_flag = true
   for loop_flag == true {
       client := redis.NewClient(&redis.Options{
                                                 Addr: address_port,
												
												 DB: 0,
                                               })
		err := client.Ping(ctx).Err();
		if err != nil{
		  fmt.Println("redis connection is not up")
		  time.Sleep(time.Second)
		}else {
		  loop_flag = false
		}  
      		
		client.Close() 
   }		
   fmt.Println("redis up")  
}




func  startup_redis_container(redis_startup_script string){
      //mt.Println("start redis container")
	  docker_control.Container_up("redis",redis_startup_script)
	  time.Sleep(time.Second*4)
	  if docker_control.Container_is_running("redis") == false{
	     panic("redis container did not start")
	  }
	 
}	 


func remove_redis_container(){
    fmt.Println("remove redis container")
	docker_control.Container_rm("redis")
}


func stop_running_containers() {
   fmt.Println("stop redis container")
   var running_containers = docker_control.Containers_ls_runing()
   for _,name := range running_containers{
      docker_control.Container_stop(name)
   }      
}

func determine_hot_start() bool {

  var running_containers = docker_control.Containers_ls_runing()
  for _,name := range running_containers{
    if name == "redis"{
	  return true
	}
  }

  
  return false
  

}
	   
func Site_Init(  site_data *map[string]interface{} ){ 
                         
	graph_container_image = (*site_data)["graph_container_image"].(string)
    graph_container_script = (*site_data)["graph_container_script"].(string)			 
						 
    var redis_startup_script = (*site_data)["redis_start_script"].(string)		
	
    //fmt.Println((*site_data)["graph_container_image"].(string))
    //fmt.Println((*site_data)["graph_container_script"].(string))
    //fmt.Println((*site_data)["redis_start_script"].(string))
    

     hot_start := determine_hot_start()
     fmt.Println("hot start",hot_start)
   
   if hot_start == false {
      
      stop_running_containers()
      remove_redis_container()
      
      startup_redis_container(redis_startup_script)
      
      wait_for_redis_connection((*site_data)["host"].(string), int((*site_data)["port"].(float64)))
      
   
      if docker_control.Image_Exists(graph_container_image)== false {
          
          docker_control.Pull(graph_container_image)
	  }
	  
      docker_control.Container_Run(graph_container_script)
	  
      
   
      graph_query.Graph_support_init(site_data)
      data_handler.Data_handler_init(site_data)
      
     
      reboot_flag := data_handler.Construct_Data_Structures(&[]string{"REBOOT_FLAG"})
      reboot_flag_driver := (*reboot_flag)["REBOOT_FLAG"].(redis_handlers.Redis_Single_Structure)
      reboot_flag_driver.Set("ACTIVE")
      
      incident_log := logging_support.Construct_incident_log( []string{"INCIDENT_LOG:SITE_REBOOT","INCIDENT_LOG"} ) 
      incident_log.Post_event(true,"reboot","reboot")

      
      
      determine_system_containers((*site_data)["site"].(string))
      find_site_containers()
      find_startup_conatiners()
      start_run_once_containers()

      start_system_containers()
      reboot_flag_driver.Set("NOT_ACTIVE")
      docker_control.Prune()
      
   }else {
         graph_query.Graph_support_init(site_data)  // only start containers that are not running
         data_handler.Data_handler_init(site_data)
         
         
         determine_system_containers((*site_data)["site"].(string))
		 find_site_containers()
         //fmt.Println("containers",site_containers)
         
		 start_stopped_system_containers()
         
         reboot_flag := data_handler.Construct_Data_Structures(&[]string{"REBOOT_FLAG"})
         reboot_flag_driver := (*reboot_flag)["REBOOT_FLAG"].(redis_handlers.Redis_Single_Structure)
         reboot_flag_driver.Set("NOT_ACTIVE")
		 docker_control.Prune()
        
		 
   }
 
   fmt.Println("allowing system containers to get started")
   time.Sleep(time.Second*5) // allow containers to startup_command
   verify_system_containers()

}	
						 





    




