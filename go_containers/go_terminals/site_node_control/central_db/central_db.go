package central_db

import (
    "strconv"
    "os"
    "lacima.com/go_terminals/site_node_control/site_init"
    "lacima.com/go_terminals/docker_control"
    "lacima.com/redis_support/generate_handlers"
     "lacima.com/redis_support/redis_handlers"
     "lacima.com/redis_support/graph_query"
)

var Connection_state int  
/*
 * 0  -- not connected
 * 1  -- redis standalone
 * 2  -- slave standalone
 * 3  -- site controller active
 */


var Redis_Up  bool
var Graph_db  float64
var Site_data map[string]interface{}

func Fill_in_slave_data(){

  /*
   * Minimium information to connect to event broker and event registry
   */
  
  Site_data = make(map[string]interface{})
  Site_data["master_flag"]  =false
  Site_data["site"]  = os.Getenv("site")
  Site_data["local_node"]  = os.Getenv("local_node")
  
  // ip of the redis server
   port,_               := strconv.Atoi(os.Getenv("port"))
  Site_data["port"]     = float64(port)
  Site_data["host"]     =   os.Getenv("host")
  
  Site_data["graph_container_image"]     = os.Getenv("graph_container_image")
  Site_data["graph_container_script"]    = os.Getenv("graph_container_script")		
  Site_data["redis_container_name"]      = os.Getenv("redis_container_name")
  Site_data["redis_container_image"]     = os.Getenv("redis_container_image")
  Site_data["redis_start_script"]        = os.Getenv("redis_start_script")
  Graph_db,_                              := strconv.ParseFloat(os.Getenv("graph_db"),64)
  Site_data["graph_db"] = Graph_db
  
}




func Generate_title( starting_title string)string{
    switch Connection_state{
        case 0:
            starting_title += " Redis db not connected "
            
        case 1:
            starting_title += " Stand Alone Master "
            
        case 2:
            starting_title += " Stand Alone Slave "
        case 3:
            starting_title += " Connected to Site Controller "
        default:
            panic("bad state")
    }
    return starting_title
}

func convert_float( input float64 )int{
    return int(input)
}



func Test_for_redis_connection()bool{
    port := convert_float(Site_data["port"].(float64))
    if site_init.Test_redis_connection( Site_data["host"].(string),port )  == true{
        return true
    }
    return false
}

func Do_master_setup(){
    
    site_init.Site_Master_Init(&Site_data)
    
    
}

func Do_slave_setup(){
    Fill_in_slave_data()
    site_init.Site_Slave_Init(&Site_data)
    
    
}


func Reload_graphic_db(){
    
 docker_control.Container_Run(Site_data["graph_container_script"].(string))   
    
}

func Setup_Structures(){
   redis_handlers.Init_Redis_Mutex()
   graph_query.Graph_support_init(&Site_data)
   data_handler.Data_handler_init(&Site_data)
}
