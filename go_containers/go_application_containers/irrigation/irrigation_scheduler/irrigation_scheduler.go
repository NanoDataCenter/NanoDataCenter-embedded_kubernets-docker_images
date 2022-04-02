package main

import (
          "fmt"
         // "strconv"
          "time"
          //"strings"
         "lacima.com/redis_support/generate_handlers"
	     "lacima.com/redis_support/graph_query"
	     "lacima.com/redis_support/redis_handlers"
	     "lacima.com/site_data"   
         "encoding/json"
          "lacima.com/go_application_containers/irrigation/irrigation_libraries/postgres_access/schedule_access"
)


type Action_data_type struct  {
    key                        string
    name                    string
    description          string
    day_mask             []bool
    dow_week_flag     bool
    doy_divisor           float64
    doy_modulus         float64
    start_time_hr       float64
    start_time_min    float64
    end_time_hr         float64
    end_time_min      float64
    steps                     []map[string]interface{}
}
    
var  action_data Action_data_type    
    
var control_block irr_sched_access.Irr_sched_access_type

func main(){
    
	var config_file = "/data/redis_configuration.json"
	var site_data map[string]interface{}
      
	site_data = get_site_data.Get_site_data(config_file)
	redis_handlers.Init_Redis_Mutex()
	graph_query.Graph_support_init(&site_data)
	data_handler.Data_handler_init(&site_data)
    setup_data_structures()
    for true {
        //fmt.Println("\nCheck Cycle \n")
        check_irrigation_jobs()
        //fmt.Println("main loop pooling loop")
		time.Sleep(time.Second * 60)
	}

}

func setup_data_structures(){
    
      control_block = irr_sched_access.Construct_irr_schedule_access()
     irr_sched_access.Delete_schedule_job()
       

    
}



func check_irrigation_jobs(){
     
     data,err :=    irr_sched_access.Action_Select_All_Raw()
     
     if err != true {
         panic("data fetch error")
     }
     
     for  index, item   := range data {
        
         parse_input(item)
         //fmt.Println("action_data",index,action_data)
          job_key := action_data.key +":"+action_data.name
        
        if check_irrigation_job() == true {
          if irr_sched_access.Check_schedule_job( job_key) == true {
            //fmt.Println("job previous scheduled")
            continue
        }
           irr_sched_access.Set_schedule_job( job_key )
            queue_irrigation_jobs(  action_data.key,  data[index] )   
        }else{
           //fmt.Println("$$$$$$$$$$$$$$$$ job not queued and entry  cleared",job_key)
           irr_sched_access. Clear_schedule_job( job_key)
         
        }
        
     }
}        


func parse_input( item map[string]interface{}){
      form_key( item )
      form_hr( item )
      form_dow(item)
     
}

func check_irrigation_job( )bool{
     // check redis data base
    if check_hour() == false{
       return false
    }
    
    if check_day() == false{
        return false
    }
    return true   
}

func check_hour()bool {
   currentTime := time.Now()      
   h := currentTime.Hour()
   m := currentTime.Minute()
   current_min := float64(h*60+m)
   
   start_min := (action_data.start_time_hr*60) + action_data.start_time_min
   end_min  := (action_data.end_time_hr*60) + action_data.end_time_min
   
   //fmt.Println("hour data",current_min,start_min,end_min)
   if start_min < end_min {
        if current_min < start_min{
            return false
        }
        if current_min > end_min {
            return false
        }
        //fmt.Println("made it to check_hour 1")
       return true
       
   }
   if  current_min < end_min {
       //fmt.Println("made it to check hour 2")
       return true
   }
   if current_min > start_min {
       //fmt.Println("made it to check hour 3")
       return true
   }
   //fmt.Println("made it to chech hour 4")
   return false
}



func check_day()bool{
    currentTime :=  time.Now()      
    dow               :=  int64(currentTime.Weekday())
    doy                :=  int64(currentTime.YearDay()) 
    if action_data.dow_week_flag == true {
        //fmt.Println("check day 1",action_data.day_mask[dow])
        return action_data.day_mask[dow]
    }
    
    time_mod := int64( doy) %  int64(action_data.doy_divisor)
    //fmt.Println("tod",time_mod,int64(action_data.doy_modulus))    
    if time_mod == int64(action_data.doy_modulus) {
        //fmt.Println("check day 2",time_mod ,time_mod)
        return true
    }
    //fmt.Println("check day 3")
    return false
        
    
}




func  form_key(item map[string]interface{}){
   //fmt.Println("item",item)
    action_data.key                       = item["server_key"].(string)
    action_data.name                  = item["name"].(string)
}

func form_dow(item map[string]interface{}){

    action_data.dow_week_flag  = item["dow_week_flag"].(bool)
    day_mask := make([]bool,7)
    
    temp_mask :=  item["day_mask"].([]interface{})
    for  index,value    := range temp_mask{
         day_mask[index]   = value.(bool)
    }
   action_data.day_mask = day_mask
   
   //temp_string :=  item["doy_divisor"].(string)
   //temp_float, _ := strconv.ParseFloat(temp_string, 64)
   //action_data.doy_divisor = item["doy_divisor"].(float64)
   action_data.doy_divisor = item["doy_divisor"].(float64)
   
   //temp_string =  item["doy_modulus"].(string) 
   //temp_float, _  = strconv.ParseFloat(temp_string, 64)
   action_data.doy_modulus = item["doy_modulus"].(float64) 
    
}

func form_hr(item map[string]interface{}){
    
   action_data.start_time_hr          =    item["start_time_hr"].(float64)  
   action_data.start_time_min       =    item["start_time_min"].(float64)
   action_data.end_time_hr            =    item["end_time_hr"].(float64)
   action_data.end_time_min         =   item["end_time_min"].(float64)

}


func queue_irrigation_jobs(   key string  ,   json_data map[string]interface{} ) {
    data :=  json_data["steps"].([]interface{})
    for  _, temp := range data {
        array_element                 :=   temp.(map[string]interface{})
        name                                := array_element["name"].(string)
        action_type                     := array_element["type"].(string)
        if action_type == "schedule" {
               handle_schedule(key,name)
        }else{
          
            fmt.Println("action",name, irr_sched_access.Queue_Action( key,  name ) )
        }
    }
}


func  handle_schedule(server_key, schedule_name string ){
     schedule_data,_   :=irr_sched_access.Select_schedule_name(server_key,schedule_name )
     for _, temp := range schedule_data{
               steps := generate_step_data(temp.Json_data)
               for _,step := range steps {
                   time           := step["time"].(float64)
                   temp         :=  step["station"].(map[string]interface {})
                    station_io := generate_station_io( temp )
                    //fmt.Println("station",time,station_io)
                   fmt.Println(irr_sched_access.Queue_Managed_Irrigation( server_key, time ,  station_io ))
               }
               
           }
           
       
}

func generate_step_data(input string)[]map[string]interface{}{
    var data []map[string]interface{}
        if err := json.Unmarshal([]byte(input), &data); err != nil {
           panic(err)
        }
        return data
}

//station:map[station_3:map[1:1] station_4:map[1:1]] time:60]
func generate_station_io(  input map[string]interface {})[]string{
  return_value := make([]string,0)
   for station, io_data :=  range input {
       for pin , _  := range io_data.(map[string]interface{}) {
          return_value= append(return_value, station+":"+pin)
       }
   }
   return return_value
}
