
package gc_support



import (
	gc "github.com/gbin/goncurses"
	//"fmt"
)

const MAX_COLOR_PAIRS       = 11
const SOFT_KEY_COLOR        = 6
const NORMAL_WINDOW_COLOR   = 4
const POP_UP_COLOR          = 5
const MENU_START            = 9

var colours = []int16{gc.C_BLACK, gc.C_BLUE, gc.C_CYAN, gc.C_GREEN,
		gc.C_MAGENTA, gc.C_RED, gc.C_WHITE, gc.C_YELLOW}

var attributes = []struct {
		attr gc.Char
		text string
	}{
		{gc.A_NORMAL, "normal"},
		{gc.A_STANDOUT, "standout"},
		{gc.A_UNDERLINE | gc.A_BOLD, "underline"},
		{gc.A_REVERSE, "reverse"},
		{gc.A_BLINK, "blink"},
		{gc.A_DIM, "dim"},
		{gc.A_BOLD, "bold"},
	}


func StartColor(){
    
    gc.StartColor()
    init_color_pairs()
}

func init_color_pairs(){
    gc.InitPair(1, gc.C_BLACK, gc.C_WHITE)
    gc.InitPair(2, gc.C_WHITE, gc.C_BLACK)
    gc.InitPair(3, gc.C_WHITE, gc.C_GREEN)
    gc.InitPair(4, gc.C_BLACK, gc.C_GREEN)
    gc.InitPair(5, gc.C_GREEN, gc.C_BLACK)
    gc.InitPair(6, gc.C_RED, gc.C_GREEN)
    gc.InitPair(7, gc.C_BLUE,  gc.C_YELLOW)
    gc.InitPair(8, gc.C_CYAN,  gc.C_RED)
    gc.InitPair(9, gc.C_RED, gc.C_BLACK)
	gc.InitPair(10, gc.C_GREEN, gc.C_BLACK)
	gc.InitPair(11, gc.C_BLUE,  gc.C_YELLOW)    
}    

        
func Set_reverse_on(scr *gc.Window ){
  scr.AttrOn(gc.A_REVERSE)
    
}

func Set_reverse_off(scr *gc.Window){
   scr.AttrOff(gc.A_REVERSE) 
}

func Set_blink_on(scr *gc.Window){
    scr.AttrOn(gc.A_BLINK)
}

func Set_blink_off(scr *gc.Window){
   scr.AttrOff(gc.A_BLINK) 
}

func Set_normal_on( scr *gc.Window){
    scr.AttrOn(gc.A_NORMAL)
}

func Set_normal_off(scr *gc.Window){
    scr.AttrOff(gc.A_NORMAL)
}

func Set_bold_on(scr *gc.Window){
    scr.AttrOn(gc.A_BOLD)
}

func Set_bold_off(scr *gc.Window){
  scr.AttrOff(gc.A_BOLD)  
}

func Set_dim_on(scr *gc.Window){
    scr.AttrOn(gc.A_DIM)
}

func Set_dim_off(scr *gc.Window){
   scr.AttrOff(gc.A_DIM) 
}






