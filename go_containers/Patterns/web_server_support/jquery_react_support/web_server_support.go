package web_support



import (
  "fmt"
  "os"
  "strings"  
  "html/template"
  "net"
  "net/http"
  "github.com/gorilla/mux"
  "lacima.com/redis_support/generate_handlers"
  "lacima.com/redis_support/redis_handlers"
  "lacima.com/redis_support/graph_query"
)


type Menu_function_type func(http.ResponseWriter, *http.Request)



type Menu_element struct{
  
    Menu_name string;
    Menu_link string;
    Menu_function  Menu_function_type  
    
    
}

type Menu_array []Menu_element

func Construct_Menu_Element( name,link string, function Menu_function_type ) Menu_element{
    var return_value Menu_element
    return_value.Menu_name      = name
    return_value.Menu_link      = link
    return_value.Menu_function  = function
    return return_value
}

var web_page_start string
var server_id        string
    

var slash_ref_function Menu_function_type

var template_array *template.Template

var router *mux.Router


func Get_Web_Start()string{
    return web_page_start
}

func Register_web_page_start( input string ){
 
    
    data_nodes                 :=  graph_query.Common_qs_search(&[]string{"WEB_MAP:WEB_MAP"})
    data_node                  :=  data_nodes[0]
    start_label_map            :=  graph_query.Convert_json_dict(data_node["start_label"])
    
    start_label , ok           :=  start_label_map[input]
    if ok == false {
        fmt.Println("bad server_id")
        fmt.Println("input",input)
        fmt.Println("map",start_label_map)
        os.Exit(1)
    }
    
    
    web_page_start = "/"+start_label+"/"
    server_id        = input
}

func Init_web_support( slash_page Menu_function_type)   {
    router = mux.NewRouter()
    slash_ref_function = slash_page
    generate_base_templates()
    
    
}

func Register_web_pages( menu_array Menu_array){
 
    //http.HandleFunc("/",slash_page)
    router.HandleFunc("/", slash_page) 
    //http.HandleFunc(web_page_start,slash_page)
    router.HandleFunc(web_page_start, slash_page) 
    for _,element := range menu_array {
       fmt.Println("link",element.Menu_link,element.Menu_function)
       //http.HandleFunc(web_page_start +element.Menu_link,element.Menu_function)
       router.HandleFunc(web_page_start+element.Menu_link, element.Menu_function)
    }
    
    
}

func Generate_single_row_menu( menu_array Menu_array )*template.Template {
    
    working_array := make([]string,0)
    working_array = append(working_array,nav_start)
    for _,menu_element := range menu_array {
        if menu_element.Menu_name != "/"{
            
            element := `<a class="dropdown-item" href="`+web_page_start +menu_element.Menu_link+`"  target="_self">`+menu_element.Menu_name+"</a>"
            working_array = append(working_array,element)
        }
    }
    working_array = append(working_array,nav_end)
    template_string :=strings.Join(working_array,"\n")
    
    template.Must(template_array.New("menu").Parse(template_string)) 
    return template_array
}

func Launch_web_server( ){
   
   display_struct_search_list := []string{"WEB_IP"}
   data_structures            :=  data_handler.Construct_Data_Structures(&display_struct_search_list)
   web_ip                     := (*data_structures)["WEB_IP"].(redis_handlers.Redis_Hash_Struct)
   ip_address                 := find_local_address()
   fmt.Println("ip_address",ip_address)
   web_ip.HSet(server_id,ip_address )
   
   data_nodes                 :=  graph_query.Common_qs_search(&[]string{"WEB_MAP:WEB_MAP"})
   data_node                  :=  data_nodes[0]
   web_port                   :=  graph_query.Convert_json_dict(data_node["port_map"])
   fmt.Println("web_port",web_port)
   fmt.Println("web_ip",web_ip.HGetAll())
   
  
   go http.ListenAndServe(web_port[server_id], router)
}
    
func find_local_address()string{
    
   conn, error := net.Dial("udp", "8.8.8.8:80")  
   if error != nil {  
      fmt.Println(error)  
  
    }  
  
    defer conn.Close()  
    ipAddress_port := conn.LocalAddr().(*net.UDPAddr).String()
    temp := strings.Split(ipAddress_port,":")
    ip_address := temp[0]
  
    return ip_address
}  

func slash_page(w http.ResponseWriter, r *http.Request){
    
 
        slash_ref_function(w,r)

}

//http://error_detection/ajax/incident/[SYSTEM:farm_system][SITE:LACIMA_SITE][CONTAINER_LIST:CONTAINER_LIST][CONTAINER:file_server][INCIDENT_LOG:managed_process_failure]

func Generate_special_route( route_string string , web_page_function Menu_function_type){
    
    fmt.Println("link",web_page_start+`ajax/`+route_string)
    
    router.HandleFunc(web_page_start+`ajax/`+route_string, web_page_function )
    
}






