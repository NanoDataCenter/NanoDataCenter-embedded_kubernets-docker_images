
from .redis_monitor_py3 import Redis_Monitor_Container
from .log_stream_events_py3 import LOG_STREAM_EVENTS_CONTAINER
from .stream_events_to_cloud_py3 import Stream_Events_To_Cloud

from .op_monitor_py3 import OP_Monitor
from .mqtt_interface_py3 import MQTT_Interface
from .irrigation_container_py3 import IRRIGATION_CONTROL
from .eto_container_py3  import ETO_CONTAINER
from .irrigation_scheduling_py3 import IRRIGATION_SCHEDULING
from .plc_io_container_py3  import PLC_IO_CONTAINER
from .modbus_server_container_py3  import MODBUS_SERVER_CONTAINER
from .redis_service_py3 import Redis_Service

from .file_server_service_py3 import File_System_Service


class Construct_Containers(object):

   def __init__(self,bc,cd,container_list):
      for i in container_list:
 
         if i == "monitor_redis":
             #print(i)
             Redis_Monitor_Container(bc,cd,i)
         elif i == "stream_events_to_log":
             #print(i)
             LOG_STREAM_EVENTS_CONTAINER(bc,cd,i)     
         elif i == "stream_events_to_cloud":
             #print(i)
             Stream_Events_To_Cloud(bc,cd,i) 

         elif i == "op_monitor":
              #print(i)
              OP_Monitor(bc,cd,i) 
         elif i == "mqtt_interface":
              MQTT_Interface(bc,cd,i)
         elif i == "irrigation_control":
              IRRIGATION_CONTROL(bc,cd,i)
              
         elif i == "eto":
              ETO_CONTAINER(bc,cd,i)   
         elif i == "irrigation_scheduling":
              IRRIGATION_SCHEDULING(bc,cd,i)     
         elif i == "plc_io":
              PLC_IO_CONTAINER(bc,cd,i)    
              
         elif i == "modbus_server":
              MODBUS_SERVER_CONTAINER(bc,cd,i)  
         elif i == "redis":
               #print(i)
               Redis_Service(bc,cd,i)     

         elif i == "file_server":
               #print(i)
               File_System_Service(bc,cd,i)                
         else:
             raise         
          

