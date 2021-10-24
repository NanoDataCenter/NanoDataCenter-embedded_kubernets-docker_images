package docker_management

import "fmt"
import "time"

import "strings"
import "strconv"
import "lacima.com/site_control_app/docker_control"

import "lacima.com/Patterns/msgpack_2"






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

   
   fmt.Println("starting container log ",time.Now().Unix())
   for _,container := range (docker_handle).containers {
       var working_values = (docker_handle).generate_parsed_fields(container)
	   if working_values != nil {
	     (*docker_handle).store_performance_data(container,"cpu", "cpu",working_values)
	     (*docker_handle).store_performance_data(container,"rss", "rss",working_values)
	     (*docker_handle).store_performance_data(container,"vsz", "vsz",working_values)
	   }
   }	  
  

}

func (docker_handle *Docker_Handle_Type) store_performance_data (tag2,tag3,data_key string, working_values *map[string]map[string]float64  ){


  tag1 := "CONTAINER"
  //fmt.Println("working_values",tag3,data_key)
  for process, data := range *working_values {
    data_element := data[data_key]
    //fmt.Println("data_element",data)
    packed_data  := msg_pack_utils.Pack_float64(data_element)
    //fmt.Println("postgress log",tag1,tag2,tag3,process,data_element,packed_data)
    (*docker_handle).logging_stream.Insert( tag1,tag2,tag3, process,"",packed_data )

  
  }
  
}



func (docker_handle *Docker_Handle_Type) generate_parsed_fields( container_name string) *map[string]map[string]float64{

  // ps headers headers = [ "USER","PID","%CPU","%MEM","VSZ","RSS","TTY","STAT","START","TIME","COMMAND", "PARAMETER1", "PARAMETER2" ]
  return_value := make(map[string]map[string]float64)
  cmd_string := "docker top "+container_name+ "  -aux "
  output := docker_control.System_shell(cmd_string)
  skip_lines := 1
  //fmt.Println("output",output)
  split_lines := strings.Split(output,"\n")
  
  if len(split_lines) <= skip_lines {
    return nil
  }
   
  process_lines := split_lines[skip_lines:]
  for _, data := range process_lines {
      if len(data) > 0 {
	       //fmt.Println("process_lines",process_lines)
  	       fields := strings.Fields(data)
           //fmt.Println("fields",fields)
		   name_list := fields[10:]
		   var process_name = strings.Join(name_list,":")
		   entry := make(map[string]float64) 
           
		   temp,err := strconv.ParseFloat(fields[2],64)
		   if err != nil {
		     panic("cpu error")
		   }	 
		   entry["cpu"] = temp
		   temp1,err := strconv.ParseFloat(fields[4],64)
		   if err != nil {
		     panic("vsz error")
		   }	
		   entry["vsz"] = temp1
		   temp2,err := strconv.ParseFloat(fields[5],64)
		   if err != nil {
              panic("rss error")
		   }
		   entry["rss"] = temp2
		   return_value[process_name] = entry
		   //fmt.Println("entry",entry)
	   }
	 }
	 //fmt.Println("return_value",return_value)
	 return &return_value
}
   
    








func (docker_handle Docker_Handle_Type)container_status_keys() []string {

  var driver = (docker_handle).hash_status
  var return_value = driver.HKeys()
  
  return return_value

}

func (docker_handle Docker_Handle_Type) hdel_container_status_key(field string )  {
  driver := (docker_handle).hash_status
  driver.HDel(field)


}

func (docker_handle Docker_Handle_Type) hget_status_value( field string ) *map[string]bool{


    driver := (docker_handle).hash_status
   
   
    v,err := msg_pack_utils.Unpack_map_string_bool(driver.HGet(field) )  
    if err != true {
        panic("error")
    }
    
	
    
	
	return &v

}




func (docker_handle Docker_Handle_Type) hset_status_values( field string, value *map[string]bool){

   // convert bool to msgpack
    var driver = (docker_handle).hash_status
 
   
    driver.HSet(field, msg_pack_utils.Pack_map_string_bool(*value))

}

func (docker_handle Docker_Handle_Type) add_incident_log(container string ,redis_container_status *map[string]bool , status bool ){

   
   var driver = (docker_handle).incident_stream
   
   
 
  driver.Log_data( msg_pack_utils.Pack_map_string_bool(*redis_container_status  ))
}
   
