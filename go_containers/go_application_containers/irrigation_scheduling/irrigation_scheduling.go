package main

import (
    
    "fmt"
    
	"time"
 
    "lacima.com/site_data"
    "lacima.com/redis_support/graph_query"
    "lacima.com/redis_support/redis_handlers"
    "lacima.com/redis_support/generate_handlers"
    "lacima.com/server_libraries/file_server_library"
	"lacima.com/cf_control"
)

type Irrigation_Scheduling_Type struct {

    irrigation_hash    redis_handlers.Redis_Hash_Struct
    completion_file       redis_handlers.Redis_Hash_Struct
    base_file             string

}

type System_Scheduling_Type struct {

    completion_file       redis_handlers.Redis_Hash_Struct
    base_file             string

}


type Scheduling_Type struct {

  job_queue           redis_handlers.Redis_Job_Queue
  system_control      System_Scheduling_Type
  irrigation_control  Irrigation_Scheduling_Type
      

}

var  CF_site_node_control_cluster cf.CF_CLUSTER_TYPE
var fs  file_server_lib.File_Server_Client_Type

func main() {
  
  var return_value Scheduling_Type
  var config_file = "/data/redis_server.json"
  var site_data_store map[string]interface{}

  site_data_store = get_site_data.Get_site_data(config_file)
  graph_query.Graph_support_init(&site_data_store)
  redis_handlers.Init_Redis_Mutex()
  data_handler.Data_handler_init(&site_data_store)
  fs = file_server_lib.File_Server_Init(&[]string{"FILE_SERVER"})
  fmt.Println((fs).Ping())	
  var fs_handle = file_server_lib.File_Server_Init(&[]string{"FILE_SERVER"})
    fmt.Println((fs_handle).Ping())	
 (CF_site_node_control_cluster).Cf_cluster_init()
 (CF_site_node_control_cluster).Cf_set_current_row("irrigation_scheduling")
  
  (&return_value).irrigation_initialize_setup()
  (&return_value).irrigation_schedule_exec()


}


func ( v* Scheduling_Type ) irrigation_initialize_setup(){

    search_path := []string{"IRRIGIGATION_SCHEDULING:IRRIGIGATION_SCHEDULING","IRRIGIGATION_SCHEDULING"}
    handlers := data_handler.Construct_Data_Structures(&search_path)
    
	v.irrigation_control.completion_file = (*handlers)["IRRIGATION_COMPLETION_DICTIONARY"].(redis_handlers.Redis_Hash_Struct)
	v.system_control.completion_file     = (*handlers)["SYSTEM_COMPLETION_DICTIONARY"].(redis_handlers.Redis_Hash_Struct)
	
    search_path = []string{"IRRIGIGATION_CONTROL:IRRIGIGATION_CONTROL","IRRIGIGATION_CONTROL"}
    handlers = data_handler.Construct_Data_Structures(&search_path)
    
	v.irrigation_control.irrigation_hash  = (*handlers)["IRRIGATION_CONTROL"].(redis_handlers.Redis_Hash_Struct)
	v.job_queue                   = (*handlers)["IRRIGATION_JOB_SCHEDULING"].(redis_handlers.Redis_Job_Queue)

    v.irrigation_control.base_file = "sprinkler_ctrl.json"
	v.system_control.base_file     = "system_actions.json"



}


func ( v* Scheduling_Type )irrigation_schedule_exec(){

  panic("done")
  //(CF_site_node_control_cluster).CF_Fork()


}



func (v* Scheduling_Type)construct_chain(){
  
   var cf_control  cf.CF_SYSTEM_TYPE
  (cf_control).Init(&CF_site_node_control_cluster , "irrigation_scheduling" ,true, time.Minute )




  (cf_control).Add_Chain("irrigation_scheduling",true)
  (cf_control).Cf_add_log_link("irrigation_scheduling")
   
   
  (cf_control).Cf_add_one_step(v.action_check_for_system_activity, make(map[string]interface{}))  
  (cf_control).Cf_add_one_step(v.sched_check_for_schedule_activity, make(map[string]interface{})) 
  
  (cf_control).Cf_add_wait_interval(time.Minute )
  (cf_control).Cf_add_reset()
  
  (cf_control).Add_Chain("clear_done_flag",true)
  (cf_control).Cf_add_log_link("clear_done_flag")

  (cf_control).Cf_add_one_step(v.action_clear_done_flag, make(map[string]interface{})) 
  (cf_control).Cf_add_one_step(v.sched_clear_done_flag, make(map[string]interface{})) 
  
  (cf_control).Cf_add_wait_interval(time.Minute )
  (cf_control).Cf_add_reset()
    
  (cf_control).Add_Chain("check_required_file",true)
  (cf_control).Cf_add_log_link("check_required_file")
  
  (cf_control).Cf_add_one_step(v.action_check_for_required_files , make(map[string]interface{})) 
  (cf_control).Cf_add_one_step(v.sched_check_for_required_files , make(map[string]interface{}))
  
  (cf_control).Cf_add_wait_interval(time.Minute )
  (cf_control).Cf_add_reset()	
	
	

}



func (v *Scheduling_Type)action_check_for_system_activity( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{




   return cf.CF_DISABLE
}
 

func (v *Scheduling_Type) sched_check_for_schedule_activity( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{



    return cf.CF_DISABLE
}  
  
 
func (v *Scheduling_Type)action_clear_done_flag( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{



    return cf.CF_DISABLE
}  
 
   
func (v *Scheduling_Type) sched_clear_done_flag( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{



   return cf.CF_DISABLE
}
  
 
func (v *Scheduling_Type) action_check_for_required_files( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{



    return cf.CF_DISABLE
} 
 
 
func (v *Scheduling_Type) sched_check_for_required_files( system interface{},chain interface{}, parameters map[string]interface{}, event *cf.CF_EVENT_TYPE)int{



   return cf.CF_DISABLE
}
  

	






/*
# 
#
# File: utilities.py
#
#
#
#





import msgpack

from common_tools.redis_support_py3.construct_data_handlers_py3 import Generate_Handlers
from common_tools.file_server_library.file_server_lib_py3 import Construct_RPC_File_Library

from common_tools.system_error_log_py3 import  System_Error_Logging
from common_tools.Pattern_tools_py3.builders.common_directors_py3 import construct_all_handlers
from common_tools.Pattern_tools_py3.factories.graph_search_py3 import common_qs_search
from common_tools.Pattern_tools_py3.factories.get_site_data_py3 import get_site_data

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
*/