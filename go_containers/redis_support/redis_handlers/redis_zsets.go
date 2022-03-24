package redis_handlers


//import "github.com/go-redis/redis/v8"



import "context"
//import "fmt"
import "github.com/go-redis/redis/v8"

type Redis_ZSet_Struct struct {
   ctx context.Context;
   client *redis.Client;
   key    string;
  
  
}



func Construct_Redis_ZSet(  ctx context.Context, client *redis.Client, key string) Redis_ZSet_Struct   {


   var return_value = Redis_ZSet_Struct{ ctx,client,key}

   return return_value

}

func ( v Redis_ZSet_Struct )ZAdd(item string, score float64){
          Lock_Redis_Mutex()
	      defer UnLock_Redis_Mutex()
          err   :=   v.client.ZAdd(v.ctx, v.key, &redis.Z{score, item} ).Err()
		 if err != nil {
	       panic(err)
	     }
}






func ( v Redis_ZSet_Struct )ZRange(start, stop int64)[]string{
          
          Lock_Redis_Mutex()
	      defer UnLock_Redis_Mutex()
          result,err   :=   v.client.ZRange(v.ctx, v.key, start, stop ).Result()
           if err != nil {
	       panic(err)
	     }
        return result
}


func (v Redis_ZSet_Struct)ZRank( item string )(int64,bool){
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
    data := v.client.ZRank(v.ctx,v.key,item)
    return_value1 := data.Val()
    return_value2 := true
    if data.Err() != nil {
        return_value2 = false
    }
    return return_value1,return_value2

}


func (v Redis_ZSet_Struct)ZScore( item string ) float64{

         Lock_Redis_Mutex()
	      defer UnLock_Redis_Mutex()    
         zRank :=   v.client.ZScore(v.ctx, v.key, item)
         return zRank.Val()
}

func (v Redis_ZSet_Struct)ZRem(item string ){
        Lock_Redis_Mutex()
        defer UnLock_Redis_Mutex()
        v.client.ZRem(v.ctx , v.key, item)
}   

func (v Redis_ZSet_Struct)ZPopmin(count int64)([]redis.Z){
         Lock_Redis_Mutex()
	    defer UnLock_Redis_Mutex()
       data,err :=  v.client.ZPopMin( v.ctx , v.key , count).Result()
       if err != nil {
           panic("error")
       }
       return data

}

func (v Redis_ZSet_Struct)ZPopmax( count int64)([]redis.Z){
         Lock_Redis_Mutex()
	    defer UnLock_Redis_Mutex()
       data,err :=  v.client.ZPopMax( v.ctx , v.key , count).Result()
       if err != nil{
           panic("error")
       }
       return data

}

func (v Redis_ZSet_Struct) Delete_All()   {
     Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
 	v.client.Del(v.ctx, v.key ).Err()

} 

/*
It("should ZPopMin", func() {
			err := client.ZAdd(ctx, "zset", &redis.Z{
				Score:  1,
				Member: "one",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = client.ZAdd(ctx, "zset", &redis.Z{
				Score:  2,
				Member: "two",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = client.ZAdd(ctx, "zset", &redis.Z{
				Score:  3,
				Member: "three",
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			members, err := client.ZPopMin(ctx, "zset").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(members).To(Equal([]redis.Z{{
				Score:  1,
				Member: "one",
			}}))

			// adding back 1
			err = client.ZAdd(ctx, "zset", &redis.Z{
				Score:  1,
				Member: "one",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			members, err = client.ZPopMin(ctx, "zset", 2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(members).To(Equal([]redis.Z{{
				Score:  1,
				Member: "one",
			}, {
				Score:  2,
				Member: "two",
			}}))

			// adding back 1 & 2
			err = client.ZAdd(ctx, "zset", &redis.Z{
				Score:  1,
				Member: "one",
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			err = client.ZAdd(ctx, "zset", &redis.Z{
				Score:  2,
				Member: "two",
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			members, err = client.ZPopMin(ctx, "zset", 10).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(members).To(Equal([]redis.Z{{
				Score:  1,
				Member: "one",
			}, {
				Score:  2,
				Member: "two",
			}, {
				Score:  3,
				Member: "three",
			}}))
		})
    
    */
