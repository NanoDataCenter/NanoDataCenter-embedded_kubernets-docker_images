package connection_management

import (
gc "github.com/gbin/goncurses"
"lacima.com/go_terminals/go_terminal_library"
"lacima.com/go_terminals/site_node_control/central_db"
"lacima.com/go_terminals/docker_control"

	
)


var current_line *gc_support.LINE_BUFFER


func Connection_management_launcher(){
    
    gc_support.Construct_Console(gc_support.Commmand_handler,connection_key_handler,init_handler)
    
    
    
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
            panic("illegal state")
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
            setup_standalone_master()
            current_line.Refresh()
    
        case gc.KEY_F2:
            setup_standalone_slave()
            
           
            
        case gc.KEY_F3:
            connect_to_site_controller()
            
        
        default:
            ;    
     }
    return return_value
}

func process_redis_connected(ch gc.Key )bool{
     return_value := false
     switch ch {
        
         case gc.KEY_F1:
            remove_existing_connection()
            current_line.Refresh()
            
        case gc.KEY_F2:
            reload_graphic_container()
            current_line.Refresh()

        
       
        default:
            ;    
     }
    return return_value
}

    

func setup_standalone_master(){
    if central_db.Test_for_redis_connection() == true {
        
        status := remove_existing_connection()
        
        
        if status == false {
            return
        }
    }
    gc_support.Pop_up_Display("Starting Containers",make([]string,0))
    central_db.Do_master_setup()
    gc_support.Pop_up_close()
    current_line.Return_cmd("Redis Server and Containers Started\n")
    current_line.Refresh()
    gc.UpdatePanels()
    gc.Update()
    central_db.Connection_state = 1
}    




func remove_existing_connection()bool{
    result := gc_support.Pop_up_confirmation("Delete And Remove All Containers",[]string{"An Existing Redis Server is Present","This command will stop all containers and",
                                                                                    "And Delete All Containers"})
    current_line.Refresh()
    if result == true{
         gc_support.Pop_up_Display("Stop Running Containers",make([]string,0))
         docker_control.Stop_Running_Containters()
         current_line.Return_cmd("Running Containers Stopped\n")
         gc_support.Pop_up_close()
         current_line.Refresh()
         
         gc.UpdatePanels()
         gc.Update()
         gc_support.Pop_up_Display("Remove  Containers",make([]string,0))
    
        docker_control.Remove_All_Containers()
        current_line.Return_cmd("Containers Removed\n")
        gc_support.Pop_up_close()
        current_line.Refresh()
        central_db.Connection_state = 0
        current_line.Refresh()
         
        gc.UpdatePanels()
        gc.Update()
    }
    return result
}

func reload_graphic_container(){
    current_line.Return_cmd("Registry Starting To Be Reloaded\n")
    central_db.Reload_graphic_db()
    current_line.Return_cmd("Registry Reloaded\n")
    
}

func verify_redis_server()bool{
   result := true
   
   if central_db.Test_for_redis_connection() == false {
        error_message := []string{"No Redis Server","Start Master Node","Or bad configuration data"}
        gc_support.Pop_up_alert("No Redis Server",error_message)
        result = false
    } 
    return result 
    
}


func setup_standalone_slave(){
   current_line.Return_cmd("Start in standalone slave\n")
   if verify_redis_server() == true {
       central_db.Do_slave_setup()
       current_line.Return_cmd("Stand only slave setup\n")
   }else{
     current_line.Return_cmd("Stand only slave setup aborted\n")
   }
    
}

func connect_to_site_controller(){
   if verify_redis_server() {
      central_db.Setup_Structures()
      current_line.Return_cmd("Connected to Site Controller")
      central_db.Connection_state = 3
   }
    
}
