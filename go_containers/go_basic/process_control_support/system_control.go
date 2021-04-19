package system_control

import "fmt"
import "time"
import "bytes"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/system_error_logging"
import "lacima.com/redis_support/graph_query"
import "lacima.com/cf_control"
import "github.com/msgpack/msgpack-go"


type System_Control_Type struct {
  system_log              *system_log.SYSTEM_LOGGING_RECORD
  site                     string
  local_node               string
  container_name           string
  file_name                string
  web_display              redis_handlers.Redis_Hash_Struct
  process_failure_status   redis_handlers.Redis_Hash_Struct
  process_failure_logs     redis_handlers.Redis_Stream_Struct
  watchdog_strobe          redis_handlers.Redis_Single_Structure
  process_map              map[string]string
  process_ctrl             []*Process_Manager_Type
  
}



func Construct_System_Control( sys_log *system_log.SYSTEM_LOGGING_RECORD,site, local_node,container_name,file_name string ) *System_Control_Type {

   var return_value  System_Control_Type
   return_value.site = site
   return_value.local_node = local_node
   return_value.container_name = container_name
   return_value.file_name = file_name
   return_value.system_log = sys_log
   return_value.process_map = make(map[string]string)
   return &return_value
}   

func ( v *System_Control_Type) Init(cf_cluster *cf.CF_CLUSTER_TYPE){

   search_list :=[]string{ "CONTAINER:"+v.container_name ,"DATA_STRUCTURES" }
   handlers := data_handler.Construct_Data_Structures(&search_list)
   v.web_display = (*handlers)["WEB_DISPLAY_DICTIONARY"].(redis_handlers.Redis_Hash_Struct)
   v.process_failure_status = (*handlers)["Process_Status"].(redis_handlers.Redis_Hash_Struct)
   v.process_failure_logs = (*handlers)["Process_Failure"].(redis_handlers.Redis_Stream_Struct)
   v.watchdog_strobe = (*handlers)["controller_watchdog"].(redis_handlers.Redis_Single_Structure)
   
   v.web_display.Delete_All()  // remove elements which may not belong
   
   search_list = []string{ "CONTAINER:"+v.container_name  }
   //fmt.Println("site",v.site)
   //fmt.Println(search_list)
   nodes := graph_query.Common_qs_search( &search_list)
   node:= nodes[0]
   process_list := node["command_list"]
   v.verify_process_map(process_list)
   //fmt.Println("process_map",process_list)
   v.construct_chains(cf_cluster)
   
  
}


func ( v *System_Control_Type ) verify_process_map( process_list string ) {

   data := graph_query.Convert_json_dict_array(process_list)
   for _,element := range data { 
     
	 key,ok := element["key"] 
	 if ok != true {
	    panic("bad key")
	 }
	 command,ok1 := element["command"]
	 if ok1 != true {
	    panic("bad command")
	 }
   
    v.process_map[key] = command
	v.process_ctrl = append(v.process_ctrl,construct_process_manager( key,command) )
   
   }


}




func ( v *System_Control_Type ) construct_chains(cf_cluster *cf.CF_CLUSTER_TYPE) {

   var cf_control  cf.CF_SYSTEM_TYPE

   (cf_control).Init(cf_cluster ,"container_process_monitor",true, time.Second)
   
   (cf_control).Add_Chain("initialization",true)
   //(cf_control).Cf_add_log_link("cf initialization chain")
   var parameters = make(map[string]interface{})
   (cf_control).Cf_add_one_step(v.launch_processes,parameters)
   (cf_control).Cf_add_enable_chains_links( []string{"monitor_active_processes"}  )
   (cf_control).Cf_add_terminate()  


   (cf_control).Add_Chain("strobe_watchdog",true)
   //(cf_control).Cf_add_log_link("strobe_watch_dog")
   var parameters1 = make(map[string]interface{})
   (cf_control).Cf_add_one_step(v.strobe_watch_dog,parameters1)
   (cf_control).Cf_add_wait_interval(time.Second*5  )
   (cf_control).Cf_add_reset()

  
    (cf_control).Add_Chain("monitor_active_processes",false)
   //(cf_control).Cf_add_log_link("strobe_watch_dog")
    var parameters2 = make(map[string]interface{})
   (cf_control).Cf_add_one_step(v.monitor_process,parameters2)
   (cf_control).Cf_add_wait_interval(time.Second*10  )
   (cf_control).Cf_add_reset()

}
   
func ( v *System_Control_Type )launch_processes( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{

  
   for _,element := range v.process_ctrl {
   
     go (*element).run()
   
   }
   
   
   return cf.CF_DISABLE
}

func ( v *System_Control_Type )strobe_watch_dog( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{

   var data = time.Now().UnixNano()
   var b bytes.Buffer	
   msgpack.Pack(&b,data)
   v.watchdog_strobe.Set(b.String())


   return cf.CF_DISABLE
}

func ( v *System_Control_Type )monitor_process( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{
   for _,element := range v.process_ctrl {
       v.update_web_display(element)
       v.log_error_message(element)
   }
   return cf.CF_DISABLE
}

func ( v *System_Control_Type )update_web_display( element  *Process_Manager_Type){

   key := (*element).key
   state := true
   if (*element).failed == true {
      state = false
   }
   var b bytes.Buffer	
   msgpack.Pack(&b,state)
   

   v.web_display.HSet(key,b.String())
}



func ( v *System_Control_Type )log_error_message( element  *Process_Manager_Type){

  key := (*element).key
  if (*element).failed == false {
     return
  }else{
    error_message := (*element).error_log
	(*element).failed = false
	 var b bytes.Buffer	
    msgpack.Pack(&b,error_message)
	new_message := b.String()
	ref_message := v.process_failure_status.HGet(key)
	
	if ref_message != new_message {
	   v.process_failure_status.HSet(key,new_message)
	   v.process_failure_logs.Xadd(new_message)
	   fmt.Println("error_message is differenet ####################")
	}else{
	   fmt.Println("error_message is the same --------------------------")
	}
  }


}

  

