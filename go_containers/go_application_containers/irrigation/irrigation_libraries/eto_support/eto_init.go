package eto_support

import (
	"fmt"
	//"strconv"
	"encoding/json"
	"lacima.com/Patterns/msgpack_2"
	"lacima.com/redis_support/redis_handlers"
	"lacima.com/server_libraries/file_server_library"
	"strings"
)

var file_server_library file_server_lib.File_Server_Client_Type

var Eto_Accumulation        redis_handlers.Redis_Hash_Struct
var Eto_Reserve             redis_handlers.Redis_Hash_Struct
var Eto_Valve_Definition    redis_handlers.Redis_Hash_Struct


func Setup_eto_handlers() {

	search_list := []string{"ETO_SETUP_PROPERTIES:ETO_SETUP_PROPERTIES","ETO_DATA_STRUCTURES"}
	eto_data_structs = data_handler.Construct_Data_Structures(&search_list)
	
    Eto_Accumulation      = (*eto_data_structs)["ETO_ACCUMULATION"].(redis_handlers.Redis_Hash_Struct)
	Eto_Reserve           = (*eto_data_structs)["ETO_RESERVE"].(redis_handlers.Redis_Hash_Struct)
	Eto_Valve_Definition  = (*eto_data_structs)["ETO_VALVE_DEFINITIONS"].(redis_handlers.Redis_Hash_Struct)
	
}





func Update_Accumulation_Tables(eto_accumulation redis_handlers.Redis_Hash_Struct,  eto float64){
    
   data :=  eto_accumulation.HGetAll()
   
   for key, msgpack_value := range data {
      previous_value,_ := msg_pack_utils.Unpack_float64(msgpack_value)  // error will result in 0.0 returned
      new_value        := previous_value +eto 
      eto_accumulation.HSet(key,msg_pack_utils.Pack_float64(new_value))
   }
}       


func GetAll_Accumulation_Tables(eto_accumulation redis_handlers.Redis_Hash_Struct)map[string]float64{
    return_value := make(map[string]float64)
    data :=  eto_accumulation.HGetAll()
   
   for key, msgpack_value := range data {
      float_value,_ := msg_pack_utils.Unpack_float64(msgpack_value)  // error will result in 0.0 returned
      return_value[key] = float_value
   }
   return return_value
}       
    
    
    

    
   
   
    
    
    
