package main

import "fmt"
import "time"
import "site_control.com/site_data"
//import "site_control.com/smtp"
import "site_control.com/site_control"
import "site_control.com/redis_support/graph_query"
import "site_control.com/redis_support/redis_handlers"
import "site_control.com/redis_support/generate_handlers"

func main(){

	var config_file = "/mnt/ssd/site_config/redis_server.json"
	var site_data_store map[string]interface{}

	site_data_store = get_site_data.Get_site_data(config_file)
/*
	var master_flag = site_data_store["master"].(bool)
	if master_flag {
	    smtp.Initialization( &site_data_store  , "System Control Startup")
    }else { 
        smtp.Initialization( &site_data_store, "Node Control Startup" )
	}
		
*/
 	  

	// do node_init
	
	graph_query.Graph_support_init(&site_data_store)
	redis_handlers.Init_Redis_Mutex()
	data_handler.Data_handler_init(&site_data_store)
	
    
	
	site_control.Site_Startup(&site_data_store)
	/*
	
	   --- other initializations
	   
	   
	*/
	
	go site_control.Execute()
   

	
	
	
   var loop_flag = true
   for loop_flag{
      time.Sleep(time.Second*100)
      fmt.Println("main is spinning")
   } 

	
	
    

}