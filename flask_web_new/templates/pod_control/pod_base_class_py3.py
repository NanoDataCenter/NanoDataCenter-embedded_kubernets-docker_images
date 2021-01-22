from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers 
import json

class Pod_Base_Class(object):

   def __init__( self ,base_self):
       self.base_self = base_self   
       self.site_data = base_self.site_data
       self.qs        = base_self.qs
       self.add_ajax_handler = base_self.add_ajax_handler
    
       
       self.assemble_handlers()






   def assemble_handlers(self):  
   
       #
       #
       # First step is to find controllers
       #
       #
   
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )

       query_list = self.qs.add_match_terminal( query_list, 
                                        relationship = "PROCESSOR" )
                                           
       processor_sets, processor_nodes = self.qs.match_list(query_list)  
       self.processor_names = []
       for i in processor_nodes:
           self.processor_names.append(i["name"])
       self.processor_names.sort()
      
       
       #
       #
       # Assemble data structures for each controller
       #
       #
       self.handlers = []
       for i in self.processor_names:
          self.handlers.append(self.assemble_data_structures(i))

   
       

 
   def assemble_data_structures(self,controller_name ):
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )

       query_list = self.qs.add_match_relationship( query_list, relationship = "PROCESSOR", label = controller_name )
       query_list = self.qs.add_match_relationship( query_list, relationship = "NODE_PROCESSES", label = controller_name )
       query_list = self.qs.add_match_terminal( query_list, 
                                        relationship = "PACKAGE" )
                                           
       package_sets, package_sources = self.qs.match_list(query_list)  
     
       package = package_sources[0] 
       data_structures = package["data_structures"]
       generate_handlers = Generate_Handlers(package,self.qs)
       handlers = {}
       handlers["ERROR_STREAM"]        = generate_handlers.construct_redis_stream_reader(data_structures["ERROR_STREAM"])
       handlers["ERROR_HASH"]        = generate_handlers.construct_hash(data_structures["ERROR_HASH"])
       handlers["WEB_COMMAND_QUEUE"]   = generate_handlers.construct_job_queue_client(data_structures["WEB_COMMAND_QUEUE"])
       
       handlers["WEB_DISPLAY_DICTIONARY"]   =  generate_handlers.construct_hash(data_structures["WEB_DISPLAY_DICTIONARY"])
       return handlers
       
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