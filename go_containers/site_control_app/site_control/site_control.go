package site_control


//import "fmt"
import "time"
import "site_control.com/docker_management"
import "site_control.com/cf_control"
import "site_control.com/site_control/site_control_up_grade"

var docker_handle docker_management.Docker_Handle_Type






func Site_Startup(cf_cluster *cf.CF_CLUSTER_TYPE , site_data *map[string]interface{}){

  
   	var container_search_list = []string{"SITE_CONTROL:SITE_CONTROL"}
    var display_struct_search_list = []string{"SITE_CONTROL:SITE_CONTROL","DOCKER_CONTROL"}
    (docker_handle).Initialize_Docker_Monitor( &container_search_list, &display_struct_search_list,site_data)
    (docker_handle).Clean_Up_Data_Structures()
	(docker_handle).Set_Initial_Hash_Values_Values()
	site_control_upgrade.Initialize_site_monitoring_data_structures( site_data)
	initialize_site_docker_monitoring(cf_cluster)
	initialize_site_docker_performance_monitoring(cf_cluster)
	site_control_upgrade.Initialize_site_monitoring_chains(cf_cluster)
}




 



  
func  initialize_site_docker_monitoring(cf_cluster *cf.CF_CLUSTER_TYPE){

   var cf_control  cf.CF_SYSTEM_TYPE

   (cf_control).Init(cf_cluster ,"site_control_docker_monitoring",true, time.Second*5)

   
   
   (cf_control).Add_Chain("container_monitoring",true)
   //(cf_control).Cf_add_log_link("container_monitor_loop")
   
   var parameters = make(map[string]interface{})
  (cf_control).Cf_add_one_step(docker_monitor,parameters)
  
   (cf_control).Cf_add_wait_interval(time.Second*15  )  // first time tick does not count aim for every 15 seconds
   (cf_control).Cf_add_reset()
  
   
}	

func  initialize_site_docker_performance_monitoring(cf_cluster *cf.CF_CLUSTER_TYPE){

   var cf_control  cf.CF_SYSTEM_TYPE

   (cf_control).Init(cf_cluster ,"site_control_docker_performance_monitoring",true, time.Minute)
   (cf_control).Add_Chain("container_performance_logs",true)
   //(cf_control).Cf_add_log_link("container_performance_loop")
   
   var  parameters = make(map[string]interface{}) 
   (cf_control).Cf_add_one_step(docker_performance_monitor,parameters)
   
   (cf_control).Cf_add_wait_interval(time.Minute*15  )
   (cf_control).Cf_add_reset()

}

func docker_monitor( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {

	// for managed containes
	
   
 
     //fmt.Println("site_control-docker_monitor")
	 (docker_handle).Monitor_Containers()
     return cf.CF_DISABLE
}


func docker_performance_monitor( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {

   //fmt.Println("site+control docker_performance_monitor") 
  (docker_handle).Log_Container_Performance_Data()
  return cf.CF_DISABLE
}
