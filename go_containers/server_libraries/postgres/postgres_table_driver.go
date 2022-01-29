package pg_drv


import (
    
    "fmt"   
    "strings"
    //"strconv"
    "time"
	"context"
	//"github.com/jackc/pgx/v4"   
	
)
    

type Table_Output_Data_Record struct {
 
    Stream_id  int64;
    Tag1       string;  
    Tag2       string;  
    Tag3       string;  
    Tag4       string;  
    Tag5       string;
    Data       string;
    Time_stamp int64;
}    



type Postgres_Table_Driver struct {
    
     Postgres_Basic_Driver
     key           string
     user          string
     password      string
     database      string
     table_name    string
     
    
}




func Construct_Postgres_Table_Driver( key,user,password,database, table_name string) Postgres_Table_Driver{
    var return_value Postgres_Table_Driver
    return_value.key            = key
    return_value.user           = user
    return_value.password       = password
    return_value.database       = database
    return_value.table_name     = table_name
    
    
    return return_value
}

func ( v  Postgres_Table_Driver )Vacuum( )bool{
    
    
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


func ( v  Postgres_Table_Driver )Drop_table(  )bool{
    script := "DROP TABLE IF EXISTS  "+v.table_name+";" 
    return v.Exec( script  )
    
}


func ( v  Postgres_Table_Driver )create_table( )bool{
   script_array := make([]string,9)
   script_array[0] = "CREATE TABLE IF NOT EXISTS  "+ v.table_name +"( "
   script_array[1] = "stream_id BIGSERIAL PRIMARY KEY,"
   script_array[2] = "tag1 text,"
   script_array[3] = "tag2 text,"
   script_array[4] = "tag3 text,"
   script_array[5] = "tag4 text,"
   script_array[6] = "tag5 text,"
   script_array[7] = "data text,"
   script_array[8] = "time bigint );"
   script := strings.Join(script_array," ")
   return v.Exec( script  )
}


func ( v Postgres_Table_Driver)Exists(tags map[string]string)(int,bool){

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

func ( v  Postgres_Table_Driver )Update( tag1,tag2,tag3,tag4,tag5,data string )bool{
    
  time_stamp    := time.Now().UnixNano()
  
     
  
   update_part := fmt.Sprint(" SET tag1 ='%s',tag2 ='%s',tag3 ='%s',tag4 ='%s',tag5 ='%s', data ='%s',time =%d ; " ,tag1,tag2,tag3,tag4,tag5,data,time_stamp)
   script := "UPDATE INTO "+v.table_name+update_part

  status :=  v.Exec( script  )
 
  return status
}

func ( v  Postgres_Table_Driver )Insert( tag1,tag2,tag3,tag4,tag5,data string )bool{
    
  time_stamp    := time.Now().UnixNano()
  
     
  
 
  script := fmt.Sprintf("INSERT INTO %s (tag1,tag2,tag3,tag4,tag5,data,time ) VALUES('%s','%s','%s','%s','%s','%s',%d);",v.table_name,tag1,tag2,tag3,tag4,tag5,data,time_stamp)

    status :=  v.Exec( script  )
 
  return status
}



func ( v   Postgres_Table_Driver)Delete_Entry( tags map[string]string)bool{
    
    where_clause := v.tags_where_clause(tags)
    script := "DELETE FROM "+v.table_name+" where "+where_clause+";"
    fmt.Println("delete script ",script)
    return v.Exec(script)
    
}



func (v Postgres_Table_Driver)tags_where_clause( tags map[string]string)string{
    
   where_clause_array := make([]string,0)
   for key,value := range tags{
       entry := `( `+key+`  = "`+value+`" )`
       where_clause_array = append(where_clause_array,entry)
   }
   where_clause := strings.Join(where_clause_array," AND ")   
   where_clause = where_clause
   return where_clause
}

func (v  Postgres_Table_Driver)Select_tags(tags map[string]string)([]Table_Output_Data_Record, bool){
   where_clause := v.tags_where_clause( tags )
   return v.Select_where(where_clause)
    
}

func (v  Postgres_Table_Driver)Select_where(where_clause string)([]Table_Output_Data_Record, bool){
    
    
    
    return_value := make([]Table_Output_Data_Record,0)
    
    script := "Select * from "+v.table_name +" where "+where_clause+";"
    //fmt.Println("script",script)
    rows, err := v.conn.Query(context.Background(), script)
    if err != nil {
      fmt.Println("err",err)
      return return_value, false
    }
    defer rows.Close()

    for rows.Next() {
            
            var item Table_Output_Data_Record
            rows.Scan(&item.Stream_id,&item.Tag1,&item.Tag2,&item.Tag3,&item.Tag4,&item.Tag5,&item.Data,&item.Time_stamp)
            
            if rows.Err() != nil {
              return return_value,false
            }
            return_value = append(return_value,item)
              
        }
    
  
    return return_value,true
    
}
func (v  Postgres_Table_Driver)Select_All()([]Table_Output_Data_Record , bool){
    
    
    
    return_value := make([]Table_Output_Data_Record,0)
    
    script := "Select * from "+v.table_name +";"
    //fmt.Println("script",script)
    rows, err := v.conn.Query(context.Background(), script)
    if err != nil {
      return return_value, false
    }
    defer rows.Close()

    for rows.Next() {
            
            var item Table_Output_Data_Record
            rows.Scan(&item.Stream_id,&item.Tag1,&item.Tag2,&item.Tag3,&item.Tag4,&item.Tag5,&item.Data,&item.Time_stamp)
            
            if rows.Err() != nil {
              return return_value,false
            }
            return_value = append(return_value,item)
              
        }
    
  
    return return_value,true
    
}


