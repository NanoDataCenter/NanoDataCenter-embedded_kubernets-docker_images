package node_control_server_lib



//import "fmt"

import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/graph_query"

type Node_Server_Client_Type struct{

   Driver_array map[string]redis_handlers.Redis_RPC_Struct
}



func find_processors()[]string {
    
  site_nodes := make([]string,0)
  nodes := graph_query.Common_qs_search(&[]string{"PROCESSOR"})
  for _,node := range nodes {
      temp := graph_query.Convert_json_string(node["name"])
      site_nodes = append(site_nodes,temp)
  }
  return site_nodes
    
}

 
    
    
func Node_Server_Init()Node_Server_Client_Type{

  processors := find_processors()
  
  var return_value Node_Server_Client_Type
  return_value.Driver_array = make(map[string]redis_handlers.Redis_RPC_Struct)
  for _,processor := range processors {
     
      temp := data_handler.Construct_Data_Structures(&[]string{"PROCESSOR:"+processor,"RPC_SERVER:NODE_CONTROL","RPC_SERVER"} )
      return_value.Driver_array[processor] = (*temp)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
      
  }
  
  return return_value
}  

func (v Node_Server_Client_Type)Ping(node string)bool{
  
       //fmt.Println("node",node)
       parameters := make(map[string]interface{})
       
       result := v.Driver_array[node].Send_json_rpc_message( "ping", parameters )
       //fmt.Println("result",result)
       if result != nil {
          return result["status"].(bool)
       }
       return false

}

func (v Node_Server_Client_Type)Reboot(node string){
  

       parameters := make(map[string]interface{})
       
       v.Driver_array[node].Post_json_rpc_message( "rebot", parameters ) 
      

}
