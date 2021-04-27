from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container


class File_System_Service(object):

    def __init__(self,bc,cd,name):
       container_run_script = "docker run   -d  --network host   --name file_server        --mount type=bind,source=/home/pi/mountpoint/lacuma_conf/site_config,target=/data/   --mount type=bind,source=/home/pi/mountpoint/lacuma_conf/files/,target=/files/  nanodatacenter/file_server /bin/bash file_server_control.bsh"   
       command_list = [  {"command":"./file_server","key":"file_server" } ]
              
       Start_Container(bc,cd,name,container_run_script,command_list,"nanodatacenter/file_server")             
       End_Container(bc,cd)  