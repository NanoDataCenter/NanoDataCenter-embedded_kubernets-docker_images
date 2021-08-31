package main

import "time"

import "fmt"
import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/go_service_containers/mqtt_to_db/mqtt_support"
import "lacima.com/go_service_containers/mqtt_to_db/mqtt_web"
var site_data_store map[string]interface{}
const config_file = "/data/redis_configuration.json"




func main(){

   
 
    site_data_store = get_site_data.Get_site_data(config_file)
    graph_query.Graph_support_init(&site_data_store)
	redis_handlers.Init_Redis_Mutex()
	data_handler.Data_handler_init(&site_data_store)

	mqtt_monitor_init()
	mqtt_monitor_exec()


}

func mqtt_monitor_init(){
    ip   := site_data_store["host"].(string)
    port := 1883
    site := site_data_store["site"].(string)
    mqtt_support.Construct_event_registry_actions(site)
    mqtt_support.Construct_mqtt_actions(ip,port)
    mqtt_web.Construct_event_registry_actions(site)
}



func mqtt_monitor_exec(){
  go mqtt_support.Monitor_devices()
  go mqtt_web.Init_site_web_server() 
  for true {
       time.Sleep(time.Second*10)
       fmt.Println("polling loop")
    }
        
}
