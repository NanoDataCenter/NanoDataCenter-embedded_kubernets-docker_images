package sys_defs

import "lacima.com/go_setup_containers/site_generation_base/system_definitions/irrigation"
import "lacima.com/go_setup_containers/site_generation_base/site_generation_utilities"

const eto_image    string   = "nanodatacenter/eto"

func construct_irrigation( master_flag bool, node_name string){
 
   containers := []string{"eto"}
   eto_command_map  := make(map[string]string)
   eto_command_map["eto"] = "./eto"   
   su.Add_container( false,"eto",eto_image, su.Managed_run,eto_command_map, su.Data_mount)
   su.Construct_service_def("irrigation",master_flag,node_name, containers, generate_irrigation_component_graph)   
}   
  
func generate_irrigation_component_graph(){
 
    irrigation.Construct_irrigation_scheduling()
    irrigation.Construct_eto_data()
    irrigation.Add_irrigation_stations_definitions()
    irrigation.Eto_valve_setup()
    
}   

      
      

      
