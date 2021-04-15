package file_server_lib





import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"

var driver redis_handlers.Redis_RPC_Struct

func File_Server_Init(){

  var search_list = []string{"FILE_SERVER","FILE_SERVER"}
  var handlers = data_handler.Construct_Data_Structures(&search_list)  
  driver = (*handlers)["FILE_SERVER_RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)

}  

func ping(path,file_name string)bool{
  

       var parameters = make(map[string]interface{})
       var result = (driver).Send_rpc_message( "ping", parameters )  
       return (*result)["status"].(bool)
}

func load_file(path,file_name string)(bool,string){
  

       var parameters = make(map[string]interface{})
       parameters["path"] = path
       parameters["file_name"] = file_name
       var result = (driver).Send_rpc_message( "load", parameters )  
       return (*result)["status"].(bool),(*result)["result"].(string)
}

func save_file(path,file_name,data string)bool{
       /*
       parameters = {}
       parameters["path"] = path
       parameters["file_name"] = file_name
       parameters["data"] = data
       return_value = self.rpc_client.send_rpc_message( method="save",parameters=parameters,timeout=3 )
       return self.check_file(file_name,return_value)
       */
	   return false
}
func file_exists(path,file_name string) bool {
      
	   /*
       parameters = {}
       parameters["path"] = path
       parameters["file_name"] = file_name
       return_value = self.rpc_client.send_rpc_message( method="file_exists",parameters=parameters,timeout=3 )
       return return_value[0]
	   */
	   return false
}
        
func file_directory(path,file string){}string {
       /*     
       parameters = {}
       parameters["path"] = path
       return_value = self.rpc_client.send_rpc_message( method="file_directory",parameters=parameters,timeout=3 )
       return self.check_file(path,return_value)
	   */
}           

func delete_file(self, path,file_name)bool {
       
       parameters = {}
       parameters["path"] = path
       parameters["file_name"] = file_name
       return_value = self.rpc_client.send_rpc_message( method="delete_file",parameters=parameters,timeout=3 )
       return return_value[0]
}
   
func  mkdir(self,path,file)bool{
      
       parameters = {}
       parameters["path"] = path
       return_value = self.rpc_client.send_rpc_message( method="make_dir",parameters=parameters,timeout=3 )
       return return_value[0]
   
 }




*/