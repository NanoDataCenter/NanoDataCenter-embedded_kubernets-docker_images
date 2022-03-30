package construct_actions


import(
   //"fmt"
    "encoding/json"
   "lacima.com/go_application_containers/irrigation/irrigation_libraries/postgres_access/schedule_access"
   
    
)

type Action_data_type struct {
    
    Server_key      string
    Name               string
    Description     string
    Data                 string 
}   


func Ajax_add_action(input string){  
     var action_data map[string]interface{}
     err :=  json.Unmarshal([]byte(input),&action_data)
     if err != nil {
         
      
       panic(err)
    
     }
    
     var access  irr_sched_access.Action_data_type
     access.Data                        = input
     access.Server_key           = action_data["server_key"].(string)
     access.Name                   = action_data["name"].(string)
     access.Description         = action_data["description"].(string)
   
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
     access.Server_key           = action_data["server_key"].(string)
     access.Name                   = action_data["name"].(string)
    
    
     
     
    
     irr_sched_access.Delete_action_data(access)
     
     
}    




    
func Ajax_post_actions(input string)string{ 
    var server_data map[string]interface{}
     err :=  json.Unmarshal([]byte(input),&server_data)
     if err != nil {
       panic(err)
     }
   //fmt.Println("server_key",server_data)
   server_key  := server_data["server_key"].(string)
    return_value,result := irr_sched_access.Select_action_data(server_key)
    
    //fmt.Println("return_value",return_value)
 
    
    if(result != true){
        panic("fail select")
    }
  
 
        
        
    bytes,_ :=  json.Marshal(return_value)
   return string(bytes)
}    



func Ajax_post_irrigation_actions(input string)string{ 
    var server_data map[string]interface{}
     err :=  json.Unmarshal([]byte(input),&server_data)
     if err != nil {
       panic(err)
     }
   //server_data map[master_server:main_server server_type:false sub_server:main_server:sub_server_1]

    server_key  := server_data["server_key"].(string)
    return_value,result := irr_sched_access.Select_irrigation_action_data(server_key)
    
 
    
    if(result != true){
        panic("fail select")
    }
  
 
        
        
    bytes,_ :=  json.Marshal(return_value)
   return string(bytes)
}    
