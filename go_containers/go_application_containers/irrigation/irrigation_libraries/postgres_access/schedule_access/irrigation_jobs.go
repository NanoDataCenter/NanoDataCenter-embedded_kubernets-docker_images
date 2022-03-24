package irr_sched_access
/*
import (
   // "fmt"
    "lacima.com/server_libraries/postgres"
)

type Job_delete_type struct{
    
   Server_type           bool    
   Master_server       string  // irrigation server id
   Sub_server             string
   Name                      string
   Queued_time         float64     
}

type Job_server_type struct{
  Server_type           bool    
   Master_server       string  // irrigation server id
   Sub_server             string    
}


type Job_data_type struct {
    
   Server_type           bool    
   Master_server       string  
   Sub_server             string
   Queued_time         float64
   
   Job_state               string

   Start_time              float64
   End_time                float64
   Current_time         float64

   Type                        string
   Name                      string
   Starting_step         float64  
   Ending_step           float64
   Current_step          float64
   Job_notes              string
}   


func   Job_Vacuum()bool{
    
 
    return control_block.job_driver.Vacuum()
    
}

func Job_Trim(time_second int64)bool{
    return control_block.job_driver.Trim(time_second)   
}

 *
 *
 *  Delete Jobs
 *
 *
 *
     
func Delete_all()bool{


}




func Delete_job_data( input Job_delete_type)  bool{
 
 
   
     where_entries := make(map[string]string)
     where_entries["tag1"] = input.Status 
     where_entries["tag2"] = input.Server_type 
     where_entries["tag3"] = input.Name
     where_entries["tag1"] = input.Master_server
     where_entries["tag2"] = input.Sub_server
     where_entries["tag3"] = input.Name    
     return control_block.job_driver.Delete_Entry(where_entries)
    
}

   

 *
 *
 *  Insert Routines
 *
 *
 *
 *
    
   
func Insert_schedule_data( input Job_data_type)bool{ 
    
     var output pg_drv.Table_Output_Data_Record
    
     output.Text1                  =   convert_bool_string(output.Server_type)    
     output.Text2                  =   input.Master_server     
     output.Text3                  =   input.Sub_server             
     output.Text4                  =   convert_bool_string(input.Completed_state)    
    output.Text5                   =   input.Type                      
     output.Text6                  =   input.Name                      
     output.Float1                = input.Starting_step
     output.Float2                = input.Ending_step
     output.Float3                = input.Current_step
     output.Float4                = input.Queued_time
     output.Data                  = input.Job_notes              
    return control_block.sched_driver.Insert(output)
}

 * 
 * 
 *   Select Modules
 * 
 * 
 * 


func Make_job_active(){
    
    
    
}

func Make_job_inactive(){
    
    
}


func Make_job_completed(){
    
    
    
}

    
func Select_irrigation_jobs_(master_controller,sub_server string)([]Job_data_type,bool){
    
    where_entries := make(map[string]string)
    where_entries["tag1"] = master_controller
    where_entries["tag2"] = sub_server
   
   
    raw_data,status := control_block.sched_driver.Select_tags(where_entries)
    return_value := make([]Job_data_type,len(raw_data))
    
    for index,value := range raw_data {
        var temp Job_data_type
              temp.Server_type             = convert_text_to_bool(temp.Text1)    
              temp.Master_server        = temp.Text2 
              temp.Sub_server             = temp.Text3
              temp.Status                     = convert_text_to_bool(temp.Text4)             
              tempType                        =  temp. Text5
              temp.Name                     =  temp.Text6  
              temp.Starting_step        = temp.Float1
              temp.Ending_step          = temp.Float2
              temp.Current_step         = temp.Float3
             temp.Current_time         = temp.Float4
             temp.Start_time              = temp.Float5 
             temp.End_time                = temp.Float6
            temp Completed_state   = convert_text_to_bool(temp.Data)
             temp.Job_notes              = temp.Text8
             return_value[index] = temp
    }
    return return_value
}
*/
