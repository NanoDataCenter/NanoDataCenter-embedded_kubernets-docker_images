package monitor_wd_logs


import "fmt"
import "time"
import "strings"
import "strconv"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/server_libraries/postgres"
import "github.com/vmihailenco/msgpack/v5"


type wd_control_type struct {
    start_delay                     int64
    trim_time                       int64
    wd_time                         int64
    max_count                       int
    status                          redis_handlers.Redis_Single_Structure
    composite_value                 redis_handlers.Redis_Hash_Struct
    wd_value                        redis_handlers.Redis_Hash_Struct
    wd_ts                           redis_handlers.Redis_Hash_Struct
    state_change_counts             redis_handlers.Redis_Hash_Struct
    wd_incidents                    pg_drv.Postgres_Stream_Driver
    trim_handle                     pg_drv.Postgres_Stream_Driver

}



type wd_record_type struct {
  name               string
  description        string
  max_time_interval  int64
  namespace          string
  key_array          []string
  key                string
  watch_dog_time     redis_handlers.Redis_Single_Structure 
  watch_dog_state    redis_handlers.Redis_Single_Structure
    
}


var wd_control       wd_control_type
var wd_records       []wd_record_type


var overall_state   bool


//var   current_time_map    map[string]int64
var   current_state_map   map[string]bool


func Init_data_structures(){
    construct_wd_data_structures()
    construct_wd_nodes()
    
    initialize_monitoring_variables()
    
}



func construct_wd_data_structures() {
   
    wd_nodes  := []string{"ERROR_DETECTION:ERROR_DETECTION", "WD_DETECTION:WD_DETECTION"   }
    nodes := graph_query.Common_qs_search(&wd_nodes)
    node  := nodes[0]
    wd_control.start_delay      = graph_query.Convert_json_int(node["start_delay"])
    wd_control.trim_time        = graph_query.Convert_json_int(node["trim_time"])
    wd_control.wd_time          = graph_query.Convert_json_int(node["wd_time"])
    wd_control.max_count        = int(graph_query.Convert_json_int(node["max_count"]))
    
    wd_data_nodes  := []string{"ERROR_DETECTION:ERROR_DETECTION", "WD_DETECTION:WD_DETECTION" ,"WATCH_DOG_DATA"  }
    handlers := data_handler.Construct_Data_Structures(&wd_data_nodes)
    wd_control.status                   = (*handlers)["WATCH_DOG_STATUS"].(redis_handlers.Redis_Single_Structure)
    wd_control.composite_value          = (*handlers)["WATCH_DOG_COMPOSITE_VALUE"].(redis_handlers.Redis_Hash_Struct)
    wd_control.wd_value                 = (*handlers)["WATCH_DOG_VALUE"].(redis_handlers.Redis_Hash_Struct)
    wd_control.wd_ts                    = (*handlers)["WATCH_DOG_STAMP"].(redis_handlers.Redis_Hash_Struct)
    wd_control.state_change_counts      = (*handlers)["STATE_CHANGE_COUNTS"].(redis_handlers.Redis_Hash_Struct)
    wd_control.wd_incidents             = (*handlers)["WATCH_DOG_INCIDENTS"].(pg_drv.Postgres_Stream_Driver)    
    wd_control.trim_handle              = (*handlers)["WATCH_DOG_INCIDENTS"].(pg_drv.Postgres_Stream_Driver)  
    
}
    



func construct_wd_nodes(){
    wd_records = make([]wd_record_type,0)
    wd_nodes  := []string{"WATCH_DOG"}
    nodes := graph_query.Common_qs_search(&wd_nodes)
    for _,node := range nodes{
        var item  wd_record_type
       
        item.name               = graph_query.Convert_json_string(node["name"])
        item.description        = graph_query.Convert_json_string(node["description"])
        item.max_time_interval  = graph_query.Convert_json_int(node["max_time_interval"])
        item.namespace          = graph_query.Convert_json_string(node["namespace"])
        item.key_array          = graph_query.Generate_key(item.namespace)
        item.key_array          = append(item.key_array,"WATCH_DOG")
        item.key                = strings.Join(item.key_array,"/")
        handlers                := data_handler.Construct_Data_Structures(&item.key_array)
        item.watch_dog_time     = (*handlers)["WATCH_DOG_TS"].(redis_handlers.Redis_Single_Structure)
        
        wd_records              = append(wd_records,item)
    }
    
}

 



func initialize_monitoring_variables(){
    
                           
    wd_control.composite_value.Delete_All()                 
    wd_control.wd_value.Delete_All()
    wd_control.wd_ts.Delete_All()
    wd_control.state_change_counts.Delete_All()
   
    
    current_state_map   = make(map[string]bool)
    
    current_time  := time.Now().Unix()
    time_stamp    := strconv.FormatInt(current_time,10)
    for _, wd_record := range wd_records{
        key := wd_record.key
        
        //current_time_map[key]  = current_time
        current_state_map[key] = true
        wd_control.wd_value.HSet(key,"true")
        wd_control.composite_value.HSet(key,"true")
        wd_control.wd_ts.HSet(key,time_stamp)
        wd_control.state_change_counts.HSet(key,"0")
    }
    
    
}


