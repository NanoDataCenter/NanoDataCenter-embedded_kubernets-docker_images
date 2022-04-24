package irrigation_information_web
     

import (
     "io"
    // "encoding/json"
     "fmt"
    
    "lacima.com/go_application_containers/irrigation/irrigation_manage/web_services/eto_adjust"
    "lacima.com/go_application_containers/irrigation/irrigation_manage/web_services/irrigation_station_channel"
    "lacima.com/go_application_containers/irrigation/irrigation_manage/web_services/irrigation_valve_group"
    "lacima.com/go_application_containers/irrigation/irrigation_manage/web_services/irrigation_operations"
    "lacima.com/go_application_containers/irrigation/irrigation_manage/web_services/irrigation_stream_data"
    "lacima.com/go_application_containers/irrigation/irrigation_manage/web_services/manage_irrigation_parameters"
    "lacima.com/go_application_containers/irrigation/irrigation_manage/web_services/past_irrigation_jobs"
    "lacima.com/go_application_containers/irrigation/irrigation_manage/web_services/manage_irrigation_queue"
    
    //"lacima.com/redis_support/generate_handlers"
    "os"
    //"lacima.com/redis_support/redis_handlers"
	//"time"
    //"lacima.com/Patterns/msgpack_2"
    

      "encoding/json"
    "net/http"
    "html/template"
    "lacima.com/Patterns/web_server_support/jquery_react_support"
   "lacima.com/go_application_containers/irrigation/irrigation_libraries/postgres_access/schedule_access"

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
    web_support.Generate_special_post_route("irrigation/eto/eto_adjust_store" , eto_adjust_store)
    web_support.Generate_special_post_route("irrigation/irrigation_manage/post_job" , irrigation_post_job)
    web_support.Generate_special_post_route("irrigation/irrigation_manage/get_actions",get_actions) 
    web_support.Generate_special_post_route("irrigation/irrigation_manage/get_schedules",get_schedules)
    web_support.Generate_special_post_route("irrigation/irrigation_manage/post_action",post_action)
    web_support.Generate_special_post_route("irrigation/irrigation_manage/post_schedule",post_schedule)
}

func initialize_handlers(){
 
    introduction_page_init()
    eto_adjust.Page_init(base_templates)
    irrigation_station_channel.Page_init(base_templates)
    irrigation_valve_group.Page_init(base_templates)
    irrigation_operations.Page_init(base_templates)
    irrigation_manage_parameters.Page_init(base_templates)
    irrigation_past_operation.Page_init(base_templates)
    irrigation_manage_queue.Page_init(base_templates)
    irrigation_streaming_data.Page_init(base_templates)
    web_support.Micro_web_page_init(base_templates)
}






func define_web_pages()*template.Template  {
 
    return_value := make(web_support.Menu_array,9)
    return_value[0] = web_support.Construct_Menu_Element( "Iintroduction page","introduction_page",introduction_page_generate)
    return_value[1] = web_support.Construct_Menu_Element( "ETO Manage","eto_manage", eto_adjust.Generate_page_adjust)
    return_value[2] = web_support.Construct_Menu_Element("Irrigation Valve Group","irrigation_valve_group",irrigation_valve_group.Generate_page_adjust)
    return_value[3] = web_support.Construct_Menu_Element("Irrigation Station Channel","irrigation_station_channel",irrigation_station_channel.Generate_page_adjust)
    return_value[4] = web_support.Construct_Menu_Element("Irrigation Operations","manual_operations",irrigation_operations.Generate_page_adjust)
    return_value[5] = web_support.Construct_Menu_Element("Irrigation Manage Parameters","irrigation_manage_parameters",irrigation_manage_parameters.Generate_page_adjust)
    return_value[6] = web_support.Construct_Menu_Element("Irrigation Past Operations","irrigation_past_operation",irrigation_past_operation.Generate_page_adjust)
    return_value[7] = web_support.Construct_Menu_Element("Irrigation Streaming Data","irrigation_streaming_data",irrigation_streaming_data.Generate_page_adjust)
    return_value[8] = web_support.Construct_Menu_Element( "Other Servers","other_servers", web_support.Micro_web_page)
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




const eto_modify_body  string =`
  This page allows modification of eto accumulation data `

const irrigation_valve_group_op  string =`
  This page allows to do diagnostics operations by valve group`

const irrigation_station_channel_op  string =`
  This page allows to do diagnostics operations by station channel `


const manual_operations_op  string =`
  This page allows to do diagnostics operations `

const manage_irrigation_parameters_op  string =`
  This page allows to do diagnostics operations `

const past_irrigation_jobs_op string =`
  This page allows to do diagnostics operations `


 const manage_irrigation_jobs_op string =`
  This page allows to do diagnostics operations `


 
const irrigation_stream_data_op  string =`
  This page allows to do diagnostics operations `

 

const application_server_body string = `
This web page lists all web servers Relating to Site Micro Services<br><br>

Clink the the link opens Web Page for the Micro Service in a separate table.`

 
   
    
 
    
func generate_intro_data()[]web_support.Accordion_Elements{

  title_array := []string{"ETO Change","Irrigation Diagnostics Valve Group","Irrigation Diagnostics Station Channel","Manual Operations","Manage Irrigation Parameters","Past Irrigation Jobs","Manage Irrigation Jobs", "Irrigation Stream Data","Application Servers"}
  body_array  := []string{eto_modify_body, irrigation_valve_group_op ,  irrigation_station_channel_op,  manual_operations_op ,manage_irrigation_parameters_op,past_irrigation_jobs_op,manage_irrigation_jobs_op, irrigation_stream_data_op,application_server_body }

                          
  return web_support.Populate_accordian_elements(title_array,body_array)
    
    
}  


/*
 * 
 * 
 * Ajax handlers 
 * 
*/

 

func post_action( w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
 input,err :=  io.ReadAll(r.Body)
  if err != nil {
      panic(err)
  }
  
  handle_post_action(string(input))
  
  output := []byte(`"SUCCESS"`)
  
   w.Write(output) 
    
}   

func   handle_post_action(raw_input string ){
       var decode_value map[string]interface{}
  
    if err := json.Unmarshal([]byte(raw_input), &decode_value); err != nil {
        panic("bad json")
    }else{
        key := decode_value["key"].(string)
        action := decode_value["action"].(string)
        
         fmt.Println("action",action, irr_sched_access.Queue_Action( key,   action ) )
    }
    
}



func post_schedule( w http.ResponseWriter, r *http.Request) {
    
    w.Header().Set("Content-Type", "application/json")


input,err :=  io.ReadAll(r.Body)
  if err != nil {
      panic(err)
  }
  handle_post_schedule(string(input))
  output := []byte(`"SUCCESS"`)
  
   w.Write(output) 
    
}   

func   handle_post_schedule(raw_input string ){
       var decode_value map[string]interface{}
  
    if err := json.Unmarshal([]byte(raw_input), &decode_value); err != nil {
        panic("bad json")
    }else{
        handle_schedule(decode_value)
    }
    
}

 


func irrigation_post_job(w http.ResponseWriter, r *http.Request) {
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
  
        parse_irrigation_direct( string(input))
      
  }
  
  output := []byte(`"SUCCESS"`)
  
   w.Write(output) 
    
}   

