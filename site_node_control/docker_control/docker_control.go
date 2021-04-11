package docker_control
import "fmt"
import "strings"
//import "os"
import "os/exec"
//import "log"
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
func isValueInList(value string, list []string) bool {
    for _, v := range list {
        if v == value {
            return true
        }
    }
    return false
}

func Filter_first(start int, text string) []string {

  var return_value  []string
  var lines = strings.Split(text,"\n")
  //fmt.Println(len(lines),lines[1])
  if len(lines) <= start {
    return return_value
  }
   
  var line1 = lines[1:]
  for _, data := range line1 {
      if len(data) > 0 {
  	       var data_list = strings.Fields(data)
           //fmt.Println(data_list[0])		   
	       return_value = append(return_value,data_list[0])
	   }
	 }
   return return_value
 }

func Filter_last(start int ,text string) []string {

  var return_value []string
  var lines = strings.Split(text,"\n")
  if len(lines) <= start {
    return return_value
  }

  
  for _, data := range lines {
       if len(data) > 0 {
	   var data_list = strings.Fields(data) 
	   return_value = append(return_value,data_list[len(data_list)-1])
	 }
   }
   return return_value
  
  
}

func Exists( container string ) bool {

   var container_list = Containers_ls_all()
   return isValueInList(container,container_list) 
    
}


func Image_Exists( container string ) bool {

   var container_list = Images()
   return isValueInList(container,container_list) 
    
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


func Containers_ls_runing() []string {
    
	output := System("docker container ls ")
	return Filter_last(1,output)
   
}

func Container_is_running( match_container string) bool {

  var running_containers = Containers_ls_runing()
  for _,container := range running_containers{
    if match_container == container {
	  return true
	}
  }
  return false
}

         
func Containers_ls_all() []string {

    
	output := System("docker container ls -a")
	return Filter_last(1,output)
   
}
       
       
func Container_start(container string) {
      System("docker start "+container)
}	  
       
func Container_stop(container string) {
       System("docker stop "+container)
}
     
func Container_Run(startup_script string)string{

   return System_shell(startup_script)
}
	 
func Container_up(container,startup_script string) {
     if Exists(container){
	    Container_start(container)
	 } else {
	   System(startup_script)
     }
}	 
 
func Container_rm(container string) string {
       if Exists(container) {
	     return System("docker rm "+container)
		} else {
		  return "OK"
		}
           
}

             
func  Prune(){
       return_value := System("docker images -qf dangling=true ")
       var return2 = Filter_first(1,return_value)
	   for _, name := range return2 {
	     //fmt.Println(i,name)
		 Image_rmi(name)
	   }
	      
       
}      
       
func Push(image string){
        System("docker push "+image)
}
      
func  Pull(image string ){
       System("docker pull "+image)
}

func  Images()[]string{
      return_value := System("docker images ")
	  var return2 = Filter_first(1,return_value)
	  
	  return return2
}
   
func  Image_rmi(deleted_image string)string {
      return System("docker rmi "+deleted_image)
}     
	  
 
func Upgrade_container(container,container_image, build_script string)string{
     if Exists(container) {
	    Container_stop(container)
		Container_rm(container)
     }
	 Image_rmi(container_image)
	 Pull(container_image)
	 return System(build_script)
}   
  
 
          