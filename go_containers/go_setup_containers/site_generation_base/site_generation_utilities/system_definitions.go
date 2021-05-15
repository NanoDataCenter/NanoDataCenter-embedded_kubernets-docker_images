package su // site_utilities

/*
 *  Support utilities for site generation
 * 
 */



type System_graph_generation func()



type system_definition struct {
   name               string
   system_flag        bool
   processor          string
   containers         []string
   graph_generation   System_graph_generation
}

var system_map      map[string]system_definition
var system_list      []system_definition
var processor_set    map[string]bool
var container_set    map[string]bool


func init_system_generation(){
  
    system_map    = make(map[string]system_definition)
    system_list    = make([]system_definition,0)
    processor_set  = make(map[string]bool)  // make sure that processor is defined and only one
    container_set  = make(map[string]bool)  // make sure a container is defined and used only once
   
}



func expand_container_definitions(){
   for _,element := range system_list {
      register_containers(element.containers)
   }
       
}

func expand_subsystem_definitions(){
 
    for _,element := range system_list {
        element.graph_generation()
    }
    
    
}

func find_system_containers(system_flag bool, processor string )[]string{
   return_value := make([]string,0)
   for _,element := range system_list {
      if element.system_flag != system_flag{
          continue
      }
      if system_flag == true {
          return_value = add_containers(return_value,element.containers)
      }else if processor == element.processor {
          return_value = add_containers(return_value,element.containers)
      }
       
   }
   return return_value
}

func add_containers( input []string, new_elements []string )[]string {
    for _,element := range new_elements {
        input = append(input,element)
    }
    return input
}


func Add_processor( processor_name string ){
 
    check_for_duplicate_processor(processor_name)
    processor_set[processor_name] = true
}

func Construct_system_def(system_name string,system_flag bool, processor_name string, containers []string, graph_generation   System_graph_generation){
 
    var system_element system_definition
    check_for_duplicate_system(system_name)
    register_system_containers(containers)
    if system_flag == true {
      system_element.name   = system_name      
      system_element.system_flag   = true
      system_element.processor   = ""
      system_element.containers    = containers
      system_element.graph_generation = graph_generation
        
    }else{
       check_for_existing_processor(processor_name) 
       system_element.name          = system_name   
       system_element.system_flag   = true
       system_element.processor     = processor_name
       system_element.containers    = containers
       system_element.graph_generation = graph_generation
        
    }
    system_map[system_name] = system_element
    system_list = append(system_list,system_element)
    
    
}  


func Generate_System_Container_list()[]string{
    
   return_value := make([]string,0)
   for _,system_element := range system_list{
        if system_element.system_flag == true {
            for _,container := range system_element.containers{
               return_value = append(return_value,container)
            }
        }     
    }
    return return_value
    
}

func Generate_Processor_Container_list(processor_name string )[]string{
   return_value := make([]string,0)
   for _,system_element := range system_list{
        if system_element.processor == processor_name {
            for _,container := range system_element.containers{
               return_value = append(return_value,container)
            }
        }     
    }
    return return_value
    
}

/*
 *   This code generates the graph System_definition
 *   for all the registered system
 * 
 * 
 */


func Construct_graph_definitions(){
   for _,system_element := range system_list{
        system_element.graph_generation()    
    }    
    
    
}


/*
 * 
 * 
 * 
 */




func check_for_duplicate_processor( processor_name string){
    if _,ok := processor_set[processor_name]; ok== true {
       panic("duplicate processor")
    }    
    
}

func check_for_existing_processor( processor_name string ){
    
     if _,ok := processor_set[processor_name]; ok== true {
       panic("processor not defined")
    }
}    

func check_for_duplicate_system( system_name string){
     if _,ok := system_map[system_name]; ok== true {
       panic("duplicate system")
    }     
    
}

func check_for_duplicate_container( container string ){
    
    
    if _,ok := container_set[container]; ok== true {
       panic("duplicate container")
    }      
    
}

func register_system_containers(containers []string){
    
    for _,container := range containers{
        check_for_duplicate_container(container)
        container_set[container] = true
        
    }
    
    
    
}














