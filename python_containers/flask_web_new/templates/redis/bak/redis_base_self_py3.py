

from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers

class Redis_Stream_Base(object):
    def __init__(self,base_self):
       qs = base_self.qs
       query_list = []
       query_list = qs.add_match_relationship( query_list,relationship="SITE",label=base_self.site_data["site"] )
       query_list = qs.add_match_relationship( query_list,relationship="CONTAINER",label="monitor_redis" )
       
       
       query_list = qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", label = "REDIS_MONITORING" )
                                           
       package_sets, package_sources = qs.match_list(query_list)  
      
       package = package_sources[0]
       generate_handlers = Generate_Handlers(package,qs)
       data_structures = package["data_structures"]     
 
       self.handlers = {}
       self.handlers["REDIS_MONITOR_KEY_STREAM"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_KEY_STREAM"])
       self.handlers["REDIS_MONITOR_CLIENT_STREAM"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_CLIENT_STREAM"])
       self.handlers["REDIS_MONITOR_MEMORY_STREAM"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_MEMORY_STREAM"])
       self.handlers["REDIS_MONITOR_CALL_STREAM"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_CALL_STREAM"])
       self.handlers["REDIS_MONITOR_CMD_TIME_STREAM"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_CMD_TIME_STREAM"])
       self.handlers["REDIS_MONITOR_SERVER_TIME"] = generate_handlers.construct_redis_stream_reader(data_structures["REDIS_MONITOR_SERVER_TIME"])
       