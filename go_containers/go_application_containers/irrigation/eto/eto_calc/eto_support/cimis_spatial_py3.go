package eto_support

import (
	"fmt"
    "github.com/tidwall/gjson"
    "lacima.com/redis_support/graph_query"
	"github.com/GiterLab/urllib"
	"time"
    "strings"
	//"strconv"
	
)



type CIMIS_ETO_SPATIAL_TYPE struct {
	calc_type  string
	key        string
	access_key string
	priority   float64
	lat        float64
	long       float64
	url        string
	eto_output  ETO_RAIN_TYPE
	
}

func Create_CIMIS_ETO_SPATIAL(key, calc_type string, node map[string]string) interface{} {
	var cimis_eto_spacial CIMIS_ETO_SPATIAL_TYPE
    
	cimis_eto_spacial.calc_type  = calc_type
	cimis_eto_spacial.key        = key
	cimis_eto_spacial.access_key = graph_query.Convert_json_string(node["access_key"])
	cimis_eto_spacial.access_key = "appKey=" + Decode_access_key(cimis_eto_spacial.access_key)
	cimis_eto_spacial.priority   = graph_query.Convert_json_float64(node["priority"])
	cimis_eto_spacial.lat        = graph_query.Convert_json_float64(node["latitude"])
    cimis_eto_spacial.long       = graph_query.Convert_json_float64(node["longitude"])
	cimis_eto_spacial.url        = graph_query.Convert_json_string(node["url"])
	//fmt.Println(cimis_eto_spacial)
	
	return cimis_eto_spacial
}

func (r CIMIS_ETO_SPATIAL_TYPE) Compute_eto() {
	if ( ETO_Exist(r.key) == true){
		fmt.Println("***************** cimis eto sat  returning")
        return
	}
    
	current_day := time.Now()
	previous_day := current_day.Add(-24 * time.Hour)
	year, month, day := previous_day.Date()
	date := fmt.Sprintf("%04d-%02d-%02d", year, month, day)

	lat_long := fmt.Sprintf(`lat=%f,lng=%f`,r.lat,r.long)
    lat_long = strings.ReplaceAll(lat_long," ","")
	url := r.url + "?" + r.access_key + "&targets=" + lat_long + "&startDate=" + date + "&endDate=" + date

    
 
	
	req := urllib.Get(url)
	byt, err := req.Bytes()
    
	if err != nil {
		panic(err)
	}
	
	value1 := gjson.Get(string(byt),"Data.Providers.0.Records.0.DayAsceEto.Value")
	eto_value := value1.Float()
    r.eto_output.Key            = r.key 
    r.eto_output.Status         = true
    r.eto_output.Date_string    = date
    r.eto_output.Priority       = r.priority 
    r.eto_output.Value          = eto_value
	
	ETO_HSet( r.key, r.eto_output   )

}

