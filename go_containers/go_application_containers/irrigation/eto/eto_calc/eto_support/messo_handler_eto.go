package eto_support

import (
    "fmt"
     "github.com/tidwall/gjson"
    "lacima.com/redis_support/graph_query"
    "github.com/GiterLab/urllib"
    "time"
)


type Messo_ETO_TYPE struct {
	calc_type  string
	key        string
	token      string
	priority   float64
	latitude   float64
	altitude   float64
	station    string
	url        string
	eto_data   []ETO_INPUT
	eto_output  ETO_RAIN_TYPE
}

func Create_Messo_ETO(key, calc_type string, node map[string]string) interface{} {
	var meso_eto Messo_ETO_TYPE

	
	meso_eto.calc_type  = calc_type
	meso_eto.key        = key
	access_key          := graph_query.Convert_json_string(node["access_key"])
	meso_eto.token      = "&token=" + Decode_access_key(access_key)
	meso_eto.priority   = graph_query.Convert_json_float64(node["priority"])
    meso_eto.latitude   = graph_query.Convert_json_float64(node["latitude"])
    meso_eto.altitude   = graph_query.Convert_json_float64(node["altitude"])
	meso_eto.station    = graph_query.Convert_json_string(node["station"])
	meso_eto.url        = graph_query.Convert_json_string(node["url"])
    meso_eto.eto_data   =  make([]ETO_INPUT,0)
	return meso_eto
}



func (r Messo_ETO_TYPE) Compute_eto() {


	if ( ETO_Exist(r.key) == true){
		fmt.Println("***************** messo eto returning")
        return
	}

	current_day := time.Now()
	previous_day := current_day.Add(-24 * time.Hour)
	year, month, day := previous_day.Date()
	previous_date := fmt.Sprintf("%04d%02d%02d", year, month, day)
    
	year, month, day = current_day.Date()
    current_date     := fmt.Sprintf("%04d%02d%02d", year, month, day)
    
    
    start_time  := `&start=`+previous_date+`0800`
    end_time    := `&end=`+current_date+`1000`
   
    
    

    
    
    ending := "&vars=relative_humidity,air_temp,solar_radiation,peak_wind_speed,wind_speed&obtimezone=local"
    url := r.url + "?stid=" + r.station + r.token + start_time + end_time + ending

    
    req := urllib.Get(url)
	byt, err := req.Bytes()

	if err != nil {
		panic(err)
	}
	data := string(byt)
    //fmt.Println("data",data)
    //peak_windspeed_result    := gjson.Get(string(data),"STATION.0.OBSERVATIONS.peak_wind_speed_set_1")
	solar_radiation_result     := gjson.Get(string(data),"STATION.0.OBSERVATIONS.solar_radiation_set_1")
    relative_humidity_result   := gjson.Get(string(data),"STATION.0.OBSERVATIONS.relative_humidity_set_1")
    air_temp_result            := gjson.Get(string(data),"STATION.0.OBSERVATIONS.air_temp_set_1")
    windspeed_result           := gjson.Get(string(data),"STATION.0.OBSERVATIONS.wind_speed_set_1")
    
    
    
    
    
    windspeed_list            := windspeed_result.Array()
    solar_radiation_list      := solar_radiation_result.Array()
    relative_humidity_list    := relative_humidity_result.Array() 
    air_temp_list            := air_temp_result.Array()
    
    check_array(len(windspeed_list),24)
    check_array(len(solar_radiation_list),24)
    check_array(len(relative_humidity_list),24)
    check_array(len(windspeed_list),24 )
    

    
    for i:= 0; i<24;i++ {
        var eto_input  ETO_INPUT
        eto_input.Wind_speed                    = windspeed_list[i].Float()
	    eto_input.Temp_C                        = air_temp_list[i].Float()
	    eto_input.Humidity                      = relative_humidity_list[i].Float()
	    eto_input.SolarRadiationWatts_m_squared = solar_radiation_list[i].Float()
	    eto_input.Delta_timestamp               = float64(1./24.)
        r.eto_data  = append(r.eto_data,eto_input)
    }
    
    eto_result := Construt_Eto_Results(r.eto_data)
    
    eto_value := eto_result.Calculate_eto(r.altitude, r.latitude ,24) 
    fmt.Println("messo eto value",eto_value)
    r.eto_output.Key            = r.key 
    r.eto_output.Status         = true
    r.eto_output.Date_string    = previous_date
    r.eto_output.Priority       = r.priority 
    r.eto_output.Value          = eto_value
    ETO_HSet(r.key,r.eto_output)
    Stream_HSet(r.key,r.eto_data)
    
}    

func check_array(input, size int){
    if input < size  {
        //fmt.Println(input,size)
        panic("not enough data")
    }
}


