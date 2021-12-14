package docker_utils

import (
    //"fmt"
	"strings"
    "lacima.com/go_terminals/go_terminal_library"
    "lacima.com/go_terminals/docker_control"
	"lacima.com/Patterns/shell_utils"
    gc "github.com/gbin/goncurses"
)





var current_line *gc_support.LINE_BUFFER


func Basic_docker_launcher( ) {
    

    gc_support.Construct_Console(commmand_handler,special_handlers,init_handler)
    
    
    
    
}

 func init_handler(){
     current_line = gc_support.Return_SubWindow_Draw_Structures()
     title   := "Docker Management "
     message := "F1: prune images F2: running containers F3: all containers F4: container images F5: stop running containers F6: rm all containers"
     gc_support.Construct_Console_Menu(current_line.Window,title,message)
}

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


func special_handlers( input gc.Key )bool{
    current_line = gc_support.Return_SubWindow_Draw_Structures()
    switch input {
        case gc.KEY_F1:
            
            current_line.Return_cmd(dangling())
            
        case gc.KEY_F2:
            current_line.Return_cmd(running_containers())
           
        case gc.KEY_F3:
            current_line.Return_cmd(all_containers())
           
        case gc.KEY_F4:
            current_line.Return_cmd(images())
            
         case gc.KEY_F5:
             current_line.Return_cmd(stop_running())
            
         case gc.KEY_F6:
             current_line.Return_cmd(rm_all_containers())
            
        default:
            ;
    }
    return false
}

func dangling()string{
   output := docker_control.Prune()
   if len(output) == 0{
       return "no images pruned \n"
   }
   return "Key F1 Pruned Images\n"+strings.Join(output,"\n")+"\n"
}

func running_containers()string{
   output := docker_control.Containers_ls_runing()
   if len(output) == 0{
       return "no running containers \n"
   }
   return "Key F2 Running Containers\n"+strings.Join(output,"\n")+"\n"
}


func all_containers()string{
   output := docker_control.Containers_ls_all()
   if len(output) == 0{
       return "no  containers \n"
   }
   return "Key F3 All Containers \n"+strings.Join(output,"\n")+"\n"
}

func images()string{
   output := docker_control.Images()
   if len(output) == 0{
       return "no  images \n"
   }
   return "Key F4 Container Images\n"+strings.Join(output,"\n")+"\n"
}

func stop_running()string{
    docker_control.Stop_Running_Containters()
    
    return "Key F5 Stop Running Containers\n"
    
}


func rm_all_containers()string{
    docker_control.Remove_All_Containers()
     return "Key F6 Remove All Containers\n"
}
