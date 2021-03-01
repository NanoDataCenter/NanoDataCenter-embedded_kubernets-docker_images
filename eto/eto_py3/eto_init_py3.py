






import json


       
class Initialize_ETO_Accumulation_Table(object):

   def __init__(self, file_system_library ):
   
       self.file_system_library = file_system_library
       
       
      
       
 
  
   
       
       
   def initialize_eto_tables(self,eto_data_handler):
       
       self.eto_data_handler = eto_data_handler
       # the eto_site table may have changed
       # need to merge old table values into the new table
       # there may be insertions as well as deletions
       try:
           response = self.file_system_library.load_file("application_files","eto_site_setup.json")
           print("response",response)
           eto_file_data = json.loads(response)
       except:
          raise
          print("eto_site_setup.json does not exit")
          return


      
       
       eto_redis_hash_data = self.eto_data_handler.hgetall()

       
       new_data = {}
       
       #
       # Step 1  Populate file dummy initial values
       #
       
       for j in eto_file_data:
          
           new_data[ j["controller"] + "|" + str(j["pin"])] = 0

       
       
       self.eto_data_handler.delete_all()
       
       #
       # merge old values and possible new values into new table.
       #
       
       for i in new_data.keys():
           
           
           if i in eto_redis_hash_data: #checking to see if entry in old table
             
              data = eto_redis_hash_data[i]   # key old values
           else:
              data = 0
              
           self.eto_data_handler.hset(i,data )         
           
        
