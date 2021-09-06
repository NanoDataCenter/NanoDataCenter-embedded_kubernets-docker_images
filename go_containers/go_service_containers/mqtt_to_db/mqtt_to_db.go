package main

import "time"

import "fmt"
import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/go_service_containers/mqtt_to_db/mqtt_client"
import "lacima.com/go_service_containers/mqtt_to_db/mqtt_web"
import "lacima.com/go_service_containers/mqtt_to_db/mqtt_db_table_trim"
import "lacima.com/go_service_containers/mqtt_to_db/mqtt_device_monitor"
import "lacima.com/go_service_containers/mqtt_to_db/mqtt_on_line_test"
import mqtt "github.com/eclipse/paho.mqtt.golang"


var site_data_store map[string]interface{}
const config_file = "/data/redis_configuration.json"

var client mqtt.Client


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
    mqtt_client.Construct_event_registry_actions(site)
    client = mqtt_client.Construct_mqtt_actions(ip,port)
    mqtt_web.Construct_event_registry_actions(site)
    mqtt_db_trim.Trim_int(3600*24) // one day trim time
    mqtt_monitor_devices.Monitor_int()
    mqtt_test.Mqtt_test_init(site)
}



func mqtt_monitor_exec(){


  go mqtt_db_trim.Trim_dbs()

  go mqtt_monitor_devices.Monitor_devices()
  go mqtt_test.Mqtt_test_tx(client)
  go mqtt_test.Mqtt_test_rx()

  go mqtt_web.Init_site_web_server()
  for true {
       time.Sleep(time.Second*10)
       fmt.Println("polling loop")
    }
        
}
