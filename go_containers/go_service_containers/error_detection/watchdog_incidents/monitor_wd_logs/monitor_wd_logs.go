package monitor_wd_logs


import "fmt"
import "time"
//import "strings"

import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/server_libraries/postgres"
import "lacima.com/Patterns/msgpack_2"



type wd_control_type struct {
    
    trim_time                       int64
    sample_time                     int64
    debounce_counts                  int64
    subsystem_id                    string
    overall_status                  redis_handlers.Redis_Hash_Struct
    debounced_status                redis_handlers.Redis_Hash_Struct
    status                          redis_handlers.Redis_Hash_Struct
    time_stamp                      redis_handlers.Redis_Hash_Struct
    description                     redis_handlers.Redis_Hash_Struct
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
  counts             int64
  watch_dog_time     redis_handlers.Redis_Single_Structure 
  
    
}


var wd_control       wd_control_type
var wd_records       []wd_record_type


var overall_state   bool


//var   current_time_map    map[string]int64
var   current_state_map  map[string]bool
var   current_count_map   map[string]int64


var msg_pack_true  string
var msg_pack_false string

func Init_data_structures(){
    msg_pack_true   = msg_pack_utils.Pack_bool(true)
    msg_pack_false  = msg_pack_utils.Pack_bool(false)
    
    
    wd_control.overall_status = get_overall_status_hash()
    construct_wd_data_structures()
    construct_wd_nodes()
    
    initialize_monitoring_variables()
    
}

func get_overall_status_hash()redis_handlers.Redis_Hash_Struct{
    
    node_search  := []string{"ERROR_DETECTION:ERROR_DETECTION", "OVERALL_STATUS"  }
    handlers := data_handler.Construct_Data_Structures(&node_search)
    return (*handlers)["OVERALL_STATUS"].(redis_handlers.Redis_Hash_Struct) 
    
    
}

func construct_wd_data_structures() {
   
    wd_nodes  := []string{"ERROR_DETECTION:ERROR_DETECTION", "WD_DETECTION:WD_DETECTION"   }
    nodes := graph_query.Common_qs_search(&wd_nodes)
    node  := nodes[0]
    
    wd_control.trim_time                = graph_query.Convert_json_int(node["trim_time"])
    wd_control.sample_time              = graph_query.Convert_json_int(node["sample_time"])
    wd_control.debounce_counts          = int64(graph_query.Convert_json_int(node["debounce_count"]))
    wd_control.subsystem_id             = graph_query.Convert_json_string(node["subsystem_id"])
    wd_data_nodes                       := []string{"ERROR_DETECTION:ERROR_DETECTION", "WD_DETECTION:WD_DETECTION" ,"WATCH_DOG_DATA"  }
    handlers                            := data_handler.Construct_Data_Structures(&wd_data_nodes)
    wd_control.debounced_status         = (*handlers)["DEBOUNCED_STATUS"].(redis_handlers.Redis_Hash_Struct)
    wd_control.status                   = (*handlers)["STATUS"].(redis_handlers.Redis_Hash_Struct)
    wd_control.time_stamp               = (*handlers)["TIME_STAMP"].(redis_handlers.Redis_Hash_Struct)
    wd_control.description              = (*handlers)["DESCRIPTION"].(redis_handlers.Redis_Hash_Struct)
    wd_control.wd_incidents             = (*handlers)["WATCH_DOG_LOG"].(pg_drv.Postgres_Stream_Driver)    
    wd_control.trim_handle              = (*handlers)["WATCH_DOG_LOG"].(pg_drv.Postgres_Stream_Driver)  
    
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
        item.max_time_interval  = item.max_time_interval*1e9
        item.namespace          = graph_query.Convert_json_string(node["namespace"])
       
        item.key_array          = graph_query.Generate_key(item.namespace)
        item.key_array          = append(item.key_array,"WATCH_DOG")
        item.key                = item.namespace
        handlers                := data_handler.Construct_Data_Structures(&item.key_array)
        item.watch_dog_time     = (*handlers)["WATCH_DOG_TS"].(redis_handlers.Redis_Single_Structure)
        
        wd_records              = append(wd_records,item)
        
    }
    
}

 


