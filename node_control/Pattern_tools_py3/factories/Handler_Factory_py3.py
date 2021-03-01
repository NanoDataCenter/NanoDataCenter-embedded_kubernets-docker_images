
from redis_support_py3.construct_data_handlers_py3 import Generate_Handlers

class Handler_Factory(Generate_Handlers):
      
   def __init__(self,package,qs ):
       Generate_Handlers.__init__(self,package,qs)
       self.package = package

 
                                         
  
   def generate_handlers(self, field_list=None,type_list=None):
       if (field_list == None) and (type_list==None):
          return self.generate_all()
       if field_list != None:
          return self.generate_specific_key(field_list)
       if search_list != None:
          return self.generate_specific_type(type_list)
       raise  ValueError("bad logic")          

   def generate_all(self):

       handlers = {}
       data_structures = self.package["data_structures"]
       for key,data in data_structures.items():
           
           if data["type"] in self.constructors:
               
               handlers[key] = self.constructors[data["type"]](data)
           else:
               raise ValueError("Bad handler key")
               
       return handlers    


   def generate_specific_key(self,key_list):
       handlers = {}
       key_set = set(key_list)
       data_structures = self.package["data_structures"]
       for key,data in data_structures.items():
           if key in key_set:
               if data["type"] in self.constructors:
                   handlers[key] = self.constructors[data["type"]](data)
               else:
                   raise ValueError("Bad handler key")
               
       return handlers  

   def generate_specific_type(self,type_list):
       handlers = {}
       type_set = set(type_list)
       data_structures = self.package["data_structures"]
       for key,data in data_structures.items():
           if data["type"] in type_set:
               if data["type"] in self.constructors:
                   handlers[key] = self.constructors[data["type"]](data)
               else:
                   raise ValueError("Bad handler key")
               
       return handlers         