package eto_support

import (
    "fmt"
     "github.com/tidwall/gjson"
    "lacima.com/redis_support/graph_query"
    "github.com/GiterLab/urllib"
    "time"
)


type Wunder_ETO_TYPE struct {
	calc_type  string
	key        string
	token      string
	priority   float64
	latitude   float64
	altitude   float64
	pws        string
	url        string
	eto_data   []ETO_INPUT
    rain_output ETO_RAIN_TYPE
	eto_output  ETO_RAIN_TYPE
	
}


func Create_Wunder_ETO(key, calc_type string, node map[string]string) interface{} {
	var wunder_eto Wunder_ETO_TYPE

    
     
	
	
	wunder_eto.calc_type  = calc_type
	wunder_eto.key        = key
	access_key          := graph_query.Convert_json_string(node["access_key"])
	wunder_eto.token      = Decode_access_key(access_key)
	wunder_eto.priority   = graph_query.Convert_json_float64(node["priority"])
    wunder_eto.latitude   = graph_query.Convert_json_float64(node["lat"])
    wunder_eto.altitude   = graph_query.Convert_json_float64(node["alt"])
	wunder_eto.pws        = graph_query.Convert_json_string(node["sub_id"])
	wunder_eto.url        = `https://api.weather.com/v2/pws/observations/hourly/7day?`
	wunder_eto.eto_data   = make([]ETO_INPUT,0) 
	return wunder_eto
}

func (r Wunder_ETO_TYPE ) Compute_eto() {



    
    
	if ( ETO_Exist(r.key) == true){
		fmt.Println("***************** wunder eto returning")
        return
	}
	
    current_day := time.Now()
    end_year, end_month, end_day  := current_day.Date()
    previous_day := current_day.Add(-24 * time.Hour)
    start_year,start_month,start_day := previous_day.Date()
    
    previous_date := fmt.Sprintf("%04d%02d%02d", end_year, end_month, end_day)
    start_time := time.Date(start_year, start_month, start_day,0,0, 0, 0, time.Local)
    end_time   := time.Date(end_year,end_month,end_day, 2, 0, 0, 0, time.Local) 
    
    start_unix_time := start_time.Unix()
    end_unix_time   := end_time.Unix()
    
	url := r.url + "stationId="+ r.pws+ "&format=json&units=e&apiKey="+r.token
	
    req := urllib.Get(url)
	byt, err := req.Bytes()

	if err != nil {
		panic(err)
	}
	//fmt.Println("byt",string(byt))
	data :=  gjson.Get(string(byt),"observations")
    data_array  := data.Array()
 
    
    rain_value := float64(0)
    eto_value  := float64(0)
    for _,value := range data_array{
       value_map := value.Map()
       epoch_time := value_map["epoch"].Int()
       
      if epoch_time > end_unix_time {
          break
      }
      if epoch_time >= start_unix_time {
          var eto_input ETO_INPUT          
          eto_input.Humidity                      = value_map["humidityAvg"].Float()
          eto_input.SolarRadiationWatts_m_squared = value_map["solarRadiationHigh"].Float()
          eto_input.Delta_timestamp               = float64(1./24.)
          temp                                    := value_map["imperial"].Map()
          eto_input.Temp_C                        =  5.*(temp["tempAvg"].Float()-32.)/9.
          eto_input.Wind_speed                     =  temp["windspeedAvg"].Float()*0.44704
          r.eto_data                               = append(r.eto_data,eto_input)
          rain_value                              = temp["precipTotal"].Float()
          
      }
       
    }
    calculator := Construt_Eto_Results(r.eto_data)
    eto_value  = calculator.Calculate_eto(r.altitude,r.latitude,24)
    fmt.Println("wunder",eto_value,rain_value)
    r.eto_output.Key            = r.key 
    r.eto_output.Status         = true
    r.eto_output.Date_string    = previous_date
    r.eto_output.Priority       = r.priority 
    r.eto_output.Value          = eto_value
    r.rain_output.Key            = r.key 
    r.rain_output.Status         = true
    r.rain_output.Date_string    = previous_date 
    r.rain_output.Priority       = r.priority 
    r.rain_output.Value          = rain_value 
    ETO_HSet(r.key,r.eto_output)
    Rain_HSet(r.key,r.rain_output)
    Stream_HSet(r.key,r.eto_data)
    //fmt.Println(r.eto_data)
    
}


