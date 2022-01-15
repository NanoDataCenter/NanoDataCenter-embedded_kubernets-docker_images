package irrigation

import (
    "lacima.com/go_setup_containers/site_generation_base/site_generation_utilities"
)




func Add_irrigation_stations_definitions(){
    
    su.Bc_Rec.Add_header_node("IRRIGATION_GROUPS","IRRIGATION_GROUPS",make(map[string]interface{}))
    su.Bc_Rec.Add_header_node("IRRIGATION_GROUP","main_location",make(map[string]interface{}))
    
    station_1 := make(map[string]interface{})
    station_1["valves"] = generate_interger_range(1,56,make([]int64,0))
    su.Bc_Rec.Add_info_node("IRRIGATION_STATION","station_1",station_1 )

    station_2 := make(map[string]interface{})
    station_2["valves"] = generate_interger_range(1,22,make([]int64,0))
    su.Bc_Rec.Add_info_node("IRRIGATION_STATION","station_2",station_2 )

    station_3 := make(map[string]interface{})
    station_3["valves"] = generate_interger_range(1,22,make([]int64,0))
    su.Bc_Rec.Add_info_node("IRRIGATION_STATION","station_3",station_3 )

    station_4 := make(map[string]interface{})
    station_4["valves"] = generate_interger_range(1,20,make([]int64,0))
    su.Bc_Rec.Add_info_node("IRRIGATION_STATION","station_4",station_4 )

    su.Bc_Rec.End_header_node("IRRIGATION_GROUP","main_location")
    su.Bc_Rec.End_header_node("IRRIGATION_GROUPS","IRRIGATION_GROUPS")
}


func Eto_valve_setup(){

    su.Bc_Rec.Add_header_node("ETO_SETUP","ETO_SETUP",make(map[string]interface{}))

    eto_info := make(map[string]interface{})
    su.Bc_Rec.Add_info_node("ETO_SETUP_PROPERTIES","ETO_SETUP_PROPERTIES",eto_info )    
    su.Cd_Rec.Construct_package("ETO_DATA_STRUCTURES")
    su.Cd_Rec.Add_hash("ETO_ACCUMULATION") 
    su.Cd_Rec.Add_hash("ETO_RESERVE")
    su.Cd_Rec.Add_hash("ETO_MIN_LEVEL")
    su.Cd_Rec.Add_hash("ETO_RECHARGE_RATE")
	su.Cd_Rec.Close_package_construction()    
    
    su.Bc_Rec.End_header_node("ETO_SETUP","ETO_SETUP")    
    
    
    
}




func generate_interger_range(start,end int64, exclusion_array []int64 )[]int64{
    
      return_value := make([]int64,0)
      
      for i := start; i<=end;i++ {
          if check_exclusion(i,exclusion_array) == true {
              return_value = append(return_value,i)
          }
      }
      return return_value    
    
}


func check_exclusion( value int64, exclusion_array []int64 )bool{
    
    for _,check := range exclusion_array {
        if value != check {
            return false
        }
    }
    return true
}

