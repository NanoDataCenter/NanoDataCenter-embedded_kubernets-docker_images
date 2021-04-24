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
  
  system_map   map[string]map[string]*CF_SYSTEM_TYPE
  current_system *CF_SYSTEM_TYPE
  current_row string

}
         
func (cf_cluster *CF_CLUSTER_TYPE) Cf_cluster_init(){

   cf_init_chain_flow_structures()
   (cf_cluster).system_map = make(map[string]map[string]*CF_SYSTEM_TYPE)
   (cf_cluster).current_system = nil
   (cf_cluster).current_row = ""

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

func (cf_cluster *CF_CLUSTER_TYPE) Cf_set_current_row( row string){
  (*cf_cluster).current_row = row
}


func (cf_cluster *CF_CLUSTER_TYPE) CF_add_cf_system(cf_system *CF_SYSTEM_TYPE,name string){

  var row = (cf_cluster ).current_row
  if row == "" {
    panic("need to set current row") 
  }
  

  if _,ok := (cf_cluster).system_map[row ];ok==false{
        (cf_cluster).system_map[row ] = make(map[string]*CF_SYSTEM_TYPE)
   }

  _,err := (cf_cluster).system_map[row ][name] 
  if err == true  {
    panic("cf_system is already defined")
  }
  
 (cf_cluster).system_map[row ][name] = cf_system
 (cf_cluster).current_system = cf_system



}


func (cf_cluster *CF_CLUSTER_TYPE) CF_Fork(){
  for _, column_data := range cf_cluster.system_map {
     for _, cf_system := range column_data {
	     go (*cf_system).Execute()
     }
  }  


}
