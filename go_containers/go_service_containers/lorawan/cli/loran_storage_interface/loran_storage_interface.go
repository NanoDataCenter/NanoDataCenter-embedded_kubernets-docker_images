package loran_server_storage_interface



import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Get_data(url_base, app_name,url_after ,password,limit, after string ){
    
    
    raw_text_stream , err := Get_raw_stream( url_base, app_name,url_after ,password,limit, after)
    if err == false{
        fmt.Println("bad")
    }
    fmt.Println("raw_text_stream",raw_text_stream)
    
    
    
    
    
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
