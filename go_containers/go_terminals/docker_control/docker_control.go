package docker_control
import "fmt"
import "strings"
//import "os"
//import "os/exec"
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

func Filter_first_special( text string) []string {

  var return_value  []string
  var lines = strings.Split(text,"\n")
  
   
  line1 := lines
  for _, data := range line1 {
      if len(data) > 0 {
  	       var data_list = strings.Fields(data)
           //fmt.Println(data_list[0])		   
	       return_value = append(return_value,data_list[0])
	   }
	 }
   return return_value
 }

func Filter_first( text string) []string {

  var return_value  []string
  var lines = strings.Split(text,"\n")
  
   
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

func Filter_last(text string) []string {

  var return_value []string
  var lines = strings.Split(text,"\n")
  
  

  lines_1 := lines[1:]
  
  for _, data := range lines_1 {
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


func Image_Exists( image string ) bool {

   var image_list = Images()
   //fmt.Println("container_list",image_list)
   //fmt.Println("image_list",image)
   return isValueInList(image,image_list) 
    
}

func handlepanic() {
  
    if a := recover(); a != nil {
        fmt.Println("RECOVER", a)
    }
}


func System_shell( script string )string{
   var script_list = strings.Fields(script) 
   defer handlepanic()
   out := gosh.Gosh(script_list).Output()
   
   return string(out)
}
/*
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
*/


func Containers_ls_runing() []string {
    
	output := System_shell("docker container ps  ")
    //fmt.Println("output",output)
	return Filter_last(output)
   
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

    
	output := System_shell("docker container ls -a ")
	return Filter_last(output)
   
}
       
       
func Container_start(container string) {
      System_shell("docker start "+container)
}	  
       
func Container_stop(container string) {
       System_shell("docker stop "+container)
}
     
func Container_Run(startup_script string)string{

   return System_shell(startup_script)
}
	 
func Container_up(container,startup_script string) {
     if Exists(container){
        
	    Container_start(container)
	 } else {
        
	   fmt.Println(System_shell(startup_script))
     }
}	 


func Stop_Running_Containters(){
 running_containers := Containers_ls_runing()
 for _,container := range running_containers {
     Container_stop(container)
 }
 
    
} 
 
func Remove_All_Containers(){
   containers := Containers_ls_all()
   for _,container := range containers {
       Container_rm(container)
   }
    
   
}

func Container_rm(container string) string {
       if Exists(container) {
	     return System_shell("docker rm "+container)
		} else {
		  return "OK"
		}
           
}

             
func  Prune()[]string{
       output := make([]string,0)
       return_value := System_shell("docker images -qf dangling=true ")
      
       var return2 = Filter_first_special(return_value)
      
	   for _, name := range return2 {
	     output = append(output,"image "+name +" Pruned ")
		 Image_rmi(name)
	   }
	   return output   
       
}      
       
func Push(image string){
        System_shell("docker push "+image)
}
      
func  Pull(image string ){
       System_shell("docker pull "+image)
}

func  Images()[]string{
      return_value := System_shell("docker images ")
	  var return2 = Filter_first(return_value)
	  
	  return return2
}
   
func  Image_rmi(deleted_image string)string {
      return System_shell("docker rmi "+deleted_image)
}     
	  
 
func Upgrade_container(container,container_image, build_script string)string{
     if Exists(container) {
	    Container_stop(container)
		Container_rm(container)
     }
	 Image_rmi(container_image)
	 Pull(container_image)
	 return System_shell(build_script)
}   
  
 
          
