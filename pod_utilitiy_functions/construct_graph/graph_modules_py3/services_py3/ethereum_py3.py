from .service_utilities_py3 import Start_Service
from .service_utilities_py3 import End_Service


class Ethereum_Service(object):

    def __init__(self,bc,cd,name):
       container_run_script = "docker run -d --name ethereum_go  -p 30303:30303     --mount type=bind,source=/mnt/ssd/ethereum,target=/data/    "  
       container_run_script = container_run_script + " nanodatacenter/ethereum-go  geth --dev --datadir /data/dev_data --ipcpath /data/geth.ipc " 
       

       
       Start_Service(bc,cd,name,container_run_script)             
       End_Service(bc,cd)  