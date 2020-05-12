

class Construct_Monitors(object):

   def __init__(self,site_data,qs,monitoring_list):
      self.site_data = site_data
      self.monitoring_list = monitoring_list
      self.monitors = {}
      for i in self.monitoring_list:
 
         if i == "CORE_OPS":
             print(i)
             self.monitors[i] = self.Generate_Core_Utilites(site_data,qs)
                  
         else:
             raise         
          
   def execute_monitors(self):
       for i in self.monitoring_list:
           self.monitors[i].execute()
