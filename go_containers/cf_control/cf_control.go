package cf

/*
 new suggested manager
 
 add erlang suppervision trees
 
 add dynmicall created thread  similar to erlang process



*/
import "time"
//import "container/list"
//import "fmt"


const CF_HALT       int = 0
const CF_CONTINUE   int = 1
const CF_DISABLE    int = 2
const CF_RESET      int = 3
const CF_TERMINATE  int= 4

	


const  CF_TIME_TICK    string = "CF_TIME_TICK"
const  CF_INIT_EVENT  string = "CF_INIT_TICK"
const  CF_START_EVENT string = "CF_START_TICK"
const  CF_TERMINATE_EVENT string = "CF_TERMINATE_EVENT"
const  CF_SYSTEM_STOP  string = "CF_SYSTEM_STOP"
const  CF_SYSTEM_RESET string = "CF_SYSTEM_RESET"

type aux_function_type func( handle interface{},parameters map[string]interface{})int

type CF_LINK_TYPE struct {

  initialized         bool
  active             bool
  parameters         map[string]interface{}
  opcode_type         string
 
}

type CF_CHAIN_TYPE struct {
  name string
  active bool
  initialized bool
  links []*CF_LINK_TYPE

}

type CF_SYSTEM_TYPE struct {
  event_queue chan CF_EVENT_TYPE
  row string
  name string
  active bool
  time_tick_duration time.Duration
  return_value interface{}
  op_code_map  map[string] CF_Function_type
  chain_map map[string]*CF_CHAIN_TYPE
  chain_order []*CF_CHAIN_TYPE
  current_chain *CF_CHAIN_TYPE
  current_link  *CF_LINK_TYPE
  current_event *CF_EVENT_TYPE
  ticker       *time.Ticker
 
}







func (system *CF_SYSTEM_TYPE )Init(cf_cluster *CF_CLUSTER_TYPE, name string ,active bool, duration time.Duration){
  (system).name = name
  (system).row = (*cf_cluster).current_row
  (system).active = active
  (system).return_value = nil
  
  (system).chain_map = make(map[string]*CF_CHAIN_TYPE)

  (system).time_tick_duration = duration
  (system).event_queue = make(chan CF_EVENT_TYPE,10)  

  (system).cf_initialize_opcodes()
  (*cf_cluster).CF_add_cf_system(system,name)
}

func ( system *CF_SYSTEM_TYPE) Add_Chain(chain_name string, state bool ){


   _,ok := (system).chain_map[chain_name]
   if ok == true{
     panic("duplicate chain name")
	}
   var temp CF_CHAIN_TYPE
   temp.name = chain_name
   temp.active = state
   temp.initialized = false
   
   (system).chain_map[chain_name] = &temp
   (system).chain_order = append((system).chain_order,&temp)
   (system).current_chain = &temp
   (system).current_link = nil
  


}

func (cf_system *CF_SYSTEM_TYPE)cf_system_execute_event(event *CF_EVENT_TYPE) {
 
  
 
  if (cf_system).active == true { // chain is  inactive throw event away
     
	 for _ , chain_data := range (cf_system).chain_order{
       
	    if (*chain_data).active == true {
		 
           (cf_system).cf_execute_chain( chain_data,event)
		} else {
		   
		}
		
	 }//for
  }   
}

func (cf_system *CF_SYSTEM_TYPE)check_for_system_events(event *CF_EVENT_TYPE) bool{

  var return_value bool
  switch (*event).Name {
    case CF_SYSTEM_STOP:
	   return_value = true
	   (cf_system).active = false
	
	case CF_SYSTEM_RESET:
	    return_value = true
		(cf_system).Initialize_Chains()
	
	default:
	   return_value = true
   }
  
   return return_value
}   


func (cf_system *CF_SYSTEM_TYPE) Execute(){

  (cf_system).Initialize_Chains()
  
  (cf_system).start_time_tick()
 
  for true { 
     var event = (cf_system).wait_for_event()
	 if (cf_system).check_for_system_events(event) {
	 
        (cf_system).cf_system_execute_event(event)
	 }
  }
}

func (cf_system *CF_SYSTEM_TYPE)Initialize_Chains(){
    if (cf_system).active == true {	
       for _ , chain_data := range (cf_system).chain_order{
	      if (*chain_data).active == true {
            (cf_system).cf_initialize_chain( chain_data)
	       }// if
	    }// for
		(cf_system).cf_system_execute_event(&cf_system_start_event)
		
	 }else{
	   ; // do nothing right now
	   
	 }
	 

}


