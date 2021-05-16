package main

/*
 * This is a utility package, ie runs and completes
 * The purpose is to load files from a command line argument
 * and store the file data in a redis db specified by a second
 * command line argument.
 * 
 * The format for running this command is
 * ./file_loader directory,db number
 * 
 * has a dependency on a /data/ mount for configuration data
 * to access redis data file_base
 * 
 */ 



import "fmt"
import "io/ioutil"
import "io/fs"
import "os"
import "strconv"
import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"

import "lacima.com/redis_support/redis_file"


var file_base string

var driver *redis_file.Redis_File_Struct

func main(){
    var config_file ="/data/redis_configuration.json"
	var site_data map[string]interface{}
	
	file_base = os.Args[1]
	file_db, err := strconv.Atoi(os.Args[2])
	if err != nil {
	   panic("bad db number")
	}
	site_data = get_site_data.Get_site_data(config_file)
    graph_query.Graph_support_init(&site_data)
	address  :=  site_data["host"].(string)
    port  := 	int(site_data["port"].(float64))
    redis_file.Create_redis_data_handle(address, port , file_db )
	driver = redis_file.Construct_File_Struct(  ) 
	(driver).FlushDB()
    //fmt.Println("made it here ",file_base)
	load_files_initial(file_base)
	
}



func load_files_initial( path  string ){

  
  
  files := file_directory( path  )
  //fmt.Println("files",files)
  for _,filename := range files {
    path_filename := path+"/"+filename
	target_filename :=  filename
    info, err := os.Stat(path_filename)
	if err == nil {
      if info.IsDir() == true {
	    create_marker(target_filename)
	    load_files(path_filename,target_filename)
	  }else{
         process_file(path_filename,target_filename)	
	  }	
	}
  }

}

func load_files( path ,target_path string ){

  fmt.Println("path",path,target_path)
  
  files := file_directory( path  )
  //fmt.Println("files",files)
  for _,filename := range files {
    path_filename := path+"/"+filename
	target_filename := target_path +"/"+ filename
    info, err := os.Stat(path_filename)
	if err == nil {
      if info.IsDir() == true {
	    create_marker(target_filename)
	    load_files(path_filename,target_filename)
	  }else{
         process_file(path_filename,target_filename)	
	  }	
	}
  }

}


func create_marker(path string){
  //fmt.Println("create directory ",path)
  driver.Set(path,"directory")
}


func process_file( path_file_name,target_name string) {

  //fmt.Println("process_file   ",path_file_name)
  var data, err = ioutil.ReadFile(path_file_name)
  
  if err == nil {
    driver.Set(target_name,string(data))
     
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
