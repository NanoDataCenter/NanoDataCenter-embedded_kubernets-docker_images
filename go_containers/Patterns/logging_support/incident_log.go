
package logging_support
import  "fmt"
import  "time"
import "bytes"
import "lacima.com/redis_support/redis_handlers"
import   "lacima.com/redis_support/generate_handlers"
import	"github.com/msgpack/msgpack-go"

type Incident_Log_Type struct {

  time           redis_handlers.Redis_Single_Structure
  status         redis_handlers.Redis_Single_Structure
  current_state  redis_handlers.Redis_Single_Structure
  last_error     redis_handlers.Redis_Single_Structure
  error_log      redis_handlers.Redis_Stream_Struct

 
}


func Construct_incident_log( search_path []string ) *Incident_Log_Type{

   var return_value Incident_Log_Type
   fmt.Println("search_path",search_path)
   handlers := data_handler.Construct_Data_Structures(&search_path)
   return_value.time          = (*handlers)["TIME_STAMP"].(redis_handlers.Redis_Single_Structure)
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

  time_stamp := time.Now().UnixNano()
  var b1 bytes.Buffer	
  msgpack.Pack(&b1,time_stamp)
  v.time.Set(b1.String())
  //fmt.Println("new_value",status,time_stamp,new_value,current_error) 
  if status == false {
     //fmt.Println("status is false")
     old_error := v.last_error.Get()
     if current_error != old_error {
	    //fmt.Println("updating error status")
        v.last_error.Set(current_error)
	    v.error_log.Xadd(current_error)
      }
  } 
}
