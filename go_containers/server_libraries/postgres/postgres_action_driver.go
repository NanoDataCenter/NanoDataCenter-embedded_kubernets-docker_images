package pg_drv
/*

import (
    
    "fmt"   
    "strings"
    //"strconv"
    "time"
	"context"
	//"github.com/jackc/pgx/v4"   
	
)
    

type Action_Output_Data_Record struct {
 
    Stream_id          int64
    Name               string
    Master_Controller  string
    Sub_Controller     string
    Action_Type        string
    Action_Subtype     string
    Start_Time_Hr      int64
    Start_Time_Min     int64
    End_Time_Hr        int64
    End_Time_Min       int64
    Active             int64
    Time_stamp         int64
}    



type Postgres_Action_Driver struct {
    
     Postgres_Basic_Driver
     key           string
     user          string
     password      string
     database      string
     table_name    string
     
    
}




func Construct_Postgres_Table_Driver( key,user,password,database, table_name string) Postgres_Action_Driver{
    var return_value Postgres_Action_Driver
    return_value.key            = key
    return_value.user           = user
    return_value.password       = password
    return_value.database       = database
    return_value.table_name     = table_name
    
    
    return return_value
}

func ( v  Postgres_Action_Driver )Vacuum( )bool{
    
    
    script := "VACUUM "+v.table_name +";"
    return v.Exec( script  )
    
}

func ( v  *Postgres_Table_Driver )Connect( ip string )bool{
    connection_url := "postgres://"+v.user+":" + v.password + "@"+ ip+":5432" + "/"+v.database 
    if v.connect(connection_url) == false {
        return false
    }
    
	
	

	v.create_table()
    
	
	return true
}


func ( v  Postgres_Table_Driver )Create_table()bool{
    v.create_table()
    
    return true
}


func ( v  Postgres_Action_Driver )Drop_table(  )bool{
    script := "DROP TABLE IF EXISTS  "+v.table_name+";" 
    return v.Exec( script  )
    
}
 

func ( v  Postgres_Action_Driver )create_table( )bool{
   script_array := make([]string,13)
   script_array[0] =   "CREATE TABLE IF NOT EXISTS  "+ v.table_name +"( "
   script_array[1] =   "stream_id BIGSERIAL PRIMARY KEY,"
   script_array[2] =   "name              text,"
   script_array[3] =   "master_controller text,"
   script_array[4] =   "sub_controller    text,"
   script_array[5] =   "action_type       text,"
   script_array[6] =   "action_subtype    text,"
   script_array[7] =   "start_time_hr     bigint,"
   script_array[8] =   "start_time_min    bigint,"
   script_array[9] =   "end_time_hr       bigint,"
   script_array[10] =  "end_time_min      bigint,"
   script_array[11] =  "active            bigint,"
   script_array[12] = "time bigint );"
   script := strings.Join(script_array," ")
   return v.Exec( script  )
}


func ( v Postgres_Action_Driver)Exists(tags map[string]string)(int,bool){

  where_clause := v.tags_where_clause(tags)
  script := fmt.Sprintf(`Select EXISTS (SELECT * FROM %s WHERE '`+where_clause +`);` ,v.table_name)

  rows, err := v.conn.Query(context.Background(), script)
  if err != nil {
      fmt.Println("err",err)
      return 0, false
  }
  defer rows.Close()
  count := 0
  for rows.Next() {
            
            count = count+1
            if rows.Err() != nil {
              return 0,false
            }

   }
    
  return count,true
    
}

 script_array[3] =   "master_controller text,"
   script_array[4] =   "sub_controller    text,"
   script_array[5] =   "action_type       text,"
   script_array[6] =   "action_subtype    text,"
   script_array[7] =   "start_time_hr     bigint,"
   script_array[8] =   "start_time_min    bigint,"
   script_array[9] =   "end_time_hr       bigint,"
   script_array[10] =  "end_time_min      bigint,"
   script_array[11] =  "active            bigint,"
   script_array[12] = "time bigint );"

func ( v  Postgres_Action_Driver )Insert( tag1,tag2,tag3,tag4,tag5,data string )bool{
    
  time_stamp    := time.Now().UnixNano()
  format_string := "INSERT INTO %s (name,master_controller,sub_controller,action_type,action_subtype,start_time_hr   data,time )
  
  script := fmt.Sprintf("INSERT INTO %s (name,master_controller,sub_controller,tag4,tag5,data,time ) VALUES('%s','%s','%s','%s','%s','%s',%d);",v.table_name,tag1,tag2,tag3,tag4,tag5,data,time_stamp)
    fmt.Println("script",script)
    status :=  v.Exec( script  )
 
  return status
}



func ( v   Postgres_Action_Driver)Delete_Entry( tags map[string]string)bool{
    
    where_clause := v.tags_where_clause(tags)
    script := "DELETE FROM "+v.table_name+" where "+where_clause+";"
    fmt.Println("delete script ",script)
    return v.Exec(script)
    
}



func (v Postgres_Action_Driver)tags_where_clause( tags map[string]string)string{
    
   where_clause_array := make([]string,0)
   for key,value := range tags{
       entry := `( `+key+`  = '`+value+`' )`
       where_clause_array = append(where_clause_array,entry)
   }
   where_clause := strings.Join(where_clause_array," AND ")   
   where_clause = where_clause
   return where_clause
}

func (v  Postgres_Action_Driver)Select_tags(tags map[string]string)([]Action_Output_Data_Record, bool){
   where_clause := v.tags_where_clause( tags )
   return v.Select_where(where_clause)
    
}

func (v Postgres_Action_Driver)Select_where(where_clause string)([]Action_Output_Data_Record bool){
    
    
    
    return_value := make([]Action_Output_Data_Record,0)
    
    script := "Select * from "+v.table_name +" where "+where_clause+";"
    //fmt.Println("script",script)
    rows, err := v.conn.Query(context.Background(), script)
    if err != nil {
      fmt.Println("err",err)
      return return_value, false
    }
    defer rows.Close()

    for rows.Next() {
            
            var item Action_Output_Data_Record
            rows.Scan(&item.Stream_id,&item.Tag1,&item.Tag2,&item.Tag3,&item.Tag4,&item.Tag5,&item.Data,&item.Time_stamp)
            
            if rows.Err() != nil {
              return return_value,false
            }
            return_value = append(return_value,item)
              
        }
    
  
    return return_value,true
    
}
func (v  Postgres_Action_Driver)Select_All()([]Action_Output_Data_Record , bool){
    
    
    
    return_value := make([]Action_Output_Data_Record,0)
    
    script := "Select * from "+v.table_name +";"
    //fmt.Println("script",script)
    rows, err := v.conn.Query(context.Background(), script)
    if err != nil {
      return return_value, false
    }
    defer rows.Close()

    for rows.Next() {
            
            var item Action_Output_Data_Record
            rows.Scan(&item.Stream_id,&item.Tag1,&item.Tag2,&item.Tag3,&item.Tag4,&item.Tag5,&item.Data,&item.Time_stamp)
            
            if rows.Err() != nil {
              return return_value,false
            }
            return_value = append(return_value,item)
              
        }
    
  
    return return_value,true
    
}


*/
