package pg_drv

import (
    
    "fmt"   
    "strings"
    //"strconv"
    "time"
	"context"

)

type Json_Table_Record struct {
     Stream_id  int64;
    Tag1        string;  
    Tag2        string;  
    Tag3        string;  
    Tag4        string;  
    Tag5        string;
    Tag6        string;  
    Tag7        string;  
    Tag8        string;  
    Tag9        string;  
    Tag10       string;
    

    Data      string;
    Time_stamp int64;
}    


type Json_Table_Driver struct {
    
     Postgres_Basic_Driver
     key           string
     user          string
     password      string
     database      string
     table_name    string
     time_limit    int64  // time length of stream in nano seconds
     
    
}




func Construct_Json_Table_Driver( key,user,password,database, table_name string)Json_Table_Driver{
    var return_value Json_Table_Driver
    return_value.key            = key
    return_value.user           = user
    return_value.password       = password
    return_value.database       = database
    return_value.table_name     = table_name

    
    return return_value
}


func ( v  *Json_Table_Driver )Connect( ip string )bool{
    connection_url := "postgres://"+v.user+":" + v.password + "@"+ ip+":5432" + "/"+v.database 
    if v.connect(connection_url) == false {
        return false
    }
    
	
	

	v.create_table()
    v.create_index()
	
	return true
}


func ( v  Json_Table_Driver )Create_table()bool{
    v.create_table()
    v.create_index()
    return true
}


func ( v  Json_Table_Driver )Drop_table(  )bool{
   
    script := "DROP TABLE IF EXISTS  "+v.table_name+";" 
    return v.Exec( script  )
    
}


func ( v  Json_Table_Driver )create_table( )bool{

   script_array := make([]string,14)
   script_array[0] = "CREATE TABLE IF NOT EXISTS  "+ v.table_name +"( "
   script_array[1] = "stream_id BIGSERIAL PRIMARY KEY,"
   script_array[2] = "Tag1 text,"
   script_array[3] = "Tag2 text,"
   script_array[4] = "Tag3 text,"
   script_array[5] = "Tag4 text,"
   script_array[6] = "Tag5 text,"
   script_array[7] = "Tag6 text,"
   script_array[8] = "Tag7 text,"
   script_array[9] = "Tag8 text,"
   script_array[10] = "Tag9 text,"
   script_array[11] = "Tag10 text,"
   script_array[12] = "Data json,"                                                                                                                                                                                                                                                                                                                                                                                                                           
   script_array[13] = "Time  bigint );"
   script := strings.Join(script_array," ")
   return v.Exec( script  )
}

func ( v  Json_Table_Driver )create_index()bool{
    script := "CREATE INDEX time_idx ON "+v.table_name+ "(Time);"
    return v.Exec( script  )
}    
    




func ( v  Json_Table_Driver )Insert(input Json_Table_Record  )bool{
    
  time_stamp    := time.Now().UnixNano()
  
     
  
 
   script := fmt.Sprintf("INSERT INTO %s "+
                        "(Tag1,Tag2,Tag3,Tag4,Tag5,Tag6,Tag7,Tag8,Tag9,Tag10,Data,Time )  "+ 
                        "VALUES('%s',  '%s',  '%s',  '%s',  '%s',    '%s',  '%s',  '%s',  '%s',  '%s',   '%s',%d)",
                        v.table_name,   input.Tag1  ,input.Tag2,   input.Tag3,  input.Tag4,   input.Tag5,
                        input.Tag6,   input.Tag7,   input.Tag8,  input.Tag9,  input.Tag10,
                          input.Data,time_stamp)
    status :=  v.Exec( script  )
 
  return status
}

/*
 *  time_time is the number of seconds in the past
 * 
 */
func ( v Json_Table_Driver )Trim( trim_time_second int64  )bool{

    current_time := time.Now().UnixNano()
    delete_time  := current_time - trim_time_second *1000000000 
    script := fmt.Sprintf("DELETE FROM %s WHERE time < %d ;",v.table_name, delete_time)
    
    return v.Exec(script)
    
}

func (v Json_Table_Driver)tags_where_clause( tags map[string]string)string{
    
   where_clause_array := make([]string,0)
   for key,value := range tags{
       entry := `( `+key+`  = '`+value+`' )`
       where_clause_array = append(where_clause_array,entry)
   }
   where_clause := strings.Join(where_clause_array," AND ")   
   where_clause = where_clause
   return where_clause
}

