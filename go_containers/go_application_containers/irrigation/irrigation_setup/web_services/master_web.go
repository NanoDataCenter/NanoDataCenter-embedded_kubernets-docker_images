package irrigation_information_web
     

import (
     //"fmt"
     "io"
     //"encoding/json"
     "fmt"
    "lacima.com/go_application_containers/irrigation/irrigation_setup/web_services/eto_setup"
    "lacima.com/go_application_containers/irrigation/irrigation_setup/web_services/eto_adjust"
    "lacima.com/go_application_containers/irrigation/irrigation_setup/web_services/schedules"
    "lacima.com/go_application_containers/irrigation/irrigation_setup/web_services/actions"
    
    //"lacima.com/redis_support/generate_handlers"
    "os"
    //"lacima.com/redis_support/redis_handlers"
	//"time"
    //"lacima.com/Patterns/msgpack_2"
    
    
    "net/http"
    "html/template"
    "lacima.com/Patterns/web_server_support/jquery_react_support"

    //"github.com/go-redis/redis/v8"
)

var server_id                       string
var base_templates                 *template.Template
var introduction_page_template     *template.Template




func Start(){
   

   server_id  =  os.Getenv("SERVER_ID")
   if server_id == "" {
       panic("bad server id ")
       
   }
   
   web_support.Register_web_page_start(server_id)
   init_web_server_pages()
   web_support.Launch_web_server()
}


func init_web_server_pages() {

    web_support.Init_web_support(introduction_page_generate)  // register page
    base_templates = define_web_pages()
    initialize_handlers()
    web_support.Generate_special_post_route("irrigation/eto/eto_setup_store" , eto_setup_store)
    web_support.Generate_special_post_route("irrigation/eto/eto_adjust_store" , eto_adjust_store)
    
    
    web_support.Generate_special_post_route("irrigation/irrigation_schedules/add_schedule",add_schedule)
    web_support.Generate_special_post_route("irrigation/irrigation_schedules/get_schedules" ,   get_schedules)
    web_support.Generate_special_post_route("irrigation/irrigation_schedules/delete_schedule" , delete_schedule)
    
    web_support.Generate_special_post_route("irrigation/irrigation_schedules/add_action",add_action)
    web_support.Generate_special_post_route("irrigation/irrigation_schedules/delete_action",delete_action) 
    web_support.Generate_special_post_route("irrigation/irrigation_schedules/get_actions",get_actions) 
    
                                             
}

func initialize_handlers(){
 
    introduction_page_init()
    eto_setup.Page_init(base_templates)
    eto_adjust.Page_init(base_templates)
    construct_schedule.Page_init(base_templates)
    construct_actions.Page_init(base_templates)
    web_support.Micro_web_page_init(base_templates)
}






func define_web_pages()*template.Template  {
 
    return_value := make(web_support.Menu_array,6)
    return_value[0] = web_support.Construct_Menu_Element( "Iintroduction page","introduction_page",introduction_page_generate)
    return_value[1] = web_support.Construct_Menu_Element( "ETO Station Setup","eto_setup", eto_setup.Generate_page_setup)
    return_value[2] = web_support.Construct_Menu_Element( "ETO Manage","eto_manage", eto_adjust.Generate_page_adjust)
    return_value[3] = web_support.Construct_Menu_Element( "Construct Schedules","construct_schedule",construct_schedule.Generate_page)
    return_value[4] = web_support.Construct_Menu_Element( "Construct Action","construct_action",construct_actions.Generate_page)
    return_value[5] = web_support.Construct_Menu_Element( "Other Servers","other_servers", web_support.Micro_web_page)
    web_support.Register_web_pages(return_value)
    return web_support.Generate_single_row_menu(return_value)
}
    

 




func introduction_page_init( ){
    introduction_page_template ,_ = base_templates.Clone()
    introduction_page_html := web_support.Generate_Introduction()+ web_support.Generate_accordian("intro_1","Description of Web Pages",  generate_intro_data())
    template.Must(introduction_page_template.New("application").Parse(introduction_page_html))    
}    

 

