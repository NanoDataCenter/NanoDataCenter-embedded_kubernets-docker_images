package site_web_server

import (
    "fmt"
    
    "net/http"
    "html/template"
    "lacima.com/Patterns/web_server_support/jquery_react_support"
    "lacima.com/redis_support/generate_handlers"
    "lacima.com/redis_support/redis_handlers"
   // "lacima.com/Patterns/msgpack"    
)




var node_ip_hash redis_handlers.Redis_Hash_Struct


func node_ip_init(){
   display_struct_search_list := []string{"NODE_MAP"}
   data_structures            :=  data_handler.Construct_Data_Structures(&display_struct_search_list)
   node_ip_hash = (*data_structures)["NODE_MAP"].(redis_handlers.Redis_Hash_Struct)
    
    
}    



func node_ip(w http.ResponseWriter, r *http.Request) {
   
   node_status_template ,_ := base_templates.Clone()
   
   display_data := generate_node_ip_data()
   node_status_html := web_support.Generate_list_link( "Node_IP_1","<center>Node IP Data</center>", display_data )
   template.Must(node_status_template.New("application").Parse(node_status_html))
 
    
   data := make(map[string]interface{})
   data["Title"] = "Node IP"
   node_status_template.ExecuteTemplate(w,"bootstrap", data)
   
    
    
}


func generate_node_ip_data()[]string {
    
  
    return_value := make([]string,0)
    all_data := node_ip_hash.HGetAll()
    fmt.Print("ip all data",all_data)
    
    for key,data := range all_data {
        
        return_value = append(return_value,"Node: "+key+"<br>IP: "+data )
        
    } 
    
    return return_value
}


