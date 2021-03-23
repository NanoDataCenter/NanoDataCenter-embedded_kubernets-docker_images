package main
import "fmt"
import "time"
import "site_control.com/site_data"
import "site_control.com/docker_management"
import "site_control.com/redis_support/graph_query"
import  "site_control.com/redis_support/generate_handlers"
var site_data map[string]interface{}


func cannot_continue(display_string string){

   
   var delay_count = time.Second*10
   for{
     fmt.Println(display_string)
	 time.Sleep(delay_count)
   }
   
}


func determine_master(site_file string) map[string]interface{} {

       var site_data = get_site_data.Get_site_data(site_file)
	   var val,ok = site_data["master"]
	   if ok&&val != true {
	       cannot_continue("Not Master -- Spining in loop")
	   }
	   
       return site_data  
       
}



func Site_Control( config_file string ) {
   site_data = determine_master(config_file)
   graph_query.Graph_support_init(&site_data)
   data_handler.Data_handler_init(&site_data)
   //fmt.Println(site_data)
   // cor_int Monitor Nodes
   // cor_int Monitor System Containers
   // cor_int Queue for Node and System Control
   // cor_int Queue for container update
   var search_path = []string{"SITE_CONTROL:SITE_CONTROL"}
   docker_management.Initialize_Docker_Monitor(&search_path)
   
   
   go docker_management.Docker_Monitor()
   go docker_management.Docker_Performance_Monitor()
   
   var loop_flag = true
   for loop_flag{
      time.Sleep(time.Second*10)
      fmt.Println("waiting for coroutinnse")
   } 



}



func main() {

	var config_file = "/mnt/ssd/site_config/redis_server.json"
	fmt.Println("main start")
    Site_Control(config_file  )
}

