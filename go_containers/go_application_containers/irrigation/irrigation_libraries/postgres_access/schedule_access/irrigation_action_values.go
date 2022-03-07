package irr_sched_access


import (
      //"fmt"
      "lacima.com/redis_support/graph_query"
	//"lacima.com/redis_support/redis_handlers"
)


func Select_irrigation_action_data( server_type bool,master_server,sub_server string)([]string ,bool){
     search_string := []string{"IRRIGATION_SERVER:"+master_server,"IRRIGATION_SUBSERVER:"+sub_server}
    if server_type == true{
       search_string = []string{"IRRIGATION_SERVER:"+master_server}
    } 
  
    properties := graph_query.Common_qs_search(&search_string)
    
    property := properties[0]
    return_value := graph_query.Convert_json_string_array( property["supported_actions"] )
    return return_value,true
    
    
}
