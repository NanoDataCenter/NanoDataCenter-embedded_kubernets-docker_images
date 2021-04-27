package main

import "fmt"
import "lacima.com/go_application_containers/site_generation_base/site_generation_utilities"


const drive_path string = "--mount type=bind,source=/home/pi/mountpoint/lacuma_conf/site_config,target=/data/"
const file_path  string = "--mount type=bind,source=/home/pi/mountpoint/lacuma_conf/files/,target=/files/"
const redis_path  string = "--mount type=bind,source=/home/pi/mountpoint/lacuma_conf/redis/,target=/data/"
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
  su.Add_mount("DATA",drive_path)
  su.Add_mount("FILE",file_path)
  su.Add_mount("REDIS_DATA",redis_path)
  
  command_map := make(map[string]string)
  command_map["redis"] = "./redis_server ./redis.conf"
  su.Add_container( true,"redis","nanodatacenter/redis","./redis_control.bsh",command_map, []string{"REDIS_DATA"})
  
  

  command_map = make(map[string]string)
  command_map["file_server"] = "./file_server"
  su.Add_container( true,"file_server","nanodatacenter/file_server","./process_control",command_map, []string{"DATA","FILE"})

  command_map = make(map[string]string)
  command_map["manage_switch_logger"] = "./manage_switch_logger"
  su.Add_container( true,"managed_switch_logger","nanodatacenter/managed_switch_logger","./process_control",command_map, []string{"DATA"})

  command_map = make(map[string]string)
  command_map["redis_monitoring"] = "./redis_monitoring"
  su.Add_container( true,"redis_monitoring","nanodatacenter/redis_monitoring","./process_control",command_map, []string{"DATA"})

}


func setup_configuration(){
 startup_containers := []string{ "nanodatacenter/lacima_configuratiion","nanodatacenter/lacima_setups"}
 irrigation_containers := []string{"managed_switch_logger","redis_monitoring"}
 system_containers := []string{"redis","file_server"}
 su.Start_site_definitions("LACIMA_SITE",system_containers,startup_containers)
 construct_site_specific_definitions()
 su.Construct_processor("irrigation_controller", irrigation_containers)
 
 su.End_site_definitions()
 su.Done()
 
}



func construct_site_specific_definitions(){ 
   
    properties := make(map[string]interface{})
    properties["ip"] = "192.168.1.45"
 
    su.Bc_Rec.Add_header_node("TP_SWITCH","switch_office",properties)
	su.Construct_incident_logging("switch_office")
    su.Bc_Rec.End_header_node("TP_SWITCH","switch_office")

    properties = make(map[string]interface{})
    properties["ip"] = "192.168.1.56"
    su.Bc_Rec.Add_header_node("TP_SWITCH","switch_garage",properties)
    su.Construct_incident_logging("switch_garage")
    su.Bc_Rec.End_header_node("TP_SWITCH","switch_garage")
    
}    