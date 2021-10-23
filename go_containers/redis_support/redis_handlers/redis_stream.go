package redis_handlers

import "fmt"
import "time"
import "context"
import "strconv"
import "github.com/go-redis/redis/v8"
import "lacima.com/Patterns/msgpack_2"

type Redis_Stream_Struct struct {
   ctx context.Context;
   client *redis.Client;
   key    string;
   depth  int64;
  
}



func Construct_Redis_Stream(  ctx context.Context, client *redis.Client, key string, depth int64 ) Redis_Stream_Struct   {


   var return_value = Redis_Stream_Struct{ ctx,client,key,depth}

   return return_value



}


func (v Redis_Stream_Struct)Get_client() *redis.Client{
  return v.client
}

func (v Redis_Stream_Struct)Get_context() context.Context{
  return v.ctx
}



func (v Redis_Stream_Struct) Delete_all() {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
   v.client.Del(v.ctx,v.key)

}

func (v Redis_Stream_Struct) Xadd(data interface{} ) string  {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
    var xdata = make(map[string]interface{})
	xdata["data"]  = msg_pack_utils.Pack_interface(data)
	xdata["time"]  = msg_pack_utils.Pack_int64(time.Now().UnixNano())
	var x_add_args = redis.XAddArgs{ v.key, v.depth, v.depth, "*" ,xdata}
	result,err := v.client.XAdd(v.ctx, &x_add_args ).Result()
	if err != nil {
	  panic(err)
	
   }
   return result
}




func (v Redis_Stream_Struct) Xlen() int64  {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
 	val, err := v.client.XLen(v.ctx, v.key ).Result()
	if err != nil {
	  panic(err)
	
   }	
	return val
}

func (v Redis_Stream_Struct) Xtrim( length int64)   {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
 	err := v.client.XTrimApprox(v.ctx, v.key,length )
	if err != nil {
	  panic(err)
	
   }
   
}


   
func (v Redis_Stream_Struct) XRange( low_ts, high_ts string) {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
 	value, err := v.client.XRange(v.ctx, v.key,low_ts,high_ts).Result()
	if err != nil {
	  panic(err)
	
   }
   fmt.Println("value",value)
   panic("done")
   
}
func (v Redis_Stream_Struct) XRangeN(low_ts, high_ts string, length int64)   {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
 	value,err := v.client.XRangeN(v.ctx, v.key,low_ts,high_ts,length ).Result()
	if err != nil {
	  panic(err)
	
   }
   fmt.Println("value",value)
   
   
}
func (v Redis_Stream_Struct)XRevRange(high_ts,low_ts string )   {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
 	value,err := v.client.XRevRange(v.ctx, v.key,high_ts,low_ts ).Result()
	if err != nil {
	  panic(err)
	
   }
   fmt.Println("value",value)
   panic("done")
}
func (v Redis_Stream_Struct) XRevRangeN(high_ts, low_ts string, length int64)   {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
 	value,err := v.client.XRevRangeN(v.ctx, v.key,high_ts,low_ts, length).Result()
	if err != nil {
	  panic(err)
	
   }
   
   for _,item := range value{
       fmt.Println(item.ID)
       
       fmt.Println(msg_pack_utils.Unpack_interface(item.Values["data"].(string)))
   }
   
   
}
 
func (v Redis_Stream_Struct)Convert_ts( time_stamp int64)string{
    return strconv.FormatInt(time_stamp,10)   
}   
   

