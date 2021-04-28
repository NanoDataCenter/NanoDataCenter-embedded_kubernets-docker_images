package irrigation_rpc


import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"

type Irrigation_Client_Type struct{

   driver redis_handlers.Redis_RPC_Struct
}




func Irrigation_RPC_Client_Init(search_list *[]string)Irrigation_Client_Type{

  var return_value Irrigation_Client_Type
  var handlers = data_handler.Construct_Data_Structures(search_list)  
  return_value.driver = (*handlers)["IRRIGATION_JOB_SERVER"].(redis_handlers.Redis_RPC_Struct)
  return return_value
}  

func (v Irrigation_Client_Type)Ping()bool{
  

       var parameters = make(map[string]interface{})
       var result = v.driver.Send_rpc_message( "ping", &parameters ) 
       //fmt.Println("ping",result)
       //fmt.Println(	(*result)["status"].(bool))   
       return (*result)["status"].(bool)
}

func (v Irrigation_Client_Type)Queue_Schedule(schedule_name string)bool {
  

       var parameters = make(map[string]interface{})
       
       parameters["schedule_name"] = schedule_name
       var result = v.driver.Send_rpc_message( "queue_schedule", &parameters )  
        return (*result)["status"].(bool)
}

func (v Irrigation_Client_Type)Get_Rain_Flag()bool {
  

       var parameters = make(map[string]interface{})
       
    
       var result = v.driver.Send_rpc_message( "rain_flag", &parameters )  
        return (*result)["status"].(bool)
}

