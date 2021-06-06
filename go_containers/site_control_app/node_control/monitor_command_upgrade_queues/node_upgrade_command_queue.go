package node_up

//import "fmt"
import "time"
import "bytes"
import "lacima.com/cf_control"
import "lacima.com/site_control_app/docker_control"
import  "lacima.com/redis_support/generate_handlers"
import  "lacima.com/redis_support/redis_handlers"
import  "lacima.com/Patterns/msgpack"
import  "lacima.com/Patterns/shell_utils"
import "github.com/msgpack/msgpack-go"



type node_upgrade_data_type struct {

 watchdog_driver redis_handlers.Redis_Single_Structure
 node_command_queue redis_handlers.Redis_Job_Queue
 node_upgrade_queue redis_handlers.Redis_Job_Queue

} 

const reboot_time_delay = time.Second*15

var node_upgrade_data_structures node_upgrade_data_type

type job_function_type func( input interface{} ) 

var job_table map[string]job_function_type


func reboot_system( input interface{}){

  if (node_upgrade_data_structures).upgrade_queue_length() != 0 {
      return  // upgrades are still waiting
	 
  }
  time.Sleep(reboot_time_delay)
  shell_utils.System_shell("reboot now")
  

}
Bc_Rec.Add_header_node("PROCESSOR_WATCHDOG","PROCESSOR_WATCHDOG", make(map[string]interface{}))
      Construct_watchdog_logging("PROCESSOR_WATCHDOG")   
	  Bc_Rec.End_header_node("PROCESSOR_WATCHDOG","PROCESSOR_WATCHDOG")
func Node_command_queue_structures(site_data *map[string]interface{}){

   var job_table = make( map[string]job_function_type )
   job_table["reboot"] = reboot_system
  
   // initialize data structures
   var search_list = []string{"PROCESSOR:"+(*site_data)["local_node"].(string),"NODE_SYSTEM","NODE_CONTROL"}
   var data_element = data_handler.Construct_Data_Structures(&search_list)  

   node_upgrade_data_structures.watchdog_driver = (*data_element)["NODE_WATCH_DOG"].(redis_handlers.Redis_Single_Structure)
   node_upgrade_data_structures.node_command_queue = (*data_element)["NODE_COMMAND_QUEUE"].(redis_handlers.Redis_Job_Queue)
   node_upgrade_data_structures.node_upgrade_queue = (*data_element)["NODE_UPGRADE_QUEUE"].(redis_handlers.Redis_Job_Queue)
   // resetting old values
   node_upgrade_data_structures.node_upgrade_queue.Delete_all()
   node_upgrade_data_structures.node_command_queue.Delete_all()
   node_upgrade_data_structures.watchdog_driver.Delete_all()
}

func ( v node_upgrade_data_type )store_time_stamp(){

  var data = time.Now().UnixNano()
  var b bytes.Buffer	
  msgpack.Pack(&b,data)
  v.watchdog_driver.Set(b.String())


}

func ( v node_upgrade_data_type )get_time_stamp()int64{

 
  var return_value = msgpack_utils.Unpack(v.watchdog_driver.Get()).(int64)
  return return_value
 
}

func ( v node_upgrade_data_type )get_command_queue()(int64, *cf.CF_EVENT_TYPE){

   var length = v.node_command_queue.Length()
   if length == 0 {
       return 0, nil
   }
   
   var val =  msgpack_utils.Unpack((v.node_command_queue).Show_next_job()).(*cf.CF_EVENT_TYPE)
   return length, val
   

}

func ( v node_upgrade_data_type )pop_command_queue(){


   (v.node_command_queue).Pop()

}


func ( v node_upgrade_data_type )upgrade_queue_length()int64{

   return v.node_command_queue.Length()


}


func ( v node_upgrade_data_type )get_upgrade_queue()(int64, *cf.CF_EVENT_TYPE){

   var length = v.node_upgrade_queue.Length()
   if length == 0 {
       return 0, nil
   }
   var val =  msgpack_utils.Unpack((v.node_upgrade_queue).Show_next_job()).(*cf.CF_EVENT_TYPE)
   return length, val
   

}

func ( v node_upgrade_data_type )pop_upgrade_queue(){


   (v.node_upgrade_queue).Pop()

}





func  Initialize_node_job_server_watch_dog_cf(cf_cluster *cf.CF_CLUSTER_TYPE){

   var cf_control  cf.CF_SYSTEM_TYPE

   (cf_control).Init(cf_cluster, "node_command_queue_watch_dog",true, time.Second)
   
   
   
   (cf_control).Add_Chain("container_monitoring",true)   // watch dog strobe
   
   var parameters = make(map[string]interface{})
   ( cf_control).Cf_add_one_step(node_strobe_watch_dog,parameters)
   (cf_control).Cf_add_wait_interval(time.Second*15  ) // every 15 seconds
   (cf_control).Cf_add_reset()
  
   (cf_control).Add_Chain("monitor_node_command_queue",true) // monitor command from site_contol
   
   parameters = make(map[string]interface{}) 
   (cf_control).Cf_add_unfiltered_element(node_process_job_queue,parameters)
   (cf_control).Cf_add_reset()
   
   (cf_control).Add_Chain("monitor_upgrade_queue",true) // monitor command from site_contol
   
   parameters = make(map[string]interface{}) 
   (cf_control).Cf_add_unfiltered_element(node_process_upgrade_queue,parameters)
   (cf_control).Cf_add_reset()
   
}

func node_strobe_watch_dog( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{

 
  (node_upgrade_data_structures ).store_time_stamp()
 
  return cf.CF_DISABLE
  
}
  
  
  
func node_process_job_queue( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{

  
  job_length, job_data := (node_upgrade_data_structures ).get_command_queue()
  //fmt.Println("node job queue")
  if job_length != 0 {
      var job_name = (*job_data).Name
      if   _,flag := job_table[(*job_data).Name]; flag == false {
	      panic("bad job name ")
	 }
	 job_table[job_name]((*job_data).Value)
     (node_upgrade_data_structures ).pop_command_queue()
  }
  return cf.CF_HALT
  
}

func node_process_upgrade_queue( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{

  
    job_length, job_data := (node_upgrade_data_structures ).get_upgrade_queue()
  //fmt.Println("node upgrade queue")
  if job_length != 0 {
    var container_image = (*job_data).Name
    docker_control.Pull(container_image)
	 
	 
     (node_upgrade_data_structures ).pop_upgrade_queue()
  }
  return cf.CF_HALT
 
  return cf.CF_HALT
  
}
