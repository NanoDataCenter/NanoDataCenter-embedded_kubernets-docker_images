 package mqtt_db_trim
 
//import "fmt"
import "time"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/server_libraries/postgres"
 
 
 
 
var postges_topic_stream    pg_drv.Postgres_Stream_Driver      // time stream for all topics
                                                       // tag1 class
                                                       // tag2 device
                                                       // tag3 topic
                                                       // tag4 msgpack handler 
                                                       // tag5 not used
                                                       // data msgpack data

                                                        
 
var  trim_time int64 

 
func Trim_int(trim_time_seconds int64) { 
    
   trim_time                     = trim_time_seconds              
   data_search_list              := []string{"MQTT_OUTPUT_SETUP:site_out_server","TOPIC_STATUS"}
   data_element                  := data_handler.Construct_Data_Structures(&data_search_list)
   postges_topic_stream          = (*data_element)["POSTGRES_DATA_STREAM"].(pg_drv.Postgres_Stream_Driver)
 

    
}




func Trim_dbs(){
    
    
   for true {
       
      postges_topic_stream.Trim(trim_time)
      postges_topic_stream.Vacuum()
      
      //fmt.Println("made it here")
      time.Sleep(time.Second *3600)

       
   }
    
}