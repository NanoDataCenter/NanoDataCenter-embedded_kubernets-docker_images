package eto_support

import (
	//"encoding/json"
    "github.com/tidwall/gjson"
	"fmt"
	"github.com/GiterLab/urllib"
	"lacima.com/redis_support/graph_query"
	
	//"strconv"
	"time"
)

/*
cimis_setup map[access_key:"ETO_CIMIS" altitude:2400 name:"ETO_CIMIS" namespace:"[SYSTEM:farm_system][SITE:LACIMA_SITE][WEATHER_STATIONS:WEATHER_STATIONS][WEATHER_STATION:ETO_CIMIS]" priority:2 station:"62" sub_id:"" type:"CIMIS" url:"http://et.water.ca.gov/api/data"]
*/

type CIMIS_ETO_TYPE struct {
	calc_type  string
	key        string
	access_key string
	priority   float64
	station    string
	url        string
	eto_data    []ETO_INPUT
	rain_output ETO_RAIN_TYPE
	eto_output  ETO_RAIN_TYPE
}

func Create_CIMIS_ETO(key, calc_type string, node map[string]string) interface{} {
	var cimis_eto CIMIS_ETO_TYPE

	cimis_eto.calc_type = calc_type
	cimis_eto.key = key
	cimis_eto.access_key = graph_query.Convert_json_string(node["access_key"])
	cimis_eto.access_key = "appKey=" + Decode_access_key(cimis_eto.access_key)
	cimis_eto.priority = graph_query.Convert_json_float64(node["priority"])
	cimis_eto.station = graph_query.Convert_json_string(node["station"])
	cimis_eto.url = graph_query.Convert_json_string(node["url"])
	
	
	return cimis_eto
}

func (r CIMIS_ETO_TYPE) Compute_eto() {

	if ( ETO_Exist(r.key) == true)&& ( Rain_Exist(r.key) == true){
		fmt.Println("***************** cimis  returning")
        return
	}
	

	current_day := time.Now()
	previous_day := current_day.Add(-24 * time.Hour)
	year, month, day := previous_day.Date()
	date := fmt.Sprintf("%04d-%02d-%02d", year, month, day)

	url := fmt.Sprintf(`%s?%s&targets=%s&startDate=%s&endDate=%s`, r.url, r.access_key, r.station, date, date)
	

	
	req := urllib.Get(url)
	byt, err := req.Bytes()

	if err != nil {
		panic(err)
	}
	value1 := gjson.Get(string(byt),"Data.Providers.0.Records.0.DayAsceEto.Value")
	eto_value := value1.Float()
    
	
    value2 := gjson.Get(string(byt),"Data.Providers.0.Records.0.DayPrecip.Value")
	
    rain_value := value2.Float()
   
	r.eto_output.Key            = r.key 
    r.eto_output.Status         = true
    r.eto_output.Date_string    = date
    r.eto_output.Priority       = r.priority 
    r.eto_output.Value          = eto_value
    r.rain_output.Key            = r.key 
    r.rain_output.Status         = true
    r.rain_output.Date_string    = date 
    r.rain_output.Priority       = r.priority 
    r.rain_output.Value          = rain_value 
    ETO_HSet(r.key,r.eto_output)
    Rain_HSet(r.key,r.rain_output)
    
}
