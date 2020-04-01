from .container_utilities_py3 import Start_Container
from .container_utilities_py3 import End_Container


class Redis_Monitor_Container(object):

    def __init__(self,bc,cd,name):
       command_list = [  { "file":"redis_monitoring_py3.py","restart":True } ]
       Start_Container(bc,cd,name,command_list)       
       cd.construct_package("REDIS_MONITORING")      
       cd.add_redis_stream("REDIS_MONITOR_KEY_STREAM")
       cd.add_redis_stream("REDIS_MONITOR_CLIENT_STREAM")
       cd.add_redis_stream("REDIS_MONITOR_MEMORY_STREAM")
       cd.add_redis_stream("REDIS_MONITOR_CALL_STREAM")
       cd.add_redis_stream("REDIS_MONITOR_CMD_TIME_STREAM")
       cd.add_redis_stream("REDIS_MONITOR_SERVER_TIME")
       
       cd.close_package_contruction()
       End_Container(bc,cd)  