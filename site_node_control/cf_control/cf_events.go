package cf

import "time"
//import "fmt"



func (system *CF_SYSTEM)wait_for_event()*map[string]interface{}{

  
   var return_value = make(map[string]interface{})
   if (*system).event_queue.Len() == 0 {
    
     time.Sleep(time.Second)
	 return_value["event_name"] = CF_TIME_TICK
	 return_value["value"] = time.Now().UnixNano()
	 //fmt.Println("time_tick")
	 
   }else{
      var e = (*system).event_queue.Front().Value
	  var temp = e.(map[string]interface{})
      return_value["event_name"] = temp["event_name"]
	  return_value["value"] = temp["value"]
   }
   return &return_value

}

func (system *CF_SYSTEM)Send_event( event_name string, value interface{} ){

   var event = make(map[string]interface{})
   event["evemt_name"] = event_name
   event["value"] = value

   (*system).event_queue.PushBack(&event)
}



