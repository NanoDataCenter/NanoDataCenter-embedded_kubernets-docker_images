
class Cloud_Site_Definitons(object):

   def __init__(self,bc,cd):
       self.bc = bc
       self.cd = cd
       
       properties = {}
       properties["port"] = 443
       properties["https"]= True
       properties["debug"]= True
       properties["modules"] = ["monitoring","system_control","mqtt_client"]
       bc.add_info_node("WEB_SERVER","WEB_SERVER",properties = properties),
                    
       
       
       
       bc.add_header_node("CLOUD_SERVICE_QUEUE")
       cd.construct_package("CLOUD_SERVICE_QUEUE_DATA")
       cd.add_job_queue("CLOUD_JOB_SERVER",2048,forward=False)
       cd.add_hash("CLOUD_SUB_EVENTS")
       cd.close_package_contruction()
   
       bc.add_header_node("CLOUD_SERVICE_HOST_INTERFACE")
       bc.add_info_node( "HOST_INFORMATION","HOST_INFORMATION",properties={"host":"192.168.1.41" ,"port": 6379, "key_data_base": 6, "key":"_UPLOAD_QUEUE_" ,"depth":1024} )
       bc.end_header_node("CLOUD_SERVICE_HOST_INTERFACE")
       bc.end_header_node("CLOUD_SERVICE_QUEUE")
   
   
       bc.add_header_node("CLOUD_BLOCK_CHAIN_SERVER")
       cd.construct_package("CLOUD_BLOCK_CHAIN_SERVER")
       cd.add_rpc_server("BLOCK_CHAIN_RPC_SERVER",{"timeout":5,"queue":"BLOCK_CHAIN_RPC_SERVER"})
       cd.close_package_contruction()
       bc.end_header_node("CLOUD_BLOCK_CHAIN_SERVER")
   
 
       

       bc.add_header_node("TICKET_CONTROL")
       bc.add_info_node( "DATA_BASE","TICKET_CONTROL", properties = {"db":"SYSTEM_CONTROL"} )
       bc.add_info_node("TABLE","TICKET_CONTROL",properties = {"name":"TICKET_CONTROL",
                    "fields":["active","create_timestamp","close_timestamp","type","subtype","title","description","resolution"   ]} )
       bc.add_info_node("VALID_TYPES","TICKET_CONTROL",properties = {"types":["OTHERS","IRRIGATION_ISSUES","IRRIGATION_EQUIPMENT","TRIMMING","NON_IRRIGATION_FIXING"]})                   
       bc.end_header_node("TICKET_CONTROL")
   
       bc.add_header_node("TICKET_LOG")
       bc.add_info_node( "DATA_BASE","TICKET_LOG", properties = {"db":"SYSTEM_CONTROL"} )
       bc.add_info_node("TABLE","TICKET_LOG",properties = {"name":"TICKET_LOG",
                    "fields":["entry_timestamp","create_timestamp","close_timestamp","type","subtype","title","description","resolution"   ]} )
                      
       bc.end_header_node("TICKET_LOG")    
       monitoring_systems = ["CORE_OPS","MONITOR_REDIS","MONITOR_SQLITE","MONITOR_BLOCK_CHAIN","MONITOR_RPI_MOSQUITTO","MONITOR_RPI_MOSQUITTO_CLIENTS"]
       bc.add_info_node( "OP_MONITOR","OP_MONITOR", properties = {"OP_MONITOR_LIST":monitoring_systems} ) 
   
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
       properties["HOST"] = "192.168.1.51"
       properties["PORT"] = 1883
       properties["BASE_TOPIC"] = "/"
       bc.add_info_node( "MQTT_SERVER","MQTT_SERVER",properties=properties )
       self.add_mqtt_monitor()
       bc.end_header_node("MQTT_DEVICES")    
   
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