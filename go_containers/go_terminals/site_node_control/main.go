package main
import (
	gc "github.com/gbin/goncurses"
    "lacima.com/go_terminals/go_terminal_library"
    "lacima.com/go_terminals/site_node_control/command_utilities"
	//"fmt"
)

   

func main(){
    
    
    gc_support.Init_SoftKey()
    defer gc.End()
   
    initial_screen()
    
	

    
}


func initial_screen(){
    title := "Site and Node Control Terminal"
    message := []string{"F1 Key ","F2 Key ","F3 Key","F4 Key ","F5 Key ","F6 Key","F7 Key ","F8 Key Command Line Utitilies"}
    help_message :=[]string{"this is a test","of the help message"}
    user_key := make([]gc_support.Soft_Key_Def,6)
    user_key[0].Label = "Pop up"
    user_key[0].Func  = pop_key_handler
    user_key[1].Label = "Nested Win"
    user_key[1].Func  = nested_key_handler
    user_key[2].Label = "Menu"
    user_key[2].Func  = menu_key_handler  
    user_key[3].Label = "Single Select"
    user_key[3].Func  =  single_key_handler
    user_key[4].Label = "Confirm"
    user_key[4].Func  = confirm_handler  
    user_key[5].Label = "Util"
    user_key[5].Func  = command_line_utilities.Command_line_launcher  
	gc_support.Construct_default_panel( title , message,help_message,user_key , nil )    
	//gc_support.Pop_up_alert(title,message)
    
}   

func pop_key_handler( input gc.Key )bool{
    
    message  := make([]string,1)
    message[0] = "test message"
    gc_support.Pop_up_alert("popup_test",message)
    return false
    
}

func nested_key_handler( input gc.Key )bool{
    
    message  := make([]string,1)
    message[0] = "nested window"
    gc_support.Construct_default_panel( "nested test" , message,make([]string,0),make([]gc_support.Soft_Key_Def,0) , nil )    
    return false
    
}

func menu_key_handler( input gc.Key ) bool{
    
   menu_screen_test()
   return false
    
}

func menu_screen_test(){
    title := "menu test"
   
    menu_items := []gc_support.Menu_records{
        {"Choice 1","test  gggg" ,false},
        {"Choice 2","test  gggg" ,true},
        {"Choice 3","test  gggg" ,false},
        {"Choice 4","test  gggg" ,true},
        {"Choice 5","test  gggg" ,false},
        {"Choice 6","test  gggg" ,true},
        {"Choice 7","test  gggg" ,false},
        {"Choice 8","test  gggg" ,true},
        {"Choice 9","test  gggg" ,false},
        {"Choice 10","test  gggg" ,true},
        {"Choice 11","test  gggg" ,false},
        {"Choice 12","test  gggg" ,true},
        {"Choice 13","test  gggg" ,false},
        {"Choice 14","test  gggg" ,true},
        {"Choice 15","test  gggg" ,false},
        {"Choice 16","test  gggg" ,true},
        {"Choice 17","test  gggg" ,false},
        {"Choice 18","test  gggg" ,true},
        {"Choice 19","test  gggg" ,false},
        {"Choice 20","test  gggg" ,true},
        {"Choice 21","test  gggg" ,false},
        {"Choice 22","test  gggg" ,true},
        {"Choice 23","test  gggg" ,false},
        {"Choice 24","test  gggg" ,true},
        {"Choice 25","test  gggg" ,false},
        {"Choice 26","test  gggg" ,true},
        {"Choice 27","test  gggg" ,false},
        {"Choice 28","test  gggg" ,true},
        {"Choice 29","test  gggg" ,false},
        {"Choice 30","test  gggg" ,true},
        {"Choice 31","test  gggg" ,false},
        {"Choice 32","test  gggg" ,true},
        {"Choice 33","test  gggg" ,false},
        {"Choice 34","test  gggg" ,true},
        {"Choice 35","test  gggg" ,false},
        {"Choice 36","test  gggg" ,true},
        {"Choice 37","test  gggg" ,false},
        {"Choice 38","test  gggg" ,true},
        {"Choice 39","test  gggg" ,false},
        {"Choice 40","test  gggg" ,true}}
    
	gc_support.Construct_menu_window( title , menu_items,1 ,false)    
	//gc_support.Pop_up_alert(title,message)
}                  
func single_key_handler( input gc.Key ) bool{
    
   single_screen_test()
   return false
    
}

func single_screen_test(){
    title := "Single Select"
   
    menu_items := []gc_support.Menu_records{
        {"Choice 1","test  gggg" ,false},
        {"Choice 2","test  gggg" ,true},
        {"Choice 3","test  gggg" ,false},
        {"Choice 4","test  gggg" ,true},
        {"Choice 5","test  gggg" ,false},
        {"Choice 6","test  gggg" ,true},
        {"Choice 7","test  gggg" ,false},
        {"Choice 8","test  gggg" ,true},
        {"Choice 9","test  gggg" ,false},
        {"Choice 10","test  gggg" ,true},
        {"Choice 11","test  gggg" ,false},
        {"Choice 12","test  gggg" ,true},
        {"Choice 13","test  gggg" ,false},
        {"Choice 14","test  gggg" ,true},
        {"Choice 15","test  gggg" ,false},
        {"Choice 16","test  gggg" ,true},
        {"Choice 17","test  gggg" ,false},
        {"Choice 18","test  gggg" ,true},
        {"Choice 19","test  gggg" ,false},
        {"Choice 20","test  gggg" ,true},
        {"Choice 21","test  gggg" ,false},
        {"Choice 22","test  gggg" ,true},
        {"Choice 23","test  gggg" ,false},
        {"Choice 24","test  gggg" ,true},
        {"Choice 25","test  gggg" ,false},
        {"Choice 26","test  gggg" ,true},
        {"Choice 27","test  gggg" ,false},
        {"Choice 28","test  gggg" ,true},
        {"Choice 29","test  gggg" ,false},
        {"Choice 30","test  gggg" ,true},
        {"Choice 31","test  gggg" ,false},
        {"Choice 32","test  gggg" ,true},
        {"Choice 33","test  gggg" ,false},
        {"Choice 34","test  gggg" ,true},
        {"Choice 35","test  gggg" ,false},
        {"Choice 36","test  gggg" ,true},
        {"Choice 37","test  gggg" ,false},
        {"Choice 38","test  gggg" ,true},
        {"Choice 39","test  gggg" ,false},
        {"Choice 40","test  gggg" ,true}}
    
    
	gc_support.Construct_menu_window( title , menu_items,1 ,true)    
	//gc_support.Pop_up_alert(title,message)
}   




func confirm_handler( input gc.Key )bool{
    title     := "test message"
    message   := []string{"test line1","test lin2","test line3"}
    gc_support.Pop_up_confirmation(title,message)
    return false
}    



