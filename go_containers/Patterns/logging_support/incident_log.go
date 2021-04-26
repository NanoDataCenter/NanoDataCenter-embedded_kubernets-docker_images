
package logging_support

import "bytes"
import "lacima.com/redis_support/redis_handlers"
import   "lacima.com/redis_support/generate_handlers"
import	"github.com/msgpack/msgpack-go"

type Incident_Log_Type struct {

  status         redis_handlers.Redis_Single_Structure
  current_state  redis_handlers.Redis_Single_Structure
  last_error     redis_handlers.Redis_Single_Structure
  error_log      redis_handlers.Redis_Stream_Struct

 
}


func Construct_incident_log( search_path []string ) *Incident_Log_Type{

   var return_value Incident_Log_Type
   handlers := data_handler.Construct_Data_Structures(&search_path)
   return_value.status        = (*handlers)["STATUS"].(redis_handlers.Redis_Single_Structure)
   return_value.current_state = (*handlers)["CURRENT_STATE"].(redis_handlers.Redis_Single_Structure)
   return_value.last_error    = (*handlers)["LAST_ERROR"].(redis_handlers.Redis_Single_Structure)
   return_value.error_log     = (*handlers)["ERROR_LOG"].(redis_handlers.Redis_Stream_Struct)
   return &return_value


}

func (v *Incident_Log_Type)Log_data( status bool, new_value, current_error string ){
   
  var b bytes.Buffer	
  msgpack.Pack(&b,status)
  v.status.Set(b.String())
  v.current_state.Set(new_value)
  old_error := v.last_error.Get()
  if current_error != old_error {
      v.last_error.Set(current_error)
	  v.error_log.Xadd(current_error)
   }
}