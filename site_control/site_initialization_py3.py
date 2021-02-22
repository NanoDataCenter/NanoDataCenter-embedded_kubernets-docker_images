import redis
import json
import time
import os
from docker_control.docker_interface_py3 import Docker_Interface
from redis_support_py3.graph_query_support_py3 import  Query_Support
redis_site_file = "/mnt/ssd/site_config/redis_server.json"

file_server_startup_script = "docker run    -d  --network host   --name file_server    --mount type=bind,source=/mnt/ssd/site_config,target=/data/ "      
file_server_startup_script = file_server_startup_script + "   --mount type=bind,source=/mnt/ssd/files/,target=/files/  nanodatacenter/file_server /bin/bash file_server_control.bsh"  
 
redis_startup_script = "docker run -d   --restart on-failure  --name redis -p 6379:6379 --mount type=bind,source=/mnt/ssd/redis,target=/data  " 
redis_startup_script = redis_startup_script + " --mount type=bind,source=/mnt/ssd/redis/config/redis.conf,target=/usr/local/etc/redis/redis.conf redis"
        
docker_control = Docker_Interface()





file_handle = open(redis_site_file,'r')
try:    
    data = file_handle.read()
    file_handle.close()
    site_data = json.loads(data)
except:
    # post appropriate error message
    raise    



if 'master' not in site_data:
   while True:  # not a site_control node
     time.sleep(5)
     
# startup redis
#pull redis container
docker_control.container_up("redis",redis_startup_script)
# startup file_server
docker_control.pull("nanodatacenter/file_server")
docker_control.container_up("file_server",file_server_startup_script)















