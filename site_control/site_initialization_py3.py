
import json
import time
import os
import redis
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



docker_control = Docker_Interface()

#
# Note these images are required to be on site_control node
#


def load_docker_image(smtp,image):
   try:
       print("pulling images")
       docker_control.pull(image)
   except:
       smtp.send_mail("load image failure",image)
       while  True:
          print("fatal error missing image ",image)
          time.sleep(3600)




site_data = get_site_data("/mnt/ssd/site_config/redis_server.json")
print("site_data",site_data)

if 'master' not in site_data:
   while True:  # not a site_control node
     time.sleep(5)

running_containers = docker_control.containers_ls_runing()
for i in required_containers:
    if i in running_containers:
        docker_control.container_stop(i)
        docker_control.container_rm(i)
docker_control.prune()
running_containers = docker_control.containers_ls_runing()

smtp =  SMTP_py3(site_data,"site_initialization")


system_images = docker_control.images()


for i in required_images:
   
   if i not in system_images:       
       load_docker_image(smtp,i)
       




if "redis" not in running_containers:
   docker_control.container_up("redis",startup_scripts["redis"])

print("waiting for redis connections")
loop_flag = True
while loop_flag:
   try:
      redis_handle = redis.StrictRedis(site_data["host"], site_data["port"], db=site_data["redis_password_db"], decode_responses=True)
      print(redis_handle.ping())
      if redis_handle.ping() == True:
         loop_flag = False
   except:
      print("exception")
      time.sleep(1)   
  
print("loading configuration graph")
os.system(graph_script)


for i in required_containers:
   if i not in running_containers:
       docker_control.container_up(i,startup_scripts[i])


running_containers = docker_control.containers_ls_runing()

print("running containers",running_containers)
os.system("python3 /mnt/ssd/site_config/passwords.py")














