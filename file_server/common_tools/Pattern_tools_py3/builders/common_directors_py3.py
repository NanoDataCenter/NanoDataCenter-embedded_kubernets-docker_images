
from common_tools.Pattern_tools_py3.factories.Handler_Factory_py3 import Handler_Factory
from common_tools.Pattern_tools_py3.factories.graph_search_py3 import common_package_search

def construct_all_handlers(site_data,qs,search_list,rpc_client=None,field_list=None,type_list=None):
    
    package = common_package_search(site_data,qs,search_list)
   
    factory = Handler_Factory(package[0],qs )
    handlers = factory.generate_handlers(field_list,type_list)
    if rpc_client != None:  #register rpc handler for sql types
       register_rpc_client(handlers,rpc_client)
    return handlers

def register_rpc_client( handlers,rpc_client):
    sql_handlers = {}
    sql_handlers["SEARCH_SQL_LOG_TABLE"] = True
    sql_handlers["SQL_LOG_TABLE"] = True
    for i,item in handlers.items():
      
      
       if item.properties["type"] in sql_handlers:
          item.set_rpc_handler(rpc_client)
          
    