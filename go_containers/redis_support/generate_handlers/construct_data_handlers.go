package data_handler

import "context"
import "fmt"
import "time"
//import "reflect"
import "strconv"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "github.com/go-redis/redis/v8"

var site string
var site_ptr *map[string]interface{}
var ctx    = context.TODO()
var client *redis.Client
var constructor_table = make(map[string]interface{})

func Data_handler_init( site_data *map[string]interface{}){

   site_ptr = site_data
   site = (*site_data)["site"].(string)
   fmt.Println("site",site)
   create_redis_data_handle()
   create_constructors(&constructor_table)

}

func create_redis_data_handle(){
  
	site = (*site_ptr)["site"].(string)
    var address =  (*site_ptr)["host"].(string)
    var port = 	int((*site_ptr)["port"].(float64))
	var address_port = address+":"+strconv.Itoa(port)
	var db = int((*site_ptr)["redis_table"].(float64))
	fmt.Println("redis data",address_port,db)
	client = redis.NewClient(&redis.Options{
                                                 Addr: address_port,
												 ReadTimeout : time.Second*30,
												 WriteTimeout : time.Second*30,
												 DB: db,
                                               })
	err := client.Ping(ctx).Err();
	if err != nil{
	         panic("redis data connect")
	  }
    fmt.Println(" redis data connection ping")	
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

   var handler_definitions = make([]map[string]interface{},0)
   var handlers            = make(map[string]interface{})
   construct_handler_definitions( search_list , &handler_definitions)
   construct_redis_handlers(&handler_definitions, &handlers) 

   return &handlers   
}

func construct_handler_definitions( search_list *[]string, handler_definitions *[]map[string]interface{} )   {

   
   
   var packages = graph_query.Common_package_search(&site,search_list)
   ///fmt.Println("packages",len(packages),packages) 
   var data_structures_json = packages[0]["data_structures"]
   //fmt.Println(data_structures_json)
   var data_structures = graph_query.Convert_json_dictionary_interface(  data_structures_json)
   
   var namespace_json = packages[0]["namespace"]
   var namespace = graph_query.Convert_json_string( namespace_json)
   
   //fmt.Println(len(data_structures),data_structures)
   
   
   for _,v := range data_structures{
     var k = v.(map[string]interface{}) 
     var key = namespace +"["+k["type"].(string)+":"+k["name"].(string) +"]"
	 k["key"]= key
	 k["client"] = client
	 *handler_definitions = append(*handler_definitions,k)
   }
}  
   
 
func construct_redis_handlers( handler_definitions *[]map[string]interface{}, handlers *map[string]interface{} ){
   var type_def string
   var name string
   var key string
   var depth int64
   var timeout int64
   for _,v := range *handler_definitions {
      type_def = v["type"].(string)
	 
	  //fmt.Println("key type",type_def)
	  if type_def == "STREAM_REDIS" {
	     key = v["key"].(string)
		 name = v["name"].(string)
		 depth = int64(v["depth"].(float64))
		 (*handlers)[name] = redis_handlers.Construct_Redis_Stream(ctx,client,key,depth)
		  
		 //var x redis_handlers.Redis_Stream_Struct
		 //x =(*handlers)[name].(redis_handlers.Redis_Stream_Struct)

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
	   
	   } else{
	   panic("Key is not expected "+type_def)
	 }
   }
	 
}






