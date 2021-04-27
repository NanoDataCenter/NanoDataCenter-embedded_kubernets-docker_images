from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container


class PLC_IO_CONTAINER(object):

     def __init__(self,bc,cd,name):
          command_list = [  { "file":"plc_io_cntrl_py3.py","restart":True } ]
          starting_command = "docker run -d  --network host   --name plc_io   --mount type=bind,source=/mnt/ssd/site_config,target=/data/ nanodatacenter/plc_io  /bin/bash ./process_control.bsh"
          Start_Container(bc,cd,name,starting_command,command_list,"nanodatacenter/plc_io")       
          End_Container(bc,cd)  