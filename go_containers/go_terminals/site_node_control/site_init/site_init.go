
package site_init

import ( 

"fmt"
"net"
"strings"
"context"
"time"
"strconv"
"encoding/json"
//"os"
//"io/ioutil"

"lacima.com/go_terminals/docker_control"
"lacima.com/redis_support/generate_handlers"
"lacima.com/redis_support/redis_handlers"
"lacima.com/redis_support/graph_query"
//"lacima.com/Patterns/logging_support"
"github.com/go-redis/redis/v8"
//"lacima.com/Patterns/msgpack_2"
//"/site_node_control/
 "lacima.com/Patterns/shell_utils" 
 "io/ioutil"
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
var node_containers = make([]map[string]string,9)

 
func Test_redis_connection( address string, port int )bool{
    
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
		   
    
func Site_Master_Init(  site_data *map[string]interface{} ){ 
    
      
      redis_handlers.Init_Redis_Mutex()

	
       
      mount_usb_drive() // mount external hard drive for storing system data
      //fill_in_site_data()
       
	  graph_container_image   := (*site_data)["graph_container_image"].(string)
      graph_container_script  := (*site_data)["graph_container_script"].(string)
      redis_container_name    = (*site_data)["redis_container_name"].(string)	 		
	  redis_container_image   := (*site_data)["redis_container_image"].(string)	 					 
      redis_startup_script    := (*site_data)["redis_start_script"].(string)		
      

      
      docker_control.Stop_Running_Containters()
     
      docker_control.Remove_All_Containers()
      if docker_control.Image_Exists(redis_container_image)== false {
          panic("container image should exit")
          
	  }
      docker_control.Container_up(redis_container_name,redis_startup_script)
      
	  time.Sleep(time.Second*4)
	  if docker_control.Container_is_running(redis_container_name) == false{
	     panic("redis container did not start")
	  }
      
      wait_for_redis_connection((*site_data)["host"].(string), int((*site_data)["port"].(float64)))
      
   
      if docker_control.Image_Exists(graph_container_image)== false {
          panic("container image should exit")
          
	  }
	 
      docker_control.Container_Run(graph_container_script)
	  
      
   
      graph_query.Graph_support_init(site_data)
      data_handler.Data_handler_init(site_data)
      store_site_data(site_data)
     
 
      
      // remove obselete data Construct_Data_Structures
      
      remove_obsolete_data_structures()
     
      determine_master_containers((*site_data)["site"].(string))
      
      
      find_master_containers_data()
      find_master_run_once_containers()
     
      start_run_once_containers()

      start_master_containers()
      
      
      
     time.Sleep(time.Second*5) // allow containers to startup_command
     verify_master_containers()
     reboot_flag_data_structures := data_handler.Construct_Data_Structures(&[]string{"REBOOT_FLAG"})
     reboot_flag_driver := (*reboot_flag_data_structures)["REBOOT_FLAG"].(redis_handlers.Redis_Single_Structure)
     reboot_flag_driver.Set("NOT_ACTIVE")
   
     
     ip_table    := data_handler.Construct_Data_Structures(&[]string{"NODE_MAP"})
     ip_driver   := (*ip_table)["NODE_MAP"].(redis_handlers.Redis_Hash_Struct)
     ip_address  := find_local_address()
     ip_driver.Delete_All()
     ip_driver.HSet((*site_data)["local_node"].(string),ip_address )    
     ip_driver.HSet("master",ip_address ) 
   
   

     docker_table    := data_handler.Construct_Data_Structures(&[]string{"DOCKER_CONTROL"})
     docker_driver   := (*docker_table)["DOCKER_DISPLAY_DICTIONARY"].(redis_handlers.Redis_Hash_Struct)
     docker_driver.Delete_All()

}	



      
func handle_mount_panic() {
  
    if a := recover(); a != nil {
        ;//fmt.Println("RECOVER", a)
    }
}

func mount_usb_drive(){
  defer handle_mount_panic()
  shell_utils.System_mshell("mount /dev/sda /home/pi/mountpoint")

}      
      
      
    
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
		   
    
    



func wait_for_redis_connection(address string, port int ) {
   
   
   var loop_flag = true
   for loop_flag == true {
      if test_redis_connection(address,port ) == true {
          return
      }
      time.Sleep(time.Second)
    
   }		
   
}

func store_site_data(site_data *map[string]interface{}){


   
    json_data, _ := json.MarshalIndent(site_data,"","")
    site_data_package := data_handler.Construct_Data_Structures(&[]string{"ENVIRONMENTAL_VARIABLES"})
    site_data_driver := (*site_data_package)["ENVIRONMENTAL_VARIABLES"].(redis_handlers.Redis_Single_Structure)
    site_data_driver.Set(string(json_data))

}

func remove_obsolete_data_structures(){
 
    current_keys      := data_handler.Get_data_keys()
    valid_set        := graph_query.Get_valid_keys()
    valid_set["data_set"] = "key_dictionary"
    data_handler.Store_Valid_Set("data_set",valid_set)
    for _,key := range current_keys{
        if _,ok := valid_set[key]; ok == false {
            
            data_handler.Remove_key(key)
        }
    }
            
}

func verify_master_containers() {

 for _,value := range site_containers{
   if docker_control.Container_is_running(value["name"] )== false{
     panic("container "+value["name"]+"is not running")
   }
 }
}

func find_local_address()string{
    
   conn, error := net.Dial("udp", "8.8.8.8:80")  
   if error != nil {  
     panic("bad ip")
  
    }  
  
    defer conn.Close()  
    ipAddress_port := conn.LocalAddr().(*net.UDPAddr).String()
    temp := strings.Split(ipAddress_port,":")
    ip_address := temp[0]
  
    return ip_address
}     
	
	
func start_run_once_containers(){
  //fmt.Println("run once ",site_run_once_containers)
  for _,value := range site_run_once_containers{
    if value["name"] == redis_container_name {
	  
	  continue
	}

	if docker_control.Image_Exists(value["container_image"]) == false{
	   
	   
	   panic("container image should exit")
	}
	
	docker_control.Container_rm(value["name"])
    
	docker_control.Container_up(value["name"],value["startup_command"])
  }
}

func start_master_containers(){
  for _,value := range site_containers{
    if value["name"] == redis_container_name {
	 
	  continue
	}
	if docker_control.Image_Exists(value["container_image"]) == false{
	   
	   panic("container image should exit")
	}
	
	docker_control.Container_rm(value["name"])
	docker_control.Container_up(value["name"],value["startup_command"])
  }
}  
	
 
func find_master_run_once_containers(){
    
    for _,service := range startup_containers{
         
	     var item = make(map[string]string,0)
	     var search_list = []string{"CONTAINER"+":"+service}
		 var container_nodes = graph_query.Common_qs_search(&search_list)
         var container_node = container_nodes[0]
		 
		 item["name"] =  graph_query.Convert_json_string(container_node["name"])
 		 item["container_image"] =graph_query.Convert_json_string(container_node["container_image"])
		 item["startup_command"] = graph_query.Convert_json_string(container_node["startup_command"])

		 site_run_once_containers = append(site_run_once_containers,item)
		 
	}
	
	
} 
 

func find_master_containers_data(){
    
    for _,service := range containers{
         
	     var item = make(map[string]string,0)
	     var search_list = []string{"CONTAINER"+":"+service}
		 var container_nodes = graph_query.Common_qs_search(&search_list)
         var container_node = container_nodes[0]
		
		 item["name"] =  graph_query.Convert_json_string(container_node["name"])
 		 item["container_image"] =graph_query.Convert_json_string(container_node["container_image"])
		 item["startup_command"] = graph_query.Convert_json_string(container_node["startup_command"])

		 site_containers = append(site_containers,item)
		 
	}
	
	
}

func  determine_master_containers(site string){
    search_list := []string{ "SITE:"+site }
   
    site_nodes := graph_query.Common_qs_search(&search_list)
    site_node  := site_nodes[0]
    
    
    containers               = graph_query.Convert_json_string_array(	site_node["containers"] )
    startup_containers       = graph_query.Convert_json_string_array(	site_node["startup_containers"] )
    //fmt.Println("containers, startup_containers",containers,startup_containers)
   
}    





func Site_Slave_Init(site_data *map[string]interface{} ){
    
 
 
    
    graph_query.Graph_support_init(site_data)
    data_handler.Data_handler_init(site_data)
    store_site_data(site_data)
    
   
    
 
    get_store_site_data(site_data)
    ip_table    := data_handler.Construct_Data_Structures(&[]string{"NODE_MAP"})
    ip_driver   := (*ip_table)["NODE_MAP"].(redis_handlers.Redis_Hash_Struct)
    ip_address  := find_local_address()
    ip_driver.HSet((*site_data)["local_node"].(string),ip_address )    
      
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
