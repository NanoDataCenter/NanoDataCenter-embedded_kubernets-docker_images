package central_db

import "strconv"
import "os"


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
  Graph_db,_                             := strconv.Atoi(os.Getenv("graph_db"))
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



