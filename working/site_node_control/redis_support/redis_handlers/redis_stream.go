package redis_handlers

import "context"
//import "fmt"
import "github.com/go-redis/redis/v8"

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

func (v Redis_Stream_Struct) Delete_all() {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
   v.client.Del(v.ctx,v.key)

}

func (v Redis_Stream_Struct) Xadd(packed_data string) string  {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
    var xdata = make(map[string]interface{})
	xdata["data"]  = packed_data
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





/* Go commands are documented here
https://github.com/sp0n-7/redis/blob/v6.15.0/commands.go#L1373



from https://sourcegraph.com/github.com/go-redis/redis@b3d392315ba16c2bfb0fcab2655e0d401f36ffa5/-/blob/commands_test.go#L3488:14
Example for stream github
("should XRevRange", func() {
			msgs, err := client.XRevRange("stream", "+", "-").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(msgs).To(Equal([]redis.XMessage{
				{ID: "3-0", Values: map[string]interface{}{"tres": "troix"}},
				{ID: "2-0", Values: map[string]interface{}{"dos": "deux"}},
				{ID: "1-0", Values: map[string]interface{}{"uno": "un"}},
			}))
			

*/