package eto_support

import (
    "fmt"
    "time"
    "lacima.com/redis_support/graph_query"   
    
    
)



type Hybrid_ETO_TYPE struct {
	calc_type         string
	key               string
	priority          float64
	latitude          float64
	altitude          float64
    base_type         string
    base_sub_type     string
    variant_type      string
    variant_sub_type  string
    variant_fields    []string
	eto_data          []ETO_INPUT
	eto_output        ETO_RAIN_TYPE
	
}


func Create_Hybrid_ETO(key, calc_type string, node map[string]string) interface{} {
	var hybrid_eto Hybrid_ETO_TYPE
    //fmt.Println("node",node)
	hybrid_eto.calc_type        = calc_type
	hybrid_eto.key              = key
	hybrid_eto.priority         =  graph_query.Convert_json_float64(node["priority"])
    hybrid_eto.latitude         =  graph_query.Convert_json_float64(node["latitude"])
    hybrid_eto.altitude         =  graph_query.Convert_json_float64(node["altitude"])
    hybrid_eto.base_type        =  graph_query.Convert_json_string(node["base_type"])
    hybrid_eto.base_sub_type    =  graph_query.Convert_json_string(node["base_sub_type"])
    hybrid_eto.variant_type     =  graph_query.Convert_json_string(node["variant_type"])
    hybrid_eto.variant_sub_type =  graph_query.Convert_json_string(node["variant_sub_type"])
    hybrid_eto.variant_fields   =  graph_query.Convert_json_string_array(node["variant_fields"])
	hybrid_eto.eto_data          =    make([]ETO_INPUT,0) 
	return hybrid_eto
}



 

func (r Hybrid_ETO_TYPE ) Compute_eto() {
	if ( ETO_Exist(r.key) == true){
		fmt.Println("***************** hybrid returning "+r.key)
        return
	}

	current_day := time.Now()
	previous_day := current_day.Add(-24 * time.Hour)
	year, month, day := previous_day.Date()
	previous_date := fmt.Sprintf("%04d%02d%02d", year, month, day) 
    // assemble base key
    base_key := r.base_type +":"+r.base_sub_type
    fmt.Println("base key",base_key)
    base_data := Stream_HGet(base_key)
    variant_key := r.variant_type +":"+r.variant_sub_type
    variant_data :=  Stream_HGet(variant_key)
    
    merged_map := r.merge_map(base_data,variant_data,r.variant_fields)
    //fmt.Println("merge_map",merged_map)
    
    merged_eto := Map_array_to_eto_array(24,merged_map)
    
    
    
    r.eto_data  = merged_eto
    eto_result := Construt_Eto_Results(r.eto_data)
    eto_value := eto_result.Calculate_eto(r.altitude, r.latitude ,24) 
    r.eto_output.Key            = r.key 
    r.eto_output.Status         = true
    r.eto_output.Date_string    = previous_date
    r.eto_output.Priority       = r.priority 
    r.eto_output.Value          = eto_value
    fmt.Println("hybrid_eto_value",r.eto_output.Value)
    ETO_HSet(r.key,r.eto_output )
    Stream_HSet(r.key,r.eto_data)
    
   
   
}

func (r Hybrid_ETO_TYPE ) merge_map(base_data,variant_data map[string][]float64, variant_fields []string)map[string][]float64{
    
    // check array lengths
    
    
    for _,field := range variant_fields {
        if _,ok := variant_data[field]; ok == false {
            panic("bad field "+field)
        }
        base_data[field] = variant_data[field]
    }
    return base_data

}
