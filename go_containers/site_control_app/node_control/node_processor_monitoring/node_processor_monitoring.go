package node_perform

/*
reference for sar function is

https://www.thegeekstuff.com/2011/03/sar-examples/

*/



import "time"
import "bytes"
import "strings"
import "strconv"
import  "lacima.com/site_control_app/docker_control"
import "lacima.com/cf_control"
import  "lacima.com/redis_support/generate_handlers"
import  "lacima.com/redis_support/redis_handlers"
import "github.com/msgpack/msgpack-go"


type processor_measure_type struct {

performance_drivers map[string]redis_handlers.Redis_Stream_Struct



}

var processor_measurement processor_measure_type

func Init_processor_data_structures(site_data *map[string]interface{}){

   
   processor_measurement.performance_drivers = make(map[string]redis_handlers.Redis_Stream_Struct)
   var search_list = []string{"PROCESSOR:"+(*site_data)["local_node"].(string),"STREAMING_LOG:processor_monitor","STREAMING_LOG"}
   var data_element = data_handler.Construct_Data_Structures(&search_list)
   
   keys := []string{"FREE_CPU","RAM","TEMPERATURE","DISK_SPACE","SWAP_SPACE","CONTEXT_SWITCHES","BLOCK_DEV","IO_SPACE","RUN_QUEUE","EDEV"}
   for _,key := range keys{
     
     processor_measurement.performance_drivers[key] = (*data_element)[key].(redis_handlers.Redis_Stream_Struct)
   }
   

}

func (processor_measure_type) log_data(key string, data map[string]interface{} ) {

  var driver = processor_measurement.performance_drivers[key]
 
  data["time"] = time.Now().UnixNano()
  var b bytes.Buffer	
  msgpack.Pack(&b,data)
  driver.Xadd(b.String())


}



func Initialize_node_processor_performance(cf_cluster *cf.CF_CLUSTER_TYPE){


   var cf_control  cf.CF_SYSTEM_TYPE
  
  (cf_control).Init(cf_cluster , "node_control_processor_monitor" ,true, time.Minute )




  (cf_control).Add_Chain("processor_monitoring",true)
  (cf_control).Cf_add_log_link("processor monitoring started ")
  
   var par1 = make(map[string]interface{})
  (cf_control).Cf_add_one_step(assemble_free_cpu,par1)
  var par2 = make(map[string]interface{})
  (cf_control).Cf_add_one_step(assemble_ram,par2)
  var par3 = make(map[string]interface{})
  (cf_control).Cf_add_one_step(assemble_temperature,par3)
 
  var par7 = make(map[string]interface{})
  (cf_control).Cf_add_one_step(assemble_disk_space,par7)
  var par8 = make(map[string]interface{})
  (cf_control).Cf_add_one_step(assemble_swap_space,par8)
  var par10 = make(map[string]interface{})
  (cf_control).Cf_add_one_step(assemble_io_space,par10)
  var par12 = make(map[string]interface{})
  (cf_control).Cf_add_one_step(assemble_block_io,par12)
  var par13 = make(map[string]interface{})
  (cf_control).Cf_add_one_step(assemble_context_switches,par13)
  var par14 = make(map[string]interface{})
  (cf_control).Cf_add_one_step(assemble_run_queue,par14)
  var par15 = make(map[string]interface{})
  (cf_control).Cf_add_one_step(assemble_net_edev,par15)
  (cf_control).Cf_add_log_link("processor monitoring done")
  (cf_control).Cf_add_wait_interval(time.Minute*9  ) // first tick is not counted sar -u 300 1 takes 5 minutes
  (cf_control).Cf_add_reset()


}

func split_lines( text string  )  []string {
   
   return strings.Split(text,"\n")
  
}

func tokenize_line( text string ) [] string{
  return strings.Fields(text) 
}

func string_to_float64( text string ) float64 {
   value,err := strconv.ParseFloat(text,64)
   if err != nil {
      
	  value = 0
   }
   return value
}

func tokens_to_dict(tokens []string, header []string, start_index int) map[string]interface{} {

    var return_value = make(map[string]interface{})
	for i:= start_index; i < len(header); i++ {
	   var key = header[i]
	   var value,err = strconv.ParseFloat(tokens[i],64)
	   if err != nil {
	     panic("bad value")
	   }
	   return_value[key]=value
	}
	return return_value


}

