package construct_schedule


import(
   //"fmt"
    "encoding/json"
   "lacima.com/go_application_containers/irrigation/irrigation_libraries/postgres_access/schedule_access"
   
    
)





 
func Ajax_add_schedule(input string){  // input master controller, sub_controller, schedule_name , schedule_data
 
     var schedule_data map[string]interface{}
     err :=  json.Unmarshal([]byte(input),&schedule_data)
     if err != nil {
       panic(err)
     }
    
     var where_entries irr_sched_access.Schedule_data_type
     where_entries.Server_key        = schedule_data["server_key"].(string)
     where_entries.Name                 = schedule_data["name"].(string)
     where_entries.Description       = schedule_data["description"].(string)
     where_entries.Json_data        =  schedule_data["json_steps"].(string)
     
     irr_sched_access.Delete_schedule_data(where_entries)
     
   
     irr_sched_access.Insert_schedule_data(where_entries)
    
 
}
    
func Ajax_delete_schedule(input string){  // input master controller, sub_controller  , schedule_name
     var schedule_data map[string]interface{}
     err :=  json.Unmarshal([]byte(input),&schedule_data)
     if err != nil {
       panic(err)
     }
     
     var where_entries irr_sched_access.Schedule_data_type
     where_entries.Server_key        = schedule_data["server_key"].(string)
     where_entries.Name                 = schedule_data["name"].(string)
    
     //fmt.Println("where_entries",where_entries)
     irr_sched_access.Delete_schedule_data(where_entries)
}    
    
func Ajax_post_schedules(input string)string{  // input master controller,sub_controller  output json data
    var server_data map[string]interface{}
     err :=  json.Unmarshal([]byte(input),&server_data)
     if err != nil {
       panic(err)
     }
     
    
    server_key     := server_data["server_key"].(string)
    
    data,_  := irr_sched_access.Select_schedule_data(server_key) 
    output := make([]map[string]interface{},len(data))
    for index,element := range data{
        temp := make(map[string]interface{})
        temp["server_key"]    = element.Server_key
        temp["name"]            = element.Name
        temp["description"]    = element.Description
        json_data                   := element.Json_data
        var temp_data interface{}        
        err :=  json.Unmarshal([]byte(json_data),&temp_data)
        if err != nil {
          panic(err)
        }
        temp["steps"] = temp_data
        output[index] = temp
    }
        
    bytes,_ :=  json.Marshal(output)
   
    return string(bytes)
}    
