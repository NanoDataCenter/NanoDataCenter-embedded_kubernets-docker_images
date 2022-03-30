package irr_sched_access

import (
    //"fmt"
    "lacima.com/server_libraries/postgres"
    //"encoding/json"
)

type Action_data_type struct {
    
    Server_key      string
    Name               string
    Description     string
    Data                 string 
}   

func   Action_Vacuum()bool{
    
 
    return control_block.sched_driver.Vacuum()
    
}

func   Action_Drop()bool{
    
 
    return control_block.action_driver.Drop_table()
    
}


func Action_Select_All()([]Action_data_type,bool){
   
    raw_data , err:=   control_block.action_driver.Select_All()
    if err != true{
        panic("db error")
    }
    return_value := make([]Action_data_type,len(raw_data))
    for index, value := range raw_data {
        var temp Action_data_type
        temp.Server_key     =   value.Tag1
        temp.Name              =   value.Tag2
        temp.Description    =   value.Tag3
        temp.Data                =   value.Data
        return_value[ index ] = temp      
    }
    return return_value, true
}  
    
    
func Delete_action_data( input Action_data_type)bool{
    
     where_entries := make(map[string]string)
   
     where_entries["Tag1"]  = input.Server_key
     where_entries["Tag2"]  = input.Name
    
     
     status :=  control_block.action_driver.Delete_Entry(where_entries)
     
     return status
}

func convert_format( input Action_data_type)pg_drv.Json_Table_Record{
    
    var output  pg_drv.Json_Table_Record
    
    
    output.Tag1        =  input.Server_key
    output.Tag2        =  input.Name
    output.Tag3        =  input.Description
    output.Tag4        =  ""
    output.Tag5        = ""
    output.Tag6        =  ""
    output.Tag7        =  ""
    output.Tag8        =  ""
    output.Tag9        =  ""
    output.Tag10      =  ""
    output.Data         = input.Data
    
    return output
    
    
}
    
func Insert_action_data( input Action_data_type)bool{ 
    
    
    output := convert_format(input)

   return control_block.action_driver.Insert(output)
    
     
}

func Select_action_data(server_key  string)([]string,bool){
    
    where_entries := make(map[string]string)
    where_entries["Tag1"] = server_key
    
    raw_data,status := control_block.action_driver.Select_tags(where_entries)
    if status == false {
        
        
    }
    
    return_value := make([]string,len(raw_data))
    
    for index,value := range raw_data {
       
       return_value[index] = value.Data
    }
    
    return return_value,status
}  
    
    
func Select_action_name(server_key,name  string)([]string,bool){
    
    where_entries := make(map[string]string)
    where_entries["Tag1"] = server_key
    where_entries["Tag2"] = name
    
    raw_data,status := control_block.action_driver.Select_tags(where_entries)
    if status == false {
        
        
    }
    
    return_value := make([]string,len(raw_data))
    
    for index,value := range raw_data {
       
       return_value[index] = value.Data
    }
    
    return return_value,status
}  
    
