package main

import "time"

import "fmt"
import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/go_service_containers/error_detection/web_services/web_components"


var site_data_store map[string]interface{}
const config_file = "/data/redis_configuration.json"




func main(){

   
 
    site_data_store = get_site_data.Get_site_data(config_file)
    graph_query.Graph_support_init(&site_data_store)
    redis_handlers.Init_Redis_Mutex()
	data_handler.Data_handler_init(&site_data_store)
    error_detection_components.Init_site_web_server()
    
    
	
	
    for true {
        
      time.Sleep(time.Minute*10)
      fmt.Println("polling loop")   
        
    }


}

