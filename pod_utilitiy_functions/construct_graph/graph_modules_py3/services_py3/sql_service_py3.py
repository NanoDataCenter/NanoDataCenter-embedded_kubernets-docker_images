from .service_utilities_py3 import Start_Service
from .service_utilities_py3 import End_Service


class SQLITE_Service(object):

    def __init__(self,bc,cd,name):
       container_run_script = "docker run    -d  --network host   --name sqlite_server    --mount type=bind,source=/mnt/ssd/site_config,target=/data/   --mount type=bind,source=/mnt/ssd/sqlite/,target=/sqlite/  nanodatacenter/sqlite_server /bin/bash sqlite_control.bsh"

       
       Start_Service(bc,cd,name,container_run_script,"nanodatacenter/sqlite_server")             
       End_Service(bc,cd)  