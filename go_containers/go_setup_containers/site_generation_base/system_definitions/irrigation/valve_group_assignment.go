package irrigation

import (
    "lacima.com/go_setup_containers/site_generation_base/site_generation_utilities"
)


var io_description_map  map[string]map[int64]string


func Add_valve_group_definitions(){
  
 io_description_map    =  make(map[string]map[int64]string)
  
  properties                             := make(map[string]interface{})
  data                                       := make([][]map[string]interface{},10)
  names                                   := make([]string,10)
  description                           := make([]string,10)
   data [0]  ,  names[0]   ,  description[0]  = add_valve_group_1( )
   data [1]  ,  names[1]   ,  description[1]  = add_valve_group_2( )
   data [2]  ,  names[2]   ,  description[2]  = add_valve_group_3( )
    data [3]  ,  names[3]   ,  description[3]  = add_valve_group_4( )
    data [4]  ,  names[4]   ,  description[4]  = add_valve_group_5( )
    data [5]  ,  names[5]   ,  description[5]  = add_valve_group_6( )
    data [6]  ,  names[6]   ,  description[6]  = add_valve_group_7( )
    data [7]  ,  names[7]   ,  description[7]  = add_valve_group_8( )
    data [8]  ,  names[8]   ,  description[8]  = add_valve_group_9( )
   data [9]  ,  names[9]   ,  description[9]  = add_valve_group_10( )
   
   properties["data"]               = data
   properties["names"]            = names
   properties["descriptions"]  = description
   properties["io_map"]           = io_description_map

  su.Bc_Rec.Add_info_node("VALVE_GROUP_DEFS","VALVE_GROUP_DEFS",properties)
}    

func add_valve_group_entry( name ,controller string, io int64)map[string]interface{}{
 
    return_value := make(map[string]interface{})
    return_value["name"]           = name
    return_value["controller"]   = controller
    return_value["io"]                = io
    if _, ok := io_description_map[controller] ; ok == false {
         temp := make(map[int64]string)
         io_description_map[controller] = temp
    }
    io_description_map[controller][io]    =  name
    
    return return_value
}




func add_valve_group_1()( []map[string]interface{},string,string ){
    name            :=  "valve group 1"
    description  := "xxxxxxxxxxxxxxx"
    return_value := make([]map[string]interface{}, 7 )
    return_value[0] = add_valve_group_entry(   "Flowers along side walk" ,"satellite_1",20)
    return_value[1] = add_valve_group_entry(  "Well Water Drip Line","satellite_1",19)
    return_value[2] = add_valve_group_entry(  "Barbecue Clover Area","satellite_1",14)
    return_value[3] = add_valve_group_entry(  "Well Clover Area","satellite_1",18)
    return_value[4] = add_valve_group_entry(  "Triangle Pool Area","satellite_1",15)
    return_value[5] = add_valve_group_entry(  "Dragon Fruit — Fruit Tree Drip Line","satellite_1",17)
    return_value[6] = add_valve_group_entry(  "Pool Fence Area","satellite_1",16)
    
   return return_value, name , description
}

func add_valve_group_2()( []map[string]interface{},string,string ){
    name            :=  "valve group 2"
     description  := "xxxxxxxxxxxxxxx"
    return_value := make([]map[string]interface{}, 4 )
    return_value[0] = add_valve_group_entry(   "Lemon Tree Drip Line near Steps" ,"satellite_1",13)
    return_value[1] = add_valve_group_entry(  "Middle Clover Near Well" ,"satellite_1",25)
    return_value[2] = add_valve_group_entry(  "Middle Clover Near Barbecue","satellite_1",12)
    return_value[3] = add_valve_group_entry(  "Drip Line along garage","satellite_1",11)
    
   return return_value, name , description
}

func add_valve_group_3()( []map[string]interface{},string,string ){
    name            :=  "valve group 3"
     description  := "xxxxxxxxxxxxxxx"
    return_value := make([]map[string]interface{}, 4 )
    return_value[0] = add_valve_group_entry(    "Flowers Toward Garage" ,"satellite_1",24)
    return_value[1] = add_valve_group_entry(  "Flowers on Opposite Side of Garage" ,"satellite_1",21)
    return_value[2] = add_valve_group_entry(   "Grass Zone Away From Door","satellite_1",22)
    return_value[3] = add_valve_group_entry(  "Grass Toward Door","satellite_1",23)
    
  return return_value, name , description
}
    

func add_valve_group_4()( []map[string]interface{},string,string ){
    name            :=  "valve group 4"
     description  := "xxxxxxxxxxxxxxx"
    return_value := make([]map[string]interface{}, 7 )
    return_value[0] = add_valve_group_entry(    "Upper area Near Property Line"  ,   "satellite_2"     ,2)
    return_value[1] = add_valve_group_entry(   "On Property Side of Valve #3"     ,    "satellite_2"     ,4)
    return_value[2] = add_valve_group_entry(   "Next to Valve 4"                             ,     "satellite_2"    ,3)
    return_value[3] = add_valve_group_entry(   "Spray Area Next to Remote"       ,     "satellite_2"    ,1)
    return_value[4] = add_valve_group_entry(    "Sprayers along Drive Way"         ,      "satellite_2"   ,5)
    return_value[5] = add_valve_group_entry(    "Lower area near Property Line" ,      "satellite_2"   ,6)    
     return_value[6] = add_valve_group_entry(    "Drip Line along bank" ,                      "satellite_2"   ,7)    
   return return_value, name , description
}
 

