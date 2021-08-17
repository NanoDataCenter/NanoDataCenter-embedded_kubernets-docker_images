package main

import (
    "fmt"
    "lacima.com/server_libraries/postgres"
    "lacima.com/site_data"
    "lacima.com/redis_support/graph_query"
    "lacima.com/redis_support/generate_handlers"


)



const start_banner string= `
**************************************************
starting test
*************************************************
`
const adhoc_banner string = `
************************************************
starting adhoc test
**************************************************
`


const managed_key  string= `
********************************************************
managed key test
***********************************************************
`


func main(){
    
    
    var site_data_store map[string]interface{}
    const config_file = "/data/redis_configuration.json"
    
    site_data_store = get_site_data.Get_site_data(config_file)
    graph_query.Graph_support_init(&site_data_store)
	
	data_handler.Data_handler_init(&site_data_store)
    data_search_list := []string{"POSTGRES_TEST:driver_test","POSTGRES_LOG:postgress_test","POSTGRES_LOG"}
	data_element := data_handler.Construct_Data_Structures(&data_search_list)

	managed_driver := (*data_element)["POSTGRES_LOG"].(pg_drv.Postgres_Stream_Driver)


    /*
     * Test Adhoc table handler
     * 
     */ 
    fmt.Println(start_banner)
    fmt.Println(adhoc_banner)
    key := "" // not used in adhoc accesses
    driver := pg_drv.Construct_Postgres_Stream_Driver(key, "admin","password","admin","stream_test",3*30*24*3600 ) 
    status := driver.Connect("localhost")
    fmt.Println("status",status)
    if status == false {
        return
    }
    test_table(driver)
    
    
    fmt.Println(managed_key)
    status = managed_driver.Connect("localhost")
    fmt.Println("status",status)
    if status == false {
        return
    }
    test_table(managed_driver)
    
}

func test_table(driver pg_drv.Postgres_Stream_Driver){
   defer driver.Close()
    //fmt.Println("drop",driver.Drop_table())
    //fmt.Println("create",driver.Create_table())
    fmt.Println("insert",driver.Insert("tag1_a","tag2_a","tag3_a","++++ data1 ..........................")) 
    fmt.Println("insert",driver.Insert("tag1_b","tag2_b","tag3_b","++++ data2 .........................."))  
    fmt.Println("insert",driver.Insert("tag1_c","tag2_c","tag3_c","+++  data3 .........................."))  
    fmt.Println("insert",driver.Insert("tag1_d","tag2_d","tag3_d","+++  data4 .........................."))  
    fmt.Println("insert",driver.Insert("tag1_e","tag2_e","tag3_e","+++  data5 .........................."))  
    fmt.Println("query")
    result,err := driver.Select_All()
    fmt.Println("err",err)
    for i,item := range result {
        fmt.Println("index",i,item )
    }
    result,err = driver.Select_after_time_stamp( 30*24*3600)
    fmt.Println("err",err)
    for i,item := range result {
        fmt.Println("index",i,item )
    }
    result,err = driver.Select_where(" tag1='tag1_a' or tag2='tag2_e' or data like '%data4%' ")
    fmt.Println("err",err)
    for i,item := range result {
        fmt.Println("index",i,item )
    }
    
    fmt.Println("trim",driver.Trim( 15  ))
    fmt.Println("vacuum",driver.Vacuum())
    fmt.Println("query")
    result,err = driver.Select_All()
    fmt.Println("err",err)
    for i,item := range result {
        fmt.Println("index",i,item )
    }
    result,err = driver.Select_where(" to_tsvector('english' , data) @@ to_tsquery('english','data5 | data1') ")
    fmt.Println("err",err)
    for i,item := range result {
        fmt.Println("index",i,item )
    }
    
    fmt.Println("drop table",driver.Drop_table())
}        
    

