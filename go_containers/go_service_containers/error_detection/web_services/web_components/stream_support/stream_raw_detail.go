
package stream_support

import (
    "fmt"
    "strings"
    //"strconv"
    "sort"
    "time"
    "net/http"
    "html/template"
    //"encoding/json"
    //"lacima.com/redis_support/graph_query"
    //"lacima.com/redis_support/generate_handlers"
   // "lacima.com/redis_support/redis_handlers"
   // "lacima.com/server_libraries/postgres"
   // "lacima.com/Patterns/msgpack_2"
   "lacima.com/Patterns/web_server_support/jquery_react_support"
   "github.com/vmihailenco/msgpack/v5"
    
    //"net/http"
    //"html/template"
    //"lacima.com/Patterns/web_server_support/jquery_react_support"
    //"lacima.com/redis_support/generate_handlers"
    //"lacima.com/redis_support/redis_handlers"
    //"lacima.com/Patterns/msgpack"
    //"github.com/msgpack/msgpack-go"
 


)

 
type raw_stream_type struct{
    value       float64  
    time_stamp  int64
}


func stream_raw_detail  (w http.ResponseWriter, r *http.Request) {
   
   vars := web_support.Get_vars(r)
   key := vars["key"]
 
    
   incident_status_template ,_ := base_templates.Clone()
   
   display_data := stream_generate_html()
   
   template.Must(incident_status_template.New("application").Parse("<center><h2>Stream Status Data </h2></center><br>"+display_data))
 
    
   data := make(map[string]interface{})
   data["Title"] = "Stream Status"
   incident_status_template.ExecuteTemplate(w,"bootstrap",data )
   
}



func stream_raw_generate_html(key string)string{ 

    
    
    key_tags := strings.Split(key,"/")
    fmt.Println(len(key_tags))
    if len(key_tags) != 6 {
       panic("bad postgres key")
    }
    
    tag1 := key_tags[1]
    tag2 := key_tags[2]
    tag3 := key_tags[3]
    tag4 := key_tags[4]
    tag5 := key_tags[5]
    where_clause   := " tag1 = '"+tag1+" tag2 = '"+tag2+" tag3 = '"+tag3+" tag4 = '"+tag4+" tag5 = '"+tag5+"'  and  time >= 0 ORDER BY time DESC LIMIT 200 "
    pg_data,status := incident_control.incident_log.Select_where(where_clause)
 
    return_value := make([]raw_stream_type,len(pg_data))
    if status != true {
        panic("bad select")
    }
    for index,data_element := range pg_data{
       
       value,err      :=    msg_pack_utils.Unpack_float64(data_element.Data)
       if err != true {
           panic("bad packed data")
       }
    
       return_value[index].value := value
       return_value[index].time_stamp := data.Time
    }
    return generate_table_display(return_value)
    
}

func generate_table_display( return                 
