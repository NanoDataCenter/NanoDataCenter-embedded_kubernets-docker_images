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


loop_flag = True
while loop_flag:
   try:
      redis_handle = redis.StrictRedis(site_data["host"], site_data["port"], db=site_data["redis_password_db"], decode_responses=True)
      if redis_handle.ping() == True:
         loop_flag = False
   except:

      time.sleep(10)   



print("redis server is up")
time.sleep(15) ## allow site services to setup

docker_control = Docker_Interface()


qs = Query_Support( site_data )
smtp =  SMTP_py3(site_data,"node_initialization")

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





















