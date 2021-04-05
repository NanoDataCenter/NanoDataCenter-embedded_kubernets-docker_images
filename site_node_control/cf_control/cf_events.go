package cf

import "time"
//import "fmt"



func (system *CF_SYSTEM_TYPE)wait_for_event()*CF_EVENT_TYPE{

  
   var return_value CF_EVENT_TYPE
   if (system).event_queue.Len() == 0 {
    
     time.Sleep(time.Duration( (system).time_tick_duration))
	 return_value.Name = CF_TIME_TICK
	 return_value.Value = time.Now().UnixNano()
	 //fmt.Println("time_tick")
	 
   }else{
      var e = (system).event_queue.Front().Value
	  return_value = e.(CF_EVENT_TYPE)
      
   }
   //fmt.Println("emitted event",return_value)
   return &return_value

}

func (system *CF_SYSTEM_TYPE)Send_event( event_name string, value interface{} ){

   var event  CF_EVENT_TYPE
   event.Name = event_name
   event.Value = value

   (system).event_queue.PushBack(&event)
}



