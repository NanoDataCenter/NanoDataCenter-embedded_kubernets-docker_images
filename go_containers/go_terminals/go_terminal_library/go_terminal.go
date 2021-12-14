package gc_support


import (
	gc "github.com/gbin/goncurses"
     
      
	//"fmt"
)




type CMD_handler func(string)string


var line_limit int

func Construct_Console_Menu(window *gc.Window, title string, message string ){
    title =  title +" F8 to return F7 screen clear"

      
    _, cols := Stdscr.MaxYX()
 
     w  := cols
    
    window.Clear()
    window.MovePrint(1, (w/2)-(len(title)/2), title)
    window.MovePrint(2,(w/2)-(len(message)/2), message)
    window.Refresh()
    current_line.sub_window.Touch()
    current_line.sub_window.Refresh()

       
    
    
}
 

func Construct_Console( cmd_handler CMD_handler, user_handler  func( gc.Key )bool, init_handler func()){
     
     
     ref_lines  = make(map[LINE_TYPE]LINE_BUFFER)
     past_lines = make([]LINE_TYPE,0)
      
    rows, cols := Stdscr.MaxYX()
    
    line_limit = 160
    if line_limit > cols {
        line_limit = cols 
    }
    
    
    
    
    
    
    window, _ := gc.NewWindow(rows,cols,0,0)
    
    window.ColorOn(NORMAL_WINDOW_COLOR)
    window.SetBackground(gc.ColorPair(NORMAL_WINDOW_COLOR))
   
    
   
   
    
    
    

   
    
    panel := gc.NewPanel(window)
    panel.Top()
    panel_list.PushFront(panel)
    gc.Cursor(0)
    gc.UpdatePanels()
    sub_window, _ := gc.NewWindow(rows-4,cols,4,0)
    sub_window.ColorOn(POP_UP_COLOR)
    sub_window.ScrollOk(true)
    sub_window.SetBackground(gc.ColorPair(POP_UP_COLOR))
    sub_window.Clear()
    
    current_line.initialize_line(rows-5,cols,window,sub_window)
    
    
     
   
   
    sub_window.Refresh()
    window.Refresh()
    
    init_handler()
    
    current_index = 0
    status   := false
    for status == false {
      
      gc.UpdatePanels()
      gc.Update()
      ch := Stdscr.GetChar()
      
      
      
      if user_handler != nil {
       
          status= user_handler(ch)
          
 
      }
      if status == false {
      
      switch ch {

            
          case gc.KEY_F1,gc.KEY_F2, gc.KEY_F3, gc.KEY_F4, gc.KEY_F5,gc.KEY_F6,gc.KEY_F9, gc.KEY_F10, gc.KEY_F11,gc.KEY_F12:
              ;
            
        case gc.KEY_RETURN, gc.KEY_ENTER:
       
            current_index = 0
            output := cmd_handler(string(current_line.line[1:]))
            current_line.Return_cmd( output)
                 
		case gc.KEY_F7:
            window.Clear()
            init_handler()
            window.Refresh()
            sub_window.Clear()
            sub_window.SetBackground(gc.ColorPair(POP_UP_COLOR))
            current_line.initialize_line(rows-5,cols,window,sub_window)
            sub_window.Refresh()        
               
            
		case gc.KEY_F8:
            sub_window.Delete()
             kill_current_panel()
             
             return
    
        case gc.KEY_DOWN:
            if current_index >0 {
               current_index -=1
               new_index := past_lines[len(past_lines)-1-current_index]
               current_line = ref_lines[new_index]
               current_line.redraw()

            }
           
        case gc.KEY_UP:
            
            if current_index <= len(past_lines)-1 {
               
               new_index := past_lines[len(past_lines)-1-current_index]
               current_line = ref_lines[new_index]
               current_line.redraw()
               if current_index < len(past_lines)-1{ 
                 current_index +=1
               }
                 
            }
           
        case gc.KEY_LEFT:
            current_index = 0
            current_line.left()
            
        case gc.KEY_RIGHT:
             current_index = 0
             current_line.right()
            
        case gc.KEY_BACKSPACE:
             current_index = 0
             current_line.backspace()
             
		default:
             current_index = 0
			 current_line.add_character(byte(ch))
            }
		}
       
    }
    kill_current_panel()
    
    
}

func map_special_character(ch gc.Key)string{
    switch ch {
        case gc.KEY_F1:
             return "F1"
        case gc.KEY_F2:
             return "F2"
        case gc.KEY_F3:
             return "F3"
        case gc.KEY_F4:
             return "F4"
        case gc.KEY_F5:
             return "F5"
        case gc.KEY_F6:
            return "F6"
        default:
            ;
    }
    return ""

}

