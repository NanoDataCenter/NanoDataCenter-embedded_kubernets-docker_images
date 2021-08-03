package main

import "fmt"

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
    fmt.Println(sqlite3_handle.Vacuum("test2.db"))
	fmt.Println("done")
	
}

/*
  sqlite_client = Construct_RPC_Library(qs,redis_site)
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
    try:
       print(sqlite_client.create_database("backup"))
    except:
        print("duplicate db")
    print(sqlite_client.version())
    print(sqlite_client.get_text("test"))
    print(sqlite_client.set_text("test","bytes"))
    print(sqlite_client.get_text("test"))
    print(sqlite_client.set_text("test","string"))
    print(sqlite_client.get_text("test"))
    print(sqlite_client.backup("test","backup"))
    temp = '''create table recipe( name text, ingredients text);'''
    print(sqlite_client.ex_exec("test",temp))
    temp = """
    insert into recipe (name, ingredients) values ('broccoli stew', 'broccoli peppers cheese tomatoes');
    insert into recipe (name, ingredients) values ('pumpkin stew', 'pumpkin onions garlic celery');
    insert into recipe (name, ingredients) values ('broccoli pie', 'broccoli cheese onions flour');
    insert into recipe (name, ingredients) values ('pumpkin pie', 'pumpkin sugar flour butter');
    """
    
    print(sqlite_client.select("test","select * from recipe"))
    print(sqlite_client.select("test","select rowid,name,ingredients from recipe"))
    # intentional bad sql
    
    #print(sqlite_client.select("test","select "))
    print(sqlite_client.close_database("test"))
    print(sqlite_client.close_database("backup"))
    print(sqlite_client.delete_database("test"))
    print(sqlite_client.delete_database("backup"))
    
*/
