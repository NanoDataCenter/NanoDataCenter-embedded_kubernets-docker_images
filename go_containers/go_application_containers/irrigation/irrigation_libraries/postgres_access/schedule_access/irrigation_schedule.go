package irr_sched_access


import (
 // "fmt"
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
  driver                 pg_drv.Postgres_Table_Driver
}



func Construct_irr_schedule_access()Irr_sched_access_type{
    var return_value Irr_sched_access_type
    construct_master_server_list(&return_value )
    construct_postgress_data_structures(&return_value)
    
    
    
    
    
    return return_value
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
  search_list := []string{"IRRIGATION_SERVERS:IRRIGATION_SERVERS","IRRIGATION_DATA_STRUCTURES"}
  handlers := data_handler.Construct_Data_Structures(&search_list)
  r.driver = (*handlers)["IRRIGATION_SCHEDULES"].(pg_drv.Postgres_Table_Driver)
 
}
   
   
   



/*
func ( v *Irr_sched_access_type)Data_Clean_Up( )bool{

; //tbd
    
}


func ( v *Irr_sched_access_type)Get_schedule( table_type, master_controller, sub_controller, table_name string )(Irr_table_data,bool){
    
    
     (v  Postgres_Table_Driver)Select_tags(tags map[string]string)([]Table_Output_Data_Record, bool){
    
    
}

func ( v *Irr_sched_access_type)Get_schedule_list(table_type, master_controller, sub_controller string)([]Irr_table_data ,bool){
    
     (v  Postgres_Table_Driver)Select_tags(tags map[string]string)([]Table_Output_Data_Record, bool){
    
    
}

func (v *Irr_sched_access_type)Insert_Modify( table_type, master_controller,sub_controller,table_name,data_json string)bool{
    
    ( v  Postgres_Table_Driver )Insert( tag1,tag2,tag3,tag4,tag5,data string )bool{
    
    
}

func ( v *Irr_sched_access_type)Schedule_delete( table_type, master_controller, sub_controller, table_name string )bool{
    
    
func ( v   Postgres_Table_Driver)Delete_Entry( tags map[string]string)bool{    
    
}



func ( v *Irr_sched_access_type)Vacuum( table_type, master_controller,  )bool{
    
   return driver.Vacuum()
    
    
}
*/
