from templates.Base_Template_Class_py3  import Base_Template_Class

class Site_Manager( Base_Template_Class):
   def __init__(self,base_self,properties):
      Base_Template_Class.__init__(self,base_self,properties)
      
      
      
   def application_page_generation(self,data): # method is to be overriden  Setup variables for macro processing
           

        web_links = list(self.base_self.web_map.keys())
        print("web_links",web_links)
        web_links.sort()
        return_value = []
        return_value.append("<h3>Web Links For Web Site</h3>")
        return_value.append("<ul>")
        for i in web_links:
            return_value.append('<li><a href="'+i+'">'+i+'</a></li>')
        return_value.append("</ul>")
        return "".join(return_value)
      
      
  