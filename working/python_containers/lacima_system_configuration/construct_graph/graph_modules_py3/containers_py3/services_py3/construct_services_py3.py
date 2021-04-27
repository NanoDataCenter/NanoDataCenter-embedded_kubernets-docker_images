
from .sql_service_py3 import SQLITE_Service
from .rpi_mosquito_service_py3 import RPI_Mosquitto_Service
from .redis_service_py3 import Redis_Service
from .ethereum_py3 import Ethereum_Service
from .file_server_service_py3 import File_System_Service


class Construct_Services(object):

   def __init__(self,bc,cd,services_list):
       for i in services_list:
 
           if i == "rpi_mosquitto":
               #print(i)
               RPI_Mosquitto_Service(bc,cd,i)
           elif i == "redis":
               #print(i)
               Redis_Service(bc,cd,i)     
           elif i == "sqlite_server":
               #print(i)
               SQLITE_Service(bc,cd,i) 
           elif i == "file_server":
               #print(i)
               File_System_Service(bc,cd,i)  
           elif i == "ethereum_go":
               #print(i)
               Ethereum_Service(bc,cd,i) 
           else:
               raise         
          

