
import hy
class Macro_Processor(object):
   def __init__(self):
       pass

   def lisp_eval(self,expression):
       x = hy.read_str(expression)
      
       return str(hy.eval(x))

   def macro_expand_file(self,start,end,file_name):
       file_handle = open(file_name,"r")
       data = file_handle.read()
       return self.macro_expand_start(start,end,data)
       
       
   def macro_expand_start(self,start,end,data):
       
       self.start = start
       self.end   = end
       
       return self.macro_expand(data)
       

   def macro_expand(self,raw_string):
       return_value = []
       start_list = raw_string.split(self.start)
      
       return_value.append(start_list.pop(0))
       
       for i in start_list:
          
           temp = i.split(self.end)
           
           return_value.append(self.lisp_eval(temp[0]))
           if len(temp) > 1 :
             return_value.append(temp[1])
       
       return "".join(return_value)
   
   def include(self,file_name):
       file = open(file_name)
       data = file.read()
       file.close()
       return data

   def include_mp_expand(self,file_name):
       file = open(file_name)
       data = file.read()
       file.close()
       return self.macro_expand(data)

if __name__ == "__main__":
    import hy
    import macro_functions
    macro_functions.test = test
   
    macro_functions.top_block = "test"
    test = 'abcdef\n<<( macro_functions.block "echo_block")>>\nabcde<<macro_functions.top_block >>\n<<(macro_functions.test "z")>><<(macro_functions.test "z" )>>'
    
    print(macro_expand("<<",">>",test))
    
    #( macro_functions.block 'test')