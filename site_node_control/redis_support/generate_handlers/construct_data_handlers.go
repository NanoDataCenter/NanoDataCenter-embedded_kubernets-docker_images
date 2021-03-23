package data_handler
import "fmt"
//import "reflect"
import "site_control.com/redis_support/graph_query"

var site string

func Data_handler_init( site_data *map[string]interface{}){

   site = (*site_data)["site"].(string)
   fmt.Println("site",site)
}



func Construct_handler_definitions( search_list *[]string ) { //[]map[string]interface{} {

   
   var packages = graph_query.Common_package_search(&site,search_list)
   ///fmt.Println("packages",len(packages),packages) 
   var data_structures_json = packages[0]["data_structures"]
   fmt.Println(data_structures_json)
   var data_structures = graph_query.Convert_json_dictionary_interface(  data_structures_json)
   
   var namespace_json = packages[0]["namespace"]
   var namespace = graph_query.Convert_json_string( namespace_json)
   var new_dictionary = make(map[string]interface{},0)
   
   for i,v := range data_structures{
     var k = v.(map[string]interface{}) 
     var key = namespace +"["+k["type"].(string)+":"+k["name"].(string) +"]"
	 k["key"]= key
	
	 new_dictionary[i] = k
   }

   for i,v := range new_dictionary{
     fmt.Println(i,v)
   }
 
}





//key = self.package["namespace"]+"["+data["type"]+":"+data["name"] +"]"