func assemble_free_cpu( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {
    
    
	var output = docker_control.System_shell("sar -u 300 1 ")
		var lines = split_lines(output)
	
	var average_line = lines[len(lines)-2]
	
	var tokens = tokenize_line(average_line)
	
	var data = tokens_to_dict(tokens,[]string{ "Time","cpu","%user" , "%nice", "%system", "%iowait" ,"%steal" ,"%idle" },2)
	
	(processor_measurement).log_data("FREE_CPU",data) 
	

  return cf.CF_DISABLE
}

func assemble_ram( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {

  var output = docker_control.System_shell("cat /proc/meminfo ")
  
  var data = make(map[string]interface{})
  var lines = split_lines(output)
  for _,line := range lines{
      var tokens = tokenize_line(line)
	  if len(tokens) == 3{
	      var key = strings.Replace(tokens[0], ":", "", -1)
          var value,err = strconv.ParseFloat(tokens[1],64)
	      if err != nil{
	         panic("bad float")
	      }
          data[key] = value
	  }
	 
  }
  
  (processor_measurement).log_data("RAM",data) 
  return cf.CF_DISABLE

}

func assemble_temperature( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {

  var output = docker_control.System_shell("vcgencmd measure_temp")
  var output1 = strings.Replace(output, "temp=", "", -1)
  var output2 = strings.Replace(output1, "'C\n", "", -1)
 
  value,err := strconv.ParseFloat(output2,64)
  if err != nil{
	  panic("bad float")
  
  }
 value = (9.0/5.0*value)+32.
 var data = make(map[string]interface{})
 data["TEMP_F"] = value
 
 (processor_measurement).log_data("TEMPERATURE",data) 
  return cf.CF_DISABLE

  
}

func assemble_disk_space( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {

  var output = docker_control.System_shell("df")
 
  var data = make(map[string]interface{})
  var lines = split_lines(output)
  for _ ,line := range lines{
    var tokens = tokenize_line(line)
	if len(tokens) > 5 {
	   if tokens[0] == "Filesystem"{
	      continue
	   }
	   if tokens[0] == "overlay"{
	      continue
	    }
	   var value_string = strings.Replace(tokens[4], "%", "", -1)
	   value,err := strconv.ParseFloat(value_string,64)
       if err != nil{
	       panic("bad float")
  
       }
	 
	   data[tokens[5]] = value	  
	   
	}
  }
  (processor_measurement).log_data("DISK_SPACE",data)
  return cf.CF_DISABLE
}



func assemble_swap_space( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {

 
    var output = docker_control.System_shell("sar -S 1 1")
    var data = make(map[string]interface{})
    var lines = split_lines(output)
    
	var key_tokens =  tokenize_line(lines[2])
	var value_tokens = tokenize_line(lines[4])
	
	var free  = string_to_float64(value_tokens[1])
	var used  = string_to_float64(value_tokens[2])
	
    data[key_tokens[1]] = free
	data[key_tokens[2]] = used
    
	
   (processor_measurement).log_data("SWAP_SPACE",data)
   return cf.CF_DISABLE
}

func assemble_context_switches( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {

   var output = docker_control.System_shell("sar -w 1 1")
      
    var data = make(map[string]interface{})
    var lines = split_lines(output)
    
	var key_tokens =  tokenize_line(lines[2])
	var value_tokens = tokenize_line(lines[4])
	
	var proc_s  = string_to_float64(value_tokens[1])
	var cswch_s  = string_to_float64(value_tokens[2])
	
    data[key_tokens[1]] = proc_s
	data[key_tokens[2]] = cswch_s
   
	
   (processor_measurement).log_data("CONTEXT_SWITCHES",data)
   return cf.CF_DISABLE
}

func assemble_block_io( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {

  
   var output = docker_control.System_shell("sar -d  3 1")
   
   var data = make(map[string]interface{})
   var lines = split_lines(output)
   var data_lines = lines[3:]
   for _,line := range data_lines {
    
     if len(line) == 0{
	   break
	 }
   var tokens = tokenize_line(line)
   
   var key = tokens[1]
   var value = string_to_float64(tokens[len(tokens)-1])
   data[key] = value
   
   }
   
  
   
   (processor_measurement).log_data("BLOCK_DEV",data)
   return cf.CF_DISABLE

  
}


func assemble_io_space( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {

  
  var output = docker_control.System_shell("sar -b 1 1")
   
  var data = make(map[string]interface{})
  var lines = split_lines(output)
  var key_line = lines[2]
  var data_line = lines[3]
  var key_tokens =  tokenize_line(key_line)
  var data_tokens = tokenize_line(data_line)
   
  for i :=1;i<len(key_tokens);i++{
     data[key_tokens[i]]  = string_to_float64(data_tokens[i])
  }
  (processor_measurement).log_data("IO_SPACE",data)
  return cf.CF_DISABLE
 
}

func assemble_run_queue( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {

    
   var output = docker_control.System_shell("sar  -q 1 1")
   
  var data = make(map[string]interface{})
  var lines = split_lines(output)
  var key_line = lines[2]
  var data_line = lines[3]
  var key_tokens =  tokenize_line(key_line)
  var data_tokens = tokenize_line(data_line)
  for i :=1;i<len(key_tokens);i++{
     data[key_tokens[i]]  = string_to_float64(data_tokens[i])
  }
  

  (processor_measurement).log_data("RUN_QUEUE",data)
  return cf.CF_DISABLE
}

func assemble_net_edev( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE) int {

  
  var output = docker_control.System_shell("sar -n EDEV  3 1")
  var data = make(map[string]interface{})
  var lines = split_lines(output)
  var data_lines = lines[3:]
  for _,line := range data_lines {
    
     if len(line) == 0{
	   break
	 }
   var tokens = tokenize_line(line)
   
   var key = tokens[1]
   var value = string_to_float64(tokens[2])
   data[key] = value
   
   }
   
  
   
   
   (processor_measurement).log_data("EDEV",data)
   return cf.CF_DISABLE

}

