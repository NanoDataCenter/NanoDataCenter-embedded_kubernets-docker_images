from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container


class Redis_Monitor_Container(object):

    def __init__(self,bc,cd,name):
      
      
       container_run_script = "docker run -d  --network host   --name monitor_redis  --mount type=bind,source=/mnt/ssd/site_config,target=/data/ nanodatacenter/redis_monitoring  /bin/bash ./process_control.bsh "
       command_list = [  {"command":"./redis_monitoring","key":"redis_monitoring" } ]
              
       Start_Container(bc,cd,name,container_run_script,command_list,"nanodatacenter/redis_monitoring")             
       End_Container(bc,cd)  