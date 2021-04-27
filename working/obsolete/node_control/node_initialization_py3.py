
import json
import time
import os
from docker_control.docker_interface_py3 import Docker_Interface
from common_tools.Pattern_tools_py3.factories.get_site_data_py3 import get_site_data
from smtp_py3.smtp_py3 import  SMTP_py3
from common_tools.Pattern_tools_py3.factories.graph_search_py3 import common_qs_search
from common_tools.Pattern_tools_py3.factories.iterators_py3 import pattern_iter_strip_list_dict
from common_tools.redis_support_py3.graph_query_support_py3 import  Query_Support
import redis



class Node_Initialization(object):

    def __init__(self,config_file):
       self.site_data = get_site_data(config_file)
       self.docker_control = Docker_Interface()
       self.smtp =  SMTP_py3(self.site_data,"node_initialization")
       self.wait_for_redis_connection()
       self.qs = Query_Support( self.site_data )
       self.find_site_containers()
       self.wait_for_site_containers()
       self.node_containers = self.find_node_containers()
       self.stop_node_containers()
       self.load_container_images()
       self.start_container_images()
       self.smtp.send_mail("node is intialized","node_initialized")

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
    
    def find_site_containers(self):
       search_list = [ "SITE_CONTROL","SITE_CONTROL" ]
       site_nodes = common_qs_search(self.site_data,self.qs,search_list)
       site_node = site_nodes[0]
       self.site_containers = []
       self.site_containers=site_node["services"]
       self.site_containers.extend(site_node["containers"])
       
         
    
   
    
    def wait_for_site_containers(self):
        loop_flag = True
        while loop_flag:
           loop_flag = False
           running_containers = self.docker_control.containers_ls_runing()
           for i in self.site_containers:
             
              if i not in running_containers:
                  loop_flag = True
           if loop_flag == True:
              time.sleep(1)
 
    

    def find_node_containers(self):
        search_list = [ ["PROCESSOR" ,self.site_data["local_node"]   ] ]
        processor_nodes = common_qs_search(self.site_data,self.qs,search_list)
        processor_node = processor_nodes[0]
        self.services = processor_node["services"]
        self.containers = processor_node["containers"]
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

    def stop_node_containers(self):
       running_containers =  running_containers = self.docker_control.containers_ls_runing()
       for i in self.node_containers:
           name = i["name"]
           if name in running_containers:
               self.docker_control.container_stop(name)
           self.docker_control.container_rm(name)
    
    def load_container_images(self):
       for i in self.node_containers:
           image = i["container_image"]
           self.load_docker_image(image)
             
    
    def start_container_images(self):
       for i in self.node_containers:
           
           self.docker_control.container_up(i["name"], i['startup_command'])

    def load_docker_image(self,image):
        if image not in self.docker_control.images():
           raise ValueError("should not happen")
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
   config_file ="/mnt/ssd/site_config/redis_server.json"
   Node_Initialization(config_file)

'''
import json
import time
import os
from docker_control.docker_interface_py3 import Docker_Interface
from common_tools.Pattern_tools_py3.factories.get_site_data_py3 import get_site_data
from smtp_py3.smtp_py3 import  SMTP_py3
from common_tools.Pattern_tools_py3.factories.graph_search_py3 import common_qs_search
from common_tools.Pattern_tools_py3.factories.iterators_py3 import pattern_iter_strip_list_dict
from common_tools.redis_support_py3.graph_query_support_py3 import  Query_Support
import redis

site_data = get_site_data("/mnt/ssd/site_config/redis_server.json")
required_containers = ["redis","file_server"]

docker_control = Docker_Interface()
smtp =  SMTP_py3(site_data,"node_initialization")
loop_flag = True
while loop_flag:
   try:
      redis_handle = redis.StrictRedis(site_data["host"], site_data["port"], db=site_data["redis_password_db"], decode_responses=True)
      if redis_handle.ping() == True:
         loop_flag = False
   except:

      time.sleep(10)   
print("redis server is up")


# need to wait for requried containers

while loop_flag:
   time.sleep(3)
   loop_flag == False;
   running_containers = docker_control.containers_ls_runing()
   for i in required_containers:
      if i not in running_containers:
          loop_flag = True

running_containers = docker_control.containers_ls_runing()
for i in required_containers:
    if i not in required_containers:
        docker_control.container_stop(i)
        docker_control.container_rm(i)
docker_control.prune()
running_containers = docker_control.containers_ls_runing()


qs = Query_Support( site_data )


#
#  Basic services are up now can use tree to find configuration
#
# find containers and services

search_list = [["PROCESSOR",site_data["local_node"]]]
processor_data = common_qs_search(site_data,qs,search_list)[0]

services = processor_data["services"]
containers = processor_data["containers"]

# find data about services
search_list = ["SERVICE"]
service_data = common_qs_search(site_data,qs,search_list)
print(service_data)

# find data about containers
search_list = ["CONTAINER"]
container_data = common_qs_search(site_data,qs,search_list)
print(container_data)

common_containers = service_data
common_containers.extend(container_data)




# find images
required_images = pattern_iter_strip_list_dict(common_containers,"container_image")

system_images = docker_control.images()
print("system_images",system_images)

print("required_images",required_images)

for i in required_images:
   if i not in system_images:

      docker.pull(i)  # put error handler aroung ithis

running_containers = docker_control.containers_ls_runing()   
print(running_containers)   
for i in common_containers:
    if i["name"] not in running_containers:
       docker_control.container_up(i["name"],i["startup_command"])
 
# ready for node operation
smtp.send_mail("node is intialized","node_initialized")


# find descriptors for containers and services

# find images for all containers and services  pull if not there

# start all container and services

'''



















