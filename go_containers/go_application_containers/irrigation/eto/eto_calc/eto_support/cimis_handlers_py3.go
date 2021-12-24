package eto_support

import (
    
    "fmt"
    "strconv"
    "time"
    "encoding/json"
    "github.com/GiterLab/urllib"
    "lacima.com/redis_support/redis_handlers"
    "lacima.com/redis_support/graph_query"
    
)

/*
cimis_setup map[access_key:"ETO_CIMIS" altitude:2400 name:"ETO_CIMIS" namespace:"[SYSTEM:farm_system][SITE:LACIMA_SITE][WEATHER_STATIONS:WEATHER_STATIONS][WEATHER_STATION:ETO_CIMIS]" priority:2 station:"62" sub_id:"" type:"CIMIS" url:"http://et.water.ca.gov/api/data"]
*/

type CIMIS_ETO_TYPE struct {
    calc_type      string
    key            string
    access_key     string
    priority       float64
    station        string 
    url            string
    eto_data       redis_handlers.Redis_Hash_Struct
    rain_data      redis_handlers.Redis_Hash_Struct
    
}


func Create_CIMIS_ETO( key,calc_type string, node map[string]string, eto_data,rain_data redis_handlers.Redis_Hash_Struct   )interface{}{
  var cimis_eto  CIMIS_ETO_TYPE
  
  cimis_eto.calc_type  = calc_type
  cimis_eto.key        = key
  cimis_eto.access_key = graph_query.Convert_json_string(node["access_key"])
  cimis_eto.access_key = "appKey=" +Decode_access_key(cimis_eto.access_key)
  cimis_eto.priority   = graph_query.Convert_json_float64(node["priority"])
  cimis_eto.station    = graph_query.Convert_json_string(node["station"]) 
  cimis_eto.url        = graph_query.Convert_json_string(node["url"])
  cimis_eto.eto_data   = eto_data
  cimis_eto.rain_data  = rain_data
  fmt.Println("cimis_eto",cimis_eto)
  return cimis_eto  
}






func (r CIMIS_ETO_TYPE)Compute_eto(){
    
   if r.eto_data.HExists(r.key ) == true {
      fmt.Println("***************** cimis eto returning")
   }
   
   current_day     := time.Now()
   previous_day    := current_day.Add(-24*time.Hour)
   year,month,day  := previous_day.Date()
   date            := fmt.Sprintf("%04d-%02d-%02d",year,month,day)
   
     
   url :=  fmt.Sprintf(`%s?%s&targets=%s&startDate=%s&endDate=%s`,r.url,r.access_key,r.station,date,date)
   fmt.Println("url",url)
   
   
   //url :== r.cimis_url + "?" + r.app_key + "&targets=" + 
   //         str(r.station) + "&startDate=" + date + "&endDate=" + date
   
   req := urllib.Get(url)
   byt, err := req.Bytes()
   
   
  

   if err != nil {
        panic(err)
	}
	
    var dat map[string]interface{}
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    
    
    
    data := dat["Data"]  //   ["Providers"][0]["Records"][0]['DayAsceEto']["Value"].(string)
    temp := data.(map[string]interface{})
    providers := temp["Providers"]
    temp1  := providers.([]interface{})
    provider_0 := temp1[0]
    
    
    //fmt.Println("provider",provider_0)
    temp2 := provider_0.(map[string]interface{})
    records := temp2["Records"]
    temp3 :=  records.([]interface{})
    record0 := temp3[0]
    fmt.Println("record0",record0)
    temp4  := record0.(map[string]interface{})
    eto_map := temp4["DayAsceEto"]
    fmt.Println("eto_map",eto_map)
    temp5 := eto_map.(map[string]interface{})
    eto_value_a := temp5["Value"]
    fmt.Println("eto_value",eto_value_a)
    eto_value_string := eto_value_a.(string)
    eto_value, err := strconv.ParseFloat(eto_value_string, 32)
    if err != nil {
        panic("bad float conversion")
    }
    fmt.Println("eto_value",eto_value)
    
    panic("done")
/* 
       need to parse data
   
       
        value = float(data["Data"]["Providers"][0]["Records"][0]['DayAsceEto']["Value"])
        print("value",value)
        print("*************** cimis made it here",value)
        date_string = str(datetime.datetime.now())
        self.eto_data.hset("CIMIS:"+str(self.station), {"eto":float(
                data["Data"]["Providers"][0]["Records"][0]['DayAsceEto']["Value"]),
                "priority":self.cimis_data["priority"],"status":"OK","time": date_string })
        self.rain_data.hset("CIMIS:"+str(self.station), {"rain":float(
                data["Data"]["Providers"][0]["Records"][0]['DayPrecip']["Value"]),
                "priority":self.cimis_data["priority"],"status":"OK","time": date_string })  
    
    */
}
