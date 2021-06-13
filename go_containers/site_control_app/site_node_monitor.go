package main

import "os"
import "fmt"
import "time"
import "strconv"
import "context"
import "net"
import "strings"

import "lacima.com/site_data"

import "lacima.com/site_control_app/site_init"
import "lacima.com/site_control_app/node_init"
import "lacima.com/site_control_app/site_control"
import "lacima.com/site_control_app/node_control"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"
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

func mount_usb_drive(){
  defer handle_mount_panic()
  fmt.Println(docker_control.System_shell("mount /dev/sda /home/pi/mountpoint"))

}



var site_data map[string]interface{}

func fill_in_site_data(){
 
   
  site_data= make(map[string]interface{})
  
  site_data["master_flag"]  = os.Getenv("master_flag")
  site_data["site"]  = os.Getenv("site")
  site_data["local_node"]  = os.Getenv("local_node")
  port,_ := strconv.Atoi(os.Getenv("port"))
  site_data["port"]  = float64(port)
  
  graph_db, _  := strconv.Atoi(os.Getenv("graph_db"))
  site_data["graph_db"]  = float64(graph_db)
  
  
  
  redis_table, _  := strconv.Atoi(os.Getenv("redis_table"))
  site_data["redis_table"] = float64(redis_table)
  
  password_table, _  := strconv.Atoi(os.Getenv("password_table"))
  site_data["password_table"] = float64(password_table)
  
  irrigation_files, _  := strconv.Atoi(os.Getenv("irrigation_files"))
  site_data["irrigation_files"] = float64(irrigation_files)
  
  
  // necessary for a new installation or corrupted installation
  site_data["graph_container_image"]   = os.Getenv("graph_container_image")
  site_data["graph_container_script"]    = os.Getenv("graph_container_script")		
  site_data["redis_start_script"]              = os.Getenv("redis_start_script")
  site_data["host"]                                    =   os.Getenv("host")
  site_data["config_file"]              =  os.Getenv("config_file")
  
  fmt.Println("config_file",site_data["config_file"])
  /*
   *   store site file
   *
   */

  get_site_data.Save_site_data(site_data["config_file"].(string)  ,site_data)
  mount_usb_drive()
  
  
}





func main(){
    
    fill_in_site_data()
  
 
	var master_flag = site_data["master_flag"].(string)
	fmt.Println("master flag",master_flag)
    redis_handlers.Init_Redis_Mutex()

	if master_flag == "true"{
       
	   site_init.Site_Init(&site_data)
       //data_handler.Data_handler_init(&site_data)
    
       
       
       
       ip_table := data_handler.Construct_Data_Structures(&[]string{"NODE_MAP"})
       ip_driver := (*ip_table)["NODE_MAP"].(redis_handlers.Redis_Hash_Struct)
       ip_address := find_local_address()
       ip_driver.HSet("SITE",ip_address )
       
       
	} else {
       
	   wait_for_redis_connection(site_data["host"].(string), int(site_data["port"].(float64)) )
       graph_query.Graph_support_init(&site_data)
       data_handler.Data_handler_init(&site_data)
	}
	
	
	
 	
    ip_table := data_handler.Construct_Data_Structures(&[]string{"NODE_MAP"})
    ip_driver := (*ip_table)["NODE_MAP"].(redis_handlers.Redis_Hash_Struct)
    ip_address := find_local_address()
    ip_driver.HSet(site_data["local_node"].(string),ip_address )    
	node_init.Node_Init(&site_data)
   
	
	
    
	(CF_site_node_control_cluster).Cf_cluster_init()
	(CF_site_node_control_cluster).Cf_set_current_row("site_node_control")
	site_control.Site_Startup(&CF_site_node_control_cluster,&site_data)
	node_control.Node_Startup(&CF_site_node_control_cluster,&site_data)
	/*
	
	   --- other initializations
	   
	   
	*/
	
	(CF_site_node_control_cluster).CF_Fork()
	
	
	
   var loop_flag = true
   for loop_flag{
      time.Sleep(time.Second*100)
      fmt.Println("main is spinning")
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

func find_local_address()string{
    
   conn, error := net.Dial("udp", "8.8.8.8:80")  
   if error != nil {  
      fmt.Println(error)  
  
    }  
  
    defer conn.Close()  
    ipAddress_port := conn.LocalAddr().(*net.UDPAddr).String()
    temp := strings.Split(ipAddress_port,":")
    ip_address := temp[0]
  
    return ip_address
}      
    
    
    


