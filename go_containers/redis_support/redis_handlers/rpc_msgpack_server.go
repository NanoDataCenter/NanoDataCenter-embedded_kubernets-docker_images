package redis_handlers

/*
func (v Redis_RPC_Struct)msgpack_start() {

   for true {
     if v.Length() != 0 {
        defer recovery()
		
		var input    = msgpack_utils.Unpack(v.Pop()).([]interface{})
		var method = string(input[0].([]byte))
		var key    = string(input[2].([]byte))
		var params_interface = input[1].(map[interface{}]interface{})
		var params = make(map[string]interface{})
		for p_key, p_value := range params_interface{
		    params[p_key.(string)] = p_value
		}
		

		if _,ok := v.rpc_handlers[method]; ok == true {
		    var response = v.rpc_handlers[method](params)
			var b bytes.Buffer	
            msgpack.Pack(&b,response)
			fmt.Println(len(b.String()),b.String())
            v.Push_Response(key,b.String())  
		}else{
		  panic("bad "+method)
		}
	
     }else{

	   time.Sleep(time.Second/10)
	 }
	 
   }
}



   
func (v Redis_RPC_Struct)Send_msgpack_rpc_message( method string, parameters *map[string]interface{}) interface{} {

   var request = make([]interface{},0)
   
   
   u2 := uuid.NewV4().String()
  
   request = append(request,method,(*parameters),u2)
 
   var b bytes.Buffer	
   msgpack.Pack(&b,request)
   v.Push(b.String())
   for i:=int64(0);i<v.timeout;i++{
      length,_:=v.client.LLen(v.ctx,u2).Result()
      if length > 0 {
	     result ,err := v.client.RPop(v.ctx ,u2 ).Result()
	     if err != nil {
	         panic(err)
	     }
		 fmt.Println("packed_result",len(result),result)
	     panic("done")
		 return nil
	   
	 }else{
	  time.Sleep(time.Second/10)
	 }
   }
   panic("communication failure")
   
   return nil

}
*/