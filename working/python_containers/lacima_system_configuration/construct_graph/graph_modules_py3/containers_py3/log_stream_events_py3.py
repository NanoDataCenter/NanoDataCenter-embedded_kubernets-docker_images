
from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container


class LOG_STREAM_EVENTS_CONTAINER(object):

     def __init__(self,bc,cd,name):
          command_list = [  { "file":"log_stream_events_py3.py","restart":True },{ "file":"block_chain_event_rpc_server_py3.py","restart":True } ]
          startup_command = "docker run -d   --name stream_events_to_log  --network host    --mount type=bind,source=/mnt/ssd/site_config,target=/data/ "
          startup_command = startup_command + "  --mount type=bind,source=/mnt/ssd/ethereum/,target=/ipc/  --mount type=bind,source=/mnt/ssd/ethereum/keystore/,target=/keystore/ "
          startup_command = startup_command + " nanodatacenter/log_stream_events  /bin/bash process_control.bsh " 

          Start_Container(bc,cd,name,startup_command, command_list,"nanodatacenter/log_stream_events")       
          End_Container(bc,cd)  