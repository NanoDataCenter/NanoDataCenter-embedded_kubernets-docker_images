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
             su.Cd_Rec.Construct_package("STREAM_SUMMARY_DATA")           
               su.Cd_Rec.Add_hash("WORKING_TABLE") 
               su.Cd_Rec.Add_hash("TIME_TABLE") 
               su.Cd_Rec.Add_hash("STATUS_TABLE")
               su.Cd_Rec.Add_hash("ERROR_TABLE") 
               su.Cd_Rec.Add_hash("ERROR_TIME")
               su.Cd_Rec.Create_postgres_stream( "DATA_STREAM","admin","password","admin",30*24*3600)  
               su.Cd_Rec.Create_postgres_stream( "INCIDENT_STREAM","admin","password","admin",30*24*3600)  
             su.Cd_Rec.Close_package_construction()
  
    
       su.Bc_Rec .End_header_node("STREAMING_LOGS","STREAMING_LOGS")    
  
       rpc_properties := make(map[string]interface{})
       
       rpc_properties["sample_time"]     = 15  // 15 minutes
       
       
       su.Bc_Rec.Add_header_node("RPC_ANALYSIS","RPC_ANALYSIS",rpc_properties)
  
       su.Construct_incident_logging("RPC_ANALYSIS" ,"RPC FAILURES",su.Emergency)
       
         
  
       su.Bc_Rec.End_header_node("RPC_ANALYSIS","RPC_ANALYSIS")    
       
       alert_properties := make(map[string]interface{})
       su.Bc_Rec.Add_header_node("ALERT_NOTIFICATION","ALERT_NOTIFICATION",alert_properties)
       
           telegram_properties := make(map[string]interface{})
           telegram_properties["valid_users"] = []string{"1575166855"}    
           su.Bc_Rec.Add_header_node("TELEGRAM_SERVER","TELEGRAM_SERVER",telegram_properties)
              su.Construct_RPC_Server("TELEGRAPH_RPC","rpc for controlling system",10,15, make( map[string]interface{}) )
           su.Bc_Rec.End_header_node("TELEGRAM_SERVER","TELEGRAM_SERVER")    
       
           
        su.Bc_Rec.End_header_node("ALERT_NOTIFICATION","ALERT_NOTIFICATION")   
       
       
       
   su.Bc_Rec.End_header_node("ERROR_DETECTION","ERROR_DETECTION")  
    
    
    
 
    
}
