package main

import "fmt"
import "lacima.com/go_application_containers/site_generation_base/site_generation_utilities"


const drive_path string = "--mount type=bind,source=/home/pi/mountpoint/lacuma_conf/site_config,target=/data/"
const file_path  string = "--mount type=bind,source=/home/pi/mountpoint/lacuma_conf/files/,target=/files/"
const command_start string = "docker run -d  --network host   --name"
const command_run   string = "docker run   -it --network host --rm  --name"



func main(){

  
  su.Setup_Site_File()
  fmt.Println(su.Site)
  su.Setup_graph_generation()
  setup_containers()
  setup_configuration()
   
}




func setup_containers(){

  su.Initialialize_container_data_structures(command_start,command_run)
  su.Add_mount("drive",drive_path)
  su.Add_mount("file",file_path)
  
  command_map := make(map[string]string)
  command_map["redis"] = "./redis_server ./redis.conf"
  su.Add_container( true,"redis","nanodatacenter/redis","./process_control",command_map, []string{"DATA"})
  
  command_map = make(map[string]string)
  command_map["generate_site"] = "./generate_site"  
  su.Add_container( false,"lacima_configuration", "nanodatacenter/lacima_configuration","./generate_site",command_map, []string{"DATA"})

  command_map = make(map[string]string)
  command_map["file_server"] = "./file_server"
  su.Add_container( true,"file_server","nanodatacenter/file_server","./process_control",command_map, []string{"DATA"})

  command_map = make(map[string]string)
  command_map["redis"] = "./redis_server ./redis.conf"
  su.Add_container( true,"managed_switch_logger","nanodatacenter/managed_switch_logger","./process_control",command_map, []string{"DATA"})

  command_map = make(map[string]string)
  command_map["redis_monitoring"] = "./redis_monitoring"
  su.Add_container( true,"monitor_redis","nanodatacenter/redis_monitoring","./process_control",command_map, []string{"DATA"})

}


func setup_configuration(){

 su.Start_site_definitions("LACIMA_SITE",[]string{"redis","lacima_configuration","file_server"})
 su.Construct_processor("irrigation_controller",[]string{"managed_switch_logger","monitor_redis"})
 su.End_site_definitions()
 su.Done()
 
}