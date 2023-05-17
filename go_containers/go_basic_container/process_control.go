package main


/*

notes from erlang supervisor

  call back --- handle specialized recovery action
  number of retries  for a process
  reset time  -- recovery process must be active for a tbd time before reset is considered successful

  types of resets
  one for one  -- just start the one process
  one for all  --- reset all if any one reset
  reset for all  all  -- reset all that are following item in the list
  
  simple-one-for-all used in dynamically created processes
  


*/
import "fmt"
import "os"
import "time"
import "strconv"
import "context"
import "github.com/go-redis/redis/v8"
import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/cf_control"
import "lacima.com/go_basic_container/process_control_support"

var ctx    = context.TODO()
var  cf_control_cluster cf.CF_CLUSTER_TYPE
var site_data_store map[string]interface{}
const config_file = "/data/redis_configuration.json"

func main(){

    container_name := os.Getenv("CONTAINER_NAME")

    fmt.Println("container_name",container_name)
    
    // wait for redis connection
    site_data_store = get_site_data.Get_site_data(config_file)
	 address :=  site_data_store["host"].(string)
     port  := 	      int(site_data_store["port"].(float64))//float 64 because of json
    
 
    wait_for_redis_connection(address, port  )

    
    
    graph_query.Graph_support_init(&site_data_store)
	redis_handlers.Init_Redis_Mutex()
	data_handler.Data_handler_init(&site_data_store)
	




    system_ctrl := system_control.Construct_System_Control(   container_name )

	cf_control_cluster.Cf_cluster_init()
	cf_control_cluster.Cf_set_current_row("container_process_monitor")
	
    (system_ctrl).Init(&cf_control_cluster)
	
	(cf_control_cluster).CF_Fork()	
	
    loop_flag := true
    for loop_flag{
      time.Sleep(time.Second*100)
      //fmt.Println("main is spinning")
    } 

	
	
    

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


func test_redis_connection( address string, port int )bool{
    
    address_port  := address+":"+strconv.Itoa(port)
   //fmt.Println("address port",address_port)
    
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

