
import json


ONE_WEEK = 24*7

class Start_Container(object):
    
    def __init__(self,bc,cd,name,command_list)
        properties = []
        properties["command_list"] = command_list
        bc.add_header_node("CONTAINER",name,properties=properties)
        cd.construct_package("DATA_STRUCTURES")
        cd.add_redis_stream("ERROR_STREAM",forward=True)
        cd.add_hash("ERROR_HASH")
        cd.add_job_queue("WEB_COMMAND_QUEUE",1)
        cd.add_hash("WEB_DISPLAY_DICTIONARY")
        cd.close_package_contruction()
        
class End_Container(object):

     def __init__(self,bc,cd)
          bc.end_header_node("CONTAINER")
          
          
          
class Controller_Monitoring_Container(object):
    
    def __init__(self,bc,cd):
       command_list = [ { "file":"pi_monitoring_py3.py","restart":True } ]
       Start_Container(bc,cd,"controller_monitor",command_list)       
       cd.construct_package("SYSTEM_MONITORING")
       cd.add_redis_stream("FREE_CPU",forward = True) # one month of data
       cd.add_redis_stream("RAM",forward = True)
       cd.add_redis_stream("DISK_SPACE",forward = True) # one month of data
       cd.add_redis_stream("TEMPERATURE",forward = True)
       cd.add_redis_stream("PROCESS_VSZ")
       cd.add_redis_stream("PROCESS_RSS")
       cd.add_redis_stream("PROCESS_CPU")
       
       cd.add_redis_stream("CPU_CORE")
       cd.add_redis_stream("SWAP_SPACE")
       cd.add_redis_stream("IO_SPACE")
       cd.add_redis_stream("BLOCK_DEV")
       cd.add_redis_stream("CONTEXT_SWITCHES")
       cd.add_redis_stream("RUN_QUEUE")       
       cd.add_redis_stream("DEV") 
       cd.add_redis_stream("SOCK") 
       cd.add_redis_stream("TCP") 
       cd.add_redis_stream("UDP") 
       cd.close_package_contruction()
       End_Container(bc,cd)

class Redis_Monitor_Container(object):

       command_list = [  { "file":"redis_monitoring_py3.py","restart":True } ]
       Start_Container(bc,cd,"controller_monitor",command_list)       
       cd.construct_package("REDIS_MONITORING")      
       cd.add_redis_stream("REDIS_MONITOR_KEY_STREAM")
       cd.add_redis_stream("REDIS_MONITOR_CLIENT_STREAM")
       cd.add_redis_stream("REDIS_MONITOR_MEMORY_STREAM")
       cd.add_redis_stream("REDIS_MONITOR_CALL_STREAM")
       cd.add_redis_stream("REDIS_MONITOR_CMD_TIME_STREAM")
       cd.add_redis_stream("REDIS_MONITOR_SERVER_TIME")
       
       cd.close_package_contruction()
       bc.end_header_node("CONTAINER")    
    

class Construct_Cloud_Controllers(object):

   def __init__(self,bc,cd):

 
       properties = []
       properties["containers"] = ["monitor_processor","monitor_redis"]
     
       bc.add_header_node("PROCESSOR","block_chain_server",properties= properties) # name is identified in site_data["local_node"]
 
       bc.end_header_node("PROCESSOR")
       
       #
       #
       #  Add other processes if desired
       #
class Construct_Containers(object)
    
    
    