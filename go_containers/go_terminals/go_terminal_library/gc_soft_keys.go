package gc_support



import (
	gc "github.com/gbin/goncurses"
   
	//"fmt"
)




func Init_SoftKey(  ){
  
  
  temp, _ := gc.Init()
  StartColor() 
  Stdscr = temp
  gc.Echo(false)
  Stdscr.Keypad(true)
  
  init_panel_handler() 

}

    


