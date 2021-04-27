class Construct_Lacima_WIFI_Devices(object):

   def __init__(self,bc,cd):
       device_number = 1
       properties = {}
       properties["ip"] = "192.168.1.110"
       properties["port"] = 6379
       properties["db"] = 11
       queues = {}
       queues["heart_beat"] = "HEART_BEAT"
       queues["reboot"] = "REBOOT"
       queues["client_input"] = "DEV_OUT"
       queues["internal_streams"] = "INTERNAL_STREAM"
       
       queues["external_streams"]="EXTERNAL_STREAM"
       properties["queues"] = queues
       bc.add_header_node("WIFI_DEVICES",properties=properties)
       cd.construct_package("WIFI_DATA_STRUCTURES")
       cd.add_hash("DEVICE_STATUS")
       cd.add_redis_stream("DEVICE_STATUS_LOG")
       
       cd.add_hash("UNKNOWN_DEVICES")
       cd.add_hash("INTERNAL_SENSORS")
       cd.add_hash("EXTERNAL_SENSORS")
       cd.add_redis_stream("INTERNAL_SENSORS_LOG",depth=device_number*26*78*12) # one week of data
       cd.add_redis_stream("EXTERNAL_SENSORS_LOG",depth=device_number*26*78*12) # one week of data
       cd.add_hash("WELL_SENSORS")
       cd.add_job_queue("ALERT_PUSH",25)
       cd.add_job_queue("TURN_ON_WATER",25)
       
       cd.close_package_contruction()
       
       properties= {}
       properties["node_id"] = "OP_1"
       properties["input_queue"] = "OP_1/INPUT"
       properties["type"] = "WATER_ALERT"
       properties["alert_queue"] = "TURN_ON_WATER"
       properties["input_queue_length"] = 1
       bc.add_info_node( "WIFI_DEVICE", properties["node_id"],properties=properties )
       bc.end_header_node("WIFI_DEVICES")  
   
   
   
   
