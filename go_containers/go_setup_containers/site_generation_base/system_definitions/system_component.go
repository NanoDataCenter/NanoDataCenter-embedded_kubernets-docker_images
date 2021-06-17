package sys_defs
import "lacima.com/go_setup_containers/site_generation_base/site_generation_utilities"

const redis_image             string   = "nanodatacenter/redis"
const lacima_site_image       string   = "nanodatacenter/lacima_site_generation"
const lacima_secrets_image    string   = "nanodatacenter/lacima_secrets"
const file_server_image       string   = "nanodatacenter/file_server"
const redis_monitor_image     string   = "nanodatacenter/redis_monitoring"






func generate_system_components(system_flag bool, processor_name string ){
   file_server_mount := []string {"DATA","FILE"}
   redis_mount       := []string{"REDIS_DATA"}
   secrets_mount     := []string{"DATA","SECRETS"}

   
   redis_monitor_command_map  := make(map[string]string)
   redis_monitor_command_map["redis_monitor"] = "./redis_monitor"

   file_server_command_map  := make(map[string]string)
   file_server_command_map["file_server"] = "./file_server"   
   
   
   
    
   null_map := make(map[string]string)
   su.Add_container( false,"redis", redis_image, "./redis_control.bsh", null_map,redis_mount)
   su.Add_container( true, "lacima_site_generation",lacima_site_image,su.Temp_run ,null_map, su.Data_mount )
   su.Add_container( true, "lacima_secrets",lacima_secrets_image,su.Temp_run ,null_map, secrets_mount)
   su.Add_container( false, "file_server",file_server_image, su.Managed_run ,file_server_command_map ,file_server_mount)
   su.Add_container( false,"redis_monitor",redis_monitor_image, su.Managed_run,redis_monitor_command_map, su.Data_mount)
   
   containers := []string{"redis","lacima_secrets","file_server","redis_monitor"}
   su.Construct_system_def("system_monitoring",true,"", containers, generate_system_component_graph) 
    
    
}


func generate_system_component_graph(){
    su.Cd_Rec.Construct_package("DATA_MAP")
    su.Cd_Rec.Add_single_element("DATA_MAP") // map of site data
    su.Cd_Rec.Close_package_contruction()
    
    su.Construct_incident_logging("SITE_REBOOT","site_reboot")
    
    su.Cd_Rec.Construct_package("REBOOT_FLAG")
    su.Cd_Rec.Add_single_element("REBOOT_FLAG") // determine if site has done all initialization
    su.Cd_Rec.Close_package_contruction()
    
    
    su.Cd_Rec.Construct_package("NODE_MAP")
    su.Cd_Rec.Add_hash("NODE_MAP") // map of node ip's
    su.Cd_Rec.Close_package_contruction()
    
    su.Cd_Rec.Construct_package("WEB_MAP")
    su.Cd_Rec.Add_hash("WEB_MAP") // map of all subsystem web servers
    su.Cd_Rec.Close_package_contruction()    
    
    
    
    
    
    su.Construct_RPC_Server("SYSTEM_CONTROL","rpc for controlling system",10,15, make( map[string]interface{}) )

    su.Cd_Rec.Construct_package("NODE_STATUS")
    su.Cd_Rec.Add_hash("NODE_STATUS")
    su.Cd_Rec.Close_package_contruction()
   
    
    
    su.Bc_Rec.Add_header_node("REDIS_MONITORING","REDIS_MONITORING",make(map[string]interface{}))
    su.Construct_streaming_logs("redis_monitor","redis_monitor",[]string{"KEYS","CLIENTS","MEMORY","REDIS_MONITOR_CMD_TIME_STREAM"})   
    su.Bc_Rec.End_header_node("REDIS_MONITORING","REDIS_MONITORING")
   
    file_server_properties := make(map[string]interface{})
    file_server_properties["directory"] = "/files"
    su.Construct_RPC_Server( "SITE_FILE_SERVER","site_file_server",30,10,file_server_properties)
    
    
 
    
    
}
