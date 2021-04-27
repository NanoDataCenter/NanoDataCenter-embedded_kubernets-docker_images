
import time
from Pattern_tools_py3.builders.common_directors_py3 import construct_all_handlers
from Pattern_tools_py3.factories.graph_search_py3 import common_qs_search 
from Pattern_tools_py3.factories.iterators_py3 import form_dictionary_from_list

class Column_Control(object):

    def __init__(self, site_data,qs, column_search_list, package_search_list= None ):
        # find controller data
        
        data = common_qs_search(site_data,qs,column_search_list)
        self.column_data = data[0]
        self.class_name = self.column_data["class"]
        column_search_list.append("CLASS_DEF")
        
        
        data = common_qs_search(site_data,qs,column_search_list)
        self.class_dict = form_dictionary_from_list(data,"name")
        print(self.class_dict.keys())
        
        if package_search_list != None:
           handlers = construct_all_handlers(site_data,qs,package_search_list)
        print("handlers",handlers)           
            
        
        exit()
        self.class_map = {}
        # instanciate classes
    


    def construct_classes(self,


    
    def process_event(self):
        while True:
           time.sleep(1)
           
           
           
           

if __name__ == "__main__":
   from redis_support_py3.graph_query_support_py3 import  Query_Support
   from Pattern_tools_py3.factories.get_site_data_py3 import get_site_data
   
   site_data = get_site_data("/mnt/ssd/site_config/redis_server.json")
   qs = Query_Support( site_data )
   column_control = Column_Control(site_data,qs,[["COLUMN_DEFINITIONS","label"],["COLUMN_ELEMENT","test_column" ]],[["COLUMN_DEFINITIONS","label"],"WIFI_DATA_STRUCTURES"] )
   

           
