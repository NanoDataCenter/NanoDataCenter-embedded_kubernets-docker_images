

   


class Functional_Monand_Failure(object):
   def __init__(self,function_key,error_handler = None):
       self.function_key = function_key
       if error_handler == None:
          self.error_handler = self.default_error_handler
       else:
          self.error_handler = error_handler       


   
   def bind(self,data):
      self.result = []
      for i in range(0,len(data)):
          try:
            
             temp = data[i]
             
             calc = temp[self.function_key]
             self.result.append(calc(temp))
          except Exception as e:
             
             self.error_handler(e,data[i])
      return self.result
      
   def bind_no_data(self,data):
      self.result = []
      for i in range(0,len(data)):
          try:
            
             temp = data[i]
            
             calc = temp[self.function_key]
             self.result.append(calc())
          except Exception as e:
             raise
             self.error_handler(e,data[i])
      return self.result
      
          
 
       

   def default_error_handler( self,exception_text):
        print("exception-->",exception_text,data[i])    
 
'''
    def make_measurement(self, *parameters):
       
        for source in self.eto_sources:
            print("source",source)
            if "calculator" in source:
                try:
                    source["calculator"].compute_previous_day()
                except Exception as tst:
                   
                    print("exception",source["name"],tst)
                    self.ds_handlers["EXCEPTION_VALUES"].hset(source["name"],str(tst))
        print("calculator done")    
'''        