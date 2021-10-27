package incident_processing


import "fmt"
import "time"
import "strings"
//import "strconv"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/server_libraries/postgres"
import "lacima.com/Patterns/msgpack_2"



type incident_control_type struct {
    
    trim_time                       int64
    sample_time                     int64
    subsystem_id                    string
    keys                            map[string]string
    overall_status                  redis_handlers.Redis_Hash_Struct
    time                            redis_handlers.Redis_Hash_Struct
    status                          redis_handlers.Redis_Hash_Struct
    last_error_data                 redis_handlers.Redis_Hash_Struct
    incident_log                    pg_drv.Postgres_Stream_Driver
    trim_handle                     pg_drv.Postgres_Stream_Driver
}



type incident_record_type struct {
  name               string
  description        string
  max_time_interval  int64
  namespace          string
  key_array          []string
  key                string
  time                redis_handlers.Redis_Single_Structure
  status              redis_handlers.Redis_Single_Structure
  last_error          redis_handlers.Redis_Single_Structure
  
    
}



var incident_control       incident_control_type
var incident_records       []incident_record_type

var valid_state  bool



func Init_data_structures(){
    incident_control.overall_status  =  get_overall_status_hash()
    construct_incident_data_structures()
    construct_incident_data_nodes()
    construct_keys()
    initialize_monitoring_variables()
    
}

func construct_keys(){
    incident_control.keys = make(map[string]string)
    for _,item := range incident_records{
        incident_control.keys[item.namespace] = item.namespace
    }
}



func get_overall_status_hash()redis_handlers.Redis_Hash_Struct{
    
    node_search  := []string{"ERROR_DETECTION:ERROR_DETECTION", "OVERALL_STATUS"  }
    handlers := data_handler.Construct_Data_Structures(&node_search)
    return (*handlers)["OVERALL_STATUS"].(redis_handlers.Redis_Hash_Struct) 
    
    
}






func construct_incident_data_structures() {
   
    node_search  := []string{"ERROR_DETECTION:ERROR_DETECTION", "INCIDENT_STREAMS:INCIDENT_STREAMS"   }
    nodes := graph_query.Common_qs_search(&node_search)
    node  := nodes[0]

    incident_control.trim_time            = graph_query.Convert_json_int(node["trim_time"])
    incident_control.sample_time          = graph_query.Convert_json_int(node["sample_time"])
    incident_control.subsystem_id         = graph_query.Convert_json_string(node["subsystem_id"])
    
    data_node_search  := []string{"ERROR_DETECTION:ERROR_DETECTION", "INCIDENT_STREAMS:INCIDENT_STREAMS" ,"INCIDENT_DATA"  }
    handlers := data_handler.Construct_Data_Structures(&data_node_search)
    incident_control.time                     = (*handlers)["TIME"].(redis_handlers.Redis_Hash_Struct)
    incident_control.status                   = (*handlers)["STATUS"].(redis_handlers.Redis_Hash_Struct)
    incident_control.last_error_data          = (*handlers)["LAST_ERROR"].(redis_handlers.Redis_Hash_Struct)
    incident_control.incident_log             = (*handlers)["INCIDENT_LOG"].(pg_drv.Postgres_Stream_Driver)
    incident_control.trim_handle              = (*handlers)["INCIDENT_LOG"].(pg_drv.Postgres_Stream_Driver)  
    
}
    

func construct_incident_data_nodes(){
    incident_records = make([]incident_record_type,0)
    incident_nodes  := []string{"INCIDENT_LOG"}
    nodes := graph_query.Common_qs_search(&incident_nodes)
   
   
    for _,node := range nodes{
        var item  incident_record_type
       
        item.name               = graph_query.Convert_json_string(node["name"])
        item.description        = graph_query.Convert_json_string(node["description"])
        
        item.namespace          = graph_query.Convert_json_string(node["namespace"])
        item.key_array          = graph_query.Generate_key(item.namespace)
        item.key_array          = append(item.key_array,"INCIDENT_LOG")
        item.key                = strings.Join(item.key_array,"/")
        
        handlers                := data_handler.Construct_Data_Structures(&item.key_array)
        item.time               = (*handlers)["TIME_STAMP"].(redis_handlers.Redis_Single_Structure)
        item.status             = (*handlers)["STATUS"].(redis_handlers.Redis_Single_Structure)
        item.last_error         = (*handlers)["LAST_ERROR"].(redis_handlers.Redis_Single_Structure)
        incident_records        = append(incident_records,item)
    }
    
}




func initialize_monitoring_variables(){
   remove_invalid_keys()
 
   validate_stored_data()
}

func remove_invalid_keys(){
    test_keys("time",incident_control.time)
    test_keys("status",incident_control.status)
    test_keys("last_error_data",incident_control.last_error_data) 
   
    
}

func test_keys(table_name string,redis_hash redis_handlers.Redis_Hash_Struct){
   valid_keys := incident_control.keys
   current_keys := redis_hash.HKeys()
   for _, key := range current_keys {
       if _,ok := valid_keys[key]; ok == false{
           fmt.Println("invalid key",table_name,key)
           redis_hash.HDel(key)
       }
   }
}

