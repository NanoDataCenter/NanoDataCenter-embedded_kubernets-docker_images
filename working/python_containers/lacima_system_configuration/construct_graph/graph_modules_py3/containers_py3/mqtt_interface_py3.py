from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container


class MQTT_Interface(object):

    def __init__(self,bc,cd,name):
       command_list = [  { "file":"mqtt_self_test_py3.py","restart":True },
       { "file":"mqtt_local_publish_server_py3.py","restart":True },
       { "file":"mqtt_redis_gateway_py3.py","restart":True },
       { "file":"mqtt_scan_data_py3.py","restart":True }
       ]
       starting_command = "docker run -d --name mqtt_interface --network host   --mount type=bind,source=/mnt/ssd/site_config,target=/data/   nanodatacenter/mqtt_interface /bin/bash process_control.bsh"
       Start_Container(bc,cd,name,starting_command,command_list,"nanodatacenter/mqtt_interface")             
       End_Container(bc,cd)  