func Generate_special_post_route( route_string string , web_page_function Menu_function_type){
    
    fmt.Println("link",web_page_start+`ajax/`+route_string)
    
    router.HandleFunc(web_page_start+`ajax/`+route_string, web_page_function ).Methods("POST")
    
}

func Generate_special_get_route( route_string string , web_page_function Menu_function_type){
    
    fmt.Println("link",web_page_start+`ajax/`+route_string)
    
    router.HandleFunc(web_page_start+`ajax/`+route_string, web_page_function ).Methods("GET")
    
}


func Generate_ajax_anchor( links []string,display string ) string {
    
    anchor := `<a href="`+web_page_start+`ajax/`+strings.Join(links,"/")+`">`+display+"</a>"
    //fmt.Println("anchor",anchor)
    return anchor
    
    
}

func Generate_ajax_anchor_target( links []string,target,display string ) string {
    
    anchor := `<a href="`+web_page_start+`ajax/`+strings.Join(links,"/")+`"  target="`+target+`">`+display+"</a>"
    //fmt.Println("anchor",anchor)
    return anchor
    
    
}

func Generate_ajax_anchor_with_parameters( links []string,parameters,display string ) string {
    
    anchor := `<a href="`+web_page_start+`ajax/`+strings.Join(links,"/")+"?"+parameters+`">`+display+"</a>"
    //fmt.Println("anchor",anchor)
    return anchor
    
    
}


func Generate_ajax_anchor_with_parameters_and_target( links []string,parameters,target,  display string ) string {
    
    anchor := `<a href="`+web_page_start+`ajax/`+strings.Join(links,"/")+"?"+parameters+`"  target="`+target+`">`+display+"</a>"
    //fmt.Println("anchor",anchor)
    return anchor
    
    
}


// <button onclick="location.href='http://www.example.com'" type="button">
//         www.example.com</button>
func Generate_link_button( link string,display string ) string {
    
    button := `<button  onclick="location.href='`+web_page_start+link+  `'"    type="button">`+display+`</button>`
    //fmt.Println("link",button)
   
    return button
    
    
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, "custom 404")
    }
}


func generate_base_templates() {
    template_array = template.New("base")
    template.Must(template_array.New("bootstrap").Parse(base_template))
    
}

func Get_vars(r *http.Request)map[string]string{

    return mux.Vars(r)

}

  



var base_template = `

<!DOCTYPE html>
<html lang="en">
  <head>
   
    
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/css/bootstrap.min.css" integrity="sha384-TX8t27EcRE3e/ihU7zmQxVncDAy5uIKz4rEkgIXeMed4M0jlfIDPvg6uqKI2xXr2" crossorigin="anonymous">
<link rel="stylesheet" type="text/css" href="https://cdn.datatables.net/1.11.0/css/jquery.dataTables.css">
   <!-- jquery & Bootstrap JS -->
<script
  src="https://code.jquery.com/jquery-3.6.0.js"
  integrity="sha256-H+K7U5CnXl1h5ywQfKtSj8PCmoN9aaq30gDh27Xc0jk="
  crossorigin="anonymous"></script>
<script src="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-ho+j7jyWK8fNQe+A12Hb8AhRq26LrZ/JpcUGGOn+Y7RsweNrtN/tE3MoK7ZeZDyx" crossorigin="anonymous"></script>
<script type="text/javascript" charset="utf8" src="https://cdn.datatables.net/1.11.0/js/jquery.dataTables.js"></script>

<script type="text/javascript" >

</script>

<style>
input[type="range"] {
  width: 100%;
  margin-bottom: 3rem;
}

.auto {
  -ms-touch-action: auto;
  touch-action: auto;
} 
</style>
<title>{{.Title}}</title>
  </head>

  <body>
  {{template "menu" }}
  <style>
    .application-example{
        margin: 20px;
    }
</style>

<div class="application-example">
   
  {{template "application"}}
  
</div> 
   

 
  </body>
</html>

`

var nav_start = `
<nav class="navbar navbar-expand-sm bg-dark navbar-dark">

  <!-- Links -->
  <ul class="navbar-nav">

    <!-- Dropdown -->
    <li class="nav-item dropdown">
      <a class="nav-link dropdown-toggle" href="#" id="navbardrop" data-toggle="dropdown">Menu</a>
      
      <div class="dropdown-menu">
`



        
var nav_end = `
      </div>
    </li>
  </ul>
  <ul class="navbar-nav">

      <button id="status_panel", class="btn " type="submit">Status</button>
  </ul>
  <nav class="navbar navbar-light bg-dark navbar-dark">
  <span class="navbar-text" >
   <h4 id ="status_display"> Status: </h4>
  </span>
  </nav>
</nav>

`
