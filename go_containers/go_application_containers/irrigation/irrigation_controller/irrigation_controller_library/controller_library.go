package irrigation_controller_library

import (
      // "fmt"
      "strings"
    "strconv"
    "encoding/json"
      "lacima.com/redis_support/graph_query"
	
)


type Registry_type struct{
 
  Master_table_list                 map[string][]string
  Valve_list                               map[string]map[string]map[string][]int
  Inverse_Valve_Map             map[string]string
  Actions                                 map[string]map[string]interface{}
  Action_list                           map[string][]string
}


func Construct_server_key( master_flag bool,  master_name, sub_name string )string{
 
       if master_flag == true {
           return  "true~"+master_name+"~"
       }
       return "false~"+master_name+"~"+sub_name
}


var Registry_data Registry_type


func Contruct_Registry(){
     Construct_master_server_list(&Registry_data)
     Registry_data.Inverse_Valve_Map  =    Construct_Inverse_Valve_Map(&Registry_data)
     Registry_data.Action_list                  =    Construct_Action_List(Registry_data.Master_table_list)
     Registry_data. Actions                     =    Construct_Actions( )
   
}



func Construct_master_server_list(r *Registry_type){

   
    r.Master_table_list           = make(map[string][]string)
    r.Valve_list                        = make(map[string]map[string]map[string][]int)
    nodes                                := graph_query.Common_qs_search(&[]string{"IRRIGATION_SERVERS:IRRIGATION_SERVERS","IRRIGATION_SERVER"})
    
    // iteration master controller
    for _,node := range nodes {
        
       name     := graph_query.Convert_json_string(node["name"])
       r.Master_table_list[name], r.Valve_list[name] = find_subnodes( r, name,true )
        
    }
   
}



func find_subnodes( r  *Registry_type, master_node string ,flag bool)([]string,map[string]map[string][]int){
    return_value2 := make(map[string]map[string][]int)
    return_value1 := make([]string,0)
    sub_nodes := graph_query.Common_qs_search(&[]string{"IRRIGATION_SERVERS:IRRIGATION_SERVERS","IRRIGATION_SERVER:"+master_node,"IRRIGATION_SUBSERVER"})
    if len(sub_nodes) == 0{
        return return_value1,return_value2
    }
    for _,sub_node := range(sub_nodes){
        name     := graph_query.Convert_json_string(sub_node["name"])
        
       
    
        byte_array := []byte(sub_node["supported_stations"])
        var data map[string][]int
        if err := json.Unmarshal(byte_array, &data); err != nil {
           panic(err)
        }
        
        return_value2[name]=data
        return_value1 = append(return_value1,name)
    }
   
    return return_value1,return_value2

}


func  Construct_Inverse_Valve_Map( r *Registry_type)map[string]string{
    return_value := make(map[string]string)
    for   master_controller, data1 := range   r.Valve_list {
        for sub_controller,data2 := range data1 {
               data_value := master_controller+":"+ sub_controller
               scan_station_channel_data( return_value,data_value,data2)
        }
    }
    return return_value
    
    
    
}


func scan_station_channel_data( return_value map[string]string, data_value string,  sub_controller_data map[string][]int){
 
    for station, station_data := range   sub_controller_data {
           for _ , channel := range station_data{
          
                key := station+":"+strconv.Itoa(channel)
                return_value[key] = data_value
            }
     }          
        
    
        
    
}

func Construct_Action_List( station_map  map[string][]string )map[string][]string{
    return_value := make(map[string][]string)
    for master_controller , sub_controller_list := range station_map{
         master_key :=    Construct_server_key(true,  master_controller, "" )
         return_value[master_key] =  select_irrigation_action_data( master_key)
         for _ , sub_controller := range sub_controller_list {
             sub_controller_key :=    Construct_server_key(false,  master_controller, sub_controller )
             return_value[sub_controller_key] =  select_irrigation_action_data(sub_controller_key)
         }        
    }
   return return_value
}

    
    
func select_irrigation_action_data( server_key  string )([]string){
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
    return return_value
    
    
}

func Construct_Actions( )map[string]map[string]interface{}{
    return_value := make(map[string]map[string]interface{})
    search_string := []string{"IRRIGATION_ACTIONS:IRRIGATION_ACTIONS","IRRIGATION_ACTION"}
    properties := graph_query.Common_qs_search(&search_string)
    
    for _,property := range properties{
        name := property["name"]
        temp := make(map[string]interface{})
        for key, value := range property{
            temp[key] = value
        }
        return_value[name] = temp
        
            
            
    
        return_value[name]["immediate"] = graph_query.Convert_json_bool(property["immediate"])
        
        
    }
   
    return return_value
}
