from op_monitor_lib.construct_monitor_py3 import Construct_Monitors



class Op_Monitor(object):

   def __init__(self,site_data, qs,monitoring_list ):
       self.site_data = site_data
       self.qs        = qs
       self.construct_monitors = Construct_Monitors(site_data,qs,monitoring_list)   
       
       
   def monitor_subsystems(self,*parameters):
       self.construct_monitors.execute_monitors()
       
 
       
def construct_op_monitoring_instance( qs, site_data ):

                   
    
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )

    query_list = qs.add_match_terminal( query_list, 
                                        relationship = "OP_MONITOR" )
                                           
    data_sets, data_sources = qs.match_list(query_list)  
 
    data = data_sources[0]
    monitoring_list = data['OP_MONITOR_LIST']
    
    print("monitoring_list", monitoring_list)
    op_monitor = Op_Monitor(site_data ,qs,monitoring_list )
    
    
    
  

    return op_monitor




def add_chains(redis_monitor, cf):
 
    cf.define_chain("make_measurements", True)
    cf.insert.log("starting op monitoring")
    cf.insert.one_step(op_monitor.monitor_subsystems)
    cf.insert_log("ending op monitoring")
    cf.insert.wait_event_count( event = "MINUTE_TICK",count = 60)
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
    #from redis_support_py3.user_data_tables_py3 import User_Data_Tables

    from py_cf_new_py3.chain_flow_py3 import CF_Base_Interpreter

    #
    #
    # Read Boot File
    # expand json file
    # 
    file_handle = open("/data/redis_server.json",'r')
    data = file_handle.read()
    file_handle.close()
    site_data = json.loads(data)
    print("made it here 1")
    #
    # Setup handle
    # open data stores instance
   
    qs = Query_Support( site_data )
    
    redis_monitor = construct_op_monitoring_instance(qs, site_data )
    print("made it here 2")
    #
    # Adding chains
    #
    cf = CF_Base_Interpreter()
    add_chains(redis_monitor, cf)
    #
    # Executing chains
    #
    print("made it here 3")
    cf.execute()