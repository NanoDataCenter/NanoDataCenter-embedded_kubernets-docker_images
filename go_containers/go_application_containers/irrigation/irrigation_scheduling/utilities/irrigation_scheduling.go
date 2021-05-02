package scheduling_utilities

/*
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

*/