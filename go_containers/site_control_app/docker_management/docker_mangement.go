package docker_management

//import "fmt"
import "bytes"

import "strings"
import "strconv"
import "lacima.com/site_control_app/docker_control"
import  "lacima.com/Patterns/msgpack"
import "github.com/msgpack/msgpack-go"






func (docker_handle *Docker_Handle_Type)Set_Initial_Hash_Values_Values(){

   var default_value = make(map[string]bool)
   default_value["managed"] = true
   default_value["active"] = true
   for _,container := range (docker_handle).containers{
      
      (docker_handle).hset_status_values(container,&default_value)
	  
   }
  
}



func (docker_handle *Docker_Handle_Type)Monitor_Containers(){
  
  check_map := make(map[string]bool)
  
  for _,container := range docker_handle.containers {
     check_map[container] = false
  }
  
  
  running_containers := docker_control.Containers_ls_runing()
  //fmt.Println("running containers",running_containers)
  for _,running_container := range running_containers {
    //fmt.Println("running container",running_container)
     _,ok := check_map[running_container]
     //fmt.Println("ok",ok)
	 if ok == true{
	    check_map[running_container] = true
	}
  }
  //fmt.Println("container map",check_map)
  for _,container := range( docker_handle).containers {
     
     //fmt.Println("container",container,(docker_handle).hget_status_value(container))
	 container_status := (docker_handle).hget_status_value(container)
	 if (*container_status)["managed"] == false {
         continue    // do nothihg if container is not monitored
     }
	 // update status values
	 
	 
	 if (*container_status)["active"] != check_map[container] {
	   
	     (*container_status)["active"] = check_map[container]
		 (docker_handle).hset_status_values(container,container_status)
         (docker_handle).add_incident_log(container,container_status,check_map[container])
		 
	 }
	
	  if check_map[container] == false{
	    
	     docker_control.Container_start(container)
		
		 
	  }
		
	 
	 
  
  
  }
  

}


func (docker_handle *Docker_Handle_Type)Log_Container_Performance_Data(){

   for _,container := range (docker_handle).containers {
       var working_values = (docker_handle).generate_parsed_fields(container)
	   if working_values != nil {
	    
	     (*docker_handle).store_performance_data(container, "cpu", "CPU",working_values)
	     (*docker_handle).store_performance_data(container, "rsz", "RSS",working_values)
	     (*docker_handle).store_performance_data(container, "vsz", "VSZ",working_values)
	   }
   }	  
  


}

func (docker_handle *Docker_Handle_Type) store_performance_data (container string,data_key string,redis_key string ,working_values *map[string]map[string]float64  ){

  
  var output_data = make(map[string]float64)
  
 
  for process, data := range *working_values {
    output_data[process] = data[data_key]
  
  }
  
  var driver_array = (docker_handle).docker_performance_drivers[container]
  var driver = driver_array[redis_key]
  
   var b bytes.Buffer	
   msgpack.Pack(&b,output_data)

  driver.Xadd(b.String())


}



func (docker_handle *Docker_Handle_Type) generate_parsed_fields( container_name string) *map[string]map[string]float64{

  // ps headers headers = [ "USER","PID","%CPU","%MEM","VSZ","RSS","TTY","STAT","START","TIME","COMMAND", "PARAMETER1", "PARAMETER2" ]
  var return_value = make(map[string]map[string]float64)
  var cmd_string = "docker top "+container_name+ "  -aux "
  var output = docker_control.System_shell(cmd_string)
  var skip_lines = 1
  var split_lines = strings.Split(output,"\n")
  
  if len(split_lines) <= skip_lines {
    return nil
  }
   
  var process_lines = split_lines[skip_lines:]
  for _, data := range process_lines {
      if len(data) > 0 {
	  
  	       var fields = strings.Fields(data)
		   var name_list = fields[10:]
		   var process_name = strings.Join(name_list,"  ")
		   var entry = make(map[string]float64) 
           
		   temp,err := strconv.ParseFloat(fields[2],64)
		   if err != nil {
		     temp = 0
		   }	 
		   entry["cpu"] = temp
		   temp1,err := strconv.ParseFloat(fields[4],64)
		   if err != nil {
		     temp1 = 0
		   }	
		   entry["vsz"] = temp1
		   temp2,err := strconv.ParseFloat(fields[5],64)
		   if err != nil {
              temp2 = 0
		   }
		   entry["rsz"] = temp2
		   return_value[process_name] = entry
		   
	   }
	 }
	 return &return_value
}
   
    








func (docker_handle Docker_Handle_Type)container_status_keys() []string {

  var driver = (docker_handle).hash_status
  var return_value = driver.HKeys()
  
  return return_value

}

func (docker_handle Docker_Handle_Type) hdel_container_status_key(field string )  {
  var driver = (docker_handle).hash_status
  driver.HDel(field)


}

func (docker_handle Docker_Handle_Type) hget_status_value( field string ) *map[string]bool{


   var driver = (docker_handle).hash_status
   
   
   var v = msgpack_utils.Unpack(driver.HGet(field) )  

 
	
	var return_value = make(map[string]bool)
	for key, value := range v.(map[interface{}]interface{}) {
	   return_value[key.(string)] = value.(bool)
	}
	return &return_value

}




func (docker_handle Docker_Handle_Type) hset_status_values( field string, value *map[string]bool){

   // convert bool to msgpack
    var driver = (docker_handle).hash_status
 
    var b bytes.Buffer	
    msgpack.Pack(&b,*value)	
    driver.HSet(field, b.String())

}

func (docker_handle Docker_Handle_Type) add_incident_log(container string ,redis_container_status *map[string]bool , status bool ){

   var return_value = make(map[string]interface{})
   var driver = (docker_handle).incident_stream
   
   
   return_value["container"] = container
   for key,value := range *redis_container_status{
      return_value[key] = value
  }
  var b bytes.Buffer	
  msgpack.Pack(&b,return_value)	
  driver.Log_data(  b.String()  )
}
   
