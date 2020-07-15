from .service_utilities_py3 import Start_Service
from .service_utilities_py3 import End_Service


class File_System_Service(object):

    def __init__(self,bc,cd,name):
       container_run_script = "docker run    -d  --network host  --restart=always --name file_server    --mount type=bind,source=/mnt/ssd/site_config,target=/data/ "      
       container_run_script = container_run_script + "   --mount type=bind,source=/mnt/ssd/files/,target=/files/  nanodatacenter/file_server /bin/bash file_server_control.bsh"  
              
       Start_Service(bc,cd,name,container_run_script,"nanodatacenter/file_server")             
       End_Service(bc,cd)  