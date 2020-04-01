
from .redis_monitor_py3 import Redis_Monitor_Container


class Construct_Containers(object):

   def __init__(self,bc,cd,container_list):
      for i in container_list:
 
         if i == "monitor_redis":
             print(i)
             Redis_Monitor_Container(bc,cd,i)  
         else:
             raise         
          

