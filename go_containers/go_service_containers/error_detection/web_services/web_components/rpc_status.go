package error_detection_components

import (
    "fmt"
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




type rpc_data_structures struct {
    
 
    description                     redis_handlers.Redis_Hash_Struct
    status                          redis_handlers.Redis_Hash_Struct
    contact_time                    redis_handlers.Redis_Hash_Struct
    loading                         redis_handlers.Redis_Hash_Struct
    length                          redis_handlers.Redis_Hash_Struct
    incident_log                    pg_drv.Postgres_Stream_Driver    

}



var rpc_control  rpc_data_structures


func rpc_status_init(){
    
   construct_rpc_data_structures()
   web_support.Generate_special_route("rpc_status/detail/{key}" , rpc_incident_status)

    
}    



func construct_rpc_data_structures() {
   

    data_node_search  := []string{"ERROR_DETECTION:ERROR_DETECTION", "RPC_ANALYSIS:RPC_ANALYSIS" ,"RPC_ANALYSIS_DATA" }
    handlers := data_handler.Construct_Data_Structures(&data_node_search)
    rpc_control.description              = (*handlers)["DESCRIPTION"].(redis_handlers.Redis_Hash_Struct)
    rpc_control.contact_time             = (*handlers)["TIME"].(redis_handlers.Redis_Hash_Struct)
    rpc_control.status                   = (*handlers)["STATUS"].(redis_handlers.Redis_Hash_Struct)
    rpc_control.loading                  = (*handlers)["LOADING"].(redis_handlers.Redis_Hash_Struct)
    rpc_control.length                   = (*handlers)["LENGTH"].(redis_handlers.Redis_Hash_Struct)
    rpc_control.incident_log             = (*handlers)["INCIDENT_LOG"].(pg_drv.Postgres_Stream_Driver)
}

func rpc_status(w http.ResponseWriter, r *http.Request) {
   
   rpc_status_template ,_ := base_templates.Clone()
   
   display_data := generate_rpc_status_html()
   
   template.Must(rpc_status_template.New("application").Parse("<center><h2>RPC Status Data </h2></center><br>"+display_data))
 
    
   data := make(map[string]interface{})
   data["Title"] = "RPC Status"
   rpc_status_template.ExecuteTemplate(w,"bootstrap",data )
   
}
 
 

 
func generate_rpc_status_html()string { 

    
    keys           :=  rpc_control.description.HKeys()
    description    :=  rpc_control.description.HGetAll()
    time_values    :=  rpc_control.contact_time.HGetAll() 
    status         :=  rpc_control.status.HGetAll()
    loading        :=  rpc_control.loading.HGetAll() 
    length         :=  rpc_control.length.HGetAll() 
    sort.Strings(keys)
    
    display_list := make([][]string,len(keys))    
    for index,key      := range keys {
       description_value,_  := msg_pack_utils.Unpack_string(description[key])
       status_value,_       := msg_pack_utils.Unpack_bool(status[key])
       status_string        := strconv.FormatBool(status_value)
       time_value, _        := msg_pack_utils.Unpack_int64(time_values[key])
       time_stamp           := time.Unix(int64(time_value),0)
       load_value,_         := msg_pack_utils.Unpack_float64(loading[key])
       load_string          := fmt.Sprintf("%f",load_value)
       length_value,_       := msg_pack_utils.Unpack_int64(length[key])         
       length_string        := strconv.FormatInt(length_value, 10)
       link                 := web_support.Generate_ajax_anchor_target([]string{"rpc_status","detail",key},"blank","Link to Detailed Data")
       
       display_list[index] = []string{description_value,link,status_string,time_stamp.Format(time.UnixDate),load_string,length_string} 
    }
    
    return web_support.Setup_data_table("topic_list",[]string{"DESCRIPTION","LINK","STATUS","TIME STAMP","Loading","QUEUE LENGTH"},display_list)
}  






func rpc_incident_status(w http.ResponseWriter, r *http.Request) {
   
   rpc_status_template ,_ := base_templates.Clone()
   
   display_data := generate_rpc_incident_status_html(r)
   
   template.Must(rpc_status_template.New("application").Parse("<center><h2>RPC  Incident Log   </h2></center><br>"+display_data))
 
    
   data := make(map[string]interface{})
   data["Title"] = "RPC Incident Log"
   rpc_status_template.ExecuteTemplate(w,"bootstrap", data)
   
}
 
 

 
func generate_rpc_incident_status_html(r *http.Request)string { 
     
    vars := web_support.Get_vars(r)
    key := vars["key"]
    return_value   := "<h5>Summary Data</h5>"
    description,_  := msg_pack_utils.Unpack_string(rpc_control.description.HGet(key))
    status_value,_ := msg_pack_utils.Unpack_bool(rpc_control.status.HGet(key))
    status_string  := strconv.FormatBool(status_value)
    time_value, _  := msg_pack_utils.Unpack_int64(rpc_control.contact_time.HGet(key))
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
    where_clause   := "tag1 = '"+key+"'  and  time >= 0 ORDER BY time DESC LIMIT 25 "
    postgres_data,_ := rpc_control.incident_log.Select_where(where_clause)
    
    length := len(postgres_data)
    if length > 1000 {
        length = 1000
    }
    
    display_list := make([][]string,length)
    
    for index,data := range postgres_data {
     
      time_sec  := data.Time_stamp / 1e9
      time_nsec := data.Time_stamp % 1e9
    
       time_stamp := time.Unix(time_sec ,time_nsec)
       
       status_bool,_ := msg_pack_utils.Unpack_bool(data.Data)
       status_string :=  strconv.FormatBool(status_bool)
       display_list[index] = []string{data.Tag1,status_string,time_stamp.Format(time.UnixDate)} 
    }

    return return_value + web_support.Setup_data_table("topic_list",[]string{"Key","Status","Time"},display_list)
}
 


