package main
import (
    "fmt"
    "os"
)
// this is a test
func main(){
  /*
  fmt.Print("Enter a number:  ")
  var input float64
  fmt.Scan("%f",&input)
  output := input *3
  fmt.Println(output)
 */
   argsWithProg := os.Args
   argsWithoutProg := os.Args[1:]

   arg := os.Args[3]
   fmt.Println(argsWithProg)
   fmt.Println(argsWithoutProg)
   fmt.Println(arg)  
}