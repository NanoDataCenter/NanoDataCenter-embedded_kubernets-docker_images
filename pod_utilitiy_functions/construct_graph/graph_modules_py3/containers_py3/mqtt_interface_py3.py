from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container


class MQTT_Interface(object):

    def __init__(self,bc,cd,name):
       command_list = [  { "file":"mqtt_self_test_py3.py","restart":True },
       { "file":"mqtt_local_publish_server_py3.py","restart":True },
       { "file":"mqtt_redis_gateway_py3.py","restart":True },
       { "file":"mqtt_scan_data_py3.py","restart":True }
       ]
       Start_Container(bc,cd,name,command_list)             
       End_Container(bc,cd)  