func Process_wd_queues(){
        
   go trim_db()
   go check_watch_dogs()
  
}





func trim_db(){
  
    for true {
      
      wd_control.trim_handle.Trim(wd_control.trim_time)
      time.Sleep(time.Hour)
      
    }
}

func check_watch_dogs(){
    timeout   := time.Duration(wd_control.wd_time)*time.Second
    //time.Sleep(time.Minute)
    for true {
      
      check_queues()
     
      time.Sleep(timeout)
      
    }
}


func check_queues(){
   current_time  := time.Now().Unix()
   time_stamp    := strconv.FormatInt(current_time,10)
   overall_state  = true
   for _,wd_record := range wd_records{
       key := wd_record.key
       check_wd_record(key, wd_record, current_time,time_stamp )
   }
   if overall_state == true {
       wd_control.status.Set("true")
   }else{
       wd_control.status.Set("false")
   }
   fmt.Println("overall status",wd_control.status.Get() )
}       
    
   
    
func check_wd_record(key string, wd_record wd_record_type, current_time int64, time_stamp string){
    
    
     previous_state := current_state_map[key]
     
     new_state  :=   determine_device_state( key,wd_record, current_time,time_stamp)
     if new_state == false {
         
         increment_count(key)
         wd_control.wd_value.HSet(key,"false")
         wd_control.composite_value.HSet(key,"false")
         
         
     }else{
         wd_control.wd_value.HSet(key,"true")
          decrement_count(key)
           
     }
     if previous_state != new_state {
         
         if new_state == true {
            
             wd_control.wd_incidents.Insert(key,"true",time_stamp ,"","","" )
         }else{
             wd_control.wd_incidents.Insert(key,"false",time_stamp ,"","","" )
             
             
         }
     }
     
     if is_count_zero(key) == true {
       
            wd_control.composite_value.HSet(key,"true")
     }
     
     fmt.Println("counts",wd_control.state_change_counts.HGet(key))
     fmt.Println("composite",wd_control.composite_value.HGet(key))
     fmt.Println("actual value",wd_control.wd_value.HGet(key))
     
}

func determine_device_state(key string, wd_record wd_record_type, current_time int64, time_stamp string)bool {
    
     
     device_state := false
     wd_time_string := wd_record.watch_dog_time.Get()
     wd_time_int64, err   :=  msg_pack_convert(wd_time_string) 
     if err != true {
         device_state = false
     }else{
       
       //fmt.Println("time",wd_time_int64,current_time - wd_record.max_time_interval,wd_record.max_time_interval,current_time)
       
       if wd_time_int64  > current_time - wd_record.max_time_interval {
           device_state = true
       }else{
           device_state = false
       }
       
     }
     if device_state == true {
         
         wd_control.wd_ts.HSet(key,time_stamp)
         current_state_map[key] = true
         //current_time_map[key]  = current_time
         
     }else{
        wd_control.wd_value.HSet(key,"false")
        current_state_map[key] = false
     }
     
     
     fmt.Println("device state",key,device_state)
     return device_state           
    
}

func msg_pack_convert( msgpack_data string )(int64,bool){

   var result int64
    err := msgpack.Unmarshal([]byte(msgpack_data), &result)
    if err != nil {
        return 0,false
    }
    result = result/1e9
    return result, true
    
}



func is_count_zero(key string )bool{
    
    temp_str     := wd_control.state_change_counts.HGet(key)
    temp_int,_     := strconv.Atoi(temp_str)
    if temp_int == 0 {
        return true
    }
    return false
    
}

func decrement_count(key string){
    
    
    temp_str     := wd_control.state_change_counts.HGet(key)
    temp_int,_   := strconv.Atoi(temp_str)
    if temp_int >0 {
       temp_int      =  temp_int - 1 
       overall_state = false
    }
    
    temp_str     = strconv.Itoa(temp_int)
    wd_control.state_change_counts.HSet(key,temp_str)
}
func increment_count(key string){
    
    
    overall_state = false
    temp_str     := wd_control.state_change_counts.HGet(key)
    temp_int,_   := strconv.Atoi(temp_str)
    temp_int      =  temp_int + 1 
    if temp_int >    wd_control.max_count {
        temp_int =   wd_control.max_count
    }
    temp_str     = strconv.Itoa(temp_int)
    wd_control.state_change_counts.HSet(key,temp_str)
}

func clear_count(key string){
    
    wd_control.state_change_counts.HSet(key,"0")
    
}


