package mqtt_client

import "fmt"
import "os"
import "time"
import mqtt "github.com/eclipse/paho.mqtt.golang"

var client mqtt.Client
var connection_status bool

func Construct_mqtt_actions( ip string, port int)mqtt.Client{  // setup receiving mqtt construct_mqtt_tasks
    
    connection_status = false    
    opts := mqtt.NewClientOptions()
    opts.AddBroker(fmt.Sprintf("tcp://%s:%d", ip, port))
    opts.SetClientID("mqtt_to_db_out")
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
   
    token = client.Subscribe(topic, 2, nil)
    token.Wait()
    return client
    

}
 
 
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
   
  
   receive_mqtt_packet(msg)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
    log_on_connection()
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
    log_off_connection()
}
