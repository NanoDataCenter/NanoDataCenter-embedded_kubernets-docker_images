
package logging_support

import "fmt"
import "time"
import "bytes"
import "lacima.com/redis_support/redis_handlers"
import   "lacima.com/redis_support/generate_handlers"
import	"github.com/msgpack/msgpack-go"

type Watch_Dog_Log_Type struct {

  watch_dog_time_stamp   redis_handlers.Redis_Single_Structure
  watch_dog_state        redis_handlers.Redis_Single_Structure
 
}


func Construct_watch_data_log( search_path []string ) *Watch_Dog_Log_Type{

   var return_value Watch_Dog_Log_Type
   fmt.Println("search_path",search_path)
   handlers := data_handler.Construct_Data_Structures(&search_path)
   fmt.Println("handlers",handlers)
   
   return_value.watch_dog_time_stamp          = (*handlers)["WATCH_DOG_TS"].(redis_handlers.Redis_Single_Structure)
   return_value.watch_dog_state          = (*handlers)["WATCH_DOG_STATE"].(redis_handlers.Redis_Single_Structure)
   return &return_value


}

func (v *Watch_Dog_Log_Type)Strobe_Watch_Dog(  ){
   

  time_stamp := time.Now().UnixNano()
  //fmt.Println("watch_dog log",time_stamp)
  var b bytes.Buffer	
  msgpack.Pack(&b,time_stamp)
  v.watch_dog_time_stamp.Set(b.String())
  

   
}
