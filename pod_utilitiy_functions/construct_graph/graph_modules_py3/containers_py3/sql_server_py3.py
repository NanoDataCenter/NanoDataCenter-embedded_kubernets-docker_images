from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container


class SQL_Server(object):

     def __init__(self,bc,cd,name):
          command_list = [  { "file":"sql_server_py3.py","restart":True } ]
          Start_Container(bc,cd,name,command_list)       
          End_Container(bc,cd)  