package node_control

import "fmt"
import "time"
import "bytes"
import "strings"
import "strconv"
import "site_control.com/docker_control"
import "site_control.com/cf_control"
import  "site_control.com/redis_support/generate_handlers"
import  "site_control.com/redis_support/redis_handlers"
import "github.com/msgpack/msgpack-go"

type processor_measure_type struct {

performance_drivers map[string]redis_handlers.Redis_Stream_Struct



}

var processor_measurement processor_measure_type

func init_processor_data_structures(site_data *map[string]interface{}){

   
   processor_measurement.performance_drivers = make(map[string]redis_handlers.Redis_Stream_Struct)
   var search_list = []string{"PROCESSOR:"+(*site_data)["local_node"].(string),"NODE_SYSTEM","PROCESSOR_MONITORING"}
   var data_element = data_handler.Construct_Data_Structures(&search_list)
   //fmt.Println("data_element",data_element)
   for key,value := range *data_element{
     //fmt.Println(key)
     processor_measurement.performance_drivers[key] = value.(redis_handlers.Redis_Stream_Struct)
   }
   

}

func (processor_measure_type) log_data(key string, data map[string]interface{} ) {

  data["time"] = time.Now().UnixNano()
  var b bytes.Buffer
	
  msgpack.Pack(&b,data)

  var driver = processor_measurement.performance_drivers[key]
  driver.Xadd(b.String())


}




func construct_processor_measurement_chains(cf_control *cf.CF_SYSTEM){



  (cf_control).Add_Chain("processor_monitoring",true)
   
  
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

  (cf_control).Cf_add_wait_interval(int64(time.Minute*15)  )
  (cf_control).Cf_add_reset()


}

func split_lines( text string  )  []string {
   
   return strings.Split(text,"\n")
  
}

func tokenize_line( text string ) [] string{
  return strings.Fields(text) 
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

func assemble_free_cpu( system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{}) int {
    return cf.CF_DISABLE
    fmt.Println("staring free cpu")

	var output = docker_control.System("sar -u 60 1 ")
	var lines = split_lines(output)
	
	var average_line = lines[len(lines)-2]
	
	var tokens = tokenize_line(average_line)
	
	var data = tokens_to_dict(tokens,[]string{ "Time","cpu","%user" , "%nice", "%system", "%iowait" ,"%steal" ,"%idle" },2)
	fmt.Println("data",data)
	
	(processor_measurement).log_data("FREE_CPU",data) 
	

  return cf.CF_DISABLE
}

func assemble_ram( system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{}) int {

  var output = docker_control.System("cat /proc/meminfo ")
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

func assemble_temperature( system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{}) int {

  var output = docker_control.System("vcgencmd measure_temp")
  var output1 = strings.Replace(output, "temp=", "", -1)
  var output2 = strings.Replace(output1, "'C\n", "", -1)
  fmt.Println(output2)
  value,err := strconv.ParseFloat(output2,64)
  if err != nil{
	  panic("bad float")
  
  }
 value = (9.0/5.0*value)+32.
 var data = make(map[string]interface{})
 data["TEMP_F"] = value
 fmt.Println(data)
 (processor_measurement).log_data("TEMPERATURE",data) 
  return cf.CF_DISABLE

  
}

func assemble_disk_space( system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{}) int {

  var output = docker_control.System("df")
 
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



func assemble_swap_space( system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{}) int {

  parse_one_line("sar -S 1 1","SWAP_SPACE")    
  return cf.CF_DISABLE
}

func assemble_io_space( system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{}) int {

  parse_one_line("sar -w 1 1","IO_SPACE")  
  return cf.CF_DISABLE
}

func assemble_block_io( system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{}) int {

  parse_multi_line("sar -d  3 1","BLOCK_DEV",-1)
  return cf.CF_DISABLE
}


func assemble_context_switches( system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{}) int {

  parse_one_line("sar -w 1 1","CONTEXT_SWITCHES") 
  return cf.CF_DISABLE
}

func assemble_run_queue( system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{}) int {

  parse_one_line("sar -q 3 1","RUN_QUEUE")   
  return cf.CF_DISABLE
}

func assemble_net_edev( system interface{},chain interface{}, parameters map[string]interface{}, event *map[string]interface{}) int {

  parse_multi_line("sar -n EDEV  3 1","EDEV",2)
  return cf.CF_DISABLE
}


func parse_multi_line(sar_command,stream_key string,ref_index int){

    var output = docker_control.System(sar_command)
    var data = make(map[string]interface{})
    var lines = split_lines(output)
    for i,line := range lines{
      fmt.Println(i,line)
    }
    panic("done")
    print("data",stream_key,data)   
   (processor_measurement).log_data(stream_key,data)
	
      
/*       f = os.popen(sar_command)
       data = f.read()
       f.close()
       lines = data.split("\n")
       i = 3
       data = {}
       while True:
          line = lines[i]
          if line == "":
             break
          line = re.sub(' +',' ',line)
          fields = line.split(" ")
          
          key = fields[1]
          value = fields[ref_index]
          data[key] = float(float(value))
          i = i+1
*/

	   
}


func parse_one_line(sar_command, stream_key string){

    var output = docker_control.System(sar_command)
    var data = make(map[string]interface{})
    var lines = split_lines(output)
    
	var key_tokens =  tokenize_line(lines[2])
	var value_tokens = tokenize_line(lines[4])
	fmt.Println(key_tokens,value_tokens,value_tokens)
    panic("done")
    print("data",stream_key,data)   
   (processor_measurement).log_data(stream_key,data)
}
/*	   
   def parse_one_line(self, sar_command, stream_field ):
        f = os.popen(sar_command)
        data = f.read()
        f.close()

        lines = data.split("\n")
        line = lines[2]
        line = re.sub(' +',' ',line)
        fields_keys = line.split(" ")
        line = lines[3]
        line = re.sub(' +',' ',line)
        fields_data = line.split(" ")
        fields_data.pop(0)
        fields_keys.pop(0)
        data = {}
        for i in range(0,len(fields_keys)):
           data[fields_keys[i]] = float(fields_data[i])
       
        print("data",data)
          
        self.ds_handlers[stream_field].push(data = data,local_node = self.site_node)
 
 
*/