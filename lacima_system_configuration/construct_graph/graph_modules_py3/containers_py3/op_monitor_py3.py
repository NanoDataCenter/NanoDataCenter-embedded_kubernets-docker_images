from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container


class OP_Monitor(object):

    def __init__(self,bc,cd,name):
       command_list = [  { "file":"op_monitoring_py3.py","restart":True } ]
       starting_command = "docker run -d  --network host   --name op_monitor  --mount type=bind,source=/mnt/ssd/site_config,target=/data/ nanodatacenter/op_monitor  /bin/bash ./process_control.bsh"
       Start_Container(bc,cd,name,starting_command,command_list,"nanodatacenter/op_monitor")             
       End_Container(bc,cd)  