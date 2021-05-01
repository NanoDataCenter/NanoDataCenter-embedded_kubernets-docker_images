class Start_Service(object):
    
    def __init__(self,bc,cd,name,startup_command,container_image):
        properties = {}
        #print("name",name,container_run_string)
        properties["container_image"] = container_image
       
        properties["startup_command"] = startup_command
        bc.add_header_node("CONTAINER",name,properties=properties)
        cd.construct_package("DATA_STRUCTURES")
        cd.add_redis_stream("ERROR_STREAM",forward=False)
        cd.add_hash("ERROR_HASH")
        cd.add_job_queue("WEB_COMMAND_QUEUE",1)
        cd.add_hash("WEB_DISPLAY_DICTIONARY")
        cd.add_redis_stream("PROCESS_VSZ")
        cd.add_redis_stream("PROCESS_RSS")
        cd.add_redis_stream("PROCESS_CPU") 
        cd.add_hash("PROCESS_CONTROL")
        cd.add_redis_stream("POWER_UP_LOG")
        cd.close_package_contruction()
        
        
class End_Service(object):

     def __init__(self,bc,cd):
          #print("end container")
          bc.end_header_node("CONTAINER")
         