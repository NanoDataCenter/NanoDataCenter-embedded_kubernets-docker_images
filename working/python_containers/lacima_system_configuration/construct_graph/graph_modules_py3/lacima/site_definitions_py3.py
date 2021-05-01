
from .construct_weather_stations_py3 import Construct_Weather_Stations
from .construct_irrigation_scheduling_py3 import Construct_Irrigation_Scheduling_Control
from .plc_measurements_py3 import Construct_Lacima_PLC_Measurements
from .construct_plc_devices_py3 import Construct_Lacima_PLC_Devices
from .wifi_devices_py3 import Construct_Lacima_WIFI_Devices
from .main_web_site_definition_py3 import Generate_Main_Server_Web_Page_Definitions

class LACIMA_Site_Definitons(object):





   def __init__(self,bc,cd):
       self.bc = bc
       self.cd = cd
       
      
       properties = {}
       properties["port"] = 8080
       properties["https"]= False
       properties["debug"]= True
       properties["modules"] = ["monitoring","mqtt_client","eto","irrigation_scheduling","irrigation_control","modbus_control"]
       properties["status_function"] = "irrigation"
       
       
       bc.add_header_node("Web_Server_Definitions")
       temp = Generate_Main_Server_Web_Page_Definitions()
       menu_definitions = temp.generate_menu_objects()
       properties["menu"] = menu_definitions
       bc.add_info_node("WEB_SERVER","MAIN_WEB_SERVER",properties = properties)
       bc.end_header_node("Web_Server_Definitions")
       
     
       
       bc.add_header_node("CLOUD_SERVICE_QUEUE")
       cd.construct_package("CLOUD_SERVICE_QUEUE_DATA")
       cd.add_job_queue("CLOUD_JOB_SERVER",2048,forward=False)
       cd.add_hash("CLOUD_SUB_EVENTS")
       cd.close_package_contruction()
   
       bc.add_header_node("CLOUD_SERVICE_HOST_INTERFACE")
       bc.add_info_node( "HOST_INFORMATION","HOST_INFORMATION",properties={"host":"192.168.1.41" ,"port": 6379, "key_data_base": 6, "key":"_UPLOAD_QUEUE_" ,"depth":1024} )
       bc.end_header_node("CLOUD_SERVICE_HOST_INTERFACE")
       bc.end_header_node("CLOUD_SERVICE_QUEUE")       
       

 
       


   
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
       
       
   