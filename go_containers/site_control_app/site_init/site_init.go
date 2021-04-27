
package site_init

import ( 
"fmt"
"context"
"time"
"strconv"



"site_control.com/smtp"

"site_control.com/docker_control"
"site_control.com/redis_support/graph_query"
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
	       fmt.Println("should not happen")
	       //panic(value["container_image"])
	       docker_control.Pull(value["container_image"])
	      }
	      docker_control.Container_rm(value["name"])
	     docker_control.Container_up(value["name"],value["startup_command"])
	}
  }
}


func start_system_containers(){
  for _,value := range site_containers{
    if value["name"] == "redis" {
	  fmt.Println("found redis")
	  continue
	}
	if docker_control.Image_Exists(value["container_image"]) == false{
	   fmt.Println("should not happen")
	   //panic(value["container_image"])
	   docker_control.Pull(value["container_image"])
	}
	docker_control.Container_rm(value["name"])
	docker_control.Container_up(value["name"],value["startup_command"])
  }
}  
	
   

func find_site_containers(){
    
    for _,service := range containers{
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
	//fmt.Println(site_containers)
}

func  determine_system_containers(){
    var search_list = []string{ "SITE_CONTROL:SITE_CONTROL" }
    var site_nodes = graph_query.Common_qs_search(&search_list)
    var site_node = site_nodes[0]
   

    containers               = graph_query.Convert_json_string_array(	site_node["containers"] )


}


func wait_for_redis_connection(address string, port int ) {
   var address_port = address+":"+strconv.Itoa(port)
   //fmt.Println("address",address_port)
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
     
}




func  startup_redis_container(redis_startup_script string){
      fmt.Println("start redis container")
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
						 
    var redis_startup_script = "docker run -d  --network host   --name redis    --mount type=bind,source=/home/pi/mountpoint/lacuma_conf/redis,target=/data    nanodatacenter/redis /bin/bash /pod_util/redis_control.bsh"
	
    var password_script ="python3 /home/pi/passwords.py"
    //var redis_image = "nanodatacenter/redis" 


   var hot_start = determine_hot_start()
   
   if hot_start == false {
  
      stop_running_containers()
      remove_redis_container()
      startup_redis_container(redis_startup_script)
      wait_for_redis_connection((*site_data)["host"].(string), int((*site_data)["port"].(float64)) )
      fmt.Println("redis is up")
  
   
      if docker_control.Image_Exists(graph_container_image)== false {
          docker_control.Pull(graph_container_image)
	  }
      docker_control.Container_Run(graph_container_script)
	  
      fmt.Println(docker_control.System_shell(password_script))
   
      graph_query.Graph_support_init(site_data)
      determine_system_containers()
      find_site_containers()
      start_system_containers()
      docker_control.Prune()
     smtp.Send_Mail("site is intialized")
   }else {
         graph_query.Graph_support_init(site_data)  // only start containers that are not running
		 find_site_containers()
		 start_stopped_system_containers()
		 docker_control.Prune()
		 
   }
   time.Sleep(time.Second*5) // allow containers to startup_command
   verify_system_containers()

}	
						 





    



