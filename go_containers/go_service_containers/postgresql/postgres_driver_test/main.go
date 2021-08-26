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
const adhoc_registry  string= `
********************************************************
ad_hoc registry test
***********************************************************
`
const managed_registry  string= `
********************************************************
managed registry test
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
    //status = managed_driver.Connect("localhost")
    //fmt.Println("status",status)
    //if status == false {
    //    return
    //}
    test_table(managed_driver)
    
    fmt.Println(adhoc_registry)
    test_object_registry()
    
    fmt.Println(managed_registry)
    test_manage_object_registry()
}

func test_table(driver pg_drv.Postgres_Stream_Driver){
   defer driver.Close()
    fmt.Println("drop",driver.Drop_table())
    fmt.Println("create",driver.Create_table())
    fmt.Println("insert",driver.Insert("tag1_a","tag2_a","tag3_a","tag4_a","tag5_a","++++ data1 ..........................")) 
    fmt.Println("insert",driver.Insert("tag1_b","tag2_b","tag3_b","tag4_a","tag5_a","++++ data2 .........................."))  
    fmt.Println("insert",driver.Insert("tag1_c","tag2_c","tag3_c","tag4_a","tag5_a","+++  data3 .........................."))  
    fmt.Println("insert",driver.Insert("tag1_d","tag2_d","tag3_d","tag4_a","tag5_a","+++  data4 .........................."))  
    fmt.Println("insert",driver.Insert("tag1_e","tag2_e","tag3_e","tag4_a","tag5_a","+++  data5 .........................."))  
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
    
func test_object_registry(){
    
    
 fmt.Println("\n\n**************************** testingn testing object_registry")
  driver := pg_drv.Construct_Registry_Driver( "","admin","password","admin", "registry_test")
  fmt.Println(driver.Connect("localhost"))
  fmt.Println("drop table",driver.Drop_table())
  defer driver.Close()
  fmt.Println("create table",driver.Create_table())
    
  fmt.Println("insert", driver.Insert( "name1","key1","properties_1","Top"))
  fmt.Println("insert", driver.Insert( "name2","key2","properties_2","Top.Science"))
  fmt.Println("insert", driver.Insert( "name3","key3","properties_3","Top.Science.Astronomy"))
  fmt.Println("insert", driver.Insert( "name4","key4","properties_4","Top.Science.Astronomy.Astrophysics"))
  fmt.Println("insert", driver.Insert( "name5","key5","properties_5", "Top.Science.Astronomy.Cosmology"))
  fmt.Println("insert", driver.Insert( "name6","key6","properties_6","Top.Hobbies"))
  fmt.Println("insert", driver.Insert( "name7","key7","properties_7","Top.Hobbies.Amateurs_Astronomy"))
  fmt.Println("insert", driver.Insert( "name8","key8","properties_8","Top.Collections"))
  fmt.Println("insert", driver.Insert( "name9","key9","properties_9","Top.Collections.Pictures"))
  fmt.Println("insert", driver.Insert( "name10","key10","properties_10","Top.Collections.Pictures.Astronomy"))
  fmt.Println("insert", driver.Insert( "name11","key11","properties_11", "Top.Collections.Pictures.Astronomy.Stars"))
  fmt.Println("insert", driver.Insert( "name12","key12","properties_12","Top.Collections.Pictures.Astronomy.Galaxies"))
  fmt.Println("insert", driver.Insert( "name14","key14","properties_14","Top.Collections.Pictures.Astronomy.Astronauts"))
  fmt.Println("insert", driver.Insert( "name15","key15","properties_15","Top.Collections.Pictures.Astronomy.Astronauts.Flight_Engineer.left_seat"))
  
  fmt.Println("select all")
  dump_records(driver.Select_All())

  
  fmt.Println("select where  path <@ 'Top.Science'")
  dump_records(driver.Select_where("path <@ 'Top.Science'"))
  
  fmt.Println("select where  path ~ '*.Astronomy.*'")
  dump_records(driver.Select_where("path ~ '*.Astronomy.*'"))

  fmt.Println("select where path ~ '*.!pictures@.Astronomy.*'")
  dump_records(driver.Select_where("path ~ '*.!pictures@.Astronomy.*'"))
  
  fmt.Println( " full text search")
  
  fmt.Println("select where path @ 'Astro*% & !pictures@'")
  dump_records(driver.Select_where("path @ 'Astro*% & !pictures@'"))
  
  fmt.Println("select where  path @ 'Astro* & !pictures@'")
  dump_records(driver.Select_where("path @ 'Astro* & !pictures@'"))
  
  
  fmt.Println("registry specific search")
  fmt.Println("select where path ~ 'Top.Collections.*.Astronomy.*.Flight_Engineer.left_seat'")
  dump_records(driver.Select_where("path ~ 'Top.Collections.*.Astronomy.*.Flight_Engineer.left_seat'"))

  
  
  
 
 
}

func test_manage_object_registry(){
    
 
   data_search_list := []string{"POSTGRES_REGISTY_TEST"}
   data_element := data_handler.Construct_Data_Structures(&data_search_list)

   driver := (*data_element)["postgress_registry_test"].(pg_drv.Registry_Driver)   
    
    
  fmt.Println("drop table",driver.Drop_table())
  defer driver.Close()
  fmt.Println("create table",driver.Create_table())
    
  fmt.Println("insert", driver.Insert( "name1","key1","properties_1","Top"))
  fmt.Println("insert", driver.Insert( "name2","key2","properties_2","Top.Science"))
  fmt.Println("insert", driver.Insert( "name3","key3","properties_3","Top.Science.Astronomy"))
  fmt.Println("insert", driver.Insert( "name4","key4","properties_4","Top.Science.Astronomy.Astrophysics"))
  fmt.Println("insert", driver.Insert( "name5","key5","properties_5", "Top.Science.Astronomy.Cosmology"))
  fmt.Println("insert", driver.Insert( "name6","key6","properties_6","Top.Hobbies"))
  fmt.Println("insert", driver.Insert( "name7","key7","properties_7","Top.Hobbies.Amateurs_Astronomy"))
  fmt.Println("insert", driver.Insert( "name8","key8","properties_8","Top.Collections"))
  fmt.Println("insert", driver.Insert( "name9","key9","properties_9","Top.Collections.Pictures"))
  fmt.Println("insert", driver.Insert( "name10","key10","properties_10","Top.Collections.Pictures.Astronomy"))
  fmt.Println("insert", driver.Insert( "name11","key11","properties_11", "Top.Collections.Pictures.Astronomy.Stars"))
  fmt.Println("insert", driver.Insert( "name12","key12","properties_12","Top.Collections.Pictures.Astronomy.Galaxies"))
  fmt.Println("insert", driver.Insert( "name14","key14","properties_14","Top.Collections.Pictures.Astronomy.Astronauts"))
  fmt.Println("insert", driver.Insert( "name15","key15","properties_15","Top.Collections.Pictures.Astronomy.Astronauts.Flight_Engineer.left_seat"))
  
  fmt.Println("select all")
  dump_records(driver.Select_All())

  
  fmt.Println("select where  path <@ 'Top.Science'")
  dump_records(driver.Select_where("path <@ 'Top.Science'"))
  
  fmt.Println("select where  path ~ '*.Astronomy.*'")
  dump_records(driver.Select_where("path ~ '*.Astronomy.*'"))

  fmt.Println("select where path ~ '*.!pictures@.Astronomy.*'")
  dump_records(driver.Select_where("path ~ '*.!pictures@.Astronomy.*'"))
  
  fmt.Println( " full text search")
  
  fmt.Println("select where path @ 'Astro*% & !pictures@'")
  dump_records(driver.Select_where("path @ 'Astro*% & !pictures@'"))
  
  fmt.Println("select where  path @ 'Astro* & !pictures@'")
  dump_records(driver.Select_where("path @ 'Astro* & !pictures@'"))
  
  
  fmt.Println("registry specific search")
  fmt.Println("select where path ~ 'Top.Collections.*.Astronomy.*.Flight_Engineer.left_seat'")
  dump_records(driver.Select_where("path ~ 'Top.Collections.*.Astronomy.*.Flight_Engineer.left_seat'"))

  
  
      
    
    
    
    
    
}

func dump_records( input_data []pg_drv.Registry_Record , status bool){
    for _,element := range input_data{
        fmt.Println(element)
    }
}

/*

ltree stores a label path.

lquery represents a regular-expression-like pattern for matching ltree values. A simple word matches that label within a path. A star symbol 


(*) matches zero or more labels. T

foo         Match the exact label path foo
*.foo.*     Match any label path containing the label foo
*.foo       Match any label path whose last label is foo

Both star symbols and simple words can be quantified to restrict how many labels they can match:
*{n}        Match exactly n labels
*{n,}       Match at least n labels
*{n,m}      Match at least n but not more than m labels
*{,m}       Match at most m labels — same as *{0,m}
foo{n,m}    Match at least n but not more than m occurrences of foo
foo{,}      Match any number of occurrences of foo, including zero


In the absence of any explicit quantifier, the default for a star symbol is to match any number of labels (that is, {,}) while the default for a non-star item is to match exactly once (that is, {1}).

There are several modifiers that can be put at the end of a non-star lquery item to make it match more than just the exact match:

@           Match case-insensitively, for example a@ matches A

*           Match any label with this prefix, for example foo* matches foobar

%           Match initial underscore-separated words
The behavior of % is a bit complicated. 
It tries to match words rather than the entire label. 
For example foo_bar% matches foo_bar_baz but not foo_barbaz. 
If combined with *, prefix matching applies to each word separately, 
for example foo_bar%* matches foo1_bar2_baz but not foo1_br2_baz.

 | (OR) to match any of those items, and you can put ! (NOT) at the start of a non-star group to match any label that doesn't match any of the alternatives. 
 A quantifier, if any, goes at the end of the group; it means some number of matches for the group as a whole (that is, some number of labels matching or not matching any of the alternatives).

Here's an annotated example of lquery:

Top.*{0,2}.sport*@.!football|tennis{1,}.Russ*|Spain
a.  b.     c.      d.                   e.
This query will match any label path that:

begins with the label Top

and next has zero to two labels before

a label beginning with the case-insensitive prefix sport

then has one or more labels, none of which match football nor tennis

and then ends with a label beginning with Russ or exactly matching Spain.

ltxtquery represents a full-text-search-like pattern for matching ltree values. An ltxtquery value contains words, possibly with the modifiers @, *, % at the end; the modifiers have the same meanings as in lquery. Words can be combined with & (AND), | (OR), ! (NOT), and parentheses. The key difference from lquery is that ltxtquery matches words without regard to their position in the label path.

Here's an example ltxtquery:

Europe & Russia*@ & !Transportation
This will match paths that contain the label Europe and any label beginning with Russia (case-insensitive), but not paths containing the label Transportation. The location of these words within the path is not important. Also, when % is used, the word can be matched to any underscore-separated word within a label, regardless of position.

Note: ltxtquery allows whitespace between symbols, but ltree and lquery do not.




*/
