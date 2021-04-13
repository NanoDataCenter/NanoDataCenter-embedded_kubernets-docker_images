package redis_handlers

import "context"
import "time"
//import "fmt"
import "bytes"
import  "lacima.com/Patterns/msgpack"
import "github.com/go-redis/redis/v8"
import "github.com/msgpack/msgpack-go"
import "github.com/satori/go.uuid"


type Time_Out_Function_Type func(  )
type Message_Handler_Type func( input string )string

type Redis_RPC_Struct struct {
   ctx context.Context;
   client *redis.Client;
   key    string;
   depth  int64;
   timeout int
   time_out_function Time_Out_Function_Type;
   rpc_handlers map[string]Message_Handler_Type;
  
}







func Construct_Redis_RPC(  ctx context.Context, client *redis.Client, key string, timeout int, depth int64 ) Redis_RPC_Struct  {


   var return_value  Redis_RPC_Struct 
   
   return_value.ctx = ctx
   return_value.client = client
   return_value.key = key
   return_value.timeout = timeout*10 // timout is in seconds sampling is in .1 seconds
   return_value. time_out_function = nil
   return_value.rpc_handlers = make(map[string]Message_Handler_Type)
   
   return return_value


}

func (v Redis_RPC_Struct) Delete_all() {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
    v.client.Del(v.ctx,v.key)

}


func (v Redis_RPC_Struct) Add_Timeout(timeout_fn Time_Out_Function_Type) {

   v.time_out_function = timeout_fn

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

	 err := v.client.LTrim(v.ctx, v.key , 0, v.depth)
	 if err != nil {
	   panic(err)
	 }	 
}

func (v Redis_RPC_Struct)rpc_start() {
  go (v).start()
}


func (v Redis_RPC_Struct)start() {

   for true {
     if v.Length() != 0 {
        defer recover()
		var input    = msgpack_utils.Unpack(v.Pop()).(map[string]interface{})
		var key      = input["key"].(string)
		var method   = input["method"].(string)
		var params   = input["params"].(string)
		if _,ok := v.rpc_handlers[method]; ok == true {
		    var response = v.rpc_handlers[method](params)
            v.Push_Response(key,response)  
		}
     }else{

	   time.Sleep(time.Second/10)
	 }
	 
   }
}
   
func (v Redis_RPC_Struct)send_rpc_message( method, parameters string )string {

   var request = make(map[string]string)
   request["method"] = method
   request["parameters"] = parameters
   u2 := uuid.NewV4().String()

   request["key"] = u2
   var b bytes.Buffer	
   msgpack.Pack(&b,request)
   v.Push(b.String())
   for i:=0;i<v.timeout;i++{
      length,_:=v.client.LLen(v.ctx,u2).Result()
      if length > 0 {
	     value ,err := v.client.RPop(v.ctx ,u2 ).Result()
	     if err != nil {
	         panic(err)
	     }
	   return value
	 }else{
	  time.Sleep(time.Second/10)
	 }
   }
   panic("communication failure")
   
   return ""

}
