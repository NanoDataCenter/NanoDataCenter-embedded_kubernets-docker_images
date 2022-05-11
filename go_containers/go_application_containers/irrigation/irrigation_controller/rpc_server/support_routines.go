package irrigation_rpc


import (
   "fmt"
    "strings"
     "lacima.com/go_application_containers/irrigation/irrigation_controller/irrigation_controller_library"
)

func verify_controller( key string )([]string,bool){
 
    key_list := strings.Split(key,"~")
    fmt.Println("key list",key_list)
    if len(key_list) < 2 {
        return key_list, false
    }
    if key_list[0] == "true" {
        return key_list,verify_master_controller(key_list)
    }
    if key_list[0] == "false"{
        if len(key_list)< 3 {
            return key_list , false
        }
        return key_list,verify_sub_controller(key_list )
    }
    return make([]string,0),false
    
}

func verify_master_controller(key_list []string)bool{
  temp_dict := irrigation_controller_library.Registry_data.Master_table_list 
  if _,ok := temp_dict[key_list[1]]; ok==false {
      return false
  }
  return true
}

func verify_sub_controller( key_list []string)bool{
    if verify_master_controller( key_list ) == false {
        return false
    }
    temp_dict := irrigation_controller_library.Registry_data.Master_table_list 
    temp := temp_dict[key_list[1]] 
  if _,ok := temp[key_list[2]]; ok==false {
      return false
  }
  return true
}

func verify_action( action string )bool{
  
   if _,ok := irrigation_controller_library.Registry_data.Actions[action];ok==false{
        return false
    }
    return true
}


func verify_action_to_controller(key_list[]string ,action string )bool{
    key := ""
    if key_list[0] == "true"{
         key =    irrigation_controller_library.Construct_server_key(true, key_list[1], "" )
    }else{
        key =  irrigation_controller_library.Construct_server_key(false, key_list[1], key_list[2] )
    }
    fmt.Println("key",key,action)
    if _,ok := irrigation_controller_library.Registry_data.Action_list[key][action];ok==false{
        return false
    }
    return true
}

func fetch_action(action string )map[string]interface{} {

  return  irrigation_controller_library.Registry_data.Actions[action]
}

func queue_action_immediate(key string, action_data map[string]interface{})bool{
    fmt.Println("immediate",key,action_data)
    return true
}
         
func queue_action_delayed(key string , action_data map[string]interface{})bool{
     fmt.Println("queued",key,action_data)
    return true  
} 

func queue_direct_io(station string ,io int64,  action bool, master_controller string ,sub_controller string)bool{
    fmt.Println("queue_direct_io",station,io,action,master_controller,sub_controller)
    return true
}
