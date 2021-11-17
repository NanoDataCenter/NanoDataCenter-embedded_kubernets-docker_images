package error_detection_components

import (
    "fmt"
    "strings"
    "strconv"
    "sort"
    "time"
    "net/http"
    "html/template"
    "encoding/json"
    "lacima.com/redis_support/graph_query"
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


type incident_control_type struct {
    

   
    time                            redis_handlers.Redis_Hash_Struct
    status                          redis_handlers.Redis_Hash_Struct
    last_error_data                 redis_handlers.Redis_Hash_Struct
    last_error_time                 redis_handlers.Redis_Hash_Struct
    description                     redis_handlers.Redis_Hash_Struct
    old_status                      redis_handlers.Redis_Hash_Struct
    review_state                    redis_handlers.Redis_Hash_Struct
    acknowlege_state                redis_handlers.Redis_Hash_Struct
    incident_log                    pg_drv.Postgres_Stream_Driver
    keys                            map[string]incident_record_type
}

type incident_record_type struct {
  name               string
  description        string
  max_time_interval  int64
  namespace          string
  key_array          []string
  key                string
  contact_time       redis_handlers.Redis_Single_Structure
  status              redis_handlers.Redis_Single_Structure
  last_error_time     redis_handlers.Redis_Single_Structure
  last_error_data     redis_handlers.Redis_Single_Structure
  
    
}

var incident_control       incident_control_type
var incident_records       []incident_record_type




func incident_status_init(){
   
   construct_incident_data_structures() 
   construct_incident_data_nodes()
   construct_keys()
   web_support.Generate_special_route("incident_status/detail/{key}" , incident_detail_function)
   web_support.Generate_special_post_route("incident_status/review" , incident_change_review)
   web_support.Generate_special_post_route("incident_status/reset" , incident_change_reset)
   web_support.Generate_special_post_route("incident_status/acknowlege" , incident_change_acknowlege)
}
   
   








func construct_incident_data_structures() {
   
    data_node_search  := []string{"ERROR_DETECTION:ERROR_DETECTION", "INCIDENT_STREAMS:INCIDENT_STREAMS" ,"INCIDENT_DATA"  }
    handlers := data_handler.Construct_Data_Structures(&data_node_search)
    incident_control.time                     = (*handlers)["TIME"].(redis_handlers.Redis_Hash_Struct)
    incident_control.status                   = (*handlers)["STATUS"].(redis_handlers.Redis_Hash_Struct)
    incident_control.last_error_data          = (*handlers)["LAST_ERROR"].(redis_handlers.Redis_Hash_Struct)
    incident_control.last_error_time          = (*handlers)["ERROR_TIME"].(redis_handlers.Redis_Hash_Struct)
    incident_control.description              = (*handlers)["DESCRIPTION"].(redis_handlers.Redis_Hash_Struct)
    incident_control.review_state             = (*handlers)["REVIEW_STATE"].(redis_handlers.Redis_Hash_Struct)
    incident_control.acknowlege_state         = (*handlers)["ACKNOWLEGE_STATE"].(redis_handlers.Redis_Hash_Struct)
    incident_control.old_status               = (*handlers)["OLD_STATUS"].(redis_handlers.Redis_Hash_Struct)
    
    incident_control.incident_log             = (*handlers)["INCIDENT_LOG"].(pg_drv.Postgres_Stream_Driver)
    
 
}

func construct_incident_data_nodes(){
    incident_records = make([]incident_record_type,0)
    incident_nodes  := []string{"INCIDENT_LOG"}
    nodes := graph_query.Common_qs_search(&incident_nodes)
   
   
    for _,node := range nodes{
        var item  incident_record_type
        //fmt.Println("node",node)
        item.name               = graph_query.Convert_json_string(node["name"])
        item.description        = graph_query.Convert_json_string(node["description"])
        
        item.namespace          = graph_query.Convert_json_string(node["namespace"])
        item.key_array          = graph_query.Generate_key(item.namespace)
        item.key_array          = append(item.key_array,"INCIDENT_LOG")
        item.key                = strings.Join(item.key_array,"/")
        
        handlers                := data_handler.Construct_Data_Structures(&item.key_array)
        item.contact_time       = (*handlers)["TIME_STAMP"].(redis_handlers.Redis_Single_Structure)
        item.status             = (*handlers)["STATUS"].(redis_handlers.Redis_Single_Structure)
        item.last_error_data    = (*handlers)["LAST_ERROR"].(redis_handlers.Redis_Single_Structure)
        item.last_error_time    = (*handlers)["ERROR_TIME"].(redis_handlers.Redis_Single_Structure)
        incident_records        = append(incident_records,item)
    }
    
}







       
    
    
func construct_keys(){
    incident_control.keys = make(map[string]incident_record_type)
    
    for _,item := range incident_records{
        
        incident_control.keys[item.namespace] = item
        
    }
}


func incident_status (w http.ResponseWriter, r *http.Request) {
   
   incident_status_template ,_ := base_templates.Clone()
   
   display_data := generate_incident_status_html()
   
   template.Must(incident_status_template.New("application").Parse("<center><h2>INCIDENT Status Data </h2></center><br>"+display_data))
 
    
   data := make(map[string]interface{})
   data["Title"] = "INCIDENT Status"
   incident_status_template.ExecuteTemplate(w,"bootstrap",data )
   
}
 
func generate_incident_status_html()string { 

   
  
    keys                 := incident_control.status.HKeys()
    sort.Strings(keys)
    
    old_status           := incident_control.old_status.HGetAll()
    review               := incident_control.review_state.HGetAll()
    status               := incident_control.status.HGetAll()
    contact_time_values  := incident_control.time.HGetAll() 
    descriptions         := incident_control.description.HGetAll()
    last_error_times      := incident_control.last_error_time.HGetAll()
    acknowlege_state      := incident_control.acknowlege_state.HGetAll()
    //last_error_data       := incident_control.last_error_data.HGetAll()
    
   
    display_list := make([][]string,len(keys))    
    for index,key             := range keys {
       status_value,_         := msg_pack_utils.Unpack_bool(status[key])
       status_string          := strconv.FormatBool(status_value)
       
       old_status_value,_         := msg_pack_utils.Unpack_bool(old_status[key])
      
       ack_value,err       := msg_pack_utils.Unpack_bool(acknowlege_state[key])
       if err != true {
           ack_value = false
       }

       review_value,err       := msg_pack_utils.Unpack_bool(review[key])
       if err != true {
           review_value = false
       }
       if (status_value == false) && (old_status_value == true ) {
           review_value = false
       }
       
       
       if status_value   == true {
           review_value = true
       }
       if old_status_value != status_value {
           incident_control.old_status.HSet(key,status_string)
       }
           
    
       
       
       review_string          := msg_pack_utils.Pack_bool(review_value)
       incident_control.review_state.HSet(key,review_string)
       
       review_display    := "Needs To Review"
       if review_value == true {
           review_display = ""
       }
       
       ack_display    := "Acknowledged"
       if ack_value == false {
           ack_display = ""
       }       
       
       description_string,_   := msg_pack_utils.Unpack_string(descriptions[key])
       
       contact_time           := format_time(msg_pack_utils.Unpack_int64(contact_time_values[key]))
       error_time             := format_time(msg_pack_utils.Unpack_int64(last_error_times[key]))
       link                   := web_support.Generate_ajax_anchor_target([]string{"incident_status","detail",key},"blank","Link to Detailed Data")
       //fmt.Println(" comparison route /error_detection/ajax/incident/{key} \n","link",link)
       
       display_list[index] = []string{link,description_string,ack_display,review_display,status_string,contact_time,error_time} 
       
    }
    
    return web_support.Setup_data_table("topic_list",[]string{"LINK","DESCRIPTION","ACK STATE","REVIEW STATE","STATUS","CONTACT TIME","FIRST_ERROR"},display_list)
}  






func format_time( time_stamp int64, err bool) string {
    
    if err != true{
        panic("bad time")
    }
    
    time_value   :=  time_stamp/1e9
    time_modulus :=  time_stamp%1e9
    time_unix    :=  time.Unix(int64(time_value),int64(time_modulus))
    
    return time_unix.Format(time.UnixDate)
    
}


func incident_change_review(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var input map[string]interface{}

  if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        fmt.Println(err)
        panic("BAD:")
    }

  key := input["key"].(string)
  review_state  := input["review_state"].(bool)
  
  review_string := msg_pack_utils.Pack_bool(review_state)
  incident_control.review_state.HSet(key,review_string)
  
  
  output := []byte(`"SUCCESS"`)
  
   w.Write(output) 
    
}

