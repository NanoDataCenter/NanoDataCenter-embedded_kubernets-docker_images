package main

import "os"
import "fmt"
import "time"
import "strconv"
import "context"


//import "lacima.com/site_data"
import  "lacima.com/site_control_app/docker_management"
import "lacima.com/site_control_app/site_init"
import "lacima.com/site_control_app/node_init"
import "lacima.com/site_control_app/site_control"
import "lacima.com/site_control_app/node_control"
//import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"

import "lacima.com/cf_control"
import "lacima.com/site_control_app/docker_control"
import "github.com/go-redis/redis/v8"

const config_file string = "/home/pi/system_config/redis_configuration.json"
var  CF_site_node_control_cluster cf.CF_CLUSTER_TYPE



func handle_mount_panic() {
  
    if a := recover(); a != nil {
        fmt.Println("RECOVER", a)
    }
}

func mount_usb_drive(mount_string string){
  defer handle_mount_panic()
  fmt.Println(docker_control.System_shell("mount /dev/sda /home/pi/mountpoint"))

}



var site_data map[string]interface{}

func fill_in_site_data(){
  site_data = make(map[string]interface{})
  site_data["master_flag"]  = os.Getenv("master_flag")
  site_data["site"]  = os.Getenv("site")
  site_data["local_node"]  = os.Getenv("local_node")
  
  // ip of the redis server
   port,_               := strconv.Atoi(os.Getenv("port"))
  site_data["port"]     = float64(port)
  site_data["host"]     =   os.Getenv("host")
  graph_db,_              := strconv.ParseFloat(os.Getenv("graph_db"),64)
  site_data["graph_db"] = graph_db
  
  
  site_data["graph_container_image"]     = os.Getenv("graph_container_image")
  site_data["graph_container_script"]    = os.Getenv("graph_container_script")		
  site_data["redis_container_name"]      = os.Getenv("redis_container_name")
  site_data["redis_container_image"]     = os.Getenv("redis_container_image")
  site_data["redis_start_script"]        = os.Getenv("redis_start_script")
  
  
 
  
  
  
}

func fill_in_slave_data(){

  /*
   * Minimium information to connect to event broker and event registry
   */
  
  site_data = make(map[string]interface{})
  site_data["master_flag"]  = os.Getenv("master_flag")
  site_data["site"]  = os.Getenv("site")
  site_data["local_node"]  = os.Getenv("local_node")
  
  // ip of the redis server
   port,_               := strconv.Atoi(os.Getenv("port"))
  site_data["port"]     = float64(port)
  site_data["host"]     =   os.Getenv("host")
  graph_db,_              := strconv.Atoi(os.Getenv("graph_db"))
  site_data["graph_db"] = graph_db
  
}


func main(){
    
    
  
    var mount_string = os.Getenv("mount_string")
	var master_flag = os.Getenv("master_flag")
	fmt.Println("master flag",master_flag)
    redis_handlers.Init_Redis_Mutex()

	if master_flag == "true"{
       
       mount_usb_drive(mount_string) // mount external hard drive for storing system data
       fill_in_site_data()
	   site_init.Site_Master_Init(&site_data)
       
	} else {
      fill_in_slave_data()
      site_init.Site_Slave_Init(&site_data)
       

	}
	
	
	
 	

	node_init.Node_Init(&site_data)
    
	
	
    
	(CF_site_node_control_cluster).Cf_cluster_init()
	(CF_site_node_control_cluster).Cf_set_current_row("site_node_control")
    
    var all_containers = make([]string,0)
    if master_flag == "true" {
	    all_containers = docker_management.Find_containers(&[]string{ "SITE:"+site_data["site"].(string) })
    }
    all_containers = append( all_containers, docker_management.Find_containers( &[]string{"NODE:"+site_data["local_node"].(string)} )...)
	
    
    node_control.Node_Startup(&CF_site_node_control_cluster,&site_data,all_containers)
    site_control.Site_Startup(&CF_site_node_control_cluster,&site_data)
	/*
	
	   --- other initializations
	   
	   
	*/
	
	(CF_site_node_control_cluster).CF_Fork()
	
	
	
   var loop_flag = true
   for loop_flag{
      time.Sleep(time.Second*100)
      //fmt.Println("main is spinning")
   } 

	
	
    

}



func wait_for_redis_connection(address string, port int ) {
   var ctx    = context.TODO()
 
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

 
    
    
    


