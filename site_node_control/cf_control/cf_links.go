package cf

//import "fmt"
func (cf_system CF_SYSTEM) cf_initialize_chain( chain_data  *CF_CHAIN){
     (*chain_data).initialized = true
	 
	 for _, link := range (*chain_data).links {
	       (cf_system).cf_initialize_link( link )
     }		   
}

func (cf_system CF_SYSTEM) cf_execute_chain( chain_data  *CF_CHAIN, event_data *map[string]interface{}){
     
	 var return_value int
	 var count int
	 count = 0
	 for _, link := range (*chain_data).links {
	      if link.active == true {
		  count = count +1
	       return_value = (cf_system).cf_execute_link(chain_data, link, event_data )
		   if (cf_system).analyize_return_code( chain_data, link,return_value) == false{
		       break; // chain processing has halted
		   }
		 }     
     }	
	 if count == 0 {
	   (cf_system).cf_disable_chain(chain_data) // zomby chain
	 }
}

func (cf_system CF_SYSTEM) cf_initialize_link( link_data  *CF_LINK){
  (*link_data).active = true
  (*link_data).initialized = false

}
func (cf_system CF_SYSTEM) cf_disable_link( link_data  *CF_LINK){
  (*link_data).active = false
  (*link_data).initialized = false

}

func (cf_system CF_SYSTEM) cf_execute_link(chain_data *CF_CHAIN, link_data  *CF_LINK, event_data *map[string]interface{})int {
  
    if (*link_data).initialized == false{
	    cf_system.op_code_map[(*link_data).opcode_type](cf_system,chain_data,(*link_data).parameters, &cf_system.init_event ) 
		(*link_data).initialized = true
     }
	 return cf_system.op_code_map[(*link_data).opcode_type](cf_system,chain_data,(*link_data).parameters,  event_data)


}

func (cf_system  CF_SYSTEM) analyize_return_code( chain_data *  CF_CHAIN, link_data  *CF_LINK, return_code int)bool {

  //fmt.Println("return_code",return_code)
  var return_value,ok = (cf_system).continue_map[return_code]
  if ok == false {
    panic("bad return code")
  }
  
 
  if return_code == CF_TERMINATE {
     (cf_system).cf_disable_chain(chain_data)
  }
  if return_code == CF_RESET{
     (cf_system).cf_enable_chain(chain_data)
  }
  if return_code == CF_DISABLE{
     (cf_system).cf_disable_link(link_data)
  }
  return return_value  


}

func (cf_system CF_SYSTEM)  cf_enable_chain( chain_data *CF_CHAIN ){
   (*chain_data).active = true
   (cf_system).cf_initialize_chain(chain_data)

}

func (cf_system CF_SYSTEM) cf_disable_chain( chain_data *CF_CHAIN ){
   (*chain_data).active = true
   (cf_system).cf_initialize_chain(chain_data)

}

func (cf_system CF_SYSTEM) CF_enable_chains( chain_list []string ){

   for _,chain_name := range chain_list{
       chain_data,ok := (cf_system).chain_map[chain_name]
	   if ok == false{
	      panic("bad chain name")
	   }
	   (cf_system).cf_enable_chain(chain_data)
    }
}

func (cf_system CF_SYSTEM) CF_disable_chains( chain_list []string ){

   for _,chain_name := range chain_list{
       chain_data,ok := (cf_system).chain_map[chain_name]
	   if ok == false{
	      panic("bad chain name")
		}
	   (cf_system).cf_disable_chain(chain_data)
    }
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