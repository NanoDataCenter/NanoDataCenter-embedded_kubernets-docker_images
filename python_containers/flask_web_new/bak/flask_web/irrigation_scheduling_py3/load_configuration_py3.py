import os
import json
import base64

class Load_Configuration_Data(object):

   def __init__( self,app, auth,render_template,request,file_server_library,url_rule_class,subsystem_name,path):
       self.app      = app
       self.auth     = auth
       self.render_template = render_template
       self.request = request
       self.file_server_library = file_server_library
       self.url_rule_class  = url_rule_class
       self.subsystem_name = subsystem_name
       self.path = path
       
   
       
       self.assemble_url_rules()
       self.path_dest = self.url_rule_class.move_directories(self.path)     
 
   def assemble_url_rules(self):
    
       self.slash_name = "/"+self.subsystem_name
       self.menu_data = {}
       self.menu_list = []
       
       

       
       function_list =   [ self.system_actions,
                           self.create_schedule,
                           self.edit_schedules,
                           self.copy_schedule,                      
                           self.delete_schedules]
                           
 
                               
   
      
       url_list = [
                      [ 'system_actions' ,'','',"System Action"  ],
                      [ 'create_schedule' ,'','',"Create Schedule"  ],
                      [ 'edit_schedules' ,'','',"Edit Schedule"   ],
                      [ 'copy_schedule' ,'','',"Copy Schedule"  ],    
                      [ 'delete_schedules','','',"Delete Schedules"  ]
        ]

       self.url_rule_class.add_get_rules(self.subsystem_name,function_list,url_list)
                  

     

       a1 = self.auth.login_required( self.update_schedule )
       self.app.add_url_rule("/ajax/update_schedule",
                          "ajax_update_schedule",a1,methods=["POST"])

                      



   def system_actions(self): 
       system_actions = json.loads(self.file_server_library.load_file( "application_files", "system_actions.json" ))   
       
       
    
       return self.render_template( self.path+"/system_actions",  
                               title="Configure System Events",
                               system_actions       =  system_actions ,
                               system_actions_json  =  json.dumps(system_actions) )

   def create_schedule(self): 
       schedule_data = self.get_schedule_data() 
       
   
       return self.render_template(self.path+"/create_schedule",
                               template_type = "add", 
                               title="Add Schedule",
                               schedule_list      =  schedule_data.keys(),
                               pin_list           =  self.file_server_library.load_file( "system_files", "controller_cable_assignment.json"),
                               schedule_data_json =  json.dumps(schedule_data)  ) 


   def copy_schedule(self):  
       schedule_data = self.get_schedule_data() 
      
       return self.render_template( self.path+"/copy_schedule", 
                               template_type = "copy", 
                               title="Copy Schedule",
                               schedule_list      =  schedule_data.keys(),
                               schedule_data_json =  json.dumps(schedule_data)  ) 


   def delete_schedules(self):  
       schedule_data = self.get_schedule_data()
       return self.render_template( self.path+"/delete_schedule", 
                               template_type = "delete", 
                               title="Delete Schedules",
                               schedule_list      =  schedule_data.keys(),
                               schedule_data_json =  json.dumps(schedule_data)  ) 



   def edit_schedules(self):
       schedule_data = self.get_schedule_data()  
       return self.render_template( self.path+"/edit_schedule", 
                               template_type = "edit",
                               title="Edit Schedule",
                               schedule_list      =  schedule_data.keys(),
                               pin_list           =  self.file_server_library.load_file( "system_files", "controller_cable_assignment.json"),
                               schedule_data_json =  json.dumps(schedule_data)  ) 

 
  
 
  
 
  


   def update_schedule(self ):
      
       return_value     = {}
       param              = self.request.get_json()
       action             = param["action"] 
       schedule           = param["schedule"] 
       schedule_data      = param["data"] 
       print("made it here",action)
       if action == "delete":
           self.delete_schedule( schedule )
           self.delete_link_file( schedule )
           
       else:
           self.save_link_file( schedule, schedule_data[schedule] )
           self.save_schedule( schedule_data[schedule]  )
           
       return json.dumps("SUCCESS")

       

   #
   #  Internal functions
   #
   def get_schedule_data( self, *args):
       sprinkler_ctrl           = json.loads(self.file_server_library.load_file( "application_files", "sprinkler_ctrl.json" ))   
       
     
       returnValue = {}
       for j in sprinkler_ctrl:
           data = json.loads(self.file_server_library.load_file( "application_files", j["link"] )) 
           
           
           j["step_number"], j["steps"], j["controller_pins"] = self.generate_steps(data)         
           returnValue[j["name"]] = j
       return returnValue

   def generate_steps( self, file_data):
       returnValue = []
       controller_pins = []
       if file_data["schedule"] != None:
           schedule = file_data["schedule"]        
           for i  in schedule:
               returnValue.append(i[0][2])
               temp = []
               for l in  i:
                   temp.append(  [ l[0], l[1][0] ] )
               controller_pins.append(temp)
  
  
       return len(returnValue), returnValue, controller_pins



   def find_sched_index( self,  name, ref_sched_data ):
      
        for i in range(0, len(ref_sched_data) ):
           
           if ref_sched_data[i]["name"] == name:
              return i
        return None

  


   def save_schedule( self, schedule_data ):
       name = schedule_data["name"]
       
       ref_sched_data  = json.loads( self.file_server_library.load_file( "application_files", "sprinkler_ctrl.json" ) )

       index = self.find_sched_index( name, ref_sched_data )
       
       if index != None:
            ref_sched_data[ index ] = schedule_data
       else:
            ref_sched_data.append( schedule_data)
        
       ref_sched_data = json.dumps(ref_sched_data)
       self.file_server_library.save_file( "application_files", "sprinkler_ctrl.json" ,ref_sched_data)
      
      

   def save_link_file( self, schedule, schedule_data ):
       
       link_data = {}
       link_data["bits"] = {'1':'C201', '3':'DS2', '2':'C2'}
       link_data["schedule"] = []
       for step in range(0,len(schedule_data["controller_pins"] ) ):
           valve_data = schedule_data["controller_pins"][step]
           time       = schedule_data["steps"][step]
           valve_return = []
           for valve_index in range(0,len(valve_data)):
                 valve_return.append( [ valve_data[valve_index][0], [ valve_data[valve_index][1] ] , time ])
           link_data["schedule"].append( valve_return )
       
       self.file_server_library.save_file( "application_files", schedule+".json" ,json.dumps(link_data))
       
   def delete_link_file( self, schedule ):
       try:
         print("schedule",schedule)
         self.file_server_library.delete_file("application_files",schedule+".json")

       except:
          raise
          
   def delete_schedule( self, schedule ):
       
       ref_sched_data  = json.loads(self.file_server_library.load_file( "application_files", "sprinkler_ctrl.json" ) ) 
       index = self.find_sched_index( schedule, ref_sched_data )
       print("ref_schedu",index,schedule,ref_sched_data)
       if index != None:
            print("saving file")
            del ref_sched_data[ index ] 

            self.file_server_library.save_file( "application_files", "sprinkler_ctrl.json" ,json.dumps(ref_sched_data ))          



