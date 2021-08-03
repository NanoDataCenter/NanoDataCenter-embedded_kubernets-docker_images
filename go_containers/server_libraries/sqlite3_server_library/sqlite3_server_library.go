package sqlite3_server_lib


//import "fmt"
//import "reflect"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"


type Sqlite3_Server_Client_Type struct{

   driver redis_handlers.Redis_RPC_Struct
}

func Sqlite3_Server_Init(search_list *[]string)Sqlite3_Server_Client_Type{

  var return_value Sqlite3_Server_Client_Type
  handlers := data_handler.Construct_Data_Structures(search_list)  
  return_value.driver = (*handlers)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
  return return_value
}  


func (v Sqlite3_Server_Client_Type)Ping()bool{
  

       parameters := make(map[string]interface{})
       
       result := v.driver.Send_json_rpc_message( "ping", parameters ) 
       return result["status"].(bool)

}


func (v Sqlite3_Server_Client_Type)List_databases()(bool , []string) {
  
       return_value := make([]string,0)
       parameters := make(map[string]interface{})
       
       result := v.driver.Send_json_rpc_message( "list_data_bases", parameters ) 
       
       
       for _,element := range result["results"].([]interface{}){
           
        return_value = append(return_value,element.(string))   
       }
       return result["status"].(bool), return_value

}

func (v Sqlite3_Server_Client_Type)Open_database(db_name string )bool{
  

       parameters := make(map[string]interface{})
       parameters["database"] = db_name
       result := v.driver.Send_json_rpc_message( "open_database", parameters ) 
       return result["status"].(bool)

}

func (v Sqlite3_Server_Client_Type)Close_database(db_name string)bool{
  

       parameters := make(map[string]interface{})
       parameters["database"] = db_name
       result := v.driver.Send_json_rpc_message( "close_database", parameters ) 
       return result["status"].(bool)

}

func (v Sqlite3_Server_Client_Type)Delete_database( db_name string )bool{
  

       parameters := make(map[string]interface{})
       parameters["database"] = db_name
       result := v.driver.Send_json_rpc_message("delete_database", parameters ) 
       return result["status"].(bool)

}



func (v Sqlite3_Server_Client_Type)Vacuum( db_name string )bool{
  

       parameters := make(map[string]interface{})
       parameters["database"] = db_name
       result := v.driver.Send_json_rpc_message("vacuum", parameters ) 
       return result["status"].(bool)

}

func (v Sqlite3_Server_Client_Type)Execute( db_name string, script string )bool{
  

       parameters := make(map[string]interface{})
       parameters["script"] = script
       parameters["database"] = db_name
       result := v.driver.Send_json_rpc_message("execute", parameters ) 
       return result["status"].(bool)

}

func (v Sqlite3_Server_Client_Type)Query( db_name string, script string)(bool,map[string]interface{}){
  

       parameters := make(map[string]interface{})
       parameters["database"]  = db_name   
       parameters["script"]    = script      
       result := v.driver.Send_json_rpc_message("query", parameters )
       
       return result["status"].(bool),result["results"].(map[string]interface{})

}
func (v Sqlite3_Server_Client_Type)List_tables( db_name string )(bool, []string){
  

       parameters := make(map[string]interface{})
       parameters["database"]  = db_name
       result := v.driver.Send_json_rpc_message("list_tables", parameters )
       database_list := parameters["results"].([]string)
       return result["status"].(bool),database_list

}




/*

    
   def filter_result(self, return_value):
       if return_value[0] == True:
           return return_value[1]
       else:
           raise ValueError(return_value[1])
  
   def list_data_bases(self):
       print("list_data_bases")
       parameters = {}
       return_value = self.rpc_client.send_rpc_message( method="list_list_data_bases",parameters=parameters,timeout=3 )
       print("return_value",return_value)
       return self.filter_result(return_value)       
       
   def create_database(self,database):
       parameters = {}
       parameters["database"] = database
       return_value = self.rpc_client.send_rpc_message( method="create_database",parameters=parameters ,timeout=3 )
       return self.filter_result(return_value)       
           
   def delete_database(self,database):
       parameters = {}
       parameters["database"] = database      
       return_value = self.rpc_client.send_rpc_message( method="delete_database", parameters=parameters ,timeout=3 )
       return self.filter_result(return_value)   
 
 
   def close_database(self,database):
       parameters = {}
       parameters["database"] = database
       return_value = self.rpc_client.send_rpc_message( method="close_database",parameters=parameters ,timeout=3 )
       return self.filter_result(return_value)              


            

   def vacuum(self,database):
       parameters = {}
       parameters["database"] = database
       return_value = self.rpc_client.send_rpc_message( method="vacuum",parameters=parameters ,timeout=3 )
       return self.filter_result(return_value)       

       
   def version(self):
       parameters = {}
       return_value = self.rpc_client.send_rpc_message( method="version",parameters=parameters ,timeout=3 )
       return self.filter_result(return_value)  
       
   def set_text(self,database,text_state):
       parameters = {}
       parameters["database"] = database  
       parameters["text_state"] = text_state       
       return_value = self.rpc_client.send_rpc_message( method="set_txt",parameters=parameters ,timeout=3)
       return self.filter_result(return_value)

       
   def get_text(self,database):
       parameters = {}
       parameters["database"] = database
       return_value = self.rpc_client.send_rpc_message( method="get_txt",parameters=parameters ,timeout=3 )
       return self.filter_result(return_value)    
       
   def backup(self,database,backup_db,pages= 0):
       parameters = {}
       parameters["database"] = database
       parameters["backup_db"] = backup_db
       parameters["pages"] = pages
       return_value = self.rpc_client.send_rpc_message( method="backup",parameters=parameters ,timeout=3 )
       return self.filter_result(return_value) 
       
   def ex_exec(self,database,script):
       parameters = {}
       parameters["database"] = database
       parameters["script"] = script
       return_value = self.rpc_client.send_rpc_message( method="execute",parameters=parameters ,timeout=3 )
       return self.filter_result(return_value)    
       
   def ex_script(self,database,script):
       parameters = {}
       parameters["database"] = database
       parameters["script"] = script
       return_value = self.rpc_client.send_rpc_message( method="execute_script",parameters=parameters ,timeout=3)
       return self.filter_result(return_value)              
       
   def commit(self,database):
       parameters = {}
       parameters["database"] = database
       return_value = self.rpc_client.send_rpc_message( method="commit",parameters=parameters ,timeout=3 )
       return self.filter_result(return_value)    
       
   def select(self,database,script):
       parameters = {}
       parameters["database"] = database
       parameters["script"] = script
       return_value = self.rpc_client.send_rpc_message( method="select",parameters=parameters ,timeout=3 )
       return self.filter_result(return_value)    
   
   
 
 
 */      
