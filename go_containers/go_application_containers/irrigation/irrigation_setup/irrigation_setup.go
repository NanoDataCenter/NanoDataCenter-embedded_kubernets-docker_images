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

import (
	//"fmt"
	"lacima.com/go_application_containers/irrigation/irrigation_setup/web_services"
   "lacima.com/go_application_containers/irrigation/irrigation_setup/data_base_cleanup"
	"lacima.com/redis_support/generate_handlers"
	"lacima.com/redis_support/graph_query"
	"lacima.com/redis_support/redis_handlers"
	"lacima.com/site_data"
	"time"
)

func main() {
	var config_file = "/data/redis_configuration.json"
	var site_data map[string]interface{}

	site_data = get_site_data.Get_site_data(config_file)
	redis_handlers.Init_Redis_Mutex()
	graph_query.Graph_support_init(&site_data)
	data_handler.Data_handler_init(&site_data)
	irrigation_information_web.Start()
    data_base_cleanup.Start()
	for true {
		time.Sleep(time.Second * 60)
		//fmt.Println("polling loop")

	}

}
