package mqtt_out_client


import "fmt"
import  "time"
import "lacima.com/server_libraries/mqtt_publish_rpc"

var mqtt_library mqtt_publish_server_lib.MQTT_Publish_Client_Type

func Test_generator_init(){
    search_list := []string{"MQTT_OUTPUT_SETUP:site_out_server","RPC_SERVER:MQTT_OUT_RPC_SERVER","RPC_SERVER"}
    mqtt_library = mqtt_publish_server_lib.MQTT_Publish_Init(&search_list)
    
    

    
}


func Test_generator_start(){
   
    topic := "mqtt_output/test_message/heart_beat"
    payload := "test message"
    for true {
      fmt.Println("sending data")
      mqtt_library.Publish(topic,payload)     
      time.Sleep(time.Minute*15)
      
    }
    
    
}    
