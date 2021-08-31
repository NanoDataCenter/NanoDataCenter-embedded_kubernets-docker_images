package mqtt_support



import "time"
import "fmt"

var device_status           map[string]bool


func initialize_device_status(){
 
    device_status = make(map[string]bool)
    for key,_ := range device_map{
        device_status[key] = true
        redis_device_status.HSet(key,"true")
    }
    
}



func Monitor_devices(){
    initialize_device_status()
    for true {
        time.Sleep(time.Second*15)
        check_all_devices()
	 }

}

func check_all_devices(){
    time_stamp := time.Now().Unix()
    for key, item := range contact_map {
         fmt.Println("key",key,item.contact_time,item.delta_time)
         time_out_value := item.contact_time+item.delta_time
         fmt.Println("time_stamp",time_stamp, time_out_value)
         if time_stamp > time_out_value {
            if  device_status[key] == true {
              fmt.Println("device change false",device_map[key].class,key)
              postgres_incident_stream.Insert(device_map[key].class,key,"false","","","")  
            }
            redis_device_status.HSet(key,"false")
            device_status[key] = false
         }else {
            if device_status[key] == false {
                fmt.Println("device change true",device_map[key].class,key)
                postgres_incident_stream.Insert(device_map[key].class,key,"true","","","")  
            }
            redis_device_status.HSet(key,"true")
            device_status[key] = true
         }
    }
}
        
