
from macro_processor_py3 import Macro_Expansion

def assemble_irrigation_queue():
   mp = Macro_Expansion()
   return_value = ""
   return_value = mp.macro_expand_file("<<",">>","macros/ajax_functions.txt")
   return_value = return_value + mp.macro_expand_file("<<",">>","macros/queue_by_schedule.txt")
 
   mp.button_title =  "Click to Change Queue Job"
   mp.button_name  =  "Job_Queue"
   mp.function_name = "job_queue"
   mp.path = "/ajax/job_queue"
   mp.load_data ="load_queue_object()"
   mp.initialization_name = "a_bind_data"
   return_value = return_value + mp.macro_expand_file("<<",">>", "macros/mode_control.txt")
   
   '''
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
   '''
   return return_value

if __name__ == "__main__":
   print(assemble_irrigation_queue())
