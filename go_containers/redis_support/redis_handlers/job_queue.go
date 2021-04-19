package redis_handlers

import "context"
//import "fmt"
import "github.com/go-redis/redis/v8"

type Redis_Job_Queue struct {
   ctx context.Context;
   client *redis.Client;
   key    string;
   depth  int64;
  
}

 

func Construct_Redis_Job_Queue(  ctx context.Context, client *redis.Client, key string, depth int64 ) Redis_Job_Queue   {


   var return_value = Redis_Job_Queue{ ctx,client,key,depth}

   return return_value



}

	
	
	
func (v Redis_Job_Queue) Delete_all()   {

    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()

    
	v.client.Del(v.ctx, v.key )
  
}

func (v Redis_Job_Queue) Length() int64  {

    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()

    result,err := v.client.LLen(v.ctx, v.key ).Result()

	if err != nil {
	  panic(err)
	
   }
   return result
}

func (v Redis_Job_Queue) Delete(index int64){

    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()

     var remove_mark = "__#####__"
     err := v.client.LSet(v.ctx,v.key, index,remove_mark)
	 	 if err != nil {
	   panic(err)
	 }
     v.client.LRem(v.ctx,v.key, 1, remove_mark)
	 
}

func (v Redis_Job_Queue) List_Range(start int64, stop int64)[]string{
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()

     value ,err := v.client.LRange(v.ctx, v.key , start, stop).Result()
	 if err != nil {
	   panic(err)
	 }
	 return value
}

func (v Redis_Job_Queue) Pop()string{
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()

     value ,err := v.client.RPop(v.ctx ,v.key ).Result()
	 if err != nil {
	   panic(err)
	 }
	 return value
}

func (v Redis_Job_Queue) Push( value string){
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()

     v.client.LPush(v.ctx ,v.key,value )

	 v.client.LTrim(v.ctx, v.key , 0, v.depth)
		 
}

func (v Redis_Job_Queue)Delete_jobs( jobs []int64){
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()

     for _,job := range jobs{
	   v.Delete(job)
	 }
		 
}

func (v Redis_Job_Queue)Show_next_job( )string{
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()

     var length = v.Length()
	 if length > 0 {
	   result,err := v.client.LIndex(v.ctx,v.key,length-1).Result()
	   if err != nil {
	   panic(err)
	    }
	   return result
	 }
	 return ""
	 		 
}

func (v Redis_Job_Queue)Push_front(value string ){
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()

     v.client.RPush(v.ctx ,v.key,value )

	 err := v.client.LTrim(v.ctx, v.key , 0, v.depth)
	 if err != nil {
	   panic(err)
	 }	 
}

/*

LIndex(key string, index int64) *StringCmd
	LInsert(key, op string, pivot, value interface{}) *IntCmd
	LInsertBefore(key string, pivot, value interface{}) *IntCmd
	LInsertAfter(key string, pivot, value interface{}) *IntCmd
	LLen(key string) *IntCmd
	LPop(key string) *StringCmd
	LPush(key string, values ...interface{}) *IntCmd
	LPushX(key string, value interface{}) *IntCmd
	LRange(key string, start, stop int64) *StringSliceCmd
	LRem(key string, count int64, value interface{}) *IntCmd
	LSet(key string, index int64, value interface{}) *StatusCmd
	LTrim(key string, start, stop int64) *StatusCmd
	
	RPop(key string) *StringCmd
	RPopLPush(source, destination string) *StringCmd
	RPush(key string, values ...interface{}) *IntCmd
	RPushX(key string, value interface{}) *IntCmd





*/