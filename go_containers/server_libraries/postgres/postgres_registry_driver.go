package pg_drv


import (
    
    "fmt"   
    "strings"
    //"strconv"
    
	"context"
	//"github.com/jackc/pgx/v4"   
	
)


type Registry_Record struct {
 
    stream_id  int64;
    name       string;
    key        string;
    properties string;  // json data
    path       string;
    
}    

type Registry_Driver struct {
    
     Postgres_Basic_Driver
     key           string
     user          string
     password      string
     database      string
     table_name    string
     
    
}




func Construct_Registry_Driver( key,user,password,database, table_name string  ) Registry_Driver{
    var return_value Registry_Driver
    return_value.key            = key
    return_value.user           = user
    return_value.password       = password
    return_value.database       = database
    return_value.table_name     = table_name
    
    return return_value
}


func ( v  *Registry_Driver )Connect( ip string )bool{
    connection_url := "postgres://"+v.user+":" + v.password + "@"+ ip+":5432" + "/"+v.database 
    if v.connect(connection_url) == false {
        return false
    }
    
	v.register_extension()
    return true

}

func ( v Registry_Driver )register_extension()bool{
   script := "CREATE EXTENSION ltree;"
   return_value := v.Exec(script)
   
   return return_value
}

func ( v  Registry_Driver )Drop_table(  )bool{
    script := "DROP TABLE IF EXISTS  "+v.table_name+";" 
    return v.Exec( script  )
    
}

func (v Registry_Driver)Create_table() bool {
    if v.create_table() == false {
        fmt.Println("create_table is false")
        return false
    }
    return v.create_index()
}
    

    
    
func ( v  Registry_Driver )create_table( )bool{
   script_array := make([]string,6)
   script_array[0] = "CREATE TABLE IF NOT EXISTS  "+ v.table_name +"( "
   script_array[1] = "stream_id BIGSERIAL PRIMARY KEY,"
   script_array[2] = "name text,"
   script_array[3] = "key  text,"
   script_array[4] = "properties text,"
   script_array[5] = "path ltree );"
   script := strings.Join(script_array," ")
   
   return v.Exec( script  )
}




func ( v  Registry_Driver )create_index()bool{
    script := "CREATE INDEX path_gist_"+v.table_name+"_idx ON "+v.table_name+" USING GIST(path);"
    if v.Exec( script  ) == false {
        fmt.Println("first index is false")
        return false
    }
    script = "CREATE INDEX path_"+v.table_name+"comments_idx ON "+ v.table_name+  " USING btree(path);"
    return v.Exec( script  )
}    



func ( v  Registry_Driver )Insert( name,key,properties,path string )bool{
    

 
  script := fmt.Sprintf("INSERT INTO %s (name,key,properties,path ) VALUES('%s','%s','%s','%s');",v.table_name,name,key,properties,path)
  
  return v.Exec( script  )
}    


func (v  Registry_Driver)Select_All()([]Registry_Record, bool){
    
    return_value := make([]Registry_Record,0)
    
    script := "Select * from "+v.table_name +";"

    rows, err := v.conn.Query(context.Background(), script)
    if err != nil {
      return return_value, false
    }
    defer rows.Close()

    for rows.Next() {
            
            var item Registry_Record
            rows.Scan(&item.stream_id,&item.name,&item.key,&item.properties,&item.path)
            if rows.Err() != nil {
              return return_value,false
            }
            return_value = append(return_value,item)
              
        }
    
  
    return return_value,true
    
}


func (v  Registry_Driver)Select_where(where_clause string)([]Registry_Record, bool){
    
    return_value := make([]Registry_Record,0)
    
    script := "Select * from "+v.table_name +" where "+where_clause+ ";"
    
    rows, err := v.conn.Query(context.Background(), script)
    if err != nil {
      return return_value, false
    }
    defer rows.Close()

    for rows.Next() {
            
            var item Registry_Record
            rows.Scan(&item.stream_id,&item.name,&item.key,&item.properties,&item.path)
            if rows.Err() != nil {
              return return_value,false
            }
            return_value = append(return_value,item)
              
        }
    
  
    return return_value,true
    
}
