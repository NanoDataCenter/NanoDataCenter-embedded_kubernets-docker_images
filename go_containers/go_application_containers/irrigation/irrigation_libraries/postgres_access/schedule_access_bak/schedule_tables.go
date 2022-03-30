package irr_sched_access

import (
   // "fmt"
    "lacima.com/server_libraries/postgres"
)


type Schedule_data_type struct {
    
    Master_server string
    Sub_server    string
    Name          string
    Description   string
    Json_data     string 
}   

type Schedule_delete_type struct {
      
    Master_server string
    Sub_server    string
    Name          string  
    
}

func   Sched_Vacuum()bool{
    
 
    return control_block.sched_driver.Vacuum()
    
}

func Schedule_Select_All()([]Schedule_data_type,bool){
    

    raw_data,status := control_block.sched_driver.Select_All()
    return_value := make([]Schedule_data_type,len(raw_data))
    
    for index,value := range raw_data {
       var temp Schedule_data_type
       temp.Master_server  = value.Tag1
       temp.Sub_server     = value.Tag2
       temp.Name           = value.Tag3
       temp.Description    = value.Tag4
       
       
       
       temp.Json_data      = value.Data
       return_value[index] = temp
       
    
    }
    return return_value,status
}  
    
    
func Delete_schedule_data( input Schedule_delete_type)bool{
    
     where_entries := make(map[string]string)
     where_entries["tag1"] = input.Master_server
     where_entries["tag2"] = input.Sub_server
     where_entries["tag3"] = input.Name
     
     return control_block.sched_driver.Delete_Entry(where_entries)
    
}
   
func Insert_schedule_data( input Schedule_data_type)bool{ 
    
     var output pg_drv.Table_Output_Data_Record
    
    
    
    output.Tag1  = input.Master_server
    output.Tag2  = input.Sub_server
    output.Tag3  = input.Name
    output.Tag4  = input.Description
    output.Tag5  = ""
    output.Data  = input.Json_data 

    return control_block.sched_driver.Insert(output)
     
   

}

func Select_schedule_data(master_controller,sub_server string)([]Schedule_data_type,bool){
    
    where_entries := make(map[string]string)
    where_entries["tag1"] = master_controller
    where_entries["tag2"] = sub_server
   
   
    raw_data,status := control_block.sched_driver.Select_tags(where_entries)
    return_value := make([]Schedule_data_type,len(raw_data))
    
    for index,value := range raw_data {
       var temp Schedule_data_type
       temp.Master_server  = value.Tag1
       temp.Sub_server     = value.Tag2
       temp.Name           = value.Tag3
       temp.Description    = value.Tag4
       
       
       
       temp.Json_data      = value.Data
       return_value[index] = temp
       
    
    }
    return return_value,status
}  
    

    
    