func incident_change_acknowlege(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var input map[string]interface{}

  if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        fmt.Println(err)
        panic("BAD:")
    }

  key := input["key"].(string)
  ack_state  := input["ack_state"].(bool)
  
  ack_string := msg_pack_utils.Pack_bool(ack_state)
  incident_control.acknowlege_state.HSet(key,ack_string)
  
  
  output := []byte(`"SUCCESS"`)
  
   w.Write(output) 
    
}


func incident_change_reset(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var input map[string]interface{}

  if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        fmt.Println(err)
        panic("BAD:")
    }

  key := input["key"].(string)
  incident_control.keys[key].last_error_time.Set(msg_pack_utils.Pack_int64(0))
  incident_control.keys[key].status.Set(msg_pack_utils.Pack_bool(true))
  incident_control.keys[key].contact_time.Set(msg_pack_utils.Pack_int64(0))
  incident_control.keys[key].last_error_data.Set(msg_pack_utils.Pack_string(""))

  

  incident_control.time.HSet(key,msg_pack_utils.Pack_int64(0))                    
  incident_control.status.HSet(key,msg_pack_utils.Pack_bool(true))                  
  incident_control.last_error_data.HSet(key,msg_pack_utils.Pack_string(""))
  incident_control.last_error_time.HSet(key,msg_pack_utils.Pack_int64(0))          
  incident_control.review_state.HSet(key,msg_pack_utils.Pack_bool(true))              
  incident_control.old_status.HSet(key,msg_pack_utils.Pack_bool(true))   
  
  fmt.Println("db delete",incident_control.incident_log.Delete_Entry(key))
            
   
  
  
  output := []byte(`"SUCCESS"`)
  
   w.Write(output) 
    
}




