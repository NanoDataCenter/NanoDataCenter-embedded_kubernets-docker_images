
package logging_support

import "fmt"
import "time"

import "lacima.com/redis_support/redis_handlers"
import   "lacima.com/redis_support/generate_handlers"
import "lacima.com/Patterns/msgpack_2"

type Watch_Dog_Log_Type struct {

  time_stamp   redis_handlers.Redis_Single_Structure

 
}


func Construct_watch_data_log( search_path []string ) *Watch_Dog_Log_Type{

   var return_value Watch_Dog_Log_Type
   fmt.Println("search_path",search_path)
   handlers := data_handler.Construct_Data_Structures(&search_path)
   fmt.Println("handlers",handlers)
   
   return_value.time_stamp          = (*handlers)["WATCH_DOG_TS"].(redis_handlers.Redis_Single_Structure)
  
   return &return_value


}

func (v *Watch_Dog_Log_Type)Strobe_Watch_Dog(  ){
   

   time_stamp           := time.Now().UnixNano()
   time_stamp_msg_pack  := msg_pack_utils.Pack_int64(time_stamp)
   v.time_stamp.Set(time_stamp_msg_pack)
  

   
}