func initialize_monitoring_variables(){
    
                           
    wd_control.debounced_status.Delete_All()                 
    wd_control.status.Delete_All()
    wd_control.time_stamp.Delete_All()
    wd_control.description.Delete_All()
   
    
    current_state_map   = make(map[string]bool)
    current_count_map  = make(map[string]int64)
    current_time  := time.Now().UnixNano()
   
    msg_pack_time    := msg_pack_utils.Pack_int64(current_time)
    
   
    for _, wd_record := range wd_records{
        key := wd_record.key
        current_state_map[key] = true
        current_count_map[key] = 0
        wd_control.description.HSet(key,wd_record.description)
        wd_control.status.HSet(key,msg_pack_true)
        wd_control.debounced_status.HSet(key,msg_pack_true)
        wd_control.time_stamp.HSet(key,msg_pack_time)
       
    }
    
    
}


func Process_wd_queues(){
        
   go trim_db()
   time.Sleep(time.Second*5)
   go check_watch_dogs()
  
}





func trim_db(){
  
    for true {
      
      wd_control.trim_handle.Trim(wd_control.trim_time)
      time.Sleep(time.Hour)
      
    }
}

func check_watch_dogs(){
    timeout   := time.Duration(wd_control.sample_time)*time.Second
    //time.Sleep(time.Minute)
    for true {
      fmt.Println("checking watchdog queue")
      check_queues()
     
      time.Sleep(timeout)
      
    }
}


func check_queues(){
   current_time  := time.Now().UnixNano()
   
   overall_state  = true
   for _,wd_record := range wd_records{
       key := wd_record.key
       check_wd_record(key, wd_record, current_time)
   }
   
   wd_control.overall_status.HSet(wd_control.subsystem_id,msg_pack_utils.Pack_bool(overall_state))
}       
    
   
    
func check_wd_record(key string, wd_record wd_record_type, current_time int64){
    
    
     previous_state := current_state_map[key]
     
     new_state  :=   determine_device_state( key,wd_record, current_time)
     current_state_map[key] = new_state
     //fmt.Println("new_state",new_state,previous_state,wd_record.namespace)
     if previous_state != new_state {
         
         if new_state == true {
             wd_control.status.HSet(key,msg_pack_true)
             //fmt.Println("transistion to monitor true",wd_record.namespace)
            
         }else{
             //fmt.Println("transistion to false",wd_record.namespace)
             current_count_map[wd_record.namespace] = wd_control.debounce_counts
             wd_control.status.HSet(key,msg_pack_false)
             wd_control.debounced_status.HSet(key,msg_pack_false)
             wd_control.wd_incidents.Insert(key,"false","" ,"","","" )
             
             
         }
     }
     //fmt.Println(new_state,current_count_map[wd_record.namespace],wd_record.namespace)
     if ( new_state == true ) && (current_count_map[wd_record.namespace] >0) {
        current_count_map[wd_record.namespace] = current_count_map[wd_record.namespace] - 1
        if current_count_map[wd_record.namespace] == 0 {
            //fmt.Println("transistion to true",wd_record.namespace)
            wd_control.debounced_status.HSet(key,msg_pack_true)
            wd_control.wd_incidents.Insert(key,"true","" ,"","","" )
        }
             
     }
}

func determine_device_state(key string, wd_record wd_record_type, current_time int64)bool {
    
     
    
     wd_time_string       :=        wd_record.watch_dog_time.Get()
     wd_time_int64, err   :=        msg_pack_utils.Unpack_int64(wd_time_string) 
     
    
     if err != true { 
         
         return false
     }
     wd_control.time_stamp.HSet(key,wd_time_string)
     //fmt.Println( wd_time_int64  - current_time + wd_record.max_time_interval )
     if wd_time_int64  > current_time - wd_record.max_time_interval {
           
           return  true
      }
           
     
     return false           
    
}





