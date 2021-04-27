from templates.Base_Template_Class_py3  import Base_Template_Class

class Past_Actions( Base_Template_Class):
   def __init__(self,base_self,properties):
       Base_Template_Class.__init__(self,base_self,properties)
       
   def render_setup(self): # method is to be overriden  Setup variables for macro processing
       self.application = "<h4>This is a Test</h4>"    

