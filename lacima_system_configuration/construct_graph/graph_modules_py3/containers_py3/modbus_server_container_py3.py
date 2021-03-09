from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container


class MODBUS_SERVER_CONTAINER(object):

     def __init__(self,bc,cd,name):
          startup_command = "docker run -d  --network host --privileged  --name modbus_server   --mount type=bind,source=/mnt/ssd/site_config,target=/data/ "
          startup_command = startup_command + " nanodatacenter/modbus_server  /bin/bash ./process_control.bsh "

          
          command_list = [  { "file":"modbus_server_relay_py3.py MAIN_SERVER  192.168.1.227 MAIN_SERVER_QUEUE","restart":True } ]
          Start_Container(bc,cd,name,startup_command,command_list,"nanodatacenter/modbus_server")       
          End_Container(bc,cd)  