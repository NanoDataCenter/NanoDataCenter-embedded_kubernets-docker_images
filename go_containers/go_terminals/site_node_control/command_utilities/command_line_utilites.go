package command_line_utilities

import (
	gc "github.com/gbin/goncurses"
    "lacima.com/go_terminals/go_terminal_library"
	//"fmt"
)








func Command_line_launcher( input gc.Key) bool {
    
    

    gc_support.Construct_Console("Command Console",commmand_handler)
    return false
    
    
    
}



func commmand_handler(input string )string{
    return "lin2 1 \nline 2 \n"
}


