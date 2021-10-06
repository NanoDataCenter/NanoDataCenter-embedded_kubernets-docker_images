package main

import "fmt"


import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/server_libraries/telegram_rpc"




func main(){

    
    
	var config_file ="/data/redis_configuration.json"
	
	

    
	site_data_store := get_site_data.Get_site_data(config_file)
    fmt.Println(site_data_store)
    graph_query.Graph_support_init(&site_data_store)
	redis_handlers.Init_Redis_Mutex()
    
	data_handler.Data_handler_init(&site_data_store)	
 	
    rpc_driver :=telegram_rpc_interface.Site_Server_Init()
     fmt.Println(rpc_driver.Ping())
    fmt.Println(rpc_driver.Send_message("telegram test message"))
    
    
   
}




