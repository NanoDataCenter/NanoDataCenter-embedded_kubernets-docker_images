package cf
import "container/list"

type CF_RETURN int

const (
    CF_HALT = iota
    CF_CONTINUE = iota
    CF_DISABLE = iota
    CF_RESET = iota
    CF_TERMINATE = iota
	
)

var CF_TIME_TICK string = "CF_TIME_TICK"
var CF_INIT = "INIT"


type EVENT_TYPE struct{
  event int
  data *map[string]interface{}
}


type aux_function_type func( handle interface{},parameters map[string]interface{})int

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

type CF_SYSTEM struct {

  op_code_map  map[string] CF_Function_type
  chain_map map[string]*CF_CHAIN
  chain_order []*CF_CHAIN
  current_chain *CF_CHAIN
  current_link  *CF_LINK
  current_event *map[string]interface{}
  event_queue *list.List
  init_event  map[string]interface{}
}








func Init( system *CF_SYSTEM ){

 
  (*system).chain_map = make(map[string]*CF_CHAIN)
  (*system).event_queue = list.New()
  (*system).init_event = make(map[string]interface{})
  (*system).init_event["event_name"] = CF_INIT
  (*system).init_event["value"] = nil  // value is not used
  cf_initialize_opcodes(system)
  
}

func Add_Chain( system *CF_SYSTEM, chain_name string, state bool ){

   var temp CF_CHAIN
   temp.name = chain_name
   temp.active = state
   temp.initialized = false
   
   (*system).chain_map[chain_name] = &temp
   (*system).chain_order = append((*system).chain_order,&temp)
   (*system).current_chain = &temp
   (*system).current_link = nil
   


}

func (cf_system CF_SYSTEM)Execute(){

  (cf_system).Initialize_Chains()
  var loop_flag = true
  for loop_flag {
     var event = (cf_system).wait_for_event()
     for _ , chain_data := range (cf_system).chain_order{
	    if (*chain_data).active == true {
           (cf_system).cf_execute_chain( chain_data,event)
		}
	 }
  }
}

func (cf_system CF_SYSTEM)Initialize_Chains(){		
    for _ , chain_data := range (cf_system).chain_order{
	   if (*chain_data).active == true {
        (cf_system).cf_initialize_chain( chain_data)
	   }
	 }

}


