



from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers

class Processor_Performance_Stream_Base(object):
   def __init__(self,base_self):
       self.qs = base_self.qs
       self.site_data = base_self.site_data
       query_list = []
       query_list = base_self.qs.add_match_relationship( query_list,relationship="SITE",label=base_self.site_data["site"] )

       query_list = base_self.qs.add_match_terminal( query_list, 
                                        relationship = "PROCESSOR" )
                                           
       controller_sets, controller_nodes = base_self.qs.match_list(query_list)  
       self.controllers = []
       for i in controller_nodes:
           self.controllers.append(i["name"])
       self.controllers.sort()
     
       self.handlers = []
       for i in  self.controllers:
          self.handlers.append(self.assemble_processor_monitoring_data_structures(i))
 
   def assemble_processor_monitoring_data_structures(self,controller ):
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )

       query_list = self.qs.add_match_relationship( query_list, relationship = "PROCESSOR", label = controller )
       query_list = self.qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", property_mask={"name":"PROCESSOR_MONITORING"} )
                                           
       package_sets, package_sources = self.qs.match_list(query_list)  
     
       package = package_sources[0] 
       data_structures = package["data_structures"]
       generate_handlers = Generate_Handlers(package,self.qs)
       handlers = {}
       handlers["FREE_CPU"] = generate_handlers.construct_redis_stream_reader(data_structures["FREE_CPU"])
       handlers["RAM"] = generate_handlers.construct_redis_stream_reader(data_structures["RAM"])
       handlers["DISK_SPACE"] = generate_handlers.construct_redis_stream_reader(data_structures["DISK_SPACE"])
       handlers["TEMPERATURE"] = generate_handlers.construct_redis_stream_reader(data_structures["TEMPERATURE"])
       handlers["PROCESS_CPU"] = generate_handlers.construct_redis_stream_reader(data_structures["PROCESS_CPU"])
       
       handlers["CPU_CORE"] = generate_handlers.construct_redis_stream_reader(data_structures["CPU_CORE"])
       handlers["SWAP_SPACE"] = generate_handlers.construct_redis_stream_reader(data_structures["SWAP_SPACE"])
       handlers["IO_SPACE"] = generate_handlers.construct_redis_stream_reader(data_structures["IO_SPACE"])
       handlers["BLOCK_DEV"] = generate_handlers.construct_redis_stream_reader(data_structures["BLOCK_DEV"])
       handlers["CONTEXT_SWITCHES"] = generate_handlers.construct_redis_stream_reader(data_structures["CONTEXT_SWITCHES"])
       handlers["RUN_QUEUE"] = generate_handlers.construct_redis_stream_reader(data_structures["RUN_QUEUE"])
       handlers["EDEV"] = generate_handlers.construct_redis_stream_reader(data_structures["EDEV"])
       return handlers 
      
      
       
from templates.common.redis_streams.redis_multi_stream_manager_py3 import Redis_Multi_Stream_Manager


