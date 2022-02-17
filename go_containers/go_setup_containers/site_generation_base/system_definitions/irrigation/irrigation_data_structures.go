package irrigation

import (
    
    "lacima.com/go_setup_containers/site_generation_base/site_generation_utilities"

)




func Add_irrigation_data_structures(){
    
    su.Bc_Rec.Add_header_node("IRRIGATION_DATA_STRUCTURES","IRRIGATION_DATA_STRUCTURES",make(map[string]interface{}))
    
    su.Bc_Rec.Add_header_node("SCHEDULE_DATA","SCHEDULE_DATA",make(map[string]interface{}))
    
    su.Cd_Rec.Construct_package("IRRIGATION_DATA")
	su.Cd_Rec.Create_postgres_table( "IRRIGATION_SCHEDULES","admin","password","admin")
    su.Cd_Rec.Create_postgres_float( "IRRIGATION_ACTIONS","admin","password","admin")
	su.Cd_Rec.Close_package_construction()
    
    su.Bc_Rec.End_header_node("SCHEDULE_DATA","SCHEDULE_DATA")
    
    
    su.Bc_Rec.End_header_node("IRRIGATION_DATA_STRUCTURES","IRRIGATION_DATA_STRUCTURES")
    
     
    
    
}
