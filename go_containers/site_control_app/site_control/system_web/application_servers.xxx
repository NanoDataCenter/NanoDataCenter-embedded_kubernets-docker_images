package site_web_server

import (
    
    
    "net/http"
    "html/template"
    "lacima.com/Patterns/web_server_support/jquery_react_support"
    "lacima.com/redis_support/generate_handlers"
    "lacima.com/redis_support/redis_handlers"
    //"lacima.com/Patterns/msgpack"    
)





var application_servers_template *template.Template
var web_servers redis_handlers.Redis_Hash_Struct



func application_servers_init(){
    display_struct_search_list := []string{"WEB_MAP"}
    data_structures            :=  data_handler.Construct_Data_Structures(&display_struct_search_list)
    web_servers = (*data_structures)["WEB_MAP"].(redis_handlers.Redis_Hash_Struct)
 
    
    application_servers_template ,_ = base_templates.Clone()
    link_array := generate_application_web_servers()
    application_servers_html := web_support.Generate_list_link_component( "applicationr_1","<center>List of Application Web Servers</center>", link_array  )
    
    
    template.Must(application_servers_template.New("application").Parse(application_servers_html))
    data := make(map[string]interface{})
    data["Title"] = "Application Server"
    
    
}    


  

func application_servers(w http.ResponseWriter, r *http.Request) {
   data := make(map[string]interface{})
   data["Title"] = "Application Server"
   
   
   application_servers_template.ExecuteTemplate(w,"bootstrap", data)
   
    
    
}


func generate_application_web_servers()[]web_support.Link_type{
 
    return_value := make([]web_support.Link_type,0)
    
    return return_value
    
}

/*

func generate_application_web_servers()[]web_support.Link_type{
 
    return_value := make([]web_support.Link_type,0)
    link_value := web_support.Link_type{ Display:"display1",Link:"\"http://127.0.0.1/\"  target=\"_blank\""}
    
    return_value = append(return_value,link_value)    
    link_value = web_support.Link_type{ Display:"display2",Link:"\"http://127.0.0.1/\"  target=\"_blank\""}
    return_value = append(return_value,link_value)  
    link_value = web_support.Link_type{ Display:"display3",Link:"\"http://127.0.0.1/\"  target=\"_blank\""}
    return_value = append(return_value,link_value)     
    link_value = web_support.Link_type{ Display:"display4",Link:"\"http://127.0.0.1/\"  target=\"_blank\""}
    return_value = append(return_value,link_value)     
    link_value = web_support.Link_type{ Display:"display5",Link:"\"http://127.0.0.1/\"  target=\"_blank\""}
    return_value = append(return_value,link_value)     
    return return_value
    
}
*/
