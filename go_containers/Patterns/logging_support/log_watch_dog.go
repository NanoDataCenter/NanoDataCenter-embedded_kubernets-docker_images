
package logging_support

//import "fmt"
import "time"
import "bytes"
import "lacima.com/redis_support/redis_handlers"
import   "lacima.com/redis_support/generate_handlers"
import	"github.com/msgpack/msgpack-go"

type Watch_Dog_Log_Type struct {

  watch_dog   redis_handlers.Redis_Single_Structure
  
 
}


func Construct_watch_data_log( search_path []string ) *Watch_Dog_Log_Type{

   var return_value Watch_Dog_Log_Type
   handlers := data_handler.Construct_Data_Structures(&search_path)
   return_value.watch_dog          = (*handlers)["WATCH_DOG"].(redis_handlers.Redis_Single_Structure)
   return &return_value


}

func (v *Watch_Dog_Log_Type)Strobe_Watch_Dog(  ){
   

  time_stamp := time.Now().UnixNano()
  //fmt.Println("watch_dog log",time_stamp)
  var b bytes.Buffer	
  msgpack.Pack(&b,time_stamp)
  v.watch_dog.Set(b.String())
  

   
}