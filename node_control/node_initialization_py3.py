import json
import time
import os
from docker_control.docker_interface_py3 import Docker_Interface
from Pattern_tools_py3.factories.get_site_data_py3 import get_site_data
from smtp_py3.smtp_py3 import  SMTP_py3
from Pattern_tools_py3.factories.graph_search_py3 import common_qs_search
from Pattern_tools_py3.factories.iterators_py3 import pattern_iter_strip_list_dict
from redis_support_py3.graph_query_support_py3 import  Query_Support



docker_control = Docker_Interface()

site_data = get_site_data("/mnt/ssd/site_config/redis_server.json")
qs = Query_Support( site_data )
#smtp =  SMTP_py3(site_data,"node_initialization")

required_containers = ["redis","file_server","sqlite_server"]

loop_flag = True
while loop_flag == True:
   running_containers = docker_control.containers_ls_runing()
   
   loop_flag = False
   for i in required_containers:
       if i not in running_containers:
           
           loop_flag = True
   if loop_flag == True:
       time.sleep(10)

# find containers and services

search_list = [["PROCESSOR",site_data["local_node"]]]
processor_data = common_qs_search(site_data,qs,search_list)[0]

services = processor_data["services"]
containers = processor_data["containers"]

search_list = ["SERVICE"]
service_data = common_qs_search(site_data,qs,search_list)
print(service_data)

search_list = ["CONTAINER"]
container_data = common_qs_search(site_data,qs,search_list)
print(container_data)

common_containers = service_data
common_containers.extend(container_data)

all_images = pattern_iter_strip_list_dict(common_containers,"container_image")

system_images = docker_control.images()
print("system_images",system_images)

print("all_images",all_images)

for i in system_images:
   if i not in system_images:
      docker.pull(i)  # put error handler aroung ithis
      
print("done")

# find descriptors for containers and services

# find images for all containers and services  pull if not there

# start all container and services





