func incident_detail_function(w http.ResponseWriter, r *http.Request) {
   
   wd_status_template ,_ := base_templates.Clone()
   
   display_data := incident_detail_status_html(r)
   
   template.Must(wd_status_template.New("application").Parse("<center><h2>INCIDENT DETAIL DATA </h2></center><br>"+display_data))
 
    
   data := make(map[string]interface{})
   data["Title"] = "INCIDENT_DETAIL_DATA"
   wd_status_template.ExecuteTemplate(w,"bootstrap", data)
   
}

func incident_detail_status_html(r *http.Request)string { 
      vars := web_support.Get_vars(r)
      key := vars["key"]
      
      contact_time   :=    format_time(msg_pack_utils.Unpack_int64(incident_control.time.HGet(key)))
      status,_       :=    msg_pack_utils.Unpack_bool(incident_control.status.HGet(key))  
      error_data,_   :=    msg_pack_utils.Unpack_string(incident_control.last_error_data.HGet(key)) 
      error_time    :=     format_time(msg_pack_utils.Unpack_int64(incident_control.last_error_time.HGet(key)))   
      description,_  :=    msg_pack_utils.Unpack_string(incident_control.description.HGet(key))            
      review_state,_  :=   msg_pack_utils.Unpack_bool(incident_control.review_state.HGet(key))  
      ack_state,_     :=   msg_pack_utils.Unpack_bool(incident_control.acknowlege_state.HGet(key))
      status_string   := strconv.FormatBool(status)
     
      return_value := generate_java_script(key,web_support.Get_Web_Start()+"ajax/incident_status/review",
                                           web_support.Get_Web_Start()+"ajax/incident_status/acknowlege",
                                           web_support.Get_Web_Start()+"ajax/incident_status/reset")
     
     
      
      
      
      return_value = return_value + `
      <button id="reset_data" type="button">Reset Results</button><BR>
      <button id="change_review" type="button">Change Review Status</button><BR>
      <button id="change_ack" type="button">Change Acknwolege Status</button><BR>`
      
      if review_state == false {
 
      return_value = return_value+ `
      
      <input type="checkbox" id="review_state" name="review_state" value="Review"  > 
      <label for="review_state"> Data Reviewed </label><br>
      
      `
      }else{
       return_value = return_value+ `    
      
      <input type="checkbox" id="review_state" name="review_state" value="Review" checked> 
      <label for="review_state"> Data Reviewed </label><br>
      `
 
      }
      if ack_state == false {
 
      return_value = return_value+ `
      
      <input type="checkbox" id="ack_state" name="ack_state" value="ACKNOWLEGE_STATE"  > 
      <label for="review_state"> ACK Reviewed </label><br>
      
      `
      }else{
       return_value = return_value+ `    
      
      <input type="checkbox" id="review_state" name="review_state" value="ACKNOWLEGE_STATE" checked> 
      <label for="review_state"> Acknowledged </label><br>
      `
 
      
      
      
          
      }
      return_value = return_value+"<h5>Summary Data</h5>"
      
      
      list_data :=  []string{  "<ul>",
                               "<li>Key:  "+key+"</li>",
                               "<li>Description: "+description+"</li>",
                               "<li>Status: "+status_string+"</li>",
                               "<li>CONTACT TIME: "+contact_time+"</li>",
                               "<li>ERROR TIME: "+error_time+"</li>",
                               "<li>Error Data: "+fmt.Sprint(error_data) +"</ul>",
                               "</ul>" }
                               
      list_string := strings.Join(list_data,"\n")
     
     
      
      return_value = return_value+list_string
      return_value = return_value + generate_time_series(key)
 
      return return_value
}



