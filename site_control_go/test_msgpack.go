
package main

import "bytes"
import "github.com/msgpack/msgpack-go"
import "fmt"


func main() {
  fmt.Println("test")
  b := &bytes.Buffer{}
  _, err := msgpack.PackUint8(b, 52)
   msgpack.PackUint8(b, 57)
  fmt.Println(err)
  /*
  retval, _, e := msgpack.Unpack(b)
  
  fmt.Println(retval,e)
    retval1, _, e1 := msgpack.Unpack(b)
  */
  for i:= 0; i<10 ; i++{
     retval,_,e := msgpack.Unpack(b)
     fmt.Println(retval,e)
  }
}
		
