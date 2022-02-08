package construct_schedule


import(
    "fmt"
    "encoding/json"
   // "lacima.com/Patterns/web_server_support/jquery_react_support"
)








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
    
}
    
func Ajax_delete_schedule(input string){  // input master controller, sub_controller  , schedule_name

}    
    
func Ajax_post_schedules(input map[string]string)string{  // input master controller,sub_controller  output json data
   
   fmt.Println(input) 
    
   schedule_data := make([]map[string]interface{},0)
   bytes,err :=  json.Marshal(schedule_data)
   if err != nil {
       panic(err)
   }
   return string(bytes)
}    
