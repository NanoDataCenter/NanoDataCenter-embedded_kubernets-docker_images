from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container


class IRRIGATION_SCHEDULING(object):

     def __init__(self,bc,cd,name):
          command_list = [  { "file":"irrigation_scheduling_py3.py","restart":True } ]
          startup_command = "docker run  -d    --network host   --name irrigation_scheduling    --mount type=bind,source=/mnt/ssd/site_config,target=/data/  "
          startup_command = startup_command + " nanodatacenter/irrigation_scheduling /bin/bash process_control.bsh "



          Start_Container(bc,cd,name,startup_command,command_list,"nanodatacenter/irrigation_scheduling")       
          End_Container(bc,cd)  