from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container


class IRRIGATION_CONTROL(object):

     def __init__(self,bc,cd,name):
          command_list = [  { "file":"log_stream_events_py3.py","restart":True },{ "file":"fill this in.py","restart":True } ]
          Start_Container(bc,cd,name,command_list,"FILL THIS IN")       
          End_Container(bc,cd)  