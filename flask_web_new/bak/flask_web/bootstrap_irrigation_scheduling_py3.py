

from flask import render_template,jsonify
from flask import request, session, url_for

from irrigation_scheduling_py3.load_configuration_py3 import Load_Configuration_Data
class Irrigation_Scheduling(object):
    def __init__(base_self,self):
        Load_Configuration_Data(self.app, self.auth,render_template,request,self.file_server_library,self.url_rule_class,"Irrigation_Scheduling","irrigation_scheduling_py3")
            