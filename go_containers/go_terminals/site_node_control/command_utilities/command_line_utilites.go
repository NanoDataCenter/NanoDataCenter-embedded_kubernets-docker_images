package command_line_utilities

import (
	//gc "github.com/gbin/goncurses"
    "lacima.com/go_terminals/go_terminal_library"
	//"fmt"
)








func Command_line_launcher( ) {
    
    

    gc_support.Construct_Console("Command Console",commmand_handler)
    
    
    
    
}



func commmand_handler(input string )string{
    return "lin2 1 \nline 2 \n"
}


