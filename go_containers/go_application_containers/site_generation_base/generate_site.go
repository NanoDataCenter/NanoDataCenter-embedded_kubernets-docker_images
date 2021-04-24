package main

import "fmt"
import "lacima.com/go_application_containers/site_generation_base/site_generation_utilities"




func main(){

  var properties map[string]interface{}  
  su.Setup_Site_File()
  fmt.Println(su.Site)
  su.Setup_graph_generation()
  
  properties =  make(map[string]interface{})
  properties["test"] = "this is a test"
  su.Bc_Rec.Add_header_node("test_relation","test_label",properties)
  su.Bc_Rec.End_header_node("test_relation","test_label")
  properties =  make(map[string]interface{})
  properties["test_xxx"] = "this is another test"
  su.Bc_Rec.Add_header_node("test_relation","test_label_a",properties)
  
  properties =  make(map[string]interface{})
  properties["test_info"] = "this is info test" 
  su.Bc_Rec.Add_info_node("info_relation","info_label",properties)
  su.Bc_Rec.End_header_node("test_relation","test_label_a")
  su.Bc_Rec.Check_namespace()
   
}


