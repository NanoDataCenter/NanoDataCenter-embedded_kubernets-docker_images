from graph_modules_py3.web_menu_constructors_py3 import Menu_Header
from graph_modules_py3.web_menu_constructors_py3 import  Menu_Element


class Generate_Main_Server_Web_Page_Definitions( object):    
   def __init__(self):
       pass

   def diagnos_list(self,path,data):
       for i in data:
           
           if i["type"] == True:
                print(path,i["display_name"] )
                self.diagnos_list(path+"/"+i["display_name"],i["children"])
           else:
                print(path,i["display_name"],i["class_name"] )
               
 

 
   def diagnos_list_top(self,i):  #### for diagonostics only
       
       
       if i["type"] == True:
          print("display_name",i["display_name"] )
          self.diagnos_list("/top",i["children"])
       else:
           print("display_name",i["display_name"] )
           print("class_name",i["class_name"] )
 
   def generate_pod_logs(self):
       pod_logs = Menu_Header(display_name="pod_logs") 
       pod_logs.add_child(Menu_Element('display_exception_status',"pod_control/view_exception_status"))
       pod_logs.add_child(Menu_Element('display_exception_logs',"pod_control/view_exception_log"))
       return pod_logs        

   def generate_pod_menus(self):
       pod_menu = Menu_Header(display_name="pod_management") 
       pod_menu.add_child(Menu_Element('start_stop_processes',"pod_control/process_control"))
       pod_menu.add_child(self.generate_pod_logs())
       return pod_menu          

   def generate_container_logs(self):
       container_logs = Menu_Header(display_name="container_logs") 
       container_logs.add_child(Menu_Element('display_exception_status',"manage_containers/view_exception_status"))
       container_logs.add_child(Menu_Element('display_exception_logs',"manage_containers/view_exception_log"))
       container_logs.add_child(Menu_Element('display_cpu_loading',"manage_containers/cpu_loading"))
       container_logs.add_child(Menu_Element('display_vsz_useage',"manage_containers/vsz"))
       container_logs.add_child(Menu_Element('display_rss_useage',"manage_containers/rss"))
       return container_logs        

   def generate_container_menus(self):
       container_menu = Menu_Header(display_name="container_management") 
       container_menu.add_child(Menu_Element('start_stop_containers',"manage_containers/start_and_stop_containers"))
       container_menu.add_child(Menu_Element('start_stop_container_processes',"manage_containers/start_and_stop_processes"))
       container_menu.add_child(self.generate_container_logs())
       return container_menu         
                      

   def generate_processor_menus(self):
       processor_menu = Menu_Header(display_name="processor_data") 
      
   
       processor_menu.add_child(Menu_Element('free_cpu',"processor/free_cpu"))
       processor_menu.add_child(Menu_Element('ram',"processor/ram"))
       processor_menu.add_child(Menu_Element('disk_space',"processor/disk_space"))
       processor_menu.add_child(Menu_Element( 'temperature',"processor/temperature"))
       processor_menu.add_child(Menu_Element( 'cpu_core',"processor/cpu_core"))
       processor_menu.add_child(Menu_Element('swap_space',"processor/swap_space"))
       
       processor_menu.add_child(Menu_Element('io_space',"processor/io_space"))
       processor_menu.add_child(Menu_Element('block_dev',"processor/block_dev"))
       processor_menu.add_child(Menu_Element('context_switches',"processor/context_switches"))
       processor_menu.add_child(Menu_Element( 'run_queue',"processor/run_queue"))
       processor_menu.add_child(Menu_Element('edev',"processor/edev"))
        
       return processor_menu         
   
    
      
   def generate_redis_menus(self):
       redis_menu = Menu_Header(display_name="redis_data") 
      
   
       redis_menu.add_child(Menu_Element("redis_key_stream","redis/key_stream"))
       redis_menu.add_child(Menu_Element("redis_client_stream","redis/client_stream"))
       redis_menu.add_child(Menu_Element("redis_memory_stream","redis/memory_stream"))
       redis_menu.add_child(Menu_Element("redis_call_stream","redis/call_stream"))
       redis_menu.add_child(Menu_Element("redis_cmd_stream","redis/command_time"))
       redis_menu.add_child(Menu_Element("redis_server_time","redis/server_time"))
       return redis_menu         


 
   def generate_menu_objects(self):
       top_menu = Menu_Header(display_name="") # parameters do not matter for top menu
       irrigation_menu = Menu_Header(display_name="irrigation_functions") 
    
       irrigation_menu.add_child(Menu_Element("control_by_schedule","irrigation/control_by_schedule"))
       irrigation_menu.add_child(Menu_Element("control_by_controler_pin","irrigation/control_by_controller"))
       irrigation_menu.add_child(Menu_Element("control_by_valve_group","irrigation/control_by_valve_group"))
       irrigation_menu.add_child(Menu_Element("set_irrigation_parameters","irrigation/set_irrigation_parameters"))
       irrigation_menu.add_child(Menu_Element("control_irrigation_queue","irrigation/control_irrigation_queue"))
       irrigation_menu.add_child(Menu_Element("list_past_actions","irrigation/past_actions"))
       irrigation_menu.add_child(Menu_Element("irrigation_time_history","irrigation/stream_manager"))
       
       
       service_menu = Menu_Header(display_name="services")
       service_menu.add_child(self.generate_redis_menus())
    
       system_menu = Menu_Header(display_name="system_functions") 
       system_menu.add_child(self.generate_processor_menus())
       system_menu.add_child(self.generate_container_menus())
       system_menu.add_child(self.generate_pod_menus())
       system_menu.add_child(service_menu)
      
       top_menu.add_child(irrigation_menu)
       top_menu.add_child(system_menu)
       top_menu.add_child(Menu_Element("site_map","common/site_manager"))
       top_menu.add_child(Menu_Element("/","common/site_manager"))
       return_value = top_menu.generate_data_structure()
       
       self.diagnos_list_top(return_value)

       return return_value