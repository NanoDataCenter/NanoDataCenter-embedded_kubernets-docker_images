package main

import (
    
    "fmt"
    "net/http"
    "net/url"
	"io/ioutil"
	"time"
    "strings"
	"strconv"
    "lacima.com/site_data"
    "lacima.com/redis_support/graph_query"
    "lacima.com/redis_support/redis_handlers"
    "lacima.com/redis_support/generate_handlers"
	"lacima.com/Patterns/msgpack_2"
	"lacima.com/Patterns/logging_support"
	"lacima.com/Patterns/secrets"
    "lacima.com/server_libraries/postgres"
   
)



var  performance_log    pg_drv.Postgres_Stream_Driver


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
    
  var config_file ="/data/redis_configuration.json"
  var site_data_store map[string]interface{}

  site_data_store = get_site_data.Get_site_data(config_file)
  graph_query.Graph_support_init(&site_data_store)
  redis_handlers.Init_Redis_Mutex()
  data_handler.Data_handler_init(&site_data_store)
  secrets.Init_file_handler(site_data_store)
  
  
  Monitor_TP_Setup()
  Monitior_TP_Switch()


}


 

func Monitor_TP_Setup(){

   // find switches
   // for each switch find data structures
   
   performance_log  = logging_support.Find_stream_logging_driver()
   
    
   search_list := []string{ "TP_SWITCH"}
   switches := graph_query.Common_qs_search(&search_list)
   
   for _,element := range switches {
      var temp switch_record_type
	  temp.ip       =  graph_query.Convert_json_string(	element["ip"] ) 
	  
	  temp.name =       graph_query.Convert_json_string(	element["name"] ) 
      cat_user_password :=  secrets.Get_Secret("TP_SWITCH",temp.name)
      temp.username , temp.password = secrets.Extract_User_Password(cat_user_password)
      temp.incident_log  = logging_support.Construct_incident_log([]string{"TP_SWITCH:"+temp.name,"INCIDENT_LOG"} )
      switch_array = append( switch_array,temp)
      
      
	 
   }
   fmt.Println("switch_array",switch_array) 
   

}





func Monitior_TP_Switch() {

   for true {
      for _, element := range switch_array {
          make_measurement(&element)
      }
      time.Sleep(time.Minute*15)
   }

}


func make_measurement( element *switch_record_type ){
    
     if make_login_post(element) == false{
         return
     }
	 raw_data,err := make_collect_data_get(element)
     
	 //fmt.Println("raw_data",raw_data)
    
	 if err == true {
         preamble := "Switch Data Collection Success \n"
         (*element).incident_log.Log_data_status(true,preamble)
	     parse_raw_data(element,raw_data)
	 }else{
         
         fmt.Println("err",err)
         preamble := "Switch Data Collection Failue \n"
         (*element).incident_log.Log_data_status(false,preamble+fmt.Sprint(err))
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
         fmt.Println("err",err)
         preamble := "Switch Login Failue \n"
         (*element).incident_log.Log_data_status(false,preamble+fmt.Sprint(err))
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
     fmt.Println(err,resp)
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

      
      //fmt.Println("raw_data",raw_data)
      data_block := extract_balance_element( raw_data, "<script>", "</script>",1 )
      //fmt.Println("data_block",data_block)
      link_data := extract_balance_element( data_block, "link_status:[", "]",1 )
      link_int := turn_to_ints(link_data)
  
      pkt_data := extract_balance_element( data_block, "pkts:[", "]",1 )
      pkt_int := turn_to_ints(pkt_data)
 
      number := len(pkt_int)/4
      valid_links    := link_int[:number]
      pkt_tx_good    := make([]int64,0)
      pkt_tx_bad     := make([]int64,0)
      pkt_rx_good    := make([]int64,0)
      pkt_rx_bad     := make([]int64,0)
  
     for i:=0;i<number*4;i+=4{
        pkt_tx_good = append(pkt_tx_good,pkt_int[i])
        pkt_tx_bad  = append(pkt_tx_bad,pkt_int[i+1])
        pkt_rx_good = append(pkt_rx_good,pkt_int[i+2])
        pkt_rx_bad  = append(pkt_rx_bad,pkt_int[i+3])
  
       }
       for j :=int64(0);j<int64(len(valid_links));j++{
           i := valid_links[j]
           switch i{
	          case 5,6: 
                    fmt.Println("j",j,i,pkt_rx_bad[j],pkt_tx_bad[j])
                    j_ascii := strconv.FormatInt(j, 10)
	                performance_log.Insert( "TP_Managed_Switch",element.name,"tx",j_ascii,"",msg_pack_utils.Pack_int64(pkt_rx_bad[j]))
                    performance_log.Insert( "TP_Managed_Switch",element.name,"rx",j_ascii,"",msg_pack_utils.Pack_int64(pkt_tx_bad[j]))
		
              default:
                  ; // do nothing as there are empty links
           }
       }

}


func extract_balance_element( input_string, start_delem, end_delem string, target_element int) string {

  temp := strings.Split(input_string,start_delem)
  
  data_elements := strings.Split(temp[target_element],end_delem )
  
  return data_elements[0]

}

func turn_to_ints( input_data string) []int64 {
  var return_value []int64
  token_elements := strings.Split(input_data,",")
  for _,i := range token_elements{
    value,_ := strconv.Atoi(i)
    return_value = append(return_value,int64(value))
  }
  return return_value
}