type LINE_TYPE   [150]byte   

type LINE_BUFFER struct{
    cursor       int
    line         LINE_TYPE 
    sub_window  *gc.Window
    Window      *gc.Window
    rows         int
    cols         int
}

var ref_lines       map[LINE_TYPE]LINE_BUFFER
var past_lines      []LINE_TYPE
var current_line    LINE_BUFFER
var current_index   int


func Return_SubWindow_Draw_Structures()*LINE_BUFFER{
    
    return &current_line
}


func (r *LINE_BUFFER )compare_buffer( line1,line2 LINE_TYPE ) bool{
    
    for index ,_ := range line1 {
     
        if line1[index] != line2[index] {
            return false
        }
        
    }
    return true
}



func (r *LINE_BUFFER )store_history(){
    
    _, ok := ref_lines[current_line.line]
    if ok == false {
        ref_lines[current_line.line] = current_line
    }
    if len(past_lines) == 0 {
        past_lines = append( past_lines, current_line.line)
    }else{
       
        last_element := past_lines[len(past_lines)-1]
        if r.compare_buffer(current_line.line,last_element ) == false{
            past_lines = append(past_lines,current_line.line)
        }
    }
    

}


func ( r *LINE_BUFFER)initialize_line(rows int,cols int, win *gc.Window, sub_window *gc.Window){
    
    r.Window     = win
    r.sub_window = sub_window
    r.rows       = rows
    r.cols       = cols
    r.cursor = 0
    for index,_ := range r.line{
        r.line[index] = ' '
    }
    r.sub_window.Move(r.rows,r.cursor)
    r.sub_window.ClearToEOL()
    r.add_character('>')
    r.sub_window.Refresh()
    
            
}

func (r *LINE_BUFFER)redraw(){
    r.sub_window.Move(r.rows,r.cursor)
    r.sub_window.ClearToEOL()
    temp := r.cursor
    r.cursor = 0
    for index :=0 ; index < temp;index++ {
      r.add_character(r.line[index])
    }
   
    r.sub_window.Refresh()
}

func (r *LINE_BUFFER)print_string( input string){
    for _,value := range input {
        r.add_character(byte(value))
    }
}

func (r *LINE_BUFFER )add_character(input byte){
   if r.cursor < line_limit {
           
            r.clear_underline()
            
            
            r.line[r.cursor ] = input
            
            r.sub_window.Move(r.rows,r.cursor)
            r.sub_window.AddChar(gc.Char(input))
            r.cursor += 1 
            r.sub_window.Move(r.rows,r.cursor)
            r.set_underline()
            r.sub_window.Move(r.rows,r.cursor)
            r.sub_window.Refresh()
    }
}
    
func (r *LINE_BUFFER )Return_cmd( output string ){    
   r.clear_underline()
   r.store_history()
   
   r.sub_window.Scroll(1)

   r.sub_window.Move(r.rows,0)
   
   r.sub_window.Print(output)
   
   r.sub_window.Refresh()   
  
   r.cursor = 0
    for index,_ := range r.line{
        r.line[index] = ' '
    }
    r.sub_window.Move(r.rows,r.cursor)
    r.sub_window.ClearToEOL()
    r.add_character('>')
    
    r.sub_window.Refresh()     


}

func (r *LINE_BUFFER )backspace(){
    
    
    r.insert_character(' ')
    r.left()
   
    
}

func (r* LINE_BUFFER)left(){
    
    if r.cursor > 1 {
       r.clear_underline()
       r.cursor = r.cursor-1
       r.sub_window.Move(r.rows,r.cursor)
       r.set_underline()
       r.sub_window.Move(r.rows,r.cursor)
       r.sub_window.Refresh()
    }
}
    


func (r* LINE_BUFFER)right(){
    
   if r.cursor < r.cols - 2 {
       r.clear_underline()
       r.cursor = r.cursor+1
       r.sub_window.Move(r.rows,r.cursor)
        
        r.set_underline()
        r.sub_window.Move(r.rows,r.cursor)
        r.sub_window.Refresh()
       
    }  
    
}

func (r *LINE_BUFFER)insert_character(input byte){
    r.sub_window.AddChar(gc.Char(input))
    r.sub_window.Refresh()
}

func (r *LINE_BUFFER)clear_underline(){
    
   ch := r.sub_window.InChar()
   ch = ch &^ gc.A_UNDERLINE
   r.sub_window.AddChar(ch)
    
    
}

func ( r *LINE_BUFFER)set_underline(){

   ch := r.sub_window.InChar()
   ch = ch | gc.A_UNDERLINE
   r.sub_window.AddChar(ch)    
    
}
