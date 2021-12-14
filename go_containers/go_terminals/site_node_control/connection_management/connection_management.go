package connection_management

import (
gc "github.com/gbin/goncurses"
"lacima.com/go_terminals/go_terminal_library"
"lacima.com/go_terminals/site_node_control/central_db"
)


var current_line *gc_support.LINE_BUFFER

func Connection_management_launcher(){
    
    gc_support.Construct_Console(nil,connection_key_handler,init_handler)
    
    
    
}

func init_handler(){
     current_line = gc_support.Return_SubWindow_Draw_Structures()
     generate_main_screen_message(current_line.Window)
}

func generate_main_screen_message(window *gc.Window){
    
    var title   string
    var message string
    connected_message := "F1: Disconnect Redis F2: Reload Registry"
    starting_msg      := "Connection MGR "
    switch central_db.Connection_state {
        case 0:
            title = central_db.Generate_title(starting_msg)
            message = "F1: Setup Standalone Master F2: Stand alone Slave F3: Connect to Site Controller"
        case 1:
            title = central_db.Generate_title(starting_msg)
            message =connected_message
        case 2:
            title = central_db.Generate_title(starting_msg)
            message = connected_message
        case 3:
            title = central_db.Generate_title(starting_msg)
            message = connected_message
        default:
            panic("illegal statue")
    }
    gc_support.Construct_Console_Menu(current_line.Window,title,message)
}
        
var connection_state int  
/*
 * 0  -- not connected
 * 1  -- redis standalone
 * 2  -- slave standalone
 * 3  -- site controller active
 */

func connection_key_handler( ch gc.Key)bool{
   
     
      connection_state := central_db.Connection_state 
 
     
      connection_state = central_db.Connection_state
          
      gc.UpdatePanels()
      gc.Update()
      
      return_value := false
      switch connection_state {
          case 0:
              return_value = process_not_connected(ch)
          case 1,2,3:
             return_value =  process_redis_connected(ch)
 
          default:
              panic("bad case")
      }
       if connection_state != central_db.Connection_state {
          generate_main_screen_message(current_line.Window)
      }
      return return_value 
 
}

func process_not_connected(ch gc.Key )bool{
     return_value := false
     switch ch {
        case gc.KEY_F1:
            central_db.Connection_state = 1
           
    
        case gc.KEY_F2:
            central_db.Connection_state = 2
           
            
        case gc.KEY_F3:
            central_db.Connection_state = 3
            
        
        default:
            ;    
     }
    return return_value
}

func process_redis_connected(ch gc.Key )bool{
     return_value := false
     switch ch {
        
         case gc.KEY_F1:
            central_db.Connection_state = 0
            
        case gc.KEY_F2:
            ; //reload graph data base

        
       
        default:
            ;    
     }
    return return_value
}

    

          
/*

func commmand_handler(input string )string{
    var return_value string
    return_value = "bad command "+input
    defer func() {
        if r := recover(); r != nil {
            return_value = input +"  bad command"
            
        }
    }()
    if len(input) > 0 {
      return_value = shell_utils.System_mshell(input)
    }
    return return_value
}
*/

