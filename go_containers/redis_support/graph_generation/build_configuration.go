package graph_generation

import "fmt"
import "context"
import "encoding/json"
import "strconv"
import "strings"
import "import "container/list"
import "github.com/go-redis/redis/v8"
import  "github.com/golang-collections/collections"


var ctx    = context.TODO()
var client *redis.Client


type Build_Configuration struct {

   sep                string
   rel_sep            string
   label_sep          string
   namespace_max_len  int
   namespace_len      int
   namespace          [][2]string  // use 
   keys               *sets.Sets 

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



func Construct_build_configuration( ) Build_Configuration {

   var return_value Build_Configuration
   return_value.sep                 = "["
   return_value.rel_sep             = ":"
   return_value.label_sep           = "]"
   return_value.namespace_max_len   = 0
   return_value.namespace_len       = 0
   return_value.keys                = sets.New()
   client.FlushDB(ctx)
   
   return return_value
}



 




func (v *Build_Configuration) Add_header_node( relation,label string, properties map[string]interface{} ){
   if label == "" {
      label = relationship
   }
   properties["name"] = label
   construct_node( True, relation, label, properties )
}


func (v *Build_Configuration)  add_info_node( relation,label string, properties map[string]interface{}  ){
     
   if label == "" {
      label = relationship
   }
   properties["name"] = label
   construct_node( True, relation, label, properties )
}


func (v *Build_Configuration) End_header_node( self, assert_namespace ){

       last_namespace := v.pop_namespace()
	   if last_namespace[0] != assert_namespace {
	      panic("unmatched namespace  expected "+last_namespace+"  got  "+assert_namespace)

}

func (v *Build_Configuration)construct_node( push_namespace bool,relationship ,label string  ,properties map[string]interface{} properties  ){
 

       redis_key, new_name_space = v.construct_basic_node( relationship,label  ) 
	   
       if redis_key in self.keys:
            raise ValueError("Duplicate Key")
       self.keys.add(redis_key)
       for i in properties.keys():
           temp = json.dumps(properties[i] )
           self.redis_handle.hset(redis_key, i, temp )
       
       
       if push_namespace == True:
          self.namespace = new_name_space
}


func construct_node( push_namespace,relationship, label,  properties, json_flag = True ):
 

       redis_key, new_name_space = self.construct_basic_node( self.namespace, relationship,label ) 
       if redis_key in self.keys:
            raise ValueError("Duplicate Key")
       self.keys.add(redis_key)
       for i in properties.keys():
           temp = json.dumps(properties[i] )
           self.redis_handle.hset(redis_key, i, temp )
       
       
       if push_namespace == True:
          self.namespace = new_name_space

//func (v *Build_Configuration) check_namespace( self ){
//       assert len(self.namespace) == 0, "unbalanced name space, current namespace: "+ json.dumps(self.namespace)
//       #print ("name space is in balance")
//}


      



   # concept of namespace name is a string which ensures unique name
   # the name is essentially the directory structure of the tree

func (v *Build_Configuration)  _convert_namespace( self, namespace):
     
       temp_value = []

       for i in namespace:
          temp_value.append(self.make_string_key( i[0],i[1] ))
       key_string = self.sep+self.sep.join(temp_value)
 
       return  key_string
  
func (v *Build_Configuration) construct_basic_node( self, namespace, relationship,label ){
       
       new_name_space = copy.copy(namespace)
       new_name_space.append( [ relationship,label ] )
       
       redis_string =  self._convert_namespace(new_name_space)

       self.redis_handle.hset(redis_string,"namespace",json.dumps(redis_string))
       self.redis_handle.hset(redis_string,"name",json.dumps(label))
       self.update_terminals( relationship, label, redis_string)
       self.update_relationship( new_name_space, redis_string )
       return redis_string, new_name_space
}

func (v *Build_Configuration)  make_string_key( relationship,label string)string{

       return relationship+v.rel_sep+label+v.label_sep
}

 
func (v *Build_Configuration)  update_relationship(  new_name_space, redis_string string ){
       for relationship,label := range new_name_space{
           client.Sadd(ctx,"@RELATIONSHIPS",relationship)
           client.Sadd(ctx,"%"+relationship,redis_string)
           client.Sadd(ctx,"#"+relationship+self.rel_sep+label,redis_string)
		}
}


func (v *Build_Configuration)update_terminals( relationship ,label, redis_string string ){
       client.Sadd(ctx,"@TERMINALS",relationship)
       client.Sadd(ctx,"&"+relationship,redis_string)
       client.Sadd(ctx,"$"+relationship+self.rel_sep+label,redis_string)
}

 
func (v *Build_Configuration)  store_keys( ){
    for i,_ := range v.keys {
       client.Sadd(ctx,"@GRAPH_KEYS", i )
	}
}       
