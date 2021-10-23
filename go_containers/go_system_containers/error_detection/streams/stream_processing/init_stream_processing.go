package monitor_streams


import "fmt"
//import "time"
//
//import "strings"
//import "encoding/json"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/server_libraries/postgres"
//import "lacima.com/Patterns/logging_support"
//import "lacima.com/Patterns/msgpack_2"
//import "github.com/vmihailenco/msgpack/v5"


type Working_Structure_Type struct {
    count int64
    status bool
}




type monitor_stream_type struct {
    
    trim_time                        int64
    sample_time                      int64
    current_time                     int64
    //stream_data                      map[string]stream_records_type
    stream_table                     redis_handlers.Redis_Hash_Struct
    time_table                       redis_handlers.Redis_Hash_Struct
    z_table                         redis_handlers.Redis_Hash_Struct
    z_time                          redis_handlers.Redis_Hash_Struct
    
    process_data_stream              pg_drv.Postgres_Stream_Driver
    filtered_data_stream             pg_drv.Postgres_Stream_Driver
    process_incident_stream          pg_drv.Postgres_Stream_Driver
    process_data_stream_trim         pg_drv.Postgres_Stream_Driver
    filtered_data_stream_trim        pg_drv.Postgres_Stream_Driver
    

}



var default_working_structure  Working_Structure_Type 
var monitor_control            monitor_stream_type 







func Init_data_structures(){
   
    
    initialize_default_structure()
    construct_monitor_control()
 
    
}


func initialize_default_structure(){
    default_working_structure.count = 0
    default_working_structure.status = true
    
}

func construct_monitor_control() {
   
    wd_nodes  := []string{"ERROR_DETECTION:ERROR_DETECTION", "STREAMING_LOGS:STREAMING_LOGS"   }
    nodes := graph_query.Common_qs_search(&wd_nodes)
    node  := nodes[0]
    
    monitor_control.trim_time       = graph_query.Convert_json_int(node["trim_time"])
    monitor_control.sample_time     = graph_query.Convert_json_int(node["sample_time"])
    fmt.Println("sample time",monitor_control.sample_time)
    
    search_list := []string{"ERROR_DETECTION:ERROR_DETECTION", "STREAMING_LOGS:STREAMING_LOGS" ,"STREAM_SUMMARY_DATA"}
    data_element := data_handler.Construct_Data_Structures(&search_list)
	 	 
	monitor_control.stream_table                =   (*data_element)["STREAM_TABLE"].(redis_handlers.Redis_Hash_Struct) 
    monitor_control.time_table                  =   (*data_element)["TIME_TABLE"].(redis_handlers.Redis_Hash_Struct) 
	monitor_control.z_table                     =   (*data_element)["STREAM_TABLE"].(redis_handlers.Redis_Hash_Struct) 
    monitor_control.z_time                     =   (*data_element)["TIME_TABLE"].(redis_handlers.Redis_Hash_Struct)     

	monitor_control.stream_table.Delete_All()
    monitor_control.time_table.Delete_All()
    monitor_control.z_table.Delete_All()
    monitor_control.z_time.Delete_All()
    
    monitor_control.process_data_stream         =   (*data_element)["LOG_STREAM"].(pg_drv.Postgres_Stream_Driver) 
    monitor_control.filtered_data_stream        =   (*data_element)["FILTERED_STREAM"].(pg_drv.Postgres_Stream_Driver)
    
	monitor_control.process_incident_stream     =   (*data_element)["INCIDENT_STREAM"].(pg_drv.Postgres_Stream_Driver) 
	
    monitor_control.process_data_stream_trim    =   (*data_element)["LOG_STREAM"].(pg_drv.Postgres_Stream_Driver) 
    monitor_control.filtered_data_stream_trim   =   (*data_element)["FILTERED_STREAM"].(pg_drv.Postgres_Stream_Driver)	
  
}
    






 
    
    
 
