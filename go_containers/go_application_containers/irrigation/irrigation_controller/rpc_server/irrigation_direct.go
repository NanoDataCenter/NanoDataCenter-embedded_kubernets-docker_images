package irrigation_rpc

import (
    "fmt"
  //  "strings"
      "strconv"
     "lacima.com/go_application_containers/irrigation/irrigation_controller/irrigation_controller_library"
)

 

func handler_irrigation_direct(parameters map[string]interface{})map[string]interface{}{
    
    defer print_data(parameters)
     station := parameters["STATION"].(string)
     io          := int64(parameters["IO"].(float64))
     action := parameters["ACTION"].(bool)
      parameters["STATUS"] = false                            
     
      if verify_valve(station,io) == false{
          return parameters
      }
      master_controller,sub_controller := find_controllers(station,io)
      parameters["STATUS"] = queue_direct_io(station,io,action, master_controller,sub_controller)

    parameters["STATUS"] = true
    
    return parameters
}

func print_data(parameters map[string]interface{}){
 
    fmt.Println(parameters)
}




func verify_valve( station string , io int64)bool{
     key := station+":"+strconv.Itoa(int(io))
    
     if _,ok := irrigation_controller_library.Registry_data.Inverse_Valve_Map[key];ok==false{
         return false
     }
    return true
    
}

func find_controllers(station string, io int64)(string,string){
    return "AAA","BBB"
}
