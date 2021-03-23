package main
import (
    "fmt"
    
)
// this is a test
func main(){
  for i := 0; i<=10;i++{
    if_func(i)
	switch_func(i)
	
  }
  
}

func if_func( x int)  {
 
   if (x & 1) == 0 {
      fmt.Println("even")
   } else {
     
	 fmt.Println("odd")
   }
}

func switch_func( x int ){

   switch x {
   case 1: fmt.Println(1)
   case 2: fmt.Println(2)  
   case 3: fmt.Println(3)
   case 4: fmt.Println(4)
   case 5: fmt.Println(5)
   case 6: fmt.Println(6)
   case 7: fmt.Println(7)
   case 8: fmt.Println(8)
   case 9: {
             fmt.Print("The number ")
             fmt.Println(9) }
   default: fmt.Println("default")
   }
}   