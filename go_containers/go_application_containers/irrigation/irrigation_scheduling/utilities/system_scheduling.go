package scheduling_utilities

import "fmt"


func (v *Scheduling_Type)check_for_system_activity_recovery(){

   v.system_error_logger()
}

func (v *Scheduling_Type)check_for_system_activity(){

    dow := get_dow()
	hour_minute  := get_hour_minute()
    fmt.Println("dow",dow)
	fmt.Println("hour_minute",hour_minute)
    for _,element := range v.system_control.scheduling_array{
	    defer v.check_for_system_activity_recovery()
		v.check_for_system_activity_loop_element(element)
	}
}
	
func (v *Scheduling_Type)check_for_system_activity_loop_element( element map[string]interface{}){

     name    := element["name"].(string)

	 if v.check_system_resumit_flag(name) == false {
	    return 
	 }
     if check_for_proper_date(element) == true{
	    fmt.Println("proper date")
	    start_time := generate_comparison_time( element["start_time"] )
		end_time   := generate_comparison_time( element["end_time"] )
		if determine_start_time( start_time,end_time ) == true {
	        v.iq.Queue_Command(name)
			v.set_system_resumit_flag(name)
		}
	 }else{
	    fmt.Println(" false proper date")
	 
	 }
	 
}


func (v *Scheduling_Type)check_system_resumit_flag( name string )bool {

   data := (v.system_control).completion_hash.HGet(name)
   fmt.Println("data",data)
   if data != "true" {
       
       return true
   }
   fmt.Println("schedule has already been queued")
   return false
   


}
          
func (v *Scheduling_Type)set_system_resumit_flag( name string ) {

   (v.system_control).completion_hash.HSet(name,"true")
 

}
    
