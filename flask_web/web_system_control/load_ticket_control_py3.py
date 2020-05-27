import os
import json
from datetime import datetime
import time
import datetime
 

class Load_Ticket_Control(object):

   def __init__( self, app, auth, request, render_template,qs,site_data,url_rule_class,path,sqlite_library):
       self.app      = app
       self.auth     = auth
       self.request  = request
       self.render_template = render_template
       self.path = path
       self.qs = qs
       self.site_data = site_data
       self.url_rule_class = url_rule_class
       self.sqlite_library = sqlite_library
       self.parse_graph_db()
       self.construct_data_bases()
       self.construct_tables()
       
       #self.assemble_url_rules()
       #self.path_dest = self.url_rule_class.move_directories(self.path)
       
  bc.add_header_node("TICKET_CONTROL")
   bc.add_info_node( "DATA_BASE","TICKET_CONTROL", properties = {"db":"SYSTEM_CONTROL.db"} )
   bc.add_info_node("TABLE","TICKET_CONTROL",properties = {"name":"TICKET_CONTROL.db",
                    "fields":[ "id INTEGER AUTOINCREMENT","active Int","create_timestamp FLOAT","close_timestamp FLOAT","type Int","subtype Text","description TEXT","resolution TEXT"   ]} )
   bc.add_info_node("VALID_TYPES","TICKET_CONTROL",properties = {"types":["OTHERS","IRRIGATION_ISSUES","IRRIGATION_EQUIPMENT","TRIMMING"]})                   
   bc.end_header_node("TICKET_CONTROL")
   def parse_graph_db(self):
       pass
       #find db
       #find table
       #find find type values
       
   def construct_data_bases(self):
       pass
       
       
   def construct_tables(self):
       pass
   
       
  
   '''
   def assemble_url_rules(self):
       
       
       
       self.slash_name = "/"+self.subsystem_name
       self.menu_data = {}
       self.menu_list = []
       
       function_list = [ self.process_control ,
                         self.display_exception_status,
                         self.display_exception_log ]
                         
       url_list = [ [ 'start_and_stop_processes','/<int:controller_id>','/0',"Stop/Start Pod Processes"  ],
                    [ 'display_exception_status','/<int:controller_id>','/0',"Pod Processes Status"  ],
                     [ 'display_exception_log','/<int:controller_id>','/0',"Pod Process Exception Log"  ] ]                            

      
       self.url_rule_class.add_get_rules(self.subsystem_name,function_list,url_list)
      

       
       
       # internal callable
       a1 = self.auth.login_required( self.load_processes )
       self.app.add_url_rule(self.slash_name+'/manage_processes/load_process',self.slash_name+"node_process_load_process",a1,methods=["POST"])
       
       # internal call
       a1 = self.auth.login_required( self.manage_processes )
       self.app.add_url_rule(self.slash_name+'/manage_processes/change_process',self.slash_name+"node_process_change_process",a1,methods=["POST"])
    
    
   '''
   #
   #
   #
   #  Web page handlers
   #
   #
   #

