package scheduling_utilities

import "fmt"

func (v *system_scheduling_type)clear_done_recovery(){

   // fill in later

}

func (v *system_scheduling_type)clear_done_flag(){
   for _, element := range v.scheduling_array{
      defer v.clear_done_recovery()
	  v.clear_done_element(element)
	}
}

func (v *system_scheduling_type)clear_done_element( element map[string]interface{} ){

    name := element["name"].(string)
	start_time := generate_comparison_time( element["start_time"] )
	end_time   := generate_comparison_time( element["end_time"] )
	if  determine_start_time( start_time,end_time) == false {
	  fmt.Println("clear event")
	  v.completion_hash.HSet(name,"false")
	}

}
