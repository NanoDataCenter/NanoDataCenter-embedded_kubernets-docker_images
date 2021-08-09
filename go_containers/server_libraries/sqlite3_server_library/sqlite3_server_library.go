
package sqlite3_server_lib


import "fmt"
//import "reflect"
import "strings"


func (v Sqlite3_Server_Client_Type)Vacuum( db_name string )bool{
  

       parameters := make(map[string]interface{})
       parameters["database"] = db_name
       result := v.driver.Send_json_rpc_message("vacuum", parameters ) 
       return result["status"].(bool)

}
/*
SELECT * 
FROM posts 
WHERE posts MATCH 'text' 
ORDER BY rank;
*/
func (v Sqlite3_Server_Client_Type)Text_search(db_name,table_name string, match_clause string)(bool,[]map[string]interface{}){
    
    
       script := "select  *  FROM "+table_name+" WHERE "+table_name+" MATCH '"+match_clause+"' ORDER BY rank ;"
       fmt.Println("text search script "+script)
       
       return v.Query(db_name,script)
}


func (v Sqlite3_Server_Client_Type)Select(db_name,table_name string, return_fields []string ,where_flag bool , where_clause string,distinct_flag bool)(bool,[]map[string]interface{}){
       script := ""
       if distinct_flag == true{
          script = "select distinct "
       }else{
          script = "select "
       }
       if len(return_fields) == 0 {
           script = script+"  *  FROM "+table_name+" "
       }else{
           script = script+"  "+strings.Join(return_fields,",") + "  FROM "+table_name+" "
       }
       if where_flag == true {
           script = script+ "  WHERE "+where_clause+" ; "
       }else{
           script = script+" ; "
       }
       
       return v.Query(db_name,script)
}


func (v Sqlite3_Server_Client_Type)List_tables( db_name string )(bool, []string){
  

       parameters := make(map[string]interface{})
       parameters["database"]  = db_name
       parameters = v.driver.Send_json_rpc_message("list_tables", parameters )
       
       return_value := make([]string,0)
 
       
       for _,element := range parameters["results"].([]interface{}){
        value_array := element.(map[string]interface{})
        value       := value_array["name"].(string)
        return_value = append(return_value,value)   
       }
      
    
       
       return parameters["status"].(bool),return_value

}

/*
Sqlite Native Table Types are
NULL        Stored as in no conversion is done.  
INTEGER     The value is a signed integer, stored in 1, 2, 3, 4, 6, or 8 bytes depending on the magnitude of the value.
REAL        The value is a floating point value, stored as an 8-byte IEEE floating point number.
TEXT        The value is a text string, stored using the database encoding (UTF-8, UTF-16BE or UTF-16LE).
BLOB        The value is a blob of data, stored exactly as it was input.

field defininations are of the form
"a TEXT"





*/

func (v Sqlite3_Server_Client_Type)Create_table( db_name, table_name string, fields []string, temp_table, not_exists bool )(bool){
  

       parameters := make(map[string]interface{})
       parameters["database"]  = db_name
       script := "create  "
       if temp_table != false {
          script=script+" TEMP "
       }
       script = script+" TABLE "
       if not_exists == true {
          script = script+" IF NOT EXISTS "
       }
       script = script+table_name+" "
       script = script+"( " +strings.Join(fields," ,  ")+" ); "
       fmt.Println(script)
       parameters["script"] = script 
       
       
       
       result := v.driver.Send_json_rpc_message("execute", parameters )
       return result["status"].(bool)

}

/* 
 Text search tables have no types or primary keys
 Just specify the column names
 SQLite will automatically add the rowid field
*/
   
func (v Sqlite3_Server_Client_Type)Create_text_search_table( db_name, table_name string, fields []string )(bool){    
 
       parameters := make(map[string]interface{})
       parameters["database"]  = db_name
       script := "create virtual table  "
       script = script+table_name+" "
       script = script+" using fts5("+strings.Join(fields,",")+",tokenize = 'porter unicode61 remove_diacritics 1'  );"
       print("script",script)
    
       parameters["script"] = script 
       parameters  = v.driver.Send_json_rpc_message("execute", parameters )
       return parameters["status"].(bool)

}
    
    
       
func (v Sqlite3_Server_Client_Type)Drop_table( db_name, table_name string )(bool){
  
      script := "DROP TABLE if exists " +table_name+" ;" 
      return v.Execute(db_name,script)
    
}



func (v Sqlite3_Server_Client_Type)Get_table_schema( db_name, table_name string )(bool,map[string]string,  string){
  
      script := "PRAGMA table_info('"+table_name+"');"
      raw_schema := ""
      return_value := make(map[string]string)
      status,query_data := v.Query(db_name,script)
      fmt.Println("Table_schema",status,query_data)
      if status == true{
          raw_schema = fmt.Sprintf( "%v", query_data)
          for _,element := range query_data {
              name                 :=  element["name"].(string)
              return_value[name]   = element["type"].(string)
          }
      }  
      fmt.Println("status",status)
      
      return status,return_value, raw_schema

}

func (v Sqlite3_Server_Client_Type)Alter_table_add_column( db_name,table_name,new_column  string )(bool){
  
      script := "ALTER TABLE "+table_name+" ADD COLUMN "+new_column+";"  
      return v.Execute(db_name,script)
      
}
  
       
func (v Sqlite3_Server_Client_Type)Alter_table_rename(db_name, old_table, new_table string )(bool){
  
      script := "ALTER TABLE "+old_table+" RENAME TO "+new_table;  
      return v.Execute(db_name,script)
    
}
   

func (v Sqlite3_Server_Client_Type)Delete_entry(db_name, table_name , where_clause string )(bool){
      
       script := "DELETE FROM "+table_name+" WHERE "+where_clause+";"
       return v.Execute(db_name,script)
}


/*
 * Note this function does multiple inserts per call
 * 
 * Entries are text -- it is up to the caller to convert float64 or integer or blob formats to string equivalents
 *
 */ 

func (v Sqlite3_Server_Client_Type)Insert_entries(db_name,table_name string,row_names[]string,row_values[][]string  )(bool){

      script_array := make([]string,0)
      for _,i := range row_values {
        
          script_array = append(script_array, "INSERT INTO "+table_name+" ("+ strings.Join(row_names,",")+" )  VALUES("+strings.Join(i,",")+")")
          
      }
      script := strings.Join(script_array,";\n")
      fmt.Println("script",script)
      return v.Execute(db_name,script)
    
}
       
       
func (v Sqlite3_Server_Client_Type)Update(db_name,table_name string,row_names,row_values[]string , where_flag bool,  where_clause string  )(bool){ 

    if len(row_names) != len(row_values){
          panic("row id and row values are not same length")    
    }      
    script := "UPDATE "+table_name+" SET  "      
    length := len(row_names)
    for i:=0;i<length;i++ {
        if i != length-1 {
            script = script +" "+row_names[i] +" = "+row_values[i]+", "
        }else{
            script = script +" "+row_names[i] +" = "+row_values[i]+" "
        }
    }
    if where_flag == true{
        script = script + "WHERE  "+where_clause+ " ; "
    }else{
           script = script+ " ;"
    }
    fmt.Println("script",script)
    return v.Execute(db_name,script)
}
    
