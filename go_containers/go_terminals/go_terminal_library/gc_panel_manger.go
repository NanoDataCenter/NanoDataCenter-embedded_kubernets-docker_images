
package gc_support


import (
	gc "github.com/gbin/goncurses"
      "container/list"
	//"fmt"
)

var panel_list *list.List

var Stdscr *gc.Window



func  init_panel_handler() {
    panel_list = list.New()
}

func kill_current_panel(){
    active := panel_list.Front()
    panel_list.Remove(active)
    if panel_list.Len() > 0 {
        panel := active.Value.(*gc.Panel)
        panel.Top()
        panel.Delete()
        panel.Window().Delete()
        active := panel_list.Front()
        panel = active.Value.(*gc.Panel)
        panel.Top()
        soft_key_populate_labels()
        
    }
    
    
}

func Pop_up_confirmation(title string,message[]string)bool{
   
   rows, cols := Stdscr.MaxYX()
   
    
    
    
     w  := len(title)+10
    
    
    
    window, _ := gc.NewWindow(rows,cols,0,0)
    window.Box(0, 0)
    window.ColorOn(POP_UP_COLOR)
    window.SetBackground(gc.ColorPair(POP_UP_COLOR))
    blank_soft_keys()
    window.MoveAddChar(2, 0, gc.ACS_LTEE)
    window.HLine(2, 1, gc.ACS_HLINE, cols-2)
    window.MoveAddChar(2, cols-1, gc.ACS_RTEE)
    window.MoveAddChar(2, 0, gc.ACS_LTEE)
    window.MovePrint(1, (w/2)-(len(title)/2), title)
    
     for index,_ := range message {
         window.MovePrint(3+index, 3, message[index])
    }   
    
    
    window.HLine(rows-3, 1, gc.ACS_HLINE, cols-2)
    window.MoveAddChar(rows-3,cols-1, gc.ACS_RTEE)
    window.MoveAddChar(rows-3, 0, gc.ACS_LTEE)
    

    window.MovePrint(rows -2,3,"F8 to confirm  and F1 reject ")
    panel := gc.NewPanel(window)
    panel.Top()
    panel_list.PushFront(panel)
    gc.Cursor(0)
    gc.UpdatePanels()
    return_value := true
    
    for true {
      gc.UpdatePanels()
      gc.Update()
      ch := Stdscr.GetChar()
      if ch == gc.KEY_F1 {
          return_value = false
          break
      }
      if ch == gc.KEY_F8 {
          return_value = true
          break
      }          

    
    }
    kill_current_panel()
    return return_value
    
}



func Pop_up_alert(title string,message[]string){
    rows, cols := Stdscr.MaxYX()

    h := 4 + len(message)
    w := calculate_message_length(message)+5
    title_length := len(title)+10
    if  title_length > w {
        w = title_length
    }
    blank_soft_keys()
    
    y := (rows-h)/2
    x := (cols-w)/2
    window, _ := gc.NewWindow(h, w, y, x)
    window.Box(0, 0)
    window.ColorOn(POP_UP_COLOR)
    window.SetBackground(gc.ColorPair(POP_UP_COLOR))
    
    window.MoveAddChar(2, 0, gc.ACS_LTEE)
    window.HLine(2, 1, gc.ACS_HLINE, w-2)
    window.MoveAddChar(2, w-1, gc.ACS_RTEE)
    window.MovePrint(1, (w/2)-(len(title)/2), title)
    for index,_ := range message {
         window.MovePrint(3+index, (w/2)-(len(message[index])/2), message[index])
    }
    panel := gc.NewPanel(window)
    panel.Top()
    panel_list.PushFront(panel)
    gc.UpdatePanels()
    gc.Update()
    gc.Cursor(0)
    Stdscr.GetChar()
 
    kill_current_panel()
    
}

func calculate_message_length( input []string )int{
    var return_value int
    return_value = 0
    for _, value := range input {
        temp := len(value)
        if temp > return_value {
            return_value = temp
        }
    }
    return return_value
}


func Construct_default_panel( title string, message []string,help []string, soft_key_constructors []Soft_Key_Def, user_handler Soft_Key_Function_Handler ){
   
   rows, cols := Stdscr.MaxYX()
    
    blank_soft_keys()
    Setup_SoftKey(soft_key_constructors,help)
    w := calculate_message_length(message)
    if len(title)> len(message) {
        w = len(title)+10
    }
    
    
    window, _ := gc.NewWindow(rows,cols,0,0)
    window.ScrollOk(true)
    window.Box(0, 0)
    window.ColorOn(NORMAL_WINDOW_COLOR)
    window.SetBackground(gc.ColorPair(NORMAL_WINDOW_COLOR))
   
    window.MoveAddChar(2, 0, gc.ACS_LTEE)
    window.HLine(2, 1, gc.ACS_HLINE, cols-2)
    window.MoveAddChar(2, cols-1, gc.ACS_RTEE)
    window.MovePrint(1, (w/2)-(len(title)/2), title)
    for index,_ := range message {
         window.MovePrint(3+index, 3, message[index])
    }
    panel := gc.NewPanel(window)
    panel.Top()
    panel_list.PushFront(panel)
    gc.Cursor(0)
    gc.UpdatePanels()
    gc.Update()
    Key_Stroke_Handler( user_handler )
   
    kill_current_panel()
    
    
}
type Menu_records struct {
    Name   string
    Description string
    State bool
    
}

