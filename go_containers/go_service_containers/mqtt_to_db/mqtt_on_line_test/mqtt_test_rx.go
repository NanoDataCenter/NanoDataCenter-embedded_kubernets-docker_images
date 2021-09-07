package mqtt_test

//import b64 "encoding/base64"
import "fmt"
//import "strings"


import "time"
import "github.com/vmihailenco/msgpack/v5"
import "lacima.com/server_libraries/postgres"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"







var redis_topic_value            redis_handlers.Redis_Hash_Struct
var redis_topic_handler          redis_handlers.Redis_Hash_Struct   
 
var postges_topic_stream    pg_drv.Postgres_Stream_Driver      // time stream for all topics
                                                       // tag1 class
                                                       // tag2 device
                                                       // tag3 topic
                                                       // tag4 msgpack handler 
                                                       // tag5 not used
                                                       // data msgpack data



var base_topic_string string

func Mqtt_test_init(site string){
    
   base_topic_string             = "/"+site+"/"
   data_search_list              := []string{"MQTT_IN_SETUP:mqtt_in_setup","TOPIC_STATUS"}
   data_element                  := data_handler.Construct_Data_Structures(&data_search_list)
   redis_topic_value             = (*data_element)["TOPIC_VALUE"].(redis_handlers.Redis_Hash_Struct)
   redis_topic_handler           = (*data_element)["TOPIC_HANDLER"].(redis_handlers.Redis_Hash_Struct)
   postges_topic_stream          = (*data_element)["POSTGRES_DATA_STREAM"].(pg_drv.Postgres_Stream_Driver)
   
    
}




func Mqtt_test_rx(){
    int_topic_handlers()
   
    time_offset := int64(30e9)
    time.Sleep(time.Second*30)
    for true {
        
       current_time := time.Now().UnixNano()
       select_time  := current_time - time_offset
       where_clause := fmt.Sprintf(`(tag1 = 'test_class') and (tag2 = 'test_device') and ( time >=  %d )  ORDER BY time ASC `,select_time)
       
    
       data_array,_ := postges_topic_stream.Select_where(where_clause )
       
       for _, data_element := range data_array{
           fmt.Println(data_element)
           process_rx_element(data_element)
       }
       
       
       time.Sleep(time.Second*30)
        
    }
}

/*
 * type Stream_Output_Data_Record struct {
 
    Stream_id  int64;  id of the record
    Tag1       class
    Tag2       device
    Tag3       full topic
    Tag4       topic_handler
    Tag5       time string
    Data       msgpack data
    Time_stamp int64;  unix time
}    
*/


type topic_handler_type  func( data  string )
type topic_handler_map_type map[string]topic_handler_type

var topic_handler_map topic_handler_map_type

func int_topic_handlers(){
    topic_handler_map = make(map[string]topic_handler_type)
    topic_handler_map["string"]               = process_string
    topic_handler_map["int32"]                = process_int32
    topic_handler_map["float64"]              = process_float
    topic_handler_map["map[string]interface"] = process_map
    topic_handler_map["[]float32"]             = process_float_array
}
  
func process_rx_element( data_element pg_drv.Stream_Output_Data_Record ){

    handler                  := data_element.Tag4
     
    
    fmt.Println("class ",data_element.Tag1)
    fmt.Println("device ", data_element.Tag2)
    fmt.Println("full topic",data_element.Tag3)
    fmt.Println("topic id ",handler)
    fmt.Println("time_string", data_element.Tag5)
    fmt.Println("time",data_element.Time_stamp)
    
    topic_handler_map[handler](data_element.Data)
   
    
    // testing redis 
    redis_data :=  redis_topic_value.HGet(data_element.Tag3)
    if redis_data !="" {
    
    handler = redis_topic_handler.HGet(data_element.Tag3)
    
    
    topic_handler_map[handler](redis_data)
    }
}

func  process_string(data string ){
   var result string
    err := msgpack.Unmarshal([]byte(data), &result)
    if err != nil {
        panic(err)
    }
    fmt.Println("string result",result)
}
   
func  process_int32(data string ){
   var result int32
    err := msgpack.Unmarshal([]byte(data), &result)
    if err != nil {
        panic(err)
    }
    fmt.Println("int32 result",result)
}


func  process_float(data string ){
   var result float64
    err := msgpack.Unmarshal([]byte(data), &result)
    if err != nil {
        panic(err)
    }
    fmt.Println("float 64 result",result)
}


func  process_map(data string ){
    var result map[string]interface{}
    err := msgpack.Unmarshal([]byte(data), &result)
    if err != nil {
        panic(err)
    }
    fmt.Println("map result",result)
}



func  process_float_array(data string ){
   var result []float32
    err := msgpack.Unmarshal([]byte(data), &result)
    if err != nil {
        panic(err)
    }
    fmt.Println("[]float32 result",result)
} 
    
