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
    keys                            map[string]incident_record_type
    description                     redis_handlers.Redis_Hash_Struct
    contact_time                    redis_handlers.Redis_Hash_Struct
    status                          redis_handlers.Redis_Hash_Struct
    last_error_data                 redis_handlers.Redis_Hash_Struct
    last_error_time                 redis_handlers.Redis_Hash_Struct
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
  contact_time       redis_handlers.Redis_Single_Structure
  status              redis_handlers.Redis_Single_Structure
  last_error_time     redis_handlers.Redis_Single_Structure
  last_error_data     redis_handlers.Redis_Single_Structure
  
    
}

type current_data_type struct{
 
    contact_time    string
    status          string
    last_error_time string
    last_error_data       string
    contact_time_unpacked int64
    status_unpacked       bool
    
}

var incident_control       incident_control_type
var incident_records       []incident_record_type
var current_data           current_data_type



func Init_data_structures(){
    
    construct_incident_data_structures()
    construct_incident_data_nodes()

    construct_keys()

   
    
    
    
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
    incident_control.description              = (*handlers)["DESCRIPTION"].(redis_handlers.Redis_Hash_Struct)
    incident_control.contact_time                     = (*handlers)["TIME"].(redis_handlers.Redis_Hash_Struct)
    incident_control.status                   = (*handlers)["STATUS"].(redis_handlers.Redis_Hash_Struct)
    incident_control.last_error_time          = (*handlers)["ERROR_TIME"].(redis_handlers.Redis_Hash_Struct)
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
        //fmt.Println("node",node)
        item.name               = graph_query.Convert_json_string(node["name"])
        item.description        = graph_query.Convert_json_string(node["description"])
        
        item.namespace          = graph_query.Convert_json_string(node["namespace"])
        item.key_array          = graph_query.Generate_key(item.namespace)
        item.key_array          = append(item.key_array,"INCIDENT_LOG")
        item.key                = strings.Join(item.key_array,"/")
        
        handlers                := data_handler.Construct_Data_Structures(&item.key_array)
        item.contact_time       = (*handlers)["TIME_STAMP"].(redis_handlers.Redis_Single_Structure)
        item.status             = (*handlers)["STATUS"].(redis_handlers.Redis_Single_Structure)
        item.last_error_data    = (*handlers)["LAST_ERROR"].(redis_handlers.Redis_Single_Structure)
        item.last_error_time    = (*handlers)["ERROR_TIME"].(redis_handlers.Redis_Single_Structure)
        incident_records        = append(incident_records,item)
    }
    
}







       
    
    
func construct_keys(){
    incident_control.keys = make(map[string]incident_record_type)
    incident_control.description.Delete_All()
    incident_control.contact_time.Delete_All()
    incident_control.status.Delete_All()
    incident_control.last_error_data.Delete_All()
    incident_control.last_error_time.Delete_All()
    for _,item := range incident_records{
        
        incident_control.keys[item.namespace] = item
        incident_control.description.HSet(item.namespace,msg_pack_utils.Pack_string(item.description))
        validate_initial_data(item)
        incident_control.contact_time.HSet(item.namespace,current_data.contact_time)
        incident_control.status.HSet(item.namespace,current_data.status)           
        incident_control.last_error_data.HSet(item.namespace,current_data.last_error_data)
        incident_control.last_error_time.HSet(item.namespace,current_data.last_error_time)
    }
}

func validate_initial_data(item incident_record_type ){
    
    valididate_last_error_data(item.last_error_data)
    validate_status(item.status)
    validate_contact_time(item.contact_time)
    validate_last_error_time(item)
} 

func validate_last_error_time(item incident_record_type){
    
     data :=   item.last_error_time.Get()
    _,err := msg_pack_utils.Unpack_int64(data)
    if err == false {
           item.last_error_time.Set(current_data.contact_time)
           data = current_data.contact_time
    
    }
    current_data.last_error_time = data
}
    
    
    
func validate_contact_time(item  redis_handlers.Redis_Single_Structure){
    msg_pack_time    := msg_pack_utils.Pack_int64(0)
    data :=   item.Get()
    _,err := msg_pack_utils.Unpack_int64(data)
    if err == false {
           item.Set(msg_pack_time)
           data = msg_pack_time
    }
    current_data.contact_time = data
    
}

func validate_status(item  redis_handlers.Redis_Single_Structure){
   msg_pack_bool    := msg_pack_utils.Pack_bool(true) 
   data :=   item.Get()
   _,err := msg_pack_utils.Unpack_bool(data)
   if err == false {
           item.Set(msg_pack_bool)
           data = msg_pack_bool
    
    }
    current_data.status = data
    
}

