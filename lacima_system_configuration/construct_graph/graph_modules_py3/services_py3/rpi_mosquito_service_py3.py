from .service_utilities_py3 import Start_Service
from .service_utilities_py3 import End_Service


class RPI_Mosquitto_Service(object):

    def __init__(self,bc,cd,name):
       container_run_script = "docker run -d  --name rpi_mosquitto  --network host     nanodatacenter/rpi_mosquitto "

       
       Start_Service(bc,cd,name,container_run_script,"nanodatacenter/rpi_mosquitto")             
       End_Service(bc,cd)  



       
       