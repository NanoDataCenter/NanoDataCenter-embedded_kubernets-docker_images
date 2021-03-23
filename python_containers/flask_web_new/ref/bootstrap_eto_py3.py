#
#
#  File: flask_web_py3.py
#
#
#
import os
import json
import redis
import urllib
from flask import render_template,jsonify
from flask import request, session, url_for


from  eto_py3.load_eto_management_py3 import Load_ETO_Management_Web
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers


class ETO_Management(object):

   def __init__(base_self ,self  ):
       
        # from graph get hash tables
       
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
       query_list = self.qs.add_match_terminal( query_list, 
                                        relationship = "WS_STATION" )
                                        
       eto_sets, eto_sources = self.qs.match_list(query_list)                                    
    
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )

       query_list = self.qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", property_mask={"name":"WEATHER_STATION_DATA"} )
                                           
       package_sets, package_sources = self.qs.match_list(query_list)  
     
       package = package_sources[0] 
       data_structures = package["data_structures"]
       generate_handlers = Generate_Handlers(package,self.qs)
       self.ds_handlers = {}
       self.ds_handlers["EXCEPTION_VALUES"] = generate_handlers.construct_hash(data_structures["EXCEPTION_VALUES"])
       self.ds_handlers["ETO_VALUES"] = generate_handlers.construct_hash(data_structures["ETO_VALUES"])
       self.ds_handlers["RAIN_VALUES"] = generate_handlers.construct_hash(data_structures["RAIN_VALUES"])
       self.ds_handlers["ETO_CONTROL"] = generate_handlers.construct_hash(data_structures["ETO_CONTROL"])
       self.ds_handlers["ETO_HISTORY"] = generate_handlers.construct_redis_stream_reader(data_structures["ETO_HISTORY"])
       self.ds_handlers["RAIN_HISTORY"] = generate_handlers.construct_redis_stream_reader(data_structures["RAIN_HISTORY"] )
       self.ds_handlers["EXCEPTION_LOG"] = generate_handlers.construct_redis_stream_reader(data_structures["EXCEPTION_LOG"] )
       self.ds_handlers["ETO_ACCUMULATION_TABLE"] = generate_handlers.construct_hash(data_structures["ETO_ACCUMULATION_TABLE"])
       
       self.redis_access.add_access_handlers("ETO_VALUES",self.ds_handlers["ETO_VALUES"],"Redis_Hash_Dictionary") 

       
       self.redis_access.add_access_handlers("RAIN_VALUES",self.ds_handlers["RAIN_VALUES"],"Redis_Hash_Dictionary") 


       
       eto_update_table = self.ds_handlers["ETO_ACCUMULATION_TABLE"]
       self.redis_access.add_access_handlers("eto_update_table",eto_update_table,"Redis_Hash_Dictionary") 
  
       
       Load_ETO_Management_Web(self.app, self.auth,request, file_server_library = self.file_server_library,path='eto_py3',url_rule_class=self.url_rule_class,
                  subsystem_name= "ETO_MANAGEMENT",render_template=render_template,redis_access = self.redis_access,eto_update_table = eto_update_table,
                     handlers=self.ds_handlers )    

       

 
   
   


   
