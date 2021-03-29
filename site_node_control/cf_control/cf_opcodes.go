package cf


import "fmt"



type CF_Function_type func( system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{} )int







func cf_initialize_opcodes(system *CF_SYSTEM){
  (*system).op_code_map = make(map[string] CF_Function_type)
  (*system).op_code_map["Log"] =  cf_op_log_message
  (*system).op_code_map["Reset"] =  cf_op_reset
  (*system).op_code_map["Terminate"] =  cf_op_termination
  (*system).op_code_map["Wait_Interval"] =  cf_op_wait_interval
  

 

}

func Cf_add_log_link(system *CF_SYSTEM, log_message string){

   var temp CF_LINK

   temp.initialized = false
   temp.active = false
   temp.parameters = make(map[string]interface{})
   temp.parameters["log_messge"] = log_message
   temp.opcode_type = "Log"
   temp.aux_function = nil // no opcode
   
   var chain *CF_CHAIN
   chain =  (*system).current_chain
   (*chain).links = append((*chain).links, &temp)   
   



}


func cf_op_log_message(system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{})int{

   if (*event)["event_name"].(string) == CF_INIT {
      fmt.Println(parameters["log_messge"].(string))
   }
   return CF_DISABLE
}


func Cf_add_reset(system *CF_SYSTEM ){

   var temp CF_LINK

   temp.initialized = false
   temp.active = false
   temp.parameters = make(map[string]interface{})
   temp.opcode_type = "Reset"
   temp.aux_function = nil // no opcode
   
   var chain *CF_CHAIN
   chain =  (*system).current_chain
   (*chain).links = append((*chain).links, &temp)   
   
}

func cf_op_reset(system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{})int{

   return CF_RESET

}

func Cf_add_terminate(system *CF_SYSTEM ){

   var temp CF_LINK

   temp.initialized = false
   temp.active = false
   temp.parameters = make(map[string]interface{})
   temp.opcode_type = "Terminate"
   temp.aux_function = nil // no opcode
   
   var chain *CF_CHAIN
   chain =  (*system).current_chain
   (*chain).links = append((*chain).links, &temp)   
}



func cf_op_termination(system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{})int{

   return CF_TERMINATE

}




func Cf_add_wait_interval(system *CF_SYSTEM, delta_duration int64 ){

   var temp CF_LINK

   temp.initialized = false
   temp.active = false
   temp.parameters = make(map[string]interface{})
   temp.parameters["delta_time"] = delta_duration
   temp.opcode_type = "Wait_Interval"
   temp.aux_function = nil // no opcode
   
   var chain *CF_CHAIN
   chain =  (*system).current_chain
   (*chain).links = append((*chain).links, &temp)   
   
}

func cf_op_wait_interval(system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{})int{

   var return_code int = CF_HALT

   if (*event)["event_name"].(string) == CF_INIT {
      parameters["ref_time"] = parameters["ref_time"].(int64) + (*event)["value"].(int64)
	  
   }
   
   if (*event)["event_name"].(string) == CF_TIME_TICK {
      if  (*event)["value"].(int64) >= parameters["ref_time"].(int64){
	    return_code = CF_DISABLE
	  }
   }else{
     ;
   }
   
   

  return return_code
}
