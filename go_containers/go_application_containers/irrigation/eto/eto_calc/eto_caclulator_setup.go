package eto_calc

import (
    "fmt"
    "time"
    "lacima.com/redis_support/graph_query"
    "lacima.com/go_application_containers/irrigation/eto/eto_calc/eto_support"
)    

type ETO_CALC_SETUP_TYPE func(string, string, map[string]string) interface{}

type ETO_CALC_TYPE func(interface{})

var eto_calc_data  map[string]ETO_Data_Lookup_Type


var eto_calc_setup map[string]ETO_CALC_SETUP_TYPE

var eto_calc_compute map[string]ETO_CALC_TYPE



var eto_calculators map[string]interface{}
var eto_key_calc_data map[string]interface{}




func initialize_eto_calc_setup() {
	eto_calculators = make(map[string]interface{})
	eto_calc_setup = make(map[string]ETO_CALC_SETUP_TYPE)
	eto_calc_setup["MESSO_ETO"] = messo_eto_setup
	eto_calc_setup["CIMIS"] = cimis_setup
	eto_calc_setup["MESSO_RAIN"] = messo_rain_setup
	eto_calc_setup["CIMIS_SAT"] = cimis_sat_setup
	eto_calc_setup["WUNDERGROUND"] = wunderground_setup
	eto_calc_setup["HYBRID_SETUP"] = hybrid_setup
	eto_calc_compute = make(map[string]ETO_CALC_TYPE)
	eto_calc_compute["MESSO_ETO"] = messo_eto_calc
	eto_calc_compute["CIMIS"] = cimis_calc
	eto_calc_compute["MESSO_RAIN"] = messo_rain_calc
	eto_calc_compute["CIMIS_SAT"] = cimis_sat_calc
	eto_calc_compute["WUNDERGROUND"] = wunderground_calc
	eto_calc_compute["HYBRID_SETUP"] = hybrid_calc
	eto_key_calc_data = make(map[string]interface{})
	
}



func messo_eto_setup(key, calc_type string, input map[string]string) interface{} {
	//fmt.Println("messo_eto_setup",input)
	return eto_support.Create_Messo_ETO(key, calc_type, input)
}

func cimis_setup(key, calc_type string, input map[string]string) interface{} {
	//fmt.Println("cimis_setup", input)

	return eto_support.Create_CIMIS_ETO(key, calc_type, input)
}
func messo_rain_setup(key, calc_type string, input map[string]string)interface{}  {
	//fmt.Println("messo_rain_setup",input)
	return eto_support.Create_Messo_Rain(key, calc_type, input)
	
}
func cimis_sat_setup(key, calc_type string, input map[string]string) interface{} {
	//fmt.Println("cimis_sat_setup",input)

    return eto_support.Create_CIMIS_ETO_SPATIAL(key, calc_type, input)

}

func wunderground_setup(key, calc_type string, input map[string]string) interface{} {
	return eto_support.Create_Wunder_ETO(key,calc_type,input)
}
func hybrid_setup(key, calc_type string, input map[string]string) interface{} {
	//fmt.Println("hybrid_setup",input)
	return eto_support.Create_Hybrid_ETO(key, calc_type, input)
}

func messo_eto_calc(input interface{}) {
	//fmt.Println("messo_eto_calc")
    data := input.(eto_support.Messo_ETO_TYPE)
	data.Compute_eto()
    
    
}


func cimis_calc(input interface{}) {
	//fmt.Println("cimis_calc")
	data := input.(eto_support.CIMIS_ETO_TYPE)
	data.Compute_eto()
   
}
func messo_rain_calc(input interface{}) {
	//fmt.Println("messo_rain_calc")
    data := input.(eto_support.Messo_RAIN_TYPE)
	data.Compute_eto()
   
   
}
func cimis_sat_calc(input interface{}) {
	//fmt.Println("cimis_sat_calc")
    data := input.(eto_support.CIMIS_ETO_SPATIAL_TYPE)
    data.Compute_eto()
    
}
func wunderground_calc(input interface{}) {
	fmt.Println("wunderground_calc")
    data := input.(eto_support.Wunder_ETO_TYPE )
    data.Compute_eto()
   
}
func hybrid_calc(input interface{}) {
	fmt.Println("hybrid_calc")
    data := input.(eto_support.Hybrid_ETO_TYPE  )
    data.Compute_eto()
   

}




func setup_eto_calculators() {
	search_list := []string{"WEATHER_STATIONS:WEATHER_STATIONS", "WEATHER_STATION"}
	nodes := graph_query.Common_qs_search(&search_list)
    
	for _, node := range nodes {
		ws_type     := graph_query.Convert_json_string(node["type"])
		sub_id      := graph_query.Convert_json_string(node["sub_id"])
        hybrid_type := graph_query.Convert_json_int(node["hybrid_flag"])
		value, ok := eto_calc_setup[ws_type]
		if ok == false {
			panic("bad calculator type")
		}
		key := ws_type + ":" + sub_id
		
		hybrid_flag := false
		if hybrid_type >0 {
            hybrid_flag = true
        }
		
		eto_data_value            := make(map[string]interface{})
        eto_data_value["hybrid"]  =  hybrid_flag
        eto_data_value["calc_id"] =  ws_type
        eto_data_value["data"]    = value(key,ws_type,node)
		eto_key_calc_data[key]    = eto_data_value
        //fmt.Println("key",key)
	}
	

	
}



func do_calculators() {
   do_non_hybrid_calculators()
   do_hybrid_calculators()
}

func do_non_hybrid_calculators(){
 	for key, data := range eto_key_calc_data {
        //fmt.Println("key",key,data)
        do_non_hybrid_calculator(key,data)

	}   
    
}

func do_hybrid_calculators(){
 	for key, data := range eto_key_calc_data {
        //fmt.Println("key",key,data)
        do_hybrid_calculator(key,data)

	}   
    

}


var ref_key string
func calculator_error(){

  if r := recover(); r != nil {
     year, month, day := time.Now().Date()
	 date := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
     fmt.Println(ref_key, date+"\n"+fmt.Sprint(r) )
     eto_exceptions.HSet(ref_key, date+"\n"+fmt.Sprint(r) )
  }
    
}


func do_non_hybrid_calculator( key string ,data interface{}){

   defer calculator_error()
   ref_key  = key
   data_map  :=  data.(map[string]interface{})
   hybrid_type := data_map["hybrid"].(bool)
   if hybrid_type == false {
       fmt.Println("non_hybrid",hybrid_type)
       ws_type   :=  data_map["calc_id"].(string)
       calc_data := data_map["data"].(interface{})
       eto_calc_compute[ws_type](calc_data)
   }
   
}

func do_hybrid_calculator( key string ,data interface{}){

   defer calculator_error()
   ref_key  = key
   data_map  :=  data.(map[string]interface{})
   hybrid_type := data_map["hybrid"].(bool)
   if hybrid_type == true {
       fmt.Println("hybrid",hybrid_type)
       ws_type   :=  data_map["calc_id"].(string)
       calc_data := data_map["data"].(interface{})
       eto_calc_compute[ws_type](calc_data)
   }
   
}
