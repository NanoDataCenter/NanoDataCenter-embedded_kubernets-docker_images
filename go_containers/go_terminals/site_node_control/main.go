package main
import (
	gc "github.com/gbin/goncurses"
    "lacima.com/go_terminals/go_terminal_library"
    "lacima.com/go_terminals/site_node_control/docker_utilities"
    //"lacima.com/go_terminals/site_node_control/site_management"
    //"lacima.com/go_terminals/site_node_control/node_management"
    //"lacima.com/go_terminals/site_node_control/redis_management"
    "lacima.com/go_terminals/site_node_control/central_db"
    "lacima.com/go_terminals/site_node_control/connection_management"
    "lacima.com/go_terminals/site_node_control/site_node_manager"
     "lacima.com/go_terminals/site_node_control/container_manager"
)

	//"fmt"




   

func main(){
    central_db.Connection_state = 0
    central_db.Fill_in_slave_data()
    
    gc_support.Init_SoftKey()
    defer gc.End()
   
    initial_screen()
    
	

    
}


func initial_screen(){
   
	gc_support.Construct_launch_panel( initial_key_handler )    
	
    
}   

func generate_main_screen_message(window *gc.Window){
    
    var title   string
    var message []string
    switch central_db.Connection_state {
        case 0:
            title = "Not Connected to Redis DB "
            message = []string{"F1: Docker Utilities","F2: Connection Mananger","F8: Exit"}
        case 1:
            title = "Connect Master Setup No Site Node controller "
            message = []string{"F1: Docker Utilities","F2: Connection Mananger", "F8: Exit"}
        case 2:
            title =  "Redis DB Connecte Slave_setup Setup No Site Node controller "
            message = []string{"F1: Docker Utilities","F2: Connection Mananger", "F8: Exit"}
        case 3:
            title = "SITE CONTROLLER ACTIVE "
            message = []string{"F1: Docker Utilities","F2: Connection Mananger","F3: Site/Node Controller","F4: Container Controller","F8: Exit"}
        default:
            panic("illegal statue")
    }
    gc_support.Launch_panel_display(window,title,message)
}
            
             
    



var connection_state int  
/*
 * 0  -- not connected
 * 1  -- redis standalone
 * 2  -- slave standalone
 * 3  -- site controller active
 */

func initial_key_handler( window *gc.Window){
   
   generate_main_screen_message(window)
   connection_state := central_db.Connection_state 
   for true {
      if connection_state != central_db.Connection_state {
          generate_main_screen_message(window)
      }
      connection_state = central_db.Connection_state
          
      gc.UpdatePanels()
      gc.Update()
      ch := gc_support.Stdscr.GetChar()
      return_value := false
      switch connection_state {
          case 0:
              return_value = process_not_connected(ch)
          case 1:
             return_value =  process_master_standalone(ch)
          case 2:
              return_value = process_slave_standalone(ch)
          case 3:
              return_value = process_site_controller_active(ch)
          default:
              panic("bad case")
      }
      if return_value == true {
          return
      }
   }
}

func process_not_connected(ch gc.Key )bool{
     return_value := false
     switch ch {
        case gc.KEY_F1:
            docker_utils.Basic_docker_launcher()
    
        case gc.KEY_F2:
            connection_management.Connection_management_launcher()
        
       
            
        case gc.KEY_F8:
            return_value = true
        default:
            ;    
     }
    return return_value
}

func process_master_standalone(ch gc.Key )bool{
     return_value := false
     switch ch {
        case gc.KEY_F1:
            docker_utils.Basic_docker_launcher()
    
        case gc.KEY_F2:
            connection_management.Connection_management_launcher()        

        
        
        case gc.KEY_F8:
            return_value = true
        default:
            ;    
     }
    return return_value
}

          
func process_slave_standalone(ch gc.Key)bool{
     return_value := false
     switch ch {
       
        case gc.KEY_F1:
            docker_utils.Basic_docker_launcher()
    
        case gc.KEY_F2:
            connection_management.Connection_management_launcher()
        
        case gc.KEY_F8:
            return_value = true
        default:
            ;    
     }
    return return_value
}

          
func process_site_controller_active(ch gc.Key )bool{
     return_value := false
     switch ch {
         case gc.KEY_F1:
            docker_utils.Basic_docker_launcher()
    
        case gc.KEY_F2:
            connection_management.Connection_management_launcher()     

         case gc.KEY_F3:
            site_node_management.Site_node_management_launcher()
            
        case gc.KEY_F4:
            container_management.Container_management_launcher()
             
        case gc.KEY_F8:
            return_value = true
        default:
            ;    
     }
    return return_value
}

          
/*
          
      switch ch {
        case gc.KEY_F1:
            docker_utils.Basic_docker_launcher()
        case gc.KEY_F2:
            site_management.Site_mngr_launcher()
        case gc.KEY_F3:
             node_management.Node_mngr_launcher()
        case gc.KEY_F4:
            redis_management.Redis_mngr_launcher()
        case gc.KEY_F5:
            ; //redis_utils.Redis_mngr_launcher()
        case gc.KEY_F6:
           ; // redis_utils.Redis_mngr_launcher()
        
        case gc.KEY_F8:
            return
        default:
            ;
      }
   }
}

*/


