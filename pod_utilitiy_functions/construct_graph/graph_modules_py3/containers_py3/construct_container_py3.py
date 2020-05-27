
from .redis_monitor_py3 import Redis_Monitor_Container
from .log_stream_events_py3 import LOG_STREAM_EVENTS_CONTAINER
from .stream_events_to_cloud_py3 import Stream_Events_To_Cloud
from .sqlite_server_py3 import SQLITE_Server
from .op_monitor_py3 import OP_Monitor
class Construct_Containers(object):

   def __init__(self,bc,cd,container_list):
      for i in container_list:
 
         if i == "monitor_redis":
             print(i)
             Redis_Monitor_Container(bc,cd,i)
         elif i == "stream_events_to_log":
             print(i)
             LOG_STREAM_EVENTS_CONTAINER(bc,cd,i)     
         elif i == "stream_events_to_cloud":
             print(i)
             Stream_Events_To_Cloud(bc,cd,i) 
         elif i == "sqlite_server":
             print(i)
             SQLITE_Server(bc,cd,i)  
         elif i == "op_monitor":
              print(i)
              OP_Monitor(bc,cd,i)              
         else:
             raise         
          

