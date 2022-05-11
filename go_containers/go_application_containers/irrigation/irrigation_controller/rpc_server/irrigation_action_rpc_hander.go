package irrigation_rpc


import (
  "fmt"
)

/*
  Master_table_list                 map[string][]string
  Valve_list                               map[string]map[string]map[string][]int
  Inverse_Valve_Map             map[string]string
  Actions                                 map[string]map[string]interface{}
  Action_list                           map[string][]string
}
parameters["COMMAND"]                = "QUEUE_ACTION"
       parameters["KEY"]                             = key 
       parameters["NAME"]                         = action
*/
func handler_actions(parameters map[string]interface{})map[string]interface{}{
    
  
    key := parameters["KEY"].(string)
    action := parameters["NAME"].(string)
    parameters["STATUS"] = false  
    key_list, status := verify_controller(key)
    if  status == false{
        return parameters
    }
    
    if verify_action( action ) == false {
        return parameters
    }
   
    if verify_action_to_controller(key_list ,action ) == false{
        return parameters
    }
    
    action_data := fetch_action(action)
    
    if action_data["immediate"] == true {
         parameters["STATUS"] =queue_action_immediate(key,action_data)
         
    }else{
         parameters["STATUS"] =queue_action_delayed(key,action_data)
    }
   fmt.Println("parameters",parameters)
    return parameters
}


