package sys_defs

import "lacima.com/go_setup_containers/site_generation_base/system_definitions/irrigation"
import "lacima.com/go_setup_containers/site_generation_base/site_generation_utilities"

func construct_irrigation( master_flag bool, node_name string){
 
   containers := make([]string,0)
   su.Construct_service_def("irrigation",master_flag,node_name, containers, generate_irrigation_component_graph)   
}   
  
func generate_irrigation_component_graph(){
 
    irrigation.Construct_irrigation_scheduling()
    irrigation.Construct_eto_data()
    
}   

      
      

      
