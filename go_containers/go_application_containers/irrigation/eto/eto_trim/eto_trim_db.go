 package eto_trim
 
//import "fmt"
import "time"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/server_libraries/postgres"
 
 
 
 
var eto_history           pg_drv.Postgres_Stream_Driver
var rain_history          pg_drv.Postgres_Stream_Driver
                                                        
 
var  trim_time int64 

 
func Trim_int(trim_time_seconds int64) { // one day trim time 
    
   trim_time                     = trim_time_seconds              
   search_list := []string{"WEATHER_DATA"}
   Eto_data_structs := data_handler.Construct_Data_Structures(&search_list)
   eto_history      = (*Eto_data_structs)["ETO_HISTORY"].(pg_drv.Postgres_Stream_Driver)
   rain_history     = (*Eto_data_structs)["RAIN_HISTORY"].(pg_drv.Postgres_Stream_Driver)
    
}




func Trim_dbs(){
    
    
   for true {
       
      eto_history.Trim(trim_time)
      eto_history.Vacuum()
      eto_history.Trim(trim_time)
      eto_history.Vacuum()
      //fmt.Println("made it here")
      time.Sleep(time.Second *3600)

       
   }
    
}
