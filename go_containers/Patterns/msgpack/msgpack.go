package msgpack_utils

import "fmt"
import "bytes"
import "github.com/msgpack/msgpack-go"

func Unpack( packed_data string )interface{}{
 
 var buf = bytes.NewBufferString(packed_data)
 unpack_data, _, err := msgpack.Unpack(buf)
  if err != nil {
		panic("bad msgpack data 1")
  }
  //fmt.Println("unpack_data",unpack_data)
  return unpack_data.Interface()
}

func Unpack_byte_array( packed_data []byte )interface{}{
 
 var buf = bytes.NewBuffer(packed_data)
 unpack_data, _, err := msgpack.Unpack(buf)
  if err != nil {
		panic("bad msgpack data 1")
  }
  //fmt.Println("unpack_data",unpack_data)
  return unpack_data.Interface()
}


func Convert_rpc_return( input interface{} ) map[string]interface{}{
   return_value := make(map[string]interface{})
   temp := input.(map[string]bool)
   for key , value := range temp{
  
	 fmt.Println(key,value)
   }
   panic("stop here")
   return return_value
}


/*
func Pack_data( data interface{})string {


    var b bytes.Buffer	
    msgpack.Pack(&b,&data)
    return (b.String())
   
   
}
*/
