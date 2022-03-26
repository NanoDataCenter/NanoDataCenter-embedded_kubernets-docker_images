package main

import (
          "fmt"
          "strconv"
          "time"
          "strings"
         "lacima.com/redis_support/generate_handlers"
	     "lacima.com/redis_support/graph_query"
	     "lacima.com/redis_support/redis_handlers"
	     "lacima.com/site_data"   
         // "encoding/json"
          "lacima.com/go_application_containers/irrigation/irrigation_libraries/postgres_access/schedule_access"
)


type Action_data_type struct  {
    key                        string
    main_controller   string
    sub_controller     string
    master_flag         bool
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
        check_irrigation_jobs()
        fmt.Println("main loop pooling loop")
		time.Sleep(time.Second * 60)
	}

}

func setup_data_structures(){
    
      control_block = irr_sched_access.Construct_irr_schedule_access()
     // setup redis job table
    // setup rpc job table
     // clean redis job table    

    
}

func clean_redis_job_table(){
    // base_time
    // get all the keys
    // for all redis keys 
    //   convert string to float64
    //   if base_time > redis_time
    //    delete redis key
    
    
}


func check_irrigation_jobs(){
     
     data,err :=    irr_sched_access.Action_Select_All()
     
     if err != true {
         panic("data fetch error")
     }
     fmt.Println("length",len(data)  )
     for  _, item   := range data {
         parse_input(item)
        if irr_sched_access.Check_schedule_job(  action_data.key ) == true {
            fmt.Println("job previous scheduled")
            continue
        }
        if check_irrigation_job() == true {
           irr_sched_access.Set_schedule_job( action_data.key)
            queue_irrigation_jobs()   
        }else{
           irr_sched_access. Clear_schedule_job( action_data.key)
            fmt.Println("job not queued")
        }
        
     }
}        


func parse_input( item map[string]interface{}){
      form_key( item )
      form_hr( item )
      form_dow(item)
      fmt.Println(action_data)
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
      // fmt.Println("made it to check hour 2")
       return true
   }
   if current_min > start_min {
      // fmt.Println("made it to check hour 3")
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
    time_mod :=  doy %  int64(action_data.doy_divisor)
    if time_mod == int64(action_data.doy_modulus) {
        //fmt.Println("check day 2",time_mod ,time_mod)
        return true
    }
    //fmt.Println("check day 3")
    return false
        
    
}


func queue_irrigation_jobs(  ) {
    
     fmt.Println("queue this job",action_data)
}



func  form_key(item map[string]interface{}){
    temp                                := "true"
    action_data.main_controller  = item["main_controller"].(string)
    action_data.sub_controller    = item["sub_controller"].(string)
    action_data.master_flag        =  item["master_flag"].(bool)
    action_data.name                   =  item["name"].(string)
    action_data.description          =  item["description"].(string)
    if  action_data.master_flag == false{
          temp = "false"
    }
    temp_list := make([]string,4)
    temp_list[0] = temp
    temp_list[1] = action_data.main_controller
    temp_list[2] = action_data.sub_controller
    temp_list[3]  = action_data.name
    action_data.key    = strings.Join(temp_list,"/")
  
    
    
}

func form_dow(item map[string]interface{}){

    action_data.dow_week_flag  = item["dow_week_flag"].(bool)
    day_mask := make([]bool,7)
    
    temp_mask :=  item["day_mask"].([]interface{})
    for  index,value    := range temp_mask{
         day_mask[index]   = value.(bool)
    }
   action_data.day_mask = day_mask
   
   temp_string :=  item["doy_divisor"].(string)
   temp_float, _ := strconv.ParseFloat(temp_string, 64)
   action_data.doy_divisor = temp_float
   
   temp_string =  item["doy_modulus"].(string)
   temp_float, _  = strconv.ParseFloat(temp_string, 64)
   action_data.doy_modulus = temp_float
    
}

func form_hr(item map[string]interface{}){
    
   action_data.start_time_hr          =    item["start_time_hr"].(float64)  
   action_data.start_time_min       =    item["start_time_min"].(float64)
   action_data.end_time_hr            =    item["end_time_hr"].(float64)
   action_data.end_time_min         =   item["end_time_min"].(float64)

}


/*
item map[day_mask:[true true true true false true true] description:action_data 1 description dow_week_flag:true doy_divisor:2 doy_modulus:0 end_time_hr:21 end_time_min:0 main_controller:main_server master_flag:false name:action start_time_hr:10 start_time_min:0 steps:[map[description: name:VALVE_RESISTANCE:valve_resistance_1 type:action] map[description:schedule 1 description name:schedule_1 type:schedule] map[description: name:VALVE_RESISTANCE:valve_resistance_1 type:action] map[description:schedule 2 description name:test_schedule_2 type:schedule]] sub_controller:main_server/sub_server_1]

item map[day_mask:[true true true true false true true] description:test 2 action dow_week_flag:true doy_divisor:2 doy_modulus:0 end_time_hr:21 end_time_min:0 main_controller:main_server master_flag:false name:test2 start_time_hr:5 start_time_min:0 steps:[map[description:schedule 1 description name:schedule_1 type:schedule]] sub_controller:main_server/sub_server_1]


item map[day_mask:[false false false false false false false] description:test 3  example dow_week_flag:false doy_divisor:2 doy_modulus:0 end_time_hr:23 end_time_min:0 main_controller:main_server master_flag:true name:test3 start_time_hr:10 start_time_min:0 steps:[] sub_controller:main_server/sub_server_1]

*/
/*
 key start_time_hr float64  
value 10
key end_time_hr float64  
value 21
key day_mask []interface {}  
value [true true true true false true true]
key end_time_min float64  
value 0
key master_flag bool  
value false
key start_time_min float64  
value 0
key dow_week_flag bool  
value true
key doy_divisor string  
value 2
key sub_controller string  
value main_server/sub_server_1
key name string  
value action
key doy_modulus string  
value 0
key sub_controller string  
value main_server/sub_server_1
key day_mask []interface {}  
value [true true true true false true true]
key start_time_hr float64  
value 5
key start_time_min float64  
value 0
key end_time_min float64  
value 0
key main_controller string  
value main_server
key description string  
value test 2 action
key steps []interface {}  
value [map[description:schedule 1 description name:schedule_1 type:schedule]]
key dow_week_flag bool  
value true
key master_flag bool  
value false
key name string  
value test2
key doy_modulus string  
value 0
key end_time_hr float64  
value 21
key doy_divisor string  
value 2
key description string  
value test 3  example
key start_time_min float64  
value 0
key dow_week_flag bool  
value false
key master_flag bool  
value true
key name string  
value test3
key steps []interface {}  
value []
key start_time_hr float64  
value 10
key end_time_hr float64  
value 23
key doy_modulus string  
value 0
key main_controller string  
value main_server
key sub_controller string  
value main_server/sub_server_1
key end_time_min float64  
value 0
key day_mask []interface {}  
value [false false false false false false false]
key doy_divisor string  
value 2

*/
