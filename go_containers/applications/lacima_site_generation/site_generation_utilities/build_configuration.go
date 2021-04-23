package graph_generation

import "fmt"
import "context"
import "encoding/json"
import "strconv"
import "strings"
import "github.com/go-redis/redis/v8"
import  "github.com/golang-collections/collections"


var ctx    = context.TODO()
var client *redis.Client


type Build_Configuration struc {

   sep        string
   rel_sep    string
   label_sep  string
   namespace  string
   keys       *Sets 




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
   return_value.sep     = "["
   return_value.rel_sep = ":"
   return_value.label_sep = "]"
   return_value.namespace = []string
   return_value.keys = sets.New()
   client.FlushDB(ctx)
   
   return return_value
}



func (v *Build_Configuration) 
 
func (v *Build_Configuration) build_namespace( self,name ):
       return_value = copy.deepcopy(self.namespace) 
       return_value.append(name)
       return return_value


func (v *Build_Configuration) pop_namespace( self ):
       del self.namespace[-1]    

func (v *Build_Configuration) add_header_node( self, relation,label=None, properties = {}, json_flag= True ):
     if label== None:
        label = relation
     properties["name"] = label
     self.construct_node( True, relation, label, properties, json_flag )

func (v *Build_Configuration) end_header_node( self, assert_namespace ):
       assert (assert_namespace == self.namespace[-1][0]) ,"miss match namespace  got  "+assert_namespace+" expected "+self.namespace[-1][0]
       del self.namespace[-1]    


func (v *Build_Configuration) check_namespace( self ):
       assert len(self.namespace) == 0, "unbalanced name space, current namespace: "+ json.dumps(self.namespace)
       #print ("name space is in balance")
      
func (v *Build_Configuration)  add_info_node( self, relation,label, properties = {}, json_flag= True ):
     
     
     self.construct_node( False, relation,  label, properties, json_flag )

   # concept of namespace name is a string which ensures unique name
   # the name is essentially the directory structure of the tree
   def construct_node(self, push_namespace,relationship, label,  properties, json_flag = True ):
 

       redis_key, new_name_space = self.construct_basic_node( self.namespace, relationship,label ) 
       if redis_key in self.keys:
            raise ValueError("Duplicate Key")
       self.keys.add(redis_key)
       for i in properties.keys():
           temp = json.dumps(properties[i] )
           self.redis_handle.hset(redis_key, i, temp )
       
       
       if push_namespace == True:
          self.namespace = new_name_space

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
