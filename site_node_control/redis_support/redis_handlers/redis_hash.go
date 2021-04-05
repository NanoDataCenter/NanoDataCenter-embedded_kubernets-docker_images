package redis_handlers


//import "github.com/go-redis/redis/v8"



import "context"
//import "fmt"
import "github.com/go-redis/redis/v8"

type Redis_Hash_Struct struct {
   ctx context.Context;
   client *redis.Client;
   key    string;
  
  
}



func Construct_Redis_Hash(  ctx context.Context, client *redis.Client, key string) Redis_Hash_Struct   {


   var return_value = Redis_Hash_Struct{ ctx,client,key}

   return return_value



}


    

func (v Redis_Hash_Struct) Delete_All(field string)   {
     Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
 	v.client.Del(v.ctx, v.key ).Err()

} 

func (v Redis_Hash_Struct) HDel(field string)   {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
 	var err = v.client.HDel(v.ctx, v.key, field ).Err()
    if err != nil {
	  panic(err)
	}
}

func (v Redis_Hash_Struct) HExists(field string) bool   {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
 	val,err := v.client.HExists(v.ctx,  v.key, field ).Result()
	if (err !=  redis.Nil) && (err != nil){
	   
	   panic(err)
	}
	    
	return val
}

func (v Redis_Hash_Struct) HGet(field string) string  {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
	
	
 	   val, err := v.client.HGet(v.ctx,  v.key, field ).Result()
	   if (err !=  redis.Nil) && (err != nil){
	        panic(err)
	   }
      
      return val 
}

func (v Redis_Hash_Struct) HGetAll() map[string]string  {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
 	val, err := v.client.HGetAll(v.ctx,  v.key ).Result()
	if (err !=  redis.Nil) && (err != nil){
	   
	   panic(err)
	}
	return val
}

func (v Redis_Hash_Struct) HKeys() []string {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
 	val,err := v.client.HKeys(v.ctx, v.key ).Result()
    
    if err != nil {
	  panic(err)
	}
	return val
}

func (v Redis_Hash_Struct) HLen() int64 {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
    val,err:= v.client.HLen(v.ctx, v.key).Result()
	if (err !=  redis.Nil) && (err != nil){
	   
	   panic(err)
	}

	return val
}

func (v Redis_Hash_Struct) HSet(field,value string)  {
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
  var err =  v.client.HSet(v.ctx, v.key,field,value).Err()
    if err != nil {
	  panic(err)
	}
	
}



