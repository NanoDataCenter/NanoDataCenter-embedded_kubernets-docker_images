package main


//import "fmt"
import "io/ioutil"
import "io/fs"
import "os"
import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"

import "lacima.com/redis_support/redis_file"


const file_base = "files"

var driver *redis_file.Redis_File_Struct

func main(){
    var config_file = "/data/redis_server.json"
	var site_data map[string]interface{}
	
	
	site_data = get_site_data.Get_site_data(config_file)
    graph_query.Graph_support_init(&site_data)
	address  :=  site_data["host"].(string)
    port  := 	int(site_data["port"].(float64))
    redis_file.Create_redis_data_handle(address,port)
	driver = redis_file.Construct_File_Struct(  ) 
	(driver).FlushDB()
	load_files(file_base)
	
}





func load_files( path string ){

  
  files := file_directory( path  )
 
  for _,filename := range files {
    path_filename := path+"/"+filename
    info, err := os.Stat(path_filename)
	if err == nil {
      if info.IsDir() == true {
	    create_marker(path_filename)
	    load_files(path_filename)
	  }else{
         process_file(path_filename)	
	  }	
	}
  }

}


func create_marker(path string){
  //fmt.Println("create directory ",path)
  driver.Set(path,"directory")
}


func process_file( path_file_name string) {

  //fmt.Println("process_file   ",path_file_name)
  var data, err = ioutil.ReadFile(path_file_name)
  
  if err == nil {
    driver.Set(path_file_name,string(data))
     
  }

}





func convert_to_file_names( input []fs.FileInfo )[]string {

  var return_value = []string{}
  for _ , file_info := range input{
     return_value = append(return_value,file_info.Name())
  }
  return return_value
}

func file_directory( path string ) []string{

  var return_value []string
  c, err := ioutil.ReadDir(path)
  if err == nil {
     return_value = convert_to_file_names(c)
  }
   return return_value
   

}


/*
test driver for reference
func test_driver(){
  files := driver.Ls("*")
  fmt.Println("data_dump","files/application_files/controller_cable_assignment.json",driver.Get("files/application_files/controller_cable_assignment.json"))
  fmt.Println("files",files)
  driver.Rm("files/application_files*")
  files = driver.Ls("*")
  fmt.Println("files",files)
}  
*/