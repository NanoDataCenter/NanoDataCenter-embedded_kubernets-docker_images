package superius_data

/*
 * This package serves two purposes
 * 1.  This package is used to remove data structures which are not used any more
 *     Event brokers must be cleared of out of date keys
 * 2.  This package will provide a dictionary and key types
 * 
 * This file will be called by the site manager and ansible installation and maintence programs
 * This Package assumes that the redis db has been connected and the tree db has been setup
 * 
 * 
 */
import "context"
import "fmt"
import "time"
import "strconv"
import "github.com/go-redis/redis/v8"




const graph_db int64 = 3
const data_db  int64 = 4

type Superius_data_type struct{
    Key_dictionary     map[string]string
    Key_list           []string
    Current_data_keys  []string
    Orphan_keys        []string
    graph_client       *redis.Client
    data_client        *redis.Client
    graph_ctx          context.Context
    data_ctx           context.Context 
 
}






func Construct_Superius_Data_Removal(host string ,port int64) Superius_data_type{
    
   var return_value    Superius_data_type
   
   return_value.graph_client = construct_client(host,port, graph_db)
   return_value.graph_ctx = context.TODO()
   
   return_value.data_client = construct_client(host,port, data_db)
   return_value.data_ctx = context.TODO()
   return_value.Key_dictionary = make(map[string]string)
   return_value.Key_list = make([]string,0)
   return_value.Current_data_keys = make([]string,0)
   return_value.Orphan_keys = make([]string,0)
   return return_value
   
}

func (v *Superius_data_type ) Remove_superius_data(){
 
 v.Mark_superius_data()
  
 // get dictionay
 // get data keys
 // for loop on data keys
    // if not in dictionay 
       // remove data key
    

}

func (v *Superius_data_type) Mark_superius_data(){
    
  v.Key_dictionary,_ = v.graph_client.HGetAll(v.graph_ctx,"key_set").Result()   
  v.Key_list,_ = v.graph_client.LRange(v.graph_ctx,"key_list",0,-1).Result()  
  v.Current_data_keys,_ = v.data_client.Keys(v.data_ctx,"*").Result()
  v.Orphan_keys = make([]string,0)
  for _,key := range v.Current_data_keys {
      
      if _,ok := v.Key_dictionary[key] ; ok == false {
         
          v.Orphan_keys = append(v.Orphan_keys,key)
      }
      
      
  }
  fmt.Println("orphan_keys",len(v.Orphan_keys),len(v.Current_data_keys),len(v.Key_dictionary))
  
  
  
}


func (v *Superius_data_type)Ping_redis_connections(){
    
    ping_test("redis_graph",v.graph_client, v.graph_ctx)
    ping_test("redis_data", v.data_client, v.data_ctx )
    
    
}


func (v *Superius_data_type)Close_handles()   {
    
    v.data_client.Close()
    v.graph_client.Close()
    
    
    
}


    
func construct_client(address string ,port, db int64)*redis.Client{
    
	address_port := address+":"+strconv.Itoa(int(port))
	
	client := redis.NewClient(&redis.Options{
                                                 Addr: address_port,
												 ReadTimeout : time.Second*30,
												 WriteTimeout : time.Second*30,
												 DB: int(db),
                                               })
	
    return client
}

func ping_test( descriptor string, client *redis.Client, ctx context.Context){
    
    err := client.Ping(ctx).Err();
    fmt.Println("err",err)
	if err != nil{
	         panic("client connect")
	  }
    fmt.Println(descriptor)	

}
