package main


import "fmt"
import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/redis_file"

//import "github.com/Shopify/go-lua"
import "lacima.com/zygo"


const file_base = "/files/"


func main(){

    

	var config_file = "/data/redis_server.json"
	
	var site_data_store map[string]interface{}

	site_data_store = get_site_data.Get_site_data(config_file)
    graph_query.Graph_support_init(&site_data_store)
	redis_handlers.Init_Redis_Mutex()
	data_handler.Data_handler_init(&site_data_store)
    
  
   env := zygo.NewZlisp()
   
   sym1 := env.MakeSymbol("foo")
sym2 := env.MakeSymbol("foo")
fmt.Println(sym1.Name() == sym2.Name() , sym1.Number() == sym2.Number()) // should be true
}
