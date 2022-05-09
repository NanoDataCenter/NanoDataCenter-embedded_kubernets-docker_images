package irrigation_rpc


import (
   // "fmt"
)


func verify_controller( key string )bool{
    
    return true
    
}



func verify_action( action string )bool{
    
    return true
}


func verify_action_to_controller(key,action string )bool{
    
    return true
}

func fetch_action(action string )map[string]interface{} {

  return make(map[string]interface{})
}

func queue_immediate(action_data map[string]interface{})bool{
    return true
}
         
func queue_delayed(action_data map[string]interface{})bool{
    return true  
} 
