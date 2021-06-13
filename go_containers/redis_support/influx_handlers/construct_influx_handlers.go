package influxdb_handler


import (
    "context"
    "fmt"
    "time"

    "github.com/influxdata/influxdb-client-go/v2"
)

var client     influxdb2.Client
var ctx        context.Context
var org_name   string

var valid_buckets map[string]string  // used as a set

//const valid_buckets const []string{"one_hour","one_day","one_week","one_month","three_month","six_month","one_year"}



type Influxdb_handle struct {
   
   org              string
   bucket           string       // handles retention policy similar to db
   measurement      string
   tags             map[string]string     
   fields           map[string]float64
   
}

/*
 * This function is called by the main functions of the application
 * 
 * 
 * 
 * 
 * 
 */


func Construct_influxdb_client( host_port,  token, org_input string, buckets_used []string){

    // Create a new client using an InfluxDB server base URL and an authentication token
    //client   = influxdb2.NewClient("http://localhost:8086", "my-token")
    
    client     = influxdb2.NewClientWithOptions(host_port,token,influxdb2.DefaultOptions().SetPrecision(time.Second))

    ctx        = context.Background()
    org_name   = org_input
    valid_buckets := make(map[string]string)
    for _,bucket_name := range buckets_used {
        valid_buckets[bucket_name] = "" // used as a set
    
    }
}


func Construct_influxdb_structure( bucket,measurement string, tags,fields []string )Influxdb_handle{
 
    var return_value              Influxdb_handle
   
    
    return_value.bucket          = bucket
    return_value.measurement     = measurement
    return_value.tags            = make(map[string]string)
    return_value.fields          = make(map[string]float64)
    for _,tag := range tags {
        return_value.tags[tag] = ""
    }
    for _,field := range fields {
        return_value.fields[field] = 0.0
    }
 
    return_value.validate_bucket(bucket)
    return return_value
    
}



func ( v Influxdb_handle )Write( tags map[string]string, fields map[string]float64 ){
    v.validate_tags(&tags)
    v.validate_fields(&fields)
    input_data := make(map[string]interface{})  // this is a hack
    for key,data := range fields {
       input_data[key] = data   
        
    }
    

    
   
    p := influxdb2.NewPoint(v.measurement,tags,input_data,time.Now()   ) 
    write_api :=client.WriteAPIBlocking(org_name,v.bucket) 
    write_api.WritePoint(ctx, p)
    
}  
   
   
func (v Influxdb_handle)internal_queury (query_string string ) {
     query_api := client.QueryAPI(org_name)
     result, err := query_api.Query(ctx,query_string)
	 if err != nil {
        panic(err)
     }
     // Iterate over query response
     for result.Next() { // figure out what to do
          fmt.Printf("value: %v\n", result.Record().Value())
     }
}

func (v Influxdb_handle)validate_bucket( bucket string ){
   
    if _,ok := valid_buckets[bucket]; ok == false {
        panic("bad bucket "+bucket)
    }
        
    
}
    
func (v Influxdb_handle)validate_fields( fields *map[string]float64 ){

    if len((*fields)) != len(v.fields) {
        panic("fields are not of same length")
    }
    for test_field,_ := range (*fields) {
       if _,ok := v.fields[test_field];ok == false{
            panic("field "+test_field )
        }
    }
}

func (v Influxdb_handle)validate_tags( tags *map[string]string  ){

    if len((*tags)) != len(v.tags) {
        panic("tags are not of same length")
    }
    for test_tag,_ := range (*tags) {
       if _,ok := v.tags[test_tag];ok == false{
            panic("bad tag "+test_tag )
        }
    }
}








