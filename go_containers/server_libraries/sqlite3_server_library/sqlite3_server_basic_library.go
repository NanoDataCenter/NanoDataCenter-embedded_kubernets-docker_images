package sqlite3_server_lib



import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"

type Sqlite_row_element []string

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




func (v Sqlite3_Server_Client_Type)Execute( db_name string, script string )bool{
  

       parameters := make(map[string]interface{})
       parameters["script"] = script
       parameters["database"] = db_name
       result := v.driver.Send_json_rpc_message("execute", parameters ) 
       return result["status"].(bool)

}

func (v Sqlite3_Server_Client_Type)Query( db_name string, script string)(bool,[]map[string]interface{}){
  

       parameters := make(map[string]interface{})
       parameters["database"]  = db_name   
       parameters["script"]    = script      
       parameters = v.driver.Send_json_rpc_message("query", parameters )
           
             
       return_array := make([]map[string]interface{},0)
       if parameters["status"].(bool) != true {
            
            return parameters["status"].(bool),return_array

       }
       
       for _,element := range parameters["results"].([]interface{}){
        value_map := element.(map[string]interface{})
       
        return_array = append(return_array,value_map)   
       }
      
       return parameters["status"].(bool),return_array

}

