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
    su.Cd_Rec.Add_hash("DATA_MAP") // map of site data
    su.Cd_Rec.Close_package_contruction()
    
    su.Cd_Rec.Construct_package("NODE_MAP")
    su.Cd_Rec.Add_hash("NODE_MAP") // map of node ip's
    su.Cd_Rec.Close_package_contruction()
    
    su.Cd_Rec.Construct_package("WEB_MAP")
    su.Cd_Rec.Add_hash("WEB_MAP") // map of all subsystem web servers
    su.Cd_Rec.Close_package_contruction()    
    
    su.Cd_Rec.Construct_package("SITE_CONTROL")
    su.Cd_Rec.Add_job_queue("WEB_COMMAND_QUEUE",10) // commands such as reboot pull container
    su.Cd_Rec.Add_redis_stream("ERROR_STREAM",1024)
    su.Cd_Rec.Add_hash("ERROR_HASH")
    su.Cd_Rec.Add_hash("WEB_DISPLAY_DICTIONARY")  //for displaying node status
    su.Cd_Rec.Close_package_contruction()
    
    su.Cd_Rec.Construct_package("DOCKER_CONTROL")
    su.Cd_Rec.Add_job_queue("DOCKER_COMMAND_QUEUE",10) //temp disable turning of containers
    su.Cd_Rec.Add_hash("DOCKER_DISPLAY_DICTIONARY")
    su.Cd_Rec.Add_redis_stream("ERROR_STREAM",1024)
    su.Cd_Rec.Close_package_contruction()

    su.Cd_Rec.Construct_package("NODE_MONITORING")
    su.Cd_Rec.Add_job_queue("WEB_COMMAND_QUEUE",1)
    su.Cd_Rec.Add_hash("NODE_STATUS")
    su.Cd_Rec.Add_redis_stream("ERROR_ STREAM",1024)
    su.Cd_Rec.Add_hash("SYSTEM_CONTAINER_IMAGES") //value list of nodes container is in
    su.Cd_Rec.Close_package_contruction()
   
    su.Cd_Rec.Construct_package("SITE_REBOOT_LOG")
    su.Cd_Rec.Add_redis_stream("SITE_REBOOT_LOG",1024)
    su.Cd_Rec.Close_package_contruction()
    
    
    su.Bc_Rec.Add_header_node("REDIS_MONITORING","REDIS_MONITORING",make(map[string]interface{}))
    su.Construct_streaming_logs("redis_monitor",[]string{"STREAMING_LOG","KEYS","CLIENTS","MEMORY","REDIS_MONITOR_CMD_TIME_STREAM"})    
    su.Bc_Rec.End_header_node("REDIS_MONITORING","REDIS_MONITORING")
   
    file_server_properties := make(map[string]interface{})
    file_server_properties["directory"] = "/files"
    su.Bc_Rec.Add_header_node("FILE_SERVER","FILE_SERVER",file_server_properties)
    su.Cd_Rec.Construct_package("FILE_SERVER")
    su.Cd_Rec.Add_rpc_server("FILE_SERVER_RPC_SERVER",30,10)
    su.Cd_Rec.Close_package_contruction()
    su.Bc_Rec.End_header_node("FILE_SERVER","FILE_SERVER")    
    
    
}
