package cf


type CF_EVENT_TYPE struct{
  Name string
  Value interface{}
}

var cf_link_init_event CF_EVENT_TYPE
var cf_system_start_event CF_EVENT_TYPE
var cf_system_terminate_event CF_EVENT_TYPE
var cf_continue_map map[int]bool


type CF_CLUSTER_TYPE struct {
  name string   // extensions to higher order
  active bool   // extensions to hiher order
  value interface{}
  
  // in future map want to add a special opcode map
  // in future add channel
  system_map   map[string]*CF_SYSTEM_TYPE
  system_order []*CF_SYSTEM_TYPE
  current_system *CF_SYSTEM_TYPE

}
         
func (cf_cluster *CF_CLUSTER_TYPE) Cf_cluster_init(name string , active bool){

   cf_init_chain_flow_structures()
   (cf_cluster).active = active
   (cf_cluster).name   = name
   (cf_cluster).value = nil
   (cf_cluster).system_map = make(map[string]*CF_SYSTEM_TYPE)
   (cf_cluster).current_system = nil

}

// initializing global variables
func cf_initialize_continue_map(){
  var temp =  make(map[int]bool)
  temp[CF_HALT] = false
  temp[CF_CONTINUE] = true
  temp[CF_DISABLE] = true
  temp[CF_RESET] = false
  temp[CF_TERMINATE] = false
  cf_continue_map = temp
}


func cf_init_chain_flow_structures(){

  cf_initialize_continue_map()
  //cf_link_init_event = make(map[string]interface{})
  cf_link_init_event.Name = CF_INIT_EVENT
  cf_link_init_event.Value =  nil
  
  //cf_system_start_event = make(map[string]interface{})
  cf_system_start_event.Name = CF_START_EVENT
  cf_system_start_event.Value = nil  // value is not used
  
  //cf_system_terminate_event = make(map[string]interface{})
  cf_system_terminate_event.Name = CF_TERMINATE_EVENT
  cf_system_terminate_event.Value = nil  // value is not used

}

func (cf_cluster *CF_CLUSTER_TYPE) CF_add_cf_system(cf_system *CF_SYSTEM_TYPE, active bool){

  (*cf_system).active = active
  var name = (*cf_system).name

  _,err := (cf_cluster).system_map[name] 
  if err == true  {
    panic("cf_system is already defined")
  }

 (cf_cluster).system_map[name] = cf_system
 (cf_cluster).system_order = append(cf_cluster.system_order,cf_system)



}


func (cf_cluster *CF_CLUSTER_TYPE) CF_Fork(){

  for _,system := range cf_cluster.system_order{
     if (*system).active == true {
        go (*system).Execute()
     }
  }  


}
