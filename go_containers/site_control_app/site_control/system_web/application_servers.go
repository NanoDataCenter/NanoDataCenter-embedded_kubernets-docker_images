package site_web_server

import (
    "fmt"
    "os"
    "net/http"
    "html/template"
    "lacima.com/Patterns/web_server_support/jquery_react_support"
)





var application_servers_template *template.Template




func application_servers_init(){
    application_servers_template ,_ = base_templates.Clone()
    link_array := generate_application_web_servers()
    application_servers_html := web_support.Generate_list_link_component( "container_1","List of Application Web Servers", link_array  )
    
    fmt.Println(application_servers_html)
    template.Must(application_servers_template.New("application").Parse(application_servers_html))
    data := make(map[string]interface{})
   data["Title"] = "Application Server"
    application_servers_template.ExecuteTemplate(os.Stdout,"bootstrap", data)
    
}    


  

func application_servers(w http.ResponseWriter, r *http.Request) {
   data := make(map[string]interface{})
   data["Title"] = "Application Server"
   fmt.Println("made it here")
   application_servers_template.ExecuteTemplate(os.Stdout,"bootstrap", data)
   application_servers_template.ExecuteTemplate(w,"bootstrap", data)
   
    
    
}

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
