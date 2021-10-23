package pg_drv

import (
    
    "fmt"   
    "strings"
    //"strconv"
    "time"
	"context"
	//"github.com/jackc/pgx/v4"   
	b64 "encoding/base64"
)
    

type Stream_Output_Data_Record struct {
 
    Stream_id  int64;
    Tag1       string;
    Tag2       string;
    Tag3       string;
    Tag4       string;
    Tag5       string;
    Data       string;
    Time_stamp int64;
}    



type Postgres_Stream_Driver struct {
    
     Postgres_Basic_Driver
     key           string
     user          string
     password      string
     database      string
     table_name    string
     time_limit    int64  // time length of stream in nano seconds
     
    
}




func Construct_Postgres_Stream_Driver( key,user,password,database, table_name string, time_limit int64) Postgres_Stream_Driver{
    var return_value Postgres_Stream_Driver
    return_value.key            = key
    return_value.user           = user
    return_value.password       = password
    return_value.database       = database
    return_value.table_name     = table_name
    return_value.time_limit     = time_limit
    
    return return_value
}


func ( v  *Postgres_Stream_Driver )Connect( ip string )bool{
    connection_url := "postgres://"+v.user+":" + v.password + "@"+ ip+":5432" + "/"+v.database 
    if v.connect(connection_url) == false {
        return false
    }
    
	
	

	v.create_table()
    v.create_index()
	
	return true
}


func ( v  Postgres_Stream_Driver )Create_table()bool{
    v.create_table()
    v.create_index()
    return true
}


func ( v  Postgres_Stream_Driver )Drop_table(  )bool{
    script := "DROP TABLE IF EXISTS  "+v.table_name+";" 
    return v.Exec( script  )
    
}


func ( v  Postgres_Stream_Driver )create_table( )bool{
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

func ( v  Postgres_Stream_Driver )create_index()bool{
    script := "CREATE INDEX time_idx ON "+v.table_name+ "(time);"
    return v.Exec( script  )
}    
    




func ( v  Postgres_Stream_Driver )Insert( tag1,tag2,tag3,tag4,tag5,data string )bool{
    
  time_stamp    := time.Now().UnixNano()
  b64_data      := b64.StdEncoding.EncodeToString([]byte(data))
     
  
 
  script := fmt.Sprintf("INSERT INTO %s (tag1,tag2,tag3,tag4,tag5,data,time ) VALUES('%s','%s','%s','%s','%s','%s',%d);",v.table_name,tag1,tag2,tag3,tag4,tag5,b64_data,time_stamp)

    status :=  v.Exec( script  )
 
  return status
}

/*
 *  time_time is the number of seconds in the past
 * 
 */
func ( v Postgres_Stream_Driver )Trim( trim_time_second int64  )bool{

    current_time := time.Now().UnixNano()
    delete_time  := current_time - trim_time_second *1000000000 
    script := fmt.Sprintf("DELETE FROM %s WHERE time < %d ;",v.table_name, delete_time)
    
    return v.Exec(script)
    
}

func ( v  Postgres_Stream_Driver )Vacuum( )bool{
    
    
    script := "VACUUM "+v.table_name +";"
    return v.Exec( script  )
    
}



func   Assemble_key( input Stream_Output_Data_Record)string{

     return_value := "/"+input.Tag1+"/"+input.Tag2+"/"+input.Tag3+"/"+input.Tag4+"/"+input.Tag5
     return return_value
    
}


func (v Postgres_Stream_Driver)Select_where(where_clause string)([]Stream_Output_Data_Record, bool){
    
    var data_value string
    
    return_value := make([]Stream_Output_Data_Record,0)
    
    script := "Select * from "+v.table_name +" where "+where_clause+ ";"
    //fmt.Println("script",script)
    rows, err := v.conn.Query(context.Background(), script)
    if err != nil {
      fmt.Println("err",err)
      return return_value, false
    }
    defer rows.Close()

    for rows.Next() {
            
            var item Stream_Output_Data_Record
            rows.Scan(&item.Stream_id,&item.Tag1,&item.Tag2,&item.Tag3,&item.Tag4,&item.Tag5,&data_value,&item.Time_stamp)
            temp, _ := b64.StdEncoding.DecodeString(data_value)
            item.Data = string(temp)
            if rows.Err() != nil {
              return return_value,false
            }
            return_value = append(return_value,item)
              
        }
    
  
    return return_value,true
    
}
func (v Postgres_Stream_Driver)Select_All()([]Stream_Output_Data_Record, bool){
    
    var data_value string
    
    return_value := make([]Stream_Output_Data_Record,0)
    
    script := "Select * from "+v.table_name +";"
    //fmt.Println("script",script)
    rows, err := v.conn.Query(context.Background(), script)
    if err != nil {
      return return_value, false
    }
    defer rows.Close()

    for rows.Next() {
            
            var item Stream_Output_Data_Record
            rows.Scan(&item.Stream_id,&item.Tag1,&item.Tag2,&item.Tag3,&item.Tag4,&item.Tag5,&data_value,&item.Time_stamp)
            temp, _ := b64.StdEncoding.DecodeString(data_value)
            item.Data = string(temp)
            if rows.Err() != nil {
              return return_value,false
            }
            return_value = append(return_value,item)
              
        }
    
  
    return return_value,true
    
}

func (v Postgres_Stream_Driver)Select_after_time_stamp_desc( timestamp int64)([]Stream_Output_Data_Record, bool){
    
    var data_value string
    
    return_value := make([]Stream_Output_Data_Record,0)
    
    current_time := time.Now().UnixNano()
    select_time  := current_time - timestamp *1000000000 
    script := fmt.Sprintf("Select * from %s WHERE time >= %d ORDER BY time DESC ;",v.table_name, select_time)
    //fmt.Println("script",script)
    rows, err := v.conn.Query(context.Background(), script)
    if err != nil {
      return return_value, false
    }
    defer rows.Close()

    for rows.Next() {
            
            var item Stream_Output_Data_Record
            rows.Scan(&item.Stream_id,&item.Tag1,&item.Tag2,&item.Tag3,&item.Tag4,&item.Tag5,&data_value,&item.Time_stamp)
            temp, _ := b64.StdEncoding.DecodeString(data_value)
            item.Data = string(temp)
            if rows.Err() != nil {
              return return_value,false
            }
            return_value = append(return_value,item)
              
        }
    
  
    return return_value,true
    
}

func (v Postgres_Stream_Driver)Select_after_time_stamp_asc( timestamp int64)([]Stream_Output_Data_Record, bool){
    
    var data_value string
    
    return_value := make([]Stream_Output_Data_Record,0)
    
    current_time := time.Now().UnixNano()
    select_time  := current_time - timestamp *1000000000 
    script := fmt.Sprintf("Select * from %s WHERE time >= %d ORDER BY time ASC  ;",v.table_name, select_time)
    //fmt.Println("script",script)
    rows, err := v.conn.Query(context.Background(), script)
    if err != nil {
      return return_value, false
    }
    defer rows.Close()

    for rows.Next() {
            
            var item Stream_Output_Data_Record
            rows.Scan(&item.Stream_id,&item.Tag1,&item.Tag2,&item.Tag3,&item.Tag4,&item.Tag5,&data_value,&item.Time_stamp)
            temp, _ := b64.StdEncoding.DecodeString(data_value)
            item.Data = string(temp)
            if rows.Err() != nil {
              return return_value,false
            }
            return_value = append(return_value,item)
              
        }
    
  
    return return_value,true
    
}

