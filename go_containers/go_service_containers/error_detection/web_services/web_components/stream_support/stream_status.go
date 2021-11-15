package stream_support


import (
    "fmt"
    "strings"
    //"strconv"
    "sort"
    "time"
    "net/http"
     "net/url"
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

 
    
    
 var base_templates  *template.Template

func  Stream_status_init(input *template.Template){
   
   base_templates = input
   init_stream_data_structures()

    web_support.Generate_special_route("stream_status/raw" ,stream_raw_detail)
    web_support.Generate_special_route("stream_status/filtered" ,stream_filtered_detail)
   //web_support.Generate_special_post_route("incident_status/review" , incident_change_review)
   //web_support.Generate_special_post_route("incident_status/reset" , incident_change_reset)
}
   
   


func Stream_status  (w http.ResponseWriter, r *http.Request) {
   
   incident_status_template ,_ := base_templates.Clone()
   
   display_data := stream_generate_html()
   
   template.Must(incident_status_template.New("application").Parse("<center><h2>Stream Status Data </h2></center><br>"+display_data))
 
    
   data := make(map[string]interface{})
   data["Title"] = "Stream Status"
   incident_status_template.ExecuteTemplate(w,"bootstrap",data )
   
}



func stream_generate_html()string{ 

   
  
    keys                 := stream_control.stream_table.HKeys()
    sort.Strings(keys)
    
    //old_status           := incident_control.old_status.HGetAll()
    //review               := incident_control.review_state.HGetAll()
    //status               := incident_control.status.HGetAll()
    
    ref_time               := time.Now().UnixNano()
   
    
    display_list := make([][]string,0)    
    for _,key             := range keys {
       
       latest_result,state := get_stream_processing_data(key)
       if state == true {
           key_encoded      :=  url.QueryEscape(key)
           parameters       :=  "key="+key_encoded
           link1             :=  web_support.Generate_ajax_anchor_with_parameters_and_target([]string{"stream_status","raw"},parameters,"_blank","RAW DATA")
           link2             :=  web_support.Generate_ajax_anchor_with_parameters_and_target([]string{"stream_status","filtered"},parameters,"_blank","FILTERED DATA")
           current_value    :=  fmt.Sprintf("%f",latest_result.median.current_value) 
           filtered_value   :=  fmt.Sprintf("%f",latest_result.median.filtered_value)
           velocity_value   :=  fmt.Sprintf("%f",latest_result.velocity.current_velocity)
           std_value        :=  fmt.Sprintf("%f",latest_result.z_data.std)
           
           z_value          :=  fmt.Sprintf("%f",latest_result.z_data.z_value)
           key_list         :=  strings.Split(key,"~+~")
           key_display      :=  strings.Join(key_list,"~")
           current_time,raw_time     :=  get_current_time(key)
           
           
           if ref_time -raw_time  < 3600 *4*1e9 {
               display_list = append(display_list, []string{link1,link2,key_display ,current_value,filtered_value,velocity_value,std_value,z_value,current_time}) 
        
               
        }
               
        }
    }
    
    return web_support.Setup_data_table("stream_list",[]string{"RAW Data","FILTERED DATA","KEY","VALUE","FILTER_VALUE","VELOCITY","STD","Z_VALUE","TIME" },display_list)
}  




func get_stream_processing_data(key_string string)( Stream_Processing_Type, bool){
    
    var return_value Stream_Processing_Type
    var state  bool

    state = false
    intermediate_data := make(map[string]interface{})
    packed_data := stream_control.stream_table.HGet(key_string)
    
    packed_byte := []byte(packed_data)
    err := msgpack.Unmarshal(packed_byte, &intermediate_data )
    //fmt.Println("err",err)
    if err == nil {
      
        state = true
	    return_value = recover_intermediate_values(intermediate_data)
        
    } 
    
    return return_value,state
    
    
}

func recover_intermediate_values( data map[string]interface{})Stream_Processing_Type{
 
    var return_value Stream_Processing_Type
    
    return_value.z_data.z_value                  = data["data_z"].(float64) 
    return_value.z_data.std                      = data["std"].(float64)
    return_value.z_data.z_state                  = data["z_state"].(bool)
    return_value.velocity.previous_value         = data["previous_value"].(float64)
    return_value.velocity.current_velocity       = data["current_velocity"].(float64)   
    return_value.velocity.lag_velocity           = data["lag_velocity"].(float64)  
    return_value.velocity.r_value                = data["r_value"].(float64)
    return_value.median.current_value            = data["current_value"].(float64)  
    return_value.median.buffer_position          = data["buffer_position"].(int64)
    return_value.median.buffer_limit             = data["buffer_limit"].(int64)  
    return_value.median.filtered_value           = data["filtered_value"].(float64)
    temp_buffer                                  := data["median_buffer"].([]interface{})
    return_value.median.median_buffer            = make([]float64,len(temp_buffer))
    for index, value := range temp_buffer{
        return_value.median.median_buffer[index] = value.(float64)
    }
    return return_value
    
    
}

func get_current_time(key string)(string,int64) {
    
    
    var intermediate_data int64
   
    intermediate_data = 0
    return_value := "NOT SPECIFIED"
    
    packed_data := stream_control.time_table.HGet(key)
    
    packed_byte := []byte(packed_data)
    err := msgpack.Unmarshal(packed_byte, &intermediate_data )
    //fmt.Println("err",err)
    if err == nil {
      
        
	    return_value = format_time(intermediate_data,true)
        
    } 
    
    return return_value, intermediate_data 
    
    
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
