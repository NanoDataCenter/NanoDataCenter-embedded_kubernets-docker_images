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
     fmt.Println("schedule_data",schedule_data)
     list_data :=  schedule_data["steps"].([]interface{})
     fmt.Println("list_data",list_data)
     bytes,err :=  json.Marshal(list_data)
     list_data_json := string(bytes)
     fmt.Println("list_data_json",list_data_json)
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

    
     where_entries := make(map[string]string)
     //where_entries["tag1"] = master_server
     //where_entries["tag2"] = sub_server
     //where_entries["tag3"] = name
     
     fmt.Println("delete",Irrigation_schedules.Delete_Entry(where_entries))
}    
    
func Ajax_post_schedules(input string)string{  // input master controller,sub_controller  output json data
    var server_data map[string]string
     err :=  json.Unmarshal([]byte(input),&server_data)
     if err != nil {
       panic(err)
     }
     
    fmt.Println("server_data",server_data)
   
    master_controller  := server_data["master_controller"]
    sub_controller     := server_data["sub_controller"]
   
    fmt.Println("master_server",master_controller)
    fmt.Println("sub_server",sub_controller)
    where_entries := make(map[string]string)
    where_entries["tag1"] = master_controller
    where_entries["tag2"] = sub_controller
    
    fmt.Println("where_entries",where_entries)
    data,result := Irrigation_schedules.Select_tags(where_entries)
    if(result != true){
        panic("fail select")
    }
    fmt.Println("data",data)
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
        
    fmt.Println("post ",bytes)
    
   return string(bytes)
}    
