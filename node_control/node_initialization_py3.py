import redis
import json
import time
import os
import sys

from docker_control.docker_interface_py3 import Docker_Interface
from redis_support_py3.graph_query_support_py3 import  Query_Support
redis_site_file = "/mnt/ssd/site_config/redis_server.json"

def wait_for_redis_db(site_data):
   
    while True:
        try:
            redis_handle = redis.StrictRedis( host = site_data["host"] , port=site_data["port"], db=site_data["graph_db"])
            temp = redis_handle.ping()
            print(temp)
            if temp == True:
              
              
               return
            else:
               raise
        except:
           print("exception")
           time.sleep(10)
           pass


def wait_for_file_server():
   loop_flag = True
   while loop_flag:
        
        running_containers = docker_control.containers_ls_runing()
        print("running_containers",running_containers)
        if 'file_server' in running_containers:
             loop_flag = False
        else:
           time.sleep(5)
        
        
        
        
docker_control = Docker_Interface()





file_handle = open(redis_site_file,'r')
try:    
    data = file_handle.read()
    file_handle.close()
    site_data = json.loads(data)
except:
    # post appropriate error message
    raise    


      
wait_for_redis_db(site_data)
wait_for_file_server()

















