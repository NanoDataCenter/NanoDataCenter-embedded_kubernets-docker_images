#
#
# File: eto.py
#
#

import datetime
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
from system_error_log_py3 import  System_Error_Logging
from Pattern_tools_py3.builders.common_directors_py3 import construct_all_handlers
from     sqlite_library.sqlite_sql_support_py3 import SQLITE_Client_Support
ONE_DAY = 24 * 3600


class Eto_Monitoring(object):
    def __init__(self,qs,site_data ):
        
        container_name = os.getenv("CONTAINER_NAME")
        self.sqlite_client = SQLITE_Client_Support(qs,site_data)
      
        self.error_logging = System_Error_Logging(qs,container_name,site_data,self.sqlite_client)

        
        search_list = ["WEATHER_STATION_DATA"]
        self.ds_handlers = construct_all_handlers(site_data,qs,search_list)
    
     


    def check_new_day_rollover( self, *parameters ):
        error_flag = True
        if  self.ds_handlers["ETO_CONTROL"].hget("ETO_UPDATE_FLAG") !=0:
            error_flag = False        
             
        if  self.ds_handlers["ETO_CONTROL"].hget("ETO_LOG_FLAG")!=0:
             error_flag = False   
        


        if error_flag == False:
            self.error_logging.log_error_message("ETO_Rollover",["action",False])
        else:
             self.error_logging.log_error_message("ETO_Rollover",["action",True])
        
        return "DISABLE"

    def check_eto_bin_update(self,*parameters):
        error_flag = True
        if  self.ds_handlers["ETO_CONTROL"].hget("ETO_UPDATE_FLAG") !=1:
            error_flag = False        
             
        if  self.ds_handlers["ETO_CONTROL"].hget("ETO_LOG_FLAG")!=1:
             error_flag = False   
        
        if  self.ds_handlers["ETO_CONTROL"].hget("ETO_UPDATE_VALUE")==None:
             error_flag = False   
       

 
        if error_flag == False:
            self.error_logging.log_error_message("ETO Update",["action",False])
        else:
             self.error_logging.log_error_message("ETO Update",["action",True])

 
        return "DISABLE"
    
  
    


def add_eto_chains(eto, cf):


    cf.define_chain("Monitor_day_rollover", True)
    cf.insert.wait_tod_le( hour =  6)
    cf.insert.wait_tod_ge( hour =  5)
    cf.insert.one_step(eto.check_new_day_rollover)
    cf.insert.wait_tod_le( hour =  6)
    cf.insert.reset()

    cf.define_chain("monitor_eto_update", True)
    cf.insert.wait_tod_le( hour =  14)
    cf.insert.wait_tod_ge( hour =  13)
    cf.insert.one_step(eto.check_eto_bin_update)
    cf.insert.wait_tod_le( hour =  14)
    cf.insert.reset()



   
if __name__ == "__main__":

    import datetime
    import time
    import string
    import urllib.request
    import math
    import redis
    import base64
    import json

    import os
    import copy
    #import load_files_py3
    from redis_support_py3.graph_query_support_py3 import  Query_Support
    import datetime
    
    from py_cf_new_py3.chain_flow_py3 import CF_Base_Interpreter

    #
    #
    # Read Boot File
    # expand json file
    # 
    file_handle = open("/data/redis_server.json",'r')
    data = file_handle.read()
    file_handle.close() 
    redis_site = json.loads(data)
     
    #
    # Setup handle
    # open data stores instance
    
    
     
       
    qs = Query_Support( redis_site )
    
  
    
    eto = Eto_Monitoring(qs,redis_site)
   
   

  
    #
    # Adding chains
    #

    cf = CF_Base_Interpreter()
    add_eto_chains(eto, cf)
    #
    # Executing chains
    #
    
    
    cf.execute()

else:
  pass
  
