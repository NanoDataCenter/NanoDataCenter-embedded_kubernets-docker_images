package node_control

import "fmt"
import "strings"
import "strconv"

import "time"
import "lacima.com/site_control_app/docker_management"
import "lacima.com/cf_control"
import "lacima.com/site_control_app/node_control/node_processor_monitoring"
import "lacima.com/Patterns/logging_support"
import "lacima.com/server_libraries/postgres"
import "lacima.com/site_control_app/docker_control"
import "lacima.com/Patterns/msgpack_2"


var docker_handle docker_management.Docker_Handle_Type
var logging_stream  pg_drv.Postgres_Stream_Driver

var local_node string



func Node_Startup(cf_cluster *cf.CF_CLUSTER_TYPE , site_data *map[string]interface{} , containers []string ){

    local_node = (*site_data)["local_node"].(string)
   
    var display_struct_search_list = []string{"DOCKER_CONTROL"}
    var incident_search_list = []string{ "INCIDENT_LOG:CONTAINER_ERROR_STREAM" ,"INCIDENT_LOG"}
    
    logging_stream  = logging_support.Find_stream_logging_driver()
    
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
   //(cf_control).Cf_add_log_link("container_monitor_loop")
   

  (cf_control).Cf_add_one_step(docker_monitor,make(map[string]interface{}))
  
   (cf_control).Cf_add_wait_interval(time.Second*15 )
   (cf_control).Cf_add_reset()
  
   
   (cf_control).Add_Chain("container_performance_logs",true)
   (cf_control).Cf_add_log_link("container_performance_loop")
   

   (cf_control).Cf_add_one_step(docker_performance_monitor,make(map[string]interface{}))
   (cf_control).Cf_add_log_link("container_performance_done")
   (cf_control).Cf_add_wait_interval(time.Minute*10 )
   (cf_control).Cf_add_reset()

   
   (cf_control).Add_Chain("processor_performance_logs",true)
   (cf_control).Cf_add_log_link("processor_performance_loop")
   

   (cf_control).Cf_add_one_step(processor_performance_monitor,make(map[string]interface{}))
   (cf_control).Cf_add_log_link("processor_performance_done")
   (cf_control).Cf_add_wait_interval(time.Minute*10 )
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



func processor_performance_monitor( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {

  working_value   := generate_parsed_fields( )
  store_performance_data("cpu",working_value["cpu"])
  store_performance_data("rss",working_value["rss"])
  store_performance_data("vsz",working_value["vsz"])
  
  return cf.CF_DISABLE
}


func store_performance_data (key string, data float64 )  {


  tag1 := "NODE_CONTROLLER"
  tag2 := local_node
  tag3 := key
  fmt.Println("pg data",tag1,tag2,tag3,data)
  packed_data  := msg_pack_utils.Pack_float64(data)
  logging_stream.Insert( tag1,tag2,tag3, "","",packed_data )

}

func  generate_parsed_fields( ) map[string]float64{

  // ps headers headers = [ "USER","PID","%CPU","%MEM","VSZ","RSS","TTY","STAT","START","TIME","COMMAND", "PARAMETER1", "PARAMETER2" ]
  return_value := make(map[string]float64)
  //cmd_string := "ps -aux | grep site_node_monitor "
  
  cmd_string := "ps -uT "
  output := docker_control.System_shell(cmd_string)
  fmt.Println("output",output)
  split_lines := strings.Split(output,"\n")
  fmt.Println("split lines",split_lines)
  data := find_process_data_lines(split_lines)
  if len(data) > 0 {
    //fmt.Println("process_lines",process_lines)
    fields := strings.Fields(data)
    fmt.Println("fields",fields)
    
    temp,err := strconv.ParseFloat(fields[2],64)
    if err != nil {
        panic("cpu error")
    }	 
    return_value["cpu"] = temp
    temp1,err := strconv.ParseFloat(fields[4],64)
    if err != nil {
        panic("vsz error")
    }	
    return_value["vsz"] = temp1
    temp2,err := strconv.ParseFloat(fields[5],64)
    if err != nil {
        panic("rss error")
    }
    return_value["rss"] = temp2
  }else{
    panic("bad data")
  }
  return return_value
}
   
func find_process_data_lines( input_lines []string )string{
    for _,input := range input_lines {
         fields := strings.Fields(input)
         if len(fields)<11 {
              continue
         }
         fmt.Println("fields",fields)
         fmt.Println("field len",len(fields))
         name_list := fields[10:]
         fmt.Println("name_list",name_list)
         if name_list[0] == "/home/pi/work/startup/site_node_monitor" {
             return input
         }
    }
    panic("process not found")
}






