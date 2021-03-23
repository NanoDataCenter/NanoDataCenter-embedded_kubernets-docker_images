from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container


class Stream_Events_To_Cloud(object):

     def __init__(self,bc,cd,name):
          startup_command = "docker run -d --name stream_events_to_cloud   --network host    --mount type=bind,source=/mnt/ssd/site_config,target=/data/  "
          startup_command = startup_command + " nanodatacenter/stream_events_to_cloud  /bin/bash process_control.bsh "

          command_list = [  { "file":"stream_events_to_cloud_py3.py","restart":True } ]
          Start_Container(bc,cd,name,startup_command,command_list,"nanodatacenter/stream_events_to_cloud")       
          End_Container(bc,cd)  