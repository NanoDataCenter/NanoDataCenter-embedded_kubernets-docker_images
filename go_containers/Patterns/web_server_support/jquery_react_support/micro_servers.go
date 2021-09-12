package web_support



import (
    
    //"fmt"
    "sort"
    "strings"
    "net/http"
    "html/template"
    "lacima.com/redis_support/generate_handlers"
    "lacima.com/redis_support/redis_handlers"
    "lacima.com/redis_support/graph_query"    
)


var base_templates *template.Template

var micro_servers_template *template.Template
var web_servers redis_handlers.Redis_Hash_Struct



func Generate_Introduction(  )string{
    
    return_array    := make([]string,1)
    return_array[0] = "<center><h2>Front End Web Server for  "+server_id+"  </h2></center>"
    return_value    := strings.Join(return_array,"<br>")
    return return_value
}
    




func Micro_web_page_init( param *template.Template){
    base_templates              = param
    
    
}    


  

func Micro_web_page(w http.ResponseWriter, r *http.Request) {
   data := make(map[string]interface{})
   data["Title"] = "Micro Servers"
   
   micro_servers_template ,_ = base_templates.Clone()
   link_array := generate_application_web_servers()
   
   micro_servers_html := generate_anchor_link_component( "micro_servers_1","<center>List of Application Web Servers</center>", link_array  )
  
   template.Must(micro_servers_template.New("application").Parse(micro_servers_html))
   
   micro_servers_template.ExecuteTemplate(w,"bootstrap", data)
   
    
    
}




func generate_application_web_servers()[]Link_type{
 
    display_struct_search_list := []string{"WEB_IP"}
    data_structures            :=  data_handler.Construct_Data_Structures(&display_struct_search_list)
    web_ip                     := (*data_structures)["WEB_IP"].(redis_handlers.Redis_Hash_Struct)
    web_ip_map                 := web_ip.HGetAll()
    
    data_nodes                 :=  graph_query.Common_qs_search(&[]string{"WEB_MAP:WEB_MAP"})
    data_node                  :=  data_nodes[0]
    web_port                   := graph_query.Convert_json_dict(data_node["port_map"])
    web_description            := graph_query.Convert_json_dict(data_node["description"])
    
    return_value := make([]Link_type,0)

    for key,ip := range web_ip_map {
        
         
         display := web_description[key]
         ip_link := "http://"+ip+web_port[key]
         var link_struct Link_type
         link_struct.Display = display
         link_struct.Link    = ip_link
         return_value = append( return_value,link_struct)
        
    }
    sort.Slice(return_value, func (i,j int)bool {
          return return_value[i].Display < return_value[j].Display
    })
    return return_value
    
}    
    
 func generate_anchor_link_component( container_id string, header string, link_array []Link_type )string{
    html_statements := make([]string,0)
    html_statements = append(html_statements,"<div id=\""+container_id+ "\" class=\"container\">")
    html_statements = append(html_statements,"<h2>"+header+"</h2>")
    html_statements = append(html_statements,`<ul class="list-group">`)
    
    
    for _,element := range link_array {
       html_statements = append(html_statements,`<li class="list-group-item">`+alink_start+element.Link+` target=_blank >`+element.Display+alink_end+"</li>")     
    }    
    
    html_statements = append(html_statements,`</ul>`)
    html_statements = append(html_statements,"</div>")
    return strings.Join(html_statements,"\n")
    


    
}     
