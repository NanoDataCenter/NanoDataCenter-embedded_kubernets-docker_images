package irr_files
import "lacima.com/server_libraries/file_server_library"
import "lacima.com/redis_support/redis_file"

type Irrigation_File_Manager_Type struct {

   fs                 file_server_lib.File_Server_Client_Type
   redis_file_driver  *redis_file.Redis_File_Struct 
   application_path   string
   system_path        string

}


func Initialization( address string, port,file_db int) Irrigation_File_Manager_Type{

  var return_value Irrigation_File_Manager_Type
  return_value.fs                      = file_server_lib.File_Server_Init(&[]string{"FILE_SERVER"})
  redis_file.Create_redis_data_handle(address, port , file_db )
  return_value.redis_file_driver       = redis_file.Construct_File_Struct(  ) 
  return_value.application_path        = "application_files"
  return_value.system_path             = "system_files"
  return return_value  
}



func( v *Irrigation_File_Manager_Type)Read_App_File( file_name string )(string,bool){
  path := v.application_path+"/"+file_name
  return v.redis_file_driver.Get(path)
}

func( v *Irrigation_File_Manager_Type)Write_App_File( file_name,data string )bool{

  path := v.application_path+"/"+file_name
  v.redis_file_driver.Set(path,data)
  return v.fs.Write_file(path,data)
}


func( v *Irrigation_File_Manager_Type)Delete_App_File( file_name string )bool{

  path := v.application_path+"/"+file_name
  v.redis_file_driver.Rm(path)
  return v.fs.Delete_file(path)

}

func( v *Irrigation_File_Manager_Type)Copy_App_File( from_file_name,to_file_name string )bool{

  from_path := v.application_path+"/"+from_file_name
  to_path   := v.application_path+"/"+to_file_name
  if from_path == to_path {
    return false
  }
  data, err := v.Read_App_File(from_path)
  if err != false {
    return v.Write_App_File(to_path,data)
  }
  return false
  


}

func( v *Irrigation_File_Manager_Type)App_Ls( file_name string )[]string{

  path := v.application_path
  return v.redis_file_driver.Ls(path)
  

}

func( v *Irrigation_File_Manager_Type)Read_Sys_File( file_name string )(string,bool){
  path := v.system_path+"/"+file_name
  return v.redis_file_driver.Get(path)
}

func( v *Irrigation_File_Manager_Type)Write_Sys_File( file_name,data string )bool{

  path := v.system_path+"/"+file_name
  v.redis_file_driver.Set(path,data)
  return v.fs.Write_file(path,data)
}


func( v *Irrigation_File_Manager_Type)Delete_Sys_File( file_name string )bool{

  path := v.system_path+"/"+file_name
  v.redis_file_driver.Rm(path)
  return v.fs.Delete_file(path)

}

func( v *Irrigation_File_Manager_Type)Copy_Sys_File( from_file_name,to_file_name string )bool{

  from_path := v.system_path+"/"+from_file_name
  to_path   := v.system_path+"/"+to_file_name
  if from_path == to_path {
    return false
  }
  data, err := v.Read_App_File(from_path)
  if err != false {
    return v.Write_App_File(to_path,data)
  }
  return false
  


}

func( v *Irrigation_File_Manager_Type)Sys_Ls( file_name string )[]string{

  path := v.system_path
  return v.redis_file_driver.Ls(path)
  

}

/*
**
**
**
*/

type Time_Type struct {
  Hour    int
  Minute  int
}

type Action_File_Type struct {
  Enable               string 
  End_time             Time_Type
  Start_time           Time_Type
  Dow                  [7]int
  Web_command_string   string
  Name                 string
}

type Action_File_Array_Type struct {

   Array_Element []Action_File_Type

}

func( v *Irrigation_File_Manager_Type)Load_Action_File( ) Action_File_Array_Type {

  
  

}








