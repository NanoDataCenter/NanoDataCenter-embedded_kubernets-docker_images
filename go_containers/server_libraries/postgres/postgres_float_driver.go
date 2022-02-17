package pg_drv


import (
    
    "fmt"   
    "strings"
    //"strconv"
    "time"
	"context"
	//"github.com/jackc/pgx/v4"   
	
)
    

type Float_Output_Data_Record struct {
 
    Stream_id  int64;
    Text1        string;  
    Text2        string;  
    Text3        string;  
    Text4        string;  
    Text5        string;
    Text6        string;  
    Text7        string;  
    Text8        string;  
    Text9        string;  
    Text10       string;
    
    Float1       float64;  
    Float2       float64;  
    Float3       float64;   
    Float4       float64;  
    Float5       float64;  
    Float6       float64;  
    Float7       float64;  
    Float8       float64;   
    Float9       float64;  
    Float10       float64;  

    Data       string;
    Time_stamp int64;
}    



type Postgres_Float_Driver struct {
    
     Postgres_Basic_Driver
     key           string
     user          string
     password      string
     database      string
     table_name    string
     
    
}




func Construct_Postgres_Float_Driver( key,user,password,database, table_name string) Postgres_Float_Driver{
    var return_value Postgres_Float_Driver
    return_value.key            = key
    return_value.user           = user
    return_value.password       = password
    return_value.database       = database
    return_value.table_name     = table_name
    
    
    return return_value
}

func ( v  Postgres_Float_Driver )Vacuum( )bool{
    
    
    script := "VACUUM "+v.table_name +";"
    return v.Exec( script  )
    
}

func ( v  *Postgres_Float_Driver )Connect( ip string )bool{
    connection_url := "postgres://"+v.user+":" + v.password + "@"+ ip+":5432" + "/"+v.database 
    if v.connect(connection_url) == false {
        return false
    }
    
	
	

	v.create_table()
    
	
	return true
}


func ( v  Postgres_Float_Driver )Create_table()bool{
    v.create_table()
    
    return true
}


func ( v  Postgres_Float_Driver )Drop_table(  )bool{
    script := "DROP TABLE IF EXISTS  "+v.table_name+";" 
    return v.Exec( script  )
    
}


func ( v  Postgres_Float_Driver )create_table( )bool{
    script_array := make([]string,24)
    script_array[0] = "CREATE TABLE IF NOT EXISTS  "+ v.table_name +"( "
    script_array[1] = "stream_id BIGSERIAL PRIMARY KEY,"
   
    script_array[2]  = "Text1         text,"
    script_array[3]  = "Text2         text,"
    script_array[4]  = "Text3         text,"
    script_array[5]  = "Text4         text,"
    script_array[6]  = "Text5         text,"
    script_array[7]  = "Text6         text," 
    script_array[8]  = "Text7         text,"
    script_array[9]  = "Text8         text," 
    script_array[10] = "Text9         text," 
    script_array[11] = "Text10        text,"
    
    script_array[12] = "Float1        double precision,"  
    script_array[13] = "Float2        double precision,"  
    script_array[14] = "Float3        double precision,"   
    script_array[15] = "Float4        double precision,"  
    script_array[16] = "Float5        double precision," 
    script_array[17] = "Float6        double precision," 
    script_array[18] = "Float7        double precision,"  
    script_array[19] = "Float8        double precision,"   
    script_array[20] = "Float9        double precision,"  
    script_array[21] = "Float10       double precision," 
    script_array[22] = "Data text,"
    script_array[23] = "Time bigint );"
    script := strings.Join(script_array," ")
   return v.Exec( script  )
}


