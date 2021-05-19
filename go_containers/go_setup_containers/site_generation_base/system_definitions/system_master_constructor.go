package sys_defs


/*
 * This is the generic part of building a system
 * 
 * 
 * 
 * 
 */

type system_dict_type func(bool, string )

var system_dict map[string]system_dict_type

func Build_System_Catalog() {
    
    system_dict                     = make(map[string]system_dict_type)
    system_dict["system_component"] = generate_system_components 
    system_dict["tp_managed_switch"] = generate_tp_monitored_switches
    // other components will come
    
    
    
    
    
    
    
}


func Add_Component(system_flag bool,processor string,component_name string){

    check_system_components(component_name )
    system_dict[component_name](system_flag,processor)    
    
}



  



func check_system_components( system_component string ){
    if _,ok := system_dict[system_component]; ok == false{
        panic("non existant compontent "+system_component)
    }
}
