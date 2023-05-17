package cf
import "time"

func (system *CF_SYSTEM_TYPE) CF_add_generic_link( link_data *CF_LINK_TYPE){

  var chain *CF_CHAIN_TYPE
  chain =  (system).current_chain
  (*chain).links = append((*chain).links, link_data)   
  
}


func (system *CF_SYSTEM_TYPE) Cf_add_log_link( log_message string){

   var temp CF_LINK_TYPE
   temp.initialized = false
   temp.active = false
   temp.parameters = make(map[string]interface{})
   temp.parameters["log_messge"] = log_message
   temp.parameters["system_name"] = (system).name
   temp.parameters["chain_name"] = (system).current_chain.name
   temp.opcode_type = "Log"
   (system).CF_add_generic_link( &temp)
  
}


func (system *CF_SYSTEM_TYPE) Cf_add_reset( ){

   var temp CF_LINK_TYPE
   temp.initialized = false
   temp.active = false
   temp.parameters = make(map[string]interface{})
   temp.opcode_type = "Reset"

   (system).CF_add_generic_link( &temp)  
   
}

func (system *CF_SYSTEM_TYPE) Cf_add_terminate( ){

   var temp CF_LINK_TYPE

   temp.initialized = false
   temp.active = false
   temp.parameters = make(map[string]interface{})
    temp.opcode_type = "Terminate"
  

   (system).CF_add_generic_link( &temp) 
}
func (system *CF_SYSTEM_TYPE) Cf_add_wait_interval( delta_duration time.Duration ){

   var temp CF_LINK_TYPE

   temp.initialized = false
   temp.active = false
   temp.parameters = make(map[string]interface{})
   temp.parameters["ref_time"] = int64(delta_duration)
   temp.opcode_type = "Wait_Interval"

   (system).CF_add_generic_link( &temp)
   
}
type CF_helper_function func( system interface{},chain interface{}, parameters map[string]interface{}, event *CF_EVENT_TYPE)int

func (system *CF_SYSTEM_TYPE) Cf_add_one_step(  helper_function CF_helper_function, parameters map[string]interface{}){

   var temp CF_LINK_TYPE

   temp.initialized = false
   temp.active = false
   temp.parameters = parameters
   temp.parameters["__helper_function__"] = helper_function
   temp.opcode_type = "One_Step"
   
   

   (system).CF_add_generic_link( &temp)
   
}


func (system *CF_SYSTEM_TYPE) Cf_add_enable_chains_links( chains []string){

   var temp CF_LINK_TYPE

   temp.initialized = false
   temp.active = false
   temp.parameters = make(map[string]interface{})
   temp.parameters["chains"] =chains
   temp.opcode_type = "Enable_Chains"
   
   

   (system).CF_add_generic_link( &temp)
   
}

func (system *CF_SYSTEM_TYPE) Cf_add_disable_chains_links( chains []string){

   var temp CF_LINK_TYPE

   temp.initialized = false
   temp.active = false
   temp.parameters = make(map[string]interface{})
   temp.parameters["chains"] =chains
   temp.opcode_type = "Disable_Chains"
   
   

   (system).CF_add_generic_link( &temp)
   
}

func (system *CF_SYSTEM_TYPE) Cf_add_unfiltered_element(helper_function CF_helper_function, parameters map[string]interface{}){

   var temp CF_LINK_TYPE

   temp.initialized = false
   temp.active = false
   temp.parameters = parameters
   temp.parameters["__helper_function__"] = helper_function
   temp.opcode_type = "Unfiltered_Element"
   
   (system).CF_add_generic_link( &temp)

} 

 
func (system *CF_SYSTEM_TYPE)Cf_wait_hour_minute_le( hour, minute int){
    
    var temp       CF_LINK_TYPE
    temp.initialized = false
    temp.active = false
    temp.parameters = make(map[string]interface{})
    temp.parameters["hour"]    = hour
    temp.parameters["minute"]  = minute
    temp.opcode_type = "Wait_hour_minute_le"
    (system).CF_add_generic_link( &temp)
    
}

func (system *CF_SYSTEM_TYPE)Cf_wait_hour_minute_ge( hour, minute int){
    
    var temp       CF_LINK_TYPE
    temp.initialized = false
    temp.active = false
    temp.parameters = make(map[string]interface{})
    temp.parameters["hour"]    = hour
    temp.parameters["minute"]  = minute
    temp.opcode_type = "Wait_hour_minute_ge"
    (system).CF_add_generic_link( &temp)
    
}
