package main


import "time"
import "io/ioutil"
import "io/fs"

import "os"
import "lacima.com/site_data"
import "lacima.com/redis_support/graph_generation"



func main(){

    
    
	var config_file = "/data/redis_server.json"
	
	var site_data_store map[string]interface{}

	site_data_store = get_site_data.Get_site_data(config_file)
    graph_query.Graph_support_init(&site_data_store)
	redis_handlers.Init_Redis_Mutex()
	data_handler.Data_handler_init(&site_data_store)	
    graph_generation.Graph_support_init(&site_data_store)
	
 
   
}


