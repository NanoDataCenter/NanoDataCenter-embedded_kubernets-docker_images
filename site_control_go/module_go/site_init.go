
package main

import ( 
"fmt"
"context"
"time"
"strconv"
"encoding/json"

"example.com/user/hello/smtp"
"example.com/user/hello/site_data"
"example.com/user/hello/docker_control"
"example.com/user/hello/redis_support/graph_query"
"github.com/go-redis/redis/v8"

)

type site_data_type map[string]interface{}
var site_data map[string]interface{}
type Database struct {
   Client *redis.Client
}
var Ctx    = context.TODO()
var graph_container_script string
var graph_container_image string
var services_json string
//var container []string
var services = make([]string,0)
var site_containers = make([]map[string]string,0)



func start_system_containers(){
  for _,value := range site_containers{
    if value["name"] == "redis" {
	  fmt.Println("found redis")
	  continue
	}
	if docker_control.Image_Exists(value["container_image"]) == false{
	   fmt.Println("should not happen")
	   docker_control.Pull(value["container_image"])
	}
	docker_control.Container_rm(value["name"])
	docker_control.Container_up(value["name"],value["startup_command"])
  }
}  
	
   

func find_site_services(){
    var temp string
    for _,service := range services{
	     var item = make(map[string]string,0)
	     var search_list = []string{"CONTAINER"+":"+service}
		 var container_nodes = graph_query.Common_qs_search(&search_list)
         var container_node = container_nodes[0]
		 //fmt.Println(container_node)
		 var err2 = json.Unmarshal([]byte(container_node["startup_command"]),&temp)
         if err2 != nil{
	         get_site_data.Can_Not_Continue("bad json data")
	     }
		 item["startup_command"] = temp
	     var err3 = json.Unmarshal([]byte(container_node["container_image"]),&temp)
         if err3 != nil{
	        get_site_data.Can_Not_Continue("bad json data")
	     }
         item["container_image"] = temp		 
		 var err4 = json.Unmarshal([]byte(container_node["name"]),&temp)
		 item["name"] = temp
		 if err4 != nil{
	        get_site_data.Can_Not_Continue("bad json data")
	     }	
		 
		 site_containers = append(site_containers,item)
		 
	}
	//fmt.Println(site_containers)
}

func  determine_graph_container(){
    var search_list = []string{ "SITE_CONTROL:SITE_CONTROL" }
    var site_nodes = graph_query.Common_qs_search(&search_list)
    var site_node = site_nodes[0]

    var graph_container_script_json = site_node["graph_container_script"]
    var graph_container_image_json = site_node["graph_container_image"]
    services_json = site_node["services"]
	
	

	var err1 = json.Unmarshal([]byte(services_json),&services)
    if err1 != nil{
	  get_site_data.Can_Not_Continue("bad json data")
	}	
	var err2 = json.Unmarshal([]byte(graph_container_script_json),&graph_container_script)
    if err2 != nil{
	  get_site_data.Can_Not_Continue("bad json data")
	}
	var err3 = json.Unmarshal([]byte(graph_container_image_json),&graph_container_image)
    if err3 != nil{
	  get_site_data.Can_Not_Continue("bad json data")
	}	
}


func wait_for_redis_connection(address string, port int ) {
   var address_port = address+":"+strconv.Itoa(port)
   //fmt.Println("address",address_port)
   fmt.Println("wait_for_redis_connection",port)
   var loop_flag = true
   for loop_flag == true {
       client := redis.NewClient(&redis.Options{
                                                 Addr: address_port,
												
												 DB: 0,
                                               })
		err := client.Ping(Ctx).Err();
		if err != nil{
		  fmt.Println("redis connection is not up")
		  time.Sleep(time.Second)
		}else {
		  loop_flag = false
		}  
      		
		
   }									   
}




func  startup_redis_container(redis_startup_script string){
      fmt.Println("start redis container")
	  docker_control.Container_up("redis",redis_startup_script)
}	 


func remove_redis_container(){
    fmt.Println("remove redis container")
	docker_control.Container_rm("redis")
}


func stop_running_containers() {
   fmt.Println("stop redis container")
   var running_containers = docker_control.Containers_ls_runing()
   for _,name := range running_containers{
      docker_control.Container_stop(name)
   }      
}


	   
func Site_Initialization( config_file string, 
                          password_script string, 
                         redis_startup_script string,
						 redis_image string) {


    

   site_data = get_site_data.Determine_master(config_file)
   
   fmt.Println("initialize Mail Server")
   smtp.Initialization(site_data,"SITE_CONTROL")
   stop_running_containers()
   remove_redis_container()
   startup_redis_container(redis_startup_script)
   wait_for_redis_connection(site_data["host"].(string), int(site_data["port"].(float64)) )
   fmt.Println("redis is up")
  
   graph_query.Graph_support_init(&site_data)
   determine_graph_container()
   docker_control.Pull(graph_container_image)
   docker_control.Container_Run(graph_container_script)
   docker_control.System(password_script)
   find_site_services()
   start_system_containers()
   docker_control.Prune()
   smtp.Send_Mail("site is intialized")
   

}	
						 


