package system_control
//import "fmt"
import "strings"
import "io"
import "os/exec"

type Process_Manager_Type struct{

  key             string
  cmd_line        string
  failed           bool
 
  error_log       string
}


func construct_process_manager( key, cmd_line string) *Process_Manager_Type {

  var return_value Process_Manager_Type
  return_value.key = key
  return_value.cmd_line = cmd_line
  return_value.error_log = ""
  return_value.failed = false
  return &return_value

}

func( v *Process_Manager_Type)run(){

   v.failed = false
  
 
      for true {
         defer recover()
         command_list :=  strings.Fields(v.cmd_line)
	     //fmt.Println("command_list",command_list)
		 args := command_list[1:]
		 command :=command_list[0]
         cmd := exec.Command(command,args...)
        
	     stderr, err := cmd.StderrPipe()
	     if err != nil {
		    panic("bad pipe")
		 }
	    // fmt.Println("error pipe is ok")
		// fmt.Println("going to start")
		 if err := cmd.Start(); err != nil {
		     panic("start is failure")
	     }
		 //fmt.Println("start is ok")
         working_buffer , _ := io.ReadAll(stderr)

         cmd.Wait()
		 //fmt.Println("process stopped")
		 if v.failed == false{
		     v.error_log = string(working_buffer)
		 }
	     v.failed = true
      }
  
}

