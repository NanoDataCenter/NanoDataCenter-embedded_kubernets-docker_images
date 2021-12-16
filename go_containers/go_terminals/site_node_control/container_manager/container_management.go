package container_management

import (
gc "github.com/gbin/goncurses"
"lacima.com/go_terminals/go_terminal_library"
"lacima.com/redis_support/generate_handlers"
"lacima.com/redis_support/redis_handlers"
//"lacima.com/redis_support/graph_query"
"lacima.com/Patterns/msgpack_2"
	
)


var current_line *gc_support.LINE_BUFFER
var hash_status redis_handlers.Redis_Hash_Struct  

func Container_management_launcher(){
    find_hash_status()
    gc_support.Construct_Console(gc_support.Commmand_handler,connection_key_handler,init_handler)

}



func init_handler(){
     current_line = gc_support.Return_SubWindow_Draw_Structures()
     generate_main_screen_message(current_line.Window)
}

func find_hash_status(){
    display_struct_search_list := &[]string{"DOCKER_CONTROL"}
    data_structures :=  data_handler.Construct_Data_Structures(display_struct_search_list)
    hash_status = (*data_structures)["DOCKER_DISPLAY_DICTIONARY"].(redis_handlers.Redis_Hash_Struct)
}   


func generate_main_screen_message(window *gc.Window){
    
   
    message        := "F1: Display Status F2: Disable Container Management F3: Enable Container Management"
    title          := "Manage Containers"

    gc_support.Construct_Console_Menu(current_line.Window,title,message)
}

func connection_key_handler(ch gc.Key )bool{
     return_value := false
     switch ch {
        
        case gc.KEY_F1:
            display_status()
            
        case gc.KEY_F2:
            disable_container_management()
    
        case gc.KEY_F3:
            enable_container_management()
       
        default:
            ;    
     }
    return return_value
}

const marker string="****************************************************************************\n"
func display_status(){
    keys := hash_status.HKeys()
    current_line.Return_cmd(marker)
    for _,container := range keys{
      status :=  hget_status_value(container)
      status_line := "Container: "+container+" Active: "+bool_display(status["active"])+ " Managed: "+bool_display(status["managed"])+"\n"
      current_line.Return_cmd(status_line)
    }
    current_line.Return_cmd(marker)
    
}

func disable_container_management(){
   keys := hash_status.HKeys()
   set   := make(map[string]bool)
   set["active"] = true
   set["managed"] = false
   current_line.Return_cmd(marker)
    for _,container := range keys{
      hset_status_values(container,set)
      status :=  hget_status_value(container)
      status_line := "Container: "+container+" Active: "+bool_display(status["active"])+ " Managed: "+bool_display(status["managed"])+"\n"
      current_line.Return_cmd(status_line)
    }   
    current_line.Return_cmd(marker)
}
func enable_container_management(){
   keys := hash_status.HKeys()
   set   := make(map[string]bool)
   set["active"] = true
   set["managed"] =true
   current_line.Return_cmd(marker)
    for _,container := range keys{
      hset_status_values(container,set)
      status :=  hget_status_value(container)
      status_line := "Container: "+container+" Active: "+bool_display(status["active"])+ " Managed: "+bool_display(status["managed"])+"\n"
      current_line.Return_cmd(status_line)
    }   
    current_line.Return_cmd(marker)
}


func bool_display( value bool)string{
    if value == true {
        return "true"
    }
    return "false"
}







func hget_status_value( field string ) map[string]bool{
   
    v,err := msg_pack_utils.Unpack_map_string_bool(hash_status.HGet(field) )  
    if err != true {
        panic("error")
    }
    
	
	return v

}




func hset_status_values( field string, value map[string]bool){

   
    hash_status.HSet(field, msg_pack_utils.Pack_map_string_bool(value))

}    
