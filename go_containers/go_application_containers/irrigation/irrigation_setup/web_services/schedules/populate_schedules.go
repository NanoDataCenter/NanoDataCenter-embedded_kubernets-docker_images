package construct_schedule


import(
    "fmt"
    "encoding/json"
   "lacima.com/go_application_containers/irrigation/irrigation_libraries/postgres_access/schedule_access"
   
    
)






func Ajax_add_schedule(input string){  // input master controller, sub_controller, schedule_name , schedule_data
 
     var schedule_data map[string]interface{}
     err :=  json.Unmarshal([]byte(input),&schedule_data)
     if err != nil {
       panic(err)
     }
    
     list_data :=  schedule_data["steps"].([]interface{})
    
     bytes,err :=  json.Marshal(list_data)
     list_data_json := string(bytes)
     
     description    := schedule_data["description"].(string)
     name           := schedule_data["name"].(string)
     master_server  := schedule_data["master_server"].(string)
     sub_server     := schedule_data["sub_server"].(string)
     
     var where_entries irr_sched_access.Schedule_delete_type
     where_entries.Master_server = master_server
     where_entries.Sub_server = sub_server
     where_entries.Name = name
     
     irr_sched_access.Delete_schedule_data(where_entries)
     
     var input_a irr_sched_access.Schedule_data_type
     input_a.Master_server = master_server
     input_a.Sub_server    = sub_server
     input_a.Name          = name
     input_a.Description   = description
     input_a.Json_data     = list_data_json
     
     irr_sched_access.Insert_schedule_data(input_a)
    
 
}
    
func Ajax_delete_schedule(input string){  // input master controller, sub_controller  , schedule_name
     var delete_data map[string]string
     err :=  json.Unmarshal([]byte(input),&delete_data)
     if err != nil {
       panic(err)
     }
     fmt.Println("delete",input,delete_data)
     master_server      := delete_data["master_server"]
     sub_server         := delete_data["sub_server"]
     schedule_name      := delete_data["name"]
    
     var where_entries  irr_sched_access.Schedule_delete_type
     where_entries.Master_server = master_server
     where_entries.Sub_server    = sub_server
     where_entries.Name          = schedule_name
     
 
     
     irr_sched_access.Delete_schedule_data(where_entries)
}    
    
func Ajax_post_schedules(input string)string{  // input master controller,sub_controller  output json data
    var server_data map[string]string
     err :=  json.Unmarshal([]byte(input),&server_data)
     if err != nil {
       panic(err)
     }
     
     
   
    master_server  := server_data["master_server"]
    sub_server     := server_data["sub_server"]
    
   
    
    data,_  := irr_sched_access.Select_schedule_data(master_server,sub_server) 
    output := make([]map[string]interface{},len(data))
    for index,element := range data{
        temp := make(map[string]interface{})
        temp["master_server"] = element.Master_server
        temp["sub_server"]    = element.Sub_server
        temp["name"]         = element.Name
        temp["description"]  = element.Description
        json_data            := element.Json_data
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
