from .docker_container_base_py3 import Docker_Base_Class
from templates.Base_Multi_Template_Class_py3  import Base_Multi_Template_Class
from flask import request
import json

class Start_Stop_Containers(Base_Multi_Template_Class,Docker_Base_Class):
   def __init__(self,base_self,parameters = None):
       Docker_Base_Class.__init__(self,base_self)
       Base_Multi_Template_Class.__init__(self,base_self,parameters)

   def change_processor_state(self):
       param = request.get_json()
      
       processor_id = int(param["processor_id"])
       process_state_json = param["process_data"]
       process_state = json.loads(process_state_json)
      
       if processor_id >= len(self.processor_names):
          
          return "BAD"
       else:
          processor_name = self.processor_names[processor_id]
         
          self.container_control_structure[processor_name]["WEB_COMMAND_QUEUE"].push(process_state)
          return json.dumps("SUCCESS")

   def load_containers(self):
       param = request.get_json()
      
       processor_id = int(param["processor_id"])
       
       if processor_id  >= len(self.processor_names):
          return "BAD"
       else:
          processor_name = self.processor_names[processor_id]
          result = self.container_control_structure[processor_name]["WEB_DISPLAY_DICTIONARY"].hgetall()

          result_json = json.dumps(result)
          
          return result_json.encode()
   

   def application_page_contruction(self):
       add_ajax_handler = self.base_self.add_ajax_handler
       self.ajax_names={}

       self.ajax_names["change_processor_state"] = "/ajax/manage_containers/change_processor_state"
       add_ajax_handler(self.ajax_names["change_processor_state"],self.change_processor_state,methods=["POST"])

       self.ajax_names["load_containers"] = "/ajax/manage_containers/load_containers, change_processor_state"
       add_ajax_handler(self.ajax_names["load_containers"],self.load_containers,methods=["POST"])


   def application_page_generation(self,processor_id,data):
       self.processor_id = processor_id
       self.processor_name = self.processor_names[processor_id]
       self.display_list = self.container_control_structure[self.processor_name]["WEB_DISPLAY_DICTIONARY"].hkeys()
       return self.generate_template()
     
   def generate_template(self):
       return_value = []
       return_value.append(self.process_html())
       return_value.append(self.process_load_javascript())
       return "\n".join(return_value)
 
 
 
   def process_html(self):
       return_value = []
       return_value.append(self.load_processor_selection_html())
       return_value.append(self.mp.macro_expand_start("{{","}}",self.load_html()))
       return "\n".join(return_value)
 
   def process_control(self,container_id):
      container_name = self.managed_container_names[container_id]
      display_list = self.docker_performance_data_structures[container_name]["WEB_DISPLAY_DICTIONARY"].hkeys()
      
      return self.render_template(self.path_dest+"/docker_process_control",
                                  display_list = display_list, 
                                  command_queue_key = "WEB_COMMAND_QUEUE",
                                  process_data_key = "WEB_DISPLAY_DICTIONARY",
                                  container_id = container_id,
                                  containers = self.managed_container_names,
                                  load_process = '"'+self.slash_name+'/load_processes/load_process'+'"',
                                  manage_process =  '"'+self.slash_name+'/manage_processes/change_process'+'"' )