
import os
import json
from datetime import datetime
import time
from  base_stream_processing.base_stream_processing_py3 import Base_Stream_Processing
from  collections import OrderedDict



         
 
  
            


class Load_ETO_Management_Web(Base_Stream_Processing):

   def __init__( self, app, auth, request, file_server_library,path,url_rule_class,subsystem_name,
                   render_template,redis_access,eto_update_table, handlers):
       
       self.app      = app
       self.auth     = auth
       self.request  = request
       self.render_template = render_template
       self.path = path
       
       
       self.url_rule_class = url_rule_class
       self.subsystem_name = subsystem_name
       self.render_template = render_template
       self.eto_update_table = eto_update_table
       self.handlers = handlers
       self.file_server_library = file_server_library
       
     
       self.assemble_url_rules()
       self.path_dest = self.url_rule_class.move_directories(self.path)  
       
      
   def assemble_url_rules(self):
    
       self.slash_name = "/"+self.subsystem_name
       self.menu_data = {}
       self.menu_list = []
       
       

       
       function_list =   [ self.eto_manage,
                           self.eto_readings,
                           self.eto_queue,
                           self.rain_queue,                      
                           self.eto_setup,
                           self.weather_station_problems ]
                           
 
                               
   
      
       url_list = [
                      [ "eto_manage" ,'','',"Manage ETO Values"  ],
                      [ "eto_readings_data" ,'','',"ETO Current Readings"  ],
                      [ 'eto_past' ,'','',"ETO Past Readings"   ],
                      [ 'rain_past' ,'','',"Rain Past Readings"  ],    
                      [ 'eto_setup','','',"ETO Setup"  ],
                      [ 'weather_station_problems','','',"Weather Station Problems"  ]
        ]

       self.url_rule_class.add_get_rules(self.subsystem_name,function_list,url_list) 
       
      
      
      
      
   # 
   #
   #
   #
   #

   def weather_station_problems(self):    
       station_data = self.handlers["EXCEPTION_VALUES"].hgetall()
       return self.render_template(self.path_dest+"/eto_weather_station_issues",station_data = station_data,station_keys = station_data.keys() )
                                
      
       
   def eto_manage( self ):
       return self.render_template(self.path_dest+"/eto_manage",eto_update_table = "eto_update_table")
       
       
   def eto_readings(self):
       eto_data =  self.handlers["ETO_VALUES"].hgetall()
       
       temp_data = {}
       for i,item in eto_data.items():
           temp_data[i] = item["priority"]
       temp_data =[(k, temp_data[k]) for k in sorted(temp_data, key=temp_data.get, reverse=True)]
       eto_keys = []
       for i in temp_data:
          
          eto_keys.append(i[0])
       eto_keys.reverse()
       rain_data =  self.handlers["RAIN_VALUES"].hgetall()
      
       temp_data = {}
       for i,item in rain_data.items():
           temp_data[i] = item["priority"]
       temp_data =[(k, temp_data[k]) for k in sorted(temp_data, key=temp_data.get, reverse=True)]
       rain_keys = []
       for i in temp_data:
          
          rain_keys.append(i[0])
       rain_keys.reverse()
   
       eto_data =  self.handlers["ETO_VALUES"].hgetall()
     
       
       rain_data = self.handlers["RAIN_VALUES"].hgetall()

       return self.render_template( self.path_dest+"/eto_readings",eto_data = eto_data,eto_keys = eto_keys, 
                               rain_data = rain_data,rain_keys =rain_keys ) 

       
   def eto_queue(self):
   
       temp_data = self.handlers["ETO_HISTORY"].revrange("+","-" , count=1000)
       temp_data.reverse()
       #print("temp_data",temp_data)
       chart_title = " ETO Log For Weather Station : "
       stream_keys,stream_range,stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="Daily ETO Rate",title_x="Date")
       '''
       eto_data =  self.handlers["ETO_VALUES"].hgetall()
       temp_data = {}
       for i,item in eto_data.items():
           temp_data[i] = item["priority"]
        
       temp_data =[(k, temp_data[k]) for k in sorted(temp_data, key=temp_data.get, reverse=True)]
       stream_keys = []
       if len(temp_data) > 0:
           for i in temp_data:
              stream_keys.append(i[0])
           stream_keys.reverse()
       if len(stream_keys )== 0:
            stream_range = []
       '''
       return self.render_template( "streams/base_stream",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     max_value = .5,
                                     min_value = 0.,
                                     
                                     )


  
   def rain_queue(self):
       temp_data = self.handlers["RAIN_HISTORY"].revrange("+","-" , count=1000)
       
       temp_data.reverse()
      
       chart_title = " Rain Log For Weather Station : "
       stream_keys,stream_range,stream_data = self.format_data_variable_title(temp_data,title=chart_title,title_y="Daily Rain",title_x="Date")

       '''
       rain_data = self.handlers["RAIN_VALUES"].hgetall()
       
       rain_data =  self.handlers["RAIN_VALUES"].hgetall()
       temp_data = {}
       
       for i,item in rain_data.items():
           temp_data[i] = item["priority"]
       
       temp_data =[(k, temp_data[k]) for k in sorted(temp_data, key=temp_data.get, reverse=True)]
       stream_keys = []
       if len(temp_data) > 0:
           for i in temp_data:
              stream_keys.append(i[0])
       
       stream_keys.reverse()

 
     
       stream_data = temp_data
       if len(stream_keys )== 0:
            stream_range = []

       '''            
       return self.render_template( "streams/base_stream",
                                     stream_data = stream_data,
                                     stream_keys = stream_keys,
                                     title = stream_keys,
                                     stream_range = stream_range,
                                     max_value = 10.,
                                     min_value = 0.,
                                      
                                     
                                     )
      

   def eto_setup(self):
       eto_data  = json.loads(self.file_server_library.load_file( "application_files","eto_site_setup.json" ))
       pin_list  = json.loads(self.file_server_library.load_file("system_files","controller_cable_assignment.json"))
       return self.render_template(self.path_dest+"/eto_setup",eto_data_json = json.dumps(eto_data),pin_list_json=json.dumps(pin_list),
                 eto_data = eto_data       )
 
     
