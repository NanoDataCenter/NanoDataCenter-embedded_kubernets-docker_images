class Monad_Class():
   def __init__(self, value):
       self.value = value

   def get(self):
       return self.value

   def __or__(self, f):
        return self.bind(f)



class Failure_Class(Monad_Class):
   def __init__(self,value,failed=False):
       Monad_Class.__init__(self,value)
       self.failed = failed

   def __str__(self):
       return ' '.join([str(self.value), str(self.failed)])
   
   def is_failed(self):
       return self.failed


	   
   def bind(self, f):
         if self.failed:
             return self
         try:
             x = f(self.get())
             
             return Failure_Class(x)
             
         except Exception as e:
             return Failure_Class(e, True)
             
 
class Failure_List_Class(Monad_Class):
   def __init__(self,value,failed=False):
       Monad_Class.__init__(self,value)
       self.failed = failed

   def __str__(self):
       return ' '.join([str(self.value), str(self.failed)])
   
   def is_failed(self):
       return self.failed


	   
   def bind(self, f):
         if self.failed:
             return self
         try:
             
             x = list(map(f, self.value))
             
             return Failure_List_Class(x)
             
         except Exception as e:
             return Failure_List_Class(e, True)
             
                         
             
from operator import neg
           
if __name__ == "__main__":
   Failure = Failure_Class(1)
   print(Failure.bind(neg))
   
   Failure = Failure_Class('t')
   print(Failure.bind(neg))
   
   Failure = Failure_Class(1)
   print(Failure|neg|neg|str)

   Failure_List =  Failure_List_Class([1,2,3,4])
   print(Failure_List|neg|str)