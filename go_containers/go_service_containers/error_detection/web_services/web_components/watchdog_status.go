package error_detection_components

import (
    //"fmt"
    //"strings"
    "strconv"
    "sort"
    "time"
    "net/http"
    "html/template"
    "lacima.com/redis_support/generate_handlers"
    "lacima.com/redis_support/redis_handlers"
    "lacima.com/server_libraries/postgres"
    "lacima.com/Patterns/msgpack_2"
    "lacima.com/Patterns/web_server_support/jquery_react_support"

    
    //"net/http"
    //"html/template"
    //"lacima.com/Patterns/web_server_support/jquery_react_support"
    //"lacima.com/redis_support/generate_handlers"
    //"lacima.com/redis_support/redis_handlers"
    //"lacima.com/Patterns/msgpack"
    //"github.com/msgpack/msgpack-go"
 


)


type wd_data_structures struct {
    
 

    status                          redis_handlers.Redis_Hash_Struct
    time_stamp                      redis_handlers.Redis_Hash_Struct
    wd_incidents                    pg_drv.Postgres_Stream_Driver
    

}


var wd_control  wd_data_structures


func watchdog_status_init(){
    
   construct_wd_data_structures()

    
}    



func construct_wd_data_structures() {
   

    wd_data_nodes                       := []string{"ERROR_DETECTION:ERROR_DETECTION", "WD_DETECTION:WD_DETECTION" ,"WATCH_DOG_DATA"  }
    handlers                            := data_handler.Construct_Data_Structures(&wd_data_nodes)
    wd_control.status                   = (*handlers)["DEBOUNCED_STATUS"].(redis_handlers.Redis_Hash_Struct)
    wd_control.time_stamp               = (*handlers)["TIME_STAMP"].(redis_handlers.Redis_Hash_Struct)
    wd_control.wd_incidents             = (*handlers)["WATCH_DOG_LOG"].(pg_drv.Postgres_Stream_Driver)    
 
}

func watchdog_status(w http.ResponseWriter, r *http.Request) {
   
   wd_status_template ,_ := base_templates.Clone()
   
   display_data := generate_wd_status_html()
   
   template.Must(wd_status_template.New("application").Parse("<center><h2>Watch Dog Status Data </h2></center><br>"+display_data))
 
    
   data := make(map[string]interface{})
   data["Title"] = "Watch Dog Status"
   wd_status_template.ExecuteTemplate(w,"bootstrap",data )
   
}
 
func generate_wd_status_html()string { 

    
  
    keys            := wd_control.status.HKeys()
    status          := wd_control.status.HGetAll()
    time_values     := wd_control.time_stamp.HGetAll() 
    sort.Strings(keys)
    
    display_list := make([][]string,len(keys))    
    for index,key := range keys {
       status_value,_ := msg_pack_utils.Unpack_bool(status[key])
       status_string  := strconv.FormatBool(status_value)
       time_value, _  := msg_pack_utils.Unpack_int64(time_values[key])
       time_stamp     := time.Unix(int64(time_value),0)
       
       display_list[index] = []string{key,status_string,time_stamp.Format(time.UnixDate)} 
    }
    
    return web_support.Setup_data_table("topic_list",[]string{"ID","STATUS","TIME STAMP"},display_list)
}  


func watchdog_incident_status(w http.ResponseWriter, r *http.Request) {
   
   wd_status_template ,_ := base_templates.Clone()
   
   display_data := generate_wd_incident_status_html()
   
   template.Must(wd_status_template.New("application").Parse("<center><h2>Watch Dog Incident Log   </h2></center><br>"+display_data))
 
    
   data := make(map[string]interface{})
   data["Title"] = "Watch Incident Log"
   wd_status_template.ExecuteTemplate(w,"bootstrap", data)
   
}
 
func generate_wd_incident_status_html()string { 

    
    postgres_data,_ := wd_control.wd_incidents.Select_after_time_stamp_desc(3600*24*30*6) // six months
    
    length := len(postgres_data)
    if length > 1000 {
        length = 1000
    }
    
    display_list := make([][]string,length)
    
    for index,data := range postgres_data {
     
      time_sec  := data.Time_stamp / 1e9
      time_nsec := data.Time_stamp % 1e9
    
       time_stamp := time.Unix(time_sec ,time_nsec)
       
       stream_id_string := strconv.FormatInt(data.Stream_id,10) 
       display_list[index] = []string{stream_id_string,data.Tag1,data.Tag2,time_stamp.Format(time.UnixDate)} 
    }

    return web_support.Setup_data_table("topic_list",[]string{"Key ID","Stats","Time"},display_list)
}
 




