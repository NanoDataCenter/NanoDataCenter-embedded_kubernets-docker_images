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

    
	v.client.Del(v.ctx, v.key )

   
}

func (v Redis_Job_Queue) Length() int64  {

    result,err := v.client.LLen(v.ctx, v.key ).Result()

	if err != nil {
	  panic("Rddis Xadd failed")
	
   }
   return result
}

func (v Redis_Job_Queue) Delete(index int64){

     var remove_mark = "__#####__"
     v.client.LSet(v.ctx,v.key, index,remove_mark)
     v.client.LRem(v.ctx,v.key, 1, remove_mark) 
}

func (v Redis_Job_Queue) List_Range(start int64, stop int64)[]string{

     value ,_ := v.client.LRange(v.ctx, v.key , start, stop).Result()
	 return value
}

func (v Redis_Job_Queue) Pop()string{

     value ,_ := v.client.RPop(v.ctx ,v.key ).Result()
	 return value
}

func (v Redis_Job_Queue) Push( value string){

     v.client.LPush(v.ctx ,v.key,value )
	 v.client.LTrim(v.ctx, v.key , 0, v.depth)
}

func (v Redis_Job_Queue)Delete_jobs( jobs []int64){

     for _,job := range jobs{
	   v.Delete(job)
	 }
		 
}

func (v Redis_Job_Queue)Show_next_job( )string{

     var length = v.Length()
	 if length > 0 {
	   result,_ := v.client.LIndex(v.ctx,v.key,length-1).Result()
	   return result
	 }
	 return ""
	 		 
}

func (v Redis_Job_Queue)Push_front(value string ){

     v.client.RPush(v.ctx ,v.key,value )
	 v.client.LTrim(v.ctx, v.key , 0, v.depth)		 
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