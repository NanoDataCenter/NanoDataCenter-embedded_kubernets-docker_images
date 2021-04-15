package main

import "fmt"

import "lacima.com/site_data"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/server_libraries/file_server_library"




func main(){

    
		
	//var config_file = "/data/redis_server.json"
	var config_file = "/home/pi/mountpoint/lacuma_conf/site_config/redis_server.json"
	var site_data_store map[string]interface{}

	site_data_store = get_site_data.Get_site_data(config_file)
    graph_query.Graph_support_init(&site_data_store)
	redis_handlers.Init_Redis_Mutex()
	data_handler.Data_handler_init(&site_data_store)	
 	file_server_lib.File_Server_Init()  
	fmt.Println(file_server_lib.file_directory(""))
    fmt.Println(file_server_lib.mkdir("test_path"))
    fmt.Println(file_server_lib.file_directory(""))
    fmt.Println(file_server_lib.save_file("test_path","test_file.test","hi\nthere\nbrown\ncow"))
    fmt.Println(file_server_lib.load_file("test_path","test_file.test"))
    fmt.Println(file_server_lib.file_directory("test_path"))
    fmt.Println(file_server_lib.file_exists("test_path","test_file.test"))
    fmt.Println(file_server_lib.delete_file("test_path","test_file.test"))
    fmt.Println(file_server_lib.file_directory("test_path"))
	
}
