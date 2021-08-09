package main

import "fmt"

import "encoding/json"
import "lacima.com/site_data"
import "lacima.com/redis_support/generate_handlers"
import "lacima.com/redis_support/graph_query"
import "lacima.com/redis_support/redis_handlers"
import "lacima.com/server_libraries/sqlite3_server_library"




func main(){

    
		
	var config_file = "/data/redis_configuration.json"
	
	var site_data_store map[string]interface{}

	site_data_store = get_site_data.Get_site_data(config_file)
    graph_query.Graph_support_init(&site_data_store)
	redis_handlers.Init_Redis_Mutex()
	data_handler.Data_handler_init(&site_data_store)	
    search_list := []string{"RPC_SERVER:SQLITE3_SERVER","RPC_SERVER"}
 	sqlite3_handle := sqlite3_server_lib.Sqlite3_Server_Init(&search_list)
    fmt.Println("ping")
    fmt.Println(sqlite3_handle.Ping())
    fmt.Println("list db")
    fmt.Println(sqlite3_handle.List_databases())
    fmt.Println(" open test1.db ")
    fmt.Println(sqlite3_handle.Open_database("test1_db"))
    fmt.Println("open test2.db ")
    fmt.Println(sqlite3_handle.Open_database("test2_db"))
    fmt.Println("list db")
    fmt.Println(sqlite3_handle.List_databases())
    fmt.Println("close test1.db")
    fmt.Println(sqlite3_handle.Close_database("test1_db"))
    fmt.Println("list database")
    fmt.Println(sqlite3_handle.List_databases())
    fmt.Println("vacuum")
    fmt.Println(sqlite3_handle.Vacuum("test2_db"))
    fmt.Println(sqlite3_handle.Delete_database("test1_db"))

    fmt.Println(sqlite3_handle.Close_database("test2_db"))
     fmt.Println(sqlite3_handle.Delete_database("test2_db"))
	fmt.Println("done  first step")
    
    
    
	fmt.Println(sqlite3_handle.Open_database("test"))
    fmt.Println("list_tables")
    fmt.Println(sqlite3_handle.List_tables("test"))
    fmt.Println("create_table",sqlite3_handle.Create_table("test","test",[]string{"a text","b text"},false,true))
    fmt.Println("list_tables")
    fmt.Println(sqlite3_handle.List_tables("test"))
    
    fmt.Println("drop_table",sqlite3_handle.Drop_table("test","test"))
    fmt.Println("list tables after drop")
    fmt.Println(sqlite3_handle.List_tables("test"))
    fmt.Println("create table",sqlite3_handle.Create_table("test","test",[]string{"a text","b text"},false,true))
    fmt.Println("list tables")
    fmt.Println(sqlite3_handle.List_tables("test"))
    fmt.Println("schema")
    fmt.Println(sqlite3_handle.Get_table_schema("test","test"))
    //fmt.Println("drop_table",sqlite3_handle.Drop_table("test","test"))
    fmt.Println("list tables after drop")
    
    
    fmt.Println(sqlite3_handle.Create_text_search_table("test","text",[]string{"a","b"}))
    fmt.Println("list tables")
    fmt.Println(sqlite3_handle.List_tables("test"))
    
    
    field_values := make([][]string,6)
    field_values[0] = []string{".1", "'broccoli peppers cheese tomatoes'"}
    field_values[1] = []string{ ".15", "'test data'"}
    field_values[2] = []string{".2", "'pumpkin onions garlic celery'"}
    field_values[3] = []string{".3", "'broccoli cheese onions flour'"}
    field_values[4] = []string{".3", "'duplicate value'"}
    field_values[5] = []string{".4", "'pumpkin sugar flour butter'"}
    fmt.Println("insert",sqlite3_handle.Insert_entries("test","text",[]string{"a","b"},field_values))
    
    fmt.Println("text search example")
    fmt.Println(sqlite3_handle.Text_search("test","text", "pumpkin OR celery"))
    
    
    fmt.Println("select")
    fmt.Println(sqlite3_handle.Select("test","text",[]string{"a","b"},false,"",false))
    fmt.Println("selecy")
    fmt.Println(sqlite3_handle.Select("test","text",[]string{"a","b"},true,"a < .2",false))
    fmt.Println("select")
    fmt.Println(sqlite3_handle.Select("test","text",[]string{"a"},true,"a >= .2",false))

    fmt.Println(sqlite3_handle.Delete_entry("test","text","a <= .2"))
    fmt.Println("select")
    fmt.Println(sqlite3_handle.Select("test","text",[]string{"a","b"},false,"",false))
    /*
    #
    #
    #  Now testing alter command
    #
    #
    
    
    #
    #
    # Testing Update Command
    #
    
    */
    fmt.Println("drop_table",sqlite3_handle.Drop_table("test","test"))
    fmt.Println("create table",sqlite3_handle.Create_table("test","test",[]string{"a text","b text"},false,true))
    
    
    field_values = make([][]string,6)
    field_values[0] = []string{".1", "'broccoli peppers cheese tomatoes'"}
    field_values[1] = []string{ ".15", "'test data'"}
    field_values[2] = []string{".2", "'pumpkin onions garlic celery'"}
    field_values[3] = []string{".3", "'broccoli cheese onions flour'"}
    field_values[4] = []string{".3", "'duplicate value'"}
    field_values[5] = []string{".4", "'pumpkin sugar flour butter'"}
    fmt.Println("insert",sqlite3_handle.Insert_entries("test","test",[]string{"a","b"},field_values))
    
    fmt.Println("update entry ",sqlite3_handle.Update("test","test", []string{"b"},[]string{"'new value'"},true,"a>.2"))

    fmt.Println("select ")
    fmt.Println(sqlite3_handle.Select("test","test",[]string{},false,"",false))
    
    fmt.Println("update ",sqlite3_handle.Update("test","test",[]string{"b"},[]string{"'new default'"},false,""))
    fmt.Println("select")
    fmt.Println(sqlite3_handle.Select("test","test",[]string{},false,"",false))

    

    fmt.Println("alter_table_rename",sqlite3_handle.Alter_table_rename("test","text","new_text" ))
    fmt.Println("tables")
    fmt.Println(sqlite3_handle.List_tables("test"))
    
    fmt.Println("insert",sqlite3_handle.Insert_entries("test","test",[]string{"a","b"},field_values))
    fmt.Println("select")
    fmt.Println(sqlite3_handle.Select("test","test",[]string{"a","b"},false,"",false))
    fmt.Println("alter_table_add_column",sqlite3_handle.Alter_table_add_column("test","test","c text"    ))
    fmt.Println("schema")
    fmt.Println(sqlite3_handle.Get_table_schema("test","test"))
    field_values = make([][]string,6)
    field_values[0] = []string{".1", "'broccoli peppers cheese tomatoes'","'c0'"}
    field_values[1] = []string{ ".15", "'test data'","'c1'"}
    field_values[2] = []string{".2", "'pumpkin onions garlic celery'","'c2'"}
    field_values[3] = []string{".3", "'broccoli cheese onions flour'","'c3'"}
    field_values[4] = []string{".3", "'duplicate value'","'c4'"}
    field_values[5] = []string{".4", "'pumpkin sugar flour butter'","'c5'"}
    fmt.Println("insert",sqlite3_handle.Insert_entries("test","test",[]string{"a","b","c"},field_values))
    fmt.Println("select")
    fmt.Println(sqlite3_handle.Select("test","test",[]string{},false,"",false))   
   
    
    /*
     * Testing JSON FUNCTIONS
     * 
     */
    
        
    fmt.Println("create JSON Table",sqlite3_handle.Create_table("test","json_test",[]string{"a text","b JSON"},false,true))
    fmt.Println("schema")
    fmt.Println(sqlite3_handle.Get_table_schema("test","json_test"))
    
    field_values = make([][]string,2)
    
    insert_data_dict      := make(map[string]interface{})
    insert_data_dict["test_string"] = "test_string_result"
    insert_data_dict["test_float"] = 3.14159
    insert_data_json,_ := json.Marshal(insert_data_dict)
    field_values[0] = []string{"'test_map'","'"+string(insert_data_json)+"'"}
    
    insert_data_dict      = make(map[string]interface{})
    insert_data_dict["test_string"] = "test_string_result"
    insert_data_dict["test_float"] = 19.555555
    insert_data_json,_ = json.Marshal(insert_data_dict)
    field_values[1] = []string{"'test_map'","'"+string(insert_data_json)+"'"}
/*    
    insert_data_list   := make([]interface{},2)
    insert_data_list[0] = "test_string_result"
    insert_data_list[1] = 3.14159 
    insert_data_json1,_ := json.Marshal(insert_data_list)
    field_values[1] = []string{"'test_array'","'"+string(insert_data_json1)+"'"}   
*/
    fmt.Println("insert",sqlite3_handle.Insert_entries("test","json_test",[]string{"a","b"},field_values))
    fmt.Println("json query")
    err , json_results := sqlite3_handle.Query("test","select json_extract(json_test.b  ,'$.test_float') from json_test")
    fmt.Println("err",err)
    for _,i := range json_results{
        fmt.Println("resultS",i)
    }
    
    
    fmt.Println("drop_table",sqlite3_handle.Drop_table("test","json_test"))
    fmt.Println("drop_table",sqlite3_handle.Drop_table("test","test"))
    fmt.Println("drop_table",sqlite3_handle.Drop_table("test","text"))
    fmt.Println("list tables after drop")
    sqlite3_handle.Close_database("test") 
    sqlite3_handle.Delete_database("test")
    sqlite3_handle.Delete_database("test")

}
//select json_extract(user.phone, '$.cell') from user;

