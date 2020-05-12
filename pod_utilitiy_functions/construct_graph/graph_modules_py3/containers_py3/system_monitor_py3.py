
from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container

class Redis_Monitor_Container(object):

    def __init__(self,bc,cd,name):
       command_list = [  { "file":"system_monitor_py3.py","restart":True } ]
       Start_Container(bc,cd,name,command_list)       
       cd.construct_package("SYSTEM_MONITOR")      
       #cd.add_managed_hash(self,name,fields,forward=False) perfored way to store field how to get field in system
       cd.add_hash(("SYSTEM_STATUS")
       cd.add_redis_stream("SYSTEM_ALERTS")
       cd.add_redis_stream("SYSTEM_PUSHED_ALERTS")
      
       
       cd.close_package_contruction()
       End_Container(bc,cd)  