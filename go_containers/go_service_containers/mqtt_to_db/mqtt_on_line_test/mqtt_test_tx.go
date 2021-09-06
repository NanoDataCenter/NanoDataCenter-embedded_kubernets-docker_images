package mqtt_test
import "fmt"

import "time"
import "github.com/vmihailenco/msgpack/v5"
import mqtt "github.com/eclipse/paho.mqtt.golang"


    
type topic_test_type struct {
    topic string
    data  string
}


func Mqtt_test_tx( client mqtt.Client){
  
  
  topic_data := generate_test_data()  
  
  for true {
        fmt.Println("transmitting data")
        for _,element := range  topic_data { 
              
              t := client.Publish(base_topic_string+element.topic, 2,false, element.data)
              _ = t.Wait()
              
        }
        time.Sleep(time.Second *30)
  }
}


func generate_test_data()[]topic_test_type{
    
    return_value := make([]topic_test_type,5)
    return_value[0] = generate_topic_1()
    return_value[1] = generate_topic_2()
    return_value[2] = generate_topic_3()
    return_value[3] = generate_topic_4()
    return_value[4] = generate_topic_5()
    return return_value
}

func generate_topic_1()topic_test_type{
 
    var return_value topic_test_type
    return_value.topic = "test_class/test_device/test_string"
    
    input_data := "test string"
    b, err := msgpack.Marshal(&input_data)
    if err != nil {
        panic(err)
    }
    return_value.data  = string(b)
    
    return return_value
    
    
}



func generate_topic_2()topic_test_type{
 
    var return_value topic_test_type
    return_value.topic = "test_class/test_device/test_int"
    
    input_data := int32(14)
    b, err := msgpack.Marshal(&input_data)
    if err != nil {
        panic(err)
    }
    return_value.data  =  string(b)
    
    return return_value
    
    
}



func generate_topic_3()topic_test_type{
 
    var return_value topic_test_type
    return_value.topic = "test_class/test_device/test_float"
    
    input_data := float64(32.)
    b, err := msgpack.Marshal(&input_data)
    if err != nil {
        panic(err)
    }
    return_value.data  =  string(b)
    
    return return_value
    
    
}


func generate_topic_4()topic_test_type{
 
    var return_value topic_test_type
    return_value.topic = "test_class/test_device/test_map"
    
    input_data := make(map[string]interface{})
    input_data["string"] = "string"
    input_data["int"]    = int16(1234)
    input_data["float"]  = float32(34.)
    
    
    b, err := msgpack.Marshal(&input_data)
    if err != nil {
        panic(err)
    }
    return_value.data  =  string(b)
    
    return return_value
    
    
}


func generate_topic_5()topic_test_type{
    var return_value topic_test_type
    return_value.topic = "test_class/test_device/test_array"
    
    input_data := []float32{ 1.1,1.2,1.3,1.4}
    b, err := msgpack.Marshal(&input_data)
    if err != nil {
        panic(err)
    }
    return_value.data  =  string(b)
    
    return return_value
    
    
}    
    


