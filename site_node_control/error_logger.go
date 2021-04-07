package main

//import "fmt"
import "os"
import "bytes"
import "io/ioutil"
import "time"
import "site_control.com/site_data"
import "site_control.com/redis_support/graph_query"
import "site_control.com/redis_support/generate_handlers"
import  "site_control.com/redis_support/redis_handlers"
import "github.com/msgpack/msgpack-go"


func main(){

  var config_file = "/mnt/ssd/site_config/redis_server.json"
  var site_data_store map[string]interface{}
  
  var file_name = os.Args[1] // location of error file
  var file_data, err = ioutil.ReadFile(file_name)
  if err != nil {
        panic("no error data file ")
  }
  //fmt.Println("file_data",file_data)
  site_data_store = get_site_data.Get_site_data(config_file)
  var local_node = site_data_store["local_node"].(string)
  
  graph_query.Graph_support_init(&site_data_store)
  redis_handlers.Init_Redis_Mutex()
  data_handler.Data_handler_init(&site_data_store)

  var search_list = []string{"PROCESSOR:"+local_node,"NODE_SYSTEM","SITE_NODE_CONTROL_LOG"}
  var data_element = data_handler.Construct_Data_Structures(&search_list)  
  //fmt.Println(data_element)
  var driver = (*data_element)["ERROR_STREAM"].(redis_handlers.Redis_Stream_Struct)
  //fmt.Println("driver",driver)
  var data = make(map[string]interface{})
  data["time"] = time.Now().UnixNano()
  data["error_trace_back"] = file_data
  var b bytes.Buffer	
  msgpack.Pack(&b,data)
  //fmt.Println(data)
  driver.Xadd(b.String())  
 

}

