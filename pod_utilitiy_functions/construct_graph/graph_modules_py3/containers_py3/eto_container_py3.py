from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container


class ETO_CONTAINER(object):

     def __init__(self,bc,cd,name):
          start_command = "docker run    -d  --network host   --name eto    --mount type=bind,source=/mnt/ssd/site_config,target=/data/  "  
          start_command =  start_command + "  --mount type=bind,source=/mnt/ssd/files/,target=/files/  nanodatacenter/eto /bin/bash process_control.bsh "
          command_list = [  { "file":"eto_py3.py","restart":True },{ "file":"eto_monitoring_py3.py","restart":True }  ]
          Start_Container(bc,cd,name,start_command, command_list,"nanodatacenter/eto")       
          End_Container(bc,cd)  