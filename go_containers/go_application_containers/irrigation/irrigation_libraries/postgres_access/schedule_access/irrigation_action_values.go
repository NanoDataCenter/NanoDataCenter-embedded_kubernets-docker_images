package irr_sched_access


import (
       //"fmt"
      "strings"
      "lacima.com/redis_support/graph_query"
	//"lacima.com/redis_support/redis_handlers"
)


func Select_irrigation_action_data( server_key  string )([]string ,bool){
     key_list  :=  strings.Split(server_key,"~")
     server_type      := key_list[0]
     master_server := key_list[1]
     sub_server       := key_list[2]
     search_string := []string{"IRRIGATION_SERVER:"+master_server,"IRRIGATION_SUBSERVER:"+sub_server}
    if server_type == "true"{
       search_string = []string{"IRRIGATION_SERVER:"+master_server}
    } 
  
    properties := graph_query.Common_qs_search(&search_string)
    
    property := properties[0]
    return_value := graph_query.Convert_json_string_array( property["supported_actions"] )
    return return_value,true
    
    
}

/*
alve_group_json ='`+control_block.Valve_group_json +`'
    valve_group = JSON.parse(valve_group_json)
    valve_group.sort()
  
    valve_group_map_json = `+ control_block.Valve_group_map_json }`
*/
func Get_Valve_Group_data(  ){
   
   
     search_string := []string{"IRRIGATION_STATIONS:IRRIGATION_STATIONS","VALVE_GROUP_DEFS:VALVE_GROUP_DEFS"}
  
  
    properties := graph_query.Common_qs_search(&search_string)
    
    property := properties[0]
    control_block.Valve_group_data  = property


    
}
    
    
    
    

    
    
    
    
    
    
    
