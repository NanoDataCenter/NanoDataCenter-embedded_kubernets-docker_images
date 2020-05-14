
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers



class  Common_Functions(object):

   def __init__(self,site_data,qs,handlers):
       self.site_data = site_data
       self.qs = qs
       self.handlers = handlers
       

  
       
   def find_processors(self):
       query_list = []
       
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
       query_list = self.qs.add_match_terminal( query_list,relationship = "PROCESSOR" )
       processor_sets, processor_lists = self.qs.match_list(query_list)  
       return_value = []
       for i in processor_lists:
          return_value.append(i["name"])
       return return_value
   
   def find_containers(self,processor_name):
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
       query_list = self.qs.add_match_terminal( query_list,relationship = "PROCESSOR",label = processor_name )
       processor_sets, processor_lists = self.qs.match_list(query_list)  
       return_value = []
       processor_data = processor_list[0]
       return processor_data[ properties["containers"]]

   def generate_hash_structures_with_processor_container(self,processor_listdata_structures,hash_flag = True)
       processor_ds = {}
       
       for i in processor_list:
           containers = self.find_containers(i)
           for j in containers:
               container_ds = {}           
               print(i,j)
               query_list = []
               query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
               query_list = self.qs.add_match_relationship( query_list,relationship = "PROCESSOR",label = i )
               query_list = self.qs.add_match_relationship( query_list,relationship = "CONTAINER",label = j )
               query_list = self.qs.add_match_terminal( query_list,relationship = "PACKAGE" )
               package_sets, package_sources = self.qs.match_list(query_list)  
               package = package_sources[0]     
               data_structures = package["data_structures"]
               print("data_structures",data_structures.keys())
               generate_handlers = Generate_Handlers( package, self.qs )
               temp = {}
               for k in key_list:
                   if hash_flag == True:
                      temp[k] = generate_handlers.construct_hash(data_structures[k] )
                   else:
                      temp[k] = generate_handlers.construct_redis_stream_writer(data_structures[k] )
               container_ds[j] = temp
           processor_ds[i] = container_ds
       return processor_ds


   
   def generate_hash_structures_with_processor(self,processor_list,search_list,key_list,hash_flag = True):
       return_value = {}
      
       for i in processor_list:
           print(i)
           query_list = []
           query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
 
           query_list = self.qs.add_match_relationship( query_list,relationship = "PROCESSOR",label = i )
           for j in range(0,len(search_list)-1):
               query_list = self.qs.add_match_relationship( query_list,relationship = search_list[j] )
           query_list = self.qs.add_match_terminal( query_list,relationship = search_list[-1] )
           package_sets, package_sources = self.qs.match_list(query_list)  
           package = package_sources[0]     
           data_structures = package["data_structures"]
           print("data_structures",data_structures.keys())
           generate_handlers = Generate_Handlers( package, self.qs )
           temp = {}
            for j in key_list:
               if hash_flag == True:
                   temp[k] = generate_handlers.construct_hash(data_structures[k] )
               else:
                   temp[k] = generate_handlers.construct_redis_stream_writer(data_structures[k] )
           return_value[i] = temp
           
       return return_value
           
           
   def generate_hash_structures_without_processor(self,search_terms,key_list,hash_flag = True):
       query_list = []
       query_list = self.qs.add_match_relationship( query_list,relationship="SITE",label=self.site_data["site"] )
       for i in range(0,len(search_list)-1):
           query_list = self.qs.add_match_relationship( query_list,relationship = search_list[i] )
       query_list = self.qs.add_match_terminal( query_list,relationship = search_list[-1] )
       package_sets, package_sources = self.qs.match_list(query_list)  
       package = package_sources[0]     
       data_structures = package["data_structures"]
       print("data_structures",data_structures.keys())
       generate_handlers = Generate_Handlers( package, self.qs )
       return_value = {}
       for j in key_list:
           if hash_flag == True:
               return_value[j] = generate_handlers.construct_hash(data_structures[j] )
           else:
               return_value[j] = generate_handlers.construct_redis_stream_writer(data_structures[k] )
       return return_value

 
       
   def general_hash_iterator(self,subsystem,check_function,hash_obj):

       error_map = {}
       status_map = {}
       error_count = 0           
       for i in hash_obj.hkeys():
           #print(i)
           store_flag, status,result =  check_function(hash_obj.hget(i))
           if store_flag == False:
              continue
           status_map[i] = [status,result]
           if status == False:
               error_count = error_count +1
               error_map[i] = [status,result]
           
       if error_count >0:
          print("failure",i)
          self.handlers["SYSTEM_STATUS"].hset(subsystem,False)
     
       return [error_count, status_map, error_map]


   def check_for_error_flag(self,sub_system,new_data):
       print("new_data",new_data)
       if new_data[0] > 0: # error count != 0
          self.handlers["SYSTEM_STATUS"].hset(subsystem,False)
          return False
       return True
          


   ###### this needs work
   def  detect_new_alert(self,subsystem,new_data,ref_data):
       try:
           if len(new_data) != len(ref_data):
              return True # wait next time for match
           for count in range(0,len(new_data)):
               test_data = ref_data[count]
               test_element = new_data[count]
               if test_data[0] != test_element[0]: # error count not the same
                         self.handlers["SYSTEM_STATUS"].hset(subsystem,False)
                         return False
               for keys in test_data[1]:
                   if keys not in test_element[1]: # process mismatch wait next iteration
                      return True
                   for key in test_data[1]:
                      if test_data[1][key][0] != test_element[1][key][0]:
                         self.handlers["SYSTEM_STATUS"].hset(subsystem,False)
                         return False
           return True
        
       except:
          print("exception") #probably empty data set
          return True          
       

   def log_alert(self,subsystem,data):
       self.handlers["SYSTEM_ALERTS"].push(data=redis_data)  
      
 
 