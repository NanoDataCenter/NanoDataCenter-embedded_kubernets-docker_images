package main


import "fmt"
import "time"
import "strconv"
import "context"
import "site_control.com/site_data"
import "site_control.com/smtp"
import "site_control.com/site_init"
import "site_control.com/node_init"
import "site_control.com/site_control"
import "site_control.com/node_control"
import "site_control.com/redis_support/graph_query"
import "site_control.com/redis_support/redis_handlers"
import "site_control.com/redis_support/generate_handlers"
import "site_control.com/cf_control"
import "site_control.com/docker_control"
import "github.com/go-redis/redis/v8"

var  CF_site_node_control_cluster cf.CF_CLUSTER_TYPE



func handle_mount_panic() {
  
    if a := recover(); a != nil {
        fmt.Println("RECOVER", a)
    }
}

func mount_usb_drive(){
  defer handle_mount_panic()
  fmt.Println(docker_control.System_shell("mount /dev/sda1 /home/pi/mountpoint"))

}


func main(){

    //mount_usb_drive()
      
		
	var config_file = "/home/pi/mountpoint/lacuma_conf/site_config/redis_server.json"
	var site_data_store map[string]interface{}

	site_data_store = get_site_data.Get_site_data(config_file)
 
	var master_flag = site_data_store["master"].(bool)
	fmt.Println("master flag",master_flag)
	if master_flag {
	    smtp.Initialization( &site_data_store  , "System Control Startup")
    }else { 
        smtp.Initialization( &site_data_store, "Node Control Startup" )
	}
	if master_flag == true{
	   site_init.Site_Init(&site_data_store)
	} else {
	   wait_for_redis_connection(site_data_store["host"].(string), int(site_data_store["port"].(float64)) )
       graph_query.Graph_support_init(&site_data_store)
	}
 	  
 
	node_init.Node_Init(&site_data_store)

	//graph_query.Graph_support_init(&site_data_store)
	redis_handlers.Init_Redis_Mutex()
	data_handler.Data_handler_init(&site_data_store)
	
    
	(CF_site_node_control_cluster).Cf_cluster_init()
	(CF_site_node_control_cluster).Cf_set_current_row("site_node_control")
	site_control.Site_Startup(&CF_site_node_control_cluster,&site_data_store)
	node_control.Node_Startup(&CF_site_node_control_cluster,&site_data_store)
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

