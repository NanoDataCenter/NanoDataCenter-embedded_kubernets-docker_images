package influxdb_handler

import (
    //"context"
    //"fmt"
   // "time"

    //"github.com/influxdata/influxdb-client-go/v2"
    "github.com/influxdata/influxdb-client-go/v2/domain"
)




/*
**  Influx bucket management  called by site controller of master node
**
**
**
*/

func Create_necessary_buckets( bucket_map map[string]int){
    bucketsAPI := client.BucketsAPI()
    	
	org, err := client.OrganizationsAPI().FindOrganizationByName(ctx, org_name)
    if err != nil{
        panic("constructing org")
    }
    for bucket_name,time_duration := range bucket_map {
        bucket_id ,err := bucketsAPI.FindBucketByName(ctx, bucket_name) 
        if err != nil {
           _,err := bucketsAPI.CreateBucketWithName(ctx, org,bucket_name,domain.RetentionRule{EverySeconds:time_duration})
           if err != nil {
               panic("cannot create bucket")
           }
        }else{
            
            (*bucket_id).RetentionRules = make([]domain.RetentionRule,0)
            (*bucket_id).RetentionRules = append((*bucket_id).RetentionRules,domain.RetentionRule{EverySeconds:time_duration})
          _,err :=  bucketsAPI.UpdateBucket(ctx ,bucket_id)
          if err != nil {
              panic("cannot update time ")
          }
        }
        
    }
    
}


func Prune_obsolete_buckets( bucket_map  map[string]int){
    bucketsAPI := client.BucketsAPI()
    bucket_list , err := bucketsAPI.GetBuckets(ctx)
    if err != nil {
        panic("cannot get bucket list")
    }
    for _,test_bucket := range (*bucket_list){
        name := test_bucket.Name
        if _,ok := bucket_map[name];ok==false{
           err := bucketsAPI.DeleteBucket(ctx , &test_bucket)
           if err != nil {
               panic("cannot delete bucket")
           }
        }
    }
}







