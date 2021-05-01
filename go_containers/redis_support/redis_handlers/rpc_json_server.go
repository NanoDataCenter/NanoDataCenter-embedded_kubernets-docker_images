package redis_handlers


//import "fmt"
import "time"
import "encoding/json"


import "github.com/satori/go.uuid"



func (v Redis_RPC_Struct)Json_Rpc_start() {
  go (v).json_start()
}




func (v Redis_RPC_Struct)json_start() {

   for true {
     if v.Length() != 0 {
        
		v.json_handler_request()
	   
     }else{

	   time.Sleep(time.Second/10)
	 }
	 
   }

}

func (v Redis_RPC_Struct)json_handler_request(){
  defer recovery()
  input := v.Pop()
  //fmt.Println(input)
  input_unmarshall := make([]interface{},0)
  byte_array := []byte(input)
  if err := json.Unmarshal( byte_array, &input_unmarshall); err != nil {
        panic(err)
    }
  method := input_unmarshall[0].(string)
  uuid   := input_unmarshall[2].(string)
  params := input_unmarshall[1].(map[string]interface{})
  if _,ok := v.rpc_handlers[method]; ok == true {
  
       response := v.rpc_handlers[method](params)
       request_json,err := json.Marshal(&response)
       if err != nil{
          panic("json marshall error")
        }  
        v.Push_Response(uuid,string(request_json))  
   }else{
		  panic("bad "+method)
   }
	


}


  
func (v Redis_RPC_Struct)Send_json_rpc_message( method string, parameters map[string]interface{}) map[string]interface{} {

   var request = make([]interface{},0)
   
   
   u2 := uuid.NewV4().String()
  
   request = append(request,method,parameters,u2)
   request_json,err := json.Marshal(&request)
   if err != nil{
     panic("json marshall error")
   }
   //fmt.Println("request",err,string(request_json))
   
   v.Push(string(request_json))
   for i:=int64(0);i<v.timeout;i++{
      length,_:=v.client.LLen(v.ctx,u2).Result()
      if length > 0 {
	     result ,err := v.client.RPop(v.ctx ,u2 ).Result()
	     if err != nil {
	         panic(err)
	     }
		 input_unmarshall := make(map[string]interface{})
         byte_array := []byte(result)
         if err := json.Unmarshal( byte_array, &input_unmarshall); err != nil {
            panic(err)
         }
		 //fmt.Println("input_unmarshall",input_unmarshall)
	     
 	     return input_unmarshall
	 }else{
	  time.Sleep(time.Second/10)
	 }
   }
   panic("communication failure")
   
   return nil

}
