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

func Start(){
    
    init_data_structures()
    
    go start_database_services()

}
func init_data_structures(){
    control_block =  irr_sched_access.Construct_irr_schedule_access()
  
}
    


func start_database_services(){
 	for true {
 
        clean_schedule_table()
        clean_action_table()
        irr_sched_access.Sched_Vacuum()
        irr_sched_access.Action_Vacuum()
        time.Sleep(time.Second * 3600) // one hour
       fmt.Println("polling loop")

	}
   
    
}

func clean_schedule_table(){
 
    all_entries,err :=  irr_sched_access.Schedule_Select_All() 
    if err != true {
        panic("bad action table access")
    }
    for _, value := range all_entries {
        key := value.Server_key 
       
       if   general_map_check(key ) == false{
           fmt.Println("key does not exit",key)
          
            irr_sched_access.Delete_schedule_data( value )
            
       }
    } 
}

func clean_action_table(){
 
    all_entries,err :=  irr_sched_access.Action_Select_All()
    if err != true {
        panic("bad action table access")
    }
    for _, value := range all_entries {
            key := value.Server_key 
            
             if   general_map_check(key ) == false{
                  fmt.Println("key does not exit",key)
                irr_sched_access.Delete_action_data(value)
             }
            
    } 
}

func general_map_check( key  string) bool{
    if _,ok := control_block.Node_list[key] ; ok == false {
        return false
    }
    return true
}


    
    


