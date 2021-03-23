


        
class Extract_and_Build_Web_Structure(object):

   def __init__(self,parent_self):
        self.parent_self = parent_self
       
        self.web_properties = self.parent_self.server_properties["menu"]
        self.generate_top_menu()
   
   def flatten_list(self,data):
       return_value = []
       for i in data:
           if type(i) == list:
              return_value.append(self.flatten_list(i))
           elif type(i)== str:
               return_value.append(i)
       return return_value
        
        
   def generate_top_menu(self):
      
       self.base_html = ""
       menu_data = []
       menu_data.append(self.define_menu_top())
       children = self.web_properties["children"]
       menu_data.append(self.generate_menu(0,"/",children))
       menu_data.append(self.define_menu_bottom())
      
       menu_text = "\n".join(menu_data)
      
       
       self.write_out_data(menu_text)
   
   def write_out_data(self,data):
       out_file = open("templates/menu/menu","w")
       out_file.write(data)
       out_file.close()

   
   def define_menu_top(self):
       return '''

    <nav class="navbar navbar-expand-lg navbar-light bg-light" id="main_navbar">
        <a class="navbar-brand" href="#"></a>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent"
            aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>
       
        <div class="collapse navbar-collapse" id="navbarSupportedContent">
            <ul class="navbar-nav mr-auto">
                
                <li class="nav-item dropdown">
                    <a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown"
                        aria-haspopup="true" aria-expanded="false">
                        Menu
                    </a>
                    <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
       '''
       
   def define_menu_bottom(self):
       return '''
                          </ul>
                </li>
               
            </ul>
           
        </div>
		<h5 id="status_display">Status: </h5>
    </nav>
       '''
       
 
   def generate_menu(self,level,base_html,children):

       return_value = []
       for i in children:
         
          if i["type"] == True:
              
              return_value.append(self.sub_element_top(i["display_name"]))
              return_value.append(self.generate_menu(level+1,base_html+i["display_name"]+"/",i["children"]))
              return_value.append(self.sub_element_bottom())
              
          else:
              return_value.append(self.generate_html_element(base_html,i))

                     
                     
       return "\n".join(return_value)      
              
              
   def generate_html_element(self,base_html,i):
        target = base_html+i["display_name"]
        return'<li><a class="dropdown-item" href="'+ target+'">'+i["display_name"]+'</a></li>\n'
       
        
   def sub_element_top(self,display):
       return_value = []
       return_value.append('<li class="nav-item dropdown">')
       return_value.append('<a class="nav-link dropdown-toggle" href="#" id="'+display+'" role="button" data-toggle="dropdown"')
       return_value.append('aria-haspopup="true" aria-expanded="false">')
       return_value.append(display)
       return_value.append('</a>')
       return_value.append('<ul class="dropdown-menu" aria-labelledby="navbarDropdown">')
       return "\n".join(return_value)
       
   
   def sub_element_bottom(self):
       return_value = []
       return_value.append('</li>')
       return_value.append('</ul>')
       return "\n".join(return_value)
               
         
   
   
