package data_handler

import "fmt"
import "context"

import "time"
//import "reflect"
//import "encoding/json"
import "strconv"
import "hash/fnv"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "github.com/go-redis/redis/v8"
import "lacima.com/server_libraries/postgres"

var site       string
var redis_ip   string
var site_ptr *map[string]interface{}
var ctx    = context.TODO()
var client *redis.Client
var constructor_table = make(map[string]interface{})

func Data_handler_init( site_data *map[string]interface{}){

   site_ptr     = site_data
   site         = (*site_data)["site"].(string)
   redis_ip     = (*site_data)["host"].(string)
   create_redis_data_handle()
   create_constructors(&constructor_table)

}

func Remove_key(key string ){
 
    client.Del(ctx,key)
}


func Get_data_keys()[]string{
    
    keys,_ := client.Keys(ctx,"*").Result()   
    return keys
    
}    

func Store_Valid_Set(dict_key string,valid_set map[string]string){
    
  for key,value := range valid_set {
      
     client.HSet(ctx,dict_key,key,value)   
      
      
  }
    

}

func get_data_base_number()int{
 
   
    search_list := []string{"SYSTEM"}
    nodes := graph_query.Full_site_serach(&search_list)
    fmt.Println("nodes",nodes)
    node := nodes[0]
    
    data_db,err :=strconv.Atoi(node["data_db"])
    if err != nil {
      panic("bad data_db")
    }
    return data_db
    
}

func create_redis_data_handle(){
  
	site = (*site_ptr)["site"].(string)
    var address =  (*site_ptr)["host"].(string)
    var port = 	int((*site_ptr)["port"].(float64))
	var address_port = address+":"+strconv.Itoa(port)
	
    data_db := get_data_base_number()
	
	client = redis.NewClient(&redis.Options{
                                                 Addr: address_port,
												 ReadTimeout : time.Second*30,
												 WriteTimeout : time.Second*30,
												 DB: data_db,
                                               })
	err := client.Ping(ctx).Err();
	if err != nil{
	         panic("redis data connect")
	  }
    
}	

func create_constructors( constructor_table *map[string]interface{}){

      (*constructor_table)["HASH" ] ="HASH"
	  (*constructor_table)["STREAM_REDIS" ] =redis_handlers.Construct_Redis_Stream
	  (*constructor_table)["JOB_QUEUE" ] ="JOB_QUEUE"
	  
	  /* handlers to be implemented
       ["SINGLE_ELEMENT" ]
       ["MANAGED_HASH"]    
       ["RPC_SERVER"]     
       ["RPC_CLIENT"] 
      */
}    

  

func Construct_Data_Structures(  search_list *[]string )  *map[string]interface{}{

   fmt.Println("$$$$$$$$$$$$$$",search_list)
   handlers             := make(map[string]interface{})
   handler_definitions  := make([]map[string]interface{},0)
   construct_handler_definitions( search_list , &handler_definitions  ) 
   

   construct_redis_handlers( &handler_definitions, &handlers )
   
   return &handlers
} 


/*
func Construct_Multiple_Data_Structures(  search_list *[]string )[]*map[string]interface{}{
    
    packages   := graph_query.Common_package_search(&site,search_list)
    handlers       := make([]*map[string]interface{},0)
    for _,package_name := range packages {
        
        handler_items := make(map[string]interface{})
        handler_definitions  := make([]map[string]interface{},0)
        data_structures_json := package_name["data_structures"]
        construct_handler_definitions( data_structures_json , &handler_definitions)
        construct_redis_handlers(&handler_definitions, &handler_items)
        handlers = append(handlers,&handler_items)
        
    }
    return handlers
}
*/

