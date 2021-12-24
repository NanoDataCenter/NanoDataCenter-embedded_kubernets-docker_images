package eto_support

import (
    
    "fmt"
    "time"
    "github.com/GiterLab/urllib"
    //"encoding/json"
    "lacima.com/redis_support/redis_handlers"
)

/*
self.eto_sources = eto_sources
        self.cimis_data = access_data
        self.app_key = "appKey=" + self.cimis_data["access_key"]
        self.cimis_url = self.cimis_data["url"]
        self.latitude = self.cimis_data["latitude"]
        self.longitude = self.cimis_data["longitude"]
        self.priority = access_data["priority"]
*/


type CIMIS_ETO_SPATIAL_TYPE struct {
    Apk_key        string 
    Cimis_url      string
    Station        string
    Redis_string   string
    Eto_data       redis_handlers.Redis_Hash_Struct
    Rain_data      redis_handlers.Redis_Hash_Struct
    //cimis_data ???? 
    
    
}


func Create_CIMIS_ETO_SPATIAL( input CIMIS_ETO_SPATIAL_TYPE   )CIMIS_ETO_SPATIAL_TYPE{
  var return_value CIMIS_ETO_SPATIAL_TYPE   
  return_value.Redis_string  = "CIMIS:"+ input.Station
    
    
  return return_value
}

func (r CIMIS_ETO_SPATIAL_TYPE)Compute_previous_day()float64{
    
   fmt.Println("compute eto")
   
   if r.Eto_data.HExists(r.Redis_string ) == true {
      fmt.Println("***************** cimis spacial eto returning")
   }
   
   current_day     := time.Now()
   previous_day    := current_day.Add(-24*time.Hour)
   year,month,day  := previous_day.Date()
   date            := fmt.Sprintf("%04d-%02d-%02d",year,month,day)
   fmt.Println("date",date)
       
   //url :=  fmt.Sprintf(`%s+"?"%s"&targets="%d+ "&startDate=" + %s + "&endDate=" + %s`,r.cimis_url,r.app_key,r.station,date,date)
   //url :== r.cimis_url + "?" + r.app_key + "&targets=" + 
   //         str(r.station) + "&startDate=" + date + "&endDate=" + date
   url := ""
   req := urllib.Get(url)
   byt, err := req.Bytes()
   
   fmt.Println(err)
   fmt.Println(byt)
   
   //lat_long = "lat=" + str(self.latitude) + ",lng=" + str(self.longitude)
   //url := r.cimis_url + "?" + r.app_key + "&targets=" + lat_long + "&startDate=" + date + "&endDate=" + date


   //if err != nil {
   //     panic(err)
   //	}
	
    //var dat map[string][string]interface{}
    //if err := json.Unmarshal(byt, &dat); err != nil {
    //    panic(err)
    //}
    //fmt.Println(dat)
 /*   
    `
               temp = float(data["Data"]["Providers"][0]
                     ["Records"][0]['DayAsceEto']["Value"])
        date_string = str(datetime.datetime.now())
        self.eto_sources.hset("CIMIS_SAT:"+str(self.longitude), 
           { "eto":temp,"priority":self.priority,"status":"OK","time": date_string}
    
    `
*/    
   return 0.0
}

/*
`
`
`


    def  compute_previous_day(self):
        if self.eto_sources.hget("CIMIS_SAT:"+str(self.longitude)) != None:
            print("*********************","am returning cimis_spatial")
            return
         
        ts=time.time() - 1 * ONE_DAY  # time is in seconds for desired day
       
        date = datetime.datetime.fromtimestamp(ts).strftime('%Y-%m-%d')
        lat_long = "lat=" + str(self.latitude) + ",lng=" + str(self.longitude)
        url = self.cimis_url + "?" + self.app_key + "&targets=" + \
            lat_long + "&startDate=" + date + "&endDate=" + date

        req = urllib.request.Request(url)
        response = urllib.request.urlopen(req)
        temp = response.read()
       
        data = json.loads(temp.decode())
       
        temp = float(data["Data"]["Providers"][0]
                     ["Records"][0]['DayAsceEto']["Value"])
        date_string = str(datetime.datetime.now())
        self.eto_sources.hset("CIMIS_SAT:"+str(self.longitude), 
           { "eto":temp,"priority":self.priority,"status":"OK","time": date_string})


``
*/
