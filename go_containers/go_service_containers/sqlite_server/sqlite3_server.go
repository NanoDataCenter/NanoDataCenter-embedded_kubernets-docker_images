package main

import "fmt"

import "time"
import "os"
import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"
import  _ "lacima.com/go_service_containers/sqlite_server/go-sqlite3"
import	"database/sql"

/*
Note 

To import a package solely for its side-effects (initialization), use the blank identifier as explicit package name:

import _ "lib/math"

import	_ "lacima.com/go_service_containers/sqlite_server/go-sqlite3"

*/

var db_handlers  map[string]*sql.DB

func main(){

    
    
	var config_file ="/data/redis_configuration.json"
	
	var site_data_store map[string]interface{}

    
	site_data_store = get_site_data.Get_site_data(config_file)

    graph_query.Graph_support_init(&site_data_store)
	redis_handlers.Init_Redis_Mutex()
    
    
    db_handlers = make(map[string]*sql.DB)
   
	data_handler.Data_handler_init(&site_data_store)	
 	  
     search_list := []string{"RPC_SERVER:SQLITE3_SERVER","RPC_SERVER"}

     handlers := data_handler.Construct_Data_Structures(&search_list)
    
     driver := (*handlers)["RPC_SERVER"].(redis_handlers.Redis_RPC_Struct)
	
	
	
	
    driver.Add_handler( "list_data_bases",list_data_bases)
    driver.Add_handler( "open_database", open_database)
    driver.Add_handler( "delete_database",delete_database)
    driver.Add_handler("close_database",close_database)
    driver.Add_handler("vacuum",vacuum)
    driver.Add_handler("execute",ex_exec)
    
    driver.Add_handler("query",query)
    driver.Add_handler("list_tables",list_tables)
    
    
	driver.Json_Rpc_start()
	// code to keep the go rpc thread running
	for true {
	  //fmt.Println("main spining")
	  time.Sleep(time.Second*10)
	}	
   
}





func valid_database(db_name string)bool{
  if _,found := db_handlers[db_name]; found{
      //fmt.Println("db name",db_name,"  found")
      return true
  }
  //fmt.Println("db name",db_name,"  not found")
  return false
    
}

func list_data_bases(parameters map[string]interface{} ) map[string]interface{}{
     
     results := make([]string,len(db_handlers))
     for key,_ := range db_handlers {
         results = append(results,key)
         
     }
     //fmt.Println("results",results)
     parameters["status"] = true
     parameters["results"] = results
     return parameters

}


func open_database(parameters map[string]interface{} ) map[string]interface{}{

    parameters["results"] = ""
    parameters["status"] = false
    db_name := parameters["database"].(string)
    
    file_path := "/files/"+db_name+".db"
    if valid_database(db_name) == false {
        db, err := sql.Open("sqlite3", file_path)
        //fmt.Println("file_path",file_path)
        //fmt.Println("err",err)
        if err != nil {
            fmt.Println(err)
            return parameters
        }
        //fmt.Println("made it here")
        db_handlers[db_name] = db
        parameters["status"] = true
        return parameters
        
        
    }
    // all ready open
    parameters["status"] = true 
    return parameters
    
}

func close_database(parameters map[string]interface{} ) map[string]interface{}{
 
    parameters["results"] = ""
    db_name := parameters["database"].(string)
    if valid_database(db_name) == true{   
        parameters["status"] = true
        db_handlers[db_name].Close()
        delete(db_handlers, db_name )
        return parameters
    }       
    parameters["status"] = false
    
    return parameters
}


func delete_database(parameters map[string]interface{} ) map[string]interface{}{ 
   
     parameters["results"] = ""
     db_name := parameters["database"].(string)
     if valid_database(db_name) == false{
         file_path := "/sqlite/"+db_name+".db"
         err := os.Remove(file_path)
         if err != nil {
             fmt.Println(err)
             parameters["status"] = false
             return parameters
         }
         parameters["status"] = true
         return parameters

     }
    parameters["status"] = false
    return parameters 
}
           
func list_tables(parameters map[string]interface{} ) map[string]interface{}{
     fmt.Println("list table start")
     parameters["script"] = "SELECT name FROM sqlite_master WHERE type='table';"
     return query(parameters)
}


func vacuum(parameters map[string]interface{} ) map[string]interface{}{ 
     db_name := parameters["database"].(string)
     parameters["results"] = ""
     parameters["status"] = false
     if valid_database(db_name)==true{
          db := db_handlers[db_name]
          
          _,err := db.Exec("VACUUM")
        fmt.Println("vacuum",err)
         if err != nil {
              fmt.Println(err)
              return parameters
		      
	      }
	      parameters["status"] = true
          return parameters
     }
     
     fmt.Println("bad db")
     return parameters
}
       
       
  

func ex_exec(parameters map[string]interface{} ) map[string]interface{}{   

    script  := parameters["script"].(string) 
    db_name := parameters["database"].(string)
    
    parameters["results"] = ""
    parameters["status"] = false
    if valid_database(db_name) == false {
        return parameters
    }
    
    
     if valid_database(db_name) == false{
         
		 return parameters
     }       
      db      := db_handlers[db_name]
    _, err := db.Exec(script)
	if err != nil {
        fmt.Println(err)
		return parameters
	}
	parameters["status"] = true
	return parameters
}

           
          


func query(parameters map[string]interface{} ) map[string]interface{}{            
   
    db_name   :=  parameters["database"].(string)   
    script    :=  parameters["script"].(string)           
    db        := db_handlers[db_name]  
    fmt.Println("query start")
    parameters["results"] = ""
    parameters["status"] = false
    if valid_database(db_name) == false {
        return parameters
    }   
     rows, err := db.Query(script )
     if err != nil {
		
		return parameters
     }           

    cols, _ := rows.Columns()
	
    
    
    result := make([]map[string]interface{},0)
    for rows.Next(){
      columns := make([]interface{}, len(cols))
      columnPointers := make([]interface{}, len(cols))
      for i := range columns {
			columnPointers[i] = &columns[i]
       }

		if err := rows.Scan(columnPointers...); err != nil {
           
		     return parameters
		}

		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
        result = append(result,m)
    } 
    
    parameters["results"] = result
    parameters["status"] = true
    
    return parameters
}
           
 
