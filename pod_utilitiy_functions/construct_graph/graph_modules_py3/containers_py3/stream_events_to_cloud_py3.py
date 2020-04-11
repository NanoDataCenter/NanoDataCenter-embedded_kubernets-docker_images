from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container


class Stream_Events_To_Cloud(object):

     def __init__(self,bc,cd,name):
          command_list = [  { "file":"stream_events_to_cloud_py3.py","restart":True } ]
          Start_Container(bc,cd,name,command_list)       
          End_Container(bc,cd)  