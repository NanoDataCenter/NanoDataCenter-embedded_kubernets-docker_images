package redis_handlers

import "context"
//import "fmt"
import "github.com/go-redis/redis/v8"

type Redis_Single_Structure struct {
   ctx context.Context;
   client *redis.Client;
   key    string;
 

}

func Construct_Redis_Single(  ctx context.Context, client *redis.Client, key string) Redis_Single_Structure   {


   var return_value =  Redis_Single_Structure{ ctx,client,key}

   return return_value



}


func (v Redis_Single_Structure) Get() string {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
    value,err :=  v.client.Get(v.ctx,v.key).Result()
	if err != nil{
	  panic(err)
	}
	return value	
}


func (v Redis_Single_Structure) Set( value string)  {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
    err :=v.client.Set(v.ctx,v.key,value,0) // no expiration
	if err != nil{
	  panic(err)
	}
	

}
