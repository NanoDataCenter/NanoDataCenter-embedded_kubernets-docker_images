package redis_handlers

import "context"
import "time"
import "fmt"

import "github.com/go-redis/redis/v8"



type Message_Handler_Type func( parameters map[string]interface{} ) map[string]interface{}

type Redis_RPC_Struct struct {
   ctx context.Context;
   client *redis.Client;
   key    string;
   depth  int64
   timeout int64
   rpc_handlers map[string]Message_Handler_Type;
  
}







func Construct_Redis_RPC(  ctx context.Context, client *redis.Client, key string, timeout int64, depth int64 ) Redis_RPC_Struct  {


   var return_value  Redis_RPC_Struct 
   
   return_value.ctx = ctx
   return_value.client = client
   return_value.key = key
   return_value.timeout = timeout*10 // timout is in seconds sampling is in .1 seconds

   return_value.rpc_handlers = make(map[string]Message_Handler_Type)
   return return_value


}

func (v Redis_RPC_Struct) Delete_all() {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
    v.client.Del(v.ctx,v.key)

}



func (v Redis_RPC_Struct) Add_handler( key string, handler Message_Handler_Type){

  v.rpc_handlers[key] = handler
  
}  


func (v Redis_RPC_Struct) Length() int64  {

    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()

    result,err := v.client.LLen(v.ctx, v.key ).Result()

	if err != nil {
	  panic(err)
	
   }
   return result
}

func (v Redis_RPC_Struct) Pop()string{
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()

     value ,err := v.client.RPop(v.ctx ,v.key ).Result()
	 if err != nil {
	   panic(err)
	 }
	 return value
}


func (v Redis_RPC_Struct) Push_Response( key, value string){
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
    v.client.LPush(v.ctx ,key,value )
    v.client.Expire(v.ctx,key,time.Minute)
}

func (v Redis_RPC_Struct) Push( value string){
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()

     v.client.LPush(v.ctx ,v.key,value )

	 v.client.LTrim(v.ctx, v.key , 0, v.depth)

}



func recovery(){
   if r := recover(); r != nil {
        
		   fmt.Println(r)
		   panic("done")   
    }

}

