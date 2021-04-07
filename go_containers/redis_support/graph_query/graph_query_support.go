 package graph_query



import "fmt"
import "context"
import "encoding/json"
import "strconv"
import "strings"
import "github.com/go-redis/redis/v8"

 
 
//type site_data_type map[string]interface{}

type Graph_data map[string]string

type query_element map[string]string
type query_type []query_element

 
var site_data     map[string]interface{}
var sep           string  = "["
var rel_sep       string = ":"
var label_sep     string = "]"
var namespace     []string
var site          string  = ""

var ctx    = context.TODO()
var client *redis.Client

func Convert_json_interface_array( json_string string) []map[string]interface{}  {

     var return_value = make([]map[string]interface{},0)
     var err2 = json.Unmarshal([]byte(json_string),&return_value)
     if err2 != nil{
	         panic("bad json data")
	  }
     return return_value  

}
func Convert_json_dictionary_interface( json_string string) map[string]interface{}  {

     var return_value = make(map[string]interface{},0)
     var err2 = json.Unmarshal([]byte(json_string),&return_value)
     if err2 != nil{
	         panic("bad json data")
	  }
     return return_value  

}

func Convert_json_string_array( json_string string)[]string{

     var return_value = make([]string,0)
     var err2 = json.Unmarshal([]byte(json_string),&return_value)
     if err2 != nil{
	         panic("bad json data")
	  }
     return return_value  

}

func Convert_json_string( json_string string) string{
   
     var return_value string;
     var err2 = json.Unmarshal([]byte(json_string),&return_value)
     if err2 != nil{
	         panic("bad json data")
	  }
     return return_value

}
 
func Graph_support_init(sdata *map[string]interface{}) {
    site_data = *sdata
	site = site_data["site"].(string)
    var address =  site_data["host"].(string)
    var port = 	int(site_data["port"].(float64))
	var address_port = address+":"+strconv.Itoa(port)
	client = redis.NewClient(&redis.Options{
                                                 Addr: address_port,
												
												 DB: 3,
                                               })
	err := client.Ping(ctx).Err();     
	if err != nil{
	         panic("redis graph connection")
	 }
    fmt.Println("redis graph ping")	
}	

func Common_package_search( site *string, search_list *[]string) []map[string]string{
   var query_list = make([]query_element,0)
   //fmt.Println("len search list",len(*search_list))
   add_match_relationship(&query_list,"SITE",site_data["site"].(string))
   //fmt.Println("building query list",len(query_list))
   for i :=0; i <len(*search_list)-1;i++{
      var search_term = (*search_list)[i]

	  var search_list = parse_search_list(search_term)
	  add_match_relationship(&query_list,search_list[0],search_list[1])
	  //fmt.Println("building query list",len(query_list))
   }
   
   var search_list_term = parse_search_list((*search_list)[len(*search_list)-1])
   
   add_match_terminal(&query_list, "PACKAGE",search_list_term[0])
   //fmt.Println("read to match query list",len(query_list))
   return match_list(&query_list)
}

func Common_qs_search(search_list *[]string)[]map[string]string{

   var query_list = make([]query_element,0)
   //fmt.Println("len search list",len(*search_list))
   add_match_relationship(&query_list,"SITE",site_data["site"].(string))
   //fmt.Println("building query list",len(query_list))
   for i :=0; i <len(*search_list)-1;i++{
      var search_term = (*search_list)[i]

	  var search_list = parse_search_list(search_term)
	  add_match_relationship(&query_list,search_list[0],search_list[1])
	  //fmt.Println("building query list",len(query_list))
   }
   
   var search_list_term = parse_search_list((*search_list)[len(*search_list)-1])
   
   add_match_terminal(&query_list,search_list_term[0],search_list_term[1])
   //fmt.Println("read to match query list",len(query_list))
   return match_list(&query_list)
}
 
func parse_search_list(search_term string)[2]string{
    var return_value [2]string
    var result = strings.Split(search_term,":")
    if len(result) <2{
	  return_value[0] = result[0]
	  return_value[1] = ""
	}else{
	  return_value[0] = result[0]
	  return_value[1] = result[1]
	}
	//fmt.Println(return_value)
	return return_value
    
}

