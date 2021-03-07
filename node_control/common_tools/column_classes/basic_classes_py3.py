
class Base_Class(object):

    def  __init__(self,sys_handle,properties):
         self.properties = properties
         self.sys_handle = sys_handle
        
         if "sub_classes" not in properties:
            properties["class_obj"] = []
         
         self.create_classes(self,properties["sub_classes"])
         
         
         if "event_list" in properties:
             pass # find event objects
         else:
            self.event_list = []
         
		 #
		 # Find event objects
		 #
	
		 self.event_queue = []
		 self.active = False
		 self.initialized = False 
		 self.vote_state = False
		 self.vote_value = None
         self.halt_code = "HALT"
         self.continue_code = "CONTINUE"
         self.disable_code = "DISABLE"

    def reset(self):
	    self.active = True
	    for i in self.properties["class_obj"]:
		   i.reset()
		
		
    def initialize(self):
	    for i in self.properties["class_obj"]:
		    i.initialize()
	
	def terminate(self):

	    for i in self.properties["class_obj"]:
		   if (i.active == True )&&(i.initialized== True):
		       i.terminate()



    def check_properties(self,required_fields):	
	    for i in required_fields:
           if i not in self.properties:
              ValueError(f"property { i } required")
              
	def analyize_return_code(self,return_value):
       if return_code == "HALT":
          return False
       if return_code == "DISABLE":
          return True
       if return_code == "CONTINUE":
          return True
       raise ValueError(f"unknown return type { return_code } ")

		
    def create_classes(self,sub_class_definitions):
        pass
        
        
        
	
	

class Leaf(Base_Class):

    def __init__(self,sys_handle,properties):
	    Base_Class.__init__(self,sys_handle,properties)
		self.check_properties("class_obj")
        if "init_function" in self.properties:
           self.init_function = self.properties["init_function"]
        else:
           self.init_function = None
           
        if "term_function" in self.properties:
           self.term_function = self.properties["term_function"]
        else:
           self.term_function = None
           
           
           
           
           
	def reset(self):
	   self.active = True
	
	
	def initialize(self):
	    
		self.initialized = True
		if self.init_function != None:
		   self.init_function(self)
		   
    def terminate(self):
		if self.term_function != None:
		   self.term_function(self)
	    
	    self.active = False
		self.initialized = False

	
    def process_event(self,event):
	    if self.initialized == False:
		   self.initialize()
		   
	    self.event_queue = []
	    self.event_queue.append(event)
		continue_flag = True
		while (len(self.event_queue) != 0) and (continue_flag == True):
		     return_value = self.base_function(self,self.event_queue.pop())
		     continue_flag = self.analyize_return_code(return_value)
        return return_value




class Iterator(Base_Class):

    def __init__(self,sys_handle,propeties):
	    Base_Class.__init__(self,sys_handle,class_obj):
		self.check_properties("class_obj")

	
	
    def process_event(self,event):
        active_link =  False
        for i in self.class_obj:
		   
		   if i.active == True:
		      active_link = True
		      if i.initialized == False:
			     i.initialize()
			  return_code = i.process_event()
              if self.analyize_return_code(return_code) == False:
                 break
        if active_link == False:
           return_value = self.disable_code
		else:
		    return_value = self.halt_code
	    return return_value





class Fork_Join(Base_Class):

    def _init__(self,sys_handle,class_objs,event_list,base_function):
	    Base_Class.__init__(self,sys_handle,class_objs,event_list,base_function)
		self.check_properties("class_obj")		   
 	
    def process_event(self,event):
        active_link =  False
        for i in self.class_obj:
		   
		   if i.active == True:
		      active_link = True
		      if i.initialized == False:
			     i.initialize()
			  i.process_event(event)
              
        if active_link == False:
           return_value = self.disable_code
		else:
		    return_value = self.continue_code
	    return return_value		


  				
class Fork_Join(Base_Class):

    def _init__(self,sys_handle,properties):
	    Base_Class.__init__(self,sys_handle,properties)
        self.check_properties("class_obj")
        

	
    def process_event(self,event):
        active_link =  False
        for i in self.class_obj:
		   
		   if i.active == True:
		      active_link = True
		      if i.initialized == False:
			     i.initialize()
			  i.process_event(event)
              
        if active_link == False:
           return_value = self.disable_code
		else:
		    return_value = self.halt_code
	    return return_value	
        


class Vote(Base_Class):

    def __init__(self,sys_handle,properties):
	    Base_Class.__init__(self,sys_handle,properties)
		self.check_properties("class_obj","vote_count","vote_function")
		
		
	
    def process_event(self,event):
       vote_count = 0
        for i in self.properties["class_obj"]:
		   
		   if i.active == True:
		      
		      if i.initialized == False:
			     i.initialize()
			  i.process_event(event)
			  if i.vote_state == True:
			     vote_count = vote_count + 1
		if vote_count >= self.properties["vote_count"]:
		    return_value = self.properties["vote_function"](self.class_obj)
	    else:
		    return_value self.halt_code
           
	    return return_value		

