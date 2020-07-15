from .service_utilities_py3 import Start_Service
from .service_utilities_py3 import End_Service


class Redis_Service(object):

    def __init__(self,bc,cd,name):
       container_run_script = "docker run -d  --restart on-failure  --name redis -p 6379:6379 --mount type=bind,source=/mnt/ssd/redis,target=/data  " 
       container_run_script = container_run_script + " --mount type=bind,source=/mnt/ssd/redis/config/redis.conf,target=/usr/local/etc/redis/redis.conf redis"
       

       
       Start_Service(bc,cd,name,container_run_script,"redis")             
       End_Service(bc,cd)  