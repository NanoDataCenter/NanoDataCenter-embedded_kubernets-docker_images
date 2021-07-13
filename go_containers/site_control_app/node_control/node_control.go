package node_control



import "time"
import "lacima.com/site_control_app/docker_management"
import "lacima.com/cf_control"
import "lacima.com/site_control_app/node_control/node_processor_monitoring"
//import "lacima.com/site_control_app/node_control/monitor_command_upgrade_queues"


var docker_handle docker_management.Docker_Handle_Type






func Node_Startup(cf_cluster *cf.CF_CLUSTER_TYPE , site_data *map[string]interface{} , containers []string ){

  
   
    var display_struct_search_list = []string{"DOCKER_CONTROL"}
    var incident_search_list = []string{ "INCIDENT_LOG:CONTAINER_ERROR_STREAM" ,"INCIDENT_LOG"}
    
    
    (docker_handle).Initialize_Docker_Monitor(containers , &display_struct_search_list,&incident_search_list,site_data)
    
	(docker_handle).Set_Initial_Hash_Values_Values()
    
	node_perform.Init_processor_data_structures(site_data )
	
	initialize_node_docker_monitoring(cf_cluster)
	
    node_perform.Initialize_node_processor_performance(cf_cluster)
	 setup_site_control(cf_cluster,site_data)
   
}



 


  
func  initialize_node_docker_monitoring(cf_cluster *cf.CF_CLUSTER_TYPE){

   var cf_control  cf.CF_SYSTEM_TYPE

   (cf_control).Init(cf_cluster ,"node_control_docker_monitoring",true, time.Second)
   
   (cf_control).Add_Chain("container_monitoring",true)
   (cf_control).Cf_add_log_link("container_monitor_loop")
   
   var parameters = make(map[string]interface{})
  (cf_control).Cf_add_one_step(docker_monitor,parameters)
  
   (cf_control).Cf_add_wait_interval(time.Second*15 )
   (cf_control).Cf_add_reset()
  
   (cf_control).Add_Chain("container_performance_logs",true)
   (cf_control).Cf_add_log_link("container_performance_loop")
   
   parameters = make(map[string]interface{}) 
   (cf_control).Cf_add_one_step(docker_performance_monitor,parameters)
   
   (cf_control).Cf_add_wait_interval(time.Minute*1 )
   (cf_control).Cf_add_reset()

   
   
   

}	


func docker_monitor( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {

	// for managed containes
	
   

    
	 (docker_handle).Monitor_Containers()
     return cf.CF_DISABLE
}


func docker_performance_monitor( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {

  
  (docker_handle).Log_Container_Performance_Data()
  return cf.CF_DISABLE
}
