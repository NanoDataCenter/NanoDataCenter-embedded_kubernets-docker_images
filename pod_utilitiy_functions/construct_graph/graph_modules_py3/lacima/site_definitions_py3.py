
from .construct_weather_stations_py3 import Construct_Weather_Stations
from .construct_irrigation_scheduling_py3 import Construct_Irrigation_Scheduling_Control
from .plc_measurements_py3 import Construct_Lacima_PLC_Measurements
from .construct_plc_devices_py3 import Construct_Lacima_PLC_Devices
from .wifi_devices_py3 import Construct_Lacima_WIFI_Devices
from graph_modules_py3.web_menu_constructors_py3 import Menu_Header
from graph_modules_py3.web_menu_constructors_py3 import  Menu_Element


class Generate_Web_Page_Definitions( object):    
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
       system_menu.add_child(service_menu)
      
       top_menu.add_child(irrigation_menu)
       top_menu.add_child(system_menu)
       top_menu.add_child(Menu_Element("site_map","common/site_manager"))
       top_menu.add_child(Menu_Element("/","common/site_manager"))
       return_value = top_menu.generate_data_structure()
       
       self.diagnos_list_top(return_value)

       return return_value

class LACIMA_Site_Definitons(object):





   def __init__(self,bc,cd):
       self.bc = bc
       self.cd = cd
       
       monitoring_systems = ["CORE_OPS","MONITOR_REDIS","MONITOR_RPI_MOSQUITTO","MONITOR_RPI_MOSQUITTO_CLIENTS"]
       bc.add_info_node( "OP_MONITOR","OP_MONITOR", properties = {"OP_MONITOR_LIST":monitoring_systems} ) 
       
       properties = {}
       properties["port"] = 8080
       properties["https"]= False
       properties["debug"]= True
       properties["modules"] = ["monitoring","mqtt_client","eto","irrigation_scheduling","irrigation_control","modbus_control"]
       properties["status_function"] = "irrigation"
       
       
       temp = Generate_Web_Page_Definitions()
       menu_definitions = temp.generate_menu_objects()
       properties["menu"] = menu_definitions
       
       
       bc.add_info_node("WEB_SERVER","WEB_SERVER",properties = properties)
       
       
      
       
       bc.add_header_node("CLOUD_SERVICE_QUEUE")
       cd.construct_package("CLOUD_SERVICE_QUEUE_DATA")
       cd.add_job_queue("CLOUD_JOB_SERVER",2048,forward=False)
       cd.add_hash("CLOUD_SUB_EVENTS")
       cd.close_package_contruction()
   
       bc.add_header_node("CLOUD_SERVICE_HOST_INTERFACE")
       bc.add_info_node( "HOST_INFORMATION","HOST_INFORMATION",properties={"host":"192.168.1.41" ,"port": 6379, "key_data_base": 6, "key":"_UPLOAD_QUEUE_" ,"depth":1024} )
       bc.end_header_node("CLOUD_SERVICE_HOST_INTERFACE")
       bc.end_header_node("CLOUD_SERVICE_QUEUE")       
       

 
       bc.add_header_node("SYSTEM_MONITOR")
       cd.construct_package("SYSTEM_MONITOR")      
       #cd.add_managed_hash(self,name,fields,forward=False) perfored way to store field how to get field in system
       cd.add_hash("SYSTEM_STATUS")
   
       cd.add_hash("MONITORING_DATA")
       cd.add_redis_stream("SYSTEM_ALERTS")
       cd.add_redis_stream("SYSTEM_PUSHED_ALERTS")
       cd.close_package_contruction()
       bc.end_header_node("SYSTEM_MONITOR")


   
       bc.add_header_node("MQTT_DEVICES")
       cd.construct_package("MQTT_DEVICES_DATA")
       cd.add_redis_stream("MQTT_INPUT_QUEUE",50000)
       cd.add_redis_stream("MQTT_SENSOR_QUEUE",10000)
       cd.add_redis_stream("MQTT_PAST_ACTION_QUEUE",300)
       cd.add_hash("MQTT_SENSOR_STATUS")
       cd.add_hash("MQTT_DEVICES")
       cd.add_hash("MQTT_SUBSCRIPTIONS")
       cd.add_hash("MQTT_CONTACT_LOG")
       cd.add_hash("MQTT_UNKNOWN_DEVICES")
       cd.add_hash("MQTT_UNKNOWN_SUBSCRIPTIONS")
       cd.add_hash("MQTT_REBOOT_LOG")
       cd.add_hash("MQTT_SERVER_STATE")
       cd.add_job_queue("MQTT_PUBLISH_QUEUE",depth= 50,forward = False)
       cd.close_package_contruction()
       properties = {}
       properties["HOST"] = "192.168.1.110"
       properties["PORT"] = 1883
       properties["BASE_TOPIC"] = "/"
       bc.add_info_node( "MQTT_SERVER","MQTT_SERVER",properties=properties )
       self.add_mqtt_monitor()
       bc.end_header_node("MQTT_DEVICES")  
       
       
       bc.add_header_node("FILE_SERVER")
       cd.construct_package("FILE_SERVER")
       cd.add_rpc_server("FILE_SERVER_RPC_SERVER",{"timeout":5,"queue":"FILE_RPC_SERVER"})

       cd.close_package_contruction()
       bc.end_header_node("FILE_SERVER")
       
       Construct_Weather_Stations(bc,cd)  
       Construct_Irrigation_Scheduling_Control(bc,cd) 
       Construct_Lacima_PLC_Measurements(bc,cd)       
       Construct_Lacima_PLC_Devices(bc,cd)
       Construct_Lacima_WIFI_Devices(bc,cd)
   
   def add_mqtt_monitor(self):
       mqtt_tag = "MQTT_SERVER_CHECK"
       properties = {}
       properties["REBOOT_FLAG"] = True
       properties["REBOOT_KEY"] = "REBOOT"
       properties["type"] = "MQTT_MONITOR"
       properties["HEART_BEAT"] = "HEART_BEAT"
       properties["HEART_BEAT_TIME_OUT"] = 120
       properties["topic"] = mqtt_tag
       properties["null_commands"] = {}
       properties["subscriptions"] = {}
       properties["subscriptions"]["REBOOT"] = True
       properties["subscriptions"]["HEART_BEAT"] = True
       properties["subscriptions"]["SERVER_CHECK"] = True
       self.bc.add_info_node( "MQTT_DEVICE",mqtt_tag,properties=properties )
       
       
   