func main() {
    var redis_startup_script = "docker run -d  --network host   --name redis    --mount type=bind,source=/mnt/ssd/redis,target=/data    nanodatacenter/redis /bin/bash /pod_util/redis_control.bsh"
	var config_file = "/mnt/ssd/site_config/redis_server.json"
    var password_script ="python3 /mnt/ssd/site_config/passwords.py"
    var redis_image = "nanodatacenter/redis" 
    Site_Initialization(config_file,password_script,redis_startup_script,redis_image  )
}


    



/*

        self.qs = Query_Support( self.site_data )
     
        self.determine_graph_container()
        
        self.load_graph_container()
        os.system(self.graph_container_script)
        
        os.system(password_script)
        
        self.site_containers = self.find_site_containers()

        self.start_site_containers()
        self.docker_control.prune()
        self.smtp.send_mail("site is intialized","site_initialization")
        

 
    
    def determine_graph_container(self):
       search_list = [ "SITE_CONTROL","SITE_CONTROL" ]
       site_nodes = common_qs_search(self.site_data,self.qs,search_list)
       site_node = site_nodes[0]
       self.graph_container_script = site_node["graph_container_script"]
       self.graph_container_image = site_node['graph_container_image']
       self.services = site_node["services"]
       self.containers = site_node["containers"]
    
    
    def load_graph_container(self):
        self.load_docker_image(self.graph_container_image)
        
       
    def find_site_containers(self):
        site_containers = []
        for i in self.services:
            item = {}
            search_list = [[ "CONTAINER",i] ]
            services = common_qs_search(self.site_data,self.qs,search_list)
            service = services[0]
            item['container_image'] = service['container_image']
            item['startup_command'] = service['startup_command']
            item["name"] =i
            site_containers.append(item)
            
        for i in self.containers:
            item = {}
            search_list = [[ "CONTAINER",i] ]
            containers = common_qs_search(self.site_data,self.qs,search_list)
            container = containers[0]
            item['container_image'] = container['container_image']
            item['startup_command'] = container['startup_command']
            item["name"] =i  
            site_containers.append(item)
 
        return site_containers        
    


    def start_site_containers(self):
       for i in self.site_containers:
           if i["name"] == "redis":
               continue
           
           self.load_docker_image(i['container_image'])
           self.docker_control.container_rm(i["name"])
           self.docker_control.container_up(i["name"],i['startup_command'])
    
    def load_docker_image(self,image):
        if image not in self.docker_control.images():
           #raise ValueError("should not happen")
           self.pull_docker_image(image)
    
    
    def pull_docker_image(self,image):
       try:
           print("pulling images")
           docker_control.pull(image)
       except:
            self.smtp.send_mail("load image failure",image)
            while  True:
               print("fatal error missing image ",image)
               time.sleep(3600)











if __name__ == "__main__":
    redis_startup_script = "docker run -d  --network host   --name redis    --mount type=bind,source=/mnt/ssd/redis,target=/data    nanodatacenter/redis /bin/bash /pod_util/redis_control.bsh"
    Site_Initialization(config_file = "/mnt/ssd/site_config/redis_server.json",
                              password_script ="python3 /mnt/ssd/site_config/passwords.py",
                              redis_startup_script = redis_startup_script,
                              redis_image = "nanodatacenter/redis"   )





'''
from docker_control.docker_interface_py3 import Docker_Interface
from common_tools.Pattern_tools_py3.factories.get_site_data_py3 import get_site_data
from smtp_py3.smtp_py3 import  SMTP_py3
redis_startup_script = "docker run -d  --network host   --name redis    --mount type=bind,source=/mnt/ssd/redis,target=/data    nanodatacenter/redis /pod_util/redis-server /pod_util/redis.conf"
#sqlite_run_script = "docker run    -d  --network host   --name sqlite_server    --mount type=bind,source=/mnt/ssd/site_config,target=/data/   --mount type=bind,source=/mnt/ssd/sqlite/,target=/sqlite/  nanodatacenter/sqlite_server /bin/bash sqlite_control.bsh"
file_server_script = "docker run   -d  --network host   --name file_server        --mount type=bind,source=/mnt/ssd/site_config,target=/data/   --mount type=bind,source=/mnt/ssd/files/,target=/files/  nanodatacenter/file_server /bin/bash file_server_control.bsh"   

required_images = ["nanodatacenter/redis","nanodatacenter/file_server","nanodatacenter/lacima_system_configuration"]
required_containers = [ "redis"  ,"file_server" ]
startup_scripts = {}
startup_scripts["redis"] = redis_startup_script
#startup_scripts["sqlite_server"] = sqlite_run_script
startup_scripts["file_server"] = file_server_script

redis_site_file ="/mnt/ssd/site_config/redis_server.json"
graph_script ="docker run   -it --network host --rm  --name lacima_system_configuration  --mount type=bind,source=/mnt/ssd/site_config,target=/data/ nanodatacenter/lacima_system_configuration /bin/bash construct_graph.bsh"


        docker_control.container_rm(i)

running_containers = docker_control.containers_ls_runing()




system_images = docker_control.images()


for i in required_images:
   
   if i not in system_images:       
       load_docker_image(smtp,i)
       




if "redis" not in running_containers:
   docker_control.container_up("redis",startup_scripts["redis"])

   
  
print("loading configuration graph")
os.system(graph_script)


for i in required_containers:
   if i not in running_containers:
       docker_control.container_up(i,startup_scripts[i])


running_containers = docker_control.containers_ls_runing()

print("running containers",running_containers)
os.system()
'''


*/
