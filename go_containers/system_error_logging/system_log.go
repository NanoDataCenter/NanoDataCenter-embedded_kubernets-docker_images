package system_log


import "time"
import "bytes"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"
import "github.com/msgpack/msgpack-go"

type SYSTEM_LOGGING_RECORD struct {
   local_node string
   container  string
   file       string
   stream_driver  redis_handlers.Redis_Stream_Struct
   hash_driver    redis_handlers.Redis_Hash_Struct
}


func Construct_system_logging( local_node, container, file string )*SYSTEM_LOGGING_RECORD {

   var return_value SYSTEM_LOGGING_RECORD
   return_value.local_node = local_node
   return_value.container = container
   return_value.file = file
   
   search_list := []string{ "SYSTEM_MONITOR","SYSTEM_MONITOR" }
   handlers := data_handler.Construct_Data_Structures(&search_list)
   return_value.stream_driver = (*handlers)["SYSTEM_ALERTS"].(redis_handlers.Redis_Stream_Struct)
   return_value.hash_driver = (*handlers)["SYSTEM_VERBS"].(redis_handlers.Redis_Hash_Struct)
   (&return_value).log_error_message("Reboot","","")
   return &return_value

}   
  



func ( v SYSTEM_LOGGING_RECORD ) log_error_message( verb,subject,obj_of string){
    log_value := make(map[string]interface{})
	
	time_value :=  time.Now().UnixNano()
	(v).update_system_verb(verb,time_value)
    log_value["local_node"] = v.local_node
    log_value["file"] = v.file
    log_value["container"] = v.container
    log_value["time"] = time_value
	log_value["verb"] = verb
	log_value["subject"] = subject
	log_value["obj_of"] = obj_of
	(v).push_stream( &log_value )

}
        
func ( v SYSTEM_LOGGING_RECORD )update_system_verb(verb string ,time_value int64 ){

   var b bytes.Buffer	
   msgpack.Pack(&b,time_value)
   v.hash_driver.HSet(verb, b.String())


}	
   


func ( v SYSTEM_LOGGING_RECORD )push_stream(stream_data *map[string]interface{}){

  var b bytes.Buffer	
  msgpack.Pack(&b,(*stream_data))
  v.stream_driver.Xadd(b.String())

}