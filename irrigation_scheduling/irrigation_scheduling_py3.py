# 
#
# File: utilities.py
#
#
#
#





import msgpack

from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
from file_server_library.file_server_lib_py3 import Construct_RPC_File_Library
from  sqlite_library.sqlite_sql_support_py3 import SQLITE_Client_Support
from system_error_log_py3 import  System_Error_Logging
from Pattern_tools_py3.builders.common_directors_py3 import construct_all_handlers
from Pattern_tools_py3.factories.graph_search_py3 import common_qs_search
from Pattern_tools_py3.factories.get_site_data_py3 import get_site_data

class Monitoring_Base(object):

   def __init__(self,error_logging,file_server_library,file_name,completion_dictionary,job_queue,active_function=None):
       
      
       self.error_logging = error_logging
       self.file_server_library = file_server_library
       
       self.file_name = file_name
       self.completion_dictionary = completion_dictionary
       self.job_queue = job_queue
       self.active_function = active_function


       
   def clear_done_flag( self, *arg ):
      try:
          dow_array = [ 1,2,3,4,5,6,0]
          dow = datetime.datetime.today().weekday()
          dow = dow_array[dow]
          item_control = json.loads(self.file_server_library.load_file("application_files",self.file_name))
          for  j in  item_control:
              name = j["name"]
              if self.determine_start_time( j["start_time"],j["end_time"]) == False: 
                 temp_1 = json.dumps( [0,-1] )
                 temp_check = self.completion_dictionary.hget(name)
                 if temp_1 != temp_check:
                     self.completion_dictionary.hset(name,temp_1)
      except:
          print("bad file")
    



       
   def check_flag( self,item ):
      try:
         data = self.completion_dictionary.hget( item )
 
         #print("data check flag",data)
         data = json.loads( data)

      except:
         #print("exception check_flag")
         data = [ 0 , -3 ]
      
      if int(data[0]) == 0 :
         return_value = True
      else:
         return_value = False
       
      
      return return_value
  

   def match_time( self, compare, value ):
     return_value = False
     if compare[0] < value[0]:
       return_value = True
     if (compare[0] ==  value[0]) and ( compare[1] <= value[1] ):
       return_value = True
     return return_value

   def determine_start_time( self, start_time,end_time ):
       return_value = False
       temp = datetime.datetime.today()
       #print("start_time",start_time,end_time)
       st_array = [ temp.hour, temp.minute ]
       if self.match_time( start_time,end_time ) == True:
	           if ( self.match_time( start_time, st_array) and 
	                self.match_time( st_array, end_time )) == True:
	              return_value = True
       else: 
	         # this is a wrap around case
          if   self.match_time( start_time,st_array) :
               return_value = True
          if  self.match_time(st_array,end_time):
                return_value = True
       #print("return_value",return_value)
       return return_value


   def schedule_doy(self,j,doy):
      divisor = int(j["day_div"])
      modulus = int(j["day_mod"])
      result = doy % divisor
      #print("doy",doy,j["name"],result==modulus,result,divisor,modulus)
      #print(j)
      #print("doy result",doy,result,modulus)
      if result == modulus:
        return True
      else:
        return False

   def schedule_dow(self,j,dow):
       if j["dow"][dow] != 0 :
          #print("dow true")
          return True
       else:
          #print("dow false")
          return False


   def check_for_proper_date(self, j,dow,doy):
       #print(j["name"],dow,doy)
       if "schedule_enable" in j:
          if j["schedule_enable"] == False:  #checking schedule global enable flag
             #print("returning false",j)
             return False
       #print("check for proper date",j['name'],dow,doy)
       #print("checing flag")
       if "day_flag" not in j:
            #print("day flag not present")
            return self.schedule_dow(j,dow)
       elif int(j["day_flag"]) > 0:

         return self.schedule_doy(j,doy)
       else:
 
         return  self.schedule_dow(j,dow)
    
   def check_for_schedule_activity( self, *args):
      #print("made it here")
      if self.active_function != None:
        
         if self.active_function() == False:
          
           return  # something like rain day has occurred
      try:
          temp = datetime.datetime.today()
          dow_array = [ 1,2,3,4,5,6,0]
          dow = datetime.datetime.today().weekday()
          doy = datetime.datetime.today().timetuple().tm_yday
          dow = dow_array[dow]
          st_array = [temp.hour,temp.minute]
          item_control = json.loads(self.file_server_library.load_file("application_files",self.file_name))
          for j in item_control:
         
              name = j["name"]
              #print( "checking schedule",name )
              if self.check_for_proper_date(j,dow,doy) == True:
         
	    
                 start_time = j["start_time"]
                 end_time   = j["end_time"]
                 #print("made it here")
                 if self.determine_start_time( start_time,end_time )  == True:
                     #print( "made it past start time",start_time,end_time )
                     if self.check_flag( name ):
                         #print( "queue in schedule ",name )
                         temp = {}
                         temp["command"] =  "QUEUE_SCHEDULE"
                         temp["schedule_name"]  = name
                         temp["step"]           = 0
                         temp["run_time"]       = 0
                         self.job_queue.push( temp )
                         print("job_queue",temp)
                         temp = [1,time.time()+60*60 ]  # +hour prevents a race condition
                         self.completion_dictionary.hset( name,json.dumps(temp) ) 
      except:
         print("missing files")


   def check_for_system_activity( self, *args):
      try:
         temp = datetime.datetime.today()
         dow_array = [ 1,2,3,4,5,6,0]
         dow = datetime.datetime.today().weekday()
         dow = dow_array[dow]
         doy = datetime.datetime.today().timetuple().tm_yday
         st_array = [temp.hour,temp.minute]
     
         sprinkler_ctrl = json.loads(self.file_server_library.load_file("application_files","system_actions.json"))

         for j in sprinkler_ctrl:
          
               name     = j["name"]
               command  = j["command_string"]
               #print( "checking activity",name)
               if self.check_for_proper_date(j,dow,doy) == True:
            
                  start_time = j["start_time"]
                  end_time   = j["end_time"]
                  #print("system activity",name,start_time,end_time)
                  if self.determine_start_time( start_time,end_time ):
                      #print("start time passed",name)
                      if self.check_flag( name ):
                          #print( "queue in schedule ",name )
                          temp = {}
                          temp["command"]        = command
                          temp["schedule_name"]  = name
                          temp["step"]           = 0
                          temp["run_time"]       = 0
                          self.job_queue.push( temp )
                          print("job queue",temp)
                          temp = [1,time.time()+60*60 ]  # +hour prevents a race condition
                          self.completion_dictionary.hset( name,json.dumps(temp) ) 
      except:
          print("missing file")
          
          

         
