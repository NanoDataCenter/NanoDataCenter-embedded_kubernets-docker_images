package main

import "fmt"
import "time"
import "io/ioutil"
import "io/fs"

import "os"
import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"

var file_base string

func main(){

    
    
	var config_file ="/data/redis_configuration.json"
	
	var site_data_store map[string]interface{}

    
	site_data_store = get_site_data.Get_site_data(config_file)
    fmt.Println(site_data_store)
    graph_query.Graph_support_init(&site_data_store)
	redis_handlers.Init_Redis_Mutex()
    properties := graph_query.Common_qs_search(&[]string{"RPC_SERVER:SITE_FILE_SERVER","RPC_SERVER"})
    fmt.Println("properties",properties)
    property := properties[0]
    file_base = graph_query.Convert_json_string(property["directory"])
    fmt.Println("file_base",file_base)
	data_handler.Data_handler_init(&site_data_store)	
 	  
     search_list := []string{"RPC_SERVER:SITE_FILE_SERVER","RPC_SERVER"}

     handlers := data_handler.Construct_Data_Structures(&search_list)
    
     driver := (*handlers)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
	
	//driver.Add_handler("ping",ping)
	
	driver.Add_handler( "read",load_file)
	driver.Add_handler( "write",save_file)
	driver.Add_handler( "file_exists",file_exists)
	driver.Add_handler( "delete_file",delete_file)
	driver.Add_handler( "file_directory",file_directory)
	driver.Add_handler( "make_dir",mkdir)
	driver.Json_Rpc_start()
	
	for true {
	  //fmt.Println("main spining")
	  time.Sleep(time.Second*10)
	}
   
}

/*
func ping( parameters map[string]interface{} ) map[string]interface{}{

   parameters["status"] = true
   return parameters

}
*/
	
func load_file( parameters map[string]interface{} ) map[string]interface{}{
   p_file_name := parameters["file_name"].(string)
  
   file_name := file_base+p_file_name
   data, err := ioutil.ReadFile(file_name)
  
  if err == nil {
        parameters["status"] = true
		parameters["results"] = string(data)
  } else {
        parameters["status"] = false
		parameters["results"] = ""
  }
  
  return parameters

}

func save_file( parameters map[string]interface{} ) map[string]interface{}{
  
   p_file_name := parameters["file_name"].(string)
 
  
   p_data := []byte(parameters["data"].(string))
   file_name := file_base+p_file_name
  //fmt.Println("save_file",file_name,p_data)
   err := ioutil.WriteFile(file_name,p_data,0666)
   //fmt.Println(err)
  if err == nil {
       parameters["status"] = true
		
  } else {
        parameters["status"] = false
		
  }

  return parameters
}

func delete_file( parameters map[string]interface{} ) map[string]interface{}{
  
   p_file_name := parameters["file_name"].(string)
   file_name := file_base+p_file_name

  err := os.Remove(file_name)
  //fmt.Println("err",err)
  if err == nil {
        parameters["status"] = true
		parameters["results"] = ""
  } else {
        parameters["status"] = false
		parameters["results"] = ""
  }

  return parameters
}



func fileExists(filename string) (bool,bool) {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false,false
    }
    return true, info.IsDir()
}


func file_exists( parameters map[string]interface{} ) map[string]interface{}{

  p_file_name := parameters["file_name"].(string)
  
  file_name := file_base+p_file_name
  exists , directory := fileExists(file_name)
  parameters["status"] = exists
  parameters["directory"] = directory

  return parameters
}
  


 
  
func mkdir( parameters map[string]interface{} ) map[string]interface{}{
   
   p_path := parameters["path"].(string)
  
   path := file_base+p_path
  //fmt.Println("path",path)
  err := os.MkdirAll(path,0666)
 
  if err == nil {
       parameters["status"] = true
		parameters["results"] = ""
  } else {
        parameters["status"] = false
		parameters["results"] = ""
  }

  return parameters
}

func convert_to_file_names( input []fs.FileInfo )[]string {

   return_value := []string{}
  for _ , file_info := range input{
     return_value = append(return_value,file_info.Name())
  }
  return return_value
}

func file_directory( parameters map[string]interface{} ) map[string]interface{}{

   p_path := parameters["path"].(string)
  
   path := file_base+p_path
  fmt.Println("path",path)
  c, err := ioutil.ReadDir(path)
  if err != nil {
        fmt.Println(err)
        panic("done")
        parameters["status"] = false
		parameters["results"] = ""
  } else {
      parameters["status"] = true
	 parameters["results"] = convert_to_file_names(c)
  }
   fmt.Println("parameters",parameters)
   return parameters

}

