package main

import "fmt"
import "os"
import "bytes"
import "io/ioutil"
import "syscall"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/site_data"
import "lacima.com/Patterns/logging_support"
import "github.com/msgpack/msgpack-go"
import ps "github.com/mitchellh/go-ps"


func main(){

  var config_file = "/data/redis_server.json"
  var site_data_store map[string]interface{}
  
  var file_name = os.Args[1] // location of error file
  var file_data, err = ioutil.ReadFile(file_name)
  if err != nil {
        panic("no error data file ")
  }
  fmt.Println("file_data",file_data)
  site_data_store = get_site_data.Get_site_data(config_file)
 
  
  graph_query.Graph_support_init(&site_data_store)
  redis_handlers.Init_Redis_Mutex()
  data_handler.Data_handler_init(&site_data_store)
  container_name := os.Getenv("CONTAINER_NAME")
  fmt.Println("container_name",container_name)
  
  
  incident_log := logging_support.Construct_incident_log([]string{"CONTAINER:"+container_name,"INCIDENT_LOG:process_control_failure","INCIDENT_LOG"} )
  
  var b bytes.Buffer	
  msgpack.Pack(&b,file_data)
  new_value := b.String()
  (*incident_log).Log_data( false, new_value, new_value )
  /*
     At this point there may be zombie spawn processes
	 this section of code will remove any processes that
	 are not either the error_logger or bash program
  */
  
   processList, _ := ps.Processes()
   for x := range processList {
        var process ps.Process
        process = processList[x]
		name := process.Executable()
		fmt.Println("name",name)
		if ( name != "bash")&&(name != "error_logger"){
		   fmt.Println("killing ",name)
		   syscall.Kill(process.Pid(),syscall.SIGINT)
		}
        fmt.Printf("%d\t%s\n",process.Pid(),process.Executable())

        // do os.* stuff on the pid
    }
  

}