func add_valve_group_5()( []map[string]interface{},string,string ){
    name            :=  "valve group 5"
     description  := "xxxxxxxxxxxxxxx"
    return_value := make([]map[string]interface{}, 4 )  
    return_value[0] = add_valve_group_entry(    "Fruit Trees on Block #9"  ,                    "satellite_2"     ,2)
    return_value[1] = add_valve_group_entry(   "Waters Bank Closest to House"     ,    "satellite_2"     ,4)
    return_value[2] = add_valve_group_entry(   "Sprinkler on Bank"                             ,    "satellite_2"    ,3)
    return_value[3] = add_valve_group_entry(   "Avocado Block6"       ,                            "satellite_2"    ,1)
  return return_value, name , description 
}
 
  

func add_valve_group_6()( []map[string]interface{},string,string ){
    name            :=  "valve group 6"
     description  := "xxxxxxxxxxxxxxx"
    return_value := make([]map[string]interface{}, 7 )
    return_value[0] = add_valve_group_entry(     "Avocado Block #4 Top",                             "satellite_2"     ,13)
    return_value[1] = add_valve_group_entry(  "Baby Tress"     ,                                                 "satellite_2"     ,14)
    return_value[2] = add_valve_group_entry(  "Avocado Block 4 Bottom",                             "satellite_2"    ,15)
    return_value[3] = add_valve_group_entry(    "Drip Line and Sprayers along Road"   ,     "satellite_2"    ,17)
    return_value[4] = add_valve_group_entry(    "Block 3 Androus Site",                                 "satellite_2"   ,16)
    return_value[5] = add_valve_group_entry(     "No Connection",                                           "satellite_2"   , 20)    
     return_value[6] = add_valve_group_entry(     "No Connection",                                          "satellite_2"   ,21 )    
   return return_value, name , description 
}
 

    

func add_valve_group_7()( []map[string]interface{},string,string ){
    name            :=  "valve group 7"
     description  := "xxxxxxxxxxxxxxx"
    return_value := make([]map[string]interface{}, 6 )
    return_value[0] = add_valve_group_entry(     "????????????????",                             "satellite_3"   ,4)
    return_value[1] = add_valve_group_entry(       "????????????"     ,                               "satellite_3"   ,3)
    return_value[2] = add_valve_group_entry(       "Avocado Block 5",                             "satellite_3"   ,1)
    return_value[3] = add_valve_group_entry(       "????????????????"   ,                          "satellite_3"   ,6)
    return_value[4] = add_valve_group_entry(       "Baby Tres Cindy Side",                     "satellite_3"   ,5)
    return_value[5] = add_valve_group_entry(     "???????????????????",                          "satellite_3"   ,2)    
   return return_value, name , description
}
     
    
func add_valve_group_8()( []map[string]interface{},string,string ){
    name            :=  "valve group 8"
     description  := "xxxxxxxxxxxxxxx"
    return_value := make([]map[string]interface{}, 7 )    
     return_value[0] = add_valve_group_entry(     "Block 1",                                    "satellite_3"   ,13)
    return_value[1] = add_valve_group_entry(       "Block 2"     ,                               "satellite_3"   ,14)
    return_value[2] = add_valve_group_entry(       "Block 1.5" ,                                 "satellite_3"   ,18)
    return_value[3] = add_valve_group_entry(       "????????????????"   ,                 "satellite_3"   ,17)
    return_value[4] = add_valve_group_entry(      "Baby Tres Cindy Side",             "satellite_3"   ,21)
    return_value[5] = add_valve_group_entry(     "Block 3 Cindy Side",                   "satellite_3"   ,15)    
    return_value[6] = add_valve_group_entry(     "???????????????????",                 "satellite_3"   ,19)    
    return return_value, name , description
}
      

     
func add_valve_group_9()( []map[string]interface{},string,string ){
    name            :=  "valve group 9"
     description  := "xxxxxxxxxxxxxxx"
    return_value := make([]map[string]interface{}, 7 )
    return_value[0] = add_valve_group_entry(       "Bottom Hill on Cindy's House Side",                                    "satellite_3"   ,11)
    return_value[1] = add_valve_group_entry(       "Bottom Hill on the opposite site of drive way"    ,              "satellite_3"   ,12)
    return return_value, name , description
}

func add_valve_group_10()( []map[string]interface{},string,string ){
    name            :=  "valve group 10"
     description  := "xxxxxxxxxxxxxxx"
    return_value := make([]map[string]interface{}, 7 )
    return_value[0] = add_valve_group_entry(     "??????????????????",                 "satellite_4"   ,1)
    return_value[1] = add_valve_group_entry(      "?????????????"     ,                     "satellite_4"   ,2)
    return_value[2] = add_valve_group_entry(      "?????????????"     ,                    "satellite_4"   ,3)
    return_value[3] = add_valve_group_entry(       "????????????????"   ,                 "satellite_4"   ,4)
    return_value[4] = add_valve_group_entry(      "?????????????????",                   "satellite_4"   ,5)
    return_value[5] = add_valve_group_entry(     "??????????????????",                   "satellite_4"   ,6)    
    return_value[6] = add_valve_group_entry(     "???????????????????",                 "satellite_4"   ,7)    
    return return_value, name , description
}
    
    
    
