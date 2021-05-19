package secrets

import "strings"
import "lacima.com/redis_support/redis_file"


var driver *redis_file.Redis_File_Struct

const secret_db int = 5

func Init_file_handler(site_data map[string]interface{} ){

	address  :=  site_data["host"].(string)
    port  := 	int(site_data["port"].(float64))
    redis_file.Create_redis_data_handle(address,port,secret_db)
	driver = redis_file.Construct_File_Struct(  ) 


}

func Get_Secret( name,field string) string {
     
     var return_value string
     var ok           bool
     
     if return_value,ok = driver.HGet(name,field);ok == false {
	     panic("non existant password")
     }
     return return_value
}

func Extract_User_Password( input string )(string,string){
    
    lines := strings.Split(input,":")
    if len (lines) < 2 {
      panic("bad format")
    }
    return lines[0],lines[1]   

    
    
}
