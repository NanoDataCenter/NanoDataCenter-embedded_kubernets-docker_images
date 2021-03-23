class Start_Container(object):
    
    def __init__(self,bc,cd,name,startup_command,command_list,container_image):
        properties = {}
        #print("name",name,startup_command,command_list,container_image)
       
        properties["container_image"] = container_image
        properties["command_list"] = command_list
        properties["startup_command"] = startup_command
        bc.add_header_node("CONTAINER",name,properties=properties)
        cd.construct_package("DATA_STRUCTURES")
        cd.add_redis_stream("ERROR_STREAM")   
        cd.add_redis_stream("PROCESS_VSZ")
        cd.add_redis_stream("PROCESS_RSS")
        cd.add_redis_stream("PROCESS_CPU") 
        cd.close_package_contruction()
        
class End_Container(object):

     def __init__(self,bc,cd):
          #print("end container")
          bc.end_header_node("CONTAINER")
          