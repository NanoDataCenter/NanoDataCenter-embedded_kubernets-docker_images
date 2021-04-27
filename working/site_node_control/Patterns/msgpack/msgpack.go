package msgpack_utils

//import "fmt"
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
/*
func Pack_data( data interface{})string {


    var b bytes.Buffer	
    msgpack.Pack(&b,&data)
    return (b.String())
   
   
}
*/