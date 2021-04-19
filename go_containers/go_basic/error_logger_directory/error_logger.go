package main

import "fmt"
import "os"
import "bytes"
import "io/ioutil"
import "time"
import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/generate_handlers"
import  "lacima.com/redis_support/redis_handlers"
import "github.com/msgpack/msgpack-go"


func main(){

  var config_file = "/data/redis_server.json"
  var site_data_store map[string]interface{}
  
  var file_name = os.Args[1] // location of error file
  var file_data, err = ioutil.ReadFile(file_name)
  if err != nil {
        panic("no error data file ")
  }
  //fmt.Println("file_data",file_data)
  site_data_store = get_site_data.Get_site_data(config_file)
 
  
  graph_query.Graph_support_init(&site_data_store)
  redis_handlers.Init_Redis_Mutex()
  data_handler.Data_handler_init(&site_data_store)
  container_name := os.Getenv("CONTAINER_NAME")
  fmt.Println("container_name",container_name)
  
  search_list := []string{"CONTAINER"+":"+container_name,"DATA_STRUCTURES"}
  //fmt.Println("search_list",search_list)
  data_element := data_handler.Construct_Data_Structures(&search_list)  
  //fmt.Println(data_element)
  driver := (*data_element)["CONTROLLER_FAILURE"].(redis_handlers.Redis_Stream_Struct)
 
  data := make(map[string]interface{})
  data["time"] = time.Now().UnixNano()
  data["error_trace_back"] = file_data
  var b bytes.Buffer	
  msgpack.Pack(&b,data)
  fmt.Println(data)
  driver.Xadd(b.String())  
 

}

