
class Irrigation_Control( object ):
   
    def __init__(self,ajax_path,file_server, qs,  ):
       self.ajax_path = ajax_path
       self.file_server = file_server
       self.qs = qs
       #
       # generate handlers
       #
       #
       #
       self.assemblers=[]
       self.assemblers["irrigation_queue"] = self.assemble_irrigation_queue()
       
       
       
       
   def assemble_irrigation_queue(self):
       return_value = ""
       return_value = self.load_template("<<",">>",[],"ajax_functions.txt")
       return_value = return_value + self.load_template("<<",">>",[],"queue_by_schedule.txt")
       macro_variables = []
       temp = ["button_title", "Click to Change Queue Job"]
       macro_variables.append(temp)
       temp = ["button_name", "Job_Queue" ]
       macro_variables.append(temp)
       temp = ["function_name","Job_Queue"]
       macro_variables.append(temp)
       temp={"path",self.ajax_path+"/queue_job"]
       return_value = return_value + self.load_template("<<",">>",macro_variables,"mode_control.txt")
       macro_variables = []
       temp = ["init_mode_control", "Click to Change Queue Job"]
       macro_variables.append(temp)
       temp = ["init_queue_by_schedule", "Job_Queue" ]
       initialization_name
<<init_mode_control>>();
   <<init_queue_by_schedule>>();
       return_value = return_value + self.load_template("<<",">>",macro_variables,"js_ready.txt") 
       return return_value
       
   def add_rule( self, rule_path):
     
   def add_ajax_rules(self):
   
   

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
      
   def get_schedule_data(self):
       sprinkler_ctrl = json.loads(self.file_server_library.load_file( "application_files", "sprinkler_ctrl.json" ))   
       returnValue = []
       for j in sprinkler_ctrl:
           temp = json.loads(self.file_server_library.load_file( "application_files", j["link"] )) 
           j["step_number"], j["steps"], j["controller_pins"] = self.generate_steps(temp)
           returnValue.append(j)
     return json.dumps(returnValue)    

     
   def queue_irrigation_jobs(self):
       schedule_data = self.get_schedule_data()
       macro_variables = []
       temp = ["schedule_data_json",schedule_data]
       return  self.load_template("{{","}}",macro_variables,self.assemblers["irrigation_queue"])  
       
   def mode_change( self):
       json_object = self.request.json
       self.handlers["IRRIGATION_JOB_SCHEDULING"].push(json_object)
       return json.dumps("SUCCESS")
   
