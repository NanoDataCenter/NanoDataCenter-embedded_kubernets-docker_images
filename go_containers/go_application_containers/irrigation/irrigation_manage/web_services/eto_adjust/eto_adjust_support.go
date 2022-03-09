package eto_adjust

import (
    //"fmt"
    //"strings"
    //"net/http"
    //"html/template"
    "encoding/json"
    //"lacima.com/redis_support/generate_handlers"
	//"lacima.com/redis_support/graph_query"
   // "lacima.com/redis_support/redis_handlers"
  
    //"lacima.com/Patterns/web_server_support/jquery_react_support"
    "lacima.com/Patterns/msgpack_2"
    //"github.com/vmihailenco/msgpack/v5"
)


func generate_json_data(){
    
    eto_adjust_raw = make(map[string]map[string]float64)
    keys := Eto_Accumulation.HKeys()
    for _,key := range keys{
         temp := make(map[string]float64)
         accumulation_value,_ := msg_pack_utils.Unpack_float64(Eto_Accumulation.HGet(key))
         
         temp["eto"] = accumulation_value
         
         reserve_value,_:= msg_pack_utils.Unpack_float64(Eto_Reserve.HGet(key))
         
         temp["reserve"] = reserve_value
         eto_adjust_raw[key] = temp
    }   
    temp,_ := json.Marshal(eto_adjust_raw)
    eto_adjust_json = string(temp)
}



func Process_new_eto_adjust(raw_input string){
    var decode_value map[string]map[string]float64
  
    if err := json.Unmarshal([]byte(raw_input), &decode_value); err != nil {
        panic("bad json")
    }else{
        store_eto_file(raw_input)
        eto_adjust_json = raw_input
    }
    
}

func store_eto_file(data string){
    var decode_value map[string]map[string]float64
    
    err1 := json.Unmarshal([]byte(data), &decode_value)
    if err1 != nil {
        panic("should not happen")
    }
    for key,value := range decode_value{
        eto_pack := msg_pack_utils.Pack_float64(value["eto"])
        reserve_pack := msg_pack_utils.Pack_float64(value["reserve"])
        Eto_Accumulation.HSet(key,eto_pack)
        Eto_Reserve.HSet(key,reserve_pack)
    }
}

