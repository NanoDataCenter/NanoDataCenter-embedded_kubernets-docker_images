package web_support

import (
   
    
    "net"
    
	"encoding/json"

    "lacima.com/redis_support/graph_query"
    "lacima.com/redis_support/redis_handlers"
    "lacima.com/redis_support/generate_handlers"   
)



func Setup_Web_Server( system_name string  )(int64,bool) {

 
    search_path := []string{"WEB_SERVICES:"+system_name}
    nodes:= graph_query.Common_qs_search(  &search_path )
	
    node := nodes[0]
	
    port_json := node["port"]
	var port int64;
    err2 := json.Unmarshal([]byte(port_json),&port)
    if err2 != nil{
	         panic("bad json data")
	}
     
	
    search_path = []string{"WEB_DISCOVERY:WEB_DISCOVERY","IRRIGIGATION_IP_LOOK_UP"}
    handlers := data_handler.Construct_Data_Structures(&search_path)
    driver := (*handlers)["WEB_IP_LOOK_UP"].(redis_handlers.Redis_Hash_Struct)
	
	ip_string := get_out_bound_IP()
	
	driver.HSet(system_name,ip_string)
	
	return port,true
	



}


// Get preferred outbound ip of this machine
func get_out_bound_IP() string {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP.String()
}
