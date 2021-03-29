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
    var master_flag = false	
	if master_flag == true {
	   site_control.Site_Initialization(&site_data_store)
	}else{
	  graph_query.Graph_support_init(&site_data_store)
	}
	// do node_init
	
	
	redis_handlers.Init_Redis_Mutex()
	data_handler.Data_handler_init(&site_data_store)
	
    
	
	site_control.Site_control_startup(&site_data_store)
    site_control.Initialize_CF()	

	
	
	
   var loop_flag = true
   for loop_flag{
      time.Sleep(time.Second*10)
      fmt.Println("main is spinning")
   } 

	
	
    

}