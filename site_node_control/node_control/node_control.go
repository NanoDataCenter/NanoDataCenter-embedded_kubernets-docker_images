package node_control


//import "fmt"
import "time"
import "site_control.com/docker_management"
import "site_control.com/cf_control"

var cf_control  cf.CF_SYSTEM
var cf_performance cf.CF_SYSTEM
var docker_handle docker_management.Docker_Handle_Type






func Node_Startup(site_data *map[string]interface{}){

   initialize_CF()
   	var container_search_list = []string{"PROCESSOR:"+(*site_data)["local_node"].(string)}
    var display_struct_search_list = []string{"PROCESSOR:"+(*site_data)["local_node"].(string),"DOCKER_CONTROL"}
    (docker_handle).Initialize_Docker_Monitor( &container_search_list, &display_struct_search_list,site_data)
    (docker_handle).Clean_Up_Data_Structures()
	(docker_handle).Set_Initial_Hash_Values_Values()
	init_processor_data_structures(site_data )
}



 

func Execute(){
   
  (cf_control).Execute()

}


func Performance_Execute(){

  (cf_performance).Execute()
}
  
func  initialize_CF(){


   (cf_control).Init("Node Control")
   
   (cf_control).Add_Chain("container_monitoring",true)
   //(cf_control).Cf_add_log_link("container_monitor_loop")
   
   var parameters = make(map[string]interface{})
  (cf_control).Cf_add_one_step(docker_monitor,parameters)
  
   (cf_control).Cf_add_wait_interval(int64(time.Second*15)  )
   (cf_control).Cf_add_reset()
  
   (cf_control).Add_Chain("container_performance_logs",true)
   //(cf_control).Cf_add_log_link("container_performance_loop")
   
   parameters = make(map[string]interface{}) 
   (cf_control).Cf_add_one_step(docker_performance_monitor,parameters)
   
   (cf_control).Cf_add_wait_interval(int64(time.Minute*15)  )
   (cf_control).Cf_add_reset()
   
   (cf_performance).Init("Performance Monitoring")
   construct_processor_measurement_chains(&cf_performance)
}	


func docker_monitor( system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{}) int {

	// for managed containes
	
   

     
	 (docker_handle).Monitor_Containers()
     return cf.CF_DISABLE
}


func docker_performance_monitor( system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{}) int {

  
  (docker_handle).Log_Container_Performance_Data()
  return cf.CF_DISABLE
}