func introduction_page_generate(w http.ResponseWriter, r *http.Request) {
   
   data := make(map[string]interface{})
   data["Title"] = "Introduction"
   introduction_page_template.ExecuteTemplate(w,"bootstrap", data)
}   


const eto_setup_body   string  = `

This page allows creation of eto resource values

`

const eto_modify_body  string =`
  This page allows modification of eto accumulation data `


const irrigation_schedule_body  string =`

This page allows creation copying deletion and modification of an irrigation schedule
`

const irrigation_maintainence_body  string =`

This page allows Modification of Irrigation Maintainece Operations
`








const application_server_body string = `
This web page lists all web servers Relating to Site Micro Services<br><br>

Clink the the link opens Web Page for the Micro Service in a separate table.`

 
   
    
    
func generate_intro_data()[]web_support.Accordion_Elements{

  title_array := []string{"ETO Setup","ETO Change","Irrigation Schedule","Maintainence Schedule","Application Servers"}
  body_array  := []string{ eto_setup_body,eto_modify_body, irrigation_schedule_body,irrigation_maintainence_body,application_server_body }

                          
  return web_support.Populate_accordian_elements(title_array,body_array)
    
    
}  


/*
 * 
 * 
 * Ajax handlers 
 * 
*/


func eto_setup_store(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  //var input interface{}

  /*if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        fmt.Println(err)
       // panic("BAD:")
    }
  */
  
  input,err :=  io.ReadAll(r.Body)
  if err != nil {
      fmt.Println(err)
  }else{   
  
     eto_setup.Process_new_eto_setup(string(input))  
      
  }
  
  output := []byte(`"SUCCESS"`)
  
   w.Write(output) 
    
}



func eto_adjust_store(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  //var input interface{}

  /*if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        fmt.Println(err)
       // panic("BAD:")
    }
  */
  
  input,err :=  io.ReadAll(r.Body)
  if err != nil {
      fmt.Println(err)
  }else{   
  
     eto_adjust.Process_new_eto_adjust(string(input))  
      
  }
  
  output := []byte(`"SUCCESS"`)
  
   w.Write(output) 
    
}

func add_schedule(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
 
  
  input,err :=  io.ReadAll(r.Body)
  if err != nil {
      fmt.Println(err)
  }else{   
  
     construct_schedule.Ajax_add_schedule(string(input))  // input master controller, sub_controller, schedule_name , schedule_data
      
  }
  
  output := []byte(`"SUCCESS"`)
  
   w.Write(output) 
    
}
func delete_schedule(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  input,err :=  io.ReadAll(r.Body)
  if err != nil {
      fmt.Println(err)
  }else{   
  
     construct_schedule.Ajax_delete_schedule(string(input))  // input master controller, sub_controller  , schedule_name
      
  }
  
  output := []byte(`"SUCCESS"`)
  
   w.Write(output) 
    
}
func get_schedules(w http.ResponseWriter, r *http.Request) {
 
    
    w.Header().Set("Content-Type", "application/json")


input,err :=  io.ReadAll(r.Body)
  if err != nil {
      panic(err)
  }else{
    output :=  construct_schedule.Ajax_post_schedules(string(input))
   
    w.Write([]byte(output) )
  }  
}

func add_action(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
 
  
  input,err :=  io.ReadAll(r.Body)
  if err != nil {
      fmt.Println(err)
  }else{   
  
     construct_actions.Ajax_add_action(string(input))  // input master controller, sub_controller, schedule_name , schedule_data
      
  }
  
  output := []byte(`"SUCCESS"`)
  
   w.Write(output) 
    
}
func delete_action(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  input,err :=  io.ReadAll(r.Body)
  if err != nil {
      fmt.Println(err)
  }else{   
  
     construct_actions.Ajax_delete_action(string(input))  // input master controller, sub_controller  , schedule_name
      
  }
  
  output := []byte(`"SUCCESS"`)
  
   w.Write(output) 
    
}

func get_actions(w http.ResponseWriter, r *http.Request) {
 
    
    w.Header().Set("Content-Type", "application/json")


input,err :=  io.ReadAll(r.Body)
  if err != nil {
      panic(err)
  }else{
    output :=  construct_actions.Ajax_post_actions(string(input))
   
    w.Write([]byte(output) )
  }  
}



 


 