from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container


class MODBUS_SERVER_CONTAINER(object):

     def __init__(self,bc,cd,name):
          command_list = [  { "file":"modbus_server_relay_py3.py MAIN_SERVER  192.168.1.227 MAIN_SERVER_QUEUE","restart":True } ]
          Start_Container(bc,cd,name,command_list)       
          End_Container(bc,cd)  