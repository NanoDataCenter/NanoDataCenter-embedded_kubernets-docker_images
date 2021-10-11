package monitor_streams


import "time"
import "fmt"



func Process_functions(){
    
    initalize_stream_processing()
    for true {
       //time.Sleep(time.Duration(monitor_control.sample_time)* time.Minute)
       process_stream_logs()
       panic("done")
       
    
    
    }
    
}

func initalize_stream_processing(){
    
    monitor_control.current_time = time.Now().UnixNano()   
    
}


func process_stream_logs(){
    
    for key, redis_stream := range monitor_control.input_stream_map{
        
      fmt.Println("key",key,redis_stream)
      fmt.Println("\n\n\n",monitor_control.description[key],monitor_control.name[key],"\n\n\n\n")
    }
     monitor_control.current_time = time.Now().UnixNano()
}
