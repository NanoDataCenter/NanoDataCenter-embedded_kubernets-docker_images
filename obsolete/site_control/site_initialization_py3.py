



import json
import time
import os
import redis
from smtp_py3.smtp_py3 import  SMTP_py3
from docker_control.docker_interface_py3 import Docker_Interface
from common_tools.Pattern_tools_py3.builders.common_directors_py3 import construct_all_handlers
from common_tools.Pattern_tools_py3.factories.graph_search_py3 import common_qs_search
from common_tools.Pattern_tools_py3.factories.get_site_data_py3 import get_site_data
from common_tools.redis_support_py3.graph_query_support_py3 import  Query_Support
#
# Note these images are required to be on site_control node
#


class Site_Initialization(object):
    def __init__(self, config_file, password_script,redis_startup_script,redis_image):
                     
        self.site_data = self.determine_master(config_file)
        self.smtp =  SMTP_py3(self.site_data,"site_initialization")
        self.docker_control = Docker_Interface()
        self.stop_running_containers()
        
        self.remove_redis_container()
        
        self.startup_redis_container(redis_startup_script)
        self.wait_for_redis_connection()
        self.qs = Query_Support( self.site_data )
     
        self.determine_graph_container()
        
        self.load_graph_container()
        os.system(self.graph_container_script)
        
        os.system(password_script)
        
        self.site_containers = self.find_site_containers()

        self.start_site_containers()
        self.docker_control.prune()
        self.smtp.send_mail("site is intialized","site_initialization")
        
    def determine_master(self,site_file):
       site_data = get_site_data(site_file)
       
       print("site_data",site_data)
       if 'master' not in site_data:
            while True:  # not a site_control node
                  time.sleep(5)
       return site_data   
    
    def stop_running_containers(self):
        running_containers = self.docker_control.containers_ls_runing()
        for i in running_containers:
            self.docker_control.container_stop(i)      
    
    
    def remove_redis_container(self):
        self.docker_control.container_rm("redis")
        
        
    def startup_redis_container(self,redis_startup_script):
        self.docker_control.container_up("redis",redis_startup_script)
    
    
    def wait_for_redis_connection(self):
        print("waiting for redis connections")
        loop_flag = True
        while loop_flag:
           try:
              redis_handle = redis.StrictRedis(self.site_data["host"], self.site_data["port"], db=self.site_data["redis_password_db"], decode_responses=True)
              print(redis_handle.ping())
              if redis_handle.ping() == True:
                    loop_flag = False
           except:
                 print("exception")
                 time.sleep(1)
    
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