/*
 print(sqlite_client.list_data_bases())
    try:
       print(sqlite_client.create_database("test"))
    except:
        print("duplicate db")
    print(sqlite_client.list_data_bases())
    print(sqlite_client.close_database("test"))
    print(sqlite_client.delete_database("test"))
    
    print(sqlite_client.list_data_bases())
    #os.system("ls /sqlite")
    try:
        print(sqlite_client.create_database("test"))
    except:
        print("duplicate db")
 
    print("tables",sqlite_client.list_tables("test"))
    print(sqlite_client.create_table("test","test",["a text","b text"],temp_table=False,not_exists=True))
    print("tables",sqlite_client.list_tables("test"))
    print("drop_table",sqlite_client.drop_table("test","test"))
    print("tables",sqlite_client.list_tables("test"))
    print(sqlite_client.create_table("test","test",["a text","b text"],temp_table=False,not_exists=True))
    print("tables",sqlite_client.list_tables("test"))
    print("schema",sqlite_client.get_table_schema("test","test"))
    print(sqlite_client.create_text_search_table("test","text",["a","b"]))
    print("tables",sqlite_client.list_tables("test"))
    print("schema",sqlite_client.get_table_schema("test","text"))
    field_values = [
    [.1, 'broccoli peppers cheese tomatoes'],
    [.15, "test data"],
    [.2, 'pumpkin onions garlic celery'],
    [.3, 'broccoli cheese onions flour'],
    [.3, 'duplicate value'],
    [.4, 'pumpkin sugar flour butter']]
    
    print("insert_composite",sqlite_client.insert_composite("test","text",["a","b"],field_values))
    print("select_composite",sqlite_client.select_composite("test","text",["a","b"]))
    print("select_composite",sqlite_client.select_composite("test","text",["a","b"],"a < .2"))
    print("select_composite",sqlite_client.select_composite("test","text",["a"],"a >= .2",True))

    print(sqlite_client.delete("test","text","a <= .2"))
    print("select_composite",sqlite_client.select_composite("test","text",["a","b"]))
    
    #
    #
    #  Now testing alter command
    #
    #
    print("tables",sqlite_client.list_tables("test"))
    print("alter_table_rename",sqlite_client.alter_table_rename("test","text","new_text" ))
    print("tables",sqlite_client.list_tables("test"))
    print(sqlite_client.create_table("test","test",["a text","b text"]))
    print("insert_composite",sqlite_client.insert_composite("test","test",["a","b"],field_values))
    print("select_composite",sqlite_client.select_composite("test","test",["a","b"]))
    print("alter_table_add_column",sqlite_client.alter_table_add_column("test","test","c text"    ))
    print("schema",sqlite_client.get_table_schema("test","test"))
    print("select_composite",sqlite_client.select_composite("test","test",["a","b","c"]))
    
    #
    #
    # Testing Update Command
    #
    #
    print("update ",sqlite_client.update("test","test",["c"],["default_value"],where_clause="a>.2"))
    print("select_composite",sqlite_client.select_composite("test","test",["a","b","c"]))
    print("update ",sqlite_client.update("test","test",["c"],["new default"]))
    print("select_composite",sqlite_client.select_composite("test","test",["a","b","c"]))

    
    print(sqlite_client.close_database("test")) 
    print(sqlite_client.delete_database("test"))
   
    
*/
