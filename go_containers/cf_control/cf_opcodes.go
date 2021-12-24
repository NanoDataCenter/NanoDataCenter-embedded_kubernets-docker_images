package cf


import "fmt"
import "time"



type CF_Function_type func( system interface{},chain interface{}, parameters map[string]interface{}, event *CF_EVENT_TYPE )int







func (system *CF_SYSTEM_TYPE) cf_initialize_opcodes(){
  (system).op_code_map = make(map[string] CF_Function_type)
  (system).add_opcode("Log",cf_op_log_message)
  (system).add_opcode("Reset", cf_op_reset )
  (system).add_opcode("Terminate", cf_op_termination )
  (system).add_opcode("Wait_Interval", cf_op_wait_interval)
  (system).add_opcode("One_Step", cf_op_one_step)
  (system).add_opcode("Enable_Chains",cf_op_enable_chains)
  (system).add_opcode("Disable_Chains",cf_op_disable_chains) 
  (system).add_opcode("Unfiltered_Element",cf_op_unfiltered_element)
  (system).add_opcode("Wait_hour_minute_le",cf_wait_hour_minute_le)
  (system).add_opcode("Wait_hour_minute_ge",cf_wait_hour_minute_ge)

}


func (system *CF_SYSTEM_TYPE) add_opcode( op_code string, function CF_Function_type){

  _ , err := (system).op_code_map[op_code]
  if err == true {
    panic("duplicate_opcode")
  }
  (system).op_code_map[op_code] = function


}


 
func cf_op_unfiltered_element( system interface{},chain interface{}, parameters map[string]interface{}, event *CF_EVENT_TYPE)int{

   var helper_function = parameters["__helper_function__"].(CF_helper_function)
   if (*event).Name == CF_INIT_EVENT {
      
     
	  return helper_function(system,chain , parameters, event)
	  
	  
   } else{
    
	 return helper_function(system,chain , parameters, event)
   }
}  

func cf_op_enable_chains(system interface{},chain interface{}, parameters map[string]interface{}, event *CF_EVENT_TYPE)int{

  if (*event).Name == CF_INIT_EVENT {
      var system = system.(*CF_SYSTEM_TYPE)
      var chain_list = parameters["chains"].([]string)
      (system).CF_enable_chains(chain_list)
   }
   return CF_DISABLE
}

func cf_op_disable_chains(system interface{},chain interface{}, parameters map[string]interface{}, event *CF_EVENT_TYPE)int{


  if (*event).Name == CF_INIT_EVENT {
      var system = system.(*CF_SYSTEM_TYPE)
      var chain_list = parameters["chains"].([]string)
      (system).CF_disable_chains(chain_list)
   }
   return CF_DISABLE
}

func cf_op_log_message(system interface{},chain interface{}, parameters map[string]interface{}, event *CF_EVENT_TYPE)int{

   
   if (*event).Name == CF_INIT_EVENT {
     var chain_name = parameters["chain_name"].(string)
	 var system_name = parameters["system_name"].(string)
	 var output = "system:  "+ system_name + "     chain:  "+chain_name +"  msg:  "+parameters["log_messge"].(string)
     fmt.Println(output)
   }
   return CF_DISABLE
}



func cf_op_reset(system interface{},chain interface{}, parameters map[string]interface{}, event *CF_EVENT_TYPE)int{

   return CF_RESET

}




func cf_op_termination(system interface{},chain interface{}, parameters map[string]interface{}, event *CF_EVENT_TYPE)int{

   return CF_TERMINATE

}






func cf_op_wait_interval(system interface{},chain interface{}, parameters map[string]interface{}, event *CF_EVENT_TYPE)int{

   var return_code int = CF_HALT
  

   if (*event).Name == CF_INIT_EVENT {
      //fmt.Println("only once")
      parameters["__count__"] = parameters["ref_time"].(int64) + time.Now().UnixNano()
	  //fmt.Println("parameters",parameters)
	  
   }
   
   if (*event).Name == CF_TIME_TICK {
      
      if  (*event).Value.(int64) >= parameters["__count__"].(int64){
	    
	    return_code = CF_DISABLE
	  }
   }else{
     ;
   }
   
   

  return return_code
}



func cf_op_one_step( system interface{},chain interface{}, parameters map[string]interface{}, event *CF_EVENT_TYPE)int{

   if (*event).Name == CF_INIT_EVENT {
      //fmt.Println("made it here")
      var helper_function = parameters["__helper_function__"].(CF_helper_function)
	  helper_function(system,chain , parameters, event)
	  
	  
   }
   return CF_DISABLE

}



func cf_wait_hour_minute_le( system interface{},chain interface{}, parameters map[string]interface{}, event *CF_EVENT_TYPE)int{
    if (*event).Name == CF_TIME_TICK {
       ref_hour           := parameters["hour"].(int)
       ref_minute         := parameters["minute"].(int)
       hour,minute,_      := time.Now().Clock() 
       if  hour < ref_hour {
           return CF_DISABLE
       }
       if hour == ref_hour {
           if minute <= ref_minute {
               return CF_DISABLE
           }
       }
    }
    return CF_HALT
}    
       
 func cf_wait_hour_minute_ge( system interface{},chain interface{}, parameters map[string]interface{}, event *CF_EVENT_TYPE)int{
    if (*event).Name == CF_TIME_TICK {
       ref_hour           := parameters["hour"].(int)
       ref_minute         := parameters["minute"].(int)
       hour,minute,_      := time.Now().Clock() 
       if  hour > ref_hour {
           return CF_DISABLE
       }
       if hour == ref_hour {
           if minute >= ref_minute {
               return CF_DISABLE
           }
       }
    }
    return CF_HALT
}         
        

    