class Processor_Free_CPU(Processor_Performance_Stream_Base,Redis_Multi_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Multi_Stream_Manager.__init__(self,base_self,parameters)
       Processor_Performance_Stream_Base.__init__(self,base_self)

   def application_generation(self,controller_id,data):
    
     temp_data = self.handlers[controller_id]["FREE_CPU"].revrange("+","-" , count=1000)
     temp_data.reverse()
     chart_title = " Free CPU Profile for Linux Controller: "+self.controllers[controller_id]
       
     self.stream_keys,self.stream_range,self.stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0
    
 
class Processor_Free_Ram(Processor_Performance_Stream_Base,Redis_Multi_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Multi_Stream_Manager.__init__(self,base_self,parameters)
       Processor_Performance_Stream_Base.__init__(self,base_self)

   def application_generation(self,controller_id,data):
     temp_data = self.handlers[controller_id]["RAM"].revrange("+","-" , count=1000)
     temp_data.reverse()
     chart_title = " Free RAM Profile for Linux Controller: "+self.controllers[controller_id]
       
     self.stream_keys,self.stream_range,self.stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0
     
     
class Processor_Free_Disk(Processor_Performance_Stream_Base,Redis_Multi_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Multi_Stream_Manager.__init__(self,base_self,parameters)
       Processor_Performance_Stream_Base.__init__(self,base_self)

   def application_generation(self,controller_id,data):
     temp_data = self.handlers[controller_id]["DISK_SPACE"].revrange("+","-" , count=1000)
     temp_data.reverse()
     chart_title = " Disk Space Utilization Linux Controller: "+self.controllers[controller_id]
       
     self.stream_keys,self.stream_range,self.stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0
     
      
class Processor_Temperature(Processor_Performance_Stream_Base,Redis_Multi_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Multi_Stream_Manager.__init__(self,base_self,parameters)
       Processor_Performance_Stream_Base.__init__(self,base_self)

   def application_generation(self,controller_id,data):
     temp_data = self.handlers[controller_id]["TEMPERATURE"].revrange("+","-" , count=1000)
     temp_data.reverse()
     chart_title = " Temperature Profile for Linux Controller: "+self.controllers[controller_id]
       
     self.stream_keys,self.stream_range,self.stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0
         
class Processor_Cpu(Processor_Performance_Stream_Base,Redis_Multi_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Multi_Stream_Manager.__init__(self,base_self,parameters)
       Processor_Performance_Stream_Base.__init__(self,base_self)

   def application_generation(self,controller_id,data):
     temp_data = self.handlers[controller_id]["PROCESS_CPU"].revrange("+","-" , count=1000)
     temp_data.reverse()
     chart_title = " % Loading for Python Process: "+self.controllers[controller_id]
       
     self.stream_keys,self.stream_range,self.stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0
         
class Processor_Core(Processor_Performance_Stream_Base,Redis_Multi_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Multi_Stream_Manager.__init__(self,base_self,parameters)
       Processor_Performance_Stream_Base.__init__(self,base_self)

   def application_generation(self,controller_id,data):
     temp_data = self.handlers[controller_id]["CPU_CORE"].revrange("+","-" , count=1000)
     temp_data.reverse()
     chart_title = " Loading for CPU Core: "+self.controllers[controller_id]
       
     self.stream_keys,self.stream_range,self.stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0

class Processor_Swap_Space(Processor_Performance_Stream_Base,Redis_Multi_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Multi_Stream_Manager.__init__(self,base_self,parameters)
       Processor_Performance_Stream_Base.__init__(self,base_self)

   def application_generation(self,controller_id,data):
     temp_data = self.handlers[controller_id]["SWAP_SPACE"].revrange("+","-" , count=1000)
     temp_data.reverse()
     chart_title = " SWAP SPACE Used: "+self.controllers[controller_id]
       
     self.stream_keys,self.stream_range,self.stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0 


class Processor_Io_Space(Processor_Performance_Stream_Base,Redis_Multi_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Multi_Stream_Manager.__init__(self,base_self,parameters)
       Processor_Performance_Stream_Base.__init__(self,base_self)

   def application_generation(self,controller_id,data):
     temp_data = self.handlers[controller_id]["IO_SPACE"].revrange("+","-" , count=1000)
     temp_data.reverse()
     chart_title = " IO Space Activity: "+self.controllers[controller_id]
       
     self.stream_keys,self.stream_range,self.stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0


class Processor_Block_Dev(Processor_Performance_Stream_Base,Redis_Multi_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Multi_Stream_Manager.__init__(self,base_self,parameters)
       Processor_Performance_Stream_Base.__init__(self,base_self)

   def application_generation(self,controller_id,data):
     temp_data = self.handlers[controller_id]["BLOCK_DEV"].revrange("+","-" , count=1000)
     temp_data.reverse()
     chart_title = " Block Space Activity: "+self.controllers[controller_id]
       
     self.stream_keys,self.stream_range,self.stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0


class Processor_Context_Switches(Processor_Performance_Stream_Base,Redis_Multi_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Multi_Stream_Manager.__init__(self,base_self,parameters)
       Processor_Performance_Stream_Base.__init__(self,base_self)

   def application_generation(self,controller_id,data):
     temp_data = self.handlers[controller_id]["CONTEXT_SWITCHES"].revrange("+","-" , count=1000)
     temp_data.reverse()
     chart_title = " Context Switches: "+self.controllers[controller_id]
       
     self.stream_keys,self.stream_range,self.stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0
 

class Processor_Run_Queue(Processor_Performance_Stream_Base,Redis_Multi_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Multi_Stream_Manager.__init__(self,base_self,parameters)
       Processor_Performance_Stream_Base.__init__(self,base_self)

   def application_generation(self,controller_id,data):
     temp_data = self.handlers[controller_id]["RUN_QUEUE"].revrange("+","-" , count=1000)
     temp_data.reverse()
     chart_title = " RUN QUEUE Activity: "+self.controllers[controller_id]
       
     self.stream_keys,self.stream_range,self.stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0


class Processor_Edev(Processor_Performance_Stream_Base,Redis_Multi_Stream_Manager):
   def __init__(self,base_self,parameters):
       Redis_Multi_Stream_Manager.__init__(self,base_self,parameters)
       Processor_Performance_Stream_Base.__init__(self,base_self)

   def application_generation(self,controller_id,data):
     temp_data = self.handlers[controller_id]["EDEV"].revrange("+","-" , count=1000)
     temp_data.reverse()
     chart_title = " NETWORK DEVICE ERRORS: for "+self.controllers[controller_id]
     self.stream_keys,self.stream_range,self.stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
     self.title = self.stream_keys
     self.max_value = 10000000
     self.min_value = 0     


 