func construct_handler_definitions( search_list *[]string, handler_definitions *[]map[string]interface{} )   {

   
   
   var packages = graph_query.Common_package_search(&site,search_list)
   if len(packages) != 1 {
       panic("bad package length "+strconv.Itoa(len(packages)))
   }
   
   var data_structures_json = packages[0]["data_structures"]
   
   
   
  
   var data_structures = graph_query.Convert_json_dictionary_interface(  data_structures_json)
   
   var namespace_json = packages[0]["namespace"]
   var namespace = graph_query.Convert_json_string( namespace_json)

   
   
   
   
   for _,v := range data_structures{
     var k = v.(map[string]interface{}) 
     var key = namespace +"["+k["type"].(string)+":"+k["name"].(string) +"]"
	 k["key"]= key
	 k["client"] = client
	 *handler_definitions = append(*handler_definitions,k)
   }
}  


 
func construct_redis_handlers( handler_definitions *[]map[string]interface{}, handlers *map[string]interface{} ){
   var type_def        string
   var name            string
   var user            string
   var password        string
   var key             string
   var depth           int64
   var timeout         int64
   var database_name   string
   var table_name      string
   var time_limit      int64
   var pg_stream       pg_drv.Postgres_Stream_Driver
   var pg_registry     pg_drv.Registry_Driver
   var pg_table        pg_drv.Postgres_Table_Driver
   var pg_float        pg_drv.Postgres_Float_Driver
   var pg_json        pg_drv.Json_Table_Driver
   for _,v := range *handler_definitions {
      type_def = v["type"].(string)
	  
	  
	  if type_def == "STREAM_REDIS" {
	     key = v["key"].(string)
		 name = v["name"].(string)
		 depth = int64(v["depth"].(float64))
		 (*handlers)[name] = redis_handlers.Construct_Redis_Stream(ctx,client,key,depth)
		  
		 

      }else if type_def == "HASH" {
	      key = v["key"].(string)
		  name = v["name"].(string)
		 (*handlers)[name] = redis_handlers.Construct_Redis_Hash(  ctx, client, key )
	     // add test code
		 
	  }else  if type_def == "JOB_QUEUE" {
	     key = v["key"].(string)
		 name = v["name"].(string)
		 depth = int64(v["depth"].(float64))
		 (*handlers)[name] = redis_handlers.Construct_Redis_Job_Queue(  ctx ,client, key , depth  )
		 
		 //var x redis_handlers.Redis_Stream_Struct
		 //x =(*handlers)[name].(redis_handlers.Redis_Stream_Struct)

	  
	   } else if type_def == "SINGLE_ELEMENT" {
	       key = v["key"].(string)
		   name = v["name"].(string)
		   (*handlers)[name] = redis_handlers.Construct_Redis_Single(  ctx ,client, key   )
	   
	   }else if type_def == "RPC_SERVER" {
	       key = v["key"].(string)
		   name = v["name"].(string)
		   timeout = int64(v["timeout"].(float64))
		   depth = int64(v["depth"].(float64))
		   (*handlers)[name] = redis_handlers.Construct_Redis_RPC(  ctx , client , key , timeout, depth  )
	   
	   } else if type_def == "POSTGRES_STREAM" {
          
		   key            = v["key"].(string)
           
           name          =   v["name"].(string)  
           user          =   v["user"] .(string) 
           password      =   v["password"].(string)  
           database_name =   v["database_name"].(string) 
           table_name    =   "T"+generate_table_name(key)
           time_limit    =   int64(v["time_limit"].(float64))
    
           
           
		   pg_stream = pg_drv.Construct_Postgres_Stream_Driver( key,user,password,database_name,table_name, time_limit) 
	       pg_stream.Connect(redis_ip)
           (*handlers)[name] = pg_stream
	   
	   } else if type_def == "POSTGRES_Registry" {
          
		   key            = v["key"].(string)
           
           name          =   v["name"].(string)  
           user          =   v["user"] .(string) 
           password      =   v["password"].(string)  
           database_name =   v["database_name"].(string) 
           table_name    =   "T"+generate_table_name(key)
           
		   pg_registry = pg_drv. Construct_Registry_Driver( key,user,password,database_name, table_name ) 
	       pg_registry.Connect(redis_ip)
           (*handlers)[name] = pg_registry	   
	   } else if type_def == "POSTGRES_TABLE" {
          
		   key            = v["key"].(string)
           
           name          =   v["name"].(string)  
           user          =   v["user"] .(string) 
           password      =   v["password"].(string)  
           database_name =   v["database_name"].(string) 
           table_name    =   "T"+generate_table_name(key)
           
		   pg_table = pg_drv. Construct_Postgres_Table_Driver( key,user,password,database_name, table_name ) 
	       pg_table.Connect(redis_ip)
           (*handlers)[name] = pg_table	   
	   } else if type_def == "POSTGRES_FLOAT" {
        
		   key            = v["key"].(string)
           
           name          =   v["name"].(string)  
           user          =   v["user"] .(string) 
           password      =   v["password"].(string)  
           database_name =   v["database_name"].(string) 
           table_name    =   "T"+generate_table_name(key)
           
		   pg_float = pg_drv. Construct_Postgres_Float_Driver( key,user,password,database_name, table_name ) 
	       pg_float.Connect(redis_ip)
           (*handlers)[name] = pg_float	   
     
        } else if type_def == "POSTGRES_JSON" {

		   key                          = v["key"].(string)           
           name                         =   v["name"].(string)  
           user                             =   v["user"] .(string) 
           password                =   v["password"].(string)  
           database_name    =   v["database_name"].(string) 
           table_name            =   "T"+generate_table_name(key)
		   pg_json = pg_drv. Construct_Json_Table_Driver( key,user,password,database_name, table_name ) 
	       pg_json.Connect(redis_ip)
           (*handlers)[name] = pg_json	   
	       
        }else if type_def == "ZSET_REDIS" {

        key = v["key"].(string)
		 name = v["name"].(string)
		
		 (*handlers)[name] = redis_handlers.Construct_Redis_ZSet(  ctx ,client, key   )
	       
 
      }else {
	   panic("Key is not expected "+type_def)
	 }
   }
	 
}

func generate_table_name( key string)string{
    h := fnv.New64a()
    h.Write([]byte(key))
    hash :=  h.Sum64()
   
    temp := int64(hash)
    if temp < 0 {
        temp = -temp
    }
   
    return strconv.FormatInt(int64(temp), 10)     
}

func convert_string_array_interface( input []interface{})[]string {
    return_value := make([]string,0)
    for _,i := range input {
        return_value = append(return_value,i.(string))
    }
    return return_value   
}   


