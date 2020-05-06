
import os
import json
from datetime import datetime
import time
from base_stream_processing.base_stream_processing_py3 import Base_Stream_Processing
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
class Processor_Monitoring(Base_Stream_Processing):
   def __init__( self, app, auth, request, render_template,qs,site_data,url_rule_class,subsystem_name,path):
       self.app      = app
       self.auth     = auth
       self.request  = request
       self.render_template = render_template
       self.path = path
       self.qs = qs
       self.site_data = site_data
       self.url_rule_class = url_rule_class
       self.subsystem_name = subsystem_name
       
       self.assemble_handlers()
       self.assemble_url_rules()
       self.path_dest = self.url_rule_class.move_directories(self.path)   
       
       
   def assemble_url_rules(self):
    
       self.slash_name = "/"+self.subsystem_name
       self.menu_data = {}
       self.menu_list = []
       
       function_list =   [ self.free_cpu,
                           self.ram,
                           self.disk_space,
                           self.temperature,                       
                           self.process_cpu_core,
                           self.process_swap_space,
                           self.process_io_space ,
                           self.process_block_dev ,
                           self.process_context_switches ,
                           self.process_run_queue, 
                           self.process_dev ,
                           self.process_sock,
                           self.process_tcp ,
                           self.process_udp ]
      
      
       url_list = [
                      [ 'free_cpu' ,'/<int:controller_id>','/0',"CPU Utilization"  ],
                      [ 'ram' ,'/<int:controller_id>','/0',"Display Ram Utilitization"  ],
                      [ 'disk_space' ,'/<int:controller_id>','/0',"Display Disk Space Utilization"   ],
                      [ 'temperature' ,'/<int:controller_id>','/0',"Display Temperature History"  ],    
                      [ 'cpu_core','/<int:controller_id>','/0',"Display CPU Core Loading"  ],
                      [ 'swap_space','/<int:controller_id>','/0',"Display Swap Space Loading"  ],
                      [ 'io_space','/<int:controller_id>','/0',"Display IO Space Loading"  ],
                      [ 'block_dev' ,'/<int:controller_id>','/0',"Display Block Device Loading"  ],
                      [ 'context_switches' ,'/<int:controller_id>','/0',"Display Context Switch Loading" ],
                      [ 'run_queue','/<int:controller_id>','/0',"Display Run Queue Loading"  ],
                      [ 'dev'  ,'/<int:controller_id>','/0',"Network Device Errors" ],
                      [ 'sock' ,'/<int:controller_id>','/0',"Display Socket Loading" ],
                      [ 'tcp'  ,'/<int:controller_id>','/0',"Display TCP Loading"  ],
                      [ 'udp'  ,'/<int:controller_id>','/0',"Display UDP Loading"  ]
        ]

       self.url_rule_class.add_get_rules(self.subsystem_name,function_list,url_list)
        
  



   def assemble_handlers(self):
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )

       query_list = self.qs.add_match_terminal( query_list, 
                                        relationship = "PROCESSOR" )
                                           
       controller_sets, controller_nodes = self.qs.match_list(query_list)  
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
                                        relationship = "PACKAGE", property_mask={"name":"SYSTEM_MONITORING"} )
                                           
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
       handlers["DEV"] = generate_handlers.construct_redis_stream_reader(data_structures["DEV"])
       handlers["SOCK"] = generate_handlers.construct_redis_stream_reader(data_structures["SOCK"])
       handlers["TCP"] = generate_handlers.construct_redis_stream_reader(data_structures["TCP"])
       handlers["UDP"] = generate_handlers.construct_redis_stream_reader(data_structures["UDP"])
       return handlers



   def free_cpu(self,controller_id):
       
       
       temp_data = self.handlers[controller_id]["FREE_CPU"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " Free CPU Profile for Linux Controller: "+self.controllers[controller_id]
       
       stream_keys,stream_range,stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       
       
       return self.render_template( "streams/stream_multi_controller",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.controllers,
                                     controller_id = controller_id
                              
             
                                     )



   def ram(self,controller_id):
      
       
       temp_data = self.handlers[controller_id]["RAM"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " Free RAM Profile for Linux Controller: "+self.controllers[controller_id]
      
       stream_keys,stream_range,stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       
      
       return self.render_template( "streams/stream_multi_controller",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.controllers,
                                     controller_id = controller_id
                                     
                                     )
            


   def disk_space(self,controller_id):
       
       temp_data = self.handlers[controller_id]["DISK_SPACE"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " Disk Space Utilization Linux Controller: "+self.controllers[controller_id]
       
       stream_keys,stream_range,stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       
       
       return self.render_template( "streams/stream_multi_controller",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.controllers,
                                     controller_id = controller_id
                                     
                                     )
            

   def temperature(self,controller_id):
       
       temp_data = self.handlers[controller_id]["TEMPERATURE"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " Temperature Profile for Linux Controller: "+self.controllers[controller_id]
       stream_keys,stream_range,stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       

       return self.render_template( "streams/stream_multi_controller",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.controllers,
                                     controller_id = controller_id
                                     
                                     )

        
 


   def process_cpu(self,controller_id):
       temp_data = self.handlers[controller_id]["PROCESS_CPU"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = "% Loading for Python Process: "
       stream_keys,stream_range,stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       

       return self.render_template( "streams/stream_multi_controller",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.controllers,
                                     controller_id = controller_id
                                     
                                     )


 
   def process_cpu_core(self,controller_id):
       temp_data = self.handlers[controller_id]["CPU_CORE"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " Temperature Profile for CPU Core: "+self.controllers[controller_id]
       stream_keys,stream_range,stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       

       return self.render_template( "streams/stream_multi_controller",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.controllers,
                                     controller_id = controller_id
                                     
                                     )



   def process_swap_space(self,controller_id):
       temp_data = self.handlers[controller_id]["SWAP_SPACE"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " SWAP SPACE Used: "+self.controllers[controller_id]
       stream_keys,stream_range,stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       

       return self.render_template( "streams/stream_multi_controller",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.controllers,
                                     controller_id = controller_id
                                     
                                     )


   def process_io_space(self,controller_id):
       temp_data = self.handlers[controller_id]["IO_SPACE"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " IO Space Activity: "+self.controllers[controller_id]
       stream_keys,stream_range,stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       

       return self.render_template( "streams/stream_multi_controller",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.controllers,
                                     controller_id = controller_id
                                     
                                     )


   def process_block_dev(self,controller_id):
       temp_data = self.handlers[controller_id]["BLOCK_DEV"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " Block Space Activity: "
       stream_keys,stream_range,stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       

       return self.render_template( "streams/stream_multi_controller",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.controllers,
                                     controller_id = controller_id
                                     
                                     )


   def process_context_switches(self,controller_id):
       temp_data = self.handlers[controller_id]["CONTEXT_SWITCHES"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " Context Switches: "+self.controllers[controller_id]
       stream_keys,stream_range,stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       

       return self.render_template( "streams/stream_multi_controller",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.controllers,
                                     controller_id = controller_id
                                     
                                     )



   def process_run_queue(self,controller_id):
       temp_data = self.handlers[controller_id]["RUN_QUEUE"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " RUN QUEUE Activity: "+self.controllers[controller_id]
       stream_keys,stream_range,stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       

       return self.render_template( "streams/stream_multi_controller",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.controllers,
                                     controller_id = controller_id
                                     
                                     )

      
   def process_dev(self,controller_id):
       temp_data = self.handlers[controller_id]["DEV"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " NETWORK DEVICE ERRORS: "
       stream_keys,stream_range,stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       

       return self.render_template( "streams/stream_multi_controller",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.controllers,
                                     controller_id = controller_id
                                     
                                     )


      

   def process_sock(self,controller_id):
       temp_data = self.handlers[controller_id]["SOCK"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " IO Space Activity: "+self.controllers[controller_id]
       stream_keys,stream_range,stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       

       return self.render_template( "streams/stream_multi_controller",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.controllers,
                                     controller_id = controller_id
                                     
                                     )



   def process_tcp(self,controller_id):
       temp_data = self.handlers[controller_id]["TCP"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " TCP Activity: "+self.controllers[controller_id]
       stream_keys,stream_range,stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       

       return self.render_template( "streams/stream_multi_controller",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.controllers,
                                     controller_id = controller_id
                                     
                                     )


       

   def process_udp(self,controller_id):
       temp_data = self.handlers[controller_id]["UDP"].revrange("+","-" , count=1000)
       temp_data.reverse()
       chart_title = " UDP Activity: "+self.controllers[controller_id]
       stream_keys,stream_range,stream_data = self.format_data(temp_data,title=chart_title,title_y="Deg F",title_x="Date")
       

       return self.render_template( "streams/stream_multi_controller",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     controllers = self.controllers,
                                     controller_id = controller_id
                                     
                                     )

 
