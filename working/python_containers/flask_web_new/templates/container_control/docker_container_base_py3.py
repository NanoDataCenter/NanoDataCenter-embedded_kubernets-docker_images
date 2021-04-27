
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers 
import json

class Docker_Base_Class(object):

   def __init__( self ,base_self):
       self.base_self = base_self   
       self.site_data = base_self.site_data
       self.qs        = base_self.qs
       self.add_ajax_handler = base_self.add_ajax_handler
    
       self.assemble_containers()
       
       self.handlers= {}
       for i in self.managed_container_names:
           self.handlers[i] = self.assemble_container_data_structures(i)


   def assemble_containers(self):  
   
       #
       #
       # First step is to find controllers
       #
       #
   
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
       query_list = self.qs.add_match_terminal( query_list,relationship="PROCESSOR" )
       processor_sets,processor_nodes = self.qs.match_list(query_list) 
       self.processor_names = []
       self.all_container_names = []
       self.managed_container_names = []
       self.container_control_structure = {}
       for i in processor_nodes:
           self.processor_names.append(i["name"])
           self.container_control_structure[i["name"]] = self.determine_container_structure(i["name"])
           
           services = set(i["services"])
           self.services = list(services)
          
           containers = set(i["containers"])
           self.containers = list(containers)
           containers_list = containers.union(services)
           self.all_container_names.extend(list(containers_list))   
           
          
       self.processor_names.sort()  
       self.all_container_names.sort()
       self.containers.sort()
       self.services.sort()
       #print(self.container_control_structure)
       #print(self.all_container_names)
       self.managed_container_names = []
       for i in processor_nodes:
           self.managed_container_names.extend(self.determine_managed_containers(i["name"]))
       self.managed_container_names.sort()
       
       
       
   def determine_managed_containers(self,processor_name):
       query_list = []
       return_value = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
       query_list = self.qs.add_match_relationship( query_list,relationship="PROCESSOR",label=processor_name )
       query_list = self.qs.add_match_terminal( query_list,relationship="CONTAINER")
       container_sets,container_nodes = self.qs.match_list(query_list)
       for i in container_nodes:
          return_value.append(i["name"])
       return return_value          

   
   def determine_container_structure(self,processor_name):
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
       query_list = self.qs.add_match_relationship( query_list,relationship="PROCESSOR",label=processor_name )
       query_list = self.qs.add_match_relationship( query_list,relationship="DOCKER_MONITOR")
       query_list = self.qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE", label = "DATA_STRUCTURES" )
                                        
       package_sets, package_nodes = self.qs.match_list(query_list)  
   
       #print("package_nodes",package_nodes)
   
       generate_handlers = Generate_Handlers(package_nodes[0],self.qs)      
       
          
         
       package_node = package_nodes[0]
       data_structures = package_node["data_structures"]
       
       #print(data_structures.keys())
       handlers = {}
       handlers["ERROR_STREAM"]        = generate_handlers.construct_redis_stream_reader(data_structures["ERROR_STREAM"])
      
       handlers["WEB_COMMAND_QUEUE"]   = generate_handlers.construct_job_queue_client(data_structures["WEB_COMMAND_QUEUE"])
       handlers["WEB_DISPLAY_DICTIONARY"] = generate_handlers.construct_hash(data_structures["WEB_DISPLAY_DICTIONARY"])
       queue_name = data_structures["DOCKER_UPDATE_QUEUE"]['queue']
       handlers["DOCKER_UPDATE_QUEUE"] = generate_handlers.construct_rpc_client( )
       handlers["DOCKER_UPDATE_QUEUE"].set_rpc_queue(queue_name)
       return handlers

   def assemble_container_data_structures(self,container_name):
       
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
       query_list = self.qs.add_match_relationship( query_list,relationship="CONTAINER",label=container_name)
       query_list = self.qs.add_match_terminal( query_list, 
                                           relationship = "PACKAGE", label = "DATA_STRUCTURES" )

       package_sets, package_nodes = self.qs.match_list(query_list)  
      
 
           
       #print("package_nodes",package_nodes)
   
       generate_handlers = Generate_Handlers(package_nodes[0],self.qs)
       data_structures = package_nodes[0]["data_structures"]
      
       handlers = {}
       handlers["ERROR_STREAM"]        = generate_handlers.construct_redis_stream_reader(data_structures["ERROR_STREAM"])
       handlers["ERROR_HASH"]        = generate_handlers.construct_hash(data_structures["ERROR_HASH"])
       handlers["WEB_COMMAND_QUEUE"]   = generate_handlers.construct_job_queue_client(data_structures["WEB_COMMAND_QUEUE"])
       handlers["WEB_DISPLAY_DICTIONARY"]   =  generate_handlers.construct_hash(data_structures["WEB_DISPLAY_DICTIONARY"])
       handlers["PROCESS_VSZ"]  = generate_handlers.construct_redis_stream_reader(data_structures["PROCESS_VSZ"])
       handlers["PROCESS_RSS"] = generate_handlers.construct_redis_stream_reader(data_structures["PROCESS_RSS"])
       handlers["PROCESS_CPU"]  = generate_handlers.construct_redis_stream_reader(data_structures["PROCESS_CPU"])
   
       return handlers
       
   def generate_container_names(self):
      return_value = []
      
      for i in range(0,len(self.container_names)):
          return_value.append('<option value="'+str(i)+'">'+self.container_names[i]+'</option>')
      return "\n".join(return_value)
 
   def load_container_selection_html(self):
       raw_html = '''
<div class="container">
<center>
<h4>Select Container</h4>
</center>

<div id="select_tag">
<center>
<select id="container_select">
  {{(self.generate_container_names )}}
</select>
</center>
</div>
       '''
       self.mp.generate_container_names = self.generate_container_names
       return self.mp.macro_expand_start("{{","}}",raw_html)
       
       
       
   def load_container_control_javascript(self):
       return '''
<script>
function change_container(event,ui)
{
  
  current_page = window.location.pathname
  
  
  current_page = current_page+"?"+$("#container_select")[0].selectedIndex
 
  window.location.href = current_page
}
</script>
       '''    
       
       
   def generate_processor_names(self):
      return_value = []
      
      for i in range(0,len(self.processor_names)):
          return_value.append('<option value="'+str(i)+'">'+self.processor_names[i]+'</option>')
      return "\n".join(return_value)
       

   def load_processor_selection_html(self):
       raw_html = '''
<div class="container">
<center>
<h4>Select Processor</h4>
</center>

<div id="select_tag">
<center>
<select id="container_select">
  {{(self.generate_processor_names )}}
</select>
</center>
</div>
       '''
       self.mp.generate_processor_names = self.generate_processor_names
       return self.mp.macro_expand_start("{{","}}",raw_html)
       
   def load_processor_control_javascript(self):
       return '''
<script>
function change_processor(event,ui)
{
  current_page = window.location.pathname
 
  
  
  current_page = current_page+"?"+$("#processor_select")[0].selectedIndex
  window.location.href = current_page
}
</script>
       '''
 
          
   def load_processes(self):
       param = self.request.get_json()
      
       container_id = int(param["container"])
       
       if container_id >= len(self.managed_container_names):
          return "BAD"
       else:
          container_name = self.managed_container_names[container_id]
          result = self.docker_performance_data_structures[container_name]["WEB_DISPLAY_DICTIONARY"].hgetall()
          result_json = json.dumps(result)
          
          return result_json.encode()
          

   def manage_processes(self):
       param = self.request.get_json()
      
       container_id = int(param["container"])
       process_state_json = param["process_data"]
       process_state = json.loads(process_state_json)
       if container_id >= len(self.managed_container_names):
          return "BAD"
       else:
          
          container_name = self.managed_container_names[container_id]
          self.docker_performance_data_structures[container_name]["WEB_COMMAND_QUEUE"].push(process_state)
          return json.dumps("SUCCESS")
          
          
   def manage_reset_upgrade(self):
       reset_data = self.request.get_json()
       print("reset_data",reset_data)
       processor_name = self.processor_names[reset_data["processor_id"]]
       handlers = self.container_control_structure[processor_name]
       rpc_client = handlers["DOCKER_UPDATE_QUEUE"]
       try:
          response = rpc_client.send_rpc_message("Upgrade",reset_data,timeout=5 )
          
          return json.dumps("SUCCESS")
       except:
         
          raise