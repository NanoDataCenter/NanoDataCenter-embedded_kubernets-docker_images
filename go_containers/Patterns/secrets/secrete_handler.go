package secrets

import "lacima.com/redis_support/redis_file"
import "lacima.com/Patterns/msgpack"

var driver *redis_file.Redis_File_Struct

const secret_db int = 5

func Init_file_handler(site_data map[string]interface{} ){

	address  :=  site_data["host"].(string)
    port  := 	int(site_data["port"].(float64))
    redis_file.Create_redis_data_handle(address,port,secret_db)
	driver = redis_file.Construct_File_Struct(  ) 


}

func Get_Secret( file_name string) map[string]map[string]string {

    return_value := make(map[string]map[string]string)
    path := "/PASSWORDS/"+file_name

    msgpack_data := (*driver).Get(path)
    unpacked_data := msgpack_utils.Unpack(msgpack_data)
	
	for up_key,up_value := range unpacked_data.(map[interface{}]interface{}) {
	   key := up_key.(string)
	   return_value[key] = make(map[string]string)
	   for up_key1,up_value1 := range up_value.(map[interface{}]interface{}){
	      key1 := up_key1.(string)
		  value1 := string(up_value1.([]uint8))
		  return_value[key][key1] = value1
	   }
	 }
	 return return_value
	
}