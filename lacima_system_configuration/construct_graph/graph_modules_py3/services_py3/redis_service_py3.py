from .service_utilities_py3 import Start_Service
from .service_utilities_py3 import End_Service


class Redis_Service(object):

    def __init__(self,bc,cd,name):
       container_run_script = "docker run -d  --network host   --name redis    --mount type=bind,source=/mnt/ssd/redis,target=/data    nanodatacenter/redis ./redis-server ./redis.conf"

       
       Start_Service(bc,cd,name,container_run_script,"redis")             
       End_Service(bc,cd)  