package cf

import "time"
//import "fmt"


func (system *CF_SYSTEM_TYPE)start_time_tick(){

  (system).ticker = time.NewTicker((system).time_tick_duration)

}


func (system *CF_SYSTEM_TYPE)wait_for_event()*CF_EVENT_TYPE{

  
   var return_value CF_EVENT_TYPE
   var loop_flag = true
   for loop_flag {
      select {
            case return_value = <-(system).event_queue:
                 loop_flag = false             
            case <-(system).ticker.C:
                return_value.Name = CF_TIME_TICK
	            return_value.Value = time.Now().UnixNano()
				loop_flag = false
            }
      
   }
   //fmt.Println("emitted event",(system).row,(system).name,return_value)
   return &return_value

}

func (system *CF_SYSTEM_TYPE)Send_event( event_name string, value interface{} ){

   var event  CF_EVENT_TYPE
   event.Name = event_name
   event.Value = value

   (system).event_queue <- event
}

