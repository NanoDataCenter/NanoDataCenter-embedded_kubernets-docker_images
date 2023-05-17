package main

/*
 * This is a utility package, ie runs and completes
 * The purpose is to load passwords from a command line argument
 * and store the passwords in a redis db specified by a second
 * command line argument.
 * 
 * The format for running this command is
 * ./file_loader password_file,db number
 * 
 * has a dependency on a /data/ mount for configuration data
 * to access redis data file_base
 * 
 */ 



import "fmt"

import "os"
import "encoding/csv"
import "strconv"
import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"

import "lacima.com/redis_support/redis_file"


var secret_file string

var driver *redis_file.Redis_File_Struct

func main(){
    var config_file ="/data/redis_configuration.json"
	var site_data map[string]interface{}
	
	secret_file = os.Args[1]
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
    fmt.Println("made it here ")
    csvReader(secret_file)
	
	
}


func csvReader(secret_file string) {
  
  recordFile, err := os.Open(secret_file)
  if err != nil {
   fmt.Println("An error encountered ::", err)
   panic("")
  }
  reader := csv.NewReader(recordFile)
  records,err1 := reader.ReadAll()
  if err1 != nil {
    fmt.Println("An error encountered ::", err1)
    panic("")
  }
  target_records := records[1:] // skip header
  //fmt.Println("target_records",target_records)
  for _,element := range target_records{
    driver.HSet(element[0],element[1],element[2])   
      
  }
 
}

     





