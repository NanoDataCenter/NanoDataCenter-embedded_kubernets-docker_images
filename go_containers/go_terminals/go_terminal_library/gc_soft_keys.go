package gc_support



import (
	gc "github.com/gbin/goncurses"
    "container/list"
	//"fmt"
)


type Soft_Key_Function_Handler func( input gc.Key )bool

type Soft_Key_Def struct {
    Label string
    Func  Soft_Key_Function_Handler
}


type Soft_key_handlers_type struct {
  
    Handler [8]Soft_Key_Def
    help    []string
    
}

var soft_key_current_entry Soft_key_handlers_type
var soft_key_current_queue *list.List


func Init_SoftKey(  ){
  
  gc.SlkInit(gc.SLK_323  )  
  temp, _ := gc.Init()
  StartColor() 
  Stdscr = temp
  gc.Echo(false)
  Stdscr.Keypad(true)
  soft_key_current_queue = list.New()
  init_panel_handler() 
  initialize_softkey_entry()
}

func Soft_key_return_function( input gc.Key )bool{
    return true
}

func NULL_key_return_function( input gc.Key )bool{
    return false
}

func Soft_key_help_function( input gc.Key )bool {
    
     Pop_up_alert("Help Information",soft_key_current_entry.help)
     return false
    
}

func initialize_softkey_entry(){
    
    for index,_ := range soft_key_current_entry.Handler{
        if index == len(soft_key_current_entry.Handler) -2 {
          soft_key_current_entry.Handler[index].Label = "HELP"
          soft_key_current_entry.Handler[index].Func = Soft_key_help_function
        }else if index == len(soft_key_current_entry.Handler)-1{
          soft_key_current_entry.Handler[index].Label = "RETURN"
          soft_key_current_entry.Handler[index].Func = Soft_key_return_function
        }else{
          soft_key_current_entry.Handler[index].Label = ""
          soft_key_current_entry.Handler[index].Func = NULL_key_return_function
        }
            
    }
    
    
}

func blank_soft_keys(){
    
    
    for index,_ := range soft_key_current_entry.Handler{
        
		gc.SlkSet(1+index, "", gc.SLK_CENTER)
        gc.SlkSet(1+index, "", gc.SLK_CENTER)
	}

	gc.SlkColor(SOFT_KEY_COLOR )
	gc.SlkNoutRefresh()
    
	Stdscr.Refresh()        
    
    
}

func soft_key_populate_labels(){
    
    
    for index,i := range soft_key_current_entry.Handler{
        
		gc.SlkSet(1+index, "", gc.SLK_CENTER)
        gc.SlkSet(1+index, i.Label, gc.SLK_CENTER)
	}

	gc.SlkColor(SOFT_KEY_COLOR )
	gc.SlkNoutRefresh()
    
	Stdscr.Refresh()    

}

 
    


func Setup_SoftKey(input []Soft_Key_Def, help[]string){
    if len(input) > 6 {
        panic("bad number of soft keys")
    }
    initialize_softkey_entry()
    
    for index,i := range input {
      soft_key_current_entry.Handler[index] = i
    }
    soft_key_current_entry.help = help
    
      
    for index,i := range soft_key_current_entry.Handler{
        
		gc.SlkSet(index+1, "", gc.SLK_CENTER)
        gc.SlkSet(index+1, i.Label, gc.SLK_CENTER)
	}
    
	gc.SlkColor(SOFT_KEY_COLOR )
	gc.SlkNoutRefresh()
	Stdscr.Refresh()  
    soft_key_current_queue.PushFront(soft_key_current_entry)
    
}    
        
        
   

func Key_Stroke_Handler( user_handler Soft_Key_Function_Handler){
    
    for {
		gc.UpdatePanels()
		gc.Update()
		ch := Stdscr.GetChar()
        if user_handler != nil {
            if user_handler(ch) == true {
                break 
            }
        }
		if soft_key_function_handler(ch) == true{
            break
        }
        if ch == 'q' {
            break
        }
        
	}   
	active := soft_key_current_queue.Front()
    soft_key_current_queue.Remove(active)
   
    if soft_key_current_queue.Len() == 0 {
        return
    }
    active = soft_key_current_queue.Front()
    soft_key_current_entry = active.Value.(Soft_key_handlers_type)
    soft_key_populate_labels()
    
}


func soft_key_function_handler( input gc.Key ) bool {
    var return_value bool
    switch input {
        
        case gc.KEY_F1:
            
            return_value = soft_key_current_entry.Handler[0].Func(input)
            
            
        case gc.KEY_F2:
            return_value = soft_key_current_entry.Handler[1].Func(input)
            
        case gc.KEY_F3:
            return_value = soft_key_current_entry.Handler[2].Func(input)
            
        case gc.KEY_F4:
            return_value = soft_key_current_entry.Handler[3].Func(input)
            
        case gc.KEY_F5:
            return_value = soft_key_current_entry.Handler[4].Func(input)
            
        case gc.KEY_F6:
            return_value = soft_key_current_entry.Handler[5].Func(input)
            
        case gc.KEY_F7:
            return_value = soft_key_current_entry.Handler[6].Func(input)
            
        case gc.KEY_F8:
            return_value = soft_key_current_entry.Handler[7].Func(input)
            
       
       
        
        default:
            return_value = false
    }
    return return_value
         
        
}        

     

            
            
    
    

/*
 * 
 * 
 * 
 * SLK_323        SlkFormat = iota // 8 labels; 3-2-3 arrangement
	SLK_44                          // 8 labels; 4-4 arrangement
	SLK_PC444                       // 12 labels; 4-4-4 arrangement
	SLK_PC444INDEX   
*/
