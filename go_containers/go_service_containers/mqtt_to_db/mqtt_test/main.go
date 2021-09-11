package main

import "time"

import "fmt"
import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"

import "lacima.com/go_service_containers/mqtt_to_db/mqtt_test/mqtt_out_client"
import "lacima.com/go_service_containers/mqtt_to_db/mqtt_test/mqtt_out_web"


import mqtt "github.com/eclipse/paho.mqtt.golang"


var site_data_store map[string]interface{}
const config_file = "/data/redis_configuration.json"

var client mqtt.Client


func main(){

   
 
    site_data_store = get_site_data.Get_site_data(config_file)
    graph_query.Graph_support_init(&site_data_store)
	data_handler.Data_handler_init(&site_data_store)
    redis_handlers.Init_Redis_Mutex()
	mqtt_monitor_init()
	mqtt_monitor_exec()


}

func mqtt_monitor_init(){
    ip   := site_data_store["host"].(string)
    port := 1883
    site := site_data_store["site"].(string)
    mqtt_out_client.Construct_event_registry_actions(site)
    mqtt_out_client.Test_generator_init()
    mqtt_out_client.Construct_mqtt_actions( ip, port )
    mqtt_out_client.Wait_for_connections()
    mqtt_test_web.Init_site_web_server()
    
}



func mqtt_monitor_exec(){

  mqtt_out_client.Test_generator_start()

  for true {
       time.Sleep(time.Second*10)
       fmt.Println("polling loop")
    }
        
}
