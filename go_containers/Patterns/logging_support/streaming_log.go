package logging_support


import "lacima.com/server_libraries/postgres"
import "lacima.com/redis_support/generate_handlers"


func Find_stream_logging_driver()pg_drv.Postgres_Stream_Driver{
    
    search_list := []string{"ERROR_DETECTION:ERROR_DETECTION", "STREAMING_LOGS:STREAMING_LOGS" ,"STREAM_SUMMARY_DATA"}
    data_element := data_handler.Construct_Data_Structures(&search_list)
	 	 
	
	return   (*data_element)[ "LOG_STREAM"].(pg_drv.Postgres_Stream_Driver) 
	     

}




