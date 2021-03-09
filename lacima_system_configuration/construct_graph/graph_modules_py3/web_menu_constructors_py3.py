class Menu_Header(object):
    def __init__(self,display_name):
        self.menu_type = True
        self.display_name = display_name
        self.children = []
        
    
    def add_child(self, child):
        self.children.append(child)
    
    def generate_data_structure(self):
        return_value = {}
        return_value["display_name"] = self.display_name
        return_value["type"] = self.menu_type
        children = []
        for i in self.children:
            children.append( i.generate_data_structure())
        return_value["children"] = children
        return return_value
        
class Menu_Element(object):
    def __init__(self,display_name,class_name,parameters={}):
        self.menu_type = False
        self.display_name = display_name
        self.class_name = class_name
        self.parameters = parameters
        
        
    
    def generate_data_structure(self):
        return_value = {}
        return_value["display_name"] = self.display_name
        return_value["class_name"] = self.class_name
        return_value["parameters"] = self.parameters
        return_value["type"] = self.menu_type
        return return_value      