func generate_java_script(key , url_path_1,url_path_2,url_path_3 string)string{


return_value := `    
 <script>
  key = "`+key+`"
  url_path_1 = "`+url_path_1+`"
  url_path_2 = "`+url_path_2+`" 
  url_path_3 = "`+url_path_3+`"
  $(document).ready(function(){` + web_support.Load_jquery_ajax_components() +
     ` 
     $("#change_review").click(function(){
         review_state = $("#review_state").is(':checked');
         data= {"review_state" : review_state,
                "key" :key}
         
         ajax_post(url_path_1, data,"review state changed","server error" )
         
    });
    $("#change_ack").click(function(){
         ack_state = $("#ack_state").is(':checked');
         data= {"ack_state" : ack_state,
                "key" :key}
         
         ajax_post(url_path_2, data,"review state changed","server error" )
         
    });
    
    
     $("#reset_data").click(function(){
         review_state = $("#review_state").is(':checked');
         data= { "key" :key}
         
         ajax_post(url_path_3, data,"entry is reset","server error" )
         
    });
  });
         
         
         
</script>`

return return_value
}


func generate_time_series(key string)string {
    
  return_value := "<h5>Time History Changes </h5>"
  where_clause   := "tag1 = '"+key+"'  and  time >= 0 ORDER BY time DESC LIMIT 10 "
  pg_data,status := incident_control.incident_log.Select_where(where_clause)

  if status != true {
      panic("bad time series data ")
  }
  for _,data := range pg_data {
    ts       := format_time(data.Time_stamp,true)
    //fmt.Println(data.Data)
    raw_data,err := msg_pack_utils.Unpack_string(data.Data)
    if err != true {
        raw_data = data.Data
    }
    return_value = return_value + "<UL>"
    return_value = return_value +"<li> Time:  "+ts+"</li>"
    return_value = return_value +"<li> Data:  "+ raw_data +"</li>"
    return_value = return_value +"</UL>"
  }
  return return_value
}
