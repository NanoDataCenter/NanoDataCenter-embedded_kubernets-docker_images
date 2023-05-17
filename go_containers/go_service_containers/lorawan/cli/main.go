package main

import (
	//"fmt"
    "lacima.com/go_service_containers/lorawan/loran_storage_interface"
	
)

func main() {
    
    url_base := "https://nam1.cloud.thethings.network/api/v3/as/applications/"
    //app_name :="lacima-ranch-test-app-1"
    app_name := "seeedec"
    url_after    := "/packages/storage/uplink_message?"
    limit  := "200"
    after := "2020-08-20T00:00:00Z"
    password := "NNSXS.5N2DRLTP3QD4SNMBXNWXZ6V3SMPEGXSW6JOT25I.7VUBLSUKWWEK4KAQUY3SP66Z6YHLQQVMRIKTWL2I7GH4GNRHETIA"
    loran_server_storage_interface.Get_data( url_base, app_name,url_after ,password,limit, after)
     //fmt.Println("err",err)
     //fmt.Println("data",data)
    
}	
