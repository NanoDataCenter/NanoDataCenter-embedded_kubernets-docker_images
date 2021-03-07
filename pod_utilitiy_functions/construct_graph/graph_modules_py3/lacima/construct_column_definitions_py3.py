

class Column_Base_Class(object):
   def __init__(self,bc,cd):
       self.bc = bc
       self.cd = cd
      
       
   def start_column_class(self,column_name,starting_class):
       self.class_name = column_name
       self.required_classes = set()
       self.defined_classes = set()
       self.required_classes.add(starting_class)
       self.bc.add_header_node("COLUMN_ELEMENT",column_name, properties = {"class":starting_class})
       
       
   def add_iterator(self,class_name,dependent_list):
       properties = {}
       properties["name"] = class_name
       properties["type"] = "iterator"
       properties["class_list"] = dependent_list
       self.bc.add_info_node("CLASS_DEF",class_name,properties)
       self.required_classes |= set(dependent_list)
       self.defined_classes.add(class_name)       
       
       
   def add_base_class(self,class_name,physical_class):      
       properties = {}
       properties["name"] = class_name
       properties["type"] = "base"
       properties["physical_class"] = physical_class
       self.bc.add_info_node("CLASS_DEF",class_name,properties)
       self.defined_classes.add(class_name)  
       self.defined_classes.add(class_name)       
         
   def end_column_class(self, column_name):
       if len(self.required_classes) != len(self.defined_classes):
           ValueError(f"Missmatch in number or required classes { self.required_classes } and defined_classes { self.defined_classes }")
       result = self.required_classes - self.defined_classes
      
       if len(result) != 0 :
          ValueError(f"Missmatch in number or required classes { self.required_classes } and defined_classes { self.defined_classes }")
       
       self.bc.end_header_node("COLUMN_ELEMENT")  


class Construct_Column_Definitions(Column_Base_Class):

   def __init__(self,bc,cd):
       Column_Base_Class.__init__(self,bc,cd)
       
       
       
       
       bc.add_header_node("COLUMN_DEFINITIONS","label")
       
       self.start_column_class("test_column","A")
       self.add_iterator("A",["B","C","D","E"])
       self.add_base_class("B","CHECK_POWER_SUPPLY")
       self.add_base_class("C","CHECK_MODBUS")
       self.add_base_class("D","CHECK_IRRIGATION")
       self.add_base_class("E","PERFORM_IRRIGATION")
       self.end_column_class("test_column")

   
   
   
       bc.add_header_node("GLOBAL_EVENTS")
       cd.construct_package("WIFI_DATA_STRUCTURES")
       cd.add_hash("DEVICE_STATUS")
       cd.add_redis_stream("DEVICE_STATUS_LOG")
       cd.close_package_contruction()
       bc.end_header_node("GLOBAL_EVENTS")  
       
       bc.end_header_node("COLUMN_DEFINITIONS") 
       


   
 