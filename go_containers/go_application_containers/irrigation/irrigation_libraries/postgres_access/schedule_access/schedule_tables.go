package irr_sched_access

import (
   // "fmt"
    "lacima.com/server_libraries/postgres"
)


type Schedule_data_type struct {
    
    Server_key    string
    Name            string
    Description   string
    Json_data     string 
}   



func   Sched_Vacuum()bool{
    
 
    return control_block.sched_driver.Vacuum()
    
}

func   Sched_Drop()bool{
    
 
    return control_block.sched_driver.Drop_table()
    
}

func Schedule_Select_All()([]Schedule_data_type,bool){
    

    raw_data,status := control_block.sched_driver.Select_All()
    return_value := make([]Schedule_data_type,len(raw_data))
    
    for index,value := range raw_data {
       var temp Schedule_data_type
       temp.Server_key     = value.Tag1
       temp.Name              = value.Tag2
       temp.Description    = value.Tag3
       
       
       
       temp.Json_data      = value.Data
       return_value[index] = temp
       
    
    }
    return return_value,status
}  
    
    
func Delete_schedule_data( input Schedule_data_type)bool{
    
     where_entries := make(map[string]string)
     where_entries["Tag1"] = input.Server_key
     where_entries["Tag2"] = input.Name
    
     
     return control_block.sched_driver.Delete_Entry(where_entries)
    
}
   
func Insert_schedule_data( input Schedule_data_type)bool{ 
    
     var output pg_drv.Json_Table_Record
    
    
    
    output.Tag1  = input.Server_key
    output.Tag2  = input.Name
    output.Tag3  = input.Description
    output.Tag4  = ""
    output.Tag5  = ""
    output.Tag6  = ""
    output.Tag7  = ""
    output.Tag8  = ""
    output.Tag9  = ""
    output.Tag10  = ""
    output.Data  = input.Json_data 

    return control_block.sched_driver.Insert(output)
     
   

}

func Select_schedule_data(server_key string)([]Schedule_data_type,bool){
    
    where_entries := make(map[string]string)
    where_entries["Tag1"] = server_key

   
   
    raw_data,status := control_block.sched_driver.Select_tags(where_entries)
    return_value := make([]Schedule_data_type,len(raw_data))
    
    for index,value := range raw_data {
       var temp Schedule_data_type
       temp.Server_key     = value.Tag1
       temp.Name              = value.Tag2
       temp.Description    = value.Tag3
    
       temp.Json_data      = value.Data
       return_value[index] = temp
       
    
    }
    return return_value,status
}  
    
func Select_schedule_name(server_key,name string)([]Schedule_data_type,bool){
    
    where_entries := make(map[string]string)
    where_entries["Tag1"] = server_key
    where_entries["Tag2"] = name

   
   
    raw_data,status := control_block.sched_driver.Select_tags(where_entries)
    return_value := make([]Schedule_data_type,len(raw_data))
    
    for index,value := range raw_data {
       var temp Schedule_data_type
       temp.Server_key     = value.Tag1
       temp.Name              = value.Tag2
       temp.Description    = value.Tag3
    
       temp.Json_data      = value.Data
       return_value[index] = temp
       
    
    }
    return return_value,status
}  
    
    
    
