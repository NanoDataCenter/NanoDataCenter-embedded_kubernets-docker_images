class Start_Service(object):
    
    def __init__(self,bc,cd,name,container_run_string):
        properties = {}
        print("name",name,container_run_string)
        properties["command_list"] = container_run_string
        bc.add_header_node("SERVICE",name,properties=properties)
        
        
class End_Service(object):

     def __init__(self,bc,cd):
          print("end container")
          bc.end_header_node("SERVICE")
         