import os
import json
from datetime import datetime
import time
import datetime
 

class Load_Ticket_Control(object):

   def __init__( self, app, auth, request, render_template,qs,site_data,url_rule_class,sqlite_client,common_qs_search,subsystem,path):
       self.app      = app
       self.auth     = auth
       self.request  = request
       self.render_template = render_template
       self.path = path
       self.subsystem = subsystem
       self.qs = qs
       self.site_data = site_data 
       self.url_rule_class = url_rule_class
       self.sqlite_client = sqlite_client
       self.common_qs_search = common_qs_search
       self.parse_graph_db()
       self.construct_data_bases()
       self.construct_tables()
       
       self.assemble_url_rules()
       self.path_dest = self.url_rule_class.move_directories(self.path)
       

   def parse_graph_db(self):
       
       self.db_nodes = self.common_qs_search(["TICKET_CONTROL","DATA_BASE"])
       self.table_nodes = self.common_qs_search(["TICKET_CONTROL","TABLE"])
       types_nodes = self.common_qs_search(["TICKET_CONTROL","VALID_TYPES"])
       #print(self.db_nodes)
       #print(self.table_nodes)
       #print(types_nodes)
       self.types = types_nodes[0]["types"]
       #print(self.types)
       
     
       self.log_nodes = self.common_qs_search(["TICKET_LOG","TABLE"])
       
       
   def construct_data_bases(self):
       self.target_db = self.db_nodes[0]["db"]
       #print(self.target_db)
       if self.target_db not in self.sqlite_client.list_data_bases():
         self.sqlite_client.create_database(self.target_db)

     
       
   def construct_tables(self):
       self.ticket_table = self.table_nodes[0]["name"]
       self.ticket_fields = self.table_nodes[0]["fields"]
       #print(self.ticket_table)
       #print(self.ticket_fields)
       #self.sqlite_client.drop_table(self.target_db,self.ticket_table)
       try:  ## cannot detect virtual table -- therefor trap exception
           self.sqlite_client.create_text_search_table(self.target_db,self.ticket_table, self.ticket_fields )
       except:
           pass
       self.ticket_table_log = self.log_nodes[0]["name"]
       self.ticket_log_fields = self.log_nodes[0]["fields"]

       #self.sqlite_client.drop_table(self.target_db,self.ticket_table_log)
       try:  ## cannot detect virtual table -- therefor trap exception
           self.sqlite_client.create_text_search_table(self.target_db,self.ticket_table_log, self.ticket_log_fields )
       except:
           pass      
  
   
   def assemble_url_rules(self):
       
       
       
       self.slash_name = "/"+self.subsystem+"/"
       self.menu_data = {}
       self.menu_list = []
       
       function_list = [ self.manage_tickets,
                         self.manage_logs ]
                         
                         
       url_list = [ [ 'manage_tickets','','',"Manage Tickets"  ],
                    [ 'manage_logs','','',"Manage Logs"  ]]
                    
       
                        

      
       self.url_rule_class.add_get_rules(self.subsystem,function_list,url_list)
      
       # internal callable
       a1 = self.auth.login_required( self.add_entry )
       self.app.add_url_rule(self.slash_name+"add_link",self.slash_name+"add_link",a1,methods=["POST"])

       # internal callable
       a1 = self.auth.login_required( self.delete_entry )
       self.app.add_url_rule(self.slash_name+"delete_link",self.slash_name+"delete_link",a1,methods=["POST"])

       # internal callable
       a1 = self.auth.login_required( self.modify_entry )
       self.app.add_url_rule(self.slash_name+"modify_link",self.slash_name+"modify_link",a1,methods=["POST"])       
    

       # internal callable
       a1 = self.auth.login_required( self.delete_log )
       self.app.add_url_rule(self.slash_name+"delete_log",self.slash_name+"delete_log",a1,methods=["POST"])      

 
   #
   #
   #
   #  Web page handlers
   #
   #+ datetime.datetime.fromtimestamp(i["create_timestamp"]).isoformat()
   #
   # "fields":[ "id INTEGER PRIMARY KEY  AUTOINCREMENT","active Int","create_timestamp FLOAT","close_timestamp FLOAT","type Int","subtype Text",""subtype TEXT","description Text","resolution TEXT"   ]} )
   def manage_tickets(self):
      
           Display_Title = "Active Tickets" 
           #print(self.ticket_fields)
           temp = self.ticket_fields
           temp.append("rowid")
           
           table_data = self.sqlite_client.select_composite(self.target_db,self.ticket_table,temp,where_clause=None,distinct_flag=False)
           #print("table_data",table_data)
          
           for i in table_data:
               i["summary_display"] = self.build_summary_display(i)
           full_link = self.slash_name+'manage_tickets'
           search_link = self.slash_name+'search_link'
           add_link = self.slash_name+"add_link"
           delete_link = self.slash_name+"delete_link"
           modify_link = self.slash_name+"modify_link"
           return self.render_template(self.path_dest+"/ticket_control",
                                       table_data=table_data,
                                       Display_Title=Display_Title,
                                       full_link=full_link,
                                       search_link=search_link,
                                       modify_link=modify_link,
                                       add_link=add_link,
                                       delete_link = delete_link)


   def manage_logs(self):
      
           Display_Title = "Manage_Logs" 
           temp = self.ticket_log_fields
           temp.append("rowid")

           table_data = self.sqlite_client.select_composite(self.target_db,self.ticket_table_log,temp,where_clause=None,distinct_flag=False)
           #print("table_data",table_data)
          
           for i in table_data:
               summary_data = "id: "+str(i["rowid"])+" type/subtype: "+i["type"]+"/"+i["subtype"]+" title: "+i["title"]
               summary_data = summary_data +" date closed: "+ i["close_timestamp"]
               i["summary_display"] = summary_data
           full_link = self.slash_name+'manage_tickets'
           search_link = self.slash_name+'search_log'
           
           delete_link = self.slash_name+"delete_log"
           export_link = self.slash_name+"export_link"
           
           return self.render_template(self.path_dest+"/log_control",
                                       table_data=table_data,
                                       Display_Title=Display_Title,
                                       full_link=full_link,
                                       search_link=search_link,
                                       export_link=export_link,
                                       delete_link = delete_link)




                                       
   def build_summary_display(self,item):
       #print("item",item)
       item["rowid"]  = str(item["rowid"])
       if type(item["close_timestamp"]) == int:
           item["close_timestamp"] = datetime.datetime.fromtimestamp(item["close_timestamp"]).isoformat()      
       item["active"] = self.filter_active(item["active"])
       item["type"] = self.filter_type(item["type"])
       item["subtype"] = str(item["subtype"])
       item["title"] = str(item["title"])
       item["create_timestamp"] = datetime.datetime.fromtimestamp(item["create_timestamp"]).isoformat()
       #print("item",item)
       return "ID: "+item["rowid"]+" "+item["active"] +" Type: "+item["type"]+" Subtype: "+item["subtype"] +"  Title: "+item["title"] +" Creation Time:  "+item["create_timestamp"]
       
       
   def filter_active(self,value):
       if value > 0:
          return "active"
       else:
          return "resolved"

   def filter_type(self,value):
       #print(value)
       value = int(value)
       types = ["OTHERS","IRRIGATION_ISSUES","IRRIGATION_EQUIPMENT","TRIMMING","NON_IRRIGATION_FIXING"]
       if len(types) >= value:
          #print("made it here")
          return types[value]
       else:
          return types[0]
       




   

       
   ## internal call
   def add_entry(self):
      values = self.request.get_json()
      values["create_timestamp"]=time.time()
      values["close_timestamp"] = 0
      #print("values",values)
      self.sqlite_client.insert_composite(self.target_db,self.ticket_table,["active","create_timestamp","close_timestamp","type","subtype","subtype","resolution","title","description"], values)
     
      return json.dumps("Success")
      
   def delete_entry(self):
       values = self.request.get_json()
       self.sqlite_client.delete(self.target_db,self.ticket_table,where_clause="rowid = "+str(values["rowid"]))
       return json.dumps("SUCCESS")
   
   def delete_log(self):
       values = self.request.get_json()
       self.sqlite_client.delete(self.target_db,self.ticket_table_log,where_clause="rowid = "+str(values["rowid"]))
       return json.dumps("SUCCESS")       
       
   def modify_entry(self):
      values = self.request.get_json()
      
      if int(values["active"]) == 0:
         
         values["close_timestamp"]=time.time()
         self.add_log_table(values)
      self.sqlite_client.update(self.target_db,self.ticket_table,["active","close_timestamp","resolution"],[values["active"],values["close_timestamp"],values["resolution"]],where_clause="rowid = "+str(values["rowid"]))
      return json.dumps("SUCCESS")
      
   def add_log_table(self,values):
       values["entry_timestamp"] = time.time()
       values["close_timestamp"] = datetime.datetime.fromtimestamp(values["close_timestamp"]).isoformat() 
       self.sqlite_client.insert_composite(self.target_db,self.ticket_table_log,["entry_timestamp","create_timestamp","close_timestamp","type","subtype","subtype","resolution","title","description"], values)