
import hy
import hy_name_space


#create data 
class Macro_Base(object):

    def __init__(self,macro_file,create_data, execute_data):
        self.create_symbols = create_symbols
        macro_file_handle = open(macro_file,"r")
        self.macro_raw = macro_file_handle.read()
        macro_file_handle.close()
        
        
        
    
    
    def lisp_eval(self, expression):
        x = hy.read_str(expression)
        return str(hy.eval(x))

    def macro_expand(self, start,end,parameters,raw_string):
        return_value = []
        start_list = raw_string.split(start)
        return_value.append(start_list.pop(0))
        for i in start_list:
           temp = i.split(end)
           print(temp)
        return_value.append(lisp_eval(temp[0]))
        if len(temp) > 1 :
             return_value.append(temp[1])
        return "".join(return_value)
   
   
    def macro_expand_file(self,start,end,path):
        macro_file = open(path,'r')
        raw_string = macro_file.read()
        return self.macro_expand(start,end,raw_string)
    






