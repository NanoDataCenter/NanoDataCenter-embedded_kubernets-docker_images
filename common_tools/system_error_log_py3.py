

import time
from Pattern_tools_py3.builders.common_directors_py3 import construct_all_handlers
import sys

class System_Error_Logging(object):
   def __init__(self,qs,container,site_data):
       self.processor = site_data["local_node"]
       self.container = container
       self.python_file = sys.argv[0]
       self.site_data = site_data
       self.qs = qs
       
       search_list = ["SYSTEM_MONITOR","SYSTEM_MONITOR"]
        
       self.handlers = construct_all_handlers(site_data,qs,search_list)
       self.log_error_message("Reboot")



        
	
   def log_error_message(self,verb,subject=None,obj_of=None):
        log_value = {}
        message = self.format_log(verb,subject,obj_of)
        self.handlers["SYSTEM_VERBS"].hset(verb,time.time())
        log_value["processor"] = self.processor
        log_value["python_file"] = self.python_file
        log_value["container"] = self.container
        log_value["error_msg"] = message
        log_value["time"] = time.time()

        self.handlers["SYSTEM_ALERTS"].push( log_value )
        


   def format_log(self,verb,subject,obj_of):
       if subject == None:
            subject = "blank"
       if obj_of == None:
            obj_of = "blank"
       if type(subject) == list:
           subject = "~".join([str(subject[0]),str(subject[1])])
       if type(obj_of) == list:
            obj_of = "~".join([str(obj_of[0]),str(obj_of[1])])
       return ":".join([str(verb),str(subject),str(obj_of)])
'''
test code
 print(self.rpc_client.list_data_bases())
       print(self.rpc_client.list_tables('default'))
       self.log_error_message("test_message")
       print(self.handlers["SYSTEM_ALERTS"].select())
       print("where ###################################")
       print(self.handlers["SYSTEM_ALERTS"].select("time> "+str(time.time()-1)))
       print(self.handlers["SYSTEM_ALERTS"].select_text_search_general("message"))
       print(self.handlers["SYSTEM_ALERTS"].select_text_search_general("not+ here"))
       self.handlers["SYSTEM_ALERTS"].trim_stream(time.time()+1)
       print(self.handlers["SYSTEM_ALERTS"].select())
       quit()
'''