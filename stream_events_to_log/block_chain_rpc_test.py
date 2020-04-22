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
 
    from redis_support_py3.graph_query_support_py3 import  Query_Support
    from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
    import datetime
    

    #send_rpc_message( self, method,parameters,timeout=30 ):

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
    query_list = qs.add_match_relationship( query_list,relationship="SITE",label=redis_site ["site"] )
    query_list = qs.add_match_relationship( query_list,relationship= "CLOUD_BLOCK_CHAIN_SERVER" )
    query_list = qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", property_mask={"name":"CLOUD_BLOCK_CHAIN_SERVER"} )
    package_sets, package_sources = qs.match_list(query_list) 
    package = package_sources[0]    
    data_structures = package["data_structures"]
    generate_handlers = Generate_Handlers( package, qs )
    rpc_client = generate_handlers.construct_rpc_client()
    queue_name = data_structures["BLOCK_CHAIN_RPC_SERVER"]['queue']
    print("queue_name",queue_name)
    
    rpc_client.set_rpc_queue(queue_name)

    parameters = {"contract_name":"EventHandler",  "start_block":0,"end_block":'latest'}
    data = rpc_client.send_rpc_message("fetch_block_chain_data",parameters) 
    print("length",len(data))
    print("\n\n")
    print("first element ",data[0])
    print("\n\n")
    print("last element ",data[-1])
    data = rpc_client.send_rpc_message("fetch_current_block_number",[])
    print(data)
     
       
       
  
