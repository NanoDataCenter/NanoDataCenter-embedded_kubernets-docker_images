package construct_actions


import(
  //  "fmt"
    "encoding/json"
   "lacima.com/go_application_containers/irrigation/irrigation_libraries/postgres_access/schedule_access"
   
    
)



func Ajax_add_action(input string){  
     var action_data map[string]interface{}
     err :=  json.Unmarshal([]byte(input),&action_data)
     if err != nil {
       panic(err)
    
     }
    
     var access  irr_sched_access.Action_data_type
     access.Data                        = input
     access.Server_type          = action_data["master_flag"].(bool)
     access.Master_server     = action_data["main_controller"].(string)
     access.Sub_server          = action_data["sub_controller"].(string)
     access.Name                   = action_data["name"].(string)
     access.Description         = action_data["description"].(string)
     access.Start_time           = (action_data["start_time_hr"].(float64)*60.)+action_data["start_time_min"].(float64)
     access.End_time             =  (action_data["end_time_hr"].(float64)*60.)+action_data["end_time_min"].(float64)
    irr_sched_access.Delete_action_data(access)    
    irr_sched_access.Insert_action_data(access) 

}   

func Ajax_delete_action(input string){  // input master controller, sub_controller  , schedule_name
     var action_data map[string]interface{}
     err :=  json.Unmarshal([]byte(input),&action_data)
     if err != nil {
       panic(err)
     }
     
     
     
    var access  irr_sched_access.Action_data_type
     access.Data                        = input
     access.Server_type          = action_data["master_flag"].(bool)
     access.Master_server     = action_data["main_controller"].(string)
     access.Sub_server          = action_data["sub_controller"].(string)
     access.Name                   = action_data["name"].(string)
    
     
     
    
     irr_sched_access.Delete_action_data(access)
     
     
}    




    
func Ajax_post_actions(input string)string{ 
    var server_data map[string]interface{}
     err :=  json.Unmarshal([]byte(input),&server_data)
     if err != nil {
       panic(err)
     }
     server_type := "false"
      if server_data["master_flag"].(bool) == true {
        server_type   = "true"
     }
   
    master_server         := server_data["master_controller"].(string)
    sub_server              := server_data["sub_controller"].(string)
   
    return_value,result := irr_sched_access.Select_action_data(server_type,master_server,sub_server)
    
 
    
    if(result != true){
        panic("fail select")
    }
  
 
        
        
    bytes,_ :=  json.Marshal(return_value)
   return string(bytes)
}    
