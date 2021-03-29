package cf

func (cf_system CF_SYSTEM) cf_initialize_chain( chain_data  *CF_CHAIN){
     (*chain_data).initialized = true
	 
	 for _, link := range (*chain_data).links {
	       (cf_system).cf_initialize_link( link )
     }		   
}

func (cf_system CF_SYSTEM) cf_execute_chain( chain_data  *CF_CHAIN, event_data *map[string]interface{}){
     
	 var return_value int
	 for _, link := range (*chain_data).links {
	       return_value = (cf_system).cf_execute_link(chain_data, link, event_data )
		   if (cf_system).analyize_return_code( chain_data, link,return_value){
		       break; // chain processing has halted
		   }
		      
     }	
}

func (cf_system CF_SYSTEM) cf_initialize_link( link_data  *CF_LINK){
  (*link_data).active = true
  (*link_data).initialized = false

}


func (cf_system CF_SYSTEM) cf_execute_link(chain_data *CF_CHAIN, link_data  *CF_LINK, event_data *map[string]interface{})int {
  
    if (*link_data).initialized == false{
	    cf_system.op_code_map[(*link_data).opcode_type](cf_system,chain_data,(*link_data).parameters, &cf_system.init_event ) 
		(*link_data).initialized = true
     }
	 return cf_system.op_code_map[(*link_data).opcode_type](cf_system,chain_data,(*link_data).parameters,  event_data)


}

func (cf_system  CF_SYSTEM) analyize_return_code( chain_data *  CF_CHAIN, link_data  *CF_LINK, return_value int)bool {

   return true  // place holder


}



func (cf_system CF_SYSTEM) CF_enable_chain( chain_list []string ){
;
}

func (cf_system CF_SYSTEM)CF_disable_chain( chain_list []string ){
;
}


/*


type CF_LINK struct {

  initialized         bool
  active             bool
  parameters         map[string]interface{}
  opcode_type         string
  aux_function aux_function_type 
}

type CF_CHAIN struct {
  name string
  active bool
  initialized bool
  links []*CF_LINK

}




*/