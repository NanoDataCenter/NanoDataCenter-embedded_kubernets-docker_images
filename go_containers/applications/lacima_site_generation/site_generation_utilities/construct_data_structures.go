package construct_data_structures


type Package_Constructor struc {
  site                string
  bc                  *Build_Constructor
  name                string
  header_node         string
  current_package     string
  current_properties  map[string]interface{}

}



func Construct_Data_Structures( site string , bc *Build_Constructor)Package_Constructor {

   var return_value Package_Constructor
   return_value.site            = site
   return_value.bc              = bc
   return_value.header_node     = "[PACK_CTRL:" +site+ "]"
   return_value.current_package = "" 
   
   return return_value

}

func ( v *Construct_Data_Structures)construct_package( package_name string ){
     if v.current_package != "" {
	     panic("current package not closed")
	 }
	 v.name = package_name
     v.current_package = self.header_node+"[PACK:"+package_name+"]"
     v.current_properties = make(map[string]interface{})
     v.properties["data_structures"]  = make(map[string]interface{})
}


func ( v *Construct_Data_Structures)close_package_contruction(){
        self.current_package = None
        v.bc.add_info_node("PACKAGE",v.name,v.properties)
}

        
func ( v *Construct_Data_Structures)create_sql_table(self,name,database_name,table_name,field_names){
       assert(name not in self.properties )
	   properties := make(map[string]|inteface{}
       properties["type"]  = "SQL_LOG_TABLE"
       properties["name"] = name  
       properties["database_name"] = database_name
       properties["table_name"]  = table_name
       properties["field_names"] = field_names       
       v.properties["data_structures"][name] = properties 
}
        
        
func ( v *Construct_Data_Structures) create_sql_text_search_table(self,name,database_name,table_name,field_names){
       assert(name not in self.properties )
       properties = {}
       properties["type"]  = "SEARCH_SQL_LOG_TABLE"
       properties["name"] = name  
       properties["database_name"] = database_name
       properties["table_name"]  = table_name
       properties["field_names"] = field_names       
       self.properties["data_structures"][name] = properties 
}
       
func ( v *Construct_Data_Structures) add_single_element(self,name,forward=False){
       assert(name not in self.properties )
       properties = {}
       properties["name"] = name
       properties["type"]  = "SINGLE_ELEMENT"
       self.properties["data_structures"][name] = properties 
}
       
func ( v *Construct_Data_Structures) add_managed_hash(self,name,fields,forward=False){
       assert(name not in self.properties )
       properties = {}
       properties["name"] = name
       properties["type"]  = "MANAGED_HASH"
       properties["forward"] =forward
       properties["fields"] = fields
       self.properties["data_structures"][name] = properties 
}
      
func ( v *Construct_Data_Structures) add_hash(self,name,forward=False){
       assert(name not in self.properties )
       properties = {}
       properties["name"] = name
       properties["type"]  = "HASH"
       properties["forward"] =forward
       self.properties["data_structures"][name] = properties 
        
}


func ( v *Construct_Data_Structures) add_redis_stream(self,name,depth=1024,forward=False){
       assert(name not in self.properties )
       properties = {}
       properties["name"] = name
       properties["type"]  = "STREAM_REDIS"
       properties["depth"]  =depth
       properties["forward"] =forward
       self.properties["data_structures"][name] = properties 
}
       

       
func ( v *Construct_Data_Structures)add_job_queue(self,name,depth,forward=False){
       assert(name not in self.properties )
       properties = {}
       properties["name"] = name
       properties["depth"] = depth
       properties["type"]  = "JOB_QUEUE"
       properties["forward"] =forward
       self.properties["data_structures"][name] = properties 
}
      
      
func ( v *Construct_Data_Structures)add_rpc_server(self,name,properties ){
       assert(name not in self.properties )
       properties["name"] = name
       properties["type"]  = "RPC_SERVER"
       self.properties["data_structures"][name] = properties 
}
      
func ( v *Construct_Data_Structures)add_rpc_client(self,name,properties){
       assert(name not in self.properties )
  
       properties["name"] = name
       properties["type"]  = "RPC_CLIENT"
       self.properties["data_structures"][name] = properties
}