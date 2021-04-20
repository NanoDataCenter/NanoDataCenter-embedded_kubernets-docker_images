package redis_file




import "context"
//import "fmt"
import "strconv"
import "time"
import "sort"

import "github.com/go-redis/redis/v8"



type Redis_File_Struct struct {
   ctx context.Context;
   client *redis.Client;
  
}


const file_db  int = 13 


var ctx    = context.TODO()
var client *redis.Client

func create_redis_data_handle( address string, port int ){
  

    Init_Redis_Mutex()
	var address_port = address+":"+strconv.Itoa(port)	
	client =redis.NewClient(&redis.Options{
                                                 Addr: address_port,
												 ReadTimeout : time.Second*30,
												 WriteTimeout : time.Second*30,
												 DB: file_db,
                                               })
	err := client.Ping(ctx).Err();
	if err != nil{
	         panic("redis file error")
	  }
    
}	


func Construct_File_Struct(  ) *Redis_File_Struct  {

   var return_value Redis_File_Struct
   return_value.ctx     = ctx
   return_value.client = client
   return &return_value


}

func ( v *Redis_File_Struct)Set(path,value string){

    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
    v.client.Set(v.ctx,path,value,0) 



}

func ( v *Redis_File_Struct)Get(path string) string {

    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
    value,err :=  v.client.Get(v.ctx,path).Result()
	if err != nil{
	  panic(err)
	}
	return value	


}

func ( v *Redis_File_Struct)Ls(pattern string)[]string{

    var return_value []string
    Lock_Redis_Mutex()
	defer UnLock_Redis_Mutex()
    result := v.client.Keys(v.ctx,pattern).Args() 
    for _,key := range result{
	   return_value = append(return_value,key.(string))
	}
	sort.Strings(return_value)
    return return_value
}

func ( v *Redis_File_Struct)Rm(pattern string){

    keys := v.Ls(pattern)
    Lock_Redis_Mutex()
   defer UnLock_Redis_Mutex()
   for _,i := range keys{
     v.client.Del(v.ctx,i)
   
   }
   
}

func ( v *Redis_File_Struct)Flush(pattern string){
 
    Lock_Redis_Mutex()
   defer UnLock_Redis_Mutex()
   v.client.FlushDB(v.ctx)
}