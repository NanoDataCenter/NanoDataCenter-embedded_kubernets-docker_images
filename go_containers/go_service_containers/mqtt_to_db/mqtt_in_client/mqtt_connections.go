package mqtt_client

import "fmt"
//import "os"
import "strconv"
import "time"
import "lacima.com/redis_support/graph_query"
import mqtt "github.com/eclipse/paho.mqtt.golang"
import "lacima.com/Patterns/logging_support"

var client mqtt.Client
var connection_status bool
var incident_log *logging_support.Incident_Log_Type

func Construct_mqtt_actions( ip string, port int)mqtt.Client{  // setup receiving mqtt construct_mqtt_tasks

    data_nodes := graph_query.Common_qs_search(&[]string{"MQTT_CLIENT_ID:MQTT_CLIENT_ID"})
    data_node  := data_nodes[0]
    incident_log = logging_support.Construct_incident_log([]string{"MQTT_IN_SETUP:site_in_server","INCIDENT_LOG:MQTT_RX_CONNECTION_LOST","INCIDENT_LOG"} )
    
   
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
        panic("mqtt time out")
    }
    if token.Error() != nil {
        panic("mqtt token error")
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
    fmt.Println("mqtt rx connected")
    log_on_connection()
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
    fmt.Println("mqtt tx connection")
    connection_status = false
   error_string := "MQTT_RX_CONNECTION_LOST "+strconv.FormatInt(time.Now().UnixNano(),10)
    incident_log.Post_event(error_string)
}
