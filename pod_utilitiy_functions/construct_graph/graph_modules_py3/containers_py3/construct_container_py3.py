
from .redis_monitor_py3 import Redis_Monitor_Container
from .log_stream_events_py3 import LOG_STREAM_EVENTS_CONTAINER

class Construct_Containers(object):

   def __init__(self,bc,cd,container_list):
      for i in container_list:
 
         if i == "monitor_redis":
             print(i)
             Redis_Monitor_Container(bc,cd,i)
         elif i == "log_stream_events":
             print(i)
             LOG_STREAM_EVENTS_CONTAINER(bc,cd,i)               
         else:
             raise         
          

