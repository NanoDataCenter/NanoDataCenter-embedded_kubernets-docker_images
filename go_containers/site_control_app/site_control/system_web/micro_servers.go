package web_support



import (
    
    
    "net/http"
    "html/template"
    "lacima.com/Patterns/web_server_support/jquery_react_support"
    "lacima.com/redis_support/generate_handlers"
    "lacima.com/redis_support/redis_handlers"
    //"lacima.com/Patterns/msgpack"    
)

var micro_servers_template *template.Template
var web_servers redis_handlers.Redis_Hash_Struct



func Micro_web_page_init(){
    display_struct_search_list := []string{"WEB_MAP"}
    data_structures            :=  data_handler.Construct_Data_Structures(&display_struct_search_list)
    web_servers = (*data_structures)["WEB_MAP"].(redis_handlers.Redis_Hash_Struct)
 
    
    micro_servers_template ,_ = base_templates.Clone()
    link_array := generate_application_web_servers()
    micro_servers_html := web_support.Generate_list_link_component( "micro_servers_1","<center>List of Application Web Servers</center>", link_array  )
    
    
    template.Must(micro_servers_template.New("application").Parse(micro_servers_html))
    
}    


  

func Micro_web_page(w http.ResponseWriter, r *http.Request) {
   data := make(map[string]interface{})
   data["Title"] = "Micro Servers"
   
   
   micro_servers_template.ExecuteTemplate(w,"bootstrap", data)
   
    
    
}


func generate_application_web_servers()[]web_support.Link_type{
 
    return_value := make([]web_support.Link_type,0)
    
    return return_value
    
}
