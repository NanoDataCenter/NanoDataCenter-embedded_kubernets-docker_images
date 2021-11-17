package error_detection_components

import (
    //"fmt"
    "strings"
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
    
 
    description                     redis_handlers.Redis_Hash_Struct
    status                          redis_handlers.Redis_Hash_Struct
    time_stamp                      redis_handlers.Redis_Hash_Struct
    wd_incidents                    pg_drv.Postgres_Stream_Driver
    

}


var wd_control  wd_data_structures


func watchdog_status_init(){
    
   construct_wd_data_structures()
   web_support.Generate_special_route("watchdog_status/detail/{key}" , watchdog_incident_status)

    
}    



func construct_wd_data_structures() {
   

    wd_data_nodes                       := []string{"ERROR_DETECTION:ERROR_DETECTION", "WD_DETECTION:WD_DETECTION" ,"WATCH_DOG_DATA"  }
    handlers                            := data_handler.Construct_Data_Structures(&wd_data_nodes)
    wd_control.status                   = (*handlers)["DEBOUNCED_STATUS"].(redis_handlers.Redis_Hash_Struct)
    wd_control.time_stamp               = (*handlers)["TIME_STAMP"].(redis_handlers.Redis_Hash_Struct)
    wd_control.wd_incidents             = (*handlers)["WATCH_DOG_LOG"].(pg_drv.Postgres_Stream_Driver)    
    wd_control.description              = (*handlers)["DESCRIPTION"].(redis_handlers.Redis_Hash_Struct)
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

    
  
    keys               := wd_control.status.HKeys()
    status             := wd_control.status.HGetAll()
    time_values        := wd_control.time_stamp.HGetAll()
    description_values := wd_control.description.HGetAll()
    sort.Strings(keys)
    
    display_list := make([][]string,len(keys))    
    for index,key := range keys {
       description    := description_values[key]
       status_value,_ := msg_pack_utils.Unpack_bool(status[key])
       status_string  := strconv.FormatBool(status_value)
       time_value, _  := msg_pack_utils.Unpack_int64(time_values[key])
       time_stamp     := time.Unix(int64(time_value),0)
       link           := web_support.Generate_ajax_anchor_target([]string{"watchdog_status","detail",key},"blank","Link to Detailed Data")
       
       display_list[index] = []string{description,link,status_string,time_stamp.Format(time.UnixDate)} 
    }
    
    return web_support.Setup_data_table("topic_list",[]string{"DESCRIPTION","LINK","STATUS","TIME STAMP"},display_list)
}  






func watchdog_incident_status(w http.ResponseWriter, r *http.Request) {
   
   wd_status_template ,_ := base_templates.Clone()
   
   display_data := generate_wd_incident_status_html(r)
   
   template.Must(wd_status_template.New("application").Parse("<center><h2>Watch Dog Incident Log   </h2></center><br>"+display_data))
 
    
   data := make(map[string]interface{})
   data["Title"] = "Watch Incident Log"
   wd_status_template.ExecuteTemplate(w,"bootstrap", data)
   
}
 
func generate_wd_incident_status_html(r *http.Request)string { 
     
    vars := web_support.Get_vars(r)
    key := vars["key"]
    return_value   := "<h5>Summary Data</h5>"
    description    := wd_control.description.HGet(key)
    status_value,_ := msg_pack_utils.Unpack_bool(wd_control.status.HGet(key))
    status_string  := strconv.FormatBool(status_value)
    time_value, _  := msg_pack_utils.Unpack_int64(wd_control.time_stamp.HGet(key))
    time_stamp     := time.Unix(int64(time_value),0)
    time_string    := time_stamp.Format(time.UnixDate)
    list_data :=  []string{  "<ul>",
                               "<li>Key:  "+key+"</li>",
                               "<li>Description: "+description+"</li>",
                               "<li>Status: "+status_string+"</li>",
                               "<li>Time: "+time_string+"</li>",
                               "</ul>" }
                               
    list_string := strings.Join(list_data,"\n")
     
     
      
      return_value = return_value+list_string
    
    
    return_value =  return_value + "<h5>Time History Changes </h5>"
    where_clause   := "tag1 = '"+key+"'  and  time >= 0 ORDER BY time DESC LIMIT 10 "
    postgres_data,_ := wd_control.wd_incidents.Select_where(where_clause)
    
    length := len(postgres_data)
    if length > 1000 {
        length = 1000
    }
    
    display_list := make([][]string,length)
    
    for index,data := range postgres_data {
     
      time_sec  := data.Time_stamp / 1e9
      time_nsec := data.Time_stamp % 1e9
    
       time_stamp := time.Unix(time_sec ,time_nsec)
       
        
       display_list[index] = []string{data.Tag1,data.Tag2,time_stamp.Format(time.UnixDate)} 
    }

    return return_value + web_support.Setup_data_table("topic_list",[]string{"Key","Stats","Time"},display_list)
}
 




