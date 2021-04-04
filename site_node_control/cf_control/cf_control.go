package cf

//import "time"
import "container/list"
//import "fmt"


var CF_HALT       int = 0
var CF_CONTINUE   int = 1
var CF_DISABLE    int = 2
var CF_RESET      int = 3
var CF_TERMINATE  int= 4
	


const  CF_TIME_TICK    string = "CF_TIME_TICK"
const  CF_INIT_EVENT  string = "CF_INIT_TICK"
const  CF_START_EVENT string = "CF_START_TICK"
const  CF_TERMINATE_EVENT string = "CF_TERMINATE_EVENT"



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
  name string
  active bool
  time_tick_duration int64
  return_value interface{}
  op_code_map  map[string] CF_Function_type
  chain_map map[string]*CF_CHAIN_TYPE
  chain_order []*CF_CHAIN_TYPE
  current_chain *CF_CHAIN_TYPE
  current_link  *CF_LINK_TYPE
  current_event *map[string]interface{}
  event_queue *list.List
 
}







func (system *CF_SYSTEM_TYPE )Init(cf_cluster *CF_CLUSTER_TYPE, name string ,active bool, duration int64){
  (system).name = name
  (system).active = active
  (system).return_value = nil
  
  (system).chain_map = make(map[string]*CF_CHAIN_TYPE)
  (system).event_queue = list.New()
  (system).time_tick_duration = duration
   

  (system).cf_initialize_opcodes()
  (*cf_cluster).CF_add_cf_system(system,active)
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

func (cf_system *CF_SYSTEM_TYPE)cf_system_execute_event(event *CF_EVENT_TYPE) bool{
 
  var return_value bool
  //fmt.Println("executed event",event)
  for _ , chain_data := range (cf_system).chain_order{
        
	    if (*chain_data).active == true {
           (cf_system).cf_execute_chain( chain_data,event)
		} // active
		return_value = (cf_system).active
		if return_value != true {
		  break
		} // return_value
	 }//for
     return return_value
}

func (cf_system *CF_SYSTEM_TYPE) Execute(){

  (cf_system).Initialize_Chains()
  (cf_system).cf_system_execute_event(&cf_system_start_event)
  var loop_flag = true
  for loop_flag { 
     var event = (cf_system).wait_for_event()
     loop_flag = (cf_system).cf_system_execute_event(event)
  }
}

func (cf_system *CF_SYSTEM_TYPE)Initialize_Chains(){		
    for _ , chain_data := range (cf_system).chain_order{
	   if (*chain_data).active == true {
        (cf_system).cf_initialize_chain( chain_data)
	   }
	 }

}


