
import redis
import msgpack
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
    
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
    #from redis_support_py3.user_data_tables_py3 import User_Data_Tables

    from py_cf_new_py3.chain_flow_py3 import CF_Base_Interpreter

    redis_handle = redis.StrictRedis(  db=11 )
    payload = {}
    payload["command"] = "SEND_ALERT"
    payload["message"] = "Water On"
    payload["duration"] =  1000
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
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=redis_site["site"] )
    query_list = qs.add_match_relationship( query_list, relationship="WIFI_DEVICES" )
    query_list = qs.add_match_terminal( query_list, relationship="PACKAGE" )
    package_sets, package_sources = qs.match_list(query_list)  
    package = package_sources[0]
    data_structures = package["data_structures"]
    #print("data_structurres",data_structures);
   
    generate_handlers = Generate_Handlers( package, qs )
    handlers = {}
    handlers["DEVICE_STATUS"] = generate_handlers.construct_hash(data_structures["DEVICE_STATUS"])
    handlers["UNKNOWN_DEVICES"] = generate_handlers.construct_hash(data_structures["UNKNOWN_DEVICES"])
    handlers["ALERT_PUSH"] = generate_handlers.construct_job_queue_client(data_structures["ALERT_PUSH"])
    handlers["INTERNAL_SENSORS"] = generate_handlers.construct_hash(data_structures["INTERNAL_SENSORS"])
    handlers["WELL_SENSORS"] = generate_handlers.construct_hash(data_structures["WELL_SENSORS"])
    handlers["TURN_ON_WATER"] = generate_handlers.construct_job_queue_client(data_structures["TURN_ON_WATER"])

    
    loop_flag = True
    while(loop_flag):
    
        data_flag, data = handlers["TURN_ON_WATER"].pop()
        
        if data_flag == True:
            print(data)
            payload["NODE_ID"] = data['NODE_ID']
            print(payload)
            handlers["ALERT_PUSH"].push(payload)

        
 
