package irr_sched_access

import (
    //"fmt"
    "lacima.com/server_libraries/postgres"
    "encoding/json"
)

type Action_data_type struct {
    
     Server_type    bool
    Master_server string
    Sub_server      string
    Name               string
    Description     string
    Data                 string 
}   

func   Action_Vacuum()bool{
    
 
    return control_block.sched_driver.Vacuum()
    
}

func Action_Select_All()([]map[string]interface{},bool){
    

   
    raw_data,status := control_block.action_driver.Select_All()
    return_value := make([]map[string]interface{},len(raw_data))
    
    for index,value := range raw_data {
        
        item := make(map[string]interface{})
        err := json.Unmarshal([]byte(value.Data),&item)
        if err != nil{
	         panic("bad json data")
	    }
        return_value[index] = item
    }
    return return_value,status
}  
    
    
func Delete_action_data( input Action_data_type)bool{
    
     where_entries := make(map[string]string)
     if input.Server_type == true {
        where_entries["Tag1"] = "true"
     }else{
         where_entries["Tag1"] = "false"
     }
     where_entries["Tag2"]  = input.Master_server
     where_entries["Tag3"]  = input.Sub_server
     where_entries["Tag4"]  = input.Name
     
     status :=  control_block.action_driver.Delete_Entry(where_entries)
     
     return status
}

func convert_format( input Action_data_type)pg_drv.Json_Table_Record{
    
    var output  pg_drv.Json_Table_Record
    if input.Server_type == true {
       output.Tag1 = "true"
     }else{
        output.Tag1= "false"
     }
    
    output.Tag2        =  input.Master_server
    output.Tag3        =  input.Sub_server
    output.Tag4        =  input.Name
    output.Tag5        =  input.Description
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

func Select_action_data(server_type, master_controller,sub_server string)([]string,bool){
    
    where_entries := make(map[string]string)
    where_entries["Tag1"] = server_type
    where_entries["Tag2"] =master_controller
    where_entries["Tag3"] = sub_server
    
  
   
    raw_data,status := control_block.action_driver.Select_tags(where_entries)
    if status == false {
        
        
    }
    
    return_value := make([]string,len(raw_data))
    
    for index,value := range raw_data {
       
       return_value[index] = value.Data
    }
    
    return return_value,status
}  
    
    
    
