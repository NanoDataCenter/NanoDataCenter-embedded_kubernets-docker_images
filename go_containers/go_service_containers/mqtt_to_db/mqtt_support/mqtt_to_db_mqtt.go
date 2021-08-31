package mqtt_support

import "fmt"
import "time"
import mqtt "github.com/eclipse/paho.mqtt.golang"



func Construct_mqtt_actions( ip string, port int){  // setup receiving mqtt construct_mqtt_tasks
    ip = "localhost"
    port = 1883
    opts := mqtt.NewClientOptions()
    opts.AddBroker(fmt.Sprintf("tcp://%s:%d", ip, port))
    opts.SetClientID("mqtt_to_db")
    opts.SetAutoReconnect(true)
    opts.SetCleanSession(true)
    opts.SetConnectRetry(true)
    opts.SetConnectRetryInterval(time.Minute)
    opts.SetConnectTimeout(time.Second*30)
    opts.SetKeepAlive(time.Second*30)
    opts.SetMaxReconnectInterval(time.Minute)
        
        
    opts.SetDefaultPublishHandler(messagePubHandler)
    opts.OnConnect = connectHandler
    opts.OnConnectionLost = connectLostHandler 
    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
    //fmt.Println("connection received")
    topic := get_monitoring_topic()
    //fmt.Println("topic",topic)
    token := client.Subscribe(topic, 2, nil)
    token.Wait()

}
 
 
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
   
   fmt.Println("message recieved")
   receive_mqtt_packet(msg)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
    log_on_connection()
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
    log_off_connection()
}
