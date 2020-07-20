from op_monitor_lib.construct_monitor_py3 import Construct_Monitors
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers


class Op_Monitor(object):

   def __init__(self,site_data, qs,monitoring_list ,handlers):
       self.site_data = site_data
       self.qs        = qs

       self.construct_monitors = Construct_Monitors(site_data,qs,monitoring_list,handlers)   
       
       
   def execute_minute(self,*parameters):
       self.construct_monitors.execute_minute()

   def execute_15_minutes(self,*parameters):
       self.construct_monitors.execute_15_minutes()

   def execute_hour(self,*parameters):
       self.construct_monitors.execute_hour()
   
   def execute_day(self,*parameters):
       self.construct_monitors.execute_day()   

 
       
def construct_op_monitoring_instance( qs, site_data ):

                   
    
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )

    query_list = qs.add_match_terminal( query_list, 
                                        relationship = "OP_MONITOR" )
                                           
    data_sets, data_sources = qs.match_list(query_list)  
 
    data = data_sources[0]
    monitoring_list = data['OP_MONITOR_LIST']
    print("monitoring_list",monitoring_list)
  
    query_list = []
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=site_data["site"] )

    query_list = qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", property_mask={"name":"SYSTEM_MONITOR"} )
                                           
    package_sets, package_sources = qs.match_list(query_list)  
 
    package = package_sources[0]
    
    #
    #  do verifications of data package
    #
    #
    #
    data_structures = package["data_structures"]
   
   
    generate_handlers = Generate_Handlers( package, qs )
    
    handlers = {}
    handlers["SYSTEM_STATUS"] = generate_handlers.construct_hash(data_structures["SYSTEM_STATUS"])
    handlers["MONITORING_DATA"] = generate_handlers.construct_hash(data_structures["MONITORING_DATA"])
    handlers["SYSTEM_ALERTS"] = generate_handlers.construct_stream_writer(data_structures["SYSTEM_ALERTS"] )
    handlers["SYSTEM_PUSHED_ALERTS"]= generate_handlers.construct_stream_writer(data_structures["SYSTEM_PUSHED_ALERTS"] )
   
    handlers["SYSTEM_STATUS"].delete_all()
   
    op_monitor = Op_Monitor(site_data ,qs,monitoring_list,handlers )
    
    
    
  

    return op_monitor




def add_chains(redis_monitor, cf):

    cf.define_chain("minute_measurements", True)
    cf.insert.log("starting minute op monitoring")
    #cf.insert.one_step(op_monitor.execute_minute)
    #cf.insert.log("ending minute op monitoring")
    cf.insert.wait_event_count( event = "MINUTE_TICK")
    cf.insert.reset()

    cf.define_chain("fifteen_minute_measurements", True)
    cf.insert.log("starting 15_minute op monitoring")
    #cf.insert.one_step(op_monitor.execute_15_minutes)
    #cf.insert.log("ending 15_minute op monitoring")
    cf.insert.wait_event_count( event = "MINUTE_TICK",count=15)
    cf.insert.reset()

 
    cf.define_chain("hour_measurements", True)
    cf.insert.log("starting hour op monitoring")
    #cf.insert.one_step(op_monitor.execute_hour)
    #cf.insert.log("ending hour op monitoring")
    cf.insert.wait_event_count( event = "HOUR_TICK")
    cf.insert.reset()

    cf.define_chain("day_measurements", True)
    cf.insert.log("starting day op monitoring")
    #cf.insert.one_step(op_monitor.execute_day)
    #cf.insert.log("ending day op monitoring")
    cf.insert.wait_event_count( event = "DAY_TICK")
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
    
    #
    # Setup handle
    # open data stores instance
   
    qs = Query_Support( site_data )
    
    op_monitor = construct_op_monitoring_instance(qs, site_data )
    print("made it here 2")
    #
    # Adding chains
    #
    cf = CF_Base_Interpreter()
    add_chains(op_monitor, cf)
    #
    # Executing chains
    #
    print("made it here 3")
    cf.execute()