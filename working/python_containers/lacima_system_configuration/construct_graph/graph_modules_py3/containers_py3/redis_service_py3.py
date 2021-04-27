from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container


class Redis_Service(object):

    def __init__(self,bc,cd,name):
       container_run_script = "docker run -d  --network host   --name redis    --mount type=bind,source=/mnt/ssd/redis,target=/data    nanodatacenter/redis ./redis-server ./redis.conf"

       
       Start_Container(bc,cd,name,container_run_script,[],"nanodatacenter/redis")             
       End_Container(bc,cd)  