package eto_support

import (
   "fmt"
   //"strconv"
   "strings"
   "lacima.com/server_libraries/file_server_library"   
   "encoding/json"
    "lacima.com/redis_support/redis_handlers"
    "lacima.com/Patterns/msgpack_2"
)

var file_server_library file_server_lib.File_Server_Client_Type

func Json_float_unmarshall( input interface{} ) string{
    
    var return_value string
     switch input.(type){
         
         case string:
             return_value = input.(string)
         case float64:
             temp := input.(float64)
             return_value =  fmt.Sprintf("%3.0f", temp)
          
         default:
             panic("bad type")
         
     }
     return_value = strings.Trim(return_value," \t")
     return return_value
    
}


func Setup_ETO_Accumulation_Table(eto_data_handler redis_handlers.Redis_Hash_Struct ){
    
    search_list := []string{"RPC_SERVER:SITE_FILE_SERVER","RPC_SERVER"}
    file_server_library = file_server_lib.File_Server_Init(&search_list)
    fmt.Println(file_server_library.Ping())
    data, err := file_server_library.Read_file( "/app_data_files/eto_site_setup.json")
    
    if err != true{
        panic("bad file read")
    }
        
    byt := []byte(data)
    var file_data []map[string]interface{}


    if err := json.Unmarshal(byt, &file_data); err != nil {
        panic(err)
    }
    
    eto_redis_hash_data := (eto_data_handler).HGetAll()

    new_data := make(map[string]float64)
   
    for _, value := range file_data {
       
        controller := value["controller"].(string)
        pin        := Json_float_unmarshall(value["pin"])
        new_data[ controller+"|"+pin ] =0
    }
    
    (eto_data_handler).Delete_All()
    //
    // merge old values and possible new values into new table.
    //
    msg_pack_zero := msg_pack_utils.Pack_float64(float64(0))
    for key,_ := range new_data {
         if value,ok := eto_redis_hash_data[key] ; ok == true {
             
              (eto_data_handler).HSet(key,value)
         }else{
            
            (eto_data_handler).HSet(key,msg_pack_zero)
         }
    }
        
    fmt.Println((eto_data_handler).HGetAll())
    
           
     
}

