
package logging_support
//import  "fmt"
import  "time"

import "lacima.com/redis_support/redis_handlers"
import  "lacima.com/redis_support/generate_handlers"
import "lacima.com/Patterns/msgpack_2"

type Incident_Log_Type struct {
  status         redis_handlers.Redis_Single_Structure
  time           redis_handlers.Redis_Single_Structure
  last_error     redis_handlers.Redis_Single_Structure
 
}
 
func Construct_incident_log( search_path []string ) *Incident_Log_Type{

   var return_value Incident_Log_Type
  
   handlers := data_handler.Construct_Data_Structures(&search_path)
   return_value.status        = (*handlers)["STATUS"].(redis_handlers.Redis_Single_Structure)
   return_value.time          = (*handlers)["TIME_STAMP"].(redis_handlers.Redis_Single_Structure)
   return_value.last_error    = (*handlers)["LAST_ERROR"].(redis_handlers.Redis_Single_Structure)
   
   return_value.status.Set(msg_pack_utils.Pack_bool(true))
   return_value.time.Set(msg_pack_utils.Pack_int64(0))
   return_value.last_error.Set(msg_pack_utils.Pack_string(""))
   
   return &return_value


}

func (v *Incident_Log_Type)Log_data(  current_error string ){
   
      time_stamp           := time.Now().UnixNano()
      time_stamp_msg_pack  := msg_pack_utils.Pack_int64(time_stamp)
      v.time.Set(time_stamp_msg_pack)
      _,err := msg_pack_utils.Unpack_string(current_error)
      if err == false {
          current_error = msg_pack_utils.Pack_string(current_error)
      }
      v.last_error.Set(current_error)
      v.status.Set(msg_pack_utils.Pack_bool(false))

      

}



func (v *Incident_Log_Type)Post_event( current_error string){
   v.Log_data(current_error)
}
 
