package main

import "fmt"

import "lacima.com/site_data"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/server_libraries/file_server_library"




func main(){

    
		
	var config_file = "/data/redis_configuration.json"
	
	var site_data_store map[string]interface{}

	site_data_store = get_site_data.Get_site_data(config_file)
    graph_query.Graph_support_init(&site_data_store)
	redis_handlers.Init_Redis_Mutex()
	data_handler.Data_handler_init(&site_data_store)	
    search_list := []string{"RPC_SERVER:SITE_FILE_SERVER","RPC_SERVER"}
 	var fs_handle = file_server_lib.File_Server_Init(&search_list)
    fmt.Println(fs_handle.Ping())	
    fmt.Println((fs_handle).Ping())
	fmt.Println((fs_handle).File_directory(""))
	fmt.Println((fs_handle).Mkdir("test_path"))
	fmt.Println((fs_handle).Mkdir("test_path_1"))
	fmt.Println((fs_handle).Mkdir("test_path_2"))
	fmt.Println((fs_handle).Mkdir("test_path_3"))
	
	fmt.Println((fs_handle).File_exists("test_path"))
	fmt.Println((fs_handle).File_exists("test_path_9"))
	fmt.Println((fs_handle).File_directory(""))
	fmt.Println((fs_handle).Write_file("test_path/test_file.test","hi\nthere\nbrown\ncow"))
	fmt.Println((fs_handle).Read_file("test_path/test_file.test"))
	fmt.Println((fs_handle).File_directory("test_path"))
	fmt.Println((fs_handle).Write_file("test_path/test_file_a.test","++++++++++++++++++++++"))
	fmt.Println((fs_handle).Read_file("test_path/test_file_a.test"))
	fmt.Println((fs_handle).File_directory("test_path"))
	fmt.Println((fs_handle).Delete_file("test_path/test_file_a.test"))
    fmt.Println((fs_handle).File_directory("test_path"))
    fmt.Println((fs_handle).File_directory("/app_data_files"))
	fmt.Println("done")
	
}
