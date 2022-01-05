package eto_web
   

import (
    
    "lacima.com/go_application_containers/irrigation/eto/eto_web/eto_history"
    "lacima.com/go_application_containers/irrigation/eto/eto_web/eto_rain_values"
    //"lacima.com/go_application_containers/irrigation/eto/eto_web/manage_eto_body"
    "lacima.com/go_application_containers/irrigation/eto/eto_web/rain_history"
    //"lacima.com/go_application_containers/irrigation/eto/eto_web/setup_eto_stations"
    "lacima.com/go_application_containers/irrigation/eto/eto_web/weather_station_problems"
    "lacima.com/go_application_containers/irrigation/eto/eto_web/eto_daily_variation_page"
    "lacima.com/redis_support/generate_handlers"
    "os"
    "lacima.com/redis_support/redis_handlers"
	//"time"
    //"lacima.com/Patterns/msgpack_2"
    "lacima.com/server_libraries/postgres" 
    
    "net/http"
    "html/template"
    "lacima.com/Patterns/web_server_support/jquery_react_support"

    //"github.com/go-redis/redis/v8"
)

var server_id                       string
var base_templates                 *template.Template
var introduction_page_template     *template.Template


var eto_exceptions        redis_handlers.Redis_Hash_Struct

var eto_data              redis_handlers.Redis_Hash_Struct
var rain_data             redis_handlers.Redis_Hash_Struct
var eto_stream_data       redis_handlers.Redis_Hash_Struct
var eto_history           pg_drv.Postgres_Stream_Driver
var rain_history          pg_drv.Postgres_Stream_Driver


func Start(){
   

   server_id  =  os.Getenv("SERVER_ID")
   if server_id == "" {
       panic("bad server id ")
       
   }
   
   web_support.Register_web_page_start(server_id)
   setup_data_structures()
   init_web_server_pages()
   web_support.Launch_web_server()
}
  
  
func setup_data_structures() {

	search_list := []string{"WEATHER_DATA"}
	Eto_data_structs := data_handler.Construct_Data_Structures(&search_list)
	
	eto_exceptions   = (*Eto_data_structs)["EXCEPTION_VALUES"].(redis_handlers.Redis_Hash_Struct)
	
	eto_data         = (*Eto_data_structs)["ETO_VALUES"].(redis_handlers.Redis_Hash_Struct)
	rain_data        = (*Eto_data_structs)["RAIN_VALUES"].(redis_handlers.Redis_Hash_Struct)
    eto_stream_data  = (*Eto_data_structs)["ETO_STREAM_DATA"].(redis_handlers.Redis_Hash_Struct)
    eto_history      = (*Eto_data_structs)["ETO_HISTORY"].(pg_drv.Postgres_Stream_Driver)
    rain_history     = (*Eto_data_structs)["RAIN_HISTORY"].(pg_drv.Postgres_Stream_Driver)
   
}


func init_web_server_pages() {

    web_support.Init_web_support(introduction_page_generate)  // register page
    base_templates = define_web_pages()
    initialize_handlers()
   
}

func initialize_handlers(){
 
    introduction_page_init()
    //manage_eto.Init(base_templates,eto_accumulation)
    //setup_eto_stations.Init(base_templates,eto_accumulation)
    eto_rain_values.Init(base_templates)
    
    eto_daily_variation_page.Init(base_templates,eto_stream_data)
    eto_history_page.Init(base_templates, eto_history)
    rain_history_page.Init(base_templates,rain_history)
    weather_station_problems.Init(base_templates,eto_exceptions)
    web_support.Micro_web_page_init(base_templates)
}






func define_web_pages()*template.Template  {
 
    return_value := make(web_support.Menu_array,9)
    return_value[0] = web_support.Construct_Menu_Element( "introduction page","introduction_page",introduction_page_generate)
    return_value[1] = web_support.Construct_Menu_Element( "Daily ETO Rain Values","eto_rain", eto_rain_values.Generate_page)
    return_value[2] = web_support.Construct_Menu_Element( "ETO Daily Variation","eto_daily_variation", eto_daily_variation_page.Generate_page)
    return_value[3] = web_support.Construct_Menu_Element( "ETO History Table","eto_history_table", eto_history_page.Generate_page_table)
    return_value[4] = web_support.Construct_Menu_Element( "ETO History Graph","eto_history_graph", eto_history_page.Generate_page_graph)
    return_value[5] = web_support.Construct_Menu_Element( "Rain History Table","rain_history_table", rain_history_page.Generate_page_table)
    return_value[6] = web_support.Construct_Menu_Element( "Rain History Graph","rain_history_graph", rain_history_page.Generate_page_graph)
    return_value[7] = web_support.Construct_Menu_Element( "Weather Station Problems","ws_problems", weather_station_problems.Generate_page)
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







const eto_rain_values_body   string  = `

This page showw the previous date eto and rain values
for the various weather stations

`

const eto_streams  string =`
  This page shows a comparision of station values `


const eto_history_body_table  string =`

This page shows the eto history for the various weather stations
In Table form
`

const eto_history_body_graph  string =`

This page shows the eto history for the various weather stations
In Table form
`



const rain_history_body_table  string =`

This page shows the rain history for the various weather stations 
In Table Form
`

const rain_history_body_graph  string =`

This page shows the rain history for the various weather stations 
In Graph Form
`
const weather_station_problems_body  string =`

This page shows the last problem a weather station experience

`







const application_server_body string = `
This web page lists all web servers Relating to Site Micro Services<br><br>

Clink the the link opens Web Page for the Micro Service in a separate table.`

 
   
    
    
func generate_intro_data()[]web_support.Accordion_Elements{

  title_array := []string{"ETO RAIN Values","ETO Stream","ETO History Table","ETO History Graph","Rain History Table","Rain History Graph","Weather Station Problems","Application Servers"}
  body_array  := []string{ eto_rain_values_body,eto_streams, eto_history_body_table,eto_history_body_graph,rain_history_body_table,rain_history_body_graph,
                             weather_station_problems_body,application_server_body }

                          
  return web_support.Populate_accordian_elements(title_array,body_array)
    
    
}
