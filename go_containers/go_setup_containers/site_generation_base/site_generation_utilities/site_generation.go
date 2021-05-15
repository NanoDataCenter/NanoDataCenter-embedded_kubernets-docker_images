package su

var working_site string



func Construt_Site(site string){
    system_containers,startup_containers := determine_system_containers()
    start_site_definitions(site , system_containers, startup_containers)
    expand_subsystem_definitions()
    add_containers_to_graph()
    add_processors_to_graph() 
    end_site_definitions(site)
}



func end_site_definitions(site_name string){

   Bc_Rec.End_header_node( "SITE",site_name )

}




func start_site_definitions(site_name string, system_containers, startup_containers   []string){
    
    working_site = site_name
    properties := make(map[string]interface{})
	properties["startup_containers"] = startup_containers
	properties["containers"] = system_containers
    Bc_Rec.Add_header_node( "SITE",site_name,  properties  )
}


func determine_system_containers()([]string,[]string) {
  return_value_1 := make([]string,0)
  return_value_2 := make([]string,0)
  system_containers :=  find_system_containers(true,"")
  for _,container := range system_containers {
     temp := container_map[container]
     if temp.temporary == true{
         return_value_2 = append(return_value_2,container)
     }else{
        return_value_1 = append(return_value_1,container)
     }
      
  }
  return return_value_1,return_value_2
}
    
    
    
func add_containers_to_graph(){
  Bc_Rec.Add_header_node( "CONTAINER_LIST","CONTAINER_LIST",make(map[string]interface{}) )  
  expand_container_definitions()
  Bc_Rec.End_header_node( "CONTAINER_LIST","CONTAINER_LIST" )  
}


func add_processors_to_graph(){
   Bc_Rec.Add_header_node( "PROCESSOR_LIST","PROCESSOR_LIST",make(map[string]interface{})  )  
   for processor,_ := range processor_set {
       
       containers := determine_processor_containers(processor)
        Construct_processor(processor, containers)
       
       
   }
   Bc_Rec.End_header_node( "PROCESSOR_LIST","PROCESSOR_LIST")     
    
}

func determine_processor_containers(processor string)[]string {
  return_value_1 := make([]string,0)
  system_containers :=  find_system_containers(false,"")
  for _,container := range system_containers {
     temp := container_map[container]
     if temp.temporary == true{
         panic("temporary containers can only be assigned to system not processor container "+container)
     }else{
        return_value_1 = append(return_value_1,container)
     }
      
  }
  return return_value_1
}