func valididate_last_error_data(item  redis_handlers.Redis_Single_Structure){
   msg_pack_string   := msg_pack_utils.Pack_string("") 
   data :=   item.Get()
    _,err := msg_pack_utils.Unpack_string(data)
    if err == false {
           item.Set(msg_pack_string)
           data = msg_pack_string
    }
    current_data.last_error_data = data
    
    
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
      //fmt.Println("checking incident logs")
      //fmt.Println("redis length",len(incident_control.status.HKeys()))
      check_incident_logs()
      //fmt.Println("keys",len(incident_control.keys))
      
      time.Sleep(timeout)
      
    }
    
}


func check_incident_logs(){

    //fmt.Println("incident records",len(incident_records))
    for index,_ := range incident_records {
        check_one_incident_log(index )
    }
    
}




func check_one_incident_log(index int ){

    item := incident_records[index]
    key  := item.namespace
     validate_new_data(item)
     new_time  := current_data.contact_time_unpacked 
     ref_time, _ := msg_pack_utils.Unpack_int64(incident_control.contact_time.HGet(key))
     //fmt.Println("new_time",new_time-ref_time)
     if new_time > ref_time   {
        status := current_data.status_unpacked
        if status == true {
             panic("should not happen")
        }else{
             //fmt.Println("process new time",index)
              process_new_status_data(item)
        }
     }
     
    
    
}

    
 
    
func validate_new_data(item incident_record_type ){
    
    check_last_error_data(item.last_error_data)
    check_status(item.status)
    check_contact_time(item.contact_time)
    check_last_error_time(item)
} 

func check_last_error_time(item incident_record_type){
    
     data :=   item.last_error_time.Get()
    _,err := msg_pack_utils.Unpack_int64(data)
    if err == false {
           panic("bad data")
    
    }
    current_data.last_error_time = data
}
    
    
    
func check_contact_time(item  redis_handlers.Redis_Single_Structure){
    
    data :=   item.Get()
    value,err := msg_pack_utils.Unpack_int64(data)
    if err == false {
           panic("bad data")
    }
    current_data.contact_time = data
    current_data.contact_time_unpacked = value
}

func check_status(item  redis_handlers.Redis_Single_Structure){
  
   data :=   item.Get()
   value,err := msg_pack_utils.Unpack_bool(data)
   if err == false {
         panic("bad data")
    
    }
    current_data.status = data
    current_data.status_unpacked = value
    
}

func check_last_error_data(item  redis_handlers.Redis_Single_Structure){
   
   data :=   item.Get()
    _,err := msg_pack_utils.Unpack_string(data)
    if err == false {
      panic("bad data")
    }
    current_data.last_error_data = data
    
    
}

       

  
    
func process_new_status_data(item incident_record_type){
    
   
    
     key := item.namespace
     incident_control.contact_time.HSet(key,current_data.contact_time) 
        
     item.last_error_time.Set(current_data.contact_time)
      
     old_status ,_ := msg_pack_utils.Unpack_bool(incident_control.status.HGet(key))
     if old_status == true {
         incident_control.last_error_time.HSet(key,current_data.contact_time)
     }
     incident_control.status.HSet(key,current_data.status)
    post_postgress_stream_data( item )

    
    
}

func post_postgress_stream_data( item incident_record_type ){
    key := item.namespace
    old_value := get_new_incident_data(key)
    new_value := item.last_error_data.Get()
    //fmt.Println("old_value",old_value)
    
    //fmt.Println("new_value",new_value)
    if old_value !=  new_value {
        log_postgress_stream( item )
    }
    
}

func get_new_incident_data( key string)string {
    
  where_clause   := "tag1 = '"+key+"'  and  time >= 0 ORDER BY time DESC LIMIT 1 "
  pg_data,status := incident_control.incident_log.Select_where(where_clause)
  if status == false {
      panic("should not happen")
  }
  if len(pg_data) == 0 {
      return ""
  }
  return pg_data[0].Data  
}


func log_postgress_stream( item incident_record_type ){
    
    current_state        :=  item.status.Get()
    current_state_bool,_   :=  msg_pack_utils.Unpack_bool(current_state)
    //fmt.Println("current_state",current_state_bool)
    current_state_value  := fmt.Sprintf("%t",current_state_bool)
    //fmt.Println("current_state",current_state_bool,current_state_value)
    incident_control.incident_log.Insert( item.namespace,current_state_value,"","","",item.last_error_data.Get())
    //fmt.Println("Postgres log table result ", status)
}
