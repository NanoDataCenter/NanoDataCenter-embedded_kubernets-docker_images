
class Cloud_Site_Mqtt_Interface(object):

   def __init__(self,bc,cd):
   
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
       cd.close_package_contruction()
       properties = {}
       properties["HOST"] = "192.168.1.51"
       properties["PORT"] = 1883
       properties["BASE_TOPIC"] = "/"
       self.bc.add_info_node( "MQTT_SERVER","MQTT_SERVER",properties=properties )
       bc.end_header_node("MQTT_DEVICES")    
   
   