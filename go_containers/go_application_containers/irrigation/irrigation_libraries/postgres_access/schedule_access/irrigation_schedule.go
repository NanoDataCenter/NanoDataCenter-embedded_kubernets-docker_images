package irr_sched_access


import (
   //"fmt"
    "encoding/json"
   "lacima.com/server_libraries/postgres"
   "lacima.com/redis_support/generate_handlers" 
    "lacima.com/redis_support/graph_query"
    
)


type Irr_sched_access_type struct{
  Table_type             string;
  Master_controller      string;
  Sub_controller         string;
  Table_name             string;
  Data_json              string;
  Master_table_list      map[string][]string
  Valve_list             map[string]map[string]interface{}
  Master_table_list_json string 
  Valve_list_json        string        
  sched_driver           pg_drv.Postgres_Table_Driver
  action_driver          pg_drv.Postgres_Float_Driver
}

var control_block Irr_sched_access_type

func Construct_irr_schedule_access()Irr_sched_access_type{
    
    construct_master_server_list(&control_block )
    construct_postgress_data_structures(&control_block)
   
    return control_block
}


   
func construct_master_server_list(r *Irr_sched_access_type){

    master_table_list := make(map[string][]string)
    valve_list        := make(map[string]map[string]interface{})
    nodes := graph_query.Common_qs_search(&[]string{"IRRIGATION_SERVERS:IRRIGATION_SERVERS","IRRIGATION_SERVER"})
    
    for _,node := range nodes {
       name     := graph_query.Convert_json_string(node["name"])
       master_table_list[name],valve_list[name] = find_subnodes( name )
        
        
    }
    
    
    r.Master_table_list = master_table_list
    
    
    temp,_ := json.Marshal(master_table_list)
    r.Master_table_list_json = string(temp)
    
    r.Valve_list        = valve_list
    temp1,_ := json.Marshal(r.Valve_list)
    r.Valve_list_json = string(temp1)
    
   
    
}
    
func find_subnodes( master_node string )([]string,map[string]interface{}){
    return_value2 := make(map[string]interface{})
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

func construct_postgress_data_structures(r *Irr_sched_access_type){
  search_list := []string{"SCHEDULE_DATA:SCHEDULE_DATA","IRRIGATION_DATA"}
  handlers := data_handler.Construct_Data_Structures(&search_list)
  
  r.sched_driver = (*handlers)["IRRIGATION_SCHEDULES"].(pg_drv.Postgres_Table_Driver)
  r.action_driver = (*handlers)["IRRIGATION_ACTIONS"].(pg_drv.Postgres_Float_Driver)
  
}
   
   
   
