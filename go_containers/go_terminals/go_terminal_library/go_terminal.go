package gc_support


import (
	gc "github.com/gbin/goncurses"
      
	//"fmt"
)


type CMD_handler func(string)string


func Construct_Console(title string,cmd_handler CMD_handler){
      rows, cols := Stdscr.MaxYX()
   
    
    
    
     w  := len(title)+10
    
    
    
    window, _ := gc.NewWindow(rows,cols,0,0)
    window.Box(0, 0)
    window.ColorOn(NORMAL_WINDOW_COLOR)
    window.SetBackground(gc.ColorPair(NORMAL_WINDOW_COLOR))
    blank_soft_keys()
    window.MoveAddChar(2, 0, gc.ACS_LTEE)
    window.HLine(2, 1, gc.ACS_HLINE, cols-2)
    window.MoveAddChar(2, cols-1, gc.ACS_RTEE)
    window.MoveAddChar(2, 0, gc.ACS_LTEE)
    window.MovePrint(1, (w/2)-(len(title)/2), title)
    
     
    
    
    window.HLine(rows-3, 1, gc.ACS_HLINE, cols-2)
    window.MoveAddChar(rows-3,cols-1, gc.ACS_RTEE)
    window.MoveAddChar(rows-3, 0, gc.ACS_LTEE)
    

    window.MovePrint(rows -2,3,"Press RTN to send command F1 to cls F8 to abort ")
    panel := gc.NewPanel(window)
    panel.Top()
    panel_list.PushFront(panel)
    gc.Cursor(0)
    gc.UpdatePanels()
    sub_window := window.Sub(rows-6,cols-4,3,2)
    defer sub_window.Delete()
    sub_window.ColorOn(POP_UP_COLOR)
    sub_window.SetBackground(gc.ColorPair(POP_UP_COLOR))
   
    sub_window.ScrollOk(true)
    sub_window.Move(rows-7,0)
    sub_window.Print("Enter Cmd")
    sub_window.Move(rows-7,0)
     
    sub_window.Scroll(1)
    sub_window.AddChar('>')
   
    sub_window.Refresh()
    var cmd_string string
    cmd_string  = ""
    current_col := 0
    for true {
      gc.UpdatePanels()
      gc.Update()
      ch := Stdscr.GetChar()
      switch ch {
		case gc.KEY_F1:
            sub_window.Erase()
			sub_window.Move(rows-7,0)
            sub_window.Print("Enter Cmd")
            sub_window.Scroll(1)
            sub_window.Move(rows-7,0)
            sub_window.AddChar('>')
            sub_window.Refresh()
            cmd_string = ""
            current_col = 0
            
            
        case gc.KEY_RETURN, gc.KEY_ENTER:
            output := cmd_handler(cmd_string)
            sub_window.Scroll(1)
            sub_window.Move(rows-7,0)
            sub_window.Print(output)
            sub_window.Scroll(1)
            sub_window.Move(rows-7,0)
            sub_window.AddChar('>')
            sub_window.Refresh()
            cmd_string = ""
            current_col = 0            
            
		case gc.KEY_F8:
             kill_current_panel()
             return
    

		default:
			if current_col < cols- 5 {
                cmd_string = cmd_string + string(ch)
                current_col = current_col + 1
                sub_window.Move(rows-7,current_col)
                
                sub_window.AddChar(gc.Char(ch))
                sub_window.Refresh()
            }
		}
    
    }
    kill_current_panel()
   
    
}

