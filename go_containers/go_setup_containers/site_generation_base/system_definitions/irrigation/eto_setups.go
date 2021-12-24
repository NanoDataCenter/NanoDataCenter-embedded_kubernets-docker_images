package irrigation

import (
    
    "lacima.com/go_setup_containers/site_generation_base/site_generation_utilities"

)



func Construct_eto_data(){
 
    
    su.Bc_Rec.Add_header_node("WEATHER_DATA_STRUCTURE","WEATHER_DATA_STRUCTURE",make(map[string]interface{}))
    
    su.Cd_Rec.Construct_package("WEATHER_DATA")
    su.Cd_Rec.Add_hash("ETO_CONTROL") 
    su.Cd_Rec.Add_hash("EXCEPTION_VALUES")
	su.Cd_Rec.Add_hash("RAIN_VALUES") // rain flag and stuff
    su.Cd_Rec.Add_hash("ETO_VALUES") 
    su.Cd_Rec.Add_hash("ETO_ACCUMULATION_TABLE")
	su.Cd_Rec.Create_postgres_stream( "ETO_HISTORY","admin","password","admin",2*365*24*3600)
    su.Cd_Rec.Create_postgres_stream( "RAIN_HISTORY","admin","password","admin",2*365*24*3600)
    su.Cd_Rec.Create_postgres_stream( "EXCEPTION_LOG","admin","password","admin",2*365*24*3600)
	su.Cd_Rec.Close_package_construction()
    
    su.Bc_Rec.End_header_node("WEATHER_DATA_STRUCTURE","WEATHER_DATA_STRUCTURE")
    
    su.Bc_Rec.Add_header_node("WEATHER_STATIONS","WEATHER_STATIONS",make(map[string]interface{}))
    add_station_cimis()
    add_station_cimis_satellite()
    add_station_messo_west_sruc1_eto()
    add_station_messo_west_sruc1_rain()
    add_station_wunderground()
    add_station_wunderground_hybrid()
    
    su.Bc_Rec.End_header_node("WEATHER_STATIONS","WEATHER_STATIONS")
    
    
    
    
}


func add_station_wunderground(){
    
 
  
  properties := make(map[string]interface{})
  properties["access_key"] = "WUNDERGROUND"
  properties["type"] = "WUNDERGROUND"
  properties["sub_id"] = "KCAMURRI101"
  properties["lat"] = float64(33.2)
  properties["alt"] = float64(2400.)
  properties["priority"] = 1000
  su.Bc_Rec.Add_info_node("WEATHER_STATION","WUNDERGROUND",properties)
}       

func add_station_cimis_satellite() {
    
       
    properties := make(map[string]interface{})
    properties["sub_id"]        = ""
    properties["access_key"] = "ETO_CIMIS_SATELLITE"
    properties["type"] = "CIMIS_SAT"
    properties["url"] = "http://et.water.ca.gov/api/data"
    properties["longitude"] = float64(-117.29945)
    properties["latitude"] = float64(33.578156)
    properties["priority"] = 4                 
    su.Bc_Rec.Add_info_node("WEATHER_STATION","ETO_CIMIS_SATELLITE",properties)
}

func add_station_cimis(){
    
 
    properties := make(map[string]interface{})
    properties["sub_id"]        = ""
    properties["access_key"] = "ETO_CIMIS"
    properties["type"] = "CIMIS"
    properties["url"] = "http://et.water.ca.gov/api/data"
    properties["station"] = "62"
    properties["priority"] = 2  
    su.Bc_Rec.Add_info_node("WEATHER_STATION","ETO_CIMIS",properties)
}


func add_station_messo_west_sruc1_eto() {

    properties := make(map[string]interface{})
    properties["sub_id"]        = ""
    properties["access_key"] = "MESSOWEST"
    properties["type"] = "MESSO_ETO"
    properties["url"] = "http://et.water.ca.gov/api/data"
    properties["station"] = "SRUC1"
    properties["altitude"] = float64(2400.)
    properties["latitude"] = float64(33.578156)
    properties["priority"] = 3  
    su.Bc_Rec.Add_info_node("WEATHER_STATION","SRUC1",properties)
}

func add_station_messo_west_sruc1_rain() {
 
    properties := make(map[string]interface{})
    properties["sub_id"]        = ""
    properties["access_key"] = "MESSOWEST"
    properties["type"] = "MESSO_RAIN"
    properties["url"]  = "http://et.water.ca.gov/api/data"
    properties["station"] = "SRUC1"
    properties["priority"] = 3  
    su.Bc_Rec.Add_info_node("WEATHER_STATION","SRUC1_RAIN",properties)
          
}       

       
func add_station_wunderground_hybrid(){
    
    properties := make(map[string]interface{})
    properties["sub_id"] = "KCAMURRI101"
    properties["type"] = "HYBRID_SETUP"
    properties["main"]  = "WUNDERGROUND"
    properties["solar"] = "SRUC1"
    properties["priority"] = 4  
    su.Bc_Rec.Add_info_node("WEATHER_STATION","WUNDERGROUND_HYBRID",properties)   
    
    
}
      
