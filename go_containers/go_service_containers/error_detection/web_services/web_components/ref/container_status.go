package site_web_server

import (
    //"fmt"
    
   
    "strconv"
    "net/http"
    "html/template"
    "lacima.com/Patterns/web_server_support/jquery_react_support"
    "lacima.com/redis_support/generate_handlers"
    "lacima.com/redis_support/redis_handlers"
    "lacima.com/Patterns/msgpack"
    //"github.com/msgpack/msgpack-go"

)




var container_hash_status redis_handlers.Redis_Hash_Struct  




func container_status_init(){
   
   display_struct_search_list := []string{"DOCKER_CONTROL"}
   data_structures            :=  data_handler.Construct_Data_Structures(&display_struct_search_list)
   container_hash_status = (*data_structures)["DOCKER_DISPLAY_DICTIONARY"].(redis_handlers.Redis_Hash_Struct)
    
}    



func container_status(w http.ResponseWriter, r *http.Request) {
   
   container_status_template ,_ := base_templates.Clone()
   
   display_data := generate_container_status_data()
   container_status_html := web_support.Generate_list_link( "continer_status_1","<center>Container Status</center>", display_data )
   
   template.Must(container_status_template.New("application").Parse(container_status_html))
   data := make(map[string]interface{})
   data["Title"] = "Container Status"
   container_status_template.ExecuteTemplate(w,"bootstrap", data)
   
    
    
}

//key redis string
//data_array redis map[interface {}]interface {}
//kv, ok := v.Interface().(map[interface{}]interface{})["XYZ"]


func generate_container_status_data()[]string {
    return_value := make([]string,0)
    all_data := container_hash_status.HGetAll()
    for key,data := range all_data {
        
        data_map := msgpack_utils.Unpack(data)
        
       
        active := data_map.(map[interface{}]interface{})["active"].(bool)
        managed := data_map.(map[interface{}]interface{})["managed"].(bool)
    
        return_value = append(return_value,"Container: "+key+"<br>Active: "+strconv.FormatBool(active)+"<br>Managed: "+strconv.FormatBool(managed) )
        
    }  
    return return_value
}


     
        
 
    





