package main

import "fmt"
import "lacima.com/go_setup_containers/site_generation_base/site_generation_utilities"
import "lacima.com/go_setup_containers/site_generation_base/system_definitions"

const drive_path string = "--mount type=bind,source=/home/pi/system_config/,target=/data/"  // path to get configuration data
const file_path  string = "--mount type=bind,source=/home/pi/mountpoint/files/,target=/files/"   // path for file server to get files
const redis_path  string = "--mount type=bind,source=/home/pi/mountpoint/redis/,target=/data/"  // path for redis server to store data
const secret_path string = "--mount type=bind,source=/home/pi/mountpoint/secrets/,target=/secrets/"


const command_start string = "docker run -d  --network host   --name"          // preambe script to start container
const command_run   string = "docker run   -it --network host --rm  --name"  // preamble script for container to run and exit




func main(){

  su.Setup_Site_File()
  
  su.Setup_graph_generation()
  setup_container_drives()
  setup_container_run_commands()
  fmt.Println("made it here")
  generate_site( "LACIMA_SITE" )
  /*
   * generate other sites
   * if needed
   */
  su.Done() //finalize graph
  
 
}



func generate_site( site_name string){
    
 
 
  add_processors()  // site dependent
  sys_defs.Build_System_Catalog()
  add_components()  // site dependent
  su.Construt_Site(site_name)    
    
}

func setup_container_run_commands(){
    
  su.Initialialize_container_data_structures(command_start,command_run)   
    
}

func setup_container_drives(){
  su.Setup_Mounts()  
  su.Add_mount("DATA",drive_path)
  su.Add_mount("FILE",file_path)
  su.Add_mount("REDIS_DATA",redis_path)    
  su.Add_mount("SECRETs",secret_path)
    
    
}

func add_processors(){
    
   su.Add_processor("irrigation_controller") 
    
    
}    

func add_components(){
    
    sys_defs.Add_Component(true,"","system_component") 
    sys_defs.Add_Component(false,"irrigation_controller", "tp_managed_switch")
    
}    
