class Base_Template_Class(object):
    def __init__(self,base_self,parameters):
        self.base_self = base_self
        self.parameters = parameters
        self.base_top = self.base_self.bp.base_page_top
        self.base_bottom = self.base_self.bp.base_page_bottom
        self.mp = self.base_self.mp

    def construct_web_page(self):
        self.mp.title = self.parameters["display_name"]
        self.construct_top = self.mp.macro_expand_start("<<",">>",self.base_top)
        self.construct_bottom = self.mp.macro_expand_start("<<",">>",self.base_bottom)
        self.application = self.application_page_contruction()
        
    def application_page_contruction(self): # method is to be overriden
        return ""
    
    def application_page_generation(self,data):
        return ""
    
    def render_page(self):
        output_page = []
        output_page.append(self.construct_top) 
        output_page.append(self.application_page_generation(self.application))
        output_page.append(self.construct_bottom)
        
        return "\n".join(output_page)
        
    
