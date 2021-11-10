package stream_support



//import "fmt"
//import "time"
//
//import "strings"
//import "encoding/json"
//import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/server_libraries/postgres"
//import "lacima.com/Patterns/logging_support"
//import "lacima.com/Patterns/msgpack_2"
//import "github.com/vmihailenco/msgpack/v5"

var Z_LEVEL  float64

type Median_Filter_Type struct {
 
    buffer_position int64
    buffer_limit    int64
    median_buffer   []float64
    filtered_value  float64
    current_value   float64
}


type Velocity_Type struct {
    
   previous_value    float64
   current_velocity  float64
   lag_velocity      float64
   r_value           float64
    
}

type Z_Type struct {
    z_value           float64
    std               float64
    z_state           bool  
    
    
}


type Stream_Processing_Type struct {
    median    Median_Filter_Type
    velocity  Velocity_Type
    z_data    Z_Type
}




type monitor_stream_type struct {
    
    stream_table                     redis_handlers.Redis_Hash_Struct
    time_table                       redis_handlers.Redis_Hash_Struct
    z_table                         redis_handlers.Redis_Hash_Struct
    z_time                          redis_handlers.Redis_Hash_Struct
    
    process_data_stream              pg_drv.Postgres_Stream_Driver
    filtered_data_stream             pg_drv.Postgres_Stream_Driver
    process_incident_stream          pg_drv.Postgres_Stream_Driver

}




var stream_control            monitor_stream_type 





func init_stream_data_structures() {
   
    
    search_list := []string{"ERROR_DETECTION:ERROR_DETECTION", "STREAMING_LOGS:STREAMING_LOGS" ,"STREAM_SUMMARY_DATA"}
    data_element := data_handler.Construct_Data_Structures(&search_list)
	 	 
	stream_control.stream_table                =   (*data_element)["STREAM_TABLE"].(redis_handlers.Redis_Hash_Struct) 
    stream_control.time_table                  =   (*data_element)["TIME_TABLE"].(redis_handlers.Redis_Hash_Struct) 
	stream_control.z_table                     =   (*data_element)["STREAM_TABLE"].(redis_handlers.Redis_Hash_Struct) 
    stream_control.z_time                     =   (*data_element)["TIME_TABLE"].(redis_handlers.Redis_Hash_Struct)     

	
    
    stream_control.process_data_stream         =   (*data_element)["LOG_STREAM"].(pg_drv.Postgres_Stream_Driver) 
    
    
    stream_control.filtered_data_stream        =   (*data_element)["FILTERED_STREAM"].(pg_drv.Postgres_Stream_Driver)
    
	stream_control.process_incident_stream     =   (*data_element)["INCIDENT_STREAM"].(pg_drv.Postgres_Stream_Driver) 
	
   
}
    
