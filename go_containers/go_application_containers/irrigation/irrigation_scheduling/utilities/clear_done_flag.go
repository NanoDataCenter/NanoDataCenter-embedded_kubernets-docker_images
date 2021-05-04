package scheduling_utilities

import "fmt"

func (v *system_scheduling_type)system_clear_done_recovery(){
    v.error_logger()

}

func (v *system_scheduling_type)system_clear_done_flag(){
   for _, element := range v.scheduling_array{
      defer v.system_clear_done_recovery()
	  v.system_clear_done_element(element)
	}
}

func (v *system_scheduling_type)system_clear_done_element( element map[string]interface{} ){

    name := element["name"].(string)
	start_time := generate_comparison_time( element["start_time"] )
	end_time   := generate_comparison_time( element["end_time"] )
	if  determine_start_time( start_time,end_time) == false {
	  fmt.Println("clear event")
	  v.completion_hash.HSet(name,"false")
	}

}



func (v *system_scheduling_type)irrigaton_done_recovery(){

  v.error_logger()

}

func (v *system_scheduling_type)irrigation_clear_done_flag(){
   for _, element := range v.scheduling_array{
      defer v.irrigaton_done_recovery()
	  v.irrigation_clear_done_element(element)
	}
}

func (v *system_scheduling_type)irrigation_clear_done_element( element map[string]interface{} ){

    name := element["name"].(string)
	start_time := generate_comparison_time( element["start_time"] )
	end_time   := generate_comparison_time( element["end_time"] )
	if  determine_start_time( start_time,end_time) == false {
	  fmt.Println("clear event")
	  v.completion_hash.HSet(name,"false")
	}

}
