package su

//import "fmt"
import "strings"

type container_descriptor struct {
    docker_image   string
    command_string string


}


var drive_mounts map[string]string

var container_map map[string]container_descriptor

var command_string_first_part string  // continually execute container
var command_string_run_part string   // container executes a script and terminates

func Initialialize_container_data_structures(start_part,run_part string){
   drive_mounts =  make(map[string]string)
   container_map = make(map[string]container_descriptor)
   command_string_first_part = start_part
   command_string_run_part =  run_part
}


func Add_mount( mount_name string , mount_path string ){
 
   if _,ok := drive_mounts[mount_name]; ok == true {
     panic("duplicate mount name "+mount_name)
   }

   drive_mounts[mount_name] = mount_path


}

func Add_container( temp_flag bool, container_name, docker_image, command_string string , mounts []string){
   
   if _,ok := container_map[container_name]; ok == true {
     panic("duplicate container name "+container_name)
   }

   var temp container_descriptor
   
   
   temp.docker_image = docker_image
   if temp_flag == false {
        temp.command_string = command_string_first_part+"  "+container_name+"  "+strings.Join(mounts,"  ")+" "+docker_image+" "+command_string
   }else{
      temp.command_string = command_string_run_part+"  "+container_name+"  "+strings.Join(mounts,"  ")+" "+docker_image+" "+command_string
   }
   container_map[container_name] = temp
   
}


func register_containers( container_list []string ){
 
   for _,container_name := range container_list {
      
       register_container(container_name)
   
   
   }

}


func register_container( container_name string){
   if _,ok := container_map[container_name]; ok == false{
      panic("container does not exist  "+container_name)
   }

   properties := make(map[string]interface{})
   properties["container_image"] = container_map[container_name].docker_image
   properties["startup_command"] = container_map[container_name].command_string
   Bc_Rec.Add_header_node("CONTAINER",container_name,properties)
  
   Cd_Rec.Construct_package("DATA_STRUCTURES")
   Cd_Rec.Add_single_element("controller_watchdog")
   Cd_Rec.Add_redis_stream("CONTROLLER_FAILURE",1024)  // container process_control failure
   Cd_Rec.Add_hash("WEB_DISPLAY_DICTIONARY") // state of process
   Cd_Rec.Add_hash("Process_Status")  // last error
   Cd_Rec.Add_redis_stream("Process_Failure",1024) // error stream of different errors
   Cd_Rec.Add_redis_stream("ERROR_STREAM",1024)  //not sure of what this is 
   Cd_Rec.Add_redis_stream("PROCESS_VSZ",1024)
   Cd_Rec.Add_redis_stream("PROCESS_RSS",1024)
   Cd_Rec.Add_redis_stream("PROCESS_CPU",1024) 
   Cd_Rec.Close_package_contruction()
 
   Bc_Rec.End_header_node("CONTAINER",container_name)
}