func ( v Postgres_Float_Driver)Exists(tags map[string]string)(int,bool){

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


func ( v  Postgres_Float_Driver )Insert(input Float_Output_Data_Record )bool{
    
  time_stamp    := time.Now().UnixNano()
  
  script := fmt.Sprintf("INSERT INTO %s "+
                        "(Text1,Text2,Text3,Text4,Text5,Text6,Text7,Text8,Text9,Text10,Float1,Float2,Float3,Float4,Float5,Float6,Float7,Float8,Float9,Float10,Data,Time )  "+ 
                        "VALUES('%s','%s','%s','%s','%s',  '%s','%s','%s','%s','%s', %f, %f , %f . %f , %f,  %f, %f , %f . %f , %f,'%s',%d)",
                        v.table_name,input.Text1,input.Text2,input.Text3,input.Text4,input.Text5,input.Text6,input.Text7,input.Text8,input.Text9,input.Text10,
                        input.Float1,input.Float2,input.Float3,input.Float4,input.Float5,input.Float6,input.Float7,input.Float8,input.Float9,input.Float10,
                        input.Data,time_stamp)
  
  fmt.Println("script",script)
    status :=  v.Exec( script  )
 
  return status
}



func ( v   Postgres_Float_Driver)Delete_Entry( tags map[string]string)bool{
    
    where_clause := v.tags_where_clause(tags)
    script := "DELETE FROM "+v.table_name+" where "+where_clause+";"
    fmt.Println("delete script ",script)
    return v.Exec(script)
    
}



func (v Postgres_Float_Driver)tags_where_clause( tags map[string]string)string{
    
   where_clause_array := make([]string,0)
   for key,value := range tags{
       entry := `( `+key+`  = '`+value+`' )`
       where_clause_array = append(where_clause_array,entry)
   }
   where_clause := strings.Join(where_clause_array," AND ")   
   where_clause = where_clause
   return where_clause
}

func (v  Postgres_Float_Driver)Select_tags(tags map[string]string)([]Float_Output_Data_Record, bool){
   where_clause := v.tags_where_clause( tags )
   return v.Select_where(where_clause)
    
}

func (v  Postgres_Float_Driver)Select_where(where_clause string)([]Float_Output_Data_Record, bool){
    
    
    
    
    
    script := "Select * from "+v.table_name +" where "+where_clause+";"
    return v.retreive_data(script )
    
}
func (v  Postgres_Float_Driver)Select_All()([]Float_Output_Data_Record , bool){
    
    script := "Select * from "+v.table_name +";"
    return v.retreive_data(script )
    
}

func (v  Postgres_Float_Driver)retreive_data(script string)([]Float_Output_Data_Record , bool){
    
    return_value := make([]Float_Output_Data_Record,0)
    
     rows, err := v.conn.Query(context.Background(), script)
    if err != nil {
      return return_value, false
    }
    defer rows.Close()

    for rows.Next() {
            
            var item Float_Output_Data_Record
            rows.Scan(&item.Stream_id,&item.Text1,&item.Text2,&item.Text3,&item.Text4,&item.Text5,&item.Text6,&item.Text7,&item.Text8,&item.Text9,&item.Text10,
                      &item.Float1,&item.Float2,&item.Float3,&item.Float4,&item.Float5,&item.Float6,&item.Float7,&item.Float8,&item.Float9,&item.Float10,
                      &item.Data,&item.Time_stamp)
            
            if rows.Err() != nil {
              return return_value,false
            }
            return_value = append(return_value,item)
              
        }
    
  
    return return_value,true
    
    
    
}

/*
 * No update right now

func ( v  Postgres_Float_Driver )Update( tag1,tag2,tag3,tag4,tag5,data string )bool{
    
  time_stamp    := time.Now().UnixNano()
  
     
  
   update_part := fmt.Sprint(" SET tag1 ='%s',tag2 ='%s',tag3 ='%s',tag4 ='%s',tag5 ='%s', data ='%s',time =%d ; " ,tag1,tag2,tag3,tag4,tag5,data,time_stamp)
   script := "UPDATE INTO "+v.table_name+update_part

  status :=  v.Exec( script  )
 
  return status
}

*/
