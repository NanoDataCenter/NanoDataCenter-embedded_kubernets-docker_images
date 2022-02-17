package construct_actions


import(
    "fmt"
    "encoding/json"
   "lacima.com/go_application_containers/irrigation/irrigation_libraries/postgres_access/schedule_access"
   
    
)














func Ajax_add_action(input string){  
     var action_data map[string]interface{}
     err :=  json.Unmarshal([]byte(input),&action_data)
     if err != nil {
       panic(err)
     }
    
     
    
    
     
     name            := action_data["name"].(string)
     description     := action_data["description"].(string)
     master_server   := action_data["master_server"].(string)
     sub_server      := action_data["sub_server"].(string)
     start_hr        := action_data["start_hr"].(float64)
     start_min       := action_data["start_min"].(float64)
     end_hr          := action_data["end_hr"].(float64)
     end_min         := action_data["end_min"].(float64)
     data            := action_data["action_data"].(string) 
     
     var delete_input irr_sched_access.Action_delete_type
     delete_input.Master_server =  master_server
     delete_input.Sub_server    =  sub_server
     delete_input.Name          =  name  
    
     
     fmt.Println("delete",irr_sched_access.Delete_action_data(delete_input))
     
     
     var input_a irr_sched_access.Action_data_type
    
     input_a.Master_server =  master_server
     input_a.Sub_server    =  sub_server
     input_a.Name          =  name  
     input_a.Description   =  description
     input_a.Start_hr      =  start_hr 
     input_a.Start_min     =  start_min
     input_a.End_hr        =  end_hr
     input_a.End_min       =  end_min
     input_a.Json_data     =  data 
     
     fmt.Println("add",irr_sched_access.Insert_action_data(input_a))
     
     
     
    
 
}
    
func Ajax_delete_action(input string){  // input master controller, sub_controller  , schedule_name
     var delete_data map[string]string
     err :=  json.Unmarshal([]byte(input),&delete_data)
     if err != nil {
       panic(err)
     }
    
     
     
     master_server  := delete_data["master_server"]
     sub_server     := delete_data["sub_server"]
     name           := delete_data["action_name"]

     var delete_input irr_sched_access.Action_delete_type
     delete_input.Master_server =  master_server
     delete_input.Sub_server    =  sub_server
     delete_input.Name          =  name  
     
     
     fmt.Println("delete",irr_sched_access.Delete_action_data(delete_input))
     
     
}    



    
func Ajax_post_actions(input string)string{ 
    var server_data map[string]string
     err :=  json.Unmarshal([]byte(input),&server_data)
     if err != nil {
       panic(err)
     }
     
    
   
    master_server  := server_data["master_server"]
    sub_server     := server_data["sub_server"]
   
    return_value,result := irr_sched_access.Select_action_data(master_server,sub_server)
    
    
    if(result != true){
        panic("fail select")
    }
  
 
        
        
    bytes,_ :=  json.Marshal(return_value)
   return string(bytes)
}    