func parse_irrigation_direct(raw_input string){
    var decode_value map[string]interface{}
  
    if err := json.Unmarshal([]byte(raw_input), &decode_value); err != nil {
        panic("bad json")
    }else{
       station := decode_value["station"].(string)
       io          := int64(decode_value["io"].(float64))
       irr_sched_access.Queue_Irrigation_Direct( station ,io  )
    }
    
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



func get_actions(w http.ResponseWriter, r *http.Request) {
 
    
    w.Header().Set("Content-Type", "application/json")


input,err :=  io.ReadAll(r.Body)
  if err != nil {
      panic(err)
  }else{
    output :=  irrigation_operations.Ajax_post_irrigation_actions(string(input))
   
    w.Write([]byte(output) )
  }  
}

func get_schedules(w http.ResponseWriter, r *http.Request) {
 
    
    w.Header().Set("Content-Type", "application/json")


input,err :=  io.ReadAll(r.Body)
  if err != nil {
      panic(err)
  }else{
    output :=  irrigation_operations.Ajax_post_schedules(string(input))
   
    w.Write([]byte(output) )
  }  
}

/*
func queue_irrigation_jobs(   json_data map[string]interface{} ) {
    data :=  json_data["steps"].([]interface{})
    for  _, temp := range data {
        array_element                 :=   temp.(map[string]interface{})
        name                                := array_element["name"].(string)
        action_type                     := array_element["type"].(string)
        if action_type == "schedule" {
               handle_schedule(key,name)
        }else{
          
            fmt.Println("action",name, irr_sched_access.Queue_Action( key,  name ) )
        }
    }
}

*/

/*
func  handle_schedule(data map[string]interface{} ){
     key := data["key"].(string)
     steps := data["schedule"].([]interface{})
     for _, temp := range steps{
         
             temp_map := temp.(map[string]interface{})
             time              := temp_map["time"].(float64)
             steps            :=   generate_step_data(temp_map["steps"].(string))
             fmt.Println("steps",key,time,steps)
            station_io := generate_station_io( steps )
            fmt.Println("station",time,station_io)
            fmt.Println(irr_sched_access.Queue_Managed_Irrigation( key, time ,  station_io ))

      }
        
}
*/
func  handle_schedule(data map[string]interface{}  ){
                key := data["key"].(string)
    
               steps := generate_step_data(data["schedule"].([]interface))
               for _,step := range steps {
                   time           := step["time"].(float64)
                   temp         :=  step["station"].(map[string]interface {})
                    station_io := generate_station_io( temp )
                    //fmt.Println("station",time,station_io)
                   fmt.Println(irr_sched_access.Queue_Managed_Irrigation( key, time ,  station_io ))
               }
               
       
           
       
}


func generate_step_data(input string)[]map[string]interface{}{
    var data []map[string]interface{}
        if err := json.Unmarshal([]byte(input), &data); err != nil {
           panic(err)
        }
        return data
}

//station:map[station_3:map[1:1] station_4:map[1:1]] time:60]
func generate_station_io(  input map[string]interface {})[]string{
  return_value := make([]string,0)
   for station, io_data :=  range input {
       for pin , _  := range io_data.(map[string]interface{}) {
          return_value= append(return_value, station+":"+pin)
       }
   }
   return return_value
}
