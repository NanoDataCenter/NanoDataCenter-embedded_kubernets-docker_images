package main

import "fmt"
import "bytes"
import "io/ioutil"
import "io/fs"
import "os"
import "lacima.com/site_data"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/redis_support/generate_handlers"
import  "lacima.com/Patterns/msgpack"
import "github.com/msgpack/msgpack-go"




func main(){

    //mount_usb_drive()
      
		
	var config_file = "/data/redis_server.json"
	var site_data_store map[string]interface{}

	site_data_store = get_site_data.Get_site_data(config_file)
    graph_query.Graph_support_init(&site_data_store)
	redis_handlers.Init_Redis_Mutex()
	data_handler.Data_handler_init(&site_data_store)	
 	  
    var search_list = []string{"FILE_SERVER","FILE_SERVER"}

    var handlers = data_handler.Construct_Data_Structures(&search_list)
    fmt.Println(handlers)
    panic("done")
	
	
	//driver.Add_Handler( "load",load_file)
	//driver.Add_Handler( "save",save_file)
	//driver.Add_Handler( "file_exists",file_exists)
	//driver.Add_Handler( "delete_file",delete_file)
	//driver.Add_Handler( "file_directory",file_directory)
	//driver.Add_Handler( "make_dir",mkdir)
	//driver.rpc_start()
   
}


	


func pack_data(input [2]string )string {

  var b bytes.Buffer	
  msgpack.Pack(&b,input)
  return b.String()

}


func load_file( parms_packed string ) string {
  
  var return_value [2]string
  var input_message    = msgpack_utils.Unpack(parms_packed).(map[string]string) 
  var path = "/files/"+input_message["path"]+"/"+input_message["file_name"]
  var data, err = ioutil.ReadFile(path)
  if err != nil {
        return_value [0] = "false"
		return_value [1] = ""
  } else {
    return_value[0] = "true"
	return_value[1] = string(data)
  }
  return pack_data(return_value)
}

func save_file( parms_packed string ) string {
  
  var return_value [2]string
  var input_message    = msgpack_utils.Unpack(parms_packed).(map[string]string)  
  var path = "/files/"+input_message["path"]+"/"+input_message["file_name"]
  
  var err = ioutil.WriteFile(path,[]byte(input_message["data"]),0666)
  if err != nil {
        return_value [0] = "false"
		return_value [1] = ""
  } else {
    return_value[0] = "true"
	return_value[1] = ""
  }
  return pack_data(return_value)
}

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}
func file_exists( parms_packed string ) string {
 
  var return_value [2]string
  var input_message    = msgpack_utils.Unpack(parms_packed).(map[string]string)  
  var path = "/files/"+input_message["path"]+"/"+input_message["file_name"] 
  if  fileExists(path) == true{
  
        return_value [0] = "false"
		return_value [1] = ""
  } else {
    return_value[0] = "true"
	return_value[1] = ""
  }
  return pack_data(return_value)
}  
  



func delete_file( parms_packed string ) string {
  
  var return_value [2]string
  
  var input_message    = msgpack_utils.Unpack(parms_packed).(map[string]string)  
  var path = "/files/"+input_message["path"]+"/"+input_message["file_name"]
  e := os.Remove(path)
  if e != nil {
         return_value [0] = "false"
		return_value [1] = ""
  } else {
    return_value[0] = "true"
	return_value[1] = ""
  }
  return pack_data(return_value)
}  
   
  



func compact_file_data( c []fs.FileInfo)string {

  var return_value []string
  for _,element := range c {
     return_value = append(return_value,element.Name())
  }
  var b bytes.Buffer	
  msgpack.Pack(&b,return_value)
  return b.String()

}

func file_directory( parms_packed string ) string {
  var return_value [2]string
  var input_message    = msgpack_utils.Unpack(parms_packed).(map[string]string)  
  var path = "/files/"+input_message["path"]+"/"+input_message["file_name"]
  c, err := ioutil.ReadDir(path)
  if err != nil {
    return_value[0] = "false"
	return_value[1] = ""
  } else {
    return_value[0] = "true"
	return_value[1] = compact_file_data(c)
  }
   return pack_data(return_value)

}

func mkdir( parms_packed string ) string {
   
  var return_value [2]string 
  var input_message    = msgpack_utils.Unpack(parms_packed).(map[string]string)  
  var path = "/files/"+input_message["path"]
  err := os.MkdirAll(path,0666)
   if err != nil {
    return_value[0] = "false"
	return_value[1] = ""
  } else {
    return_value[0] = "true"
	return_value[1] = ""
  }
   return pack_data(return_value)
  
}




