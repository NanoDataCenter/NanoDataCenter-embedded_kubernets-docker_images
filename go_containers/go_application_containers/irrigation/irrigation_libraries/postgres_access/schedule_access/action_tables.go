package irr_sched_access

import (
    "lacima.com/server_libraries/postgres"
)

type Action_data_type struct {
    
    Master_server string
    Sub_server    string
    Name          string
    Description   string
    Start_hr      float64
    Start_min     float64
    End_hr        float64
    End_min       float64
    Json_data     string 
}   

type Action_delete_type struct {
      
    Master_server string
    Sub_server    string
    Name          string  
    
}
    
    
func Delete_action_data( input Action_delete_type)bool{
    
     where_entries := make(map[string]string)
     where_entries["tag1"] = input.Master_server
     where_entries["tag2"] = input.Sub_server
     where_entries["tag3"] = input.Name
     
     return control_block.action_driver.Delete_Entry(where_entries)
    
}
   
func Insert_action_data( input Action_data_type)bool{ 
    
    var output  pg_drv.Float_Output_Data_Record
    
    output.Text1        =  input.Master_server
    output.Text2        =  input.Sub_server
    output.Text3        =  input.Name
    output.Text4        =  input.Description
    output.Text5        =  ""
    output.Text6        =  ""
    output.Text7        =  ""
    output.Text8        =  ""
    output.Text9        =  ""
    output.Text10       =  ""
    
    output.Float1       =  input.Start_hr
    output.Float2       =  input.Start_min
    output.Float3       =  input.End_hr 
    output.Float4       =  input.End_min
    output.Float5       =  0.0
    output.Float6       =  0.0
    output.Float7       =  0.0
    output.Float8       =  0.0  
    output.Float9       =  0.0
    output.Float10      =  0.0 
    output.Data         = input.Json_data
    
    return control_block.action_driver.Insert(output)
     
   

}

func Select_action_data(master_controller,sub_server string)([]Action_data_type,bool){
    
    where_entries := make(map[string]string)
    where_entries["text1"] = master_controller
    where_entries["text2"] = sub_server
   
   
    raw_data,status := control_block.action_driver.Select_tags(where_entries)
    return_value := make([]Action_data_type,len(raw_data))
    
    for index,value := range raw_data {
       var temp Action_data_type
       temp.Master_server  = value.Text1
       temp.Sub_server     = value.Text2
       temp.Name           = value.Text3
       temp.Description    = value.Text4
       temp.Start_hr       = value.Float1
       temp.Start_min      = value.Float2
       temp.End_hr         = value.Float3
       temp.End_hr         = value.Float4
       temp.Json_data      = value.Data
       return_value[index] = temp
       
    
    }
    return return_value,status
}  
    
    
    
