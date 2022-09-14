package loran_server_storage_interface



import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Get_data(url_base, app_name,url_after ,password,limit, after string )([]Loran_raw_data_stream, bool){
    
    return_value := make([]Loran_raw_data_stream,0)
    raw_text_stream , err := Get_raw_stream( url_base, app_name,url_after ,password,limit, after)
    if err == false{
        return return_value, false
    }
    for _, raw_text_record := range raw_text_stream {
          var item  Loran_raw_data_stream
          return_value.app_id              := extract_string_field(
          return_value.device_id        := extract_string_field(
          return_value.f_port              :=  extract_int_field(
          return_value.f_cnt                :=  extract_int_field(
          return_value.raw_payload   :=  extract_string_field(
          return_value.meta_data       :=  extract_string_field(
          return_value.time_stamp     :=  extract_string_field(
          
        
        
    }
    
    
    
    
    
}


func Get_raw_stream( url_base, app_name,url_after ,password,limit, after string )( string,bool){
	client := &http.Client{}
	request_string :=  url_base+app_name+url_after+"limit="+limit+"&after="+after
	fmt.Println("request_string",request_string)
	req, err := http.NewRequest("GET", request_string, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+password )
	req.Header.Set("Accept", "text/event-stream")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
    return string(bodyText),true
}
