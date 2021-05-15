package main

//import "fmt"
import "os"
import "time"
import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/cf_control"

import "lacima.com/go_basic_container/process_control_support"

var  cf_control_cluster cf.CF_CLUSTER_TYPE
var site_data_store map[string]interface{}
const config_file = "/data/redis_configuration.json"

func main(){

    container_name := os.Getenv("CONTAINER_NAME")
   
    site_data_store = get_site_data.Get_site_data(config_file)
	
    graph_query.Graph_support_init(&site_data_store)
	redis_handlers.Init_Redis_Mutex()
	data_handler.Data_handler_init(&site_data_store)
	




    system_ctrl := system_control.Construct_System_Control(   container_name )

	cf_control_cluster.Cf_cluster_init()
	cf_control_cluster.Cf_set_current_row("container_process_monitor")
	
    (system_ctrl).Init(&cf_control_cluster)
	
	(cf_control_cluster).CF_Fork()	
	
    loop_flag := true
    for loop_flag{
      time.Sleep(time.Second*100)
      //fmt.Println("main is spinning")
    } 

	
	
    

}





