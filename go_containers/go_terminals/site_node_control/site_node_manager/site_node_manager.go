package site_node_management



import (
gc "github.com/gbin/goncurses"
"lacima.com/go_terminals/go_terminal_library"
"lacima.com/server_libraries/node_control_rpc"
"lacima.com/server_libraries/site_control_rpc"
//"lacima.com/go_terminals/site_node_control/central_db"
//"lacima.com/go_terminals/docker_control"

	
)


var current_line *gc_support.LINE_BUFFER
var node_rpc_structures node_control_server_lib.Node_Server_Client_Type
var site_rpc_structure  site_control_server_lib.Site_Server_Client_Type

func Site_node_management_launcher(){
    node_rpc_structures = node_control_server_lib.Node_Server_Init()
    site_rpc_structure  = site_control_server_lib.Site_Server_Init()
    gc_support.Construct_Console(gc_support.Commmand_handler,site_key_handler,init_handler)

}



func init_handler(){
     current_line = gc_support.Return_SubWindow_Draw_Structures()
     generate_main_screen_message(current_line.Window)
}

func generate_main_screen_message(window *gc.Window){
    
   
    message := "F1: Ping Site F2 Ping All Node F3 Reset Site F4 Reset Node"
    title   := "Manage Site/Node"

    gc_support.Construct_Console_Menu(current_line.Window,title,message)
}


func bool_display( value bool)string{
    if value == true {
        return "true"
    }
    return "false"
}


const marker string="****************************************************************************\n"

func site_key_handler(ch gc.Key )bool{
     return_value := false
     switch ch {
        case gc.KEY_F1:
             current_line.Return_cmd(marker)
             status := site_rpc_structure.Ping()
             status_line := "Site    Active: "+bool_display(status)+"\n"
             current_line.Return_cmd(status_line)
            current_line.Return_cmd(marker)
    
        case gc.KEY_F2:
           current_line.Return_cmd(marker)
           for processor, _ := range node_rpc_structures.Driver_array {
               status := node_rpc_structures.Ping(processor)
               status_line := "Node: "+processor+" Active: "+bool_display(status)+"\n"
               current_line.Return_cmd(status_line)
           }
           current_line.Return_cmd(marker)
           
        case gc.KEY_F3:
             ;// Reset Site Controller
             
        case gc.KEY_F4:
            ;// Reset Node
        
        default:
            ;    
     }
    return return_value
}
