package irr_sched_access

import (
    "lacima.com/server_libraries/postgres"
)

type Action_data_type struct {
    
     Server_type bool
    Master_server string
    Sub_server    string
    Name             string
    Start_time      float64
    End_time      float64
    Data     string 
}   



    
func Delete_action_data( input Action_data_type)bool{
    
     where_entries := make(map[string]string)
     if input.Server_type == true {
        where_entries["tag1"] = "true"
     }else{
         where_entries["tag1"] = "false"
     }
     where_entries["tag2"] = input.Master_server
     where_entries["tag3"] = input.Sub_server
     where_entries["tag4"] = input.Name
     
     return control_block.action_driver.Delete_Entry(where_entries)
    
}

func convert_format( input Action_data_type)pg_drv.Float_Output_Data_Record{
    
    var output  pg_drv.Float_Output_Data_Record
    if input.Server_type == true {
       output.Text1 = "true"
     }else{
        output.Text1= "false"
     }
    
    output.Text2        =  input.Master_server
    output.Text3        =  input.Sub_server
    output.Text4        =  input.Name
    output.Text5        =  ""
    output.Text6        =  ""
    output.Text7        =  ""
    output.Text8        =  ""
    output.Text9        =  ""
    output.Text10      =  ""
    
    
    
    output.Float1       =  input.Start_time
    output.Float2       =  input.End_time
    output.Float3       =  0.0 
    output.Float4       =  0.0
    output.Float5       =  0.0
    output.Float6       =  0.0
    output.Float7       =  0.0
    output.Float8       =  0.0  
    output.Float9       =  0.0
    output.Float10      =  0.0 
    output.Data         = input.Data
    
    return output
    
    
}
    
func Insert_action_data( input Action_data_type)bool{ 
    
    Delete_action_data(input)
    output := convert_format(input)

   return  control_block.action_driver.Insert(output)

}

func Select_action_data(master_controller,sub_server string)([]string,bool){
    
    where_entries := make(map[string]string)
    where_entries["text1"] = master_controller
    where_entries["text2"] = sub_server
   
   
    raw_data,status := control_block.action_driver.Select_tags(where_entries)
    return_value := make([]string,len(raw_data))
    
    for index,value := range raw_data {
       
       return_value[index] = value.Data
    }
    return return_value,status
}  
    
    
    