func Construct_menu_window( title string,menu_items []Menu_records, ncols int, single_select bool )(bool,[]Menu_records){
   
   rows, cols := Stdscr.MaxYX()
   
    
    
    
     w  := len(title)+10
    
    
    
    window, _ := gc.NewWindow(rows,cols,0,0)
    window.Box(0, 0)
    window.ColorOn(POP_UP_COLOR)
    window.SetBackground(gc.ColorPair(POP_UP_COLOR))
   
    window.MoveAddChar(2, 0, gc.ACS_LTEE)
    window.HLine(2, 1, gc.ACS_HLINE, cols-2)
    window.MoveAddChar(2, cols-1, gc.ACS_RTEE)
    window.MoveAddChar(2, 0, gc.ACS_LTEE)
    window.MovePrint(1, (w/2)-(len(title)/2), title)
    
    
    window.HLine(rows-3, 1, gc.ACS_HLINE, cols-2)
    window.MoveAddChar(rows-3,cols-1, gc.ACS_RTEE)
    window.MoveAddChar(rows-3, 0, gc.ACS_LTEE)
    
    if single_select == false {
        window.MovePrint(rows -2,3,"Press Enter to toggle state  F8 to return  and F1 to abort")
    }else{
         window.MovePrint(rows -2,3,"Press Enter to Select/return and F1 to abort")
    }
    panel := gc.NewPanel(window)
    panel.Top()
    panel_list.PushFront(panel)
    gc.Cursor(0)
    gc.UpdatePanels()
    gc.Update()
    status, data := create_menu(window,menu_items,ncols,single_select) 
    kill_current_panel()
    return status,data
    
}




func create_menu( menuwin *gc.Window,menu_records []Menu_records, ncols int, single_select bool )(bool,[]Menu_records){


    blank_soft_keys()
 	items := make([]*gc.MenuItem, len(menu_records))
	for i, val := range menu_records  {
		items[i], _ = gc.NewItem(val.Name, val.Description)
        
		defer items[i].Free()
	}

	// create the menu
	menu, _ := gc.NewMenu(items)
	defer menu.Free()

    
    
    rows, cols := Stdscr.MaxYX()
    rows = rows

	menu.SetWindow(menuwin)
	dwin := menuwin.Derived(rows-5, cols-10, 3, 3)
	menu.SubWindow(dwin)
    
 
    menu.SetForeground(gc.ColorPair(MENU_START) | gc.A_REVERSE)
	menu.SetBackground(gc.ColorPair(MENU_START+1) | gc.A_BOLD)
	menu.Grey(gc.ColorPair(MENU_START+2) | gc.A_BOLD)
    //O_ONEVALUE   = C.O_ONEVALUE   // Only one item can be selected
	//O_SHOWDESC   = C.O_SHOWDESC   // Display item descriptions
	//O_ROWMAJOR   = C.O_ROWMAJOR   // Display in row-major order
	//O_IGNORECASE = C.O_IGNORECASE // Ingore case when pattern-matching
	//O_SHOWMATCH  = C.O_SHOWMATCH  // Move cursor to item when pattern-matching
	//O_NONCYCLIC  = C.O_NONCYCLIC  // Don't wrap next/prev item
 
    menu.Option(gc.O_SHOWDESC,true)
    if single_select == true {
       menu.Option(gc.O_ONEVALUE,true)
    }else{
        menu.Option(gc.O_ONEVALUE,false)
    }
   
	menu.Format(rows-7, ncols)
	menu.Mark(" * ")

	
	
	

	menu.Post()
	defer menu.UnPost()
	menuwin.Refresh()
    
    for i, val := range menu.Items()  {
        if single_select == false{    
		  val.SetValue(menu_records[i].State)
        }else{
            val.SetValue(false)
        }
        

	}
    menuwin.Refresh()
    
    
    
	for {
        gc.UpdatePanels()
		gc.Update()
        ch := Stdscr.GetChar()
        
		switch ch {
		case gc.KEY_F1:
			return false, make([]Menu_records,0)
        case gc.KEY_RETURN, gc.KEY_ENTER:
            if single_select == false {
			  menu.Driver(gc.REQ_TOGGLE)
            }else{
                current_value :=  menu.Current(nil)
                return_value := make([]Menu_records,0)
                var temp Menu_records
                temp.Name = (*current_value).Name()
                temp.Description = (*current_value).Description()
                temp.State = true
                return_value = append(return_value,temp)
                return true ,return_value
            }
            
		case gc.KEY_F8:
           if single_select == false{
              return_value := make([]Menu_records,0)
              for index, item := range menu.Items() {
                 if item.Value() {
                    menu_records[index].State = true
                    return_value = append(return_value,menu_records[index])
                  }else{
                     menu_records[index].State = false
                  }
                }
             return true,return_value
           }
    

		default:
			menu.Driver(gc.DriverActions[ch])
		}
	}
	
}   





