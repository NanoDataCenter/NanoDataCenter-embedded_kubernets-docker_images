
from Pattern_tools_py3.factories.Handler_Factory_py3 import Handler_Factory
from Pattern_tools_py3.factories.graph_search_py3 import common_package_search

def construct_all_handlers(site_data,qs,search_list):
    package = common_package_search(site_data,qs,search_list)
    handler = Handler_Factory(package[0],qs )
    return handler.generate_all()

