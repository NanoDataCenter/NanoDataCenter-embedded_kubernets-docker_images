package eto_support

import (
    "fmt"
     "github.com/tidwall/gjson"
    "lacima.com/redis_support/graph_query"
    "github.com/GiterLab/urllib"
    "time"
)


type Messo_RAIN_TYPE struct {
	calc_type  string
	key        string
	token      string
	station    string
	url        string
	priority   float64
    rain_output ETO_RAIN_TYPE
	
}



func Create_Messo_Rain(key, calc_type string, node map[string]string) interface{} {
	var meso_rain Messo_RAIN_TYPE

    
	
	
	
	meso_rain.calc_type  = calc_type
	meso_rain.key        = key
	access_key          := graph_query.Convert_json_string(node["access_key"])
	meso_rain.token      = "&token=" + Decode_access_key(access_key)
	meso_rain.priority   = graph_query.Convert_json_float64(node["priority"])
	meso_rain.station    = graph_query.Convert_json_string(node["station"])
	meso_rain.url        = graph_query.Convert_json_string(node["url"])
	
	return meso_rain
}

func (r Messo_RAIN_TYPE) Compute_eto() {


	if ( ETO_Exist(r.key) == true){
		fmt.Println("***************** messo rain returning")
        return
	}

	current_day := time.Now()
	previous_day := current_day.Add(-24 * time.Hour)
	year, month, day := previous_day.Date()
	previous_date := fmt.Sprintf("%04d%02d%02d", year, month, day)
    
	year, month, day = current_day.Date()
    current_date     := fmt.Sprintf("%04d%02d%02d", year, month, day)
    
   
    start_time  := `&start=`+previous_date+`0800`
    end_time    := `&end=`+current_date+`0900`
    //fmt.Println("dates",previous_date,current_date)
    


    url := r.url + "?stid=" + r.station + r.token + start_time + end_time + "&obtimezone=local" 

    //fmt.Println("url",url)
    req := urllib.Get(url)
	byt, err := req.Bytes()

	if err != nil {
		panic(err)
	}
	data := string(byt)
    //fmt.Println("data",data)
    rain_result     := gjson.Get(string(data),"STATION.0.OBSERVATIONS.total_precip_value_1")
    //fmt.Println("\nrain_result",rain_result)

    rain_value := rain_result.Float()/25.4
    
    r.rain_output.Key            = r.key 
    r.rain_output.Status         = true
    r.rain_output.Date_string    = previous_date 
    r.rain_output.Priority       = r.priority 
    r.rain_output.Value          = rain_value 

    Rain_HSet(r.key,r.rain_output)
}
    

