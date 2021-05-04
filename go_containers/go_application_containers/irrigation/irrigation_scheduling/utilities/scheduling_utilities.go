package scheduling_utilities

import      "time"
//import "os"
import      "fmt"
import      "strconv"


func generate_comparison_time( input interface{} )[]int{

  return_value := make([]int,0)  
  for _, item := range input.([]interface{}) {
     return_value = append(return_value,int(item.(float64)))
  }
  fmt.Println("return_value",return_value)
  return return_value

}


func determine_start_time( start_time []int,end_time []int) bool {

       return_value := false

 
       st_array := get_hour_minute()
       if match_time( start_time,end_time ) == true {
	           if ( match_time( start_time, st_array) && 
	                match_time( st_array, end_time )) == true{
					 return_value = true
				}
       }else{
	      // this is a wrap around case
          if   match_time( start_time,st_array) == true {
               return_value = true
		  }
          if match_time(st_array,end_time) == true {
                return_value = true
		  }
    
	   }
	   fmt.Println("return_value",return_value)
       return return_value
}


func  match_time( compare, value []int )bool{
     return_value := false
     if compare[0] < value[0]{
        return true
	 }
     if (compare[0] ==  value[0]) && ( compare[1] <= value[1] ){
       return true
	 }
     return return_value
}


func check_for_proper_date( element map[string]interface{} ) bool {

   enable,ok := element["schedule_enable"]
   if ok == true {
      if enable.(bool) == false {
	    return false
	  }
   
   }
   day_flag,ok1 := element["day_flag"]
   if ok1 == true {
      value,err := strconv.Atoi(day_flag.(string))
	  if err != nil{
	     panic("bad day flag")
	  }
      if value != 0 {
	    return schedule_doy(element)
	  }
   
   }
   return schedule_dow(element )

}

  
func schedule_dow( element map[string]interface{} )bool{
   p_dow_slice := element["dow"].([]interface{})
   dow_slice := make([]int,0)
   for _,item := range p_dow_slice{
     dow_slice = append(dow_slice,int(item.(float64)))
   
   }
   
   fmt.Println("dow_slice",dow_slice)
   dow := get_dow()
   if dow_slice[dow] == 0 {
      return false
   }else{
      return true
   }
}  

func schedule_doy(element map[string]interface{}) bool {
      doy     := get_doy()
      divisor,err:= strconv.Atoi(element["day_div"].(string))
      modulus,err1 := strconv.Atoi(element["day_mod"].(string))
	  if err != nil {
	    panic("bad day div")
	  }
	  if err1 != nil{
	    panic("bad day mod")
	  }
      result  := doy % divisor
     
      if result == modulus {
        return true
	  }
	  return false

}
  


func get_dow() int {

  return int(time.Now().Weekday())

}

func get_hour_minute() []int{
 
  t := time.Now()
  return []int{int(t.Hour()),int(t.Second())}   
   

}//dow_array := []int{ 1,2,3,4,5,6,0}



func get_doy() int {

  return int(time.Now().YearDay())

}