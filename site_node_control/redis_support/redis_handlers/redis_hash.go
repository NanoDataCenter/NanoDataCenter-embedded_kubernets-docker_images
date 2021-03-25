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

 	v.client.Del(v.ctx, v.key )

} 

func (v Redis_Hash_Struct) HDel(field string)   {

 	v.client.HDel(v.ctx, v.key, field )

}

func (v Redis_Hash_Struct) HExists(field string) bool   {

 	val, _ := v.client.HExists(v.ctx,  v.key, field ).Result()

	return val
}

func (v Redis_Hash_Struct) HGet(field string) string  {

 	val, _ := v.client.HGet(v.ctx,  v.key, field ).Result()

	return val
}

func (v Redis_Hash_Struct) HGetAll() map[string]string  {

 	val, _ := v.client.HGetAll(v.ctx,  v.key ).Result()


	return val
}

func (v Redis_Hash_Struct) HKeys() []string {

 	val, _ := v.client.HKeys(v.ctx, v.key ).Result()

	return val 
}

func (v Redis_Hash_Struct) HLen() int64 {

    val,_ := v.client.HLen(v.ctx, v.key).Result()
 	

	return val
}

func (v Redis_Hash_Struct) HSet(field,value string)  {

    v.client.HSet(v.ctx, v.key,field,value)
 	
	
}



