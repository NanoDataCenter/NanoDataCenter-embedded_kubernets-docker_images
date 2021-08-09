package "sqlite3_table_handlers"





type SQLITE3_TABLE_RECORD struct {
   handler_name        string;
   database_name       string;
   table_name          string;
   fields              []string;

}