class System_Monitoring(Monitoring_Base): 
   
   def __init__(self, error_logging,file_server_library,completion_dictionary,job_queue):
       
       Monitoring_Base.__init__(self,error_logging,file_server_library,"system_actions.json",completion_dictionary,job_queue)
 
   def check_for_required_files( self, *args):
       if self.file_server_library.file_exists("application_files","system_actions.json") == False:
            self.error_logging.log_error_message("NO_FILES",subject="system_actions.json")

       
   '''    


   '''               
  
class Irrigation_Schedule_Monitoring(Monitoring_Base):
   def __init__(self,error_logging, file_server_library,completion_dictionary,job_queue,irrigation_control ):
       Monitoring_Base.__init__(self,error_logging,file_server_library,"sprinkler_ctrl.json",completion_dictionary,job_queue,self.rain_check)
       self.irrigation_control = irrigation_control
       
       

          
   def rain_check(self):
       print("rain_flag",self.irrigation_control.hget("RAIN_FLAG"))
       return not self.irrigation_control.hget("RAIN_FLAG")


   def check_for_required_files( self, *args):
       if self.file_server_library.file_exists("application_files","sprinkler_ctrl.json") == False:
            self.error_logging.log_error_message("NO_FILES",subject="sprinkler_ctrl.json")
   
 



def add_chains(cf,sched,action):


   cf.define_chain( "irrigation_scheduling", True )
   cf.insert.log("irrigation_scheduling")
   cf.insert.one_step( action.check_for_system_activity  )
   cf.insert.one_step( sched.check_for_schedule_activity  )
   cf.insert.wait_event_count( event =  "MINUTE_TICK"  )
   cf.insert.reset()
    
   cf.define_chain("clear_done_flag",True)
   cf.insert.log("clear done chain")
   cf.insert.one_step(action.clear_done_flag )
   cf.insert.one_step(sched.clear_done_flag )
   cf.insert.wait_event_count( event =  "MINUTE_TICK"  )
   cf.insert.reset()
   
   cf.define_chain("check_required_file",True)
   cf.insert.log("check required files")
   cf.insert.one_step(action.check_for_required_files )
   cf.insert.one_step(sched.check_for_required_files )
   cf.insert.wait_event_count( event =  "HOUR_TICK"  )
   cf.insert.reset()

   #
   #
   # internet time update
   #
   #
  

   


if __name__ == "__main__":


    import datetime
    import time
    import string
    import urllib.request
    import math
    import redis
    import base64
    import json

    import os
    import copy
    
    from redis_support_py3.graph_query_support_py3 import  Query_Support
    import datetime
    

    from py_cf_new_py3.chain_flow_py3 import CF_Base_Interpreter
    
    

    site_data = get_site_data()
     
    #
    # Setup handle
    # open data stores instance
    

    qs = Query_Support( site_data )
    
    file_server_library = Construct_RPC_File_Library(qs,site_data)
    
    container_name = os.getenv("CONTAINER_NAME")
    sqlite_client = SQLITE_Client_Support(qs,site_data)
      
    error_logging = System_Error_Logging(qs,container_name,site_data,sqlite_client)   
    
    search_list = ["IRRIGIGATION_SCHEDULING_CONTROL_DATA"]
    handlers = construct_all_handlers(site_data,qs,search_list,field_list = ["IRRIGATION_JOB_SCHEDULING","SYSTEM_COMPLETION_DICTIONARY"])   
    job_queue = handlers["IRRIGATION_JOB_SCHEDULING"]
    completion_dictionary  = handlers["SYSTEM_COMPLETION_DICTIONARY"]

    search_list = ["IRRIGATION_CONTROL_MANAGEMENT"]
    handlers = construct_all_handlers(site_data,qs,search_list,field_list = ["IRRIGATION_CONTROL"])  
    irrigation_control =  handlers["IRRIGATION_CONTROL"] 
    
    sched        = Irrigation_Schedule_Monitoring(error_logging, file_server_library,completion_dictionary,job_queue,irrigation_control )
    action       = System_Monitoring(error_logging,file_server_library,completion_dictionary,job_queue)

 
    cf = CF_Base_Interpreter()
    add_chains(cf,sched,action)
    #
    # Executing chains
    #
    
    cf.execute()

else:
  pass    

