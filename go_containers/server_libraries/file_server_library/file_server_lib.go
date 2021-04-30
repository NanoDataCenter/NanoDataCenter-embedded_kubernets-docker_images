package file_server_lib



//import "fmt"

import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"


type File_Server_Client_Type struct{

   driver redis_handlers.Redis_RPC_Struct
}




func File_Server_Init(search_list *[]string)File_Server_Client_Type{

  var return_value File_Server_Client_Type
  var handlers = data_handler.Construct_Data_Structures(search_list)  
  return_value.driver = (*handlers)["FILE_SERVER_RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
  return return_value
}  

func (v File_Server_Client_Type)Ping()bool{
  /*

       var parameters = make(map[string]interface{})
       result := msgpack_utils.Convert_rpc_return(v.driver.Send_rpc_message( "ping", &parameters ))
     
       //fmt.Println("ping",result)
       //fmt.Println(	(*result)["status"].(bool))   
       return result["status"].(bool)
*/
   return true
}

func (v File_Server_Client_Type)Read_file(file_name string)(string,bool) {
  

       var parameters = make(map[string]interface{})
       
       parameters["file_name"] = file_name
       result := v.driver.Send_json_rpc_message( "read", parameters ) 
	   status := result["status"].(bool)
	   file_data := result["results"].(string)

	   
       return file_data,status

}

func (v File_Server_Client_Type)Write_file(file_name,data string)bool{
/*     
	   var parameters = make(map[string]interface{})
       parameters["file_name"] = file_name
	   parameters["data"] = data
       result := msgpack_utils.Convert_rpc_return(v.driver.Send_rpc_message( "write", &parameters ) ) 
	
       return result["status"].(bool)
*/
       return true
}

func (v File_Server_Client_Type)Delete_file(file_name string)bool {
/*
       var parameters = make(map[string]interface{})
       parameters["file_name"] = file_name
	   result := msgpack_utils.Convert_rpc_return(v.driver.Send_rpc_message( "delete_file",&parameters))
	  
       return result["status"].(bool)
*/
       return true  
}

func (v File_Server_Client_Type)File_exists(file_name string)( bool,bool) {
/*      
	   var parameters = make(map[string]interface{})
       parameters["file_name"] = file_name
	   result := msgpack_utils.Convert_rpc_return(v.driver.Send_rpc_message( "file_exists",&parameters))
	
       return result["status"].(bool),result["directory"].(bool)
*/
      return true, true
}
        
func (v File_Server_Client_Type)File_directory(path string)([]string , bool){
/*  
       var parameters = make(map[string]interface{})
       parameters["path"] = path
	   
      return_value_interface := msgpack_utils.Convert_rpc_return(v.driver.Send_rpc_message( "file_directory" ,&parameters))
	  
       status := return_value_interface["status"].(bool)
	   var results_map = return_value_interface["results"].([]interface{})
	   
	   var results =  make([]string,0)
       for _,item := range results_map{
	      results = append(results, string(item.([]byte)))
	   }
	   
	   return  results , status
*/
     return []string{},false
}           


   
func  (v File_Server_Client_Type)Mkdir(path string)bool{
 /*     
	   var parameters = make(map[string]interface{})
       parameters["path"] = path
	   result := msgpack_utils.Convert_rpc_return(v.driver.Send_rpc_message( "make_dir",&parameters))
	   
       return result["status"].(bool)
*/
   return true
 }





