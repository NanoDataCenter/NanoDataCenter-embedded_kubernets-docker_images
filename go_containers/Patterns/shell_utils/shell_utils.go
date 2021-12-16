package shell_utils


import "fmt"
import "strings"
import "os/exec"
import "github.com/polydawn/gosh"

/*
Gosh(args ...interface{}) Command {
	return enclose(bake(Opts{
		Launcher: ExecLauncher,
		Env:      getOsEnv(),
		In:       os.Stdin,
		Out:      os.Stdout,
		Err:      os.Stderr,
		OkExit:   []int{0},
	}, args...))
}
*/



func System_mshell( script string )string{
   var script_list = strings.Fields(script) 

   out := gosh.Gosh(script_list).Output()
   return string(out)
}


func System_shell( script string )string{
   var script_list = strings.Fields(script) 

   out := gosh.Gosh(script_list).Output()
   fmt.Println(out)
   return string(out)
}
		
func System( script string )string{

   //fmt.Println("script",script)
   var script_list = strings.Fields(script) 
   

   command := script_list[0]
   //fmt.Println(command)
   //fmt.Println(script_list[1])
   //fmt.Println(len(script_list))
   args := script_list[1:]
   out,_ := exec.Command(command, args...).Output()
   
   return string(out)
}
