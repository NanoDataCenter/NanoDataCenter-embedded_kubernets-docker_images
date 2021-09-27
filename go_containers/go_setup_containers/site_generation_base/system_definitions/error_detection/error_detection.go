package error_detection

import "lacima.com/go_setup_containers/site_generation_base/site_generation_utilities"



func Construct_definitions(){
    
  error_detection_properties := make(map[string]interface{})
  error_detection_properties["subsystems"] = []string{"watch_dog","incident","rpc","streaming"}
  su.Bc_Rec.Add_header_node("ERROR_DETECTION","ERROR_DETECTION",error_detection_properties)  
    
  
       su.Cd_Rec.Construct_package("OVERALL_STATUS")
           su.Cd_Rec.Add_hash("OVERALL_STATUS")
       su.Cd_Rec.Close_package_construction()
   
  
  
       wd_detection_properties := make(map[string]interface{})
       wd_detection_properties["trim_time"]       = 3600*24*30
       wd_detection_properties["sample_time"]     = 10  // 30 seconds
       wd_detection_properties["debounce_count"]       = 5
       wd_detection_properties["subsystem_id"]    = "watch_dog"
       su.Bc_Rec.Add_header_node("WD_DETECTION","WD_DETECTION",wd_detection_properties)
  
 
            su.Cd_Rec.Construct_package("WATCH_DOG_DATA")
                su.Cd_Rec.Add_hash("DEBOUNCED_STATUS")   
                su.Cd_Rec.Add_hash("STATUS")   
                su.Cd_Rec.Add_hash("TIME_STAMP") 
                su.Cd_Rec.Create_postgres_stream( "WATCH_DOG_LOG","admin","password","admin",30*24*3600)  
            su.Cd_Rec.Close_package_construction()
  
       su.Bc_Rec.End_header_node("WD_DETECTION","WD_DETECTION")  
  
  
       incident_properties := make(map[string]interface{})
       incident_properties["trim_time"]       = 3600*24*30
       incident_properties["sample_time"]     = 15  // 30 seconds
       incident_properties["subsystem_id"]     = "incident"
       
       su.Bc_Rec.Add_header_node("INCIDENT_STREAMS","INCIDENT_STREAMS",incident_properties)
  
           su.Cd_Rec.Construct_package("INCIDENT_DATA")
               su.Cd_Rec.Add_hash("TIME") 
               su.Cd_Rec.Add_hash("STATUS") 
               su.Cd_Rec.Add_hash("LAST_ERROR")   
               su.Cd_Rec.Create_postgres_stream( "INCIDENT_LOG","admin","password","admin",30*24*3600)  
           su.Cd_Rec.Close_package_construction()
  
       su.Bc_Rec .End_header_node("INCIDENT_STREAMS","INCIDENT_STREAMS")  
       
       streaming_properties := make(map[string]interface{})
       streaming_properties["sample_time"]     = 15  // 30 seconds
       
       su.Bc_Rec.Add_header_node("STREAMING_LOGS","STREAMING_LOGS",streaming_properties)
  
        
  
       su.Bc_Rec .End_header_node("STREAMING_LOGS","STREAMING_LOGS")    
  
       rpc_properties := make(map[string]interface{})
       
       rpc_properties["sample_time"]     = 15  // 15 minutes
       
       
       su.Bc_Rec.Add_header_node("RPC_ANALYSIS","RPC_ANALYSIS",rpc_properties)
  
       su.Construct_incident_logging("RPC_ANALYSIS" ,"RPC FAILURES")
       
         
  
       su.Bc_Rec.End_header_node("RPC_ANALYSIS","RPC_ANALYSIS")    
       
       
   su.Bc_Rec.End_header_node("ERROR_DETECTION","ERROR_DETECTION")  
    
    
    
    
    
}
