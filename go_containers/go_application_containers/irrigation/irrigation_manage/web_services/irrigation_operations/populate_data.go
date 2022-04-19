package irrigation_operations


import(
   //"fmt"
    "encoding/json"
   "lacima.com/go_application_containers/irrigation/irrigation_libraries/postgres_access/schedule_access"
   
    
)


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

