package eto_support

import (
	// "encoding/json"
   
	//"fmt"
	
	
	"lacima.com/redis_support/redis_handlers"

    "github.com/vmihailenco/msgpack/v5"
)





type ETO_RAIN_TYPE struct {
    Key         string  
    Status      bool    
    Date_string string 
    Priority    float64 
    Value       float64 
} 
    



var eto_redis_store redis_handlers.Redis_Hash_Struct


var rain_redis_store redis_handlers.Redis_Hash_Struct


var stream_store  redis_handlers.Redis_Hash_Struct

func Init_Record_Store( eto_hash, rain_hash,eto_stream_hash redis_handlers.Redis_Hash_Struct ){
    
    eto_redis_store  = eto_hash
    rain_redis_store = rain_hash
    stream_store     = eto_stream_hash
    
}


func ETO_RAIN_STREAM_CLEAR(){
    eto_redis_store.Delete_All()
    rain_redis_store.Delete_All()
    stream_store.Delete_All()
}


func ETO_Exist(key string)bool{
    
   return eto_redis_store.HExists(key)  
    
}

func Rain_Exist(key string )bool{
    
   return rain_redis_store.HExists(key)  
    
}

func Stream_Exist(key string )bool{

    
   return stream_store.HExists(key)  
    
}

func ETO_HKeys()[]string{
    
   return eto_redis_store.HKeys() 
    
}

func Rain_HKeys()[]string{
    
   return rain_redis_store.HKeys()  
    
}

func Stream_HKeys()[]string{

    
   return stream_store.HKeys() 
    
}
  
    

func ETO_HSet( key string,input ETO_RAIN_TYPE  ){
    

    temp := Eto_rain_to_map(input)
    b, err := msgpack.Marshal(&temp)
    if err != nil {
        panic(err)
    }
    

    eto_redis_store.HSet(key,string(b))
   
    
}
    
func Rain_HSet(key string, input ETO_RAIN_TYPE  ){
       

    temp := Eto_rain_to_map(input)
    b, err := msgpack.Marshal(&temp)
    if err != nil {
        panic(err)
    }
  

    rain_redis_store.HSet(key,string(b))
    
}

func Stream_HSet( key string, input []ETO_INPUT ){
    
    item := Eto_input_to_map(input)
    
    
    b, err := msgpack.Marshal(&item)
    if err != nil {
        panic(err)
    }
   
    
    stream_store.HSet(key,string(b))
   
}


func ETO_HGet( key string )ETO_RAIN_TYPE {
    
   item := make(map[string]interface{})
    
    status := eto_redis_store.HExists(key)
    if status == false {
        panic("non existant key")
    }
    
    input:= eto_redis_store.HGet(key)
    
    err := msgpack.Unmarshal([]byte(input), &item)
    if err != nil {
        panic("bad msgpack")
    }
       
    
    
    return Map_to_eto_rain(item)

}
    
func Rain_HGet( key string )ETO_RAIN_TYPE{
    
    item := make(map[string]interface{})
    
    status := rain_redis_store.HExists(key)
    if status == false {
        panic("non existant key")
    }
    
    input:= rain_redis_store.HGet(key)
    
    err := msgpack.Unmarshal([]byte(input), &item)
    if err != nil {
        panic("bad msgpack")
    }
    
    return Map_to_eto_rain(item)
}


func Eto_rain_to_map( input ETO_RAIN_TYPE )map[string]interface{}{
    
    return_value := make(map[string]interface{})
    
    return_value["key"]         = input.Key
    return_value["status"]      = input.Status
    return_value["date_string"] = input.Date_string
    return_value["priority"]    = input.Priority
    return_value["value"   ]    = input.Value

    return return_value
}

func Map_to_eto_rain( input map[string]interface{})ETO_RAIN_TYPE{
    var return_value ETO_RAIN_TYPE
    return_value.Key            = input ["key"].(string)
    return_value.Status         = input["status"].(bool)
    return_value.Date_string    = input["date_string"].(string)
    return_value.Priority       = input["priority"].(float64)
    return_value.Value          = input["value"].(float64)
    
    return return_value
}

func Stream_HGet( key string )map[string][]float64{
    var item map[string][]float64
    
    status := stream_store.HExists(key)
    if status == false {
        panic("non existant key")
    }
    
    input:= stream_store.HGet(key)
    
    err := msgpack.Unmarshal([]byte(input), &item)
    if err != nil {
        panic("bad msgpack")
    }
    
    return item

}
 

func Eto_input_to_map(input []ETO_INPUT ) map[string][]float64{
   return_value := make(map[string][]float64)   
   return_value["wind_speed"]                    = process_eto_input_map("wind_speed",input)
   return_value["temp_C"]                        = process_eto_input_map("temp_C",input)
   return_value["humidity"]                      = process_eto_input_map("humidity",input)
   return_value["SolarRadiationWatts_m_squared"] = process_eto_input_map("SolarRadiationWatts_m_squared",input)
   return_value["delta_timestamp"]               = process_eto_input_map("delta_timestamp",input)
   return return_value
}

func process_eto_input_map( field string,input_list []ETO_INPUT )[]float64{
    length       := len(input_list)
    return_value := make([]float64,length)
    for index,input := range input_list {
        switch(field){
            case "wind_speed":
                    return_value[index] = input.Wind_speed        
            
            case "temp_C":    
                     return_value[index] = input.Temp_C     
                
            case "humidity":
                     return_value[index] = input.Humidity      
                
            case "SolarRadiationWatts_m_squared":
                     return_value[index] = input.SolarRadiationWatts_m_squared      
	       
            case "delta_timestamp":
                     return_value[index] = input.Delta_timestamp     
            default:
                panic("bad field")
        }
    }
    return return_value
}
    
func Map_array_to_eto_array(length int ,input_map map[string][]float64)[]ETO_INPUT{
   return_value := make([]ETO_INPUT,length)
   
   for i:=0;i<length;i++ {
       var item ETO_INPUT
       
	   item.Wind_speed                    = input_map["wind_speed"][i]
	   item.Temp_C                        = input_map["temp_C"][i]
	   item.Humidity                      = input_map["humidity"][i]
	   item.SolarRadiationWatts_m_squared = input_map["SolarRadiationWatts_m_squared"][i]
	   item.Delta_timestamp               = input_map["delta_timestamp"][i]
	   return_value[i] = item
    }
    
    return return_value

}
    
    
    
    
