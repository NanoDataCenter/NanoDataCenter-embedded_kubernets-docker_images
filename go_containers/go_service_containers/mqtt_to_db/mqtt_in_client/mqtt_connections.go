package mqtt_client

import "fmt"
import "os"
import "time"
import "lacima.com/redis_support/graph_query"
import mqtt "github.com/eclipse/paho.mqtt.golang"

var client mqtt.Client
var connection_status bool

func Construct_mqtt_actions( ip string, port int)mqtt.Client{  // setup receiving mqtt construct_mqtt_tasks

    data_nodes := graph_query.Common_qs_search(&[]string{"MQTT_CLIENT_ID:MQTT_CLIENT_ID"})
    data_node  := data_nodes[0]
    
    mqtt_client_map := graph_query.Convert_json_dict( data_node["mqtt_client_id_map"] )
    
    if _,ok := mqtt_client_map["mqtt_input_server" ]; ok == false {
        panic("no existant mqtt client map")
    }
    
    mqtt_client_id := mqtt_client_map["mqtt_input_server"]
    
    
    connection_status = false    
    opts := mqtt.NewClientOptions()
    opts.AddBroker(fmt.Sprintf("tcp://%s:%d", ip, port))
    opts.SetClientID( mqtt_client_id )
    opts.SetAutoReconnect(true)
    opts.SetCleanSession(true)
    opts.SetConnectRetry(true)
    opts.SetConnectRetryInterval(time.Second *5 )
    opts.SetConnectTimeout(time.Second*30)
    opts.SetKeepAlive(time.Second*5)
    opts.SetMaxReconnectInterval(time.Second*5)
        
        
    opts.SetDefaultPublishHandler(messagePubHandler)
    opts.OnConnect = connectHandler
    opts.OnConnectionLost = connectLostHandler 
    client := mqtt.NewClient(opts)
    
    token := client.Connect()
    if token.WaitTimeout(time.Second*30) == false {
        os.Exit(1)
    }
    if token.Error() != nil {
        os.Exit(1)
    }
    
    topic := get_monitoring_topic()
    client.Subscribe("$SYS/#", 1, messagePubHandler)
    client.Subscribe(topic, 2, messagePubHandler)
    
   
    return client
  

}
 
 
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
   
  
   receive_mqtt_packet(msg)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
    log_on_connection()
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
    fmt.Printf("lost connection reason ; %v\n",err)
    log_off_connection()
}
