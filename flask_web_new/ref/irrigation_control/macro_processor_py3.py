import hy

class Macro_Expansion(object):
    def __init__(self):
        pass

    def lisp_eval(self,expression):

        x = hy.read_str(expression)
   
        return str(hy.eval(x))



    def macro_expand(self,start,end,raw_string):
   

       return_value = []
       start_list = raw_string.split(start)
       return_value.append(start_list.pop(0))

       for i in start_list:
     
           temp = i.split(end)
           print(temp)
           return_value.append(self.lisp_eval(temp[0]))
           if len(temp) > 1 :
              return_value.append(temp[1])
     
       return "".join(return_value)
   
   
    def macro_expand_file(self,start,end,path):
         macro_file = open(path,'r')
         raw_string = macro_file.read()
         print(raw_string)
         return self.macro_expand(start,end,raw_string)
    














if __name__ == "__main__":
    import hy
    import macro_functions
    macro_functions.test = test
   
    macro_functions.top_block = "test"
    test = 'abcdef\n<<( macro_functions.block "echo_block")>>\nabcde<<macro_functions.top_block >>\n<<(macro_functions.test "z")>><<(macro_functions.test "z" )>>'
    
    print(macro_expand("<<",">>",test))
    
    #( macro_functions.block 'test')