package main

import (
    
    "fmt"
    "net/http"
    "net/url"
	"io/ioutil"
	"time"
    "strings"
	"strconv"
	//"reflect"
	"bytes"
    "lacima.com/site_data"
    "lacima.com/redis_support/graph_query"
    "lacima.com/redis_support/redis_handlers"
    "lacima.com/redis_support/generate_handlers"
	"github.com/msgpack/msgpack-go"
	"lacima.com/Patterns/logging_support"
	"lacima.com/Patterns/secrets"
)






type switch_record_type  struct  {
  name     string
  ip       string
  username string
  password string
  incident_log             *logging_support.Incident_Log_Type
  


}


var switch_array []switch_record_type


var passwords map[string]map[string]string

func main() {
    
  var config_file = "/data/redis_server.json"
  var site_data_store map[string]interface{}

  site_data_store = get_site_data.Get_site_data(config_file)
  graph_query.Graph_support_init(&site_data_store)
  redis_handlers.Init_Redis_Mutex()
  data_handler.Data_handler_init(&site_data_store)
  secrets.Init_file_handler(site_data_store)
  passwords = secrets.Get_Secret("TP_MANAGED_SWITCHES")
  //fmt.Println("secrets",passwords)
  Monitor_TP_Setup()
  Monitior_TP_Switch()


}




func Monitor_TP_Setup(){

   // find switches
   // for each switch find data structures
   
   
   
    
   search_list := []string{ "TP_SWITCH"}
   switches := graph_query.Common_qs_search(&search_list)
   
   for _,element := range switches {
      var temp switch_record_type
	  temp.ip       =  graph_query.Convert_json_string(	element["ip"] ) 
	  
	  temp.name =       graph_query.Convert_json_string(	element["name"] ) 
	  temp.username = passwords[temp.name]["user"]
	  temp.password = passwords[temp.name]["password"]
	  temp.incident_log  = logging_support.Construct_incident_log([]string{"TP_SWITCH:"+temp.name,"INCIDENT_LOG"} )
      switch_array = append( switch_array,temp)
	  //fmt.Println(temp)
	 
   }


}





func Monitior_TP_Switch() {

   for true {
      for _, element := range switch_array {
          make_measurement(&element)
      }
      time.Sleep(time.Minute*15)
   }

}

func handlepanic() {
  
    if a := recover(); a != nil {
      
        fmt.Println("RECOVER", a)
    }
}
func make_measurement( element *switch_record_type ){
     defer handlepanic()
     make_login_post(element)
	 raw_data,err := make_collect_data_get(element)
	
	 if err == true {
	   parse_raw_data(element,raw_data)
	 }
	 
}




func make_login_post( element *switch_record_type ) bool{

    var return_value bool
	var user = (*element).username
	var password = (*element).password
	
	
	

	client := http.Client{Timeout: 5 * time.Second, }

     data := url.Values{
        "logon":       {"Login"},
        "username": {user},
		"password": {password},
    }
    target_url := "http://"+(*element).ip+"/logon.cgi"
    _, err := client.PostForm(target_url, data)
 
   
    if err != nil {
        return_value = false
    }else{
	  return_value = true
	}
	return return_value
	
}

func make_collect_data_get(element *switch_record_type )(string,bool){

    var success bool
    target_url := "http://"+(*element).ip+ "/PortStatisticsRpm.htm"
	
	client := http.Client{Timeout: 5 * time.Second, }
    resp, err := client.Get(target_url)
     //fmt.Println(err,resp)
    if err != nil {
        success = false
    } else{
	   success = true
	}
	

    defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
    s := string(bodyText)
    
	return s,success
}



func parse_raw_data(element *switch_record_type,raw_data string ) {

   defer recover()
  //fmt.Println("raw_data",raw_data)
  data_block := extract_balance_element( raw_data, "<script>", "</script>",1 )
  //fmt.Println("data_block",data_block)
  link_data := extract_balance_element( data_block, "link_status:[", "]",1 )
  link_int := turn_to_ints(link_data)
  
  pkt_data := extract_balance_element( data_block, "pkts:[", "]",1 )
  pkt_int := turn_to_ints(pkt_data)
 
  number := len(pkt_int)/4
  valid_links := link_int[:number]
  var pkt_tx_good []int
  var pkt_tx_bad  []int
  var pkt_rx_good []int
  var pkt_rx_bad  []int
  
  for i:=0;i<number*4;i+=4{
     pkt_tx_good = append(pkt_tx_good,pkt_int[i])
     pkt_tx_bad  = append(pkt_tx_bad,pkt_int[i+1])
     pkt_rx_good = append(pkt_rx_good,pkt_int[i+2])
     pkt_rx_bad  = append(pkt_rx_bad,pkt_int[i+3])
  
  }
  log_data:= make(map[string]interface{})
  
  log_data["pkt_tx_bad"]  =  pkt_tx_bad
  log_data["pkt_rx_bad"]  = pkt_rx_bad
  log_data["valid_links"] = valid_links
  //fmt.Println("log_data",log_data)
  
  var b bytes.Buffer	
  msgpack.Pack(&b,log_data)
  current_value := b.String()
  
   
 
  
  var ok_flag = true 
  element_tx := log_data["pkt_tx_bad"].([]int)
  element_rx := log_data["pkt_rx_bad"].([]int)
  switch_links := log_data["valid_links"].([]int)
  for _,i := range switch_links{
    switch i{
	   case 5,6: {
         //fmt.Println("i",i,element_tx[i],element_rx[i])
	     if (element_tx[i] > 0 ) || (element_rx[i] > 0 ) {
		    ok_flag = false
		} // if
		}// case
		}// switch
   }// for

  (*element).incident_log.Log_data( ok_flag,  current_value, current_value )
  fmt.Println("Message logged",ok_flag)

}

func log_differences( number int, a,b string){

  for i:= 0;i<number;i++{
    if a[i] != b[i]{
	  fmt.Println("bad",a[i],b[i])
	}
  }
}

func extract_balance_element( input_string, start_delem, end_delem string, target_element int) string {

  temp := strings.Split(input_string,start_delem)
  
  data_elements := strings.Split(temp[target_element],end_delem )
  
  return data_elements[0]

}

func turn_to_ints( input_data string) []int {
  var return_value []int
  token_elements := strings.Split(input_data,",")
  for _,i := range token_elements{
    value,_ := strconv.Atoi(i)
    return_value = append(return_value,value)
  }
  return return_value
}




