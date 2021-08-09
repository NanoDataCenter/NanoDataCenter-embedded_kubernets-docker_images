package "sqlite3_table_handlers"


type SQLITE3_TEXT_TABLE_RECORD struct {
   handler_name        string;
   database_name       string;
   table_name          string;
   fields              []string;

}




func Construct_document_table_record( handler_name, database_name string, fields []string )SQLITE3_TEXT_TABLE_RECORD{
    
    var return_value  SQLITE3_TEXT_TABLE_RECORD
    
    return_value.handler_name  = handler_name
    return_value.database_name = database_name
    return_value.fields        = fields
    // create table
    
    
    return return_value
    
}
