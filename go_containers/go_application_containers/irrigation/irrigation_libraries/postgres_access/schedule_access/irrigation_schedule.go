package irr_sched_access


import (
   "fmt"
    "encoding/json"
   "lacima.com/server_libraries/postgres"
   "lacima.com/redis_support/generate_handlers" 
    "lacima.com/redis_support/graph_query"
    "lacima.com/redis_support/redis_handlers"
    
)


type Irr_sched_access_type struct{
  Table_type                        string;
  Master_controller           string;
  Sub_controller                 string;
  Table_name                     string;
  Data_json                         string;
  Node_list                          map[string]bool
  Master_table_list            map[string][]string
  
  Valve_list                         map[string]map[string]interface{}
  Master_table_list_json  string 
  Valve_list_json               string        
  sched_driver                  pg_drv.Json_Table_Driver
  action_driver                 pg_drv.Json_Table_Driver
  redis_hash_driver         redis_handlers.Redis_Hash_Struct
}

var control_block Irr_sched_access_type

func Construct_irr_schedule_access()Irr_sched_access_type{
    
    construct_master_server_list(&control_block )
    construct_postgress_data_structures(&control_block)
    irrigation_RPC_Client_Init()
    return control_block
}

func Construct_server_key( master_flag bool,  master_name, sub_name string )string{
 
       if master_flag == true {
           return  "true~"+master_name+"~"+sub_name
       }
       return "false~"+master_name+"~"+sub_name
}

   
func construct_master_server_list(r  *Irr_sched_access_type){

    r.Node_list                 =  make(map[string]bool)
    r.Master_table_list = make(map[string][]string)
    r.Valve_list        = make(map[string]map[string]interface{})
    nodes := graph_query.Common_qs_search(&[]string{"IRRIGATION_SERVERS:IRRIGATION_SERVERS","IRRIGATION_SERVER"})
    
    for _,node := range nodes {
       name     := graph_query.Convert_json_string(node["name"])
       
       key :=  Construct_server_key(true,name,"")
       fmt.Println("key",key)
       r.Node_list[key] = true
       r.Master_table_list[name], r.Valve_list[name] = find_subnodes( r, name,true )
        
        key =  Construct_server_key(false,name,"")
       fmt.Println("key",key)
       r.Node_list[key] = true
       r.Master_table_list[name], r.Valve_list[name] = find_subnodes( r, name,false )
       
        
    }
    
    
    
    
    
    temp,_ := json.Marshal(r.Master_table_list)
    r.Master_table_list_json = string(temp)
    
    temp1,_ := json.Marshal(r.Valve_list)
    r.Valve_list_json = string(temp1)
    
   
    
}
    
func find_subnodes( r  *Irr_sched_access_type, master_node string ,flag bool)([]string,map[string]interface{}){
    return_value2 := make(map[string]interface{})
    return_value1 := make([]string,0)
    sub_nodes := graph_query.Common_qs_search(&[]string{"IRRIGATION_SERVERS:IRRIGATION_SERVERS","IRRIGATION_SERVER:"+master_node,"IRRIGATION_SUBSERVER"})
    if len(sub_nodes) == 0{
        return return_value1,return_value2
    }
    for _,sub_node := range(sub_nodes){
        name     := graph_query.Convert_json_string(sub_node["name"])
        key :=  Construct_server_key(flag,master_node,name)
        fmt.Println("key",key)
       r.Node_list[key] = true
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
  r.redis_hash_driver   = (*handlers)["IRRIGATION_HISTORY_HASH"].(redis_handlers.Redis_Hash_Struct)
  r.sched_driver = (*handlers)["IRRIGATION_SCHEDULES"].(pg_drv.Json_Table_Driver)
  r.action_driver = (*handlers)["IRRIGATION_ACTIONS"].(pg_drv.Json_Table_Driver)
  
}


func Check_schedule_job( key string )bool{
    temp := control_block.redis_hash_driver.HGet(key)
    return_value := false
    if temp == "true" {
        return_value = true
    }
    //fmt.Println("temp",temp,return_value)
    return return_value
}

func Set_schedule_job( key string ){
 
       control_block.redis_hash_driver.HSet(key,"true")
      //  fmt.Println(Get_all_keys())
}
    
func Clear_schedule_job( key string ){
    control_block.redis_hash_driver.HSet(key,"false")
     // fmt.Println(Get_all_keys())
  
}

func Get_all_keys( )map[string]string{
    
    return control_block.redis_hash_driver.HGetAll()
    
}

func Delete_schedule_job(){
  //  fmt.Println(Get_all_keys())
   
    
    
    control_block.redis_hash_driver.Delete_All()
}
        
    
