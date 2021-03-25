package docker_management
import "github.com/msgpack/msgpack-go" 
import "bytes"
import "site_control.com/redis_support/redis_handlers"
import "fmt"
var display_hash  redis_handlers.Redis_Hash_Struct
var display_error redis_handlers.Redis_Stream_Struct


func initialize_redis_display_output(){
   display_hash = (*Docker_Display_Structures)["DOCKER_DISPLAY_DICTIONARY"].(redis_handlers.Redis_Hash_Struct)
   display_error = (*Docker_Display_Structures)["ERROR_STREAM"].(redis_handlers.Redis_Stream_Struct)
   var data = make(map[string]bool)
   data["active"] = true
   data["status"] = true
   set_docker_display_dictionary("test",&data)

}


func set_docker_display_dictionary(field string,value *map[string]bool){

    // convert bool to msgpack
    var b bytes.Buffer
	
	
	
	
	msgpack.Pack(&b,*value)
	var output = b.String()
	fmt.Println("msgpack output",output)
	display_hash.HSet(field,output)
	
    var input = display_hash.HGet(field)
	var byte_input = []byte(input)
	var c = bytes.NewBuffer(byte_input)
	var input_value,cnt,err = msgpack.Unpack(c) 
    fmt.Println(input_value,cnt,err)
	kv, ok := input_value.Interface().(map[interface{}]interface{})
    fmt.Println(kv,ok)
	var y = make(map[string]bool)
	for key,value := range kv{
	  y[key.(string)] = value.(bool)
	}
	fmt.Println(y)
  
}
/*
func log_docker_error_stream( container string, action string){


}x := map[interface{}]interface{}{"XYZ": "Hello world!"}
	var buf bytes.Buffer
	_, err := Pack(&buf, x)
	if err != nil {
		t.Error(err)
		return
	}
	v, _, err := Unpack(&buf)
	if err != nil {
		t.Error(err)
		return
	}

*/