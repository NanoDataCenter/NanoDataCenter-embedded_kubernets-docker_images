package data_base_cleanup

import (
   "fmt"
    "time"
     "lacima.com/go_application_containers/irrigation/irrigation_libraries/postgres_access/schedule_access"
)

/*
 * Action and Schedule Data Base Cleansing
 * 
 */
var control_block  irr_sched_access. Irr_sched_access_type
var map_list_data   map[string][]string

var  master_map_keys map[string]bool
var master_map            map[string]map[string]bool

func Start(){
    
    init_data_structures()
    
    go start_database_services()

}
func init_data_structures(){
    control_block =  irr_sched_access.Construct_irr_schedule_access()
    map_list_data  = control_block.Master_table_list  
    master_map_keys = make(map[string]bool)
    master_map         =  make(map[string]map[string]bool)
    for key , value := range map_list_data {
         master_map_keys[key] = true
         temp := make(map[string]bool)
         for _, list_value := range value {
             temp[list_value] = true
         }
         master_map[key] = temp
    }
  
}
    


func start_database_services(){
 	for true {
 
      clean_schedule_table()
       clean_action_table()
       irr_sched_access.Action_Vacuum()
       irr_sched_access.Sched_Vacuum() 
       time.Sleep(time.Second * 3600) // one hour
      fmt.Println("polling loop")

	}
   
    
}

func clean_schedule_table(){
 
    all_entries,err := irr_sched_access.Schedule_Select_All() 
    if err != true {
        panic("bad action table access")
    }
    for _, value := range all_entries {
         master_server :=   value.Master_server 
        sub_server        :=  value.Sub_server    
       name                 :=   value.Name   
       
       if   general_map_check(master_server,sub_server) == false{
           var input irr_sched_access.Schedule_delete_type
          input.Master_server = master_server
           input.Sub_server       = sub_server
           input.Name               = name
            irr_sched_access.Delete_schedule_data( input )
            
       }
    } 
}

func clean_action_table(){
 
    all_entries,err := irr_sched_access.Action_Select_All()
    if err != true {
        panic("bad action table access")
    }
    for _, value := range all_entries {
            var  temp  irr_sched_access.Action_data_type
            temp.Server_type      =   value["master_flag"].(bool)
            temp.Master_server  =   value["main_controller"].(string)
            temp.Sub_server       =   value["sub_controller"].(string)
            temp.Name                 =  value["name"].(string)
            
            if temp.Server_type == false {
                if general_map_check(temp.Master_server,temp.Sub_server) == false{
                    
                    irr_sched_access.Delete_action_data(temp)
                }
                   
            
            }else{
                if check_action_map_list(temp.Master_server) == false{
                  
                     irr_sched_access.Delete_action_data(temp)
                }
                
                
            }
                
    } 
}

func general_map_check( master, sub string) bool{
    if _, ok := master_map_keys[master]; ok == false {
        fmt.Println("fail master",master)
    
        return false
    }
    if _, ok := master_map[master][sub]; ok == false{
        fmt.Println("fail sub",sub)
        return false
    }
    return true
}

func check_action_map_list(master string)bool{
   if _, ok := master_map_keys[master]; ok == false {
        fmt.Println("fail master",master)
    
        return false
    }
    return true
    
    
}
    
    


