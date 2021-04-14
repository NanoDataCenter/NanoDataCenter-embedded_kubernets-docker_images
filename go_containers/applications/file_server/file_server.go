package main

import "fmt"
import "time"
import "io/ioutil"

import "os"
import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"




func main(){

    //mount_usb_drive()
      
		
	//var config_file = "/data/redis_server.json"
	var config_file = "/home/pi/mountpoint/lacuma_conf/site_config/redis_server.json"
	var site_data_store map[string]interface{}

	site_data_store = get_site_data.Get_site_data(config_file)
    graph_query.Graph_support_init(&site_data_store)
	redis_handlers.Init_Redis_Mutex()
	data_handler.Data_handler_init(&site_data_store)	
 	  
    var search_list = []string{"FILE_SERVER","FILE_SERVER"}

    var handlers = data_handler.Construct_Data_Structures(&search_list)
    
    var driver = (*handlers)["FILE_SERVER_RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
	
	driver.Add_handler("ping",ping)
	
	driver.Add_handler( "load",load_file)
	driver.Add_handler( "save",save_file)
	driver.Add_handler( "file_exists",file_exists)
	driver.Add_handler( "delete_file",delete_file)
	driver.Add_handler( "file_directory",file_directory)
	driver.Add_handler( "make_dir",mkdir)
	driver.Rpc_start()
	
	for true {
	  fmt.Println("main spining")
	  time.Sleep(time.Second*10)
	}
   
}


func ping( parameters *map[string]interface{} ) *map[string]interface{}{

   (*parameters)["status"] = true
   return parameters

}
	
func load_file( parameters *map[string]interface{} ) *map[string]interface{}{

  var path = "/files/"+(*parameters)["path"].(string)+"/"+(*parameters)["file_name"].(string)
  var data, err = ioutil.ReadFile(path)
  if err != nil {
        (*parameters)["status"] = true
		(*parameters)["results"] = data
  } else {
        (*parameters)["status"] = false
		(*parameters)["results"] = ""
  }
  return parameters

}

func save_file( parameters *map[string]interface{} ) *map[string]interface{}{
  
 
  var path = "/files/"+(*parameters)["path"].(string)+"/"+(*parameters)["file_name"].(string)
  
  var err = ioutil.WriteFile(path,[]byte((*parameters)["data"].(string)),0666)
  if err != nil {
       (*parameters)["status"] = true
		(*parameters)["results"] = ""
  } else {
        (*parameters)["status"] = false
		(*parameters)["results"] = ""
  }
  return parameters
}

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}
func file_exists( parameters *map[string]interface{} ) *map[string]interface{}{
 
 
  var path = "/files/"+(*parameters)["path"].(string)+"/"+(*parameters)["file_name"].(string)
  if  fileExists(path) == true{
 
       (*parameters)["status"] = true
		(*parameters)["results"] = ""
  } else {
        (*parameters)["status"] = false
		(*parameters)["results"] = ""
  }
  return parameters
}
  

func delete_file( parameters *map[string]interface{} ) *map[string]interface{}{
  
 var path = "/files/"+(*parameters)["path"].(string)+"/"+(*parameters)["file_name"].(string)
  err := os.Remove(path)
  if err != nil {
       (*parameters)["status"] = true
		(*parameters)["results"] = ""
  } else {
        (*parameters)["status"] = false
		(*parameters)["results"] = ""
  }
  return parameters
}
 
  
func mkdir( parameters *map[string]interface{} ) *map[string]interface{}{
   
  var path = "/files/"+(*parameters)["path"].(string)+"/"+(*parameters)["file_name"].(string)
  err := os.MkdirAll(path,0666)
  if err != nil {
       (*parameters)["status"] = true
		(*parameters)["results"] = ""
  } else {
        (*parameters)["status"] = false
		(*parameters)["results"] = ""
  }
  return parameters
}


func file_directory( parameters *map[string]interface{} ) *map[string]interface{}{
var path = "/files/"+(*parameters)["path"].(string)+"/"+(*parameters)["file_name"].(string)
  c, err := ioutil.ReadDir(path)
  if err != nil {
        (*parameters)["status"] = false
		(*parameters)["results"] = ""
  } else {
      (*parameters)["status"] = true
		(*parameters)["results"] = c
  }
   return parameters

}