func validate_stored_data(){
    
    
                        
    validate_time("time",incident_control.time)
    validate_bool("status",incident_control.status)
    validate_string("last_error_data",incident_control.last_error_data)
    
    
      
}


func validate_time(id_tag string,redis_hash redis_handlers.Redis_Hash_Struct)  {
   msg_pack_time    := msg_pack_utils.Pack_int64(0)
   keys := redis_hash.HKeys()
   for _, key := range keys {
       data :=redis_hash.HGet(key)
       _,err := msg_pack_utils.Unpack_int64(data)
       if err == false {
           //fmt.Println("time bad key",id_tag,key)
           redis_hash.HSet(key,msg_pack_time)
       }
   }
}    
    
func validate_string(id_tag string,redis_hash redis_handlers.Redis_Hash_Struct)  {
   msg_pack_string   := msg_pack_utils.Pack_string("")
   keys := redis_hash.HKeys()
   for _, key := range keys {
       data :=redis_hash.HGet(key)
       _,err := msg_pack_utils.Unpack_string(data)
       if err == false {
           //fmt.Println("string bad key",id_tag,key)
           redis_hash.HSet(key,msg_pack_string)
       }
   }
}    
    
func validate_bool(id_tag string,redis_hash redis_handlers.Redis_Hash_Struct) {
   msg_pack_bool    := msg_pack_utils.Pack_bool(true)
   keys := redis_hash.HKeys()
   for _, key := range keys {
       data :=redis_hash.HGet(key)
       _,err := msg_pack_utils.Unpack_bool(data)
       if err == false {
           //fmt.Println("bool bad key",id_tag, key)
           redis_hash.HSet(key,msg_pack_bool)
       }
   }
}    







func Process_incident_structures(){
        
   go trim_db()
   go process_incident_logs()
  
}





func trim_db(){
  
    for true {
      
      incident_control.trim_handle.Trim(incident_control.trim_time)
      time.Sleep(time.Hour)
      
    }
}

func process_incident_logs(){
    timeout   := time.Duration(incident_control.sample_time)*time.Second
    //time.Sleep(time.Minute)
    for true {
      fmt.Println("checking incident logs")
      check_incident_logs()
     
      time.Sleep(timeout)
      
    }
}


func check_incident_logs(){

    valid_state = true
    for index,_ := range incident_records {
        check_one_incident_log(index )
    }
    incident_control.overall_status.HSet(incident_control.subsystem_id,msg_pack_utils.Pack_bool(valid_state))
}




func check_one_incident_log(index int ){

    item := incident_records[index]
     
     new_time, err  := msg_pack_utils.Unpack_int64(item.time.Get())
     //fmt.Println(err,item.namespace,new_time)
    if (err == true){
        
        ref_time, err := msg_pack_utils.Unpack_int64(incident_control.time.HGet(item.namespace))
        //fmt.Println("\n\nref_time",ref_time,item.namespace)
        if compare_time( err, new_time,ref_time ) == true {
           status,_ := msg_pack_utils.Unpack_bool(item.status.Get())
           //fmt.Println("status",status)
           if status == true {
             
              panic("should not happen")  
           }else{
              valid_state = false
              //fmt.Println("false path",item.namespace)
              process_false_status_data(index)
           }

        }
    }
    
    
}
func compare_time( err bool, new_time, ref_time int64 ) bool {
   return_value := false
   //fmt.Println("new time",err,new_time-ref_time,new_time,ref_time)
   if err == false {
       //fmt.Println("time error false")
       return_value = false
   }
   if new_time > ref_time {
       //fmt.Println("new_time")
       return_value = true
   }else{
     ;//fmt.Println("old time")
   }
   
   return return_value

}

   
    
    
    

    
func process_false_status_data(index int){
    
    item := incident_records[index] 
    key  := item.namespace

    post_postgress_stream_data( item )
    
    incident_control.time.HSet(key,item.time.Get()) 
    incident_control.status.HSet(key,item.status.Get())   
    incident_control.last_error_data.HSet(key,item.last_error.Get())
    
    
}

func post_postgress_stream_data( item incident_record_type ){
    key := item.namespace
    old_value := incident_control.last_error_data.HGet(key)
    new_value := item.last_error.Get()
    fmt.Println("old_value",old_value)
    
    fmt.Println("new_value",new_value)
    if old_value !=  new_value {
        log_postgress_stream( item )
    }
    
}


func log_postgress_stream( item incident_record_type ){
    
    current_state        :=  item.status.Get()
    current_state_bool,_   :=  msg_pack_utils.Unpack_bool(current_state)
    //fmt.Println("current_state",current_state_bool)
    current_state_value  := fmt.Sprintf("%t",current_state_bool)
    //fmt.Println("current_state",current_state_bool,current_state_value)
    status := incident_control.incident_log.Insert( item.namespace,current_state_value,"","","",item.last_error.Get())
    fmt.Println("Postgres log table result ", status)
}