func ( v    Json_Table_Driver)Delete_Entry( tags map[string]string)bool{
    
    where_clause := v.tags_where_clause(tags)
    script := "DELETE FROM "+v.table_name+" where "+where_clause+";"
    //fmt.Println("delete script ",script)
    return v.Exec(script)
    
}

func ( v Json_Table_Driver )Vacuum( )bool{
    
    
    script := "VACUUM "+v.table_name +";"
    return v.Exec( script  )
    
}




func (v  Json_Table_Driver)Select_tags(tags map[string]string)([]Json_Table_Record, bool){
   where_clause := v.tags_where_clause( tags )
   
   //fmt.Println(" select clause ",where_clause)
   return v.Select_where(where_clause)
    
}


func (v Json_Table_Driver)Select_where(where_clause string)([]Json_Table_Record, bool){
    

    
    return_value := make([]Json_Table_Record,0)
    
    script := "Select * from "+v.table_name +" where "+where_clause+ ";"
    //fmt.Println("script",script)
    rows, err := v.conn.Query(context.Background(), script)
    if err != nil {
      fmt.Println("err",err)
      return return_value, false
    }
    defer rows.Close()

    for rows.Next() {
            
             var item Json_Table_Record
            rows.Scan(&item.Stream_id,&item.Tag1,&item.Tag2,&item.Tag3,&item.Tag4,&item.Tag5,&item.Tag6,&item.Tag7,&item.Tag8,&item.Tag9,&item.Tag10, &item.Data,&item.Time_stamp)
           
            if rows.Err() != nil {
            fmt.Println("rows err",rows.Err())
              return return_value,false
            }
            return_value = append(return_value,item)
              
        }
    
  
    return return_value,true
    
}
func (v Json_Table_Driver)Select_All()([]Json_Table_Record, bool){
    
    
    
    return_value := make([]Json_Table_Record,0)
    
    script := "Select * from "+v.table_name +";"
    //fmt.Println("script",script)
    rows, err := v.conn.Query(context.Background(), script)
    if err != nil {
      return return_value, false
    }
    defer rows.Close()

    for rows.Next() {
            
            var item Json_Table_Record
            rows.Scan(&item.Stream_id,&item.Tag1,&item.Tag2,&item.Tag3,&item.Tag4,&item.Tag5,&item.Tag6,&item.Tag7,&item.Tag8,&item.Tag9,&item.Tag10, &item.Data,&item.Time_stamp)
            
            if rows.Err() != nil {
              return return_value,false
            }
            return_value = append(return_value,item)
              
        }
    
  
    return return_value,true
    
}




func (v Json_Table_Driver)Select_after_time_stamp_desc_tag1(tag1 string, timestamp int64)([]Json_Table_Record, bool){
    
   
    
    return_value := make([]Json_Table_Record,0)
    
    current_time := time.Now().UnixNano()
    select_time  := current_time - timestamp *1000000000 
    script := fmt.Sprintf("Select * from %s WHERE tag1 = '%s' and time >= %d ORDER BY time DESC ;",v.table_name, tag1, select_time)
    //fmt.Println("script",script)
    rows, err := v.conn.Query(context.Background(), script)
    if err != nil {
      fmt.Println(err)
      return return_value, false
    }
    defer rows.Close()

    for rows.Next() {
            
            var item Json_Table_Record
              rows.Scan(&item.Stream_id,&item.Tag1,&item.Tag2,&item.Tag3,&item.Tag4,&item.Tag5,&item.Tag6,&item.Tag7,&item.Tag8,&item.Tag9,&item.Tag10, &item.Data,&item.Time_stamp)
            
            if rows.Err() != nil {
              return return_value,false
            }
            return_value = append(return_value,item)
              
        }
    
  
    return return_value,true
    
}

func (v Json_Table_Driver)Select_after_time_stamp_asc( timestamp int64)([]Json_Table_Record, bool){
    
  
    return_value := make([]Json_Table_Record,0)
    
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
            
            var item Json_Table_Record
            rows.Scan(&item.Stream_id,&item.Tag1,&item.Tag2,&item.Tag3,&item.Tag4,&item.Tag5,&item.Tag6,&item.Tag7,&item.Tag8,&item.Tag9,&item.Tag10, &item.Data,&item.Time_stamp)
            
            if rows.Err() != nil {
              return return_value,false
            }
            return_value = append(return_value,item)
              
        }
    
  
    return return_value,true
    
}

