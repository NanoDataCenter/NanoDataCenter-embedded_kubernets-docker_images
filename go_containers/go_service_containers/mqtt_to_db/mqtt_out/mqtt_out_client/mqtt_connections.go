package mqtt_out_client

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
func Construct_mqtt_actions( ip string, port int){  // setup receiving mqtt construct_mqtt_tasks
    
    data_nodes := graph_query.Common_qs_search(&[]string{"MQTT_CLIENT_ID:MQTT_CLIENT_ID"})
    data_node  := data_nodes[0]
    incident_log = logging_support.Construct_incident_log([]string{"MQTT_OUTPUT_SETUP:site_out_server","INCIDENT_LOG:MQTT_TX_CONNECTION_LOST","INCIDENT_LOG"} )
  
    
    //su.Bc_Rec.Add_header_node("MQTT_OUTPUT_SETUP","site_out_server",make(map[string]interface{}))
    //su.Construct_incident_logging("MQTT_TX_CONNECTION_LOST","MQTT_TX_CONNECTION_LOST",su.Error)
  
    mqtt_client_map := graph_query.Convert_json_dict( data_node["mqtt_client_id_map"] )
    
    if _,ok := mqtt_client_map["mqtt_output_server" ]; ok == false {
        panic("no existant mqtt client map")
    }
    
    mqtt_client_id := mqtt_client_map["mqtt_output_server"]
    
    
    
    
    connection_status = false    
    opts := mqtt.NewClientOptions()
    opts.AddBroker(fmt.Sprintf("tcp://%s:%d", ip, port))
    opts.SetClientID(mqtt_client_id)
    opts.SetAutoReconnect(true)
    opts.SetCleanSession(true)
    opts.SetConnectRetry(true)
    opts.SetConnectRetryInterval(time.Second *5 )
    opts.SetConnectTimeout(time.Second*30)
    opts.SetKeepAlive(time.Second*5)
    opts.SetMaxReconnectInterval(time.Second*5)
     

    opts.OnConnect = connectHandler
    opts.OnConnectionLost = connectLostHandler 
    client = mqtt.NewClient(opts)
    token := client.Connect()
    if token.WaitTimeout(time.Second*30) == false {
        panic("mqtt lost time out")
    }
    if token.Error() != nil {
         panic("mqtt token error")
    }
    
    set_up_rpc_server()
    

}
 
func Wait_for_connections(){
    for connection_status == false {
        ;
    }
}


var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
    
    log_on_connection()
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
    log_off_connection()
}

func log_on_connection(){
    fmt.Println("mqtt out connection on")
    connection_status = true
  
    
}


func log_off_connection(){
    fmt.Println("mqtt out connection off")
    connection_status = false
    error_string := "MQTT_TX_CONNECTION_LOST "+strconv.FormatInt(time.Now().UnixNano(),10)
    incident_log.Post_event(error_string)
    
}
