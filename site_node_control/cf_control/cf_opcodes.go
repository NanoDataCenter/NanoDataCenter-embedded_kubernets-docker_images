package cf


import "fmt"
import "time"



type CF_Function_type func( system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{} )int







func (system *CF_SYSTEM) cf_initialize_opcodes(){
  (system).op_code_map = make(map[string] CF_Function_type)
  (system).op_code_map["Log"] =  cf_op_log_message
  (system).op_code_map["Reset"] =  cf_op_reset
  (system).op_code_map["Terminate"] =  cf_op_termination
  (system).op_code_map["Wait_Interval"] =  cf_op_wait_interval
  (system).op_code_map["One_Step"] =  cf_op_one_step

 

}

func (system *CF_SYSTEM) Cf_add_log_link( log_message string){

   var temp CF_LINK
   var chain *CF_CHAIN
   chain =  (system).current_chain
   temp.initialized = false
   temp.active = false
   temp.parameters = make(map[string]interface{})
   temp.parameters["log_messge"] = log_message
   temp.parameters["system_name"] = system.name
   temp.parameters["chain_name"] = chain.name
   temp.opcode_type = "Log"
   
   

   (*chain).links = append((*chain).links, &temp)   
   



}


func cf_op_log_message(system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{})int{

   
   if (*event)["event_name"].(string) == CF_INIT {
     var chain_name = parameters["chain_name"].(string)
	 var system_name = parameters["system_name"].(string)
	 var output = "system:  "+ system_name + "     chain:  "+chain_name +"  msg:  "+parameters["log_messge"].(string)
     fmt.Println(output)
   }
   return CF_DISABLE
}

func (system *CF_SYSTEM) Cf_add_reset( ){

   var temp CF_LINK

   temp.initialized = false
   temp.active = false
   temp.parameters = make(map[string]interface{})
   temp.opcode_type = "Reset"
 
   
   var chain *CF_CHAIN
   chain =  (system).current_chain
   (*chain).links = append((*chain).links, &temp)   
   
}

func cf_op_reset(system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{})int{

   return CF_RESET

}

func (system *CF_SYSTEM) Cf_add_terminate( ){

   var temp CF_LINK

   temp.initialized = false
   temp.active = false
   temp.parameters = make(map[string]interface{})
   temp.opcode_type = "Terminate"
   
   
   var chain *CF_CHAIN
   chain =  (system).current_chain
   (*chain).links = append((*chain).links, &temp)   
}



func cf_op_termination(system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{})int{

   return CF_TERMINATE

}




func (system *CF_SYSTEM) Cf_add_wait_interval( delta_duration int64 ){

   var temp CF_LINK

   temp.initialized = false
   temp.active = false
   temp.parameters = make(map[string]interface{})
   temp.parameters["ref_time"] = delta_duration
   temp.opcode_type = "Wait_Interval"
   
   
   var chain *CF_CHAIN
   chain =  (system).current_chain
   (*chain).links = append((*chain).links, &temp)   
   
}

func cf_op_wait_interval(system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{})int{

   var return_code int = CF_HALT
  

   if (*event)["event_name"].(string) == CF_INIT {
      //fmt.Println("only once")
      parameters["__count__"] = parameters["ref_time"].(int64) + time.Now().UnixNano()
	  //fmt.Println("parameters",parameters)
	  
   }
   
   if (*event)["event_name"].(string) == CF_TIME_TICK {
      
      if  (*event)["value"].(int64) >= parameters["__count__"].(int64){
	    
	    return_code = CF_DISABLE
	  }
   }else{
     ;
   }
   
   

  return return_code
}

type CF_helper_function func( system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{})int

func (system *CF_SYSTEM) Cf_add_one_step(  helper_function CF_helper_function, parameters map[string]interface{}){

   var temp CF_LINK

   temp.initialized = false
   temp.active = false
   temp.parameters = parameters
   temp.parameters["__helper_function__"] = helper_function
   temp.opcode_type = "One_Step"
   
   
   var chain *CF_CHAIN
   chain =  (system).current_chain
   (*chain).links = append((*chain).links, &temp)   
   
}


func cf_op_one_step( system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{})int{


   
  

   if (*event)["event_name"].(string) == CF_INIT {
      
      var helper_function = parameters["__helper_function__"].(CF_helper_function)
	  helper_function(system,chain , parameters, event)
	  
	  
   }
   return CF_DISABLE

}


