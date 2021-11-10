package monitor_streams

import "time"



func Trim_time_series(){

    for true {
        time.Sleep(time.Hour)
        execute_time_series_trim()
        
        
        
    }
}


func execute_time_series_trim(){
    monitor_control.filtered_log_stream_trim  .Trim( monitor_control.trim_time )     
    monitor_control.filtered_data_stream_trim.Trim( monitor_control.trim_time)     
    monitor_control.filtered_incident_stream_trim.Trim( monitor_control.trim_time)    
}
