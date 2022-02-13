package construct_schedule


import(
    "fmt"
    "encoding/json"
    "lacima.com/server_libraries/postgres"
    "lacima.com/redis_support/generate_handlers"

    
)
var  Irrigation_schedules            pg_drv.Postgres_Table_Driver
func initialize_irrigation_schedule_data_structures(){
    
 	search_list                     := []string{"IRRIGATION_DATA_STRUCTURES:IRRIGATION_DATA_STRUCTURES","IRRIGATION_SCHEDULES:IRRIGATION_SCHEDULES","IRRIGATION_SCHEDULES"}
	schedule_structs                := data_handler.Construct_Data_Structures(&search_list)
	
    Irrigation_schedules            = (*schedule_structs)["IRRIGATION_SCHEDULES"].(pg_drv.Postgres_Table_Driver)
	
    Irrigation_schedules.Create_table()
}






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
     
     where_entries := make(map[string]string)
     where_entries["tag1"] = master_server
     where_entries["tag2"] = sub_server
     where_entries["tag3"] = name
     
     fmt.Println("delete",Irrigation_schedules.Delete_Entry(where_entries))
     fmt.Println("add",Irrigation_schedules.Insert( master_server,sub_server,name,description,"",list_data_json ))
     fmt.Println(Irrigation_schedules.Select_All())
 
}
    
func Ajax_delete_schedule(input string){  // input master controller, sub_controller  , schedule_name
     var delete_data map[string]string
     err :=  json.Unmarshal([]byte(input),&delete_data)
     if err != nil {
       panic(err)
     }
    
     master_controller  := delete_data["master_controller"]
     sub_controller     := delete_data["sub_controller"]
     schedule_name      := delete_data["schedule_name"]
    
    
     where_entries := make(map[string]string)
     where_entries["tag1"] = master_controller
     where_entries["tag2"] = sub_controller
     where_entries["tag3"] = schedule_name
     fmt.Println("where entries",where_entries)
     
     fmt.Println("delete",Irrigation_schedules.Delete_Entry(where_entries))
}    
    
func Ajax_post_schedules(input string)string{  // input master controller,sub_controller  output json data
    var server_data map[string]string
     err :=  json.Unmarshal([]byte(input),&server_data)
     if err != nil {
       panic(err)
     }
     
    
   
    master_controller  := server_data["master_controller"]
    sub_controller     := server_data["sub_controller"]
   
    
    where_entries := make(map[string]string)
    where_entries["tag1"] = master_controller
    where_entries["tag2"] = sub_controller
    
   
    data,result := Irrigation_schedules.Select_tags(where_entries)
    if(result != true){
        panic("fail select")
    }
  
    return_value := make([]map[string]interface{},0)
    
    
    for _,input := range data{
        return_entry := make(map[string]interface{})
        return_entry["master_server"]   = input.Tag1
        return_entry["sub_server"]      = input.Tag2
        return_entry["name"]            = input.Tag3
        return_entry["description"]     = input.Tag4
        temp                            := input.Data
        var temp2 interface{}
        err :=  json.Unmarshal([]byte(temp),&temp2)
        if err != nil {
         panic(err)
        }
        return_entry["steps"] = temp2
        return_value = append(return_value,return_entry)
    }
        
    bytes,_ :=  json.Marshal(return_value)
        
   
    
   return string(bytes)
}    