func add_match_relationship(query_list *[]query_element, relationship string, label string){
       //fmt.Println("match terminal",relationship,label)
       var temp = make(map[string]string)
       temp["relationship"] = relationship
       temp["label"]        = label
       temp["type"] = "RELATIONSHIPS"
       *query_list= append(*query_list,temp)
       
}

func add_match_terminal(query_list *[]query_element, relationship string, label string){
       //fmt.Println("match terminal",relationship,label)
       var temp = make(map[string]string)
       temp["relationship"] = relationship
       temp["label"]        = label
       temp["type"] = "MATCH_TERMINAL"
       *query_list = append(*query_list,temp)

}


func match_list( query_list  *[]query_element)[]map[string]string{
  
  //fmt.Println("start match")
  var starting_set,_ =  client.SMembersMap(ctx , "@GRAPH_KEYS").Result() 
  //fmt.Println("initial set", len(starting_set))
  //fmt.Println("query_list",len(*query_list))
  for _,value := range *query_list{
     //fmt.Println("iteration",value)
     if value["type"] == "MATCH_TERMINAL"{
	   //fmt.Println("termainal path")
       starting_set = match_terminal_relationship( value["relationship"], value["label"] , starting_set)
	   //fmt.Println(len(starting_set))
      }else{
	     //fmt.Println("relation path")
	     starting_set = match_relationship(value["relationship"], value["label"] , starting_set )
		 //fmt.Println(len(starting_set))
         } 
      if len(starting_set) == 0 {
	       //fmt.Println("early exit")
           return make([]map[string]string,0)
	   }
  		 
     }            
  //fmt.Println("done",len(starting_set))        
  var return_value = return_data(starting_set)
  return return_value
}


 


func match_relationship( relationship, label string , starting_set  map[string]struct{} ) map[string]struct{}{
       var return_value = make( map[string]struct{},0)
       
       if label == ""{
          
          if flag,_ :=client.SIsMember(ctx , "@RELATIONSHIPS", relationship).Result(); flag==true{
              return_value,_ = client.SMembersMap(ctx ,"%"+relationship).Result()
              return_value = intersection(return_value,starting_set)
			}
       }else{   
             if flag,_:=client.SIsMember(ctx , "@RELATIONSHIPS", relationship).Result();flag==true{
			    //fmt.Println("pass relationships")
                if flag1,_:=client.Exists(ctx,  "#"+relationship+rel_sep+label).Result();flag1!=0{
				    //fmt.Println("pass exists")
                    return_value,_ = client.SMembersMap(ctx ,"#"+relationship+rel_sep+label).Result()
					//fmt.Println("member length",len(return_value))
                    return_value = intersection(return_value,starting_set)
			    }
			}
	   }		
       return return_value

}

 
	   
func match_terminal_relationship( relationship, label string , starting_set  map[string]struct{} ) map[string]struct{}{
       var return_value = make( map[string]struct{},0)
       
       if label == ""{
          
          if flag,_ :=client.SIsMember(ctx ,  "@TERMINALS", relationship).Result();flag==true{
              return_value,_ = client.SMembersMap(ctx ,"&"+relationship).Result()
              return_value = intersection(return_value,starting_set)
			}
       }else{   
             if flag,_ :=client.SIsMember(ctx ,  "@TERMINALS", relationship).Result();flag==true{
			     //fmt.Println("terminal")
                if flag1,_:=client.Exists(ctx,  "$"+relationship+rel_sep+label).Result();flag1 != 0{
				    //fmt.Println("pass exists")
                    return_value,_ = client.SMembersMap(ctx ,"$"+relationship+rel_sep+label).Result()
                    return_value = intersection(return_value,starting_set)
			    }
			}
	   }		
       return return_value

}






func return_data( key_set  map[string]struct{})[]map[string]string{
       var return_value = make([]map[string]string,0)
       for i,_ := range key_set{
           var data,_ = client.HGetAll(ctx ,i).Result()
           return_value = append(return_value,data)
	   }
	   return return_value
}

func intersection(new_set, starting_set  map[string]struct{})map[string]struct{}{
     var return_value = make( map[string]struct{},0)
     for i,_ := range new_set {
	       var _, found = starting_set[i]
		   if found{
		     return_value[i] = starting_set[i]
		   }
		
     }
	 return return_value
}		


