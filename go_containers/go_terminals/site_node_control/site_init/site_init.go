
package site_init

import ( 

"fmt"
"net"
"strings"
"context"
"time"
"strconv"
"encoding/json"
"io/ioutil"


"lacima.com/redis_support/generate_handlers"
"lacima.com/redis_support/redis_handlers"
"lacima.com/site_control_app/docker_control"
"lacima.com/redis_support/graph_query"
"lacima.com/Patterns/logging_support"
"github.com/go-redis/redis/v8"
"lacima.com/Patterns/msgpack_2"
"/site_node_control/

)




//type Database struct {
//   Client *redis.Client
//}
var ctx    = context.TODO()

var redis_container_name string
var services_json string
//var container []string
var containers = make([]string,0)
var startup_containers = make([]string,0)
var site_run_once_containers = make([]map[string]string,0)
var site_containers = make([]map[string]string,0)


 
func test_redis_connection( address string, port int )bool{
    
    address_port  := address+":"+strconv.Itoa(port)

    
   client := redis.NewClient(&redis.Options{
                                              Addr: address_port,
                                              DB: 0,
                                        })
     err := client.Ping(ctx).Err()
     client.Close() 
     if err != nil{
		  return false
     }
     return true
}
		   
    
    



func get_store_site_data(site_data *map[string]interface{}){
    
    
    site_data_package := data_handler.Construct_Data_Structures(&[]string{"ENVIRONMENTAL_VARIABLES"})
    site_data_driver  := (*site_data_package)["ENVIRONMENTAL_VARIABLES"].(redis_handlers.Redis_Single_Structure)
    json_data         := []byte(site_data_driver.Get())
    // test json data
    var err1 = json.Unmarshal(json_data,&site_data)
    if err1 != nil{
	 panic("bad json site data")
	}
	site := (*site_data)["site"].(string)
    search_path := []string{"SITE:"+site}
	site_nodes := graph_query.Common_qs_search(&search_path)
	
    site_node := site_nodes[0]
    config_file := graph_query.Convert_json_string(site_node["file_path"])
   
    err := ioutil.WriteFile(config_file, json_data, 0644)
    
    if err != nil{
        fmt.Println(err)
        panic("bad  file write")
    }
   
    
}
					 
						 

func Site_Slave_Init(site_data *map[string]interface{} )bool{
    
 
    status := test_redis_connection(((*site_data)["host"].(string), int((*site_data)["port"].(float64)))
    central.Redis_Up = status
    if status == false {
        
        return false
    
    }else{
      
      graph_query.Graph_support_init(site_data)  // only start containers that are not running
      data_handler.Data_handler_init(site_data) 
      get_store_site_data(site_data) 
    }
    return true
    
 
    
}

 

    




