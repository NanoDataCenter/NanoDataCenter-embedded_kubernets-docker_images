class Start_Service(object):
    
    def __init__(self,bc,cd,name,container_run_string,container_image):
        properties = {}
        print("name",name,container_run_string)
        properties["command_list"] = container_run_string
        properties["container_image"] = container_image
        bc.add_header_node("SERVICE",name,properties=properties)
        
        
class End_Service(object):

     def __init__(self,bc,cd):
          print("end container")
          bc.end_header_node("SERVICE")
         