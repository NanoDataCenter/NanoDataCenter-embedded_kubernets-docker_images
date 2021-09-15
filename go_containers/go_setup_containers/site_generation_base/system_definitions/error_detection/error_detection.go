package error_detection

import "lacima.com/go_setup_containers/site_generation_base/site_generation_utilities"



func Construct_definitions(){
    
  error_detection_properties := make(map[string]interface{})
  su.Bc_Rec.Add_header_node("ERROR_DETECTION","ERROR_DETECTION",error_detection_properties)  
    
  
  
  
  wd_detection_properties := make(map[string]interface{})
  wd_detection_properties["start_delay"]  = 60*5 // 5 minutes
  wd_detection_properties["trim_time"]    = 3600*24*30
  wd_detection_properties["wd_time"]      = 30  // 30 seconds
  
  
  su.Bc_Rec.Add_header_node("WD_DETECTION","WD_DETECTION",wd_detection_properties)
  su.Cd_Rec.Construct_package("WATCH_DOG_DATA")
  su.Cd_Rec.Add_hash("WATCH_DOG_VALUE")   // a full length topic and a marshalled data value
  su.Cd_Rec.Add_hash("WATCH_DOG_STAMP") // a full length topic and a unix time in seconds as a string
  su.Cd_Rec.Add_hash("STATE_CHANGE_COUNTS")  
  su.Cd_Rec.Create_postgres_stream( "WATCH_DOG_INCIDENTS","admin","password","admin",30*24*3600)  
  su.Cd_Rec.Close_package_construction()
  
  su.Bc_Rec.End_header_node("WD_DETECTION","WD_DETECTION")  
  
  
  
  su.Bc_Rec.End_header_node("ERROR_DETECTION","ERROR_DETECTION")  
    
    
    
    
    
}
