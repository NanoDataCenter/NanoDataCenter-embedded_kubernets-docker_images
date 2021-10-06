package main

import "fmt"
import "time"

import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"

import "lacima.com/Patterns/secrets"
import "lacima.com/go_system_containers/alert_handlers/telegram/api"



func main(){

    
    
	var config_file ="/data/redis_configuration.json"
	
	

    
	site_data_store := get_site_data.Get_site_data(config_file)
    fmt.Println(site_data_store)
    graph_query.Graph_support_init(&site_data_store)
	redis_handlers.Init_Redis_Mutex()
    
	data_handler.Data_handler_init(&site_data_store)	
 	
    
    initialize_telegraph_server(site_data_store)
    
    
     search_list := []string{"RPC_SERVER:SITE_FILE_SERVER","RPC_SERVER"}

     handlers := data_handler.Construct_Data_Structures(&search_list)
    
     driver := (*handlers)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
	
	
	
	driver.Add_handler( "send_message",send_message)
	
	driver.Json_Rpc_start()
	
	for true {
	  //fmt.Println("main spining")
	  time.Sleep(time.Second*10)
	}
   
}

func initialize_telegraph_server(site_data_store map[string]interface{}){

    properties          := graph_query.Common_qs_search(&[]string{"TELEGRAM_SERVER:TELEGRAM_SERVER"})
    fmt.Println("properties",properties)
    property            := properties[0]
    valid_contact_list  := graph_query.Convert_json_string_array(property["valid_users"])
 
    secrets.Init_file_handler(site_data_store)
    telebot_token := secrets.Get_Secret("telegram","token")
    telegram.Init(telebot_token,valid_contact_list)
}



	
func send_message( parameters map[string]interface{} ) map[string]interface{}{

  
  sent_message     := parameters["message"].(string)
  
  err              :=  telegram.Send_message(sent_message)
  
  if err == true{
        parameters["status"]  = true
		parameters["results"] = ""
  } else {
        parameters["status"]  = false
		parameters["results"] = ""
  }
  
  return